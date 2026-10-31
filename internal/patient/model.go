package patient

import "time"

type Patient struct {
	ID                  string
	FullName            string
	PrimaryDoctorName   string
	MedicalRecordNumber string
	DateOfBirth         time.Time
	Gender              string
	Phone               string
	BloodGroup          string
	Notes               string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type UpdatePatientRequest struct {
	FullName            *string `json:"fullName"`
	PrimaryDoctorName   *string `json:"primaryDoctorName"`
	MedicalRecordNumber *string `json:"medicalRecordNumber"`
	DateOfBirth         *string `json:"dateOfBirth"`
	Gender              *string `json:"gender"`
	Phone               *string `json:"phone"`
	BloodGroup          *string `json:"bloodGroup"`
	Notes               *string `json:"notes"`
}

type PatientResponse struct {
	ID                  string    `json:"id"`
	FullName            string    `json:"fullName"`
	PrimaryDoctorName   string    `json:"primaryDoctorName"`
	MedicalRecordNumber string    `json:"medicalRecordNumber"`
	DateOfBirth         string    `json:"dateOfBirth"`
	Gender              string    `json:"gender"`
	Phone               string    `json:"phone,omitempty"`
	BloodGroup          string    `json:"bloodGroup,omitempty"`
	Notes               string    `json:"notes,omitempty"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
