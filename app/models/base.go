package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"golang_udemy/todo_app_heroku/config"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

/*
const (
	tableNameUser    = "users"
	tableNameToDo    = "todos"
	tableNameSession = "sessions"
)
*/

func init() {

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Println(err)
	}

	/*
		Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

		if err != nil {
			log.Println(err)
		}

		// ' ではなく `である点は注意
		cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME)`, tableNameUser)

		_, err = Db.Exec(cmdU)
		if err != nil {
			log.Println(err)
		}

		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			user_id INTEGER,
			created_at DATETIME)`, tableNameToDo)
		Db.Exec(cmdT)

		cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			email STRING,
			user_id INTEGER,
			created_at DATETIME)`, tableNameSession)
		_, err = Db.Exec(cmdS)
		if err != nil {
			log.Println(err)
		}
	*/
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	// sha1でハッシュ化
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
