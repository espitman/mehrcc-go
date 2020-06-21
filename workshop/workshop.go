package workshop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type Workshop struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Sessions     string `json:"sessions"`
	Prerequisite bool   `json:"prerequisite"`
	Image        string `json:"image"`
	Theory       bool   `json:"theory"`
	Practical    bool   `json:"practical"`
}

var Workshops []Workshop

func GetWorkshops(wg *sync.WaitGroup) {
	cache, e := redis.Get("workshops")
	if e != nil {
		api := "http://api.mehrcc.com/wp-json/v3/workshops?count=8"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				Workshops = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []Workshop
			redis.Set("workshops", body)
			json.Unmarshal(body, &data)
			Workshops = data
			wg.Done()
		}
	} else {
		var data []Workshop
		json.Unmarshal([]byte(cache), &data)
		Workshops = data
		wg.Done()
	}
}
