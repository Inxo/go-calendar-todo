package main

import (
	"database/sql"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
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

	width, height := screen.Size()

	eventAddedCh := make(chan struct{}, 1)

	// Create event form
	var form *EventForm = nil
	formShow := false
	// Запускаем главный цикл
	for {
		drawUI(screen, currentDate, db, form)
		choice := getUserInput(screen, form)

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
		case "a":
			if formShow == false {
				form := createEventForm(screen, db, width/2, height-2, width/2, 3, currentDate, eventAddedCh)
				currentDate = time.Now()
				println("AAA")
				form.draw()
				formShow = true
			}

		case "e":
			if formShow == true {
				text := form.inputText.GetText()
				addEventToDB(db, currentDate, text)
				println("End input")
				form = nil
				screen.Show()
				formShow = false
			}
		case "q":
			return
		default:
		}

		screen.Show()
	}
}

// Функция для отрисовки пользовательского интерфейса
func drawUI(screen tcell.Screen, date time.Time, db *DB, form *EventForm) {
	screen.Clear()
	width, _ := screen.Size()
	drawCalendar(screen, 0, 0, date)
	drawEventList(screen, width/2, 0, width/2, date, db)
	screen.Show()
}

// Функция для печати текста в центре консоли
func printCentered(screen tcell.Screen, fg, bg tcell.Color, msg string) {
	w, h := screen.Size()
	xPos, yPos := w/2-len(msg)/2, h/2
	printCoor(screen, xPos, yPos, tcell.ColorDefault, tcell.ColorDefault, msg)
}
