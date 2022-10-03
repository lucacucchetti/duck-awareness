package controllers

import (
	"encoding/json"
	. "example/hello/src/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTextNote(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		noteId := c.Param("noteId")
		note, _ := repo.GetTextNote(noteId)
		noteJson, err := json.Marshal(note)
		if err != nil {
			handleRequestError(c, err)
			return
		}
		c.Data(http.StatusOK, gin.MIMEJSON, noteJson)
	}
}
