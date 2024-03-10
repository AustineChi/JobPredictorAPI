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
		log.Println("Error in Migration", err)
	}
	//verifies if a connection to the database is still alive, establishing a connection if necessary.
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("error in getting sql")
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Println("Error in connection", err)
		return nil, err
	}
	return db, nil
}
