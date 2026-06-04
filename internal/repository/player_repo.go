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
			email, password_hash, username, class_key, level, exp, max_exp, gold, stat_points,
			max_slots, max_weight, base_stats, secondary_stats,
			equipment, inventory, skills, loadout, active_buffs
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
		) RETURNING id;
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
		p.Email, p.PasswordHash, p.Username, p.ClassKey, p.Level, p.Exp, p.MaxExp, p.Gold, p.StatPoints,
		p.MaxSlots, p.MaxWeight, baseJSON, secJSON,
		equipJSON, invJSON, skillsJSON, loadoutJSON, buffsJSON,
	).Scan(&p.ID)

	return err
}

func GetPlayerByIdentifier(identifier string) (*models.Player, error) {
	query := `
		SELECT id, email, password_hash, username, class_key, level, exp, max_exp, gold, stat_points,
		       max_slots, max_weight, base_stats, secondary_stats,
		       equipment, inventory, skills, loadout, active_buffs
		FROM players
		WHERE username = $1 OR email = $1
	`
	
	var p models.Player
	var baseJSON, secJSON, equipJSON, invJSON, skillsJSON, loadoutJSON, buffsJSON []byte

	err := DB.QueryRow(context.Background(), query, identifier).Scan(
		&p.ID, &p.Email, &p.PasswordHash, &p.Username, &p.ClassKey, &p.Level, &p.Exp, &p.MaxExp, &p.Gold, &p.StatPoints,
		&p.MaxSlots, &p.MaxWeight, &baseJSON, &secJSON,
		&equipJSON, &invJSON, &skillsJSON, &loadoutJSON, &buffsJSON,
	)

	if err != nil {
		return nil, err
	}

	// แปลง JSONB กลับมาเป็น Struct ของ Go
	json.Unmarshal(baseJSON, &p.BaseStats)
	json.Unmarshal(secJSON, &p.SecStats)
	json.Unmarshal(equipJSON, &p.Equipment)
	json.Unmarshal(invJSON, &p.Inventory)
	json.Unmarshal(skillsJSON, &p.Skills)
	json.Unmarshal(loadoutJSON, &p.Loadout)
	json.Unmarshal(buffsJSON, &p.Buffs)

	return &p, nil
}