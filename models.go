package main

import "time"

type Todo struct {
	ID       int        `db:"id"`
	Date     *time.Time `db:"date"`
	Deadline *time.Time `db:"deadline"`
	Done     bool       `db:"done"`
	Text     string     `db:"text"`
	Category string     `db:"category"`
}

func NewTodo(text string) Todo {
	return Todo{
		Done:     false,
		Text:     text,
		Category: "default",
	}
}

type Category struct {
	ID   int
	Name string
}
