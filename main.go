package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/brianloveswords/airtable"
	"github.com/gin-gonic/gin"
)

type quote struct {
	airtable.Record
	Fields struct {
		Body      string `json:"body"`
		Author    string `json:"author"`
		Source    string `json:"source"`
		SourceURL string `json:"source_url"`
	}
}

func main() {
	airtableClient := airtable.Client{
		APIKey: os.Getenv("AIRTABLE_API_KEY"),
		BaseID: os.Getenv("AIRTABLE_BASE_ID"),
	}
	airtableTable := airtableClient.Table(os.Getenv("AIRTABLE_TABLE_NAME"))

	quotes := []quote{}
	err := airtableTable.List(&quotes, &airtable.Options{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		random := rand.Intn(len(quotes))
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, quotes[random].Fields)
	})

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
