package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_proj/client"
	"payment_proj/common"
	"payment_proj/model"
)

func PublicKey(c *gin.Context)  {
	publicKey, err := client.GetPublicKey()
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
	}else {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg: "OK",
			Code: common.Success,
			Data: publicKey,
		})
	}
}

func SignUp(c *gin.Context) {
	userInfo := &model.User{}
	err := c.ShouldBindJSON(userInfo)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
	}
	err = client.SignUp(userInfo)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
	}
	c.JSON(http.StatusOK, "")
}

func MakeTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}