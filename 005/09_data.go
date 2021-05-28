package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Mydata2 struct {
	ID   int
	Name string
	Mail string
	Age  int
}

func (m *Mydata2) Str2() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

func main() {
	con, er := sql.Open("sqlite3", "data.db")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	q := "select * from mydata"
	rs, er := con.Query(q)
	if er != nil {
		panic(er)
	}

	for rs.Next() {
		var md Mydata2
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str2())
	}
}
