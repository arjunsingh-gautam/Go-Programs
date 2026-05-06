package utils

func Add(a int, b int) int {
	return a + b
}

// unexported function because it starts with small case letter
func subtract(a int, b int) int {
	return a - b
}