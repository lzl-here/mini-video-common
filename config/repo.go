package config

type RepoConfig struct {
	*DBRepoConfig
	*CacheRepoConfig
}

type DBRepoConfig struct {
	DBHost      string
	DBPort      int
	DBUser      string
	DBPass      string
	DBTableName string
}

type CacheRepoConfig struct {
	CacheHost string
	CachePort int
	CacheUser string
	CachePass string
}
