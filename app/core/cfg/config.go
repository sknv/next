package cfg

import (
	"fmt"
	"time"
)

const (
	modeRelease = "release"
)

type Config struct {
	Mode string `long:"mode" env:"MODE" default:"debug" description:"application mode"`

	MailBaseURL  string `long:"mail-base-url" env:"MAIL_BASE_URL" required:"true" description:"base url to prepend to all urls in mail"`
	MailFrom     string `long:"mail-from" env:"MAIL_FROM" required:"true" description:"mail from header"`
	MailHost     string `long:"mail-host" env:"MAIL_HOST" required:"true" description:"smtp host"`
	MailPort     string `long:"mail-port" env:"MAIL_PORT" required:"true" description:"smtp port"`
	MailUsername string `long:"mail-username" env:"MAIL_USERNAME" required:"true" description:"mail username"`
	MailPassword string `long:"mail-password" env:"MAIL_PASSWORD" required:"true" description:"mail password"`

	MongoAddrs    []string      `long:"mongo-addrs" env:"MONGO_ADDRS" env-delim:"," required:"true" description:"mongo server addresses"`
	MongoDatabase string        `long:"mongo-database" env:"MONGO_INITDB_DATABASE" required:"true" description:"mongo database"`
	MongoUsername string        `long:"mongo-username" env:"MONGO_INITDB_ROOT_USERNAME" required:"true" description:"mongo username"`
	MongoPassword string        `long:"mongo-password" env:"MONGO_INITDB_ROOT_PASSWORD" required:"true" description:"mongo password"`
	MongoTimeout  time.Duration `long:"mongo-timeout" env:"MONGO_TIMEOUT" default:"45s" description:"time to wait for a mongo to respond"`

	SecretKey string `long:"secret-key" env:"SECRET_KEY" required:"true" description:"secret key"`
}

func (c *Config) IsRelease() bool {
	return c.Mode == modeRelease
}

func (c *Config) GetMailAddr() string {
	return fmt.Sprintf("%s:%s", c.MailHost, c.MailPort)
}
