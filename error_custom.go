package main

import "fmt"

// custom error
type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

// custom error
type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

// example function
func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "felix" {
		return &notFoundError{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("who", nil)

	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation eror :", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not found error :", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error :", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation eror :", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found eror :", finalError.Error())
		default:
			fmt.Println("unknown error :", err.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
