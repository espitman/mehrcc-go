package slider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type Slider struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Link     bool   `json:"link"`
	LinkType bool   `json:"linkType"`
}

var Sliders []Slider

func GetSliders(wg *sync.WaitGroup) {
	cache, e := redis.Get("sliders")
	if e != nil {
		api := "http://api.mehrcc.com/wp-json/v3/sliders"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				Sliders = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []Slider
			redis.Set("sliders", body)
			json.Unmarshal(body, &data)
			Sliders = data
			wg.Done()
		}
	} else {
		var data []Slider
		json.Unmarshal([]byte(cache), &data)
		Sliders = data
		wg.Done()
	}
}
