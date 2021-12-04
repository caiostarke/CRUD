package jokingscontroller

import (
	"goWebApp/entities"
	"goWebApp/models"
	"html/template"
	"net/http"
	"strconv"
)

func RealIndex(w http.ResponseWriter, r *http.Request ) {
	tmp, _ := template.ParseFiles("views/jokings/real_index.html")
	tmp.Execute(w, nil)
}

func FlashJokings(w http.ResponseWriter, r *http.Request ) {
	var FlashModels models.FlashModel
	flashcards, _ := FlashModels.FindAll() 

	data := map[string]interface{}{
		"flashcard": flashcards,
	}

	tmp, _ := template.ParseFiles("views/jokings/flashcards.html")
	tmp.Execute(w, data)
}

func ContentJokings(w http.ResponseWriter, r *http.Request ) {
	tmp, _ := template.ParseFiles("views/jokings/content.html")
	tmp.Execute(w, nil)
}


func View(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, err := strconv.ParseInt(query.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	var FlashModel models.FlashModel
	flashcard, _ := FlashModel.Find(id)

	data := map[string]interface{}{
		"flashcard":flashcard,
	}
	if err != nil {
		panic(err)
	}

	tmp, _ := template.ParseFiles("views/jokings/flashcard.html")
	tmp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/jokings/flashadd.html")
	tmp.Execute(w, nil)
}

func ProcessAdd(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var flashcard entities.Flashcard

	flashcard.Front = r.FormValue("front")
	flashcard.Back = r.Form.Get("back")

	var FlashModel models.FlashModel
	FlashModel.Create(&flashcard)
	http.Redirect(w, r, "/jokings/flashcards", http.StatusSeeOther)
}