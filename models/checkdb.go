package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" //sqlite3
)

//SqliteInstance a global instance of the sqlite3
var SqliteInstance *sql.DB

const dbFile = "./.issue.sqlite.db"

//init database.Will auto create table needed if table not exist.
func init() {
	_, err := os.Stat(dbFile)
	shouldInitTable := !os.IsExist(err)

	SqliteInstance, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		defer SqliteInstance.Close()
		SqliteInstance = nil
		panic(err)
	}

	if shouldInitTable {
		initTable()
	}
}

func initTable() {
	if SqliteInstance == nil {
		panic("Error when trying to initialize table: SqliteInstance is nil")
	}
	createTableArticle()
	createTableComment()
}

func createTable(sql string) {
	_, err := SqliteInstance.Exec(sql)
	if err != nil {
		log.Fatal("Error when creating table.")
		log.Fatal(err)
	}
}

func createTableArticle() {
	sqlString := `
	CREATE TABLE IF NOT EXISTS articles(
		id integer primary key autoincrement,
		parent_id integer NOT NULL,
		title text NOT NULL,
		create_date datetime NOT NULL,
		type text NOT NULL,
		content text NOT NULL DEFAULT '',
		visits integer NOT NULL DEFAULT 0
	)
	`
	createTable(sqlString)
}

func createTableComment() {
	sqlString := `
	CREATE TABLE IF NOT EXISTS comments(
		id integer primary key autoincrement,
		articleid integer NOT NULL,
		nick_name text NOT NULL,
		ip integer NOT NULL,
		content text NOT NULL,
		create_date datetime NOT NULL,
		foreign key(articleid) references articles(id)
	)
	`
	createTable(sqlString)
}
