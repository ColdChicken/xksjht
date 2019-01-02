package options

import (
	"flag"

	"github.com/vharitonsky/iniflags"
)

type WWWOptions struct {
	// 默认日志级别
	LogLevel string
	LogFile  string

	// HTTP服务监听地址
	HTTPAddress string
	HTTPPort    uint64

	// 是否使用HTTPS
	EnableTls bool

	// HTTPS服务监听地址
	HTTPSAddr string
	// TLS
	CertFile string
	KeyFile  string

	// 静态文件路径
	StaticFilePath string
	// 页面模板文件路径
	TemplateFilePath string

	// 密码加密盐
	PasswordSalt string
	// Cookie01
	Cookie01 string
	// Cookie02
	Cookie02 string

	// MySQL dataSourceName
	DataSourceName string
	// MySQL 最大连接数
	DBMaxOpenConn int
	// MySQL 最大闲置连接数
	DBMaxIdleConn int

	// 本地图片存放路径
	LocalPicRootPath string

	// 图片的默认对外地址根路径
	PicExternalRootPath string

	// 文章的默认创建人
	DefaultArticleCreater string

	// 文章二维码
	ArticleQRCodeURL string

	// 星空工具id
	XKGJId uint64
	// 关于星空id
	XKGYId uint64
}

var Options WWWOptions

func (o *WWWOptions) InitOptions() {
	flag.StringVar(&o.LogLevel, "log_level", "DEBUG", "Log Level")
	flag.StringVar(&o.LogFile, "log_file", "D:\\logs\\xksjht.log", "Log File")
	flag.StringVar(&o.HTTPAddress, "http_address", "0.0.0.0", "HTTP Address")
	flag.Uint64Var(&o.HTTPPort, "http_port", 8888, "HTTP Port")
	flag.StringVar(&o.DataSourceName, "dsn", "root:rootroot@tcp(127.0.0.1:3306)/xksj?autocommit=0&collation=utf8_general_ci", "MySQL DataSourceName")
	flag.IntVar(&o.DBMaxOpenConn, "max_open_conn", 32, "MySQL Max Open Connections")
	flag.IntVar(&o.DBMaxIdleConn, "max_idle_conn", 16, "MySQL Max Idle Connections")
	flag.StringVar(&o.StaticFilePath, "static_file_path", "", "StaticFilePath")
	flag.StringVar(&o.TemplateFilePath, "template_file_path", "", "TemplateFilePath")
	flag.StringVar(&o.PasswordSalt, "password_salt", "", "PasswordSalt")
	flag.StringVar(&o.Cookie01, "cookie_01", "", "Cookie01")
	flag.StringVar(&o.Cookie02, "cookie_02", "", "Cookie02")
	flag.StringVar(&o.DefaultArticleCreater, "default_article_creater", "CC", "DefaultArticleCreater")
	flag.StringVar(&o.LocalPicRootPath, "local_pic_root_path", "D:\\logs\\pics", "LocalPicRootPath")
	flag.StringVar(&o.PicExternalRootPath, "pic_external_root_path", "http://192.168.1.102:8888/v1/api/file/pic/download", "PicExternalRootPath")
	flag.StringVar(&o.HTTPSAddr, "https_addr", "0.0.0.0:443", "HTTPSAddr")
	flag.StringVar(&o.CertFile, "cert_file", "./key/server.pem", "CertFile")
	flag.StringVar(&o.KeyFile, "key_file", "./key/server.key", "KeyFile")
	flag.BoolVar(&o.EnableTls, "enable_tls", false, "EnableTls")
	flag.StringVar(&o.ArticleQRCodeURL, "article_qrcode_url", "http://192.168.1.102:8888/article", "ArticleQRCodeURL")
	flag.Uint64Var(&o.XKGJId, "xkgj_id", 4, "XKGJ id")
	flag.Uint64Var(&o.XKGYId, "gyxk_id", 4, "XKGYId id")

	iniflags.Parse()
}
