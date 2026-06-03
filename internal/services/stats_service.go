package services

import "dkrpg2-backend/internal/models"

// ฟังก์ชันคำนวณสเตตัสรองจากสเตตัสหลัก
func CalculateSecondaryStats(base models.BaseStats, activeBuffs []models.Buff) models.SecStats {
	secStats := models.SecStats{
		AttackPower: base.STR * 2,
		MagicPower:  base.INT * 2,
		Defense:     base.STR / 2,
		Accuracy:    90 + (base.AGI / 2), // พื้นฐานแม่นยำ 90% + โบนัสจาก AGI
		Evasion:     5 + (base.AGI / 3),  // พื้นฐานหลบหลีก 5% + โบนัสจาก AGI
	}

	// ลอจิกคำนวณบัฟที่ส่งผลต่อสเตตัสรอง
	for _, buff := range activeBuffs {
		if buff.EffectType == "increase_evasion" {
			secStats.Evasion += buff.Value
		}
		// เพิ่มเงื่อนไขบัฟอื่นๆ ตามลอจิกเวอร์ชันแรกได้เลย
	}

	return secStats
}