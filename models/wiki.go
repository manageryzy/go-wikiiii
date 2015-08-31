package models
import (
	"regexp"
	"strings"
)

const MaxIncludeLayers  = 5

//获得一个页面
func PageGet(title string)(res string){
	if isCacheExist(title) {
		res = PageCacheGet(title)
		return
	}else {
		res =  PageRender(title)
		pageCacheAdd(title,res)
		return
	}
}


//强制渲染一个页面
func PageRender(title string)(res string){
	page,exist :=  pageGetSQL(title)

	if !exist{
		res = ""
		return
	}

	page = pageRenderInclude(page,0)
	page = pageRenderLinks(page)
	res = page

	return
}

func pageGetSQL(title string)(res string,exist bool)  {
	p := Page{Title: title}
	err := O.Read(&p)

	if err != nil {
		exist = false
		res = ""
		return
	}

	exist = true
	res = p.Page

	return
}

func pageRenderInclude(content string,layers int)(res string)  {
	if layers > MaxIncludeLayers{
		res = "<pre>Error: Too much layers included!!!</pre>"
		return
	}

	re := regexp.MustCompile("{{.*}}")
	includes := re.FindAllString(content,-1)

	for _,include := range includes{
		title := strings.Trim(include,"{} ")

		subPage ,exist := pageGetSQL(title)
		if exist {
			r := strings.NewReplacer(include, pageRenderInclude(subPage,layers + 1))
			content = r.Replace(content)
		}else {
			r := strings.NewReplacer(include, "<pre>Error: Include page not found!</pre>")
			content = r.Replace(content)
		}
	}

	res = content
	return
}

func pageRenderLinks(content string) (res string) {
	res = content
	return
}