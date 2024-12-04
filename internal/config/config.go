package config

import (
	"github.com/spf13/viper"
)

// AppConfig defines application config
type AppConfig struct {
	Secret     string
	Brokers    string
	EventTopic string
	DBConnStr  string
}

// InitConfig initialises configuration from file and envs
func InitConfig() (AppConfig, error) {
	// Load config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}

	// Override with environment variables (if present)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.BindEnv("jwt.secret", "APP_JWT_SECRET")
	viper.BindEnv("kafka.brokers", "APP_KAFKA_BROKERS")
	viper.BindEnv("api.event_topic", "APP_API_EVENT_TOPIC")
	viper.BindEnv("database.connection", "APP_DB_CONN")

	jwtSecret := viper.GetString("jwt.secret")
	kafkaBrokers := viper.GetString("kafka.brokers")
	eventTopic := viper.GetString("api.event_topic")
	dbConnStr := viper.GetString("database.connection")

	return AppConfig{
		Secret:     jwtSecret,
		Brokers:    kafkaBrokers,
		EventTopic: eventTopic,
		DBConnStr:  dbConnStr,
	}, nil
}
