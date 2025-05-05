package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 匿名导入 SQLite 驱动
)

type Employees struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int64  `db:"salary"`
}

func main() {
	db, err := sqlx.Connect("sqlite3", "./practice.db")
	if err != nil {
		panic("无法连接数据库")
	}
	err = db.Ping()
	if err != nil {
		panic("数据库不可达到")
	}
	fmt.Println("成功连接到数据库")

	//emp:=Employees{Name:"dami",Department: "科技部门", Salary:5000}

	//err = createTable(db)
	//if err != nil {
	//	panic("建表失败：" + err.Error())
	//}
	//_, err = db.NamedExec(`INSERT INTO employees (name,department,salary) VALUES ("yummy","科技部",5000)`, &Employees{})
	//if err != nil {
	//	panic("插入失败：" + err.Error())
	//}
	//
	var employees []Employees
	err = db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		panic("查询失败" + err.Error())
	}
	for _, vel := range employees {

		fmt.Printf("姓名：%s,部门：%s，薪酬:%d\n", vel.Name, vel.Department, vel.Salary)
	}

	err = db.Select(&employees, "SELECT * FROM employees ORDER BY  salary DESC  limit 1  ")
	if err != nil {
		panic("查询失败" + err.Error())
	}

	fmt.Printf("姓名：%s,部门：%s，薪酬:%d\n", employees[0].Name, employees[0].Department, employees[0].Salary)

}

func createTable(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS  employees(
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  name TEXT NOT NULL,
		  department TEXT NOT NULL,
		  salary BIGINT NOT NULL
		)`)
	return err
}
