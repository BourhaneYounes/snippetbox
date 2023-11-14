package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := ``
}
