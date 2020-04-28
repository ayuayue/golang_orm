package cmd_test

import (
	"fmt"
	"golang_orm"
)

func main() {
	engine, _ := golang_orm.NewEngine("sqlite3", "go_orm.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop tables if exists user;").Exec()
	_, _ = s.Raw("create table user(name text);").Exec()
	_, _ = s.Raw("create table user(name text);").Exec()
	result, _ := s.Raw("insert into user(name),value(?),(?)", "TOM", "TIMI").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success ,%d affected\n", count)
}
