package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	Name string
	Age int
	School string
	Experience string
}

func main() {

}
