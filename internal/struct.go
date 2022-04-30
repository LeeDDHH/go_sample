package internal

import (
	"fmt"
	"strconv"
)

type Mydata struct {
	Name string
	Data []int
}

type intp int

var mydata struct {
	Name string
	Data []int
}

func Struct(){
	mydata.Name = "Taro"
	mydata.Data = []int{12, 20, 30}
	fmt.Println(mydata)
}

func Struct2(){
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{"Hanako", []int{90, 80, 70}}

	fmt.Println(taro)
	fmt.Println(hanako)
}

func StructPassByValue(){
	taro := Mydata{"Taro", []int{10, 20, 30}}
	fmt.Println(taro)
	taro = structValueRev(taro)
	fmt.Println(taro)
}

func structValueRev(md Mydata) Mydata {
	od := md.Data
	nd := []int{}
	for i := len(od) -1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
	return md
}

func StructPassByReference(){
	taro := Mydata{"Taro", []int{10, 20, 30}}
	fmt.Println(taro)
	structReferenceRev(&taro)
	fmt.Println(taro)
}

func structReferenceRev(md *Mydata) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) -1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}

func NewAndMake() {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 5)
	fmt.Println(taro)
}

func ExecStructMethod() {
	taro := Mydata{"Hanako", []int{98,76,54,32,10}}
	taro.PrintMyData()
}

func (md Mydata) PrintMyData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ",  md.Name)
	fmt.Println("Data: ",  md.Data)
	fmt.Println("*** end ***")
}

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n/2); i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}

func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

func (num *intp) doPromise() {
	pf := num.PrimeFactor()
	*num = intp(pf[len(pf)-1])
}

func ExtensionType() {
	s := Input("type a number")
	n, _ := strconv.Atoi(s)

	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())

	x.doPromise()
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
