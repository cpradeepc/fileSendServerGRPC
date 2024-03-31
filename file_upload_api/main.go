package main

import (
	"log"
	"net/http"

	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", uploadFile)
	r.Run(":9001")
}

func uploadFile(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		log.Println("error while uploading file", err)
		return
	}

	filePath := filepath.Join("upload", f.Filename)
	err = c.SaveUploadedFile(f, filePath)
	if err != nil {
		log.Println("error while save uploading file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fialed to save file"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"msg": "save"})
}
