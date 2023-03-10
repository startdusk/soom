package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) Routes() http.Handler {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	// =========================================================================
	// Chat API
	{
		chatAPI := apiV1.Group("/chat")
		chatAPI.GET("/:chat_name", s.GetChat)
		chatAPI.POST("", s.CreateChat)
	}

	return router
}
