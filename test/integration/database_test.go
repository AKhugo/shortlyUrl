package integration

import (
	"os"
	"testing"

	"shortlyurl/internal/config"
	"shortlyurl/internal/database"

	"github.com/joho/godotenv"
)

func TestDatabaseConnection(t *testing.T) {
	// Skip if not in integration test environment
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TESTS=true to run.")
	}

	// Load test environment
	err := godotenv.Load("../fixtures/.env.test")
	if err != nil {
		// Fallback to regular .env for local testing
		err = godotenv.Load("../../.env")
		if err != nil {
			t.Fatal("Error loading .env file:", err)
		}
	}

	// Load config
	config.LoadConfig()

	// Test database connection
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Test ping
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	t.Log("Database connection test passed!")
}

func TestDatabaseConnectionWithInvalidCredentials(t *testing.T) {
	// Skip if not in integration test environment
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test. Set INTEGRATION_TESTS=true to run.")
	}

	// Set invalid config
	originalConfig := config.DbConfig
	config.DbConfig = config.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "invalid_user",
		Password: "invalid_password",
		DBName:   "invalid_db",
		SSLMode:  "disable",
	}

	// Test should fail
	_, err := database.Connect()
	if err == nil {
		t.Fatal("Expected connection to fail with invalid credentials")
	}

	// Restore original config
	config.DbConfig = originalConfig

	t.Log("Invalid credentials test passed!")
}
