package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleRequestError(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}