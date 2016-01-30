package models

import (
	"github.com/astaxie/beego/cache"
)

const MaxCacheTime = 3600*1e9

var PageCache cache.Cache

func PageCacheGet(title string) (res string ,exist bool) {
	d := PageCache.Get(title)
	if d == nil{
		res = ""
		exist = false
	}else {
		res = d.(string)
		exist = true
	}
	return
}

func pageCacheAdd(title string, content string) {
	PageCache.Put(title, content, MaxCacheTime)
}

func pageCacheRemove(title string) {
	PageCache.Delete(title)
}
