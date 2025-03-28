package domain

type IPersona interface{
	Create(persona *Persona)(uint, error)
}