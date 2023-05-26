package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func AddMultipleImage(c *gin.Context, Images []*multipart.FileHeader) ([]string, error) {
	fileArray := make([]string, 0)

	for _, Image := range Images {
		file, err := Image.Open()

		if err != nil {

			return fileArray, err
		}

		defer file.Close()
		uploadDir := "../images/"

		// Create the directory if it doesn't exist
		err = os.MkdirAll(uploadDir, os.ModePerm)

		filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), Image.Filename) // name the file
		filePath := filepath.Join(uploadDir, filename)                          //specify image path

		out, err := os.Create(filePath)

		if err != nil {

			return fileArray, err
		}

		defer out.Close()

		_, err = io.Copy(out, file)

		if err != nil {

			return fileArray, err
		}

		fileArray = append(fileArray, filePath)

	}
	return fileArray, nil
}

func AddSingleImage(c *gin.Context, n string) (string, error) {
	var filePath string

	file, header, err := c.Request.FormFile(n)

	if err != nil {

		return "", err
	}
	uploadDir := "../images/"

	// Create the directory if it doesn't exist
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), header.Filename) // name the file
	filePath = filepath.Join(uploadDir, filename)                            //specify image path
	out, err := os.Create(filePath)

	if err != nil {

		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {

		return "", err
	}

	return filePath, nil
}
