// internal/services/stats_service.go
package services

import "dkrpg2-backend/internal/models"

// คำนวณ SecStats จากค่าที่ Frontend ส่งมา (ใช้สำหรับสร้างตัวละครใหม่)
func CalculateSecondaryStatsFromClient(base models.BaseStats, def, hpRegen, mpRegen, acc, block, dmgRed, critRate, critDmg, dodge, ignoreBlock int) models.SecStats {
	return models.SecStats{
		Atk:         base.STR * 2,
		Matk:        base.INT * 2,
		Def:         def,
		Acc:         acc,
		Eva:         dodge,
		Block:       block,
		IgnoreBlock: ignoreBlock,
		CritRate:    critRate,
		CritDmg:     critDmg,
		DmgRed:      dmgRed,
		HpRegen:     hpRegen,
		MpRegen:     mpRegen,
	}
}

// คำนวณ SecStats แบบพื้นฐาน (ใช้สำหรับคำนวณใหม่ระหว่างเล่นเกม เช่น เมื่อเพิ่มสเตตัส)
func CalculateSecondaryStats(base models.BaseStats, oldSecStats models.SecStats, activeBuffs []models.Buff) models.SecStats {
	
	secStats := oldSecStats
	secStats.Atk = base.STR * 2
	secStats.Matk = base.INT * 2
	secStats.Eva = oldSecStats.Eva + (base.AGI / 4)

	for _, buff := range activeBuffs {
		switch buff.EffectType {
		case "increase_atk":
			secStats.Atk += buff.Value
		case "increase_eva":
			secStats.Eva += buff.Value
		case "increase_acc":
			secStats.Acc += buff.Value
		case "increase_crit":
			secStats.CritRate += buff.Value
		}
	}

	if secStats.HpRegen < 1 {
		secStats.HpRegen = 1
	}
	if secStats.MpRegen < 1 {
		secStats.MpRegen = 1
	}

	return secStats
}