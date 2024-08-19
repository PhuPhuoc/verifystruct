package verifystruct

import "testing"

func TestVerify(t *testing.T) {
    s := MyStruct{Name: "", Age: 25}
    err := Verify(s)
    if err == nil {
        t.Errorf("Expected an error, got nil")
    }

    s = MyStruct{Name: "John", Age: 25}
    err = Verify(s)
    if err != nil {
        t.Errorf("Did not expect an error, got %v", err)
    }
}
