package main

import (
	"fmt"
)

type person struct {
	FstName string
	LstName string
	Age     int
	pinCode int
}
type contact struct {
	Email string
	Cell  int
}

type adress struct {
	HouseNo int
	Area    string
	state   string
}

type Employee struct {
	person_Details person
	person_contact contact
	person_adress  adress
}

func main() {

	emp()
	LAB()
}

func Structs() {
	var prince person
	fmt.Println("Default prince: ", prince)

	prince.FstName = "Ram"
	prince.LstName = "Bihari"
	prince.Age = 24
	prince.pinCode = 123123
	//fmt.Println("Prince person data: ", prince)

	person1 := person{
		FstName: "Akash",
		LstName: "Mane",
		Age:     23,
		pinCode: 654321,
	}
	fmt.Println("person1 person data: ", person1)
	//fmt.Println("person1 person data: ", person1.FstName)
	//fmt.Println("person1 person data: ", person1.pinCode)
}
func emp() {

	var emp1 Employee
	emp1.person_Details = person{
		FstName: "AAA",
		LstName: "BBB",
		Age:     11,
		pinCode: 000,
	}
	emp1.person_contact.Email = "abc123@gmail.com"
	emp1.person_contact.Cell = 075077664354

	emp1.person_adress = adress{
		HouseNo: 01,
		Area:    "Kartik Nagar",
		state:   "Karnataka",
	}

	fmt.Println("EMP1 Detaails: ", emp1.person_contact.Email)

}

type Books struct {
	name   string
	count  int
	Author string
}
type BooksPublish struct {
	date string
	day  string
	time string
}
type BooksCells struct {
	Price1Book   int
	Price10Book  int
	Price100Book int
}

type Lab struct {
	Book_Name    Books
	Book_Publish BooksPublish
	Book_Prices  BooksCells
}

func LAB() {
	var book1 Lab

	book1.Book_Name = Books{
		name:   "GO",
		count:  20,
		Author: "AR Rehman",
	}
	book1.Book_Publish = BooksPublish{
		date: "12/12/2002",
		day:  "Monday",
		time: "9:31AM",
	}
	book1.Book_Prices = BooksCells{
		Price1Book:   100,
		Price10Book:  900,
		Price100Book: 8000,
	}

	fmt.Println("Book1 detail in lab: ", book1)
}
