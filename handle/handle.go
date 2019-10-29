package handle

import (
	"Haoran/haoran"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liuzl/gocc"
)

func Router(r *gin.Engine) {
	r.GET("/ping", ping)
	r.GET("/readme", readme)
	r.GET("/chapters", chapters)
	r.GET("/paragraphs", paragraphs)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func readme(c *gin.Context) {
	res, err := haoran.Readme()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, res)
}

func chapters(c *gin.Context) {
	chaps := haoran.Chapters()

	t2s, err := gocc.New("t2s")
	if err != nil {
		log.Fatal(err)
	}

	for k, _ := range chaps {
		chaps[k], err = t2s.Convert(chaps[k])
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(200, gin.H{
		"chapters": chaps,
	})
}

func paragraphs(c *gin.Context) {
	paras := haoran.Paragraphs()

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(paras))

	t2s, err := gocc.New("t2s")
	if err != nil {
		log.Fatal(err)
	}

	res, err := t2s.Convert(paras[index])
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"index":     index,
		"paragraph": res,
	})
}
