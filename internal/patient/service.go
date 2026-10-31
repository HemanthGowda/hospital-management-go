package patient

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/HemanthGowda/hospital-management-go/internal/domain"
)

var (
	mrnPattern        = regexp.MustCompile(`^MRN-\d{6}$`)
	bloodGroupPattern = regexp.MustCompile(`^(A|B|AB|O)[+-]$`)
	phonePattern      = regexp.MustCompile(`^[0-9+() .\-]{7,20}$`)
	genders           = map[string]struct{}{
		"FEMALE":  {},
		"MALE":    {},
		"OTHER":   {},
		"UNKNOWN": {},
	}
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Update(ctx context.Context, patientID string, request UpdatePatientRequest) (PatientResponse, error) {
	if violations := validateUpdate(request); len(violations) > 0 {
		return PatientResponse{}, domain.ValidationFailed(violations)
	}

	existing, err := s.repository.FindByID(ctx, patientID)
	if errors.Is(err, ErrNotFound) {
		return PatientResponse{}, domain.NotFound("Patient not found: " + patientID)
	}
	if err != nil {
		return PatientResponse{}, err
	}

	medicalRecordNumber := normalizeMRN(*request.MedicalRecordNumber)
	taken, err := s.repository.MedicalRecordNumberUsedByOther(ctx, medicalRecordNumber, patientID)
	if err != nil {
		return PatientResponse{}, err
	}
	if taken {
		return PatientResponse{}, domain.Conflict("Medical record number is already assigned to another patient")
	}

	dateOfBirth, err := time.Parse("2006-01-02", strings.TrimSpace(*request.DateOfBirth))
	if err != nil {
		return PatientResponse{}, domain.ValidationFailed([]domain.Violation{{
			Field:   "dateOfBirth",
			Message: "Date of birth must be a valid date in YYYY-MM-DD format",
		}})
	}
	if dateOfBirth.After(today()) {
		return PatientResponse{}, domain.BadRequest("Date of birth cannot be in the future")
	}

	existing.FullName = strings.TrimSpace(*request.FullName)
	existing.PrimaryDoctorName = strings.TrimSpace(*request.PrimaryDoctorName)
	existing.MedicalRecordNumber = medicalRecordNumber
	existing.DateOfBirth = dateOfBirth
	existing.Gender = strings.ToUpper(strings.TrimSpace(*request.Gender))
	existing.Phone = ""
	if request.Phone != nil {
		existing.Phone = strings.TrimSpace(*request.Phone)
	}
	existing.BloodGroup = ""
	if request.BloodGroup != nil {
		existing.BloodGroup = strings.ToUpper(strings.TrimSpace(*request.BloodGroup))
	}
	existing.Notes = ""
	if request.Notes != nil {
		existing.Notes = strings.TrimSpace(*request.Notes)
	}

	updated, err := s.repository.Update(ctx, existing)
	if errors.Is(err, ErrNotFound) {
		return PatientResponse{}, domain.NotFound("Patient not found: " + patientID)
	}
	if errors.Is(err, ErrDuplicate) {
		return PatientResponse{}, domain.Conflict("Medical record number is already assigned to another patient")
	}
	if err != nil {
		return PatientResponse{}, err
	}
	return toResponse(updated), nil
}

func validateUpdate(request UpdatePatientRequest) []domain.Violation {
	var violations []domain.Violation

	switch {
	case request.FullName == nil || strings.TrimSpace(*request.FullName) == "":
		violations = append(violations, domain.Violation{Field: "fullName", Message: "Full name is required"})
	case utf8.RuneCountInString(strings.TrimSpace(*request.FullName)) < 2 || utf8.RuneCountInString(strings.TrimSpace(*request.FullName)) > 160:
		violations = append(violations, domain.Violation{Field: "fullName", Message: "Full name must be between 2 and 160 characters"})
	}

	switch {
	case request.PrimaryDoctorName == nil || strings.TrimSpace(*request.PrimaryDoctorName) == "":
		violations = append(violations, domain.Violation{Field: "primaryDoctorName", Message: "Primary doctor name is required"})
	case utf8.RuneCountInString(strings.TrimSpace(*request.PrimaryDoctorName)) > 160:
		violations = append(violations, domain.Violation{Field: "primaryDoctorName", Message: "Primary doctor name must be between 1 and 160 characters"})
	}

	switch {
	case request.MedicalRecordNumber == nil || strings.TrimSpace(*request.MedicalRecordNumber) == "":
		violations = append(violations, domain.Violation{Field: "medicalRecordNumber", Message: "Medical record number is required"})
	case !mrnPattern.MatchString(normalizeMRN(*request.MedicalRecordNumber)):
		violations = append(violations, domain.Violation{Field: "medicalRecordNumber", Message: "Medical record number must look like MRN-100001"})
	}

	switch {
	case request.DateOfBirth == nil || strings.TrimSpace(*request.DateOfBirth) == "":
		violations = append(violations, domain.Violation{Field: "dateOfBirth", Message: "Date of birth is required"})
	default:
		if _, err := time.Parse("2006-01-02", strings.TrimSpace(*request.DateOfBirth)); err != nil {
			violations = append(violations, domain.Violation{Field: "dateOfBirth", Message: "Date of birth must be a valid date in YYYY-MM-DD format"})
		}
	}

	switch {
	case request.Gender == nil || strings.TrimSpace(*request.Gender) == "":
		violations = append(violations, domain.Violation{Field: "gender", Message: "Gender is required"})
	default:
		if _, ok := genders[strings.ToUpper(strings.TrimSpace(*request.Gender))]; !ok {
			violations = append(violations, domain.Violation{Field: "gender", Message: "Gender must be FEMALE, MALE, OTHER, or UNKNOWN"})
		}
	}

	if request.Phone != nil && strings.TrimSpace(*request.Phone) != "" && !phonePattern.MatchString(strings.TrimSpace(*request.Phone)) {
		violations = append(violations, domain.Violation{Field: "phone", Message: "Phone must be between 7 and 20 characters"})
	}

	if request.BloodGroup != nil && strings.TrimSpace(*request.BloodGroup) != "" && !bloodGroupPattern.MatchString(strings.ToUpper(strings.TrimSpace(*request.BloodGroup))) {
		violations = append(violations, domain.Violation{Field: "bloodGroup", Message: "Blood group must be one of A+, A-, B+, B-, AB+, AB-, O+, O-"})
	}

	if request.Notes != nil && utf8.RuneCountInString(*request.Notes) > 2000 {
		violations = append(violations, domain.Violation{Field: "notes", Message: "Notes must be at most 2000 characters"})
	}

	return violations
}

func normalizeMRN(value string) string {
	compact := strings.ReplaceAll(strings.TrimSpace(value), " ", "")
	return strings.ToUpper(compact)
}

func today() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func toResponse(patient Patient) PatientResponse {
	return PatientResponse{
		ID:                  patient.ID,
		FullName:            patient.FullName,
		PrimaryDoctorName:   patient.PrimaryDoctorName,
		MedicalRecordNumber: patient.MedicalRecordNumber,
		DateOfBirth:         patient.DateOfBirth.UTC().Format("2006-01-02"),
		Gender:              patient.Gender,
		Phone:               patient.Phone,
		BloodGroup:          patient.BloodGroup,
		Notes:               patient.Notes,
		CreatedAt:           patient.CreatedAt.UTC(),
		UpdatedAt:           patient.UpdatedAt.UTC(),
	}
}
