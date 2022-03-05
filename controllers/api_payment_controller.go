package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_proj/client"
	"payment_proj/common"
	"payment_proj/model"
	"strconv"
)

// Encrypt user password
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


// Used for user to sign up
func CreateUser(c *gin.Context) {
	userInfo := &model.User{}
	err := c.ShouldBindJSON(userInfo)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
		return
	}
	err = client.SignUp(userInfo)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
	}
	c.JSON(http.StatusOK, model.ApiStatus{
		Msg:  "OK",
		Code: common.Success,
	})
}


// Create wallet for certain user
func CreateWallet(c *gin.Context) {
	walletInfo := &model.Wallet{}
	err := c.ShouldBindJSON(walletInfo)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
		return
	}
	client.InitWalletByUserId(walletInfo.UserID, walletInfo.Balance)
	c.JSON(http.StatusOK, model.ApiStatus{
		Msg:  "OK",
		Code: common.Success,
	})
}


// Get user info for certain user
func GetUser(c *gin.Context) {
	userIDStr := c.Query("userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
		return
	}
	userInfo, err := client.GetUserInfo(uint(userID))
	c.JSON(http.StatusOK, model.UserResponse{
		Msg:  "OK",
		Code: common.Success,
		Data: userInfo,
	})
}

func CreateTransaction(c *gin.Context) {
	transaction := &model.Transaction{}
	err := c.ShouldBindJSON(transaction)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
		return
	}
	err = client.MakeTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusOK, model.ApiStatus{
			Msg:  err.Error(),
			Code: common.GeneralErrorCode,
		})
	}
	c.JSON(http.StatusOK, model.ApiStatus{
		Msg:  "OK",
		Code: common.Success,
	})
}