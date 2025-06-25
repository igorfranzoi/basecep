package model

type Estado struct {
	Model
	Nome string `gorm:"size:250;not null"`
	UF   string `gorm:"size:2;unique;not null"`
	//Cidades []Cidade `gorm:"foreignKey:EstadoID"`
}
