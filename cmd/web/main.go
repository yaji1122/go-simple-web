package main

import (
	"fmt"
	"github.com/yaji1122/go-simple-web/pkg/handler"
	"net/http"
)


const port = ":8080"

func main() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("Starting application on port 8080 http://localhost:8080"))
	_ = http.ListenAndServe(port, nil)
}




//func Divide(w http.ResponseWriter, r *http.Request) {
//	f, err := divideValues(100.0, 0.0)
//	if err != nil {
//		fmt.Fprintf(w, fmt.Sprintf("Error Message: %s", err))
//	} else {
//		fmt.Fprintf(w, fmt.Sprintf( "%f divided by %f is %f", 100.0, 0.0, f))
//	}
//}

//func divideValues(x, y float32) (float32, error) {
//	var result float32
//	if y <= 0.0 {
//		return result, errors.New("Can't not divide by 0")
//	}
//	result = x / y
//	return result, nil
//}
////add two values
//func addValues(x, y int) (int, error) {
//	var sum int
//	sum = x + y
//	return sum, nil
//}



