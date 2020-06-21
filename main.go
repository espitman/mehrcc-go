package main

import (
	"mehrcc/article"
	"mehrcc/consultant"
	"mehrcc/course"
	"mehrcc/news"
	"mehrcc/project"
	"mehrcc/slider"
	"mehrcc/workshop"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var wg sync.WaitGroup

func main() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	r.GET("/home", func(c *gin.Context) {
		wg.Add(7)
		go consultant.GetConsultants(&wg)
		go project.GetProjects(&wg)
		go article.GetArticles(&wg)
		go news.GetNews(&wg)
		go workshop.GetWorkshops(&wg)
		go slider.GetSliders(&wg)
		go course.GetCourses(&wg)
		wg.Wait()

		c.JSON(200, gin.H{
			"consultants": consultant.Consultants,
			"projects":    project.Projects,
			"articles":    article.Articles,
			"news":        news.AllNews,
			"sliders":     slider.Sliders,
			"courses":     course.Courses,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
