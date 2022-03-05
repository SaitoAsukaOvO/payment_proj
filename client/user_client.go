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
	wallet := InitWalletByUserId(user.ID, 0)
	var wallets []model.Wallet
	wallets = append(wallets, wallet)
	user.Wallets = wallets

	userChain.Create(user)
	return nil
}

func InitWalletByUserId(userID uint, balance float64) model.Wallet {
	wallet := model.Wallet{
		UserID:  userID,
		Balance: balance,
	}

	// insert wallet record
	DB.Model(&wallet).Create(&wallet)
	return wallet
}

func GetUserInfo(UserID uint) (model.User, error) {
	user := model.User{}
	err := DB.Model(&user).Where("ID = ? ", UserID).Find(&user).Error
	var wallets []model.Wallet
	var transactions []model.Transaction

	// Find wallets for a certain user
	err = DB.Model(&model.Wallet{}).Where("user_id = ? ", UserID).Find(&wallets).Error
	if err != nil {
		return user, err
	}
	user.Wallets = wallets

	err = DB.Model(&model.Transaction{}).Where("user_id = ? or to_user_id = ?", UserID, UserID).Find(&transactions).Error
	if err != nil {
		return user, err
	}
	user.Transactions = transactions
	return user, nil
}