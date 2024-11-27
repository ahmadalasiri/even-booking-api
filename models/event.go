package models

import (
	"event-booking-api/db"
	"time"
)

type Event struct {
	ID       int
	Name     string `binding:"required"`
	DateTime time.Time
	UserID   int
}

func (e Event) Save() (Event, error) {
	query := `
		INSERT INTO events (name, user_id)
		VALUES ($1, $2)
		RETURNING id, name, date_time, user_id
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return Event{}, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(e.Name, e.UserID).Scan(&e.ID, &e.Name, &e.DateTime, &e.UserID)

	if err != nil {
		return Event{}, err
	}

	return e, nil

}

func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id string) (Event, error) {
	query := `SELECT * FROM events WHERE id = $1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return Event{}, err
	}
	defer stmt.Close()

	var event Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.DateTime, &event.UserID)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func (e Event) UpdateEvent(id string) (Event, error) {
	query := `
		UPDATE events
		SET name = $1
		WHERE id = $2
		RETURNING id, name, date_time, user_id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return Event{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(e.Name, id).Scan(&e.ID, &e.Name, &e.DateTime, &e.UserID)
	if err != nil {
		return Event{}, err
	}
	return e, nil
}

func DeleteEvent(id string) error {
	query := `DELETE FROM events WHERE id = $1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
