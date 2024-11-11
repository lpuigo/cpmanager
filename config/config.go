package config

type Config struct {
	Dir_Images string
}

func Set() *Config {
	return &Config{
		Dir_Images: "dist/image",
	}
}
