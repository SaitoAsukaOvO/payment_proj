package client

import (
	"errors"
	"fmt"
	"payment_proj/common"
	"payment_proj/config"
	"payment_proj/model"
)

func GetPublicKey() (string,error) {
	configMap := config.LoadAppConfig()

	publicKey, ok := configMap.Security[common.PublicKey]
	if !ok {
		return publicKey, errors.New("Public key is missing in config")
	}
	return publicKey, nil
}

func SignUp(user *model.User) error{
	chain := DB.Model(&model.User{})
	chain.Create(user)
	if chain == nil {
		errMsg := fmt.Sprintf("Fail to create user")
		return errors.New(errMsg)
	}
	return nil
}