package model

import (
	"time"

	"github.com/igorfranzoi/basecep/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID        uuid.UUID             `json:"id" gorm:"primary_key"`            //ID de identificação do registro
	CreatedAt time.Time             `json:"created_at" gorm:"autoCreateTime"` //Timestamp de criação do registro
	UpdatedAt time.Time             `json:"updated_at" gorm:"autoUpdateTime"` //Timestamp de atualização do registro
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"softDelete:nano"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Model) BeforeCreate(tx *gorm.DB) error {

	utils.ReturnUUID(base)

	return nil
}
