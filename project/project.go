package project

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type Project struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Target string `json:"target"`
	Year   string `json:"year"`
	Image  string `json:"image"`
}

var Projects []Project

func GetProjects(wg *sync.WaitGroup) {
	cache, e := redis.Get("projects")
	if e != nil {
		api := "https://api.mehrcc.com/wp-json/v3/projects?count=5"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				Projects = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []Project
			redis.Set("projects", body)
			json.Unmarshal(body, &data)
			Projects = data
			wg.Done()
		}
	} else {
		var data []Project
		json.Unmarshal([]byte(cache), &data)
		Projects = data
		wg.Done()
	}
}
