package models

import (
	"github.com/astaxie/beego/cache"
	"html/template"
)

const MaxCacheTime = 3600

var PageCache cache.Cache

func PageCacheGet(title string) template.HTML {
	return PageCache.Get(title).(template.HTML)
}

func isCacheExist(title string) (res bool) {
	return PageCache.IsExist(title)
}

func pageCacheAdd(title string, content template.HTML) {
	PageCache.Put(title, content, MaxCacheTime)
}

func pageCacheRemove(title string) {
	PageCache.Delete(title)
}
