package models

import (
	"github.com/astaxie/beego/cache"
)

const MaxCacheTime = 3600

var PageCache cache.Cache

func PageCacheGet(title string) string {
	return PageCache.Get(title).(string)
}

func isCacheExist(title string) (res bool) {
	return PageCache.IsExist(title)
}

func pageCacheAdd(title string, content string) {
	PageCache.Put(title, content, MaxCacheTime)
}

func pageCacheRemove(title string) {
	PageCache.Delete(title)
}
