package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadService(cfgFile string) *AppServiceConfig {
	if cfgFile == "" {
		panic("cfgFile 不能为空")
	}
	if err := loadFromFile(cfgFile); err != nil {
		panic(err)
	}

	return &AppServiceConfig{
		&RepoConfig{
			&DBRepoConfig{
				DBHost: getStrEnvOrDefault("DB_HOST", "localhost"),
				DBPort: getIntEnvOrDefault("DB_PORT", 3306),
				DBUser: getStrEnvOrDefault("DB_USER", "root"),
				DBPass: getStrEnvOrDefault("DB_PASS", "root"),
				DBName: getStrEnvOrDefault("DB_TABLE_NAME", "test_db"),
			},
			&CacheRepoConfig{
				CacheHost: getStrEnvOrDefault("CACHE_HOST", "localhost"),
				CachePort: getIntEnvOrDefault("CACHE_PORT", 6379),
				CacheUser: getStrEnvOrDefault("CACHE_USER", ""),
				CachePass: getStrEnvOrDefault("CACHE_PASS", ""),
			},
		},
		&ServiceConfig{
			ServiceAddress:  getStrEnvOrDefault("SERVICE_ADDRESS", "localhost"),
			ServiceGrpcPort: getIntEnvOrDefault("SERVICE_GRPC_PORT", 9090),
			ServiceDesc:     getStrEnvOrDefault("SERVICE_DESC", "暂无描述"),
			ServiceID:       getIntEnvOrDefault("SERVICE_ID", -1),
			ServiceName:     getStrEnvOrDefault("SERVICE_NAME", "暂无名称"),
		},
		&CosConfig{
			CosBucket:    getStrEnvOrDefault("COS_BUCKET", ""),
			CosHost:      getStrEnvOrDefault("COS_HOST", ""),
			CosSecretID:  getStrEnvOrDefault("COS_SECRET_ID", ""),
			CosSecretKey: getStrEnvOrDefault("COS_SECRET_KEY", ""),
		},
	}
}

func LoadGateway(cfgFile string) *AppGatewayConfig {
	asc := LoadService(cfgFile)
	agc := &AppGatewayConfig{
		AppServiceConfig: asc,
		GatewayConfig: &GatewayConfig{
			GatewayHttpPort: getIntEnvOrDefault("GATEWAY_HTTP_PORT", 8080),
		},
	}
	return agc
}

func loadFromFile(cfg string) error {
	if err := godotenv.Load(cfg); err != nil {
		return err
	}
	return nil
}

func getStrEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getIntEnvOrDefault(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return defaultValue
}

func getFloatEnvOrDefault(key string, defaultValue float64) float64 {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			return v
		}
	}
	return defaultValue
}

func getBoolEnvOrDefault(key string, defaultValue bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.ParseBool(value); err == nil {
			return v
		}
	}
	return defaultValue
}
