package main

import (
	"gin_mall/conf"
	"gin_mall/routes"
)

func main() {
	//fmt.Println("Hello World")
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
