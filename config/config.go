package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	MAX_GOROUTINES int `env:"MAX_GOROUTINES,default=10"`

	/*
		Database struct {
			Host     string `env:"DATABASE_HOST,required"`
			Port     int    `env:"DATABASE_PORT,default=5432"`
			User     string `env:"DATABASE_USER,required"`
			Password string `env:"DATABASE_PASSWORD,required"`
			DbName   string `env:"DATABASE_DB_NAME,required"`
		}*/
}

var C config

// ReadConfig read config
func ReadConfig() error {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

/*
// PgConn the connection string to the pg database
func PgConn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DbName)
}

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
