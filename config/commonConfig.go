package config

func GetServerConfig() {
	var serverConfig = make(map[string]string)
	serverConfig["HOST"] = "0.0.0.0"                     //监听地址
	serverConfig["PORT"] = "8080"                        //监听端口
	serverConfig["VIEWS_PATTERN"] = "easy-gin/views/*/*" //视图模板路径pattern
	serverConfig["ENV"] = "debug"                        //环境模式 release/debug/test
	return
}

func GetDbConfig() map[string]string {
	dbConfig := make(map[string]string)

	dbConfig["DB_HOST"] = "127.0.0.1" //主机
	dbConfig["DB_PORT"] = "3306"      //端口
	dbConfig["DB_NAME"] = "tg_common"    //数据库
	dbConfig["DB_USER"] = "root"      //用户名
	dbConfig["DB_PWD"] = "12345678"           //密码
	dbConfig["DB_CHARSET"] = "utf8"

	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return dbConfig
}
