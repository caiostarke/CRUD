package main

import (
	"goWebApp/controllers/jokingscontroller"
	"goWebApp/controllers/productcontroller"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/views/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", jokingscontroller.RealIndex)

	mux.HandleFunc("/jokings/flashcards", jokingscontroller.FlashJokings)
	mux.HandleFunc("/jokings/flashcard/add", jokingscontroller.Add)
	mux.HandleFunc("/jokings/flashcard/processadd", jokingscontroller.ProcessAdd)

	mux.HandleFunc("/jokings/flashcard", jokingscontroller.View)
	mux.HandleFunc("/jokings/flashcard/edit", jokingscontroller.Edit)
	mux.HandleFunc("/jokings/flashcard/update", jokingscontroller.Update)
	mux.HandleFunc("/jokings/flashcard/delete", jokingscontroller.Delete)


	mux.HandleFunc("/list", productcontroller.Index)
	mux.HandleFunc("/product", productcontroller.Index)
	mux.HandleFunc("/product/index", productcontroller.Index)
	mux.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/processadd", productcontroller.ProcessAdd)
	mux.HandleFunc("/product/delete", productcontroller.Delete)
	mux.HandleFunc("/product/edit", productcontroller.Edit)
	mux.HandleFunc("/product/update", productcontroller.Update)


	http.ListenAndServe(":8080", mux)
}
