package unit

import (
	"os"
	"testing"

	"shortlyurl/internal/config"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected config.DatabaseConfig
	}{
		{
			name: "should load all environment variables correctly",
			envVars: map[string]string{
				"DB_HOST":     "testhost",
				"DB_PORT":     "5433",
				"DB_USER":     "testuser",
				"DB_PASSWORD": "testpass",
				"DB_NAME":     "testdb",
				"DB_SSLMODE":  "require",
			},
			expected: config.DatabaseConfig{
				Host:     "testhost",
				Port:     "5433",
				User:     "testuser",
				Password: "testpass",
				DBName:   "testdb",
				SSLMode:  "require",
			},
		},
		{
			name:    "should handle empty environment variables",
			envVars: map[string]string{},
			expected: config.DatabaseConfig{
				Host:     "",
				Port:     "",
				User:     "",
				Password: "",
				DBName:   "",
				SSLMode:  "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Backup original env vars
			originalEnvs := make(map[string]string)
			envKeys := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE"}
			for _, key := range envKeys {
				originalEnvs[key] = os.Getenv(key)
			}

			// Clean env vars
			for _, key := range envKeys {
				os.Unsetenv(key)
			}

			// Set test env vars
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Test LoadConfig
			config.LoadConfig()

			// Verify results
			if config.DbConfig.Host != tt.expected.Host {
				t.Errorf("Expected Host %s, got %s", tt.expected.Host, config.DbConfig.Host)
			}
			if config.DbConfig.Port != tt.expected.Port {
				t.Errorf("Expected Port %s, got %s", tt.expected.Port, config.DbConfig.Port)
			}
			if config.DbConfig.User != tt.expected.User {
				t.Errorf("Expected User %s, got %s", tt.expected.User, config.DbConfig.User)
			}
			if config.DbConfig.Password != tt.expected.Password {
				t.Errorf("Expected Password %s, got %s", tt.expected.Password, config.DbConfig.Password)
			}
			if config.DbConfig.DBName != tt.expected.DBName {
				t.Errorf("Expected DBName %s, got %s", tt.expected.DBName, config.DbConfig.DBName)
			}
			if config.DbConfig.SSLMode != tt.expected.SSLMode {
				t.Errorf("Expected SSLMode %s, got %s", tt.expected.SSLMode, config.DbConfig.SSLMode)
			}

			// Restore original env vars
			for key, value := range originalEnvs {
				if value != "" {
					os.Setenv(key, value)
				} else {
					os.Unsetenv(key)
				}
			}
		})
	}
}
