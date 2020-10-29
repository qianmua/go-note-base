package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

/**
 * @author
 * @date 2020/10/29 16:56
 * create by Goland
 * @version 1.0
 */ 
 
/**
sql
 */


func TestDb(t *testing.T) {

	/**

	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	 */
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select uid ,name , address , addrName from users")
	defer rows.Close()
	checkErr(err)
	for rows.Next(){
		var (
			uid *string
			name *string
			address *string
			addrName *string
		)
		err := rows.Scan(&uid, &name, &address, &addrName)
		checkErr(err)
		fmt.Println(*uid)
		//fmt.Printf(" , %v , %v , %v \n"  , *name , *address , *addrName)
	}

}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

