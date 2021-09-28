package usecase

import "olive/domain"

type ReservationRepository interface {
	FindByID(id int) (domain.Reservation, error)
	Store(domain.Reservation) (domain.Reservation, error)
	Update(domain.Reservation) (domain.Reservation, error)
	DeleteByID(domain.Reservation) error
	FindAll() (domain.Reservations, error)
}
