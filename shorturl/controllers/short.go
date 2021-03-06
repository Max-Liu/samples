package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/beego/samples/shorturl/models"
)

var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

type ShortController struct {
	beego.Controller
}

func (this *ShortController) Post() {
	var result ShortResult
	longurl := this.Input().Get("longurl")
	beego.Info(longurl)
	result.UrlLong = longurl
	urlmd5 := models.GetMD5(longurl)
	beego.Info(urlmd5)
	if urlcache.IsExist(urlmd5) {
		result.UrlShort = string(urlcache.Get(urlmd5).([]uint8))
	} else {
		result.UrlShort = models.Generate()
		err := urlcache.Put(urlmd5, result.UrlShort, 0)
		if err != nil {
			beego.Info(err)
		}
		err = urlcache.Put(result.UrlShort, longurl, 0)
		if err != nil {
			beego.Info(err)
		}
	}
	this.Data["json"] = result
	this.ServeJson()
}
