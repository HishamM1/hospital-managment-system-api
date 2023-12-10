package controllers

import (
	"main/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type roomRequest struct {
	Type      string `json:"Type" binding:"required,min=1,max=255" required:"$field is required" min:"$field must be at least 1 characters" max:"$field must be at most 255 characters"`
	Number    string `json:"Number" binding:"required" required:"$field is required"`
	WardboyID uint   `json:"WardboyID" binding:"required" required:"$field is required"`
}

type roomResponse struct {
	roomRequest
	ID      uint           `json:"ID"`
	Wardboy models.Wardboy `json:"Wardboy"`
}

var room_val = g.Validator(roomRequest{})

func CreateRoom(c *gin.Context) {
	var data roomRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": room_val.DecryptErrors(err)})
		return
	}

	// Check if wardboy exists
	var wardboy models.Wardboy

	result := db.First(&wardboy, data.WardboyID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Wardboy not found"})
		return
	}

	room := models.Room{}
	room.Type = data.Type
	room.Number = data.Number
	room.WardboyID = data.WardboyID
	room.Wardboy = &wardboy

	result = db.Create(&room)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response roomResponse
	response.ID = room.ID
	response.Type = room.Type
	response.Number = room.Number
	response.WardboyID = room.WardboyID
	response.Wardboy = *room.Wardboy

	c.JSON(200, response)
}

func GetRooms(c *gin.Context) {
	var rooms []models.Room

	result := db.Preload("Wardboy").Find(&rooms)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, rooms)
}

func GetRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var room models.Room

	result := db.Preload("Wardboy").First(&room, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, room)
}

func UpdateRoom(c *gin.Context) {
	var data roomRequest

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": room_val.DecryptErrors(err)})
		return
	}

	var room models.Room

	result := db.First(&room, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}

	// Check if wardboy exists
	var wardboy models.Wardboy

	result = db.First(&wardboy, data.WardboyID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Wardboy not found"})
		return
	}

	room.ID = uint(id)
	room.Type = data.Type
	room.Number = data.Number
	room.WardboyID = data.WardboyID
	room.Wardboy = &wardboy

	result = db.Save(&room)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response roomResponse
	response.ID = room.ID
	response.Type = room.Type
	response.Number = room.Number
	response.WardboyID = room.WardboyID
	response.Wardboy = *room.Wardboy

	c.JSON(200, response)
}

func DeleteRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var room models.Room

	result := db.First(&room, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}

	result = db.Delete(&room, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Room deleted successfully"})
}
