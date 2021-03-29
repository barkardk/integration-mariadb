package main

import (
	"fmt"
	"os"
)

const (
	MariaDBRootUser     = "root"
	MariaDBRootPassword = "secret"
	MariaDBHost         = "127.0.0.1"
	MariaDBClientPort   = "3306"
)

type MDB struct {
	MariaDBRootUser     string
	MariaDBRootPassword string
	MariaDBHost         string
	MariaDBClientPort   string
}

func GetMariadbConfig() *MDB {
	return &MDB{
		MariaDBRootUser:     MariaDBRootUser,
		MariaDBRootPassword: MariaDBRootPassword,
		MariaDBHost:         MariaDBHost,
		MariaDBClientPort:   MariaDBClientPort,
	}
}
/*func createDatabase(name string) error {
	return nil
}*/

func main() {
	m := MDB{
		MariaDBRootUser:     os.Getenv("MARIADB_ROOT_USER"),
		MariaDBRootPassword: os.Getenv("MARIADB_ROOT_PASSWORD"),
		MariaDBHost:         os.Getenv("MARIADB_HOST"),
		MariaDBClientPort:   os.Getenv("MARIADB_CLIENT_PORT"),
	}
	cf := GetMariadbConfig()
	fmt.Printf("hello\n %v %v", cf.MariaDBRootPassword, m.MariaDBRootPassword)
}
