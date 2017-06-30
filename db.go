package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	database *sqlx.DB
}

func openDb() *DB {
	database, err := sqlx.Open("sqlite3", "./todo.db")
	if err != nil {
		fmt.Println("Couldn't open the database")
		os.Exit(2)
	}

	_, err = database.Exec("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, text TEXT, done INTEGER, date DATETIME, deadline DATETIME, category TEXT)")
	if err != nil {
		fmt.Println("Couldn't create the todos table")
		os.Exit(2)
	}

	return &DB{database}
}

func (d *DB) addTodo(todo Todo) {
	_, err := d.database.Exec("INSERT INTO todos (date,deadline,done,text,category) VALUES (?,?,?,?,?)", todo.Date, todo.Deadline, todo.Done, todo.Text, todo.Category)
	if err != nil {
		fmt.Println("Couldn't insert the TODO:", err)
		os.Exit(3)
	}
}

func (d *DB) getTodos(all bool, category *string) []Todo {
	todos := []Todo{}

	queryString := "SELECT * FROM todos"
	if !all {
		queryString += " WHERE done = 0"
	}
	var err error
	if category != nil {
		if !all {
			queryString += " AND"
		} else {
			queryString += " WHERE"
		}
		queryString += " category = ?"
		err = d.database.Select(&todos, queryString, *category)
	} else {
		err = d.database.Select(&todos, queryString)
	}

	if err != nil {
		fmt.Println("Couldn't retrieve the TODOs:", err)
		os.Exit(3)
	}
	return todos
}

func (d *DB) getTodo(id int) (*Todo, error) {
	todo := Todo{}
	err := d.database.Get(&todo, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (d *DB) setDone(id int, done bool) (*Todo, error) {
	todo, err := d.getTodo(id)
	if err != nil {
		return nil, err
	}
	_, err = d.database.Exec("UPDATE todos SET done = ? WHERE id = ?", done, todo.ID)
	return todo, err
}

func (d *DB) setDate(id int, time *time.Time) (*Todo, error) {
	todo, err := d.getTodo(id)
	if err != nil {
		return nil, err
	}
	_, err = d.database.Exec("UPDATE todos SET date = ? WHERE id = ?", time, todo.ID)
	return todo, err
}

func (d *DB) setDeadline(id int, time *time.Time) (*Todo, error) {
	todo, err := d.getTodo(id)
	if err != nil {
		return nil, err
	}
	_, err = d.database.Exec("UPDATE todos SET deadline = ? WHERE id = ?", time, todo.ID)
	return todo, err
}
