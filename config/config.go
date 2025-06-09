package config

type Config struct {
	DirAsset               string
	DirPersisterConsultant string
	DirPersisterUser       string
	SessionKey             string
}

func Set() Config {
	return Config{
		DirAsset:               "dist",
		DirPersisterConsultant: "ressources/consultant",
		DirPersisterUser:       "ressources/user",
		SessionKey:             "cpmanager-session-key-replace-in-production",
	}
}
