# Console Calendar and Events App

## Description:
The Console Calendar and Events App is a simple command-line application written in Go (Golang) that allows users to view a calendar for the current month and navigate to previous or next months using arrow keys. The application also displays a list of events associated with specific dates, loaded from an SQLite database.

## Features:

- Calendar Display: The application shows a calendar for the current month, with the ability to move back and forth between months using the left and right arrow keys.
- Event List: The right half of the console displays a list of events associated with specific dates. Events are stored in an SQLite database, and the app fetches events for the selected date.
- Navigation: Users can navigate through the calendar and event list using arrow keys (left, right, and up-down keys).
- Event Management: The app can manage events stored in the SQLite database. Users can add, edit, and delete events for specific dates.

## Usage:
- Navigation: Use the left and right arrow keys to navigate to the previous or next month on the calendar.
- Event List: The application fetches and displays events for the selected date in the right panel.
- Add/Manage Events: The user can manage events through the SQLite database. Events can be added, edited, or deleted for specific dates.
