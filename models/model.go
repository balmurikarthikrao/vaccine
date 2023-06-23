package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Beneficiary struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DOB         string    `json:"dob"`
	SSN         int       `json:"ssn"`
	PhoneNumber int       `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Appointment struct {
	ID            int       `json:"id"`
	BeneficiaryID int       `json:"beneficiaryId"`
	Date          string    `json:"date"`
	TimeSlot      string    `json:"timeSlot"`
	DoseType      string    `json:"doseType"`
	VaccineCenter string    `json:"vaccineCenter"`
	CreatedAt     time.Time `json:"createdAt"`
}

func CreateBeneficiary(db *sql.DB, beneficiary Beneficiary) (Beneficiary, error) {

	time := time.Now()
	insertBeneficiaryQuery := "INSERT INTO beneficiaries (name, dob, ssn, phone, created_at) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(insertBeneficiaryQuery, beneficiary.Name, beneficiary.DOB, beneficiary.SSN, beneficiary.PhoneNumber, time)
	if err != nil {
		return Beneficiary{}, fmt.Errorf("error db.Exec %s", err)
	}

	// Get the inserted beneficiary ID
	beneficiaryID, err := result.LastInsertId()
	if err != nil {
		return Beneficiary{}, fmt.Errorf("error LastInsertId %s", err)
	}

	beneficiary.ID = int(beneficiaryID)
	beneficiary.CreatedAt = time
	return beneficiary, nil
}
