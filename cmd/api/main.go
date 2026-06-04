package main

import (
	"net/http"

	"dkrpg2-backend/internal/repository" // <-- เพิ่มเข้ามา
	"dkrpg2-backend/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateCharacterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	ClassKey string `json:"class_key" binding:"required"`
	Stats    struct {
		HP  int `json:"hp" binding:"required"`
		MP  int `json:"mp" binding:"required"`
		STR int `json:"str" binding:"required"`
		AGI int `json:"agi" binding:"required"`
		INT int `json:"int" binding:"required"`
		Def int `json:"def" binding:"required"`
	} `json:"stats" binding:"required"`
	CombatStats struct {
		HpRegen     int `json:"hp_regen"`
		MpRegen     int `json:"mp_regen"`
		Acc         int `json:"acc"`
		Block       int `json:"block"`
		DmgRed      int `json:"dmg_red"`
		CritRate    int `json:"crit_rate"`
		CritDmg     int `json:"crit_dmg"`
		Dodge       int `json:"dodge"`
		IgnoreBlock int `json:"ignore_block"`
	} `json:"combat_stats" binding:"required"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ครบถ้วน หรือรูปแบบอีเมลผิด"})
			return
		}

		// 🔒 เข้ารหัส Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการเข้ารหัสข้อมูล"})
			return
		}

		// ✅ ใช้ค่าจาก Frontend โดยตรง (รวม Combat Stats)
		newPlayer := services.CreateNewCharacterFromClient(
			req.Username,
			req.ClassKey,
			req.Stats.HP,
			req.Stats.MP,
			req.Stats.STR,
			req.Stats.AGI,
			req.Stats.INT,
			req.Stats.Def,
			req.CombatStats.HpRegen,
			req.CombatStats.MpRegen,
			req.CombatStats.Acc,
			req.CombatStats.Block,
			req.CombatStats.DmgRed,
			req.CombatStats.CritRate,
			req.CombatStats.CritDmg,
			req.CombatStats.Dodge,
			req.CombatStats.IgnoreBlock,
		)
		newPlayer.Email = req.Email
		newPlayer.PasswordHash = string(hashedPassword)

		if err := repository.CreatePlayer(&newPlayer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ชื่อตัวละครหรืออีเมลนี้มีคนใช้แล้ว"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "สร้างตัวละครสำเร็จ", "data": newPlayer})
	})
	
	r.POST("/api/character/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณากรอกข้อมูลให้ครบถ้วน"})
			return
		}

		// ค้นหาด้วย Email หรือ Username ก็ได้
		player, err := repository.GetPlayerByIdentifier(req.Identifier)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ใช้งานนี้ในระบบ"})
			return
		}

		// 🔓 ตรวจสอบรหัสผ่านว่าตรงกับ Hash ในระบบหรือไม่
		err = bcrypt.CompareHashAndPassword([]byte(player.PasswordHash), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านไม่ถูกต้อง!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "เข้าสู่ระบบสำเร็จ", "data": player})
	})

	r.Run(":8080")
}