package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insert(b *testing.B, db *gorm.DB) {
	user := User{Name: "Jtip"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

// go test -bench=, -benchmem
func BenchmarkMaxOpenCon1(b *testing.B) {
	dsn := "root:root@tcp(127.0.0.1:3306)/shopdev?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if exists
		if err := db.Migrator().DropTable(&User{}); err != nil {

		}

	}
	// create table if dont have
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB form gorm.DB: %v", err)

	}
	// set up parameter connection
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			insert(b, db)
		}
	})
}

// 462	   2453705 ns/op	    5942 B/op	      75 allocs/op
//2912	    382302 ns/op	    5749 B/op	      73 allocs/op