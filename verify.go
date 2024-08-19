package verifystruct

import "fmt"

// Example struct for demonstration
type MyStruct struct {
	Name string
	Age  int
}

// Function to verify struct
func Verify(s MyStruct) error {
	if s.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if s.Age <= 0 {
		return fmt.Errorf("age must be positive")
	}
	return nil
}
