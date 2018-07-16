package models

import (
	"database/sql"
	"log"
	"sort"
	"time"
)

//Article is a Go bean
type Article struct {
	ID         int       `json:"id"`
	ParentID   int       `json:"parent_id"`
	Title      string    `json:"title"`
	CreateDate string    `json:"date"`
	Type       string    `json:"type"`
	Content    string    `json:"content"`
	Visits     string    `json:"visits"`
	Data       []Article `json:"nodes"`
}

var (
	folder  = "Folder"
	article = "Article"
)

//AddMenu add menu to DB
func (ar *Article) AddMenu() (ret error) {
	con, err := sql.Open(DBName, ConnectString)
	defer con.Close()
	if err != nil {
		ret = err
	}
	tm := time.Now()
	t := FormatTime(&tm)

	rs, err := con.Exec("INSERT INTO articles(parent_id,title,create_date,type) VALUES(?,?,?,?)", ar.ParentID, ar.Title, t, folder)
	if err != nil {
		log.Println(err.Error())
		ret = err
	}
	log.Println(rs)
	ret = nil
	return
}

// GetMenu return menu hierarchy
func GetMenu() (jsonObj []Article, err error) {
	con, err := sql.Open(DBName, ConnectString)
	defer con.Close()
	if err != nil {
		return
	}
	rows, err := con.Query("select id,parent_id,title,create_date,type from articles")
	defer rows.Close()
	if err != nil {
		return
	}
	Map := make(map[int]*Article)
	for rows.Next() {
		t := new(time.Time)
		ar := new(Article)
		err = rows.Scan(&ar.ID, &ar.ParentID, &ar.Title, &t, &ar.Type)
		if err != nil {
			log.Println(err.Error())
			return
		}
		ar.CreateDate = FormatTime(t)
		Map[ar.ID] = ar
	}
	retArray := make([]Article, 0, 64)
	for key, value := range Map {
		if value.ParentID == 0 {
			retArray = append(retArray, *value)
		} else {
			kid := Map[key]
			father := Map[value.ParentID]
			father.Data = append(father.Data, *kid)
		}
	}
	sort.Sort(atcls(retArray))
	jsonObj = retArray

	return
}

type atcls []Article

func (t atcls) Len() int {
	return len(t)
}
func (t atcls) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t atcls) Less(i, j int) bool {
	return t[i].ID < t[j].ID
}
