package config

type Config struct {
	Dir_Asset string
}

func Set() Config {
	return Config{
		Dir_Asset: "dist",
	}
}
