package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Database struct {
	Pools []Pool `json:"pools"`
}

type Pool struct {
	Uid      string   `json:"uid"`
	Language string   `json:"language"`
	PoolName string   `json"pool_name"`
	Currency string   `json:"currency"`
	Amount   int      `json:"amount"`
	Members  []Person `json:"members"`
}

type Person struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (p *Person) toString() string {
	return fmt.Sprintf("%s|%s|%s|", p.FirstName, p.LastName, p.Email)
}

func TestJson() []byte {
	person := &Person{
		Email:     "normananton03@gmail.com",
		FirstName: "Anton",
		LastName:  "Norman",
	}
	pool := &Pool{
		Uid:      "123asd456qwe",
		Currency: "kr",
		Amount:   200,
		Language: "SEK",
		Members:  []Person{*person},
	}
	db := &Database{
		Pools: []Pool{*pool},
	}
	dbString, _ := json.Marshal(db)
	return dbString
}

func LoadJSON() *Database {
	var db Database

	file, err := os.Open("./db.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	fileBytes, _ := io.ReadAll(file)
	json.Unmarshal(fileBytes, &db)
	return &db
}

func CreatePool(UUID, currency, language, poolName string, amount int) error {
	pool := &Pool{
		Uid:      UUID,
		Currency: currency,
		Language: language,
		Amount:   amount,
		PoolName: poolName,
		Members:  []Person{},
	}
	db := LoadJSON()
	db.Pools = append(db.Pools, *pool)

	jsonString, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile("./db.json", jsonString, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func DeletePool(UUID string) error {
	db := LoadJSON()
	var index int
	for i, pool := range db.Pools {
		if pool.Uid == UUID {
			index = i
			break
		}
	}

	if index == 0 {
		db.Pools = db.Pools[1:]
	} else if index == len(db.Pools)-1 {
		db.Pools = db.Pools[:len(db.Pools)-1]
	} else {
		db.Pools = append(db.Pools[:index], db.Pools[index+1:]...)
	}
	jsonString, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile("./db.json", jsonString, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func GetPool(UUID string) ([]byte, error) {
	db := LoadJSON().Pools
	for _, pool := range db {
		if pool.Uid == UUID {
			jsonPool, err := json.Marshal(pool)
			if err != nil {
				return []byte(""), fmt.Errorf("Error")
			}
			return jsonPool, nil
		}
	}
	return []byte(""), fmt.Errorf("No such pool")
}

func InsertMember(NewMember *addMemberRequest) (*Person, error) {
	db := LoadJSON()
	for i, pool := range db.Pools {
		if pool.Uid == NewMember.Uid {
			db.Pools[i].Members = append(db.Pools[i].Members, NewMember.Person)
			jsonString, err := json.MarshalIndent(db, "", "\t")
			if err != nil {
				return nil, fmt.Errorf("Something went wrong")
			}
			if err := os.WriteFile("./db.json", jsonString, os.ModePerm); err != nil {
				return nil, err
			}
			return &NewMember.Person, nil
		}
	}
	return nil, fmt.Errorf("Pool: %s does not exist", NewMember.Uid)
}
