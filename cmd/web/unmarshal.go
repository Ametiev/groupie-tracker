package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	artists_url   = "https://groupietrackers.herokuapp.com/api/artists"
	relations_url = "https://groupietrackers.herokuapp.com/api/relation"
)

func Get_artists() ([]Artist, error) {
	res, err := http.Get(artists_url)
	if err != nil {
		log.Println(err)
		return Artists, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return Artists, err
	}
	err = json.Unmarshal(body, &Artists)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return Artists, nil
}

func UnmarshallRelations() error {
	res, err := http.Get(relations_url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	unerr := json.Unmarshal(body, &Relation)
	if unerr != nil {
		log.Println("Unmar Rel")
		return unerr
	}
	for i := range Artists {
		Artists[i].DatesLocation = Relation.Index[i].DatesLocation
		FormatDates(i)
		FindSolo(i)
	}
	return nil
}
func FormatDates(i int) {
	res := make(map[string][]string)
	for key := range Artists[i].DatesLocation {
		res[strings.Title(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(key, "_", " "), "-", ", "), " usa", " USA"), " uk", " UK"))] = Artists[i].DatesLocation[key]
		delete(Artists[i].DatesLocation, key)
	}
	Artists[i].DatesLocation = res
}
func FindSolo(i int) {
	if len(Artists[i].Members) == 1 {
		Artists[i].Solo = true
	}
}
