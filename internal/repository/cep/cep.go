package cep

import (
	"context"

	"github.com/google/uuid"
	"github.com/igorfranzoi/basecep/internal/model"
	"gorm.io/gorm"
)

type CEPRepository interface {
	Create(ctx context.Context, cep *model.CEP) error
	Update(ctx context.Context, cep *model.CEP) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.CEP, error)
	GetByCode(ctx context.Context, codigo string) (*model.CEP, error)
}

type cepRepository struct {
	DB *gorm.DB
}

func NewCEPRepository(db *gorm.DB) CEPRepository {
	return &cepRepository{DB: db}
}

func (r *cepRepository) Create(ctx context.Context, cep *model.CEP) error {
	return r.DB.WithContext(ctx).Create(cep).Error
}

func (r *cepRepository) Update(ctx context.Context, cep *model.CEP) error {
	return r.DB.WithContext(ctx).Save(cep).Error
}

func (r *cepRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&model.CEP{}, id).Error
}

func (r *cepRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.CEP, error) {
	var cep model.CEP
	if err := r.DB.WithContext(ctx).First(&cep, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cep, nil
}

func (r *cepRepository) GetByCode(ctx context.Context, codigo string) (*model.CEP, error) {
	var cep model.CEP
	if err := r.DB.WithContext(ctx).Preload("Logradouros").First(&cep, "codigo = ?", codigo).Error; err != nil {
		return nil, err
	}
	return &cep, nil
}
