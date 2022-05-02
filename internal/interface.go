package internal

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Datafunc interface {
	Initial(name string, data []int)
	PrintData()
}

type DefaultData interface {
	SetValue(vals map[string]string)
	PrintData()
}

type SampleData struct {
	Name string
	Data []int
}

type Sample2Data struct {
	Name string
	Mail string
	Age int
}

type General interface {}
type GData interface {
	Set(nm string, g General)
	Print()
}
type GDataImpl struct {
	Name string
	Data General
}

type FData interface {
	Set(nm string, g General) FData
	Print()
}

type NData struct {
	Name string
	Data int
}

type SData struct {
	Name string
	Data string
}

type AIData struct {
	Name string
	Data []int
}

type ASData struct {
	Name string
	Data []string
}



func (sm *SampleData) Initial(name string, data []int) {
	sm.Name = name
	sm.Data = data
}

func (sm *SampleData) Check() {
	fmt.Printf("Check! [%s]", sm.Name)
}

func (sd *SampleData) SetValue(vals map[string]string) {
	sd.Name = vals["name"]
	var valt = []string{}
	if strings.Contains(vals["data"], ",") {
		valt = strings.Split(vals["data"], ", ")
	} else {
		valt = strings.Split(vals["data"], " ")
	}
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	sd.Data = vali
}

func (sm *SampleData) PrintData() {
	if sm != nil {
		fmt.Println("Name: ", sm.Name)
		fmt.Println("Data: ", sm.Data)
	} else {
		fmt.Println("** This is Nil value. **")
	}
}

func (sd *Sample2Data) SetValue(vals map[string]string) {
	sd.Name = vals["name"]
	sd.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	sd.Age = n
}

func (sm *Sample2Data) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", sm.Name, sm.Age)
	fmt.Printf("mail: %s.\n", sm.Mail)
}

func (gd *GDataImpl) Set(nm string, g General) {
	gd.Name = nm
	gd.Data = g
}

func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>> ", gd.Name)
	fmt.Println(gd.Data)
}

func (nd *NData) Set(nm string, g General) FData {
	nd.Name = nm
	if reflect.TypeOf(g).Kind() == reflect.Int {
		nd.Data = g.(int)
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

func (sd *SData) Set(nm string, g General) FData {
	sd.Name = nm
	if reflect.TypeOf(g).Kind() == reflect.String {
		sd.Data = g.(string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func (aid *AIData) Set(nm string, g General) FData {
	aid.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf(0)) {
		aid.Data = g.([]int)
	}
	return aid
}

func (aid *AIData) Print() {
	fmt.Printf("* %s %d *\n", aid.Name, aid.Data)
}

func (aid *ASData) Set(nm string, g General) FData {
	aid.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) {
		aid.Data = g.([]string)
	}
	return aid
}

func (aid *ASData) Print() {
	fmt.Printf("* %s %s *\n", aid.Name, aid.Data)
}





func InterfaceInitByStructType() {
	var ob SampleData = SampleData{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
	ob.Check()
}

func InterfaceInitByInterfaceType() {
	var ob Datafunc = new(SampleData)
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
}

func InterfaceInitWithOtherStructType() {
	ob := [2]DefaultData{}

	ob[0] = new(SampleData)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})

	ob[1] = new(Sample2Data)
	ob[1].SetValue(map[string]string{
		"name": "Mami",
		"mail": "mami@mume.mo",
		"age": "34",
	})

	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}

func NilReciverExample() {
	var ob *SampleData
	ob.PrintData()
	ob = &SampleData{}
	ob.SetValue(map[string]string{
		"name": "Jiro",
		"data": "123 456 789",
	})
	ob.PrintData()
}

func EmptyInterfaceExample() {
	var data = []GDataImpl{}
	data = append(data, GDataImpl{"Taro", 123})
	data = append(data, GDataImpl{"Hanako", "hello!"})
	data = append(data, GDataImpl{"Taro", []int{123, 456, 789}})

	for _, ob := range data {
		ob.Print()
	}
}

func TypeAssertionExample() {
	var data = []FData{}
	data = append(data, new(NData).Set("Taro", 123))
	data = append(data, new(SData).Set("Taro", "hello!"))
	data = append(data, new(NData).Set("Hanako", 98700))
	data = append(data, new(SData).Set("Sachiko", "happy?"))

	for _, ob := range data {
		ob.Print()
	}
}

func TypeAssertionArrayExample() {
	var data = []FData{}
	data = append(data, new(AIData).Set("Taro", []int{1, 2, 3}))
	data = append(data, new(ASData).Set("Taro", []string{"hello", "bye"}))
	data = append(data, new(AIData).Set("Hanako", 98700))
	data = append(data, new(ASData).Set("Sachiko", "happy?"))

	for _, ob := range data {
		ob.Print()
	}
}
