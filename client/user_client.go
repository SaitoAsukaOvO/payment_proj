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

// This function is used to create a user
func SignUp(user *model.User) error {
	userChain := DB.Model(&model.User{})
	if userChain == nil {
		errMsg := fmt.Sprintf("Fail to create user")
		return errors.New(errMsg)
	}

	// init a wallet when user is creating
	wallet := InitWalletByUserId(user.ID)
	var wallets []model.Wallet
	wallets = append(wallets, wallet)
	user.Wallets = wallets

	userChain.Create(user)
	return nil
}

func InitWalletByUserId(userID uint) model.Wallet {
	return model.Wallet{
		UserID:  userID,
		Balance: 0,
	}
}

func GetUserInfo(UserID uint) (model.User, error) {
	user := model.User{}
	err := DB.Model(&user).Where("ID = ? ", UserID).Find(&user).Error
	// Get all wallets
	var wallets []model.Wallet

	// Find wallets for a certain user
	err = DB.Model(&model.Wallet{}).Where("user_id = ? ", UserID).Find(&wallets).Error
	if err != nil {
		return user, err
	}
	user.Wallets = wallets
	return user, nil
}