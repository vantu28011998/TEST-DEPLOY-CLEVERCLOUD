package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	dsn := "host=ec2-35-168-145-180.compute-1.amazonaws.com user=petvagkxcfzouq password=3ffd359bafd92627b144bc789de31cee749d7ce118be6b6fb49826f99df9fb03 dbname=d9f6b3th4lqqop port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
