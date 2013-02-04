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
	DB_HOST  = "mongodb://timeChamber:123456@linus.mongohq.com:10086/physicsolympiad"
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
}

func olympiadHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "ParseForm:", err)
		return
	}

	collection := req.FormValue("collection")
	document, err := strconv.ParseInt(req.FormValue("document"), 10, 0)

	if err != nil {
		fmt.Fprintln(w, "ParseInt:", err)
		return
	}

	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		fmt.Fprintln(w, "Dial:", err)
		return
	}

	defer session.Close()

	c := session.DB(DATABASE).C(collection)

	var olympiads []Olympiad
	err = c.Find(nil).Limit(100).All(&olympiads)
	//err = c.Find(bson.M{"year":1998}).Limit(100).All(&olympiads)

	if err != nil {
		fmt.Fprintln(w, "Find olympiad:", err)
		return
	}

	var this Olympiad

	for _, olympiad := range olympiads {
		if olympiad.Year == document {
			this = olympiad
		}
	}

	var categories []Category
	c = session.DB(DATABASE).C("categories")
	err = c.Find(nil).All(&categories)

	if err != nil {
		fmt.Fprintln(w, "Find categories:", err)
		return
	}
	
	data := dict{
		"Collection": collection, 
		"Olympiads":  olympiads, 
		"Document":   this, 
		"Categories": categories,
	}
	
	tmpl := template.Must(template.ParseFiles("templates/olympiad.html", "templates/categories.html"))
	err = tmpl.Execute(w, data)

	if err != nil {
		fmt.Fprintln(w, "Execute:", err)
		return
	}
}
