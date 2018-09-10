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

	// 静态文件路径
	StaticFilePath string
	// 页面模板文件路径
	TemplateFilePath string

	// MySQL dataSourceName
	DataSourceName string
	// MySQL 最大连接数
	DBMaxOpenConn int
	// MySQL 最大闲置连接数
	DBMaxIdleConn int
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

	iniflags.Parse()
}
