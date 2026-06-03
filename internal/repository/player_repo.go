package repository

import (
	"context"
	"encoding/json"

	"dkrpg2-backend/internal/models"
)

// ฟังก์ชันบันทึกตัวละครลงฐานข้อมูล
func CreatePlayer(p *models.Player) error {
	query := `
		INSERT INTO players (
			username, class_key, level, exp, max_exp, gold, stat_points,
			max_slots, max_weight, base_stats, secondary_stats,
			equipment, inventory, skills, loadout, active_buffs
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		) RETURNING id; -- ขอ ID ที่ฐานข้อมูลสร้างขึ้นกลับมาด้วย
	`

	// แปลง Struct ซับซ้อนให้กลายเป็น JSON ก่อนเซฟ
	baseJSON, _ := json.Marshal(p.BaseStats)
	secJSON, _ := json.Marshal(p.SecStats)
	equipJSON, _ := json.Marshal(p.Equipment)
	invJSON, _ := json.Marshal(p.Inventory)
	skillsJSON, _ := json.Marshal(p.Skills)
	loadoutJSON, _ := json.Marshal(p.Loadout)
	buffsJSON, _ := json.Marshal(p.Buffs)

	// ยิงคำสั่ง INSERT ลงฐานข้อมูลและนำ ID (UUID) ที่ได้มาใส่ในตัวแปร p.ID
	err := DB.QueryRow(context.Background(), query,
		p.Username, p.ClassKey, p.Level, p.Exp, p.MaxExp, p.Gold, p.StatPoints,
		p.MaxSlots, p.MaxWeight, baseJSON, secJSON,
		equipJSON, invJSON, skillsJSON, loadoutJSON, buffsJSON,
	).Scan(&p.ID)

	return err
}