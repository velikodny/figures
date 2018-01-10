package figures

import "fmt"

type Figures interface {
	Square() int
}

func Print(array []Figures){

	var sum int
	sum = 0

	for  _, f := range array {
		sum = sum + f.Square()
	}

	fmt.Printf("%v", sum)
}
