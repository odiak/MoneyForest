package config

import (
	"os"

	"strconv"

	"github.com/go-pg/pg"
)

var HostName string
var Port uint
var PgOptions *pg.Options

func init() {
	HostName = "localhost"
	Port = parseUint(os.Getenv("APP_PORT"), 8000)

	switch os.Getenv("APP_ENV") {
	case "test":
		PgOptions = &pg.Options{
			User:     "kaido",
			Addr:     "127.0.0.1:5432",
			Database: "money_forest_test",
		}
	default:
		PgOptions = &pg.Options{
			User:     "kaido",
			Addr:     "127.0.0.1:5432",
			Database: "money_forest",
		}
	}
}

func parseUint(str string, defaultValue uint) uint {
	if str != "" {
		n, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return uint(n)
	}
	return defaultValue
}
