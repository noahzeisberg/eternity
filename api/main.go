package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	Client = &http.Client{Transport: &http.Transport{}}
	Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	School = "hbs-hattersheim"
	Class  = "08G2"
)

func main() {
	engine := gin.New()
	engine.GET("/substitutions", func(c *gin.Context) {
		t := c.Query("t")
		add, err := strconv.Atoi(t)
		datetime := time.Now().AddDate(0, 0, add)
		payload, err := GetSubstitutionPlan(datetime)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, payload)
	})
	engine.GET("/chores", func(c *gin.Context) {
		c.JSON(200, GetChores())
	})
	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
