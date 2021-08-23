package model_test

import "todo/model"

func InitMySQL() {
	model.MySQL.Database = "todo_test"
	model.MySQL.Password = "daisuki"
	model.MySQL.User = "root"
	model.MySQL.Host = "127.0.0.1"
}
