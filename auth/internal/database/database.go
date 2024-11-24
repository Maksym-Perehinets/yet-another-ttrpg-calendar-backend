package database

import (
	"auth/interfaces"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"

	"auth/internal/models"
)

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("AUTH_DB_DATABASE")
	password   = os.Getenv("AUTH_DB_PASSWORD")
	username   = os.Getenv("AUTH_DB_USERNAME")
	port       = os.Getenv("AUTH_DB_PORT")
	host       = os.Getenv("AUTH_DB_HOST")
	schema     = os.Getenv("AUTH_DB_SCHEMA")
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

	// Auto-migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate User model: %v", err)
	}

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

// Health checks the health of the database connection by pinging it.
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

// AlreadyExists checks if a record with the given column and value already exists in the database.
// If the record exists, it returns error.
func (s *service) AlreadyExists(c string, v string, m any) error {
	log.Printf("Checking if record in colum %s with %s already exists", c, v)
	var count int64

	if err := s.db.Model(&m).Where(fmt.Sprintf("%s = ?", c), v).Count(&count).Error; err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}

	if count != 0 {
		return fmt.Errorf("record with %s %s already exist", c, v)
	}

	return nil
}

// GetUser retrieves a user from the database based on the column and value provided.
func (s *service) GetUser(c string, v string) (*models.User, error) {
	var user models.User
	result := s.db.Model(&models.User{}).Where(fmt.Sprintf("%s = ?", c), v).First(&user)
	if result.Error != nil {
		log.Printf("Error querying database: %v", result.Error)
		return nil, fmt.Errorf("internal server error")
	}
	return &user, nil
}

// Close closes the database connection.
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
