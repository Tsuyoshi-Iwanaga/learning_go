package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
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

	nm := input("name")
	ml := input("mail")
	age := input("age")
	ag, _ := strconv.Atoi(age)

	qry := "insert into mydata (name, mail, age) values (?, ?, ?)"
	con.Exec(qry, nm, ml, ag)
	showRecord(con)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg, ": ")
	scanner.Scan()
	return scanner.Text()
}

func showRecord(con *sql.DB) {
	qry := "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str2())
	}
}

func mydatafmRws(rs *sql.Rows) *Mydata2 {
	var md Mydata2
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
