package consultant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehrcc/redis"
	"net/http"
	"sync"
)

type Consultant struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Field  string `json:"field"`
	Degree string `json:"degree"`
}

var Consultants []Consultant

func GetConsultants(wg *sync.WaitGroup) {
	cache, e := redis.Get("consultants")
	if e != nil {
		api := "https://api.mehrcc.com/wp-json/v3/consultants?count=15"
		resp, err := http.Get(api)
		if err == nil {
			defer resp.Body.Close()
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				Consultants = nil
				fmt.Println(e)
				wg.Done()
			}
			var data []Consultant
			redis.Set("consultants", body)
			json.Unmarshal(body, &data)
			Consultants = data
			wg.Done()
		}
	} else {
		var data []Consultant
		json.Unmarshal([]byte(cache), &data)
		Consultants = data
		wg.Done()
	}
}
