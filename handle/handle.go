package handle

import (
	"Haoran/haoran"
	"strings"

	//"log"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/liuzl/gocc"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Readme(c *gin.Context) {
	res, err := haoran.Readme()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, res)
}

func Chapters(c *gin.Context) {
	chaps := haoran.Chapters()

	//t2s, err := gocc.New("t2s")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for k, _ := range chaps {
	//	chaps[k], err = t2s.Convert(chaps[k])
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	c.JSON(200, gin.H{
		"chapters": chaps,
	})
}

func ChaptersRand(c *gin.Context) {
	chaps := haoran.ChapterRand()
	paragraphs := make([]string, 0)
	// content = append(content, chaps.Chapter)
	for _, v := range chaps.Paragraphs {
		paragraphs = append(paragraphs, v.Content)
	}
	content := strings.Join(paragraphs, "\n\n")

	//t2s, err := gocc.New("t2s")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for k, _ := range chaps {
	//	chaps[k], err = t2s.Convert(chaps[k])
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	c.JSON(200, gin.H{
		"chapter": chaps.Chapter,
		"content": content,
	})
}

func ChapterDetial(c *gin.Context) {
	cID := c.Query("chapterID")
	id, err := strconv.Atoi(cID)
	if err != nil {
		c.Status(500)
		return
	}

	res := haoran.ChapterDetail(id)
	c.JSON(200, *res)
}

func Paragraphs(c *gin.Context) {
	res := haoran.Paragraphs()

	c.JSON(200, *res)
}

func ParagraphsDetial(c *gin.Context) {
	cID := c.Query("chapterID")
	cId, err := strconv.Atoi(cID)
	if err != nil {
		c.Status(500)
		return
	}

	pID := c.Query("paragraphID")
	pId, err := strconv.Atoi(pID)
	if err != nil {
		c.Status(500)
		return
	}
	res := haoran.ParagraphsDetial(cId, pId)
	c.JSON(200, *res)
}

