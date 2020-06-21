package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type Term struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Article struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Title   string `json:"title"`
	Lead    string `json:"lead"`
	Author  string `json:"author"`
	File    string `json:"file"`
	Image   string `json:"image"`
	Terms   []Term `json:"terms"`
	TermIds []int  `json:"termIds"`
}

var Articles []Article

func GetArticles(wg *sync.WaitGroup) {
	cache, e := redis.Get("articles")
	if e != nil {
		api := "https://api.mehrcc.com/wp-json/v3/articles?count=3"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				Articles = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []Article
			redis.Set("articles", body)
			json.Unmarshal(body, &data)
			Articles = data
			wg.Done()
		}
	} else {
		var data []Article
		json.Unmarshal([]byte(cache), &data)
		Articles = data
		wg.Done()
	}
}
