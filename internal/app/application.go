package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reservation-api/internal/app/config"
	repositoryCommand "reservation-api/internal/command/core/repository"
	commandUseCase "reservation-api/internal/command/core/usecase"
	handlerCommand "reservation-api/internal/command/infraestructure/entrypoints/handlers"
	handlerMapper "reservation-api/internal/command/infraestructure/entrypoints/handlers/mapper"
	routerCommand "reservation-api/internal/command/infraestructure/router"
	commandDAO "reservation-api/internal/command/infraestructure/storage/dao"
	commandMapperDAO "reservation-api/internal/command/infraestructure/storage/mapper"
	repositoryQuery "reservation-api/internal/query/core/repository"
	handlerQuery "reservation-api/internal/query/infraestructure/entrypoints/handlers"
	routerQuery "reservation-api/internal/query/infraestructure/router"
	queryDAO "reservation-api/internal/query/infraestructure/storage/dao"
)

func Start() {
	doWiring()
}

func doWiring()  {
	db := config.ConnectDatabase()
	config.ExecuteMigrations(db)
	engine := gin.New()
	initDependenciesCommand(db, engine)
	initDependenciesQuery(db, engine)
	engine.Run()

}

func initDependenciesCommand(db *gorm.DB, engine *gin.Engine)  {

	reservationMapper := commandMapperDAO.ReservationMapper{}
	reservationDAO := commandDAO.NewPostgresReservationDAO(db, reservationMapper)
	reservationRepository := repositoryCommand.NewReservationRepository(reservationDAO)
	reservationCreateUseCase := commandUseCase.NewCreateUseCase(reservationRepository)
	apiResponseProvider := handlerCommand.ApiResponseProvider{}
	requestMapper := handlerMapper.RequestMapper{}
	reservationHandler := handlerCommand.NewReservationHandler(reservationCreateUseCase, requestMapper, apiResponseProvider)
	routerCommand.RegisterCommandEndpoints(engine, reservationHandler)
}

func initDependenciesQuery(db *gorm.DB, engine *gin.Engine)  {

	reservationDAO := queryDAO.NewPostgresReservationDAO(db)
	reservationRepository := repositoryQuery.NewReservationRepository(reservationDAO)
	apiResponseProvider := handlerQuery.ApiResponseProvider{}
	reservationHandler := handlerQuery.NewReservationHandler(reservationRepository, apiResponseProvider)
	routerQuery.RegisterQueryEndpoints(engine, reservationHandler)

}
