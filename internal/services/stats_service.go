// internal/services/stats_service.go
package services

import "dkrpg2-backend/internal/models"

// ฟังก์ชันคำนวณสเตตัสรองจากสเตตัสหลักและบัฟ
func CalculateSecondaryStats(base models.BaseStats, activeBuffs []models.Buff) models.SecStats {
	// คำนวณค่าเริ่มต้นตามสูตร game-logic.js ฉบับสมบูรณ์
	secStats := models.SecStats{
		Atk:         base.STR * 2, // พลังโจมตีพื้นฐาน
		Matk:        base.INT * 2, // พลังเวทพื้นฐาน
		Def:         0,            // พลังป้องกันพื้นฐาน (ส่วนใหญ่จะได้จากไอเทมสวมใส่)
		Acc:         5,            // แม่นยำเริ่มต้น 5% (ไว้หักลบหลบหลีกเป้าหมาย)
		Eva:         base.AGI / 4, // หลบหลีกเริ่มต้น (ได้โบนัส 1% ทุกๆ 4 AGI)
		Block:       0,            // โอกาสบล็อกเริ่มต้น
		IgnoreBlock: 0,            // เจาะเกราะเริ่มต้น
		CritRate:    5,            // โอกาสคริติคอลเริ่มต้น 5% (ไม่อิง AGI)
		CritDmg:     150,          // ความแรงคริติคอลเริ่มต้น 150%
		DmgRed:      0,            // ลดความเสียหายเริ่มต้น
		HpRegen:     int(float64(base.MaxHP) * 0.05), // ฟื้นฟู HP 5% จาก MaxHP
		MpRegen:     int(float64(base.MaxMP) * 0.05), // ฟื้นฟู MP 5% จาก MaxMP
	}

	// ลอจิกคำนวณบัฟที่ส่งผลต่อสเตตัสรอง
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
		// สามารถเพิ่ม case สำหรับบัฟอื่นๆ ได้ในอนาคต
		}
	}

	// กฎเหล็ก: การฟื้นฟูเลือดและมานาขั้นต่ำต้องเป็น 1 เสมอ
	if secStats.HpRegen < 1 {
		secStats.HpRegen = 1
	}
	if secStats.MpRegen < 1 {
		secStats.MpRegen = 1
	}

	return secStats
}