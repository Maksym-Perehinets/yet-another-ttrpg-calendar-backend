package database

import (
	"fmt"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/interfaces"
	"github.com/Maksym-Perehinets/yet-another-ttrpg-calendar-backend/calendar/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

// Service represents a service that interacts with a database.

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("CALENDAR_DB_DATABASE")
	password   = os.Getenv("CALENDAR_DB_PASSWORD")
	username   = os.Getenv("CALENDAR_DB_USERNAME")
	port       = os.Getenv("CALENDAR_DB_PORT")
	host       = os.Getenv("CALENDAR_DB_HOST")
	schema     = os.Getenv("CALENDAR_DB_SCHEMA")
	dbInstance *service
)

func New() interfaces.Service {
	if dbInstance != nil {
		return dbInstance
	}

	// Construct the connection string for GORM
	log.Printf("Connecting to database: %s", database)
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		host, username, password, database, port, schema)

	// Initialize GORM
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	// Auto-migrate the Users model
	if err := db.AutoMigrate(&models.Locations{}, &models.TTRPGSystems{}, models.Users{}, &models.Games{}); err != nil {
		log.Fatalf("Failed to auto-migrate Users model: %v", err)
	}

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	stats := make(map[string]string)

	// Attempt to ping the database
	sqlDB, err := s.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("failed to retrieve sql.DB instance: %v", err)
		log.Printf("Health check failed: %v", err)
		return stats
	}

	err = sqlDB.Ping()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("Health check failed: %v", err)
		return stats
	}

	// Database is up
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to retrieve sql.DB instance: %w", err)
	}

	log.Printf("Disconnected from database: %s", database)
	return sqlDB.Close()
}

// DB returns the GORM DB instance.
func (s *service) DB() *gorm.DB {
	log.Printf("Returning GORM DB instance for application use")
	return s.db
}

// ElementExists checks if an element exists in the database with same fields.
//func (s *service) ElementExists(fields map[string]interface{}, model interface{}) bool {
//	var count int64
//
//	rs := s.db.Model(model).Where(fields).Count(&count)
//	if rs.Error != nil {
//		log.Printf("Error while checking if element exists: %v", rs.Error)
//		return false
//	}
//
//	return count == 0
//}
