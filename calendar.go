package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"strconv"
	"time"
)

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

func printCoor(screen tcell.Screen, x int, y int, fg, bg tcell.Color, msg string) {
	style := tcell.StyleDefault.Background(bg).Foreground(fg)
	for i, c := range msg {
		screen.SetContent(x+i, y, c, nil, style)
	}
}
