package client

import (
	"errors"
	"payment_proj/model"
)

func MakeTransaction(transaction *model.Transaction) error {
	var user model.User
	userID := transaction.UserID
	DB.Model(&user).Where("ID = ? ", userID).Find(&user)

	// TODO Need to decrypt pwd using private key
	pwd := user.Password
	if pwd != transaction.Password {
		return errors.New("invalid transaction password")
	}
	err := DB.Model(&transaction).Create(&transaction).Error
	if err != nil {
		return err
	}

	// Insert Transaction record
	err = updateWallet(transaction.UserID, transaction.ToUserID, transaction.FromWalletId, transaction.ToWalletId, transaction.Amount)
	if err != nil {
		return err
	}

	return nil
}

func updateWallet(fromUser uint, toUser uint, fromWallet int, toWallet int, amount float64) error {
	// TODO Need To Use Optimization Lock To Ensure Transaction
	var payerWallet model.Wallet
	var receiverWallet model.Wallet
	DB.Model(&model.Wallet{}).Where("user_id = ? and id = ?", fromUser, fromWallet).Find(&payerWallet)
	DB.Model(&model.Wallet{}).Where("user_id = ? and id = ?", toUser, toWallet).Find(&receiverWallet)
	payerBalance := payerWallet.Balance
	receiverBalance := receiverWallet.Balance
	updatedPayerBalance := payerBalance - amount
	updatedReceiverBalance := receiverBalance + amount

	// Check if it is available to make a transaction
	if updatedPayerBalance < 0 {
		return errors.New("insufficient balance")
	}
	DB.Model(&model.Wallet{}).Where("user_id = ? and id = ?", fromUser, fromWallet).Update("balance", updatedPayerBalance)
	DB.Model(&model.Wallet{}).Where("user_id = ? and id = ?", toUser, toWallet).Update("balance", updatedReceiverBalance)
	return nil
}
