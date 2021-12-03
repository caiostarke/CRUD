package main

import (
	"goWebApp/controllers/jokingscontroller"
	"goWebApp/controllers/productcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", jokingscontroller.RealIndex)
	http.HandleFunc("/jokings/flashcards", jokingscontroller.FlashJokings)
	http.HandleFunc("/jokings/contents", jokingscontroller.ContentJokings)

	http.HandleFunc("/list", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/processadd", productcontroller.ProcessAdd)
	http.HandleFunc("/product/delete", productcontroller.Delete)
	http.HandleFunc("/product/edit", productcontroller.Edit)
	http.HandleFunc("/product/update", productcontroller.Update)

	http.HandleFunc("/product/view", jokingscontroller.View)
	
	http.ListenAndServe(":8080", nil)
}
