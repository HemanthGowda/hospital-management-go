package patient

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/HemanthGowda/hospital-management-go/internal/db"
)

const patientColumns = `
	id::text, full_name, primary_doctor_name, medical_record_number, date_of_birth,
	gender, phone, blood_group, notes, created_at, updated_at
`

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) FindByID(ctx context.Context, id string) (Patient, error) {
	if _, err := uuid.Parse(id); err != nil {
		return Patient{}, ErrNotFound
	}

	row := r.pool.QueryRow(ctx, `
		SELECT `+patientColumns+`
		FROM patients
		WHERE id = $1
	`, id)
	return scanPatient(row)
}

func (r *PostgresRepository) MedicalRecordNumberUsedByOther(ctx context.Context, medicalRecordNumber, id string) (bool, error) {
	var used bool
	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(
			SELECT 1 FROM patients
			WHERE medical_record_number = $1 AND id <> $2
		)
	`, medicalRecordNumber, id).Scan(&used)
	return used, err
}

func (r *PostgresRepository) Update(ctx context.Context, patient Patient) (Patient, error) {
	row := r.pool.QueryRow(ctx, `
		UPDATE patients
		SET full_name = $2,
		    primary_doctor_name = $3,
		    medical_record_number = $4,
		    date_of_birth = $5,
		    gender = $6,
		    phone = $7,
		    blood_group = $8,
		    notes = $9,
		    updated_at = now()
		WHERE id = $1
		RETURNING `+patientColumns+`
	`, patient.ID, patient.FullName, patient.PrimaryDoctorName, patient.MedicalRecordNumber,
		patient.DateOfBirth, patient.Gender, nullable(patient.Phone), nullable(patient.BloodGroup), nullable(patient.Notes))

	updated, err := scanPatient(row)
	if errors.Is(err, ErrNotFound) {
		return Patient{}, ErrNotFound
	}
	if db.IsUniqueViolation(err) {
		return Patient{}, ErrDuplicate
	}
	return updated, err
}

func (r *PostgresRepository) ExistsByMedicalRecordNumber(ctx context.Context, medicalRecordNumber string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM patients WHERE medical_record_number = $1)
	`, medicalRecordNumber).Scan(&exists)
	return exists, err
}

func (r *PostgresRepository) Insert(ctx context.Context, patient Patient) (Patient, error) {
	if patient.ID == "" {
		patient.ID = uuid.NewString()
	}

	row := r.pool.QueryRow(ctx, `
		INSERT INTO patients (
			id, full_name, primary_doctor_name, medical_record_number, date_of_birth,
			gender, phone, blood_group, notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING `+patientColumns+`
	`, patient.ID, patient.FullName, patient.PrimaryDoctorName, patient.MedicalRecordNumber,
		patient.DateOfBirth, patient.Gender, nullable(patient.Phone), nullable(patient.BloodGroup), nullable(patient.Notes))

	inserted, err := scanPatient(row)
	if db.IsUniqueViolation(err) {
		return Patient{}, ErrDuplicate
	}
	return inserted, err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanPatient(row scanner) (Patient, error) {
	var patient Patient
	var phone, bloodGroup, notes *string
	err := row.Scan(
		&patient.ID,
		&patient.FullName,
		&patient.PrimaryDoctorName,
		&patient.MedicalRecordNumber,
		&patient.DateOfBirth,
		&patient.Gender,
		&phone,
		&bloodGroup,
		&notes,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return Patient{}, ErrNotFound
	}
	if err != nil {
		return Patient{}, err
	}
	patient.Phone = deref(phone)
	patient.BloodGroup = deref(bloodGroup)
	patient.Notes = deref(notes)
	patient.DateOfBirth = patient.DateOfBirth.UTC()
	patient.CreatedAt = patient.CreatedAt.UTC()
	patient.UpdatedAt = patient.UpdatedAt.UTC()
	return patient, nil
}

func deref(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func nullable(value string) any {
	if value == "" {
		return nil
	}
	return value
}
