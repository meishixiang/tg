// mysql db drives
package drivers

import (
	"boxDemo/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDb *gorm.DB // db pool instance
var MysqlDbErr error // db err instance

func init() {
	// get db config
	dbConfig := config.GetDbConfig()
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
	)

	// connect and open db connection
	MysqlDb, MysqlDbErr = gorm.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		panic("database data source name error: " + MysqlDbErr.Error())
	}
}
