package main

import (
	"fmt"
	"github.com/yaji1122/go-simple-web/pkg/config"
	"github.com/yaji1122/go-simple-web/pkg/handler"
	"github.com/yaji1122/go-simple-web/pkg/render"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	//使用同一個 AppConfig
	var appConfig config.AppConfig
	//產生 Template Cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error Creating Template Cache")
	}
	//將產生的Template Cache指定到 AppConfig中
	appConfig.TemplateCache = tc
	//傳入 AppConfig
	appConfig.UseCache = false //dev mode 設為False
	// render pkg 設定 appConfig
	render.NewTemplates(&appConfig)
	//set repo
	repo := handler.NewRepo(&appConfig)
	handler.NewHandlers(repo)

	//here move the routes to the router.go
	//http.HandleFunc("/", handler.Repo.Home)
	//http.HandleFunc("/about", handler.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port 8080 http://localhost:8080"))
	//_ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
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
