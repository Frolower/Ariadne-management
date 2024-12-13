package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateStageHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")

		chIDStr := c.Param("championship_id")
		chID, err := strconv.Atoi(chIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid championship ID"})
			return
		}

		var stage models.Stage
		if err := c.ShouldBindJSON(&stage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		stage.ChampionshipID = chID

		if err := services.CreateStage(db, userID.(int), &stage); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create stage"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Stage created", "stage_id": stage.StageID})
	}
}

func GetStagesByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		stages, err := services.GetStagesByUser(db, userID.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching stages"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"stages": stages})
	}
}

func GetStageByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		stageIDStr := c.Param("stage_id")
		stageID, err := strconv.Atoi(stageIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stage ID"})
			return
		}

		stage, err := services.GetStageByID(db, userID.(int), stageID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stage not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"stage": stage})
	}
}

func UpdateStageHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		stageIDStr := c.Param("stage_id")
		stageID, err := strconv.Atoi(stageIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stage ID"})
			return
		}

		var st models.Stage
		if err := c.ShouldBindJSON(&st); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := services.UpdateStage(db, userID.(int), stageID, &st); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update stage"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Stage updated"})
	}
}

func DeleteStageHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		stageIDStr := c.Param("stage_id")
		stageID, err := strconv.Atoi(stageIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stage ID"})
			return
		}

		if err := services.DeleteStage(db, userID.(int), stageID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete stage"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Stage deleted"})
	}
}