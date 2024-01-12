// event_list.go

package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

// EventForm Structure to hold the event form state
type EventForm struct {
	screen       tcell.Screen
	x, y, width  int
	eventDate    time.Time
	inputText    *tview.InputField
	addButton    *tview.Button
	eventAddedCh chan<- struct{}
}

// Function to create the event form
func createEventForm(screen tcell.Screen, db *DB, x, y, width, height int, date time.Time, eventAddedCh chan<- struct{}) *EventForm {
	inputText := tview.NewInputField()
	inputText.SetLabel("Event: ")
	inputText.SetFormAttributes(1, tcell.ColorDefault, tcell.ColorBlack, tcell.ColorYellow, tcell.ColorBlack)

	addButton := tview.NewButton("Add")
	addButton.SetSelectedFunc(func() {
		if inputText.GetText() != "" {
			description := inputText.GetText()
			addEventToDB(db, date, description)
			eventAddedCh <- struct{}{}
			inputText.SetText("")
		}
	})

	form := &EventForm{
		screen:       screen,
		x:            x,
		y:            y,
		width:        width,
		eventDate:    date,
		inputText:    inputText,
		addButton:    addButton,
		eventAddedCh: eventAddedCh,
	}

	return form
}

// Function to handle user input for the event form
func (form *EventForm) handleInput(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyEnter:
		form.addButton.SetSelectedFunc(TestButton)
	default:
		form.inputText.InputHandler()(event, nil)
	}
}

func TestButton() {
	println("Button")
}

// Function to draw the event form
func (form *EventForm) draw() {
	form.screen.SetContent(form.x, form.y, ' ', nil, tcell.StyleDefault)
	form.inputText.Draw(form.screen)
	form.addButton.Draw(form.screen)
}
