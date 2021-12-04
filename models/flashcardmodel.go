package models

import (
	"goWebApp/config"
	"goWebApp/entities"
)

type FlashModel struct {}

func (*FlashModel) FindAll() ([]entities.Flashcard, error) {
	db, err := config.GetDb()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from flashcards")
		if err2 != nil {
			return nil, err2
		} else {
			var flashcards []entities.Flashcard
			for rows.Next() {
				var flashcard entities.Flashcard
				rows.Scan(&flashcard.Id, &flashcard.Front, &flashcard.Back)
				flashcards = append(flashcards, flashcard)
			}
			rows.Close()
			return flashcards, nil
		}
	}
}

func (*FlashModel) Find(id int64) (entities.Flashcard, error) {
	db, err := config.GetDb()

	if err != nil {
		return entities.Flashcard{}, err
	} else {
		rows, err2 := db.Query("select * from flashcards where id=?", id)
		if err2 != nil {
			return entities.Flashcard{}, err2
		} else {
			var flashcard entities.Flashcard
			for rows.Next() {
				rows.Scan(&flashcard.Id, &flashcard.Front, &flashcard.Back)
			}
			rows.Close()
			return flashcard, nil
		}
	}
}

func (*FlashModel) Create(flashcard *entities.Flashcard) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into flashcards (front, back) values (?, ?)", flashcard.Front, flashcard.Back)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}