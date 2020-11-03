package test

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)
/**
 * @author
 * @date 2020/10/30 16:40
 * create by Goland
 * @version 1.0
 */

/**
orm
 */
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}
func init(){
	orm.RegisterDriver("mysql" , orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8", 30)
	//注册定义的model
	orm.RegisterModel(new(User))

	// 创建table
	orm.RunSyncdb("default", false, true)
}

func TestOrmM1(t *testing.T) {
	fmt.Println("A")
}


 
