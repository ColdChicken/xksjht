package handle

import (
	"be/server"
)

/*
* InitHandle负责做Handle到实际URL的映射工作,
* handle包下的handle如果要被实际使用,则都需要在此进行注册
 */
func InitHandle(r *server.WWWMux) {
	// api相关的接口
	initAPIMapping(r)
}

func initAPIMapping(r *server.WWWMux) {
	// 获取新闻
	r.RegistURLMapping("/v1/api/news/list", "GET", ajaxListNews)
}
