package repository

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strconv"
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

func (m MysqlRepository) FindUser(sessionId string) (*User, error) {
	id, err := strconv.Atoi(sessionId)
	if err != nil {
		return nil, err
	}
	rows, err := m.client.Query("SELECT user_id, company_id FROM users WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}

	users := make([]User, 1)
	for rows.Next() {
		var userId int
		var companyId int
		err = rows.Scan(&userId, &companyId)
		if err != nil {
			return nil, err
		}
		users[0] = User{
			userId:    userId,
			companyId: companyId,
		}
	}

	return &users[0], nil

}
