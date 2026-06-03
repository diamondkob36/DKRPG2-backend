package repository

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// ตัวแปรส่วนกลางสำหรับเรียกใช้ฐานข้อมูล
var DB *pgxpool.Pool

func InitDB() {
	// โหลดไฟล์ .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ ไม่พบไฟล์ .env (จะใช้ Environment Variable ของระบบแทน)")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ กรุณากำหนด DATABASE_URL ในไฟล์ .env")
	}

	// สร้าง Connection Pool ไปที่ Supabase
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("❌ เชื่อมต่อฐานข้อมูลไม่สำเร็จ: ", err)
	}

	DB = pool
	log.Println("✅ เชื่อมต่อฐานข้อมูล Supabase PostgreSQL สำเร็จ!")
}