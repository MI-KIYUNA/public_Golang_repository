package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// データベース接続初期化（LaravelのDB::connectionみたいな）
func InitDB() {
	// .env読み込み
	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルが見つかりません")
	}

	// DSN構築（LaravelのDB_HOSTとかと同じ）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// DB接続
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // SQLログ表示
	})

	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}

	log.Println("✅ データベース接続成功")
}