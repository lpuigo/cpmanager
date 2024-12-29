package config

type Config struct {
	DirAsset               string
	DirPersisterConsultant string
}

func Set() Config {
	return Config{
		DirAsset:               "dist",
		DirPersisterConsultant: "ressources/consultant",
	}
}
