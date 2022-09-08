package server

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func handling(dbConnection *gorm.DB) {
	http.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			sendError(w, r, http.StatusBadRequest, "/create_event available only for POST method")
			return
		}

		data, err := createEvent(dbConnection, w, r)
		if err.Err.Code != 0 {
			sendError(w, r, err.Err.Code, err.Err.Text)
		}

		sendData(w, r, data)
	})

	http.HandleFunc("/update_event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			sendError(w, r, http.StatusBadRequest, "/update_event available only for POST method")
			return
		}

		data, err := updateEvent(dbConnection, w, r)
		if err.Err.Code != 0 {
			sendError(w, r, err.Err.Code, err.Err.Text)
		}

		sendData(w, r, data)
	})

	http.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			sendError(w, r, http.StatusBadRequest, "/delete_event available only for POST method")
			return
		}

		data, err := deleteEvent(dbConnection, w, r)
		if err.Err.Code != 0 {
			sendError(w, r, err.Err.Code, err.Err.Text)
		}

		sendData(w, r, data)
	})

	http.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			sendError(w, r, http.StatusBadRequest, "/events_for_day available only for GET method")
			return
		}

		data, err := eventsForDay(dbConnection, w, r)
		if err.Err.Code != 0 {
			sendError(w, r, err.Err.Code, err.Err.Text)
		}

		sendData(w, r, data)
	})

	http.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			sendError(w, r, http.StatusBadRequest, "/events_for_week available only for GET method")
			return
		}

		data, err := eventsForWeek(dbConnection, w, r)
		if err.Err.Code != 0 {
			sendError(w, r, err.Err.Code, err.Err.Text)
		}

		sendData(w, r, data)
	})

	http.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			sendError(w, r, http.StatusBadRequest, "/events_for_month available only for GET method")
			return
		}

		data, err := eventsForMonth(dbConnection, w, r)
		if err.Err.Code != 0 {
			sendError(w, r, err.Err.Code, err.Err.Text)
		}

		sendData(w, r, data)
	})
}

// sendData отправляет клиенту запрошенные данные.
func sendData(w http.ResponseWriter, r *http.Request, data []byte) {
	log.Printf("[INFO] %s %v: params: %v\n", r.Method, r.URL.String(), r.Form)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
}

// sendError отправляет клиенту информацию об ошибке.
func sendError(w http.ResponseWriter, r *http.Request, errCode int, errText string) {
	log.Printf("[WARNING] code: %v error: %v request: %v\n",
		errCode, errText, r.URL.String())
	errInfo := map[string]interface{}{
		"error": map[string]interface{}{
			"text": errText,
			"code": errCode,
		},
	}

	data, err := json.Marshal(errInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
