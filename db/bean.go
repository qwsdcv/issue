package db

type Article struct {
	ParentID int    `json:"parent_id"`
	Title    string `json:"title"`
}
