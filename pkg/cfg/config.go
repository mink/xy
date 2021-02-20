package cfg

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"os"
	"strconv"
)

type Config struct {
	Host string
	Port int
	Database Database
}

type Database struct {
	User string
	Pass string
	Host string
	Port int
	Name string
	CharSet string
}

func (config Config) SocketAddress() string {
	return config.Host + ":" + fmt.Sprint(config.Port)
}

func (db Database) DataSource() string {
	return db.User +":" + db.Pass + "@(" + db.Host + ":" + fmt.Sprint(db.Port) + ")/" + db.Name + "?charset=" + db.CharSet
}

func LoadDatabase(config *Config) {
	config.Database = Database{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
		CharSet: os.Getenv("DB_CHARSET"),
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		fmt.Println("Error parsing DB_PORT environment variable")
		os.Exit(1)
	}

	config.Database.Port = port

	_ = orm.RegisterDataBase(
		"default",
		"mysql", "root:@(127.0.0.1:3306)/special-chainsaw?charset=utf8",
	)
}
