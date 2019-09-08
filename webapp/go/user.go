package main

import (
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

const (
	USER_CACHE_KEY = string("USERS")
	// UESR_NUM_KEY   = string("U-ID-")
)

// func makeKey(userID int64) string {
// 	return fmt.Sprintf("%s%d", USER_CACHE_KEY, userID)
// }

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

func (r *Redisful) InitUsers() error {
	rows, err := dbx.Query("SELECT id, account_name, num_sell_items FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var u UserSimple
		if err := rows.Scan(&u.ID, &u.AccountName, &u.NumSellItems); err != nil {
			return err
		}
		r.SetHashToCache(USER_CACHE_KEY, u.ID, u)
	}
	return nil
}

func (r *Redisful) AddUser(u UserSimple) error {
	err := r.SetHashToCache(USER_CACHE_KEY, u.ID, u)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redisful) GetUserSimpleByID(userID int64) (UserSimple, error) {
	var user UserSimple
	err := r.GetHashFromCache(USER_CACHE_KEY, strconv.Itoa(int(userID)), &user)
	if err != nil {
		return user, err
	}
	return user, err
}
