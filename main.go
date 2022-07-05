package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// user info
type User struct {
	Group    string
	Password string
	Port     int
}

func main() {

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	// 	AllowCredentials: true,
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	// 	return origin == "https://github.com"
	// 	// },
	// 	MaxAge: 12 * time.Hour,
	// }))
	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*")

	// build authorized group

	yfile, err := ioutil.ReadFile("people.yml")

	if err != nil {

		log.Fatal(err)
	}

	data := make(map[string]User)

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	data_user := make(map[string]string)
	for k, v := range data {

		data_user[k] = v.Password
		fmt.Print(data_user)
	}

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/", gin.BasicAuth(data_user))

	authorized.Static("/cours", "./cours")
	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "cours.html", gin.H{
			"user": c.MustGet(gin.AuthUserKey).(string),
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
