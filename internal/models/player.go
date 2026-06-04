package models

import "time"

// ตัวละครผู้เล่น
type Player struct {
	ID           string            `json:"id"`
	Email        string            `json:"email"`
	Username     string            `json:"username"`
	PasswordHash string            `json:"-"`
	ClassKey     string            `json:"class_key"`
	Level        int               `json:"level"`
	Exp          int               `json:"exp"`
	MaxExp       int               `json:"max_exp"`
	Gold         int               `json:"gold"`
	StatPoints   int               `json:"stat_points"`

	// ระบบความจุ
	MaxSlots  int `json:"max_slots"`
	MaxWeight int `json:"max_weight"`

	// สเตตัสและข้อมูลสวมใส่
	BaseStats BaseStats         `json:"base_stats"`
	SecStats  SecStats          `json:"secondary_stats"`
	Equipment map[string]string `json:"equipment"`
	Inventory []ItemInstance `json:"inventory"`

	// ระบบสกิล
	Skills  map[string]int `json:"skills"`
	Loadout []string       `json:"loadout"`

	Buffs []Buff `json:"active_buffs"`
}

type BaseStats struct {
	MaxHP int `json:"max_hp"`
	MaxMP int `json:"max_mp"`
	HP    int `json:"hp"`
	MP    int `json:"mp"`
	STR   int `json:"str"`
	AGI   int `json:"agi"`
	INT   int `json:"int"`
}

type SecStats struct {
	Atk         int `json:"atk"`
	Matk        int `json:"matk"`
	Def         int `json:"def"`
	Acc         int `json:"acc"`
	Eva         int `json:"eva"`
	Block       int `json:"block"`
	IgnoreBlock int `json:"ignore_block"`
	CritRate    int `json:"crit_rate"`
	CritDmg     int `json:"crit_dmg"`
	DmgRed      int `json:"dmg_red"`
	HpRegen     int `json:"hp_regen"`
	MpRegen     int `json:"mp_regen"`
}

type Buff struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	EffectType  string    `json:"effect_type"`
	Value       int       `json:"value"`
	IsTurnBased bool      `json:"is_turn_based"`
	Duration    int       `json:"duration"`
	AppliedAt   time.Time `json:"applied_at"`
}

type ItemInstance struct {
	UID           string `json:"uid"`
	ItemRefID     string `json:"item_ref_id"`
	Quantity      int    `json:"quantity"`
	UpgradeLevel  int    `json:"upgrade_level"`
	Durability    int    `json:"durability"`
	MaxDurability int    `json:"max_durability"`
	SlotIndex     int    `json:"slot_index"`
}