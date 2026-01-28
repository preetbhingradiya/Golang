package intermediate

import "fmt"

type Address struct {
	city   string
	state  string
}

type PhoneCell struct {
	countryCode string
	number      string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address  //embedded struct
	PhoneCell //anonymous  field struct
}

type MyInt int

func Struct() {
	/**Structs are custom data types that group together related fields. They are used to create complex data structures.**/

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		// age:       30,
		address: Address{
			city:  "New York",
			state: "NY",
		},
		PhoneCell: PhoneCell{
			countryCode: "+1",
			number:      "123-456-7890",
		},
	}
	fmt.Println("Person 1:", p1)
	fmt.Println("Age :", p1.age) //if the field is not initialized default value is assigned of that type
	fmt.Println("City :", p1.address.city)
	fmt.Println("Phone Number :", p1.number) //accessing anonymous field directly

	//You can also create struct instances without field names
	p2 := Person{"Jane", "Smith", 25, Address{"Los Angeles", "CA"}, PhoneCell{"+1", "987-654-3210"}}
	fmt.Println("Person 2:", p2)

	//Anonymous Structs
	employee := struct {
		id     int
		Name   string
		salary float64
	}{
		id:     101,
		Name:   "Alice",
		salary: 75000.50,
	}

	fmt.Println("Employee:", employee)

	//Using Methods with Structs
	p1.changePersoneName("Kaliya")
	fmt.Println("Full Name of Person 1:", p1.fullName())

	//method instance with any type
	num := MyInt(42)
	fmt.Println("MyInt value:", num.double())
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) changePersoneName(firstName string) {
	p.firstName = firstName
}

func (m MyInt) double() MyInt {
	return m * 2
}


