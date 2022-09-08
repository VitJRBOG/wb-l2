package server

import (
	"dev11/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func createEvent(dbConnection *gorm.DB, w http.ResponseWriter, r *http.Request) ([]byte, Error) {
	event := models.Event{}

	userID := r.FormValue("user_id")
	if userID == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}

	var err error
	event.UserID, err = strconv.Atoi(userID)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	title := r.FormValue("title")
	if title == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'title' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.Title = title

	description := r.FormValue("descr")
	if description == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'descr' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.Description = description

	date := r.FormValue("e_date")
	if date == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'e_date' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.Date, err = time.Parse("2006-01-02", date)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'date' value have a incorrect format",
			Code: http.StatusBadRequest,
		}}
	}

	err = event.Create(dbConnection)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	result := NewResult([]models.Event{event})
	data, err := result.ParseToJSON()
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	return data, Error{}
}

func updateEvent(dbConnection *gorm.DB, w http.ResponseWriter, r *http.Request) ([]byte, Error) {
	event := models.Event{}

	id := r.FormValue("id")
	if id == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	var err error
	event.ID, err = strconv.Atoi(id)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	userID := r.FormValue("user_id")
	if userID == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.UserID, err = strconv.Atoi(userID)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	title := r.FormValue("title")
	if title == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'title' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.Title = title

	description := r.FormValue("descr")
	if description == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'descr' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.Description = description

	date := r.FormValue("e_date")
	if date == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'e_date' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	event.Date, err = time.Parse("2006-01-02", date)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'e_date' value have a incorrect format",
			Code: http.StatusBadRequest,
		}}
	}

	err = event.Update(dbConnection)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	result := NewResult([]models.Event{event})
	data, err := result.ParseToJSON()
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	return data, Error{}
}

func deleteEvent(dbConnection *gorm.DB, w http.ResponseWriter, r *http.Request) ([]byte, Error) {
	event := models.Event{}

	id := r.FormValue("id")
	if id == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	var err error
	event.ID, err = strconv.Atoi(id)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	err = event.Delete(dbConnection)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	result := NewResult([]models.Event{event})
	data, err := result.ParseToJSON()
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	return data, Error{}
}

func eventsForDay(dbConnection *gorm.DB, w http.ResponseWriter, r *http.Request) ([]byte, Error) {
	strUserID := r.FormValue("user_id")
	if strUserID == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	values := r.URL.Query()
	date, err := time.Parse("2006-01-02", values.Get("e_date"))
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'e_date' value have a incorrect format",
			Code: http.StatusBadRequest,
		}}
	}

	events, err := models.SelectEventsByDay(dbConnection, userID, date)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	result := NewResult(events)
	data, err := result.ParseToJSON()
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	return data, Error{}
}

func eventsForWeek(dbConnection *gorm.DB, w http.ResponseWriter, r *http.Request) ([]byte, Error) {
	strUserID := r.FormValue("user_id")
	if strUserID == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	values := r.URL.Query()
	date, err := time.Parse("2006-01-02", values.Get("e_date"))
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'e_date' value have a incorrect format",
			Code: http.StatusBadRequest,
		}}
	}

	events, err := models.SelectEventsByWeek(dbConnection, userID, date)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	result := NewResult(events)
	data, err := result.ParseToJSON()
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	return data, Error{}
}

func eventsForMonth(dbConnection *gorm.DB, w http.ResponseWriter, r *http.Request) ([]byte, Error) {
	strUserID := r.FormValue("user_id")
	if strUserID == "" {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' value couldn't be empty",
			Code: http.StatusBadRequest,
		}}
	}
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'user_id' must be integer",
			Code: http.StatusBadRequest,
		}}
	}

	values := r.URL.Query()
	date, err := time.Parse("2006-01-02", values.Get("e_date"))
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: "'e_date' value have a incorrect format",
			Code: http.StatusBadRequest,
		}}
	}

	events, err := models.SelectEventsByMonth(dbConnection, userID, date)
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	result := NewResult(events)
	data, err := result.ParseToJSON()
	if err != nil {
		return nil, Error{Err: ErrorInfo{
			Text: err.Error(),
			Code: http.StatusServiceUnavailable,
		}}
	}

	return data, Error{}
}

// Result хранит слайс Event.
type Result struct {
	Result []models.Event `json:"result"`
}

// NewResult получает слайс Event, помещает его в структуру Result и возвращает ее.
func NewResult(e []models.Event) Result {
	return Result{
		Result: e,
	}
}

// ParseToJSON преобразовывает Result в слайс байт.
func (r *Result) ParseToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		log.Printf("unable parse to json: %v\n", err)
		return nil, err
	}

	return data, nil
}

// Error хранит ErrorInfo с информацией об ошибке.
type Error struct {
	Err ErrorInfo `json:"error"`
}

// ErrorInfo хранит информацию об ошибке: текст и код.
type ErrorInfo struct {
	Text string `json:"text"`
	Code int    `json:"code"`
}
