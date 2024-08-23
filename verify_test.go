package verifystruct

// Define a struct to test verification
type testModel struct {
	Name  string `json:"name" verify:"required=true,type=string,min=4,max=20"`
	Age   int    `json:"age" verify:"required=true,type=int,min=4,max=20"`
	Email string `json:"email" verify:"type=email,min=4,max=20"`
}
