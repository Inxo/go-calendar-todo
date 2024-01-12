package main

import (
	"github.com/gdamore/tcell/v2"
	"strconv"
)

// Функция для получения выбора пользователя
func getUserInput(screen tcell.Screen, form *EventForm) string {
	event := screen.PollEvent()

	switch eventR := event.(type) {
	case *tcell.EventKey:
		// Handle input for the event form
		if form != nil {
			switch eventR.Key() {
			case tcell.KeyEnter:
				//text := form.inputText.GetText()
				return "e"
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return "q"
			default:
				form.handleInput(eventR)
			}
		} else {
			switch eventR.Key() {
			case tcell.KeyLeft:
				return "l"
			case tcell.KeyRight:
				return "r"
			case tcell.KeyUp:
				return "u"
			case tcell.KeyDown:
				return "d"
			case tcell.KeyRune:
				switch eventR.Rune() {
				case 97:
					return "a"

				case ' ':
					return "n"
				}
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return "q"
			default:
				return strconv.Itoa(int(tcell.KeyRune))
			}
		}
	}

	return ""
}
