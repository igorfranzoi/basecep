package cep

import (
	"context"

	"github.com/google/uuid"
	"github.com/igorfranzoi/basecep/internal/model"
	"gorm.io/gorm"
)

type CityRepository interface {
	Create(ctx context.Context, cidade *model.Cidade) error
	Update(ctx context.Context, cidade *model.Cidade) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Cidade, error)
}

type CidadeRepository struct {
	DB *gorm.DB
}

func NewCidadeRepository(db *gorm.DB) *CidadeRepository {
	return &CidadeRepository{DB: db}
}

func (r *CidadeRepository) Create(ctx context.Context, cidade *model.Cidade) error {
	return r.DB.WithContext(ctx).Create(cidade).Error
}

func (r *CidadeRepository) Update(ctx context.Context, cidade *model.Cidade) error {
	return r.DB.WithContext(ctx).Save(cidade).Error
}

func (r *CidadeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&model.Cidade{}, id).Error
}

func (r *CidadeRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Cidade, error) {
	var cidade model.Cidade

	if err := r.DB.WithContext(ctx).First(&cidade, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cidade, nil
}
