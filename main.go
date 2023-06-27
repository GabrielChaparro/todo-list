package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)

	router.Run("localhost:8080")
}

type task struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Date       string `json:"date"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

// tasks slice to seed record album data.
var tasks = []task{
	{ID: "1", Title: "Notes", Content: "Example content list", Date: "2019-01-01", CreatedBy: "Gabriel", ModifiedBy: "Gabriel"},
	{ID: "2", Title: "Pass", Content: "Example content list", Date: "2019-01-01", CreatedBy: "Gabriel", ModifiedBy: "Gabriel"},
	{ID: "3", Title: "Supermarket", Content: "Example content list", Date: "2019-01-01", CreatedBy: "Gabriel", ModifiedBy: "Gabriel"},
}

// getTasks responds with the list of all albums as JSON.
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}
