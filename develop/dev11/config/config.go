package config

import (
	"flag"
	"log"

	"github.com/go-ini/ini"
)

// Server хранит информацию о поднимаемом сервере.
type Server struct {
	Port string
}

// GetServerConf загружает конфигурацию сервера из файла filename
// и возвращает структуру.
func GetServerConf(filename string) Server {
	cfg, err := ini.Load(filename)
	if err != nil {
		log.Fatalf("unable to load server configs from file %s: %v\n",
			filename, err.Error())
	}

	servCfg := Server{}

	err = cfg.MapTo(&servCfg)
	if err != nil {
		log.Fatalf("unable to map server configs from file %s: %v\n",
			filename, err.Error())
	}

	return servCfg
}

// DBConnection хранит информацию для подключения к базе данных.
type DBConnection struct {
	Host     string
	Port     string
	DBName   string
	Login    string
	Password string
}

// GetDBConnectionConf парсит флаги запуска программы
// и возвращает структуру с информацией для подключения
// к базе данных.
func GetDBConnectionConf() DBConnection {
	dbHost := flag.String("dbhost", "", "")
	dbPort := flag.String("dbport", "", "")
	dbName := flag.String("dbname", "", "")
	dbLogin := flag.String("dblogin", "", "")
	dbPassword := flag.String("dbpassword", "", "")

	flag.Parse()

	return DBConnection{
		Host:     *dbHost,
		Port:     *dbPort,
		DBName:   *dbName,
		Login:    *dbLogin,
		Password: *dbPassword,
	}
}
