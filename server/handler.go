package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/startdusk/soom/chat"
)

type Service struct {
	ChatService chat.Service
}

func (s *Service) GetChat(c *gin.Context) {
	name := c.Param("chat_name")

	// TODO: using chat service

	c.JSON(http.StatusOK, gin.H{"chat_name": name})
}

func (s *Service) CreateChat(c *gin.Context) {

}
