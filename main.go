package main

import "todo/model"

func main() {
	InitMySQL()
	model.CreateTableUser()
	model.CreateTableTask()
	Server()
}

func InitMySQL() {
	model.MySQL.Database = "todo"
	model.MySQL.Password = "daisuki"
	model.MySQL.User = "root"
	model.MySQL.Host = "127.0.0.1"
}
