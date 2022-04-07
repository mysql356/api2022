package structure

type Address struct {
	City  string
	State string
}

type Contact struct {
	Email  string
	Mobile int
}

//Nested
type Employee struct {
	Name string
	Age  int
	Add  Address
	Cont []Contact `json:"contact"`
}

//Promoted
type Person struct {
	Name string
	Age  int
	Address
	Cont []Contact `json:"contact"`
}
