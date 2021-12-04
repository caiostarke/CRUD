package main

import (
	"goWebApp/controllers/jokingscontroller"
	"goWebApp/controllers/multicontroller"
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

	// Multi Languages Cards

	mux.HandleFunc("/multicards", multicontroller.Index)
	mux.HandleFunc("/multicards/multicard", multicontroller.View)

	http.ListenAndServe(":8080", mux)
}
