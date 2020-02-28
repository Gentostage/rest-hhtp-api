package model

type Post struct {
	ID     int    `json: id`
	UserId int    `json: user_id`
	Text   string `json: text`
	Title  string `json: title`
}
