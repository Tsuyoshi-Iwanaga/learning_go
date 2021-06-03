package main

import (
	"github.com/Tsuyoshi-Iwanaga/learning_go/006_2/my"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	my.Migrate()
}
