package models

import (
	"log"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
)

const _Ip = "db_ip"
const _Port = "db_port"
const _Db = "db_db"
const _User = "db_user"
const _Password = "db_passworld"

//ConnectString parameter passing to Db.Open. Use sql.Open(db.DBName, db.ConnectString) to open a Connection
var ConnectString string

//DBName parameter passing to Db.Open. Use sql.Open(db.DBName, db.ConnectString) to open a Connection
var DBName = "mysql"

//init database.Will auto create table needed if table not exist.
func init() {

	ip := beego.AppConfig.String(_Ip)
	port := beego.AppConfig.String(_Port)
	db := beego.AppConfig.String(_Db)
	user := beego.AppConfig.String(_User)
	password := beego.AppConfig.String(_Password)

	log.Printf("ip:[%s],port:[%s],user:[%s],password:[%s],db:[%s]", ip, port, user, password, db)

	config := mysql.NewConfig()
	log.Printf("Default DSN:%s", config.FormatDSN())

	if ip != "" || port != "" {
		addr := fmt.Sprintf("%s:%s", ip, port)
		config.Addr = addr
	}
	if user != "" {
		config.User = user
	} else {
		log.Panicf("DB User Not Set. Please set \"db_user=xxxx\" in conf/app.conf")
	}
	if password != "" {
		config.Passwd = password
	} else {
		log.Panicf("DB Password Not Set. Please set \"db_passworld=xxxx\" in conf/app.conf")
	}
	config.Net = "tcp"
	if db != "" {
		config.DBName = db
	} else {
		log.Panicf("DB Name Not Set. Please set \"db_db=xxxx\" in conf/app.conf")
	}

	ConnectString = config.FormatDSN()

	log.Printf("DSN set as:%s", ConnectString)
}
