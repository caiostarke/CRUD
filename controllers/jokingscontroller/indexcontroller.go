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


func Edit(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, err := strconv.ParseInt(query.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	var FlashModel models.FlashModel
	flashcard, _ := FlashModel.Find(id)

	data := map[string]interface{} {
		"flashcard": flashcard,
	}

	tmp, _ := template.ParseFiles("views/jokings/flashedit.html")
	tmp.Execute(w, data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var flashcard entities.Flashcard
	flashcard.Id, _ = strconv.ParseInt(r.Form.Get("id"), 10, 64)
	flashcard.Front = r.Form.Get("front")
	flashcard.Back = r.Form.Get("back")

	var FlashModel models.FlashModel
	FlashModel.Update(flashcard)
	http.Redirect(w, r, "/jokings/flashcards", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var FlashModel models.FlashModel
	FlashModel.Delete(id)
	http.Redirect(w, r, "/jokings/flashcards", http.StatusSeeOther)
}