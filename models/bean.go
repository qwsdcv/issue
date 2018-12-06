package models

import (
	"database/sql"
	"log"
	"sort"
	"strconv"
	"time"
)

//Article is a Go bean for table articles
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

//Comment is a Go bean for table comments
type Comment struct {
	ID         int    `json:"id"`
	ArticleID  int    `json:"parent_id"`
	NickName   string `json:"nick_name"`
	IP         string `json:"ip"`
	Content    string `json:"content"`
	CreateDate string `json:"date"`
}

var (
	folder  = "Folder"
	article = "Article"
)

//AddMenu add menu to DB
func (ar *Article) AddMenu() (ret error) {
	con := SqliteInstance
	tm := time.Now()
	t := FormatTime(&tm)

	rs, err := con.Exec("INSERT INTO articles(parent_id,title,create_date,type) VALUES(?,?,?,?)", ar.ParentID, ar.Title, t, ar.Type)
	if err != nil {
		log.Println(err.Error())
		ret = err
	}
	log.Println(rs)
	ret = nil
	return
}

//AddComment add comment to DB
func (one *Comment) AddComment() (ret error) {
	con := SqliteInstance
	tm := time.Now()
	t := FormatTime(&tm)

	rs, err := con.Exec("INSERT INTO comments(articleid,nick_name,ip,content,create_date) VALUES(?,?,?,?,?)", one.ArticleID, one.NickName, one.IP, one.Content, t)
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
	con := SqliteInstance
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
		if value.ParentID != 0 {
			kid := Map[key]
			father := Map[value.ParentID]
			father.Data = append(father.Data, *kid)
		}
	}
	for _, value := range Map {
		if value.ParentID == 0 {
			retArray = append(retArray, *value)
		}
	}
	sortRecursive(atcls(retArray))
	jsonObj = retArray

	return
}

func sortRecursive(res atcls) {
	sort.Sort(atcls(res))
	for _, value := range res {
		if len(value.Data) > 0 {
			sortRecursive(value.Data)
		}
	}
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

//GetContent return the Article whose Id is @id
func GetContent(id string) (ar Article, err error) {
	con := SqliteInstance
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, err = con.Exec("update articles set visits=visits+1 where id=?", intID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	rows := con.QueryRow("select id,parent_id,title,create_date,type,content,visits from articles where id=?", intID)
	t := new(time.Time)
	var content sql.NullString
	err = rows.Scan(&ar.ID, &ar.ParentID, &ar.Title, &t, &ar.Type, &content, &ar.Visits)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if content.Valid {
		ar.Content = content.String
	}
	ar.CreateDate = FormatTime(t)

	return
}

//SetContent update DB with content.
func SetContent(id, content string) (err error) {
	con := SqliteInstance
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, err = con.Exec("update articles set content=? where id=?", content, intID)
	return
}

//GetComment return Comment as Array
func GetComment(articleID string) (jsonObj []Comment, err error) {
	con := SqliteInstance
	rows, err := con.Query("select id,nick_name,ip,content,create_date from comments where articleid=? order by create_date desc", articleID)
	defer rows.Close()
	if err != nil {
		return
	}
	retArray := make([]Comment, 0, 64)
	for rows.Next() {
		t := new(time.Time)
		one := new(Comment)
		err = rows.Scan(&one.ID, &one.NickName, &one.IP, &one.Content, &t)
		if err != nil {
			log.Println(err.Error())
			return
		}
		one.CreateDate = FormatTime(t)
		retArray = append(retArray, *one)
	}

	jsonObj = retArray

	return
}
