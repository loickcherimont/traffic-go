package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loickcherimont/traffic/model"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		// Inject into context data fetch from API
		c.Data(http.StatusOK, "string", fetchData())
		var fetch model.Fetch

		if err := c.ShouldBindJSON(&fetch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mytotal":   fetch.Total,
			"myresults": fetch.Results,
		})

		c.String(http.StatusOK, "Results: \n%s", fetchData())
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

// Call API
// Fetch results after calling
// Gin ?
func fetchData() []byte {
	resp, err := http.Get("https://data.loire-atlantique.fr/api/explore/v2.1/catalog/datasets/224400028_info-route-departementale/records?limit=5")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return []byte(body)
}
