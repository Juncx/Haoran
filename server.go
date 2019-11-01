package main

import (
	"Haoran/handle"

	"github.com/gin-gonic/gin"
)

func init() {
}

func main() {
	r := gin.Default()
	handle.Router(r)

	//go r.RunTLS(":443", "../ssl/server.crt", "../ssl/server.key")
	r.Run(":8080")
}
