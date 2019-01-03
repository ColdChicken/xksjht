package handle

import (
	"be/options"
	"be/server"
	"net/http"
)

/*
* InitHandle负责做Handle到实际URL的映射工作,
* handle包下的handle如果要被实际使用,则都需要在此进行注册
 */
func InitHandle(r *server.WWWMux) {
	// 初始化静态文件路径
	initStaticFileMapping(r)
	// 初始化管理控制台相关页面
	initAdminPortalMapping(r)
	// 初始化管理控制台ajax
	initAjaxMapping(r)
	// api相关的接口
	initAPIMapping(r)

	// 代理
	r.SetProxy("/v1/tp/{catalog}/{action}", options.Options.SweetFishAddress)
	r.SetProxy("/v1/tp/{catalog}/{action}/{detail}", options.Options.SweetFishAddress)
}

func initStaticFileMapping(r *server.WWWMux) {
	fs := http.FileServer(http.Dir(options.Options.StaticFilePath))
	r.GetRouter().PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func initAdminPortalMapping(r *server.WWWMux) {
	// 后台
	r.RegistURLMapping("/ht", "GET", showIndexHtml)
	r.RegistURLMapping("/ht/ologin.html", "GET", showLoginHtml)

	// 前台
	r.RegistURLMapping("/", "GET", showZiXunHtml)
	r.RegistURLMapping("/zx", "GET", showZiXunHtml)
	r.RegistURLMapping("/bj", "GET", showBiJiHtml)
	r.RegistURLMapping("/gj", "GET", showGongJuHtml)
	r.RegistURLMapping("/gy", "GET", showGuanYuHtml)
	r.RegistURLMapping("/article/{id}", "GET", showArticleHtml)
	r.RegistURLMapping("/robots.txt", "GET", showRobotsHtml)

	// 默认路由
	r.GetRouter().NotFoundHandler = http.HandlerFunc(server.AccessLogHandler(showZiXunHtml))
}

func initAjaxMapping(r *server.WWWMux) {
	// 注销
	r.RegistURLMapping("/v1/ajax/auth/logout", "GET", ajaxLogout)
	// 用户认证密码并生成token
	r.RegistURLMapping("/v1/ajax/auth/token", "POST", ajaxGenTokenByUMAndPassword)
	// 获取用户信息
	r.RegistURLMapping("/v1/ajax/auth/info", "GET", ajaxGetUserInfo)
	// 列出文章
	r.RegistURLMapping("/v1/ajax/article/listbyfilter", "POST", ajaxListArticlesByFilter)
	// 创建文章
	r.RegistURLMapping("/v1/ajax/article/create", "POST", ajaxCreateArticle)
	// 更新文章
	r.RegistURLMapping("/v1/ajax/article/update", "POST", ajaxUpdateArticle)
	// 删除文章
	r.RegistURLMapping("/v1/ajax/article/delete", "POST", ajaxDeleteArticle)
	// 列出被删除的文章
	r.RegistURLMapping("/v1/ajax/article/listdeleted", "POST", ajaxListDeletedArticles)
	// 上传图片
	r.RegistURLMapping("/v1/ajax/file/pic/upload", "POST", ajaxUploadPic)
	// 下载图片
	r.RegistURLMapping("/v1/ajax/file/pic/download/{location}", "GET", ajaxDownloadPic)
	// 列出所有图片信息
	r.RegistURLMapping("/v1/ajax/file/pic/list", "POST", ajaxListPics)
	// 删除图片
	r.RegistURLMapping("/v1/ajax/file/pic/delete", "POST", ajaxDeletePic)
}

func initAPIMapping(r *server.WWWMux) {
	// 列出文章
	r.RegistURLMapping("/v1/api/article/listbyfilter", "POST", apiListArticles)
	r.RegistURLMapping("/v1/api/article/listbyfilter", "GET", apiListArticles)
	// 获取文章详情
	r.RegistURLMapping("/v1/api/article/getbyid", "POST", apiGetArticleById)
	r.RegistURLMapping("/v1/api/article/getbyid", "GET", apiGetArticleById)
	// 下载图片
	r.RegistURLMapping("/v1/api/file/pic/download/{location}", "GET", apiDownloadPic)
}
