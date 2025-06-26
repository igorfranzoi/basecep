package cep

import (
	"context"

	"github.com/google/uuid"
	"github.com/igorfranzoi/basecep/internal/model"
	"gorm.io/gorm"
)

type EstadoRepository interface {
	Create(ctx context.Context, estado *model.Estado) error
	Update(ctx context.Context, estado *model.Estado) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Estado, error)
	GetByUF(ctx context.Context, uf string) (*model.Estado, error)
}

type estadoRepository struct {
	DB *gorm.DB
}

func NewEstadoRepository(db *gorm.DB) EstadoRepository {
	return &estadoRepository{DB: db}
}

func (r *estadoRepository) Create(ctx context.Context, estado *model.Estado) error {
	return r.DB.WithContext(ctx).Create(estado).Error
}

func (r *estadoRepository) Update(ctx context.Context, estado *model.Estado) error {
	return r.DB.WithContext(ctx).Save(estado).Error
}

func (r *estadoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&model.Estado{}, id).Error
}

func (r *estadoRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Estado, error) {
	var estado model.Estado
	if err := r.DB.WithContext(ctx).First(&estado, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &estado, nil
}

func (r *estadoRepository) GetByUF(ctx context.Context, uf string) (*model.Estado, error) {
	var estado model.Estado
	if err := r.DB.WithContext(ctx).First(&estado, "uf = ?", uf).Error; err != nil {
		return nil, err
	}
	return &estado, nil
}
