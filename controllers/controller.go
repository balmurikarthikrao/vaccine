package controllers

import (
	"database/sql"
	"net/http"
	"vaccine/models"

	"github.com/gin-gonic/gin"
)

type VaccinationController interface {
	CreateBeneficiary(c *gin.Context)
	CreateAppointment(c *gin.Context)
}

type Vaccination struct {
	Db *sql.DB
}

func NewController(db *sql.DB) VaccinationController {
	return &Vaccination{
		Db: db,
	}
}

func (v *Vaccination) CreateBeneficiary(c *gin.Context) {

	var beneficiary models.Beneficiary
	if err := c.ShouldBindJSON(&beneficiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ShouldBindJSON": err.Error()})
		return
	}

	beneficiaryResp, err := models.CreateBeneficiary(v.Db, beneficiary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register beneficiary", "message": err})
	}

	c.JSON(http.StatusOK, beneficiaryResp)
	return
}

func (v *Vaccination) CreateAppointment(c *gin.Context) {

}
