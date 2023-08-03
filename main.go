package main

import (
	"database/sql"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

func main() {
	// Инициализация экрана
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

	db, err := sql.Open("sqlite3", "events.db")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// Создание таблицы events, если она не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_date DATE,
			description TEXT
		)
	`)
	if err != nil {
		panic(err)
	}

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
			fmt.Println("Quit.")
			return
		default:
			//return
			//fmt.Println("Некорректный выбор. Попробуйте еще раз.")
		}
	}
}

// Функция для отрисовки пользовательского интерфейса
func drawUI(screen tcell.Screen, date time.Time, db *sql.DB) {
	screen.Clear()

	width, _ := screen.Size()

	// Отрисовка календаря в левой панели
	drawCalendar(screen, 0, 0, date)

	// Отрисовка списка событий в правой панели
	drawEventList(screen, width/2, 0, width/2, date, db)

	screen.Show()
}

// Функция для отрисовки календаря
func drawCalendar(screen tcell.Screen, x, y int, date time.Time) {
	printCalendar(screen, x, y, date)
}

// Функция для печати календаря текущего месяца
func printCalendar(screen tcell.Screen, x, y int, date time.Time) {
	screen.Clear()

	weekdays := [7]string{"Mo", "Tu", "We", "Th", "Fr", "St", "Su"}

	color := tcell.ColorDefault
	bgColor := tcell.ColorDefault
	for col, weekday := range weekdays {
		x, y := x+col*3, y+1
		color = tcell.ColorGreen
		for i, c := range weekday {
			style := tcell.StyleDefault.Background(bgColor).Foreground(color)
			screen.SetContent(x+i, y, c, nil, style)
		}
	}
	monthCalendar := getMonthCalendar(date)

	//monthRow := 0
	for row, week := range monthCalendar {
		//monthRow = row

		for col, day := range week {
			color = tcell.ColorDefault
			bgColor = tcell.ColorDefault

			x, y := col*3, row+2
			if day == date.Day() {
				color = tcell.ColorWhite
				bgColor = tcell.ColorBlue
			}
			style := tcell.StyleDefault.Background(bgColor).Foreground(color)

			str := strconv.Itoa(day)
			if day < 10 {
				str = "0" + str
			}
			if day == 0 {
				str = "  "
			}
			for i, c := range str {
				screen.SetContent(x+i, y, c, nil, style)
			}
		}
	}

	monthName := date.Format("January")
	year := date.Year()
	w, _ := screen.Size()
	msg := monthName + " " + fmt.Sprintf("%d", year)
	xPos, yPos := 0, 0
	printCoor(screen, xPos, yPos, tcell.ColorRed, tcell.ColorDefault, msg)
	msg = "Events"
	printCoor(screen, w/2-len(msg)+7*3, yPos, tcell.ColorRed, tcell.ColorDefault, msg)

	screen.Show()
}

// Функция для получения календаря месяца
func getMonthCalendar(date time.Time) [][]int {
	var monthCalendar [][]int

	year, month, _ := date.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, date.Location())
	numDays := lastDay.Day()
	firstWeekday := int(firstDay.Weekday())
	if firstWeekday == 0 {
		firstWeekday = 7 // Переход от воскресенья (0) к понедельнику (7)
	}

	week := make([]int, 7)
	day := 1
	for i := 1; i <= numDays+firstWeekday-1; i++ {
		if i >= firstWeekday {
			week[(i-1)%7] = day
			day++
		}

		if i%7 == 0 || i == numDays+firstWeekday-1 {
			monthCalendar = append(monthCalendar, week)
			week = make([]int, 7)
		}
	}

	return monthCalendar
}

// Функция для получения выбора пользователя
func getUserInput(screen tcell.Screen) string {
	event := screen.PollEvent()
	switch event := event.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyLeft:
			return "l"
		case tcell.KeyRight:
			return "r"
		case tcell.KeyUp:
			return "u"
		case tcell.KeyDown:
			return "d"
		case tcell.KeyRune:
			switch event.Rune() {
			case ' ':
				return "n"
			}
		case tcell.KeyEsc, tcell.KeyCtrlC:
			return "q"
		}
	}
	return ""
}

// Функция для отрисовки списка событий
func drawEventList(screen tcell.Screen, x, y, width int, date time.Time, db *sql.DB) {
	// Запрос событий на выбранную дату из базы данных
	rows, err := db.Query("SELECT description FROM events WHERE event_date = ?", date.Format("2006-01-02"))
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
		//drawLine(screen, x+4, tcell.StyleDefault)
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

// Функция для печати текста в центре консоли
//func printCentered(screen tcell.Screen, fg, bg tcell.Color, msg string) {
//	w, h := screen.Size()
//	xPos, yPos := w/2-len(msg)/2, h/2
//
//	printCoor(screen, xPos, yPos, fg, bg, msg)
//}

func printCoor(screen tcell.Screen, x int, y int, fg, bg tcell.Color, msg string) {
	style := tcell.StyleDefault.Background(bg).Foreground(fg)
	for i, c := range msg {
		screen.SetContent(x+i, y, c, nil, style)
	}
}
