package models

import "time"

// ตัวละครผู้เล่น
type Player struct {
	ID        string      `json:"id"`
	Username  string      `json:"username"`
	ClassKey  string      `json:"class_key"`
	Level     int         `json:"level"`
	BaseStats BaseStats   `json:"base_stats"`
	SecStats  SecStats    `json:"secondary_stats"`
	Buffs     []Buff      `json:"active_buffs"` // เก็บสถานะบัฟที่กำลังทำงาน
}

// สเตตัสหลัก (Primary Stats)
type BaseStats struct {
	MaxHP int `json:"max_hp"`
	MaxMP int `json:"max_mp"`
	STR   int `json:"str"`
	AGI   int `json:"agi"`
	INT   int `json:"int"`
}

// สเตตัสรอง (Secondary Stats) ครบถ้วนตาม game-logic.js
type SecStats struct {
	Atk         int `json:"atk"`
	Matk        int `json:"matk"`
	Def         int `json:"def"`
	Acc         int `json:"acc"`          // แม่นยำ (หักลบหลบหลีก)
	Eva         int `json:"eva"`          // หลบหลีก (ในโค้ดเก่าใช้ชื่อ dodge)
	Block       int `json:"block"`        // โอกาสบล็อก (%)
	IgnoreBlock int `json:"ignore_block"` // เจาะเกราะ/ลดบล็อก (%)
	CritRate    int `json:"crit_rate"`    // โอกาสคริติคอล (%)
	CritDmg     int `json:"crit_dmg"`     // ความแรงคริติคอล (เริ่มต้น 150)
	DmgRed      int `json:"dmg_red"`      // ลดความเสียหาย (สูงสุด 40%)
	HpRegen     int `json:"hp_regen"`
	MpRegen     int `json:"mp_regen"`
}	

// โครงสร้างของบัฟและดีบัฟ
type Buff struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	EffectType  string    `json:"effect_type"` // เช่น "increase_evasion", "poison"
	Value       int       `json:"value"`
	IsTurnBased bool      `json:"is_turn_based"` // แยกบัฟแบบนับเทิร์นกับเรียลไทม์
	Duration    int       `json:"duration"`      // จำนวนเทิร์น หรือ วินาทีที่เหลือ
	AppliedAt   time.Time `json:"applied_at"`
}