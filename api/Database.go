package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Database struct{
    Pools []Pool `json:"pools"`
}

type Pool struct{
    Uid     string   `json:"uid"`
    Members []Person `json:"members"`
}

type Person struct{
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

func TestJson() []byte {
    person := &Person{
        Email: "normananton03@gmail.com",
        FirstName: "Anton",
        LastName: "Norman",
    }
    pool := &Pool{
        Uid: "123asd456qwe",
        Members: []Person{*person},
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

func CreatePool(UUID string) error {
    pool := &Pool{
        Uid: UUID,
        Members: []Person{},
    }
    db := LoadJSON()
    db.Pools = append(db.Pools, *pool)

    jsonString, err := json.Marshal(db) 
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
        if pool.Uid == UUID{
            jsonPool, err := json.Marshal(pool)
            if err != nil {
                return []byte(""), fmt.Errorf("Error")
            }
            return jsonPool, nil
        }
    }
    return []byte(""), fmt.Errorf("No such pool")
}
