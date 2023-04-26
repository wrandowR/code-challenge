package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	MAX_GOROUTINES int `env:"MAX_GOROUTINES,default=10"`

	FromEmail string `env:"FROM_EMAIL,required"`
	Database  struct {
		Host     string `env:"DATABASE_HOST,required"`
		Port     int    `env:"DATABASE_PORT,default=5432"`
		User     string `env:"DATABASE_USER,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
		DbName   string `env:"DATABASE_DB_NAME,required"`
	}
}

var c config

// ReadConfig read config
func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	return err
}

// PgConn the connection string to the pg database
func PgConn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DbName)
}

/*

// PgConnMigration returns the config string for migration
func PgConnMigration() *string {
	if c.Migrate {
		pgconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.DbName)
		return &pgconn
	}

	return nil
}*/

//

func MaxGoroutines() int {
	return c.MAX_GOROUTINES
}

func FromEmail() string {
	return c.FromEmail
}

/*
// SetConfigs set configs
func SetConfigs(Host string, DbName string, User string, Password string, Migrate bool) error {

	c.Database.Host = Host
	c.Database.DbName = DbName
	c.Database.User = User
	c.Database.Password = Password
	c.Migrate = Migrate

	c.CurrencyLayerAPIKEY = "test"

	err := ReadConfig()
	if err != nil {
		return err
	}

	return nil
}
*/
