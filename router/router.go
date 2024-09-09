package router

import (
	"Haoran/handle"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/ping", handle.Ping)
	r.GET("/readme", handle.Readme)
	r.GET("/chapters", handle.Chapters)
	r.GET("/chapter", handle.ChaptersRand)
	r.GET("/chapters/detial", handle.ChapterDetial)
	r.GET("/paragraphs", handle.Paragraphs)
	r.GET("/paragraphs/detial", handle.ParagraphsDetial)
}
