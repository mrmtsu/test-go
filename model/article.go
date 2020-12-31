package model

type Article struct {
	ID    int    `db:"id"`
	TITLE string `db:"title"`
}
