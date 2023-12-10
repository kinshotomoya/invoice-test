package repository

import (
	"crypto/sha256"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type MysqlRepository struct {
	client *sql.DB
}

func NewMysqlRepository(viper *viper.Viper) (*MysqlRepository, error) {
	conf := mysql.Config{
		User:   viper.GetString("MYSQL_DB_USER"),
		Passwd: viper.GetString("MYSQL_DB_PASSWORD"),
		Net:    "tcp",
		Addr:   viper.GetString("MYSQL_DB_URL"),
		DBName: viper.GetString("MYSQL_DB_NAME"),
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{
		client: db,
	}, nil

}

func (m MysqlRepository) FindUser(email string, password string) (*User, error) {
	rows, err := m.client.Query("SELECT solt FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	var solt string
	for rows.Next() {
		err = rows.Scan(&solt)
		if err != nil {
			return nil, err
		}
	}

	hashedPassword := sha256.Sum256([]byte(password + solt))
	rows, err = m.client.Query("SELECT user_id, company_id FROM users WHERE email = ? AND password = ?", email, hashedPassword)
	if err != nil {
		return nil, err
	}

	var userId int
	var companyId int
	for rows.Next() {
		err = rows.Scan(&userId, &companyId)
		if err != nil {
			return nil, err
		}
	}
	return &User{
		userId:    userId,
		companyId: companyId,
	}, nil

}
