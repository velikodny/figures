package main

import (
	"github.com/TestProjects/figures/figures"
)

type Circle struct {
	r int
}

type Square struct {
	side int
}



func main() {

	circle := new(Circle)
	square := new(Square)

	var array []figures.Figures
	array = append(array, circle, square)

	//array := make([]figures.Figures, 2)
	//array[0] = circle
	//array[1] = square
	figures.Print(array)


}

func (c *Circle) Square() int {
	return 3
}

func (s *Square) Square() int {
	return 1
}


//import "fmt"
//
//func main() {
//
//	//str := flag.String()
//	str := "Numbers: "
//	enType := 2
//
//	switch enType {
//	case 1:
//		str = encode1(str)
//	case 2:
//		str = encode2(str)
//	case 3:
//		str = encode3(str)
//	}
//
//	fmt.Printf("%v", str)
//
//}
//
//func encode1 (str string) string{
//	str = str + "1"
//	return str
//}
//
//func encode2 (str string) string{
//	str = str + "2"
//	return str
//}
//
//func encode3 (str string) string{
//	str = str + "3"
//	return str
//}
//
//func one () {
//
//}
//
//func two ( {
//
//}


