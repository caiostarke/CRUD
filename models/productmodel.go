package models

import (
	"goWebApp/config"
	"goWebApp/entities"
)

type ProductModel struct {
}

func (p *ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDb()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			rows.Close()
			return products, nil
		}
	}
}

func (p *ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDb()

	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id=?", id)
		if err2 != nil {
			return entities.Product{}, err2
		} else {
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
			}
			return product, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into product (name, price, quantity, description) values (?, ?, ?, ?)", product.Name, product.Price, product.Quantity, product.Description)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ProductModel) Update(product entities.Product) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update product set name = ?, price = ?, quantity = ?, description = ? where id = ?",
		product.Name, product.Price, product.Quantity, product.Description, product.Id)
		
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ProductModel) Delete(id int64) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from product where id=?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
