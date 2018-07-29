package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func AddUserCredential(loginName string, pwd string) error {
	stmIns,err:dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")//指令预编译，防止数据库攻击
	if err!=nil{
		return err
	}
	stmIns.Exec(loginName,pwd)
	defer stmIns.Close()
	return nil
}

func GetUSerCredential(loginName string) (string, error) {
	stmOut,err:=dbConn.Prepare("SELETCT pwd FROM users WHERE login_name=?")
	if err!=nil{
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	stmOut.QueryRow(loginName).Scan(&pwd)
	defer stmOut.Close()
	return pwd,nil
}

func DeleteUser(loginName string,pwd string) error{
	stmDel,err:=dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err!=nil{
		log.Printf()
	}
}
