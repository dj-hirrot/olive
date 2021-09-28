package infrastructure

import (
	"olive/interfaces/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	reservationController := controllers.NewReservationController(NewMySqlDb())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error { return c.JSON(200, "I am server running...") })
	e.GET("/reservations", func(c echo.Context) error { return reservationController.GetReservations(c) })
	e.GET("/reservations/:id", func(c echo.Context) error { return reservationController.GetReservation(c) })
	e.POST("/reservations", func(c echo.Context) error { return reservationController.CreateReservation(c) })
	e.PUT("/reservations/:id", func(c echo.Context) error { return reservationController.UpdateReservation(c) })
	e.DELETE("/reservations/:id", func(c echo.Context) error { return reservationController.DeleteReservation(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
