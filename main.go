package main

import (
	"database/sql"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"time"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	defer screen.Fini()

	// Инициализация кодировки терминала
	encoding.Register()

	err = screen.Init()
	if err != nil {
		panic(err)
	}

	db, err := connectDB("events.db")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db.db)

	// Получаем текущую дату
	currentDate := time.Now()

	// Запускаем главный цикл
	for {
		drawUI(screen, currentDate, db)
		choice := getUserInput(screen)

		// Обработка выбора пользователя
		switch choice {
		case "l":
			currentDate = currentDate.AddDate(0, 0, -1) // Переключаемся на предыдущий день
		case "r":
			currentDate = currentDate.AddDate(0, 0, 1) // Переключаемся на следующий день
		case "d":
			currentDate = currentDate.AddDate(0, 0, 7) // Переключаемся на следующую неделю
		case "u":
			currentDate = currentDate.AddDate(0, 0, -7) // Переключаемся на следующую неделю
		case "n":
			currentDate = time.Now()
		case "q":
			return
		default:
		}
	}
}

// Функция для отрисовки пользовательского интерфейса
func drawUI(screen tcell.Screen, date time.Time, db *DB) {
	screen.Clear()
	width, _ := screen.Size()
	drawCalendar(screen, 0, 0, date)
	drawEventList(screen, width/2, 0, width/2, date, db)
	screen.Show()
}

// Функция для печати текста в центре консоли
//func printCentered(screen t cell.Screen, fg, bg t cell.Color, msg string) {
//	w, h := screen.Size()
//	xPos, yPos := w/2-len(msg)/2, h/2
//
//	printCoordinate(screen, xPos, yPos, fg, bg, msg)
//}
