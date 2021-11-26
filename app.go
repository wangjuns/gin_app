package main

import (
	"crypto/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@."

func randomPassword(length int) string {
	l := len(chars)
	b := make([]byte, length)
	rand.Read(b)

	for i := range b {
		b[i] = chars[int(b[i])%l]
	}
	return string(b)
}

func main() {
	r := gin.Default()
	r.GET("/pwd/:len", func(c *gin.Context) {
		len, _ := strconv.Atoi(c.Param("len"))

		c.JSON(200, randomPassword(len))
	})

	r.Run()
}
