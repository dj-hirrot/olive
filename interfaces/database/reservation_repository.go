package database

import "olive/domain"

type ReservationRepository struct {
	SqlHandler
}

func (rr *ReservationRepository) FindByID(id int) (r domain.Reservation, err error) {
	if err = rr.Find(&r, id).Error; err != nil {
		return
	}
	return
}

func (reservationRepository *ReservationRepository) Store(r domain.Reservation) (reservation domain.Reservation, err error) {
	if err = reservationRepository.Create(&r).Error; err != nil {
		return
	}
	reservation = r
	return
}

func (reservationRepository *ReservationRepository) Update(r domain.Reservation) (reservation domain.Reservation, err error) {
	if err = reservationRepository.Save(&r).Error; err != nil {
		return
	}
	reservation = r
	return
}

func (reservationRepository *ReservationRepository) DeleteByID(reservation domain.Reservation) (err error) {
	if err = reservationRepository.Delete(&reservation).Error; err != nil {
		return
	}
	return
}

func (reservationRepository *ReservationRepository) FindAll() (reservations domain.Reservations, err error) {
	if err = reservationRepository.Find(&reservations).Error; err != nil {
		return
	}
	return
}
