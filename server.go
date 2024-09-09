package main

import (
	"Haoran/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)

	//go r.RunTLS(":443", "../ssl/server.crt", "../ssl/server.key")
	r.Run(":8088")
}
