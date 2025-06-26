package cep

import (
	"context"

	"github.com/google/uuid"
	"github.com/igorfranzoi/basecep/internal/model"
	"gorm.io/gorm"
)

type BairroRepository interface {
	Create(ctx context.Context, bairro *model.Bairro) error
	Update(ctx context.Context, bairro *model.Bairro) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Bairro, error)
}

type bairroRepository struct {
	DB *gorm.DB
}

func NewBairroRepository(db *gorm.DB) BairroRepository {
	return &bairroRepository{DB: db}
}

func (r *bairroRepository) Create(ctx context.Context, bairro *model.Bairro) error {
	return r.DB.WithContext(ctx).Create(bairro).Error
}

func (r *bairroRepository) Update(ctx context.Context, bairro *model.Bairro) error {
	return r.DB.WithContext(ctx).Save(bairro).Error
}

func (r *bairroRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&model.Bairro{}, id).Error
}

func (r *bairroRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Bairro, error) {
	var bairro model.Bairro
	if err := r.DB.WithContext(ctx).First(&bairro, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bairro, nil
}
