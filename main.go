package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	MariaDBRootUser     = "root"
	MariaDBRootPassword = "secret"
	MariaDBHost         = "127.0.0.1"
	MariaDBClientPort   = "3306"
	logLevel            = "info"
	connStr             = MariaDBRootUser + ":" + MariaDBRootPassword + "@tcp(" + MariaDBHost + ":" + MariaDBClientPort + ")/"
)

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

func getDatabaseConnection(name string) (conn *sql.DB, err error) {
	conn, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to connect to to mysql %v", err)
	}
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS" + name)
	if err != nil {
		log.Fatalf("could not create database %v err: %v", name, err)
	}
	return
}

func (canineModel CanineModel) insertCanines(canines *Canines) (int64, error) {
	log.Debugf("Insert values %v, %v, %v, %v", canines.Breed, canines.IsHypoAllergenic, canines.LifeExpectancy, canines.Origin)
	result, err := canineModel.Db.Exec("INSERT INTO CANINES(breed, isHypoAllergenic, lifeExpectancy, origin) values(?,?,?,?)", canines.Breed, canines.IsHypoAllergenic, canines.LifeExpectancy, canines.Origin)
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
	log.Debugf("Create table %v: ", name)
	result, err := canineModel.Db.Exec("CREATE TABLE IF NOT EXISTS " + name + " ( " +
		"'Id' INT NOT NULL PRIMARY KEY AUTO_INCREMENT," +
		"'Breed' VARCHAR(150) NOT NULL," +
		"'IsHypoAllergenic' BOOL NOT NULL" +
		"'LifeExpectancy' INT" +
		"'Origin' VARCHAR(150)" +
		")")
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
	MariaDBRootUser, _ = os.LookupEnv("MARIADB_ROOT_USER")
	MariaDBRootPassword, _ = os.LookupEnv("MARIADB_ROOT_PASSWORD")
	MariaDBHost, _ = os.LookupEnv("MARIADB_HOST")
	MariaDBClientPort, _ = os.LookupEnv("MARIADB_CLIENT_PORT")

	log.Debugf("%v, %v, %v, %v", MariaDBRootPassword, MariaDBRootUser, MariaDBHost, MariaDBClientPort)

	db, err := getDatabaseConnection("dogsDB")
	if err != nil {
		log.Fatalf("Could not connect to database %v:", err)
	} else {
		canineModel := CanineModel{
			Db: db,
		}
		_, err1 := canineModel.createTable("Canines")
		if err1 != nil {
			log.Error(err1)
		}
		canines := Canines{
			Breed: "Schnauzer",
			IsHypoAllergenic: true,
			LifeExpectancy: 14,
			Origin: "Germany",
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
