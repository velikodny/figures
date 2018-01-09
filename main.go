package main

type Shape struct {
	name string
}

type Figures interface {
	Init()
	Square()
}

func main() {

	return returnSquare(new(Shape),"treug", "kvadr")

}

func returnSquare (s *Shape, figuresName ...string){

}

func generateArrayFigures (s *Figures) Init(userShape ...string){

	for _, shapeName := range userShape  {

	}
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


