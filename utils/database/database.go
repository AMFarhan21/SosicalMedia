package database

import (
	"database/sql"
	"fmt"
	"log"
	"socialmedia/utils/config"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	Gorm  *gorm.DB
	Mongo *mongo.Client
	Raw   *sql.DB
}

func GetDatabaseConnection(cfg *config.Config, db string) (*DBConn, error) {
	switch db {
	case "GORM":
		return GetGormConnection(cfg)
	case "MONGO":
		return GetMongoConnection(cfg)
	case "RAW":
		return GetRawSQLConnection(cfg)
	}
	return nil, fmt.Errorf("unknown database type: %v", db)
}

func GetGormConnection(cfg *config.Config) (*DBConn, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgreDatabase.Host, cfg.PostgreDatabase.DBPort, cfg.PostgreDatabase.User, cfg.PostgreDatabase.Password, cfg.PostgreDatabase.Name, cfg.PostgreDatabase.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)

	log.Print("Successfully connected to the database")
	return &DBConn{
		Gorm: db,
	}, nil
}

func GetMongoConnection(cfg *config.Config) (*DBConn, error) {
	dsn := cfg.MongoDatabase

	client, err := mongo.Connect(options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	log.Print("Successfully connected to the mongoDB")
	return &DBConn{
		Mongo: client,
	}, nil
}

func GetRawSQLConnection(cfg *config.Config) (*DBConn, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgreDatabase.Host, cfg.PostgreDatabase.DBPort, cfg.PostgreDatabase.User, cfg.PostgreDatabase.Password, cfg.PostgreDatabase.Name, cfg.PostgreDatabase.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Print("Successfully connected to the database")
	return &DBConn{
		Raw: db,
	}, nil
}
