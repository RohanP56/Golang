# Structure or Struct
- More like creating class in java 
- In Go, a struct is a composite data type that groups togeather variable under single name.
- Each field in struct can have different data types and structs are used to create more complex data structure.
- **If we create a structure and don't assign any value then it will assign default values.**

``` Go
// creating a "Person" structure
type Person struct {
	FirstName string
	LststName string
	age       int
}
```

- **Another way to create Person**
``` Go
// creating a "Person" structure
type Person struct {
	FirstName string
	LastName  string
	age       int
}
func main() {
    //creating person in another way
	person1 := Person{
		FirstName: "Keanu",
		LastName:  "Reeves",
		age:       60,
	}
}
```

- We can also create using **new** keyword
``` Go
type Person struct {
	FirstName string
	LastName  string
	age       int
}

func main() {
    // we can also create using new keyword
	var person2 = new(Person)
	person2.FirstName = "Jhonny"
	person2.LastName = "Depp"
	person2.age = 62
	fmt.Println("Person person2:", person2)   // This will print like a pointer
}
```
- Assigning value in Struct, like creating object and assign value in Java.

``` Go
func main() {
	var tom Person
	fmt.Println("Person tom: ", tom)  // currentlu this is null 
	tom.FirstName = "Tom"
	tom.LastName = "Cruise"
	tom.age = 60
	fmt.Println("Person tom: ", tom)
}
```

- ***There can be structure inside structure***
``` Go
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

```