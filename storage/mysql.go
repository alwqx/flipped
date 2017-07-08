package storage

import (
	"database/sql"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/config"
	_ "github.com/go-sql-driver/mysql"
)

var onceInitDB sync.Once

type MysqlClient struct {
	Address  string
	Port     string
	User     string
	Password string
	DBName   string
	DBPath   string
}

func NewMysqlClient() *MysqlClient {
	fc := config.GetConfig("")
	mc := &MysqlClient{
		Address:  fc.Mysql.Address,
		Port:     fc.Mysql.Port,
		User:     fc.Mysql.User,
		Password: fc.Mysql.Password,
		DBName:   fc.Mysql.DBName,
	}

	mc.initialize()

	return mc
}

func (mc *MysqlClient) initialize() {
	mc.DBPath = mc.User + ":" + mc.Password +
		"@tcp(" + mc.Address + ":" +
		mc.Port + ")/"

	logrus.Infof("dbpath is %s", mc.DBPath)

	useDB := `USE ` + mc.DBName + ";"
	dbSchema := `
		CREATE DATABASE IF NOT EXISTS ` + mc.DBName + ` 
			DEFAULT CHARACTER SET utf8mb4
			DEFAULT COLLATE utf8mb4_unicode_ci;
	`

	heartbeatTable := `
		CREATE TABLE IF NOT EXISTS ` + mc.DBName + `.heartbeat (
			id INT(64) NOT NULL AUTO_INCREMENT,
			timestamp VARCHAR(20) NOT NULL,
			heartrate INT(16) NOT NULL,
			PRIMARY KEY (id)
		);
	`
	heartstrengthTable := `
		CREATE TABLE IF NOT EXISTS ` + mc.DBName + `.heartstrength (
			id INT(64) NOT NULL AUTO_INCREMENT,
			timestamp VARCHAR(20) NOT NULL,
			heartstrength INT(16) NOT NULL,
			PRIMARY KEY (id)
		);
	`

	db := mc.CreateBareDB()
	defer db.Close()

	onceInitDB.Do(func() {
		if _, err := db.Exec(dbSchema); err != nil {
			logrus.Fatalf("create database %s error: %v", mc.DBName, err)
		}

		if _, err := db.Exec(useDB); err != nil {
			logrus.Fatalf("use db %s error: %v", mc.DBName, err)
		}

		if _, err := db.Exec(heartbeatTable); err != nil {
			logrus.Fatalf("create table heartbeat error: %v", err)
		}

		if _, err := db.Exec(heartstrengthTable); err != nil {
			logrus.Fatalf("create table heartstrength error: %v", err)
		}
	})
}

// CreateBareDB create *sql.DB without connecting to specific database
func (mc *MysqlClient) CreateBareDB() *sql.DB {
	db, err := sql.Open("mysql", mc.DBPath)
	if err != nil {
		logrus.Fatal("setup up db error: ", err)
	}

	return db
}

// CreateDB create *sql.DB connecting to specific database
func (mc *MysqlClient) CreateDB() *sql.DB {
	db, err := sql.Open("mysql", mc.DBPath+mc.DBName)
	if err != nil {
		logrus.Fatal("setup up db error:", err)
	}
	logrus.Info("create db with dbname ", mc.DBName)

	return db
}

// STMTFactory return *sql.Stmt and handle error
func (mc *MysqlClient) STMTFactory(predSQL string, db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare(predSQL)
	if err != nil {
		logrus.Fatalf("prepare sql %s error: %s", predSQL, err.Error())
	}
	return stmt
}

func (mc *MysqlClient) Insert(sql, dbname string, args ...interface{}) (sql.Result, error) {
	db := mc.CreateDB()
	defer db.Close()
	stmt := mc.STMTFactory(sql, db)
	defer stmt.Close()

	result, err := stmt.Exec(args...)

	return result, err
}

func (mc *MysqlClient) QueryOne(sql, dbname string, args ...interface{}) *sql.Row {
	db := mc.CreateDB()
	defer db.Close()
	stmt := mc.STMTFactory(sql, db)
	defer stmt.Close()

	return stmt.QueryRow(args...)
}

func (mc *MysqlClient) QueryMore(sql string, args ...interface{}) (*sql.Rows, error) {
	db := mc.CreateDB()
	defer db.Close()
	stmt := mc.STMTFactory(sql, db)
	defer stmt.Close()

	return stmt.Query(args...)
}
