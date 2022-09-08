package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// Event хранит информацию о событии.
type Event struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"column:user_id" json:"user_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:descr" json:"descr"`
	Date        time.Time `gorm:"column:e_date" json:"e_date"`
}

// Create добавляет запись в таблицу events.
func (e *Event) Create(dbConnection *gorm.DB) error {
	err := dbConnection.Create(e).Error
	if err != nil {
		log.Printf("unable create new data in table 'events': %v\n", err)
		return err
	}
	return nil
}

// Update обновляет существующую запись в таблице events.
func (e *Event) Update(dbConnection *gorm.DB) error {
	err := dbConnection.Model(e).Updates(Event{
		ID:          e.ID,
		UserID:      e.UserID,
		Date:        e.Date,
		Title:       e.Title,
		Description: e.Description,
	}).Error
	if err != nil {
		log.Printf("unable update id='%d' field in table 'events': %v\n",
			e.ID, err)
		return err
	}
	return nil
}

// Delete удаляет существующую запись в таблице events.
func (e *Event) Delete(dbConnection *gorm.DB) error {
	err := dbConnection.Unscoped().Find(e).Delete(e).Error
	if err != nil {
		log.Printf("unable delete id='%d' field in table 'events': %v\n",
			e.ID, err)
		return err
	}
	return nil
}

// SelectEventsByDay возвращает слайс Event за указанную date.
func SelectEventsByDay(dbConnection *gorm.DB, userID int, date time.Time) ([]Event, error) {
	events, err := selectEventsByDateInterval(dbConnection, userID,
		date, date)
	if err != nil {
		log.Printf("unable select events for '%v' day: %v\n", date, err.Error())
		return nil, err
	}

	return events, nil
}

// SelectEventsByWeek возвращает слайс Event, входящих в интервал
// между beginDate и beginDate + 7 дней (включительно).
func SelectEventsByWeek(dbConnection *gorm.DB, userID int, beginDate time.Time) ([]Event, error) {
	endDate := beginDate.AddDate(0, 0, 7)
	events, err := selectEventsByDateInterval(dbConnection, userID,
		beginDate, endDate)
	if err != nil {
		log.Printf("unable select events for '%v'-'%v' date interval: %v\n",
			beginDate, endDate, err.Error())
		return nil, err
	}

	return events, nil
}

// SelectEventsByMonth возвращает слайс Event, входящиз в интервал
// между beginDate и beginDate + 1 месяц (включительно).
func SelectEventsByMonth(dbConnection *gorm.DB, userID int, beginDate time.Time) ([]Event, error) {
	endDate := beginDate.AddDate(0, 1, 0)
	events, err := selectEventsByDateInterval(dbConnection, userID,
		beginDate, endDate)
	if err != nil {
		log.Printf("unable select events for '%v'-'%v' date interval: %v\n",
			beginDate, endDate, err.Error())
		return nil, err
	}

	return events, nil
}

// selectEventsByDateInterval находит в таблице events записи за указанный интервал e_date
// и возвращает слайс структур Event с данными из этих записей.
func selectEventsByDateInterval(dbConnection *gorm.DB, userID int, begin, end time.Time) ([]Event, error) {
	events := []Event{}
	err := dbConnection.Where("user_id = ? and e_date >= ? and e_date <= ?",
		userID, begin, end).Find(&events).Error
	if err != nil {
		log.Printf("unable select fields from table 'events': %v\n", err)
		return nil, err
	}

	return events, nil
}
