package cep

import (
	"context"

	"github.com/google/uuid"
	"github.com/igorfranzoi/basecep/internal/model"
	"gorm.io/gorm"
)

type LogradouroRepository interface {
	Create(ctx context.Context, logradouro *model.Logradouro) error
	Update(ctx context.Context, logradouro *model.Logradouro) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Logradouro, error)
}

type logradouroRepository struct {
	DB *gorm.DB
}

func NewLogradouroRepository(db *gorm.DB) LogradouroRepository {
	return &logradouroRepository{DB: db}
}

func (r *logradouroRepository) Create(ctx context.Context, logradouro *model.Logradouro) error {
	return r.DB.WithContext(ctx).Create(logradouro).Error
}

func (r *logradouroRepository) Update(ctx context.Context, logradouro *model.Logradouro) error {
	return r.DB.WithContext(ctx).Save(logradouro).Error
}

func (r *logradouroRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&model.Logradouro{}, id).Error
}

func (r *logradouroRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Logradouro, error) {
	var logradouro model.Logradouro
	if err := r.DB.WithContext(ctx).First(&logradouro, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &logradouro, nil
}
