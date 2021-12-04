package models

import (
	"goWebApp/config"
	"goWebApp/entities"
)

type MultiModel struct{}

func (*MultiModel) FindAll() ([]entities.Multicard, error) {
	db, err := config.GetDb()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from multicard")
		if err2 != nil {
			return nil, err2
		} else {
			var multicards []entities.Multicard
			for rows.Next() {
				var multicard entities.Multicard
				rows.Scan(&multicard.Id, &multicard.English, &multicard.Italian, &multicard.Spanish, &multicard.Portuguese)
				multicards = append(multicards, multicard)
			}
			rows.Close()
			return multicards, nil
		}

	}
}

func (*MultiModel) Find(id int64) (entities.Multicard, error) {
	db, err := config.GetDb()

	if err != nil {
		return entities.Multicard{}, err
	} else {
		rows, err2 := db.Query("select * from multicard where id=?", id)
		if err2 != nil {
			return entities.Multicard{}, err2
		} else {
			var multicard entities.Multicard
			for rows.Next() {
				rows.Scan(&multicard.Id, &multicard.English, &multicard.Italian, &multicard.Spanish, &multicard.Portuguese)
			}
			rows.Close()
			return multicard, nil
		}
	}
}

func (*MultiModel) Create(multicard *entities.Multicard) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into multicard (english, italian, spanish, portuguese) values (?, ?, ?, ?)", multicard.English, multicard.Italian, multicard.Spanish, multicard.Portuguese)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*MultiModel) Update(multicard entities.Multicard) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update multicard set english=?, italian=?, spanish=?, portuguese=? where id=?", multicard.English, multicard.Italian, multicard.Spanish, multicard.Portuguese, multicard.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*MultiModel) Delete(id int64) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from multicard where id=?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}