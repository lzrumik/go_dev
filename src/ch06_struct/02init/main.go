package main

import (
	"fmt"
	"reflect"
	"time"
)

func main(){
	type Employee struct {
		ID        int
		Name,Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee

	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	position := &dilbert.Position
	fmt.Println(reflect.TypeOf(*position))

	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	fmt.Println(dilbert)
}
