package config

type Config struct {
	Dir_Images string
	Dir_Css    string
	Dir_Script string
}

func Set() *Config {
	return &Config{
		Dir_Images: "dist/images",
		Dir_Css:    "dist/css",
		Dir_Script: "dist/script",
	}
}
