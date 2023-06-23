package router

import (
	"database/sql"
	"vaccine/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter will handle all routes
func NewRouter(db *sql.DB) *gin.Engine {

	// set server mode
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	vaccineController := controllers.NewController(db)

	v1 := router.Group("/v1")
	{
		v1.POST("/beneficiaries", vaccineController.CreateBeneficiary)
		v1.POST("/appointments", vaccineController.CreateAppointment)
		v1.DELETE("/appointments/:id", vaccineController.Delete)

	}

	return router
}
