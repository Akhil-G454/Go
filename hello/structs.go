package main

type address struct {
	name    string
	street  string
	city    string
	state   string
	pincode int
}

// func structs() {
// 	// first way of declaration
// 	var a address
// 	// you can access the individual fields of the struct using (.) operator
// 	a.name = "Akhil"
// 	a.street = "Padmasri Hills"
// 	a.city = "Hyderabad"
// 	a.state = "Telangana"
// 	a.pincode = 500086

// 	// second way of declaration
// 	var b = address{"Akhil", "Padmasri Hills", "Hyderabad", "Telangana", 500086}

// 	// if you want to initialize only some of the variables in the struct then you can use the
// 	// key: value type declaration

// 	var c = address{name: "Akhil", street: "Padmasri Hills", pincode: 500086}

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }
