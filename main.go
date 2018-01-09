package main

type Circle struct {
	r int
}

type Square struct {
	side int
}

type Figures interface {
	Square()
}

func main() {

	circle := new(Circle)
	square := new(Square)

	array := make([]Figures, 2)
	array[0] = circle
	array[0] = square


	//figures := generateArray(circle, square)
	//return returnSquare(new(Shape),"treug", "kvadr")

}

func (c *Circle) Square() {

}

func (s *Square) Square() {

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


