package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"vaccine/models"

	"github.com/gin-gonic/gin"
)

type VaccinationController interface {
	CreateBeneficiary(c *gin.Context)
	CreateAppointment(c *gin.Context)
	Delete(c *gin.Context)
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

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the appointment slot is available
	count, err := models.CheckAppointmentAvailable(v.Db, appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check appointment availability", "message": err})
		return
	}
	if count >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment slot is full"})
		return
	}

	// Check if the beneficiary already has two appointments
	count, err = models.CheckMultiple(v.Db, appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check beneficiary appointments"})
		return
	}

	if count >= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Beneficiary already has two appointments"})
		return
	}

	// Check if the slot is available for the requested dose type
	count, err = models.CheckAppointmentAvailable(v.Db, appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check appointment availability for dose"})
		return
	}

	if count >= 15 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No more appointments available for the requested dose type"})
		return
	}

	// Insert appointment into the database
	result, err := models.InsertAppintment(v.Db, appointment)

	// Get the inserted appointment ID
	appointmentID, _ := result.LastInsertId()
	appointment.ID = int(appointmentID)

	c.JSON(http.StatusOK, appointment)

}

func (v *Vaccination) Delete(c *gin.Context) {

	appointmentID, _ := strconv.Atoi(c.Param("id"))

	// Delete the appointment from the database
	deleteAppointmentQuery := "DELETE FROM appointments WHERE id = ?"
	_, err := v.Db.Exec(deleteAppointmentQuery, appointmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment cancelled successfully"})
	return
}
