package application

import(
	"Rshortylong/Persona/domain"
)
	 


type CreatePersona struct {
	db domain.IPersona
}

func NewCreatePersona(db domain.IPersona)*CreatePersona{
	return &CreatePersona{db: db}
}

func ( uc *CreatePersona) CreatePersona(persona domain.Persona)(uint, error){
	return uc.db.Create(&persona)
}