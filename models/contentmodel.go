package models

import (
	"goWebApp/config"
	"goWebApp/entities"
)

type ContentModel struct{}

func (*ContentModel) FindAll() ([]entities.Content, error) {
	db, err := config.GetDb()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from contente")
		if err2 != nil {
			return nil, err2
		} else {
			var contents []entities.Content
			for rows.Next() {
				var content entities.Content
				rows.Scan(&content.Id, &content.Title, &content.Description, &content.Content)
				contents = append(contents, content)
			}
			rows.Close()
			return contents, nil
		}
	}
}
