package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
)

var (
	Client = &http.Client{Transport: &http.Transport{}}
	Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	School = "hbs-hattersheim"
	Class  = "08G2"
)

func main() {
	engine := gin.Default()
	engine.GET("/substitutions", func(c *gin.Context) {
		payload, err := GetSubstitutionPlan()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, payload)
	})
	engine.GET("/chores", func(c *gin.Context) {
		firstStudent, secondStudent := GetChores()
		c.JSON(200, []string{firstStudent, secondStudent})
	})
	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
