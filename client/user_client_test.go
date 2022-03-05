package client

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/liuhongdi/unittest04/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

var mock sqlmock.Sqlmock
var db *sql.DB

func init() {
	var err error
	db, mock, err = sqlmock.New()
	if nil != err {
		log.Fatalf("Init sqlmock failed, err %v", err)
	}
	global.DBLink, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn: db,
	}), &gorm.Config{})
	if nil != err {
		log.Fatalf("Init DB with sqlmock failed, err %v", err)
	}
}


func TestGetUserInfo(t *testing.T) {
	var id uint
	id = 1
	usr, _ := GetUserInfo(id)
	assert.Equal(t, usr.ID, uint(1))
	assert.Equal(t, usr.Phone, "17712910130")
	assert.Equal(t, usr.Password, "12345")
	assert.Equal(t, len(usr.Wallets) > 0, true)
	assert.Equal(t, len(usr.Transactions)>0, true)
}