package course

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type Course struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Teacher string `json:"teacher"`
	Date    bool   `json:"date"`
	Time    string `json:"time"`
	Year    string `json:"year"`
	Month   string `json:"month"`
	Day     string `json:"day"`
}

var Courses []Course

func GetCourses(wg *sync.WaitGroup) {
	cache, e := redis.Get("courses")
	if e != nil {
		api := "http://api.mehrcc.com/wp-json/v3/courses"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				Courses = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []Course
			redis.Set("courses", body)
			json.Unmarshal(body, &data)
			Courses = data
			wg.Done()
		}
	} else {
		var data []Course
		json.Unmarshal([]byte(cache), &data)
		Courses = data
		wg.Done()
	}
}
