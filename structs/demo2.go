package structs

type customer struct {
	firstName string
	lastName  string
	Age       int
}

func AddCustomer(firstName string, lastName string, age int) customer {
	result := customer{firstName: firstName, lastName: lastName, Age: age}

	if result.firstName == "" || result.Age == 0 || result.lastName == "" {
		return customer{}
	}
	return result
}

func (c *customer) Ageup() {
	c.Age++
}
