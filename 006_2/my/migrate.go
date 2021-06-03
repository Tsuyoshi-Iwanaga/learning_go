package my

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Migrate
func Migrate() {
	db, er := gorm.Open("sqlite3", "data.db")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})
}
