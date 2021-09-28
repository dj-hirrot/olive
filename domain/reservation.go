package domain

import "time"

type Reservation struct {
	ID            int       `gorm:"primary_key" json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        int       `json:"status"`
	Fee           int       `json:"fee"`
	DepartureDate time.Time `json:"departure_date"`
	ArrivalDate   time.Time `json:"arrival_date"`
	CretaedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type Reservations []Reservation
