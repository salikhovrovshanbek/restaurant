package postgres

import (
	"fmt"
	structs "github.com/gokurs/Projects/restaurant/repository/struct"
	"github.com/lib/pq"
	"math/rand"
	"time"
)

func (u PostgresRepo) Ingredient(id string) ([]structs.Product, error) {
	products_id := make([]string, 0)
	r := u.db.QueryRow(`SELECT products_id FROM food WHERE id=$1`, id)
	if err := r.Scan(pq.Array(&products_id)); err != nil {
		return nil, err
	}
	Products := make([]structs.Product, 0)
	fmt.Println(products_id)
	for _, v := range products_id {
		var pr structs.Product
		fmt.Println(v)
		r := u.db.QueryRow(`SELECT * FROM products WHERE id=$1`, v)
		if err := r.Scan(&pr.Id, &pr.Name, &pr.CreatedAt); err != nil {
			return nil, err
		}
		Products = append(Products, pr)
	}
	return Products, nil
}
func (u PostgresRepo) Set(sum int) ([][]structs.MenyuJson, error) {
	rand.Seed(time.Now().UnixNano())
	combos := make([][]structs.MenyuJson, 0)
	combo := make([]structs.MenyuJson, 0)
	rowf, err := u.db.Query(`SELECT id,name,price from food WHERE price<=$1`, sum)
	if err != nil {
		return nil, err
	}
	for rowf.Next() {
		menyus := make([]structs.MenyuJson, 0)
		menyud := make([]structs.MenyuJson, 0)
		mf := structs.MenyuJson{}
		if err := rowf.Scan(&mf.Id, &mf.Name, &mf.Price); err != nil {
			return nil, err
		}
		//if mf.Name == "Somsa" {
		//	mf.Name = "2xSomsa"
		//	mf.Price *= 2
		//}
		if uint32(sum) <= mf.Price {
			combo = make([]structs.MenyuJson, 0)
			combo = append(combo, mf)
			combos = append(combos, combo)
			continue
		}
		sumd := uint32(sum) - mf.Price
		rowd, err := u.db.Query(`SELECT id,name,price from drinks WHERE price<=$1`, sumd)
		if err != nil {
			return nil, err
		}
		for rowd.Next() {
			md := structs.MenyuJson{}
			if err := rowd.Scan(&md.Id, &md.Name, &md.Price); err != nil {
				return nil, err
			}
			menyud = append(menyud, md)
		}
		if len(menyud) == 0 {
			combo = make([]structs.MenyuJson, 0)
			combo = append(combo, mf)
			combos = append(combos, combo)
			continue
		}
		md := menyud[rand.Intn(len(menyud))]
		if sumd <= md.Price {
			combo = make([]structs.MenyuJson, 0)
			combo = append(combo, mf, md)
			combos = append(combos, combo)
			continue
		}
		sums := sumd - md.Price
		rows, err := u.db.Query(`SELECT id,name,price from salad WHERE price<=$1`, sums)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			ms := structs.MenyuJson{}
			if err := rows.Scan(&ms.Id, &ms.Name, &ms.Price); err != nil {
				return nil, err
			}
			menyus = append(menyus, ms)
		}
		if len(menyus) == 0 {
			combo = make([]structs.MenyuJson, 0)
			combo = append(combo, mf, md)
			combos = append(combos, combo)
			continue
		}
		ms := menyus[rand.Intn(len(menyus))]
		combo = make([]structs.MenyuJson, 0)
		combo = append(combo, mf, md, ms)
		combos = append(combos, combo)
		println(sums)
	}
	return combos, nil
}
func (u PostgresRepo) ShopCombo(table_id, food_id, salad_id, drinck_id string) error {
	if food_id != "" {
		if err := u.Shop(table_id, food_id, "", ""); err != nil {
			return err
		}
	}
	if salad_id != "" {
		if err := u.Shop(table_id, "", salad_id, ""); err != nil {
			return err
		}
	}
	if drinck_id != "" {
		if err := u.Shop(table_id, "", "", drinck_id); err != nil {
			return err
		}
	}
	return nil
}
func (u PostgresRepo) DeleteBasket(table_id, food_id, salad_id, drinck_id string) error {
	if food_id != "" {
		_, err := u.db.Exec(
			`DELETE FROM basket WHERE chek_id=(select id FROM chek WHERE tables_id=$1 and payment=false) and product_id=$2`,
			table_id, food_id)
		if err != nil {
			return err
		}
	}
	if salad_id != "" {
		_, err := u.db.Exec(
			`DELETE FROM basket WHERE chek_id=(select id FROM chek WHERE tables_id=$1 and payment=false) and product_id=$2`,
			table_id, salad_id)
		if err != nil {
			return err
		}
	}
	if drinck_id != "" {
		_, err := u.db.Exec(
			`DELETE FROM basket WHERE chek_id=(select id FROM chek WHERE tables_id=$1 and payment=false) and product_id=$2`,
			table_id, drinck_id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (u PostgresRepo) TableList() ([]structs.Table, error) {
	row, err := u.db.Query(`SELECT * FROM tables`)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	tables := make([]structs.Table, 0)
	for row.Next() {
		var t structs.Table
		if err := row.Scan(&t.Id, &t.Number, &t.Buys); err != nil {
			return nil, err
		}
		tables = append(tables, t)
	}
	return tables, nil
}
