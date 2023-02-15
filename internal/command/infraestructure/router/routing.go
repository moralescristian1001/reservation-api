package router

import (
	"github.com/gin-gonic/gin"
	"reservation-api/internal/command/infraestructure/entrypoints/handlers"
)

func RegisterCommandEndpoints(engine *gin.Engine, reservationHandler *handlers.ReservationHandler) {
	engine.POST("/v1/reservation", func(c *gin.Context) {
		reservationHandler.HandleReservation(c)
	})
}
