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

	for true {
		s := input("id")
		if s == "" {
			break
		}

		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}

		rs, er := con.Query("select * from mydata where id = ?", n)
		if er != nil {
			panic(er)
		}
		defer con.Close()

		for rs.Next() {
			var md Mydata2
			er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
			if er != nil {
				panic(er)
			}
			fmt.Println(md.Str2())
		}
	}
	fmt.Println("*** end ***")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
