package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	structs "github.com/gokurs/Projects/restaurant/repository/struct"
	"github.com/google/uuid"
	"time"
)

func (u PostgresRepo) Food1() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM food WHERE category=1`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	foodl := make([]structs.MenyuJson, 0)
	food1 := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&food1.Id, &food1.Name, &food1.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		foodl = append(foodl, food1)
	}
	return foodl, nil
}
func (u PostgresRepo) Food2() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM food WHERE category=2`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	foodl := make([]structs.MenyuJson, 0)
	food2 := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&food2.Id, &food2.Name, &food2.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		foodl = append(foodl, food2)
	}
	return foodl, nil
}

func (u PostgresRepo) Food3() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM food WHERE category=3`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	foodl := make([]structs.MenyuJson, 0)
	food3 := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&food3.Id, &food3.Name, &food3.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		foodl = append(foodl, food3)
	}
	return foodl, nil
}

func (u PostgresRepo) Salad() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM salad`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	saladl := make([]structs.MenyuJson, 0)
	for rows.Next() {
		salad := structs.MenyuJson{}
		if err = rows.Scan(&salad.Id, &salad.Name, &salad.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		saladl = append(saladl, salad)
	}
	return saladl, nil
}

func (u PostgresRepo) Drinks() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM drinks`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	drinkl := make([]structs.MenyuJson, 0)
	for rows.Next() {
		drink := structs.MenyuJson{}
		if err = rows.Scan(&drink.Id, &drink.Name, &drink.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		drinkl = append(drinkl, drink)
	}
	return drinkl, nil
}

func (u PostgresRepo) Shop(table_id, food_id, salad_id, drinck_id string) error {
	var tip uint8
	var id string
	if food_id != "" {
		tip = 1
		row1 := u.db.QueryRow("SELECT id FROM food WHERE id=$1", food_id)
		if err := row1.Scan(&food_id); errors.Is(err, sql.ErrNoRows) {
			return errors.New("There is no this type of food")
		} else if err != nil {
			print("222")
			return err
		}
		id = food_id
	} else if salad_id != "" {
		tip = 2
		row2 := u.db.QueryRow("SELECT id FROM salad WHERE id=$1", salad_id)
		if err := row2.Scan(&salad_id); errors.Is(err, sql.ErrNoRows) {
			return errors.New("There is no this type of salad")
		} else if err != nil {
			return err
		}
		id = salad_id
	} else if drinck_id != "" {
		tip = 3
		row3 := u.db.QueryRow("SELECT id FROM drinks WHERE id=$1", drinck_id)
		if err := row3.Scan(&drinck_id); errors.Is(err, sql.ErrNoRows) {
			return errors.New("There is no this type of drinks")
		} else if err != nil {
			return err
		}
		id = drinck_id
	}

	var chek_id string
	row := u.db.QueryRow("SELECT id FROM chek WHERE tables_id=$1 AND payment=false", table_id)

	if err := row.Scan(&chek_id); err != nil {
		return err
	}
	fmt.Println(chek_id)
	_, err := u.db.Exec("INSERT INTO basket VALUES ($1,$2,$3)", chek_id, id, tip)
	if err != nil {
		print("333")
		return err
	}
	return nil
}
func (u PostgresRepo) OpenChek(table_id string) error {
	row := u.db.QueryRow("SELECT busy FROM tables WHERE id=$1", table_id)
	var b bool
	if err := row.Scan(&b); err != nil {
		return err
	}
	if b == false {
		_, err := u.db.Exec("INSERT INTO chek (id,tables_id,exit_time) VALUES ($1,$2,$3)", uuid.NewString(), table_id, time.Now())
		if err != nil {
			return err
		}
		//tableni busy sini true qilish
		_, err = u.db.Exec("Update tables SET busy=true WHERE id=$1", table_id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Table is busy!")
}

func (u PostgresRepo) Chek(table_id string) (map[string]uint32, error) {
	mp := make(map[string]uint32)
	var id string
	row := u.db.QueryRow("SELECT id FROM chek WHERE tables_id=$1 AND payment=false", table_id)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	row2, err := u.db.Query("SELECT product_id,tip FROM basket WHERE chek_id=$1", id)
	if err != nil {
		return nil, err
	}
	list := map[string]uint8{}
	for row2.Next() {
		var (
			s   string
			tip uint8
		)
		if err = row2.Scan(&s, &tip); err != nil {
			return nil, err
		}
		list[s] = tip
	}
	var summa uint32
	for id, tip := range list {
		var (
			nomi  string
			narxi uint32
		)
		if tip == 1 {
			row := u.db.QueryRow("SELECT name,price FROM food WHERE id=$1", id)
			if err = row.Scan(&nomi, &narxi); err != nil {
				return nil, err
			}
		} else if tip == 2 {
			row2 := u.db.QueryRow("SELECT name,price FROM salad WHERE id=$1", id)
			if err = row2.Scan(&nomi, &narxi); err != nil {
				return nil, err
			}
		} else if tip == 3 {
			row3 := u.db.QueryRow("SELECT name,price FROM drinks WHERE id=$1", id)
			if err = row3.Scan(&nomi, &narxi); err != nil {
				return nil, err
			}
		}
		mp[nomi] = narxi
		summa += narxi

	}
	mp["Jami:"] = summa
	_ = u.db.QueryRow("Update tables SET busy=false WHERE id=$1", table_id)
	_ = u.db.QueryRow("Update chek SET payment=true WHERE tables_id=$1", table_id)
	return mp, nil
}
