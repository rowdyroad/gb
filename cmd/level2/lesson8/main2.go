package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"sync/atomic"
)

var (
	dbGuard int64
	db   *sql.DB
)

func Db() (*sql.DB, error) {
	if !atomic.CompareAndSwapInt64(&dbGuard, 0, 1) {
		return nil, errors.New("connecting")
	}
	defer atomic.StoreInt64(&dbGuard, 0)
	var err error
	if db != nil {
		err = db.Ping()
	}
	if db == nil || err != nil {
		db, err = sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/test")
		if err != nil {
			return nil, err
		}
		err = db.Ping()
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
func main() {
	db, err := Db()
	if err != nil {
		panic(err)
	}
	var v int
	r := db.QueryRow("SELECT 1")
	err = r.Scan(&v)
	fmt.Println(v, err)
}
