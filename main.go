package main

import (
	"poker/server/conf"
	"poker/server/gate"
	"poker/server/login"

	"poker/github/dolotech/leaf"

	// "net/http"
	"flag"
	"poker/github/dolotech/lib/db"
	"poker/server/game"
	"poker/server/model"

	"github.com/golang/glog"
)

var Commit = ""
var BUILD_TIME = ""
var VERSION = ""

var createdb bool

func init() {
	flag.StringVar(&conf.Server.WSAddr, "addr", ":8989", "websocket port")
	flag.IntVar(&conf.Server.MaxConnNum, "maxconn", 20000, "Max Conn Num")
	flag.StringVar(&conf.Server.DBUrl, "sql", "root:34652402@tcp(127.0.0.1:3306)/poker?charset=utf8", "mysql addr")
	flag.BoolVar(&createdb, "createdb", true, "initial database")

	flag.Parse()

	db.Init(conf.Server.DBUrl)
	if !createdb {
		createDb()
	}
}

func main() {

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

	// for test client
	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	// err := http.ListenAndServe(":12345", nil)
	// if err != nil {
	// 	glog.Fatalf("ListenAndServe: %v ", err)
	// }
}

func createDb() {
	// 建表,只维护和服务器game里面有关的表
	err := db.C().Sync(model.User{}, model.Room{})
	if err != nil {
		glog.Errorln(err)
	}
}
