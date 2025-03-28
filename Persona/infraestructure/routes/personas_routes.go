package routes

import (
	"Rshortylong/Persona/infraestructure/adapters/http/middleware"
	"Rshortylong/Persona/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPersonaRoutes(r *gin.Engine) {
	personas := r.Group("/personas")
	personas.Use(middleware.CorsMiddleware())
	{
		personas.POST("", controllers.AddPersona)
	
	}
}