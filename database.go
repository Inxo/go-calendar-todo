package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DB Структура для работы с базой данных
type DB struct {
	db *sql.DB
}

// Функция для подключения к базе данных
func connectDB(dbName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	// Создание таблицы events, если она не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_date DATE,
			description TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

// Метод для получения событий на выбранную дату из базы данных
func (db *DB) getEvents(date time.Time) ([]string, error) {
	rows, err := db.db.Query("SELECT description FROM events WHERE event_date = ?", date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var events []string
	for rows.Next() {
		var description string
		err = rows.Scan(&description)
		if err != nil {
			return nil, err
		}
		events = append(events, description)
	}

	return events, nil
}
