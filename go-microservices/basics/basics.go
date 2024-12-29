package main

import (
	"fmt"
	"go-react-ecommerce-app/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

// The default format for %v is:

// bool:                    %t
// int, int8 etc.:          %d
// uint, uint8 etc.:        %d, %#x if printed with %#v
// float32, complex64, etc: %g
// string:                  %s
// chan:                    %p
// pointer:                 %p

type ProductDetails struct {
	Name  string
	Price float64
	Stock int
}

func (p *ProductDetails) Calculate(qty int) float64 {
	fmt.Printf("%p --- Calculate --- \n", &p)
	return p.Price * float64(qty)
}

func (p *ProductDetails) ReduceStock(qty int) {
	fmt.Printf("%p --- ReduceStock --- \n", &p)
	if p.Stock >= qty {
		p.Stock -= qty
	}

}

func main() {
	fmt.Println("I am main function")

	app := fiber.New()

	// MyHelperFunction()

	configs.LoadAppSettings()

	// Basic Types: int, float, string, bool

	var age int64

	var height float64

	// var firstName string

	var _ bool

	age = 34

	height = 175.5

	firstName := "Sreeprabhu"

	fmt.Println(age, height)

	fmt.Printf("My Age is: %v\n", age)
	fmt.Printf("My Height is: %v\n", height)
	fmt.Printf("My First Name is: %v\n", firstName)

	// Composite Types: array, map, slice, struct

	// Array

	// var myFamily [3]string

	// myFamily[0] = "Achan"
	// myFamily[1] = "Amma"
	// myFamily[2] = "Myself"

	myFamily := [3]string{"Achan", "Amma", "Myself"}

	myFamily[2] = "Rony"

	fmt.Printf("My Family: %v\n", myFamily)

	// Multi Dimensional Arrays

	myCourses := [2][2]string{
		{"Go", "React"},
		{"AWS", "GCP"},
	}

	fmt.Printf("Available Courses: %v\n", myCourses)

	// Slice

	var myFriends []string

	myFriends = append(myFriends, "Vipin")
	myFriends = append(myFriends, "Sajin")

	// myFriends = []string{"Achan", "Amma", "Myself"}

	fmt.Printf("My Friends: %v\n", myFriends)

	// Map

	myWishlist := make(map[string]int)

	myWishlist["first"] = 10
	myWishlist["second"] = 20

	delete(myWishlist, "second")
	firstWish := myWishlist["first"]
	log.Println(firstWish)

	fmt.Printf("My Wishlist: %v\n", myWishlist)

	// Struct

	type Details struct {
		Description string
		images      string
	}

	type Product struct {
		Name  string `json:"product_name"`
		Price float64
		Details
	}

	var product Product
	product = Product{
		Name:  "MacPro",
		Price: 9000.50,
		Details: Details{
			Description: "An incredible machine",
			images:      "http://macproimage.jpg",
		},
	}

	product.Name = "Macbook Pro"

	fmt.Println("Product struct: %v\n", product)

	// Pointer Types: *

	/*
		- A pointer is a variable that holds memory address
		- Jay -> Laptop
		- Guest -> Jay -> Laptop
	*/

	jay := "laptop"
	fmt.Println(jay)
	fmt.Println(&jay)
	var guest *string

	guest = &jay // now guest holds the address of jay

	fmt.Println(guest)
	fmt.Println(*guest)

	// Conditional Statements

	/*
	 - if else
	 - switch case
	 - select
	*/

	// if else
	if age > 65 {
		fmt.Println("Senior Citizen")
	} else if age > 17 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}

	// switch
	seatClass := "firstClass"

	switch seatClass {
	case "FirstClass":
		fmt.Println("You will get free drinks")

	case "BusinessClass":
		fmt.Println("You will get more leg room")

	default:
		fmt.Println("You need to pay for services")
	}

	// for loop

	for i := 0; i < 5; i++ {
		fmt.Println("I am a for loop!")

		myNewFriend := fmt.Sprintf("Friend %d", i)
		myFriends = append(myFriends, myNewFriend)
	}

	fmt.Println(myFriends)

	for index, value := range myFriends {
		fmt.Println(index, value)
	}

	// loop until a condition met

	// isOver := 0

	// for {
	// 	isOver++

	// 	fmt.Println(isOver)
	// 	if isOver > 99 {
	// 		fmt.Println("it's really over now")
	// 		return
	// 	}
	// }

	// infinite loop

	// for {
	// 	fmt.Println("You are in an infinite loop!")
	// }

	// Recceiver Function

	p := ProductDetails{
		Name:  "Macbook Pro",
		Price: 9000,
		Stock: 5,
	}

	fmt.Printf("%p --- 1 ---\n", &p)

	fmt.Printf("Total Amount: %f\n", p.Calculate(2))

	fmt.Printf("%p --- 2 --- \n", &p)
	p.ReduceStock(2)
	fmt.Printf("%p --- 3 --- \n", &p)

	fmt.Printf("Updated Stock: %v\n", p.Stock)

	app.Listen("localhost:9000")
}
