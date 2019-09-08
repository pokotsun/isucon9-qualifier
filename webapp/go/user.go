package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

const (
	userSimpleHashKey = "USER_SIMPLE_HASH"
)

func getUser(r *http.Request) (user User, errCode int, errMsg string) {
	session := getSession(r)
	userID, ok := session.Values["user_id"]
	if !ok {
		return user, http.StatusNotFound, "no session"
	}

	err := dbx.Get(&user, "SELECT * FROM `users` WHERE `id` = ?", userID)
	if err == sql.ErrNoRows {
		return user, http.StatusNotFound, "user not found"
	}
	if err != nil {
		log.Print(err)
		return user, http.StatusInternalServerError, "db error"
	}

	return user, http.StatusOK, ""
}

func getUserSimpleByID(q sqlx.Queryer, userID int64) (userSimple UserSimple, err error) {
	user := User{}
	err = sqlx.Get(q, &user, "SELECT * FROM `users` WHERE `id` = ?", userID)
	if err != nil {
		return userSimple, err
	}
	userSimple.ID = user.ID
	userSimple.AccountName = user.AccountName
	userSimple.NumSellItems = user.NumSellItems
	return userSimple, err
}

func (r *Redisful) SetUserSimpleToRedis(us UserSimple) error {
	err := r.SetHashToCache(userSimpleHashKey, us.ID, us)
	return err
}

func (r *Redisful) GetUserSimpleFromRedis(id int, us UserSimple) error {
	err := r.GetHashFromCache(userSimpleHashKey, strconv.Itoa(id), us)
	return err
}

func (r *Redisful) InitializeUserSimple() error {
	var userSimples []UserSimple
	err := dbx.Select(&userSimples, "SELECT id, account_name, num_sell_items FROM users")
	if err != nil {
		return err
	}
	for _, us := range userSimples {
		err := r.SetUserSimpleToRedis(us)
		if err != nil {
			return err
		}
	}
	return nil
}
