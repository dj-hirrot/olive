package usecase

import "olive/domain"

type ReservationInteractor struct {
	ReservationRepository ReservationRepository
}

func (interactor *ReservationInteractor) ReservationById(id int) (reservation domain.Reservation, err error) {
	reservation, err = interactor.ReservationRepository.FindByID(id)
	return
}

func (interactor *ReservationInteractor) Reservations() (reservations domain.Reservations, err error) {
	reservations, err = interactor.ReservationRepository.FindAll()
	return
}

func (interactor *ReservationInteractor) Add(r domain.Reservation) (reservation domain.Reservation, err error) {
	reservation, err = interactor.ReservationRepository.Store(r)
	return
}

func (interactor *ReservationInteractor) Update(r domain.Reservation) (reservation domain.Reservation, err error) {
	reservation, err = interactor.ReservationRepository.Update(r)
	return
}

func (interactor *ReservationInteractor) DeleteById(reservation domain.Reservation) (err error) {
	err = interactor.ReservationRepository.DeleteByID(reservation)
	return
}
