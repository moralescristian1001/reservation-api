package router

import (
	"github.com/gin-gonic/gin"
	"reservation-api/internal/query/infraestructure/entrypoints/handlers"
)

func RegisterQueryEndpoints(engine *gin.Engine, reservationHandler *handlers.ReservationHandler) {
	engine.GET("/v1/reservation", func(c *gin.Context) {
		reservationHandler.GetReservations(c)
	})
}
