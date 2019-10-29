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

	r.Run()
}
