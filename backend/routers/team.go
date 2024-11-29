package services

import (
	"Ariadne_Management/models"
	servicies "Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func CreateTeam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var team models.Team
		if err := c.ShouldBindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Delegate the user registration logic to the services layer
		if err := servicies.CreateTeam(db, &team); err != nil {
			log.Printf("Error creating team: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create team"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Team created successfully"})
	}
}