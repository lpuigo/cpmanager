package config

type Config struct {
	DirAsset               string
	DirPersisterConsultant string
	SessionKey             string
}

func Set() Config {
	return Config{
		DirAsset:               "dist",
		DirPersisterConsultant: "ressources/consultant",
		SessionKey:             "cpmanager-session-key-replace-in-production",
	}
}
