package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to db:%v", err)
		return nil, err
	}
	err = db.AutoMigrate(&User{}, &SavedJob{}, &JobRecommendation{}, &Notification{}, &Job{}, &Interaction{})
	if err != nil {
		log.Fatal("Error in Migration", err)
	}
	log.Println("Migrations Successful")
	//verifies if a connection to the database is still alive, establishing a connection if necessary.
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("error in getting sql connection")
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error in connection", err)
		return nil, err
	}
	log.Println("connected to DB")
	return db, nil
}
