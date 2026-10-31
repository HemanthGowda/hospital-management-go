package seed

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/HemanthGowda/hospital-management-go/internal/config"
	"github.com/HemanthGowda/hospital-management-go/internal/password"
	"github.com/HemanthGowda/hospital-management-go/internal/patient"
	"github.com/HemanthGowda/hospital-management-go/internal/user"
)

func Run(ctx context.Context, cfg config.Config, users user.Repository, patients patient.Repository, passwords password.Encoder) error {
	if err := seedAdmin(ctx, cfg, users, passwords); err != nil {
		return err
	}
	return seedPatients(ctx, patients)
}

func seedAdmin(ctx context.Context, cfg config.Config, users user.Repository, passwords password.Encoder) error {
	email := strings.ToLower(strings.TrimSpace(cfg.SeedAdminEmail))
	exists, err := users.ExistsByEmail(ctx, email)
	if err != nil || exists {
		return err
	}

	hash, err := passwords.Encode(cfg.SeedAdminPassword)
	if err != nil {
		return err
	}

	_, err = users.Create(ctx, user.User{
		Email:        email,
		PasswordHash: hash,
		FullName:     strings.TrimSpace(cfg.SeedAdminFullName),
		Roles:        []user.Role{user.RoleAdmin},
		Enabled:      true,
	})
	if errors.Is(err, user.ErrDuplicate) {
		return nil
	}
	return err
}

func seedPatients(ctx context.Context, patients patient.Repository) error {
	samples := []patient.Patient{
		{
			FullName:            "Ada Lovelace",
			PrimaryDoctorName:   "Dr. Grace Hopper",
			MedicalRecordNumber: "MRN-100001",
			DateOfBirth:         time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
			Gender:              "FEMALE",
			Phone:               "+15551230001",
			BloodGroup:          "O+",
			Notes:               "Reference patient used by the starter seed.",
		},
		{
			FullName:            "Alan Turing",
			PrimaryDoctorName:   "Dr. Florence Nightingale",
			MedicalRecordNumber: "MRN-100002",
			DateOfBirth:         time.Date(1912, 6, 23, 0, 0, 0, 0, time.UTC),
			Gender:              "MALE",
			BloodGroup:          "AB+",
			Notes:               "Reference patient used by the starter seed.",
		},
		{
			FullName:            "Katherine Johnson",
			PrimaryDoctorName:   "Dr. Jonas Salk",
			MedicalRecordNumber: "MRN-100003",
			DateOfBirth:         time.Date(1918, 8, 26, 0, 0, 0, 0, time.UTC),
			Gender:              "FEMALE",
			BloodGroup:          "A-",
			Notes:               "Reference patient used by the starter seed.",
		},
	}

	for _, sample := range samples {
		exists, err := patients.ExistsByMedicalRecordNumber(ctx, sample.MedicalRecordNumber)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		if _, err := patients.Insert(ctx, sample); err != nil && !errors.Is(err, patient.ErrDuplicate) {
			return err
		}
	}
	return nil
}
