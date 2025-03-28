package infraestructure

import (
	"Rshortylong/Persona/domain"
	"Rshortylong/core"
	"fmt"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones:", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql MySQL) Create(persona *domain.Persona) (uint, error) {
	query := `
		INSERT INTO persona 
		(edad, nombre, sexo, genero) 
		VALUES (?, ?, ?, ?)
	`

	res, err := mysql.conn.ExecutePreparedQuery(
		query,
		persona.Edad,
		persona.Nombre,
		persona.Sexo,
		persona.Genero,
	)
	
	if err != nil {
		return 0, fmt.Errorf("error creating persona: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %v", err)
	}

	return uint(id), nil
}