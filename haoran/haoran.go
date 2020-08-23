package haoran

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/liuzl/gocc"
)

const READMEPATH = "haoran/lunyu/README.md"
const CONTENTPATH = "haoran/lunyu/lunyu.json"

//const CONTENTPATH = "haoran/lunyu/lunyu_new.json"
var LUNYUALL []Chapter

func (m *LunYu) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

/*
 *func cache() {
 *    // redis client
 *    rdb := redis.NewClient(&redis.Options{
 *        Addr:     "localhost:6379", // use default Addr
 *        Password: "",               // no password set
 *        DB:       0,                // use default DB
 *    })
 *    defer rdb.Close()
 *
 *    // set redis zset
 *    for k, v := range LUNYUALL {
 *        mes := redis.Z{
 *            Score:  float64(k),
 *            Member: &v,
 *        }
 *
 *        err := rdb.ZAdd("lunyu_all", mes).Err()
 *        if err != nil {
 *            panic(err)
 *        }
 *    }
 *
 *}
 */

func init() {
	content, err := readFile(CONTENTPATH)
	if err != nil {
		log.Fatal(err)
	}

	// parse LunYu raw data
	var LunYuRaw []LunYu
	err = json.Unmarshal([]byte(content), &LunYuRaw)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range LunYuRaw {
		c := Chapter{
			CId:     k,
			Chapter: v.Chapter,
		}

		for kk, vv := range v.Paragraphs {
			p := ParagraphDetial{
				CId:     k,
				PId:     kk,
				Chapter: v.Chapter,
				Content: vv,
			}
			c.Paragraphs = append(c.Paragraphs, p)
		}

		LUNYUALL = append(LUNYUALL, c)
	}

	// init LunYu cache
	//cache()
}

func Readme() (string, error) {
	return readFile(READMEPATH)
}

func Chapters() []string {
	var res []string
	for _, v := range LUNYUALL {
		res = append(res, v.Chapter)
	}

	return res
}

func ChapterDetail(chapterID int) *Chapter {
	return &LUNYUALL[chapterID]
}

func Paragraphs() *ParagraphDetial {
	rand.Seed(time.Now().UnixNano())
	cId := rand.Intn(len(LUNYUALL))
	pId := rand.Intn(len(LUNYUALL[cId].Paragraphs))

	return &LUNYUALL[cId].Paragraphs[pId]
}

func ParagraphsDetial(chapterID int, paragraphID int) *ParagraphDetial {
	res := LUNYUALL[chapterID].Paragraphs[paragraphID]

	t2s, err := gocc.New("t2s")
	if err != nil {
		log.Fatal(err)
	}

	res.Chapter, err = t2s.Convert(res.Chapter)
	if err != nil {
		log.Fatal(err)
	}

	res.Content, err = t2s.Convert(res.Content)
	if err != nil {
		log.Fatal(err)
	}

	return &res
}

func readFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)

	return string(content), err
}
