package pointers

func NumberAddOne(number *int) int {
	*number = *number + 1

	return *number

}
