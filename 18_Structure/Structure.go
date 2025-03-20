package main

import "fmt"

// creating a "Person" structure
type Person struct {
	FirstName string
	LastName  string
	age       int
}

type contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

// Employee Structute can store multiple type of structure
type Employee struct {
	Person_Details Person
	Person_Contact contact
	Person_Address Address
}

func main() {
	//creating a person
	var tom Person
	fmt.Println("Person tom: ", tom)
	// adding value in person
	tom.FirstName = "Tom"
	tom.LastName = "Cruise"
	tom.age = 60
	fmt.Println("Person tom:", tom)

	//creating person in another way
	person1 := Person{
		FirstName: "Keanu",
		LastName:  "Reeves",
		age:       60,
	}
	fmt.Println("Person person1:", person1)

	// we can also create using new keyword
	var person2 = new(Person)
	person2.FirstName = "Jhonny"
	person2.LastName = "Depp"
	person2.age = 62
	fmt.Println("Person person2:", person2) // This will print like a pointer

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------")

	var employee Employee
	employee.Person_Details = Person{
		FirstName: "Satya",
		LastName:  "Nadella",
		age:       62,
	}

	employee.Person_Contact.Email = "snceo@microsoft.com"
	employee.Person_Contact.Phone = "+1 56982 34789"

	employee.Person_Address = Address{
		House: 99,
		Area:  "Bay Area",
		State: "Sanfrancisco",
	}

	fmt.Println("Access all information togeather: ", employee)
	fmt.Println("Access Person contact: ", employee.Person_Contact)
	fmt.Println("Access House number from adress: ", employee.Person_Address.House)
}
