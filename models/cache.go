package models
import "time"

const cacheExpressTime  = 3600

type pageCacheItem struct {
	title string
	content string
	update int64
}

var pageCache = map[string]pageCacheItem{}

func PageCacheGet(title string)(res string){
	cache,e := pageCache[title]
	if !e {
		res = ""
		return
	}

	res = cache.content
	return
}

func isCacheExist(title string)(res bool){
	cache,e := pageCache[title]
	if !e {
		res = false
		return
	}

	if time.Now().Unix() -  cache.update > cacheExpressTime{
		delete(pageCache,title)
		res = false
		return
	}

	res = true
	return
}

func pageCacheAdd(title string,content string){
	o := pageCacheItem{}
	o.content = content
	o.update = time.Now().Unix() - 3590
	o.title = title

	pageCache[title] = o

}

func pageCacheGC()  {
	for title,cache := range pageCache{
		if time.Now().Unix() -  cache.update > cacheExpressTime{
			delete(pageCache,title)
		}
	}
}