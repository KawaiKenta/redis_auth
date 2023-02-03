package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"kk-rschian.com/redis_auth/config"
	"kk-rschian.com/redis_auth/const/models"
	"kk-rschian.com/redis_auth/service/database"
)

func main() {
	// 扱い方の表示
	if len(os.Args) != 2 {
		fmt.Println("database migration tool for redis_auth")
		fmt.Println("usage:")
		fmt.Println("\ttable:up\tcreate table if not exist")
		fmt.Println("\ttable:down\tdrop table if it exist")
		fmt.Println("\tseed:create\tcreate sample data for each tables (tables is need before execute)")
		os.Exit(0)
	}

	config.Setup()
	database.Setup()
	defer database.Close()

	flag := os.Args[1]
	switch flag {
	case "table:up":
		if err := database.DB.AutoMigrate(
			&models.User{},
		); err != nil {
			panic(err)
		}
		fmt.Printf("Tables are created in %s\n", config.Database.Name)
	case "table:down":
		if err := database.DB.Migrator().DropTable(
			&models.User{},
		); err != nil {
			panic(err)
		}
		fmt.Printf("All tables in %s is droped\n", config.Database.Name)
	case "seed:create":
		if err := userSeeds(); err != nil {
			panic(err)
		}
		println("users are created")
	default:
		fmt.Println("wrong argument is passed")
		fmt.Println("usage:")
		fmt.Println("\ttable:up\tcreate table if not exist")
		fmt.Println("\ttable:down\tdrop table if it exist")
		fmt.Println("\tseed:create\tcreate sample data for each tables (tables is need before execute)")
	}

}

func userSeeds() error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := models.User{
			Name:     "user" + strconv.Itoa(i+1),
			Email:    "sample" + strconv.Itoa(i+1) + "@gmail.com",
			Password: string(hash),
		}

		if err := database.DB.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
