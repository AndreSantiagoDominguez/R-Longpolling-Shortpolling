package controllers

import (
	"Rshortylong/Persona/application"
	"Rshortylong/Persona/domain"
	"Rshortylong/Persona/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePersonaController struct {
	app *application.CreatePersona
}

func NewCreatePersonaController() *CreatePersonaController {
	mysql := infraestructure.GetMySQL()
	app := application.NewCreatePersona(mysql)
	return &CreatePersonaController{app: app}
}

func (cp_c *CreatePersonaController) AddPersona(c *gin.Context) {
	var persona domain.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	id, err := cp_c.app.CreatePersona(persona)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Error creando persona: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"type": "persona",
			"id":   id,
			"attributes": gin.H{
				"nombre": persona.Nombre,
				"edad":   persona.Edad,
				"sexo":   persona.Sexo,
				"genero": persona.Genero,
			},
		},
	})
}