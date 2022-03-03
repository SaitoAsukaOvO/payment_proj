package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func MakeTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}