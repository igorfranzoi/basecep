package model

type CEP struct {
	Model
	Codigo       string       `gorm:"size:8;unique;not null"`
	Logradouros  []Logradouro `gorm:"many2many:cep_logradouros;"`
	NumeroInicio *string      `gorm:"size:10"`
	NumeroFim    *string      `gorm:"size:10"`
	Complemento  string       `gorm:"size:250"`
}
