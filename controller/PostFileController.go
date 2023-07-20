package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
)

func handleImage(data []byte) error {

	file, err := os.Create("./image.jpg")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	numBytes, err := file.Write(data)
	if err != nil {
		return err
	}

	// 校验写入的字节数
	if numBytes != len(data) {
		return errors.New("didn't write expected number of bytes")
	}

	return nil

}

func ImgPost(c *gin.Context) {
	data, _ := c.GetRawData()

	err := handleImage(data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "image saved"})
}
