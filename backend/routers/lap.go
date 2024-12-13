package routers

import (
	"Ariadne_Management/models"
	"Ariadne_Management/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateLapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
			return
		}

		var lap models.Lap
		if err := c.ShouldBindJSON(&lap); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		lap.SessionID = sessionID

		if err := services.CreateLap(db, userID.(int), &lap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create lap"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Lap created"})
	}
}

func GetLapsByUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		laps, err := services.GetLapsByUser(db, userID.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch laps"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"laps": laps})
	}
}

func GetLapByKeyHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		lapNumberStr := c.Param("lap_number")

		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
			return
		}

		lapNumber, err := strconv.Atoi(lapNumberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lap_number"})
			return
		}

		lap, err := services.GetLapByKey(db, userID.(int), sessionID, lapNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lap not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"lap": lap})
	}
}

func UpdateLapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		lapNumberStr := c.Param("lap_number")

		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
			return
		}

		lapNumber, err := strconv.Atoi(lapNumberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lap_number"})
			return
		}

		var data struct {
			LapTime string `json:"lap_time"`
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := services.UpdateLap(db, userID.(int), sessionID, lapNumber, data.LapTime); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update lap"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Lap updated"})
	}
}

func DeleteLapHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		sessionIDStr := c.Param("session_id")
		lapNumberStr := c.Param("lap_number")

		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
			return
		}

		lapNumber, err := strconv.Atoi(lapNumberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lap_number"})
			return
		}

		if err := services.DeleteLap(db, userID.(int), sessionID, lapNumber); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete lap"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Lap deleted"})
	}
}
