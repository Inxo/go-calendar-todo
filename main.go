package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strconv"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Получаем текущую дату
	currentDate := time.Now()

	// Запускаем главный цикл
	for {
		printCalendar(currentDate)
		choice := getUserInput()

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
			fmt.Println("Выход из программы.")
			return
		default:
			return
			//fmt.Println("Некорректный выбор. Попробуйте еще раз.")
		}
	}
}

// Функция для печати календаря текущего месяца
func printCalendar(date time.Time) {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	weekdays := [7]string{"Mo", "Tu", "We", "Th", "Fr", "St", "Su"}

	color := termbox.ColorDefault
	for col, weekday := range weekdays {
		x, y := col*3, 1
		color = termbox.ColorGreen
		for i, c := range weekday {
			termbox.SetCell(x+i, y, c, color, termbox.ColorDefault)
		}
	}
	monthCalendar := getMonthCalendar(date)

	//monthRow := 0
	for row, week := range monthCalendar {
		//monthRow = row
		for col, day := range week {
			x, y := col*3, row+2
			color = termbox.ColorDefault
			if day == date.Day() {
				color = termbox.ColorRed // Цвет текущего дня (красный)
			}

			str := strconv.Itoa(day)
			if day < 10 {
				str = "0" + str
			}
			if day == 0 {
				str = "  "
			}
			for i, c := range str {
				termbox.SetCell(x+i, y, c, color, termbox.ColorDefault)
			}
		}
	}

	monthName := date.Format("January")
	year := date.Year()
	w, _ := termbox.Size()
	msg := monthName + " " + fmt.Sprintf("%d", year)
	xPos, yPos := 0, 0
	printCoor(xPos, yPos, termbox.ColorRed, termbox.ColorDefault, msg)
	msg = "Tasks"
	printCoor(w/2-len(msg)+7*3, yPos, termbox.ColorRed, termbox.ColorDefault, msg)

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
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
func getUserInput() string {
	event := termbox.PollEvent()
	if event.Type == termbox.EventKey {
		switch event.Key {
		case termbox.KeyArrowLeft:
			return "l"
		case termbox.KeyArrowRight:
			return "r"
		case termbox.KeyArrowUp:
			return "u"
		case termbox.KeyArrowDown:
			return "d"
		case termbox.KeySpace:
			return "n"
		case termbox.KeyEsc, termbox.KeyCtrlC:
			return "q"
		}
	}
	return ""
}

// Функция для печати текста в центре консоли
func printCentered(fg, bg termbox.Attribute, msg string) {
	w, h := termbox.Size()
	xPos, yPos := w/2-len(msg)/2, h/2

	printCoor(xPos, yPos, fg, bg, msg)
}

func printCoor(x int, y int, fg, bg termbox.Attribute, msg string) {
	for i, c := range msg {
		termbox.SetCell(x+i, y, c, fg, bg)
	}
}
