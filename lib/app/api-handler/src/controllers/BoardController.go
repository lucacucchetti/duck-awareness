package controllers

import (
	"encoding/json"
	. "example/hello/src/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoard(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		boardId := c.Param("boardId")
		board, _ := repo.GetBoard(boardId)
		boardJson, err := json.Marshal(board)
		if err != nil {
			handleRequestError(c, err)
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, boardJson)
	}
}
