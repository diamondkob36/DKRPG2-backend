package services

import (
	"dkrpg2-backend/internal/models"
)

// ฟังก์ชันสร้างตัวละครจากข้อมูลที่ Frontend ส่งมา (รวม Combat Stats)
func CreateNewCharacterFromClient(username, classKey string, hp, mp, str, agi, intStat, def, hpRegen, mpRegen, acc, block, dmgRed, critRate, critDmg, dodge, ignoreBlock int) models.Player {
	baseStats := models.BaseStats{
		MaxHP: hp,
		HP:    hp,
		MaxMP: mp + (intStat * 10),
		MP:    mp + (intStat * 10),
		STR:   str,
		AGI:   agi,
		INT:   intStat,
	}

	// แจกอาวุธเริ่มต้นตามอาชีพ
	startWeapon := "wooden_sword"
	if classKey == "mage" {
		startWeapon = "novice_staff"
	} else if classKey == "rogue" {
		startWeapon = "novice_dagger"
	}

	initialItem := models.ItemInstance{
		UID:           "item-" + username + "-1", // สร้าง UID แบบง่ายๆ ไปก่อน (เช่น item-diamond-1)
		ItemRefID:     startWeapon,               // รหัสไอเทม
		Quantity:      1,                         // ได้ 1 ชิ้น
		UpgradeLevel:  0,                         // ตีบวก +0
		Durability:    50,                        // ความทนทานเริ่มต้น
		MaxDurability: 50,                        // ความทนทานสูงสุด
		SlotIndex:     0,                         // วางไว้ช่องแรกสุดของกระเป๋า
	}

	player := models.Player{
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
		Inventory:  []models.ItemInstance{initialItem},
		
		Skills:  make(map[string]int),
		Loadout: make([]string, 6),
		Buffs:   []models.Buff{},
	}

	// ใช้ค่าที่ส่งมาจาก Frontend โดยตรง
	player.SecStats = CalculateSecondaryStatsFromClient(baseStats, def, hpRegen, mpRegen, acc, block, dmgRed, critRate, critDmg, dodge, ignoreBlock)

	return player
}