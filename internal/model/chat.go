package model

type Chat struct {
	ID   int64 `db:"id"`
	Info Info  `db:"info"`
}

type User struct {
	ID     int64  `db:"id"`
	Name   string `db:"name"`
	From   int64  `db:"user_id"`
	ChatID int64  `db:"chat_id"`
	Text   string `db:"text"`
}

type Info struct {
	ChatID  int64    `db:"chat_id"`
	Owner   int64    `db:"owner"`
	Users   []string `db:"users"`
	From    int64    `db:"user_id"`
	Content string   `db:"content"`
}
