package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(cnf *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := cnf.GetString("DATABASE_HOST")
	port := cnf.GetInt("DATABASE_PORT")
	username := cnf.GetString("DATABASE_USERNAME")
	password := cnf.GetString("DATABASE_PASSWORD")
	name := cnf.GetString("DATABASE_NAME")
	maxConn := cnf.GetInt("DATABASE_MAX_CONNECTION")
	idleConn := cnf.GetInt("DATABASE_IDLE_CONNECTION")
	timeMaxLifeConn := cnf.GetInt("DATABASE_TIME_MAX_CONNECTION")
	idleMaxTimeConn := cnf.GetInt("DATABASE_TIME_IDLE_CONNECTION")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)
	dbGormConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		},
		),
	})

	if err != nil {
		log.Warnf("Error while connect to database : %v\n", err)

	}

	db, err := dbGormConnection.DB()
	if err != nil {
		log.Warnf("Error while connect to database : %v\n", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(idleConn)
	db.SetConnMaxLifetime(time.Duration(timeMaxLifeConn) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(idleMaxTimeConn) * time.Minute)

	return dbGormConnection
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
