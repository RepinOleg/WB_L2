package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
	UserID int       `json:"user_id"`
}

type Calendar struct {
	Events map[int][]*Event
}

// Logger - для логирования запросов
type Logger struct {
	handler http.Handler
}

// Конструктор логгера
func newLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v\n", r.Method, r.URL, time.Since(start))
}

// decode декодирует данные из reader в json
func (e *Event) decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		log.Printf("%+v error from decoder", err)
		return err
	}
	return nil
}

// validate проверяет наличие данных в обязательных полях
func (e *Event) validate() error {
	switch {
	case e.UserID <= 0:
		return fmt.Errorf("invalid user_id")
	case e.ID <= 0:
		return fmt.Errorf("invalid event_id")
	case e.Title == "":
		return fmt.Errorf("invalid title")
	default:
		return nil
	}
}

// Create добавление события в календарь
func (c *Calendar) Create(ev *Event) error {
	if events, ok := c.Events[ev.UserID]; ok {
		for _, event := range events {
			if event.ID == ev.ID {
				return fmt.Errorf("event with id = %v for user with id %v already exists", ev.ID, ev.UserID)
			}
		}
	}
	c.Events[ev.UserID] = append(c.Events[ev.UserID], ev)

	return nil
}

// Update обновление информации о событии в календаре
func (c *Calendar) Update(ev *Event) error {
	if _, ok := c.Events[ev.UserID]; !ok {
		return fmt.Errorf("user with ID = %v doesn't exist", ev.UserID)
	}

	for i, event := range c.Events[ev.UserID] {
		if event.ID == ev.ID {
			// Обновляем событие в слайсе
			c.Events[ev.UserID][i] = ev
			return nil
		}
	}

	return fmt.Errorf("event with ID %v not found (user %v)", ev.ID, ev.UserID)
}

func (c *Calendar) Delete(ev *Event) error {
	if _, ok := c.Events[ev.UserID]; !ok {
		return fmt.Errorf("user %v doesn't exist", ev.UserID)
	}

	events := c.Events[ev.UserID]
	for i, event := range events {
		if event.ID == ev.ID {
			// Удаляем элемент из слайса
			c.Events[ev.UserID] = append(events[:i], events[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("can't find event with %v id for %v user id", ev.ID, ev.UserID)
}

func Response(w http.ResponseWriter, r any, status int) {
	resp := struct {
		Result any `json:"result"`
	}{Result: r}

	res, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func ErrorResponse(w http.ResponseWriter, e string, status int) {
	errResp := struct {
		Error string `json:"error"`
	}{Error: e}

	res, err := json.MarshalIndent(errResp, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	var ev = new(Event)

	if err := ev.decode(r.Body); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ev.validate(); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := calendar.Create(ev); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	Response(w, "Мероприятие добавлено в календарь", http.StatusCreated)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var ev = new(Event)

	if err := ev.decode(r.Body); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ev.validate(); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := calendar.Update(ev); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	Response(w, "Данные о мероприятии успешно обновлены", http.StatusCreated)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var ev = new(Event)

	if err := ev.decode(r.Body); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := calendar.Delete(ev); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	Response(w, "Данные о мероприятии успешно удалены", http.StatusCreated)
}

func (c *Calendar) getEventsForDay(userID int, date time.Time) ([]*Event, error) {
	var res []*Event

	if _, ok := c.Events[userID]; !ok {
		return nil, fmt.Errorf("user %v doesn't exist", userID)
	}

	for _, ev := range c.Events[userID] {
		if ev.Date.Year() == date.Year() && ev.Date.Month() == date.Month() && ev.Date.Day() == date.Day() {
			res = append(res, ev)
		}
	}

	return res, nil
}

func (c *Calendar) getEventsForWeek(userID int, date time.Time) ([]*Event, error) {
	var res []*Event

	if _, ok := c.Events[userID]; !ok {
		return nil, fmt.Errorf("user with id =  %v doesn't exist", userID)
	}

	for _, ev := range c.Events[userID] {
		y1, w1 := ev.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			res = append(res, ev)
		}
	}

	return res, nil
}

func (c *Calendar) getEventsForMonth(userID int, date time.Time) ([]*Event, error) {
	var res []*Event

	if _, ok := c.Events[userID]; !ok {
		return nil, fmt.Errorf("user with id =  %v doesn't exist", userID)
	}

	for _, ev := range c.Events[userID] {
		if ev.Date.Year() == date.Year() && ev.Date.Month() == date.Month() {
			res = append(res, ev)
		}
	}

	return res, nil
}

const dateString = "2006-01-02"

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	var ev []*Event

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateString, r.URL.Query().Get("date"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ev, err = calendar.getEventsForDay(userID, date); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	Response(w, ev, http.StatusOK)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	var ev []*Event

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateString, r.URL.Query().Get("date"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ev, err = calendar.getEventsForWeek(userID, date); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	Response(w, ev, http.StatusOK)
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	var ev []*Event

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateString, r.URL.Query().Get("date"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ev, err = calendar.getEventsForMonth(userID, date); err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	Response(w, ev, http.StatusOK)
}

var calendar Calendar = Calendar{
	Events: make(map[int][]*Event)}

func main() {
	mux := http.NewServeMux()
	// GET
	mux.HandleFunc("/events_for_day", eventsForDayHandler)
	mux.HandleFunc("/events_for_week", eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", eventsForMonthHandler)
	// POST
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)
	mux.HandleFunc("/delete_event", deleteEventHandler)

	wMux := newLogger(mux)

	log.Fatalln(http.ListenAndServe(":8080", wMux))

}
