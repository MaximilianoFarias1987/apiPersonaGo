package commons

import (
	"crud-api-rest-go/models"

	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	// server := "localhost"
	// user := "sa"
	// password := "1234"
	// database := "dbpruebago"

	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;", server, user, password, database)
	dsn := "sqlserver://sa:1234@LAPTOP-9ALNHCMO:49213?database=dbpruebago"
	db, error := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}
	// defer db.Close()

	return db

}

//Funcion que hace la migracion a la bd

func Migrate() {
	db := GetConnection()

	log.Println("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
