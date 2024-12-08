package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/base-project-golang/config"
	"github.com/moura95/base-project-golang/internal/application/service"
	"github.com/moura95/base-project-golang/internal/infrastructure/repository"
	"go.uber.org/zap"
)

func CreateRoutesV1(store *sqlx.DB, cfg *config.Config, router *gin.Engine, log *zap.SugaredLogger) {
	routes := router.Group("/")
	// Instance  Repository Postgres
	repo := repository.NewStorageRepository(store)
	// Instance  Service with Postgres
	sv := service.NewStorageService(repo, *cfg, log)

	// Init all Routers
	NewReceiverRouter(sv, log).SetupContractRoute(routes)

}
