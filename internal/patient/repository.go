package patient

import (
	"context"
	"errors"
)

var (
	ErrNotFound  = errors.New("patient not found")
	ErrDuplicate = errors.New("patient already exists")
)

type Repository interface {
	FindByID(ctx context.Context, id string) (Patient, error)
	MedicalRecordNumberUsedByOther(ctx context.Context, medicalRecordNumber, id string) (bool, error)
	Update(ctx context.Context, patient Patient) (Patient, error)
	ExistsByMedicalRecordNumber(ctx context.Context, medicalRecordNumber string) (bool, error)
	Insert(ctx context.Context, patient Patient) (Patient, error)
}
