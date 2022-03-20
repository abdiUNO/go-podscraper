package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"path"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func loadEnv() error {
	goPath := os.Getenv("GOPATH")
	log.Println(goPath)
	filePath := path.Join(goPath, "src/github.com/abdiUNO/go-podscraper/.env")
	log.Println(filePath)

	e := godotenv.Load(filePath)

	if e != nil {
		fmt.Println(e)
		return e
	}
	return nil
}

func init() {
	if e := loadEnv(); e != nil {
		log.Fatal(e)
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbType := os.Getenv("DB_TYPE")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	fmt.Println(dbUri)

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")

	db.Debug().AutoMigrate(&Podcast{}, &Rank{}, &Genre{})
}

func GetDB() *gorm.DB {
	return db
}

type GormModel struct {
	ID        string     `gorm:"primary_key;type:varchar(255);"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-";sql:"index"`
}

func (model *GormModel) BeforeCreate(scope *gorm.Scope) error {
	u1 := uuid.Must(uuid.NewV4())
	err := scope.SetColumn("ID", u1.String())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
