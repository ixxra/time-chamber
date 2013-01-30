package main

import (
	"fmt"
	"net/http"
	"labix.org/v2/mgo"
	"strconv"
	"html/template"
//   "labix.org/v2/mgo/bson"	
)

const DB_LOCATION = "localhost:27017"
const DATABASE = "physicsolympiad"

type dict map[string] interface{}

type Olympiad struct {
	Year int64  `year`
	City string  `city`
	Country string `country`
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

	session, err := mgo.Dial(DB_LOCATION)
	if err != nil {
		fmt.Fprintln(w, "Dial:", err)
        return
	}

	defer session.Close()
	
	c := session.DB(DATABASE).C(collection)
	
	var olympiads []Olympiad
	err = c.Find(nil).Limit(100).All(&olympiads)
	//err = c.Find(bson.M{"year":1998}).Limit(100).All(result)
	
	if err != nil {
		fmt.Fprintln(w, "Find:", err)
		return
	}
	
	var this Olympiad
	
	for _, olympiad := range olympiads {
		if olympiad.Year == document {
			this = olympiad
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/olympiad.html", "templates/navbar.html"))
	err = tmpl.Execute(w, dict {"Collection": collection ,"Olympiads": olympiads, "Document": this})
	
	if err != nil {
		fmt.Fprintln(w, "Execute:", err)
		return
	}
}

