package model

import (
	"firebase.google.com/go/db"
	"fmt"
	"log"
	"strings"
)

type Antrian struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}

func AddAntrian() (bool, error) {
	_, _, dataAntrian := GetAntrian()
	var Id string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		Id = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}
	antrian := Antrian{
		Id:     Id,
		Status: false,
	}
	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func GetAntrian() (bool, error, []map[string]interface{}) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
		return false, err, nil
	}

	return true, nil, data
}

func UpdateAntrian(idAntrian string) (bool, error) {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian{
		Id:     idAntrian,
		Status: true,
	}
	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func DeleteAntrian(idAntrian string) (bool, error) {

	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}