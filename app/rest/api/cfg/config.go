package cfg

type Config struct {
	Addr string `long:"api-addr" env:"API_ADDR" default:"localhost:8080" description:"rest api address"`
}
