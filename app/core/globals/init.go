package globals

func Init() {
	InitConfig()
	cfg := GetConfig()

	InitHTML(cfg)
	InitJWTAuth(cfg)
	InitMongoSession(cfg)
}
