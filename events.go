package main

import (
	"database/sql"
	"github.com/gdamore/tcell/v2"
	"time"
)

// Функция для отрисовки списка событий
func drawEventList(screen tcell.Screen, x, y, width int, date time.Time, db *DB) {
	// Запрос событий на выбранную дату из базы данных
	rows, err := db.db.Query("SELECT description FROM events WHERE event_date = ?", date.Format("2006-01-02"))
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	// Перемещаем курсор на начало правой панели
	x = width / 2
	y = 0

	// Выводим заголовок списка событий
	y++
	drawText(screen, x+3, y, tcell.StyleDefault.Foreground(tcell.ColorYellow), "События на "+date.Format("02.01.2006"))
	y++

	// Выводим список событий
	for rows.Next() {
		var description string
		err = rows.Scan(&description)
		if err != nil {
			panic(err)
		}

		drawText(screen, x+3, y, tcell.StyleDefault, description)
		//drawLine(screen, x+4, t cell.StyleDefault)
		y++
	}
}

// Функция для вывода текста на экран
func drawText(screen tcell.Screen, x, y int, style tcell.Style, text string) {
	for _, ch := range text {
		screen.SetContent(x, y, ch, nil, style)
		x++
	}
}
