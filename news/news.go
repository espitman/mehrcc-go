package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type News struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Lead  string `json:"lead"`
	Time  Time   `json:"time"`
	Image string `json:"image"`
}

type Time struct {
	DayWNum   string `json:"dayWNum"`
	DayWName  string `json:"dayWName"`
	MonthName string `json:"monthName"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Hour      string `json:"hour"`
	Min       string `json:"min"`
	Sec       string `json:"sec"`
}

var AllNews []News

func GetNews(wg *sync.WaitGroup) {
	cache, e := redis.Get("news")
	if e != nil {
		api := "http://api.mehrcc.com/wp-json/v3/news?count=3"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				AllNews = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []News
			redis.Set("news", body)
			json.Unmarshal(body, &data)
			AllNews = data
			wg.Done()
		}
	} else {
		var data []News
		json.Unmarshal([]byte(cache), &data)
		AllNews = data
		wg.Done()
	}
}
