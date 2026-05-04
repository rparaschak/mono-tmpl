package storage

type Config struct {
	Endpoint        string   `env:"STORAGE_ENDPOINT"         envDefault:"localhost:5004"`
	AccessKeyID     string   `env:"STORAGE_ACCESS_KEY_ID"    envDefault:"minioadmin"`
	SecretAccessKey string   `env:"STORAGE_SECRET_ACCESS_KEY" envDefault:"minioadmin"`
	UseSSL          bool     `env:"STORAGE_USE_SSL"          envDefault:"false"`
	Region          string   `env:"STORAGE_REGION"           envDefault:"us-east-1"`
	Buckets         []string `env:"STORAGE_BUCKETS"          envDefault:"files"`
}
