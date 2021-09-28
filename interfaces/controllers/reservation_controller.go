package controllers

import (
	"strconv"

	"olive/domain"
	"olive/interfaces/database"
	"olive/usecase"
)

type ReservationController struct {
	Interactor usecase.ReservationInteractor
}

func NewReservationController(sqlHandler database.SqlHandler) *ReservationController {
	return &ReservationController{
		Interactor: usecase.ReservationInteractor{
			ReservationRepository: &database.ReservationRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ReservationController) CreateReservation(c Context) (err error) {
	u := domain.Reservation{}
	c.Bind(&u)
	reservation, err := controller.Interactor.Add(u)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, reservation)
	return
}

func (controller *ReservationController) GetReservations(c Context) (err error) {
	reservations, err := controller.Interactor.Reservations()

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, reservations)
	return
}

func (controller *ReservationController) GetReservation(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation, err := controller.Interactor.ReservationById(id)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, reservation)
	return
}

func (controller *ReservationController) UpdateReservation(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := domain.Reservation{ID: id}
	c.Bind(&u)

	reservation, err := controller.Interactor.Update(u)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, reservation)
	return
}

func (controller *ReservationController) DeleteReservation(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	reservation := domain.Reservation{ID: id}

	err = controller.Interactor.DeleteById(reservation)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, reservation)
	return
}
