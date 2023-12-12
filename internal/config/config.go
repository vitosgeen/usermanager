package config

import (
	"usermanager/internal/apperrors"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

const (
	postgresPrefix = "POSTGRES_"
	redisPrefix    = "REDIS_"
	jwtPrefix      = "JWT_"
)

type Config struct {
	Environment    string `env:"ENVIRONMENT,required"`
	LogLevel       string `env:"LOG_LEVEL,required"`
	Port           string `env:"PORT,required"`
	PortGrpc       string `env:"PORT_GRPC,required"`
	PortGrpcClient string `env:"PORT_GRPC_CLIENT,required"`
	Postgres       *PostgresConfig
	Redis          *RedisConfig
	Jwt            *JwtConfig
}

type PostgresConfig struct {
	PostgresHost     string `env:"HOST,required"`
	PostgresPort     string `env:"PORT,required"`
	PostgresUser     string `env:"USER,required"`
	PostgresPass     string `env:"PASS,required"`
	PostgresDBName   string `env:"DBNAME,required"`
	PostgresTimezone string `env:"TIMEZONE,required"`
}

type RedisConfig struct {
	RedisHost      string `env:"HOST,required"`
	RedisPort      string `env:"PORT,required"`
	RedisUser      string `env:"USER,required"`
	RedisPassword  string `env:"PASS,required"`
	RedisDB        string `env:"DBNAME,required"`
	RedisDefaultdb string `env:"DBNAME,required"`
	Password       string `env:"PASS,required"`
}

type JwtConfig struct {
	Secret string `env:"SECRET,required"`
	Ttl    int    `env:"TTL,required"`
}

func NewConfig(envStr string) (*Config, error) {
	err := godotenv.Load(envStr)
	if err != nil {
		return nil, apperrors.EnvConfigLoadError.AppendMessage(err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return cfg, apperrors.EnvConfigParseError.AppendMessage(err)
	}

	postgresCfg := &PostgresConfig{}
	opts := env.Options{
		Prefix: postgresPrefix,
	}
	if err := env.ParseWithOptions(postgresCfg, opts); err != nil {
		return cfg, apperrors.EnvConfigPostgresParseError.AppendMessage(err)
	}
	cfg.Postgres = postgresCfg

	redisCfg := &RedisConfig{}
	opts = env.Options{
		Prefix: redisPrefix,
	}
	if err := env.ParseWithOptions(redisCfg, opts); err != nil {
		return cfg, apperrors.EnvConfigRedisParseError.AppendMessage(err)
	}
	cfg.Redis = redisCfg

	jwtCfg := &JwtConfig{}
	opts = env.Options{
		Prefix: jwtPrefix,
	}
	if err := env.ParseWithOptions(jwtCfg, opts); err != nil {
		return cfg, apperrors.EnvConfigJwtParseError.AppendMessage(err)
	}
	cfg.Jwt = jwtCfg
	return cfg, nil
}
