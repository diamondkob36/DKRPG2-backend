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

// สเตตัสรอง (Secondary Stats)
type SecStats struct {
	Atk      int `json:"atk"`       // พลังโจมตี (Attack)
	Matk     int `json:"matk"`      // พลังเวท (Magic Attack)
	Def      int `json:"def"`       // พลังป้องกัน (Defense)
	Acc      int `json:"acc"`       // ความแม่นยำ (Accuracy) - เริ่มต้น 5%
	Eva      int `json:"eva"`       // อัตราหลบหลีก (Evasion) 
	CritRate int `json:"crit_rate"` // โอกาสคริติคอล (Critical Rate) - เริ่มต้น 5% อิงจากไอเทมเท่านั้น
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