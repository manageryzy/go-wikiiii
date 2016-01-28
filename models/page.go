package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"regexp"
	"strings"
)

const MaxIncludeLayers = 5

const markdownExtensions = 0 |
blackfriday.EXTENSION_NO_INTRA_EMPHASIS         			 |// ignore emphasis markers inside words
blackfriday.EXTENSION_TABLES                                 |// render tables
blackfriday.EXTENSION_FENCED_CODE                            |// render fenced code blocks
blackfriday.EXTENSION_AUTOLINK                               |// detect embedded URLs that are not explicitly marked
blackfriday.EXTENSION_STRIKETHROUGH                          |// strikethrough text using ~~test~~
blackfriday.EXTENSION_LAX_HTML_BLOCKS                        |// loosen up HTML block parsing rules
blackfriday.EXTENSION_SPACE_HEADERS                          |// be strict about prefix header rules
blackfriday.EXTENSION_HARD_LINE_BREAK                        |// translate newlines into line breaks
blackfriday.EXTENSION_TAB_SIZE_EIGHT                         |// expand tabs to eight spaces instead of four
blackfriday.EXTENSION_FOOTNOTES                              |// Pandoc-style footnotes
blackfriday.EXTENSION_HEADER_IDS                             |// specify header IDs  with {#id}
blackfriday.EXTENSION_TITLEBLOCK                             |// Titleblock ala pandoc
blackfriday.EXTENSION_AUTO_HEADER_IDS                        |// Create the header ID from the text
blackfriday.EXTENSION_BACKSLASH_LINE_BREAK                   |// translate trailing backslashes into line breaks
blackfriday.EXTENSION_DEFINITION_LISTS                       // render definition lists

const htmlFlags = 0 |
	blackfriday.HTML_USE_XHTML |
	blackfriday.HTML_USE_SMARTYPANTS |
	blackfriday.HTML_SMARTYPANTS_FRACTIONS |
	blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

//获得一个页面
func PageGet(title string) (res template.HTML) {
	if isCacheExist(title) {
		res = PageCacheGet(title)
		return
	} else {
		res = template.HTML(PageRender(title))
		pageCacheAdd(title, res)
		return
	}
}

//强制渲染一个页面
func PageRender(title string) (res string) {
	page, exist, safe := PageGetSQL(title)

	if !exist {
		res = "本页面没有内容，点击<a href=\"/edit/" + title + "\">此处</a>编辑"
		return
	}

	page = PageRenderString(page, safe)
	res = pageRefreshCategory(page, title)
	return
}

func PageRenderString(page string, safe bool) (res string) {
	page = pageRenderInclude(page, 0)
	page = pageRenderLinks(page)
	page = pageRenderMarkdown(page)

	if safe {
		page = string(bluemonday.UGCPolicy().SanitizeBytes([]byte(page)))
	}

	res = page

	return
}

func PageGetCategory(title string) []orm.Params {
	var maps []orm.Params
	O.QueryTable("categories").Filter("title", title).Values(&maps)
	return maps
}

func CategoryGetPages(cat string) []orm.Params {
	var maps []orm.Params
	O.QueryTable("categories").Filter("category", cat).Values(&maps)
	return maps
}

func PageGetSQL(title string) (res string, exist bool, safe bool) {
	p := Page{Title: title}
	err := O.Read(&p)

	if err != nil {
		exist = false
		res = ""
		return
	}

	if p.Safe == 0 {
		safe = false
	} else {
		safe = true
	}

	exist = true
	res = p.Page

	return
}

func pageRenderInclude(content string, layers int) (res string) {
	if layers > MaxIncludeLayers {
		res = "<pre>Error: Too much layers included!!!</pre>"
		return
	}

	re := regexp.MustCompile("{{.*}}")
	includes := re.FindAllString(content, -1)

	for _, include := range includes {
		title := strings.Trim(include, "{} ")

		subPage, exist, _ := PageGetSQL(title)
		if exist {
			r := strings.NewReplacer(include, pageRenderInclude(subPage, layers+1))
			content = r.Replace(content)
		} else {
			r := strings.NewReplacer(include, "<pre>Error: Include page not found!</pre>")
			content = r.Replace(content)
		}
	}

	res = content
	return
}

func pageRenderMarkdown(content string) (res string) {
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")
	unsafe := blackfriday.MarkdownOptions([]byte(content), renderer, blackfriday.Options{
		Extensions: markdownExtensions})
	res = string(unsafe)
	return
}

func pageRenderLinks(content string) (res string) {
	re := regexp.MustCompile("\\[\\[.*\\]\\]")
	links := re.FindAllString(content, -1)

	for _, include := range links {
		title := strings.Trim(include, "[] ")

		r := strings.NewReplacer(include, "<a href=\"/page/"+title+"\" >"+title+"</a>")
		content = r.Replace(content)
	}
	res = content
	return
}

func pageRefreshCategory(content string, title string) (res string) {
	content = pageRenderInclude(content, 0)

	O.QueryTable("categories").Filter("title", title).Delete()

	re := regexp.MustCompile("\\[{.*}\\]")
	cats := re.FindAllString(content, -1)
	for _, include := range cats {
		category := strings.Trim(include, "[`] ")

		cat := Categories{Title: title, Category: category}
		O.Insert(&cat)

		r := strings.NewReplacer(include, "")
		content = r.Replace(content)
	}

	res = content
	return
}

//更新页面
func PageEdit(title string, content string, uid int, safe bool, fileName string, reason string) (res bool) {
	user := User{Uid: uid}
	err := O.Read(&user)
	if err != nil {
		return false
	}

	pageCacheRemove(title)

	err = ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		println(err.Error())
		return false
	}

	pageRefreshCategory(content, title)

	p := Page{Title: title, Page: content, Uid: uid}
	if safe {
		p.Safe = 1
	} else {
		p.Safe = 0
	}
	num, err := O.Update(&p)

	if err != nil {
		return false
	}

	if num == 0 {
		_, err := O.Insert(&p)

		if err != nil {
			return false
		}
	}

	h := History{Title: title, Path: fileName, Uid: uid, Reason: reason, Name: user.Name}
	O.Insert(&h)

	return true
}
