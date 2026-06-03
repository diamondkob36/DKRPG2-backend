package services

import (
	"dkrpg2-backend/internal/models"
)

// ฟังก์ชันจำลองฐานข้อมูลอาชีพเริ่มต้น (เทียบเท่า classData.js)
func getBaseClassStats(classKey string) models.BaseStats {
	switch classKey {
	case "knight":
		return models.BaseStats{MaxHP: 150, MaxMP: 30, STR: 15, AGI: 5, INT: 5}
	case "mage":
		return models.BaseStats{MaxHP: 80, MaxMP: 120, STR: 3, AGI: 7, INT: 20}
	case "rogue":
		return models.BaseStats{MaxHP: 100, MaxMP: 50, STR: 10, AGI: 18, INT: 5}
	default: // novice
		return models.BaseStats{MaxHP: 100, MaxMP: 50, STR: 10, AGI: 10, INT: 10}
	}
}

func CreateNewCharacter(username string, classKey string) models.Player {
	baseStats := getBaseClassStats(classKey)
	
	// เซ็ต HP/MP ให้เต็ม
	baseStats.HP = baseStats.MaxHP
	baseStats.MP = baseStats.MaxMP + (baseStats.INT * 10) // สูตรคำนวณ MaxMP ตาม game-logic.js
	baseStats.MaxMP = baseStats.MP

	// แจกอาวุธเริ่มต้นตามอาชีพ
	startWeapon := "wooden_sword"
	if classKey == "mage" {
		startWeapon = "novice_staff"
	} else if classKey == "rogue" {
		startWeapon = "novice_dagger"
	}

	player := models.Player{
		ID:         "mock-uuid-1234", // (เดี๋ยวเราจะใช้ UUID จริงจาก Supabase ภายหลัง)
		Username:   username,
		ClassKey:   classKey,
		Level:      1,
		Exp:        0,
		MaxExp:     100,
		Gold:       0,
		StatPoints: 5,
		MaxSlots:   32,
		MaxWeight:  60,
		BaseStats:  baseStats,
		Equipment:  make(map[string]string),
		Inventory: map[string]int{
			"potion_s":  3,
			startWeapon: 1,
		},
		Skills:  make(map[string]int),
		Loadout: make([]string, 6), // ช่องสกิล 6 ช่อง
		Buffs:   []models.Buff{},
	}

	// คำนวณสเตตัสรอง (เรียกใช้ฟังก์ชันที่คุณแก้ไว้แล้ว)
	player.SecStats = CalculateSecondaryStats(baseStats, player.Buffs)

	return player
}