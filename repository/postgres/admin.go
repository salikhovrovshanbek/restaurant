package postgres

import (
	structs "github.com/gokurs/Projects/restaurant/repository/struct"
	"time"
)

func (u PostgresRepo) CountUsers() (map[int]int, error) {
	SumUsers := map[int]int{}
	rows, err := u.db.Query("SELECT tables_id FROM chek WHERE exit_time>$1 and payment=true", time.Now().Add(-time.Hour*24))
	if err != nil {
		return map[int]int{}, err
	}
	countTables := map[string]int{} //aynan bitta tabledagi odamlar soni uchun
	for rows.Next() {
		var table_id string
		if err := rows.Scan(&table_id); err != nil {
			return nil, err
		}

		_, ok := countTables[table_id]
		if !ok {
			countTables[table_id] = 1
		} else {
			countTables[table_id] += 1 //bitta tabledagi odamlar soni uchun
		}
	}
	s := 0
	for id, val := range countTables {
		row := u.db.QueryRow("SELECT number FROM tables WHERE id=$1", id)
		var number int /// table raqami
		row.Scan(&number)
		SumUsers[number] = val // bitta tabledagi odamlar soni
		s += val               // ummiy bir kunlik odamlar soni
	}
	SumUsers[0] = s // nolinchi indeksda hardoim ummiy odamlar soni turadi
	return SumUsers, nil
}

func (u PostgresRepo) CountSum() (map[int]int, error) {
	CountSum := map[int]int{} // ummiy bir kunlik budjet
	rows, err := u.db.Query("SELECT id, tables_id FROM chek WHERE exit_time>$1 and payment=true", time.Now().Add(-time.Hour*24))
	if err != nil {
		return nil, err
	}
	countChek := map[string]string{}
	for rows.Next() {
		var chek_id, table_id string
		if err := rows.Scan(&chek_id, &table_id); err != nil {
			return nil, err
		}
		countChek[chek_id] = table_id
	}
	sum := 0
	for chek, table := range countChek {
		sum_table := 0
		row, err := u.db.Query("SELECT product_id, tip FROM basket WHERE chek_id=$1", chek)
		if err != nil {
			return nil, err
		}

		for row.Next() {
			var product_id string
			var tip, price int
			if err = row.Scan(&product_id, &tip); err != nil {
				return nil, err
			}
			if tip == 1 {
				row := u.db.QueryRow("SELECT price FROM food WHERE id=$1", product_id)
				row.Scan(&price)
			} else if tip == 2 {
				row := u.db.QueryRow("SELECT price FROM salad WHERE id=$1", product_id)
				row.Scan(&price)
			} else if tip == 3 {
				row := u.db.QueryRow("SELECT price FROM drinks WHERE id=$1", product_id)
				row.Scan(&price)
			}
			sum_table += price
		}
		rt := u.db.QueryRow("SELECT number FROM tables WHERE id=$1", table)
		var num int
		err = rt.Scan(&num)
		CountSum[num] = sum_table
		sum += sum_table
	}
	CountSum[0] = sum
	return CountSum, nil
}

func (u PostgresRepo) ProductList() ([]structs.Product, error) {
	rows, err := u.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	products := []structs.Product{}
	for rows.Next() {
		product := structs.Product{}
		if err := rows.Scan(&product.Id, &product.Name, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (u PostgresRepo) UpdateProduct(id string) error {
	if _, err := u.db.Exec("UPDATE products SET created_at=$1 WHERE id=$2", time.Now(), id); err != nil {
		return err
	}
	return nil
}
