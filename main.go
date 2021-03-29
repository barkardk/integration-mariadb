package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	MariaDBRootUser     = "root"
	MariaDBRootPassword = "secret"
	MariaDBHost         = "127.0.0.1"
	MariaDBClientPort   = "3306"
)

type Connection struct {
	MariaDBRootUser     string
	MariaDBRootPassword string
	MariaDBHost         string
	MariaDBClientPort   string
}

var logLevel = "info"

type Canines struct {
	Id               int64
	Breed            string
	IsHypoAllergenic bool
	LifeExpectancy   int32
	Origin           string
}

type CanineModel struct {
	Db *sql.DB
}

func getEnv(key string, fallBack string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallBack
}

func getDatabaseConnection(dbName string, c *Connection) (conn *sql.DB, err error) {
	cStr := c.MariaDBRootUser + ":" + c.MariaDBRootPassword + "@tcp(" + c.MariaDBHost + ":" + c.MariaDBClientPort + ")/"
	log.Debugf("connection string %v", cStr)
	conn, err = sql.Open("mysql", cStr)
	if err != nil {
		log.Fatalf("failed to connect to to mysql %v", err)
	}
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatalf("could not create database %v err: %v", dbName, err)
	}
	return
}

func (canineModel CanineModel) insertCanines(canines *Canines) (int64, error) {
	log.Debugf("Insert values %v, %v, %v, %v", canines.Breed, canines.IsHypoAllergenic, canines.LifeExpectancy, canines.Origin)
	result, err := canineModel.Db.Exec("INSERT INTO dogsDB.CANINES(breed, isHypoAllergenic, lifeExpectancy, origin) values(?,?,?,?)", canines.Breed, canines.IsHypoAllergenic, canines.LifeExpectancy, canines.Origin)
	if err != nil {
		log.Errorf("Unexpected result trying to insert entries %v", err)
	} else {
		canines.Id, _ = result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
	return -1, nil
}

func (canineModel CanineModel) createTable(name string) (int64, error) {
	log.Debugf("CREATE TABLE IF NOT EXISTS %v: ", name)
	result, err := canineModel.Db.Exec("CREATE TABLE IF NOT EXISTS " + name + "( " +
		"Id INT NOT NULL PRIMARY KEY AUTO_INCREMENT," +
		"Breed VARCHAR(150) NOT NULL," +
		"IsHypoAllergenic BOOL NOT NULL," +
		"LifeExpectancy INT," +
		"Origin VARCHAR(150) ) ENGINE=INNODB;")

	if err != nil {
		log.Errorf("Encountered issues creating table %v err: %v", name, err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
	return -1, nil
}

func main() {
	logLevel, _ = os.LookupEnv("LOG_LEVEL")
	ll, err := log.ParseLevel(logLevel)
	if err != nil {
		ll = log.DebugLevel
	}
	log.SetLevel(ll)
	c := Connection{
		MariaDBRootUser:     getEnv("MARIADB_ROOT_USER", MariaDBRootUser),
		MariaDBRootPassword: getEnv("MARIADB_ROOT_PASSWORD", MariaDBRootPassword),
		MariaDBHost:         getEnv("MARIADB_HOST", MariaDBHost),
		MariaDBClientPort:   getEnv("MARIADB_CLIENT_PORT", MariaDBClientPort),
	}
	dbName := "dogsDB"

	db, err := getDatabaseConnection(dbName, &c)
	if err != nil {
		log.Fatalf("Could not connect to database %v:", err)
	} else {
		canineModel := CanineModel{
			Db: db,
		}
		_, err1 := canineModel.createTable(dbName + ".CANINES")
		if err1 != nil {
			log.Error(err1)
		}

		canines := Canines{
			Breed:            "Schnauzer",
			IsHypoAllergenic: true,
			LifeExpectancy:   14,
			Origin:           "Germany",
		}
		rowsAffected, err2 := canineModel.insertCanines(&canines)
		if err2 != nil {
			log.Error(err2)
		} else {
			log.Info("RowsAffected:", rowsAffected)
			log.Info("Canine Information")
			log.Info("Id:", canines.Id)
			log.Info("Breed", canines.Breed)
			log.Info("IsHypoAllergenic", canines.IsHypoAllergenic)
			log.Info("LifeExpectancy", canines.LifeExpectancy)
			log.Info("Origin", canines.Origin)

		}
	}

}
