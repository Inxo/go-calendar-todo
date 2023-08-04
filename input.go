package main

import "github.com/gdamore/tcell"

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
