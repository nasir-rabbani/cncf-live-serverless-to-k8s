package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Text string `json:"text" binding:"required"`
}

type Response struct {
	WordCount     int    `json:"word_count"`
	CharCount     int    `json:"char_count"`
	SentenceCount int    `json:"sentence_count"`
	Source        string `json:"source"`
}

func handler(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request": err.Error()})
		return
	}

	if req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request": "missing text"})
		return
	}

	words := strings.Fields(req.Text)
	wordCount := len(words)
	charCount := len(req.Text)

	re := regexp.MustCompile(`\w+[.!?]`)
	sentenceCount := len(re.FindAllString(req.Text, -1))

	response := Response{
		WordCount:     wordCount,
		CharCount:     charCount,
		SentenceCount: sentenceCount,
		Source:        "k8s",
	}

	c.JSON(http.StatusOK, response)
}
func main() {
	r := gin.Default()
	r.POST("/analyze", handler)
	r.Run(":8083")
}
