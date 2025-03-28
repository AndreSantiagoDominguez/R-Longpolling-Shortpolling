package domain

type Persona struct {
	Id     int    `json:"id"`
	Edad   int    `json:"edad"`
	Nombre string `json:"nombre"`
	Sexo   bool   `json:"sexo"`
	Genero string `json:"genero"`
}

func NewPersona(edad int, nombre string, sexo bool, genero string) *Persona {
	return &Persona{
		Id:     0,
		Edad:   edad,
		Nombre: nombre,
		Sexo:   sexo,
		Genero: genero,
	}
}

// Getters
func (p *Persona) GetId() int {
	return p.Id
}

func (p *Persona) GetEdad() int {
	return p.Edad
}

func (p *Persona) GetNombre() string {
	return p.Nombre
}

func (p *Persona) GetSexo() bool {
	return p.Sexo
}

func (p *Persona) GetGenero() string {
	return p.Genero
}

// Setters
func (p *Persona) SetId(id int) {
	p.Id = id
}

func (p *Persona) SetEdad(edad int) {
	p.Edad = edad
}

func (p *Persona) SetNombre(nombre string) {
	p.Nombre = nombre
}

func (p *Persona) SetSexo(sexo bool) {
	p.Sexo = sexo
}

func (p *Persona) SetGenero(genero string) {
	p.Genero = genero
}