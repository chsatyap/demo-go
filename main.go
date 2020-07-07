package main

import "fmt"

const (
	productPath string = "/tmp/products.json"
	sellerPath  string = "/tmp/sellers.json"
)

func main() {
	s := Seller{
		Name: "SomeSeller",
		Address: Address{
			City: "Gotham",
		},
	}
	var i GenericInterface
	i = &s

	givenCity := "Gotham"

	switch i.(type) {
	case interface{}:
		fmt.Println("What to do")
	case Product:
		fmt.Println(i.(Product).DeliversTo(givenCity))
	case Seller:
		fmt.Println(i.(Seller).DeliversTo(givenCity))
	}
	
	var arr []int
	arr = nil
	
	if arr != nil {
	    for _, val := range arr {
		    print(val)
	    }
	}
	
	if arr == nil || len(arr) == 0 {
        	println("err! arr is not populated.")
	}
}
