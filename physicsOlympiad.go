package main

import (
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	"net/http"
	"strconv"

	//"labix.org/v2/mgo/bson"	
)

const (
	DB_LOCAL = "localhost"
	DB_HOST  = "mongodb://timeChamber:123456@linus.mongohq.com:10029/physicsolympiad"
	DATABASE = "physicsolympiad"
)

type dict map[string]interface{}

type Olympiad struct {
	Year    int64  `year`
	City    string `city`
	Country string `country`
}

type Category struct {
	Name       string `name`
	Collection string `collection`
	ShortName  string `shortName`
	ExamPrefix string `examPrefix`
}

func olympiadHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "ParseForm:", err)
		return
	}

	collection := req.FormValue("collection")

	if collection == "" {
		collection = "iphos"
	}

	document := req.FormValue("document")
	var year int64
	if document == "" {
		year = 0
	} else {
		year, err = strconv.ParseInt(document, 10, 0)

		if err != nil {
			year = 0
		}
	}

	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		fmt.Fprintln(w, "Dial:", err)
		return
	}

	defer session.Close()

	c := session.DB(DATABASE).C(collection)

	var olympiads []Olympiad
	err = c.Find(nil).Sort("-$natural").Limit(100).All(&olympiads)
	//err = c.Find(bson.M{"year":1998}).Limit(100).All(&olympiads)

	if err != nil {
		fmt.Fprintln(w, "Find olympiad:", err)
		return
	}

	var this Olympiad

	for _, olympiad := range olympiads {
		if olympiad.Year == year {
			this = olympiad
		}
	}

	if year == 0 {
		this = olympiads[0]
	}

	var categories []Category
	c = session.DB(DATABASE).C("categories")
	err = c.Find(nil).All(&categories)

	if err != nil {
		fmt.Fprintln(w, "Find categories:", err)
		return
	}

	var prefix string

	for _, v := range categories {
		if v.Collection == collection {
			prefix = v.ExamPrefix
		}
	}

	data := dict{
		"Collection": collection,
		"Olympiads":  olympiads,
		"Document":   this,
		"Categories": categories,
		"ExamPrefix": prefix,
	}

	tmpl := template.Must(template.ParseFiles("templates/olympiad.html", "templates/categories.html"))
	err = tmpl.Execute(w, data)

	if err != nil {
		fmt.Fprintln(w, "Execute:", err)
		return
	}
}
