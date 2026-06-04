package main

import (
	"net/http"

	"dkrpg2-backend/internal/repository"
	"dkrpg2-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type CreateCharacterRequest struct {
	Username string `json:"username" binding:"required"`
	ClassKey string `json:"class_key" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
}

func main() {
	// 1. เรียกใช้งานการเชื่อมต่อฐานข้อมูล
	repository.InitDB()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/api/character/create", func(c *gin.Context) {
		var req CreateCharacterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ครบถ้วน หรือรูปแบบผิด"})
			return
		}

		newPlayer := services.CreateNewCharacter(req.Username, req.ClassKey)

		if err := repository.CreatePlayer(&newPlayer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถบันทึกข้อมูลตัวละครได้: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "สร้างตัวละครสำเร็จ",
			"data":    newPlayer,
		})
	})
	
	r.POST("/api/character/login", func(c *gin.Context) {
		var req LoginRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาส่งชื่อตัวละคร"})
			return
		}

		player, err := repository.GetPlayerByUsername(req.Username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบตัวละครชื่อนี้ในระบบ!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "เข้าสู่ระบบสำเร็จ",
			"data":    player,
		})
	})

	r.Run(":8080")
}