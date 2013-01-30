package main

import (
	"fmt"
//	"encoding/csv"
	"log"
	"net/http"
//	"os"
	"labix.org/v2/mgo"
	"strconv"
//    "labix.org/v2/mgo/bson"	
)

const DB_LOCATION = "localhost:27017"

type Olympiad struct {
	Year int64  `year`
	City string  `city`
	Country string `country`
}

func olympiadHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Print("olympiadHandler", err.Error())
		fmt.Fprintln(w, "Error:", err.Error())
		return
	}

	collection := req.FormValue("collection")
	document, _ := strconv.ParseInt(req.FormValue("document"), 10, 0)

	fmt.Fprintln(w, "collection:", collection, "& document:", document)

	session, err := mgo.Dial(DB_LOCATION)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
        return
	}

	defer session.Close()
	
	c := session.DB("test").C(collection)
	
	n, _ := c.Count()
	fmt.Fprintln(w, n, "documents in collection")
	
	var olympiads []Olympiad
	err = c.Find(nil).Limit(100).All(&olympiads)
	//err = c.Find(bson.M{"year":1998}).Limit(100).All(result)
	
	if err != nil {
		fmt.Fprintln(w, "error", err)
		return
	}
	
	fmt.Fprintln(w, olympiads)
	
	var this Olympiad
	
	for _, v := range olympiads {
		if v.Year == document {
			this = v
		}
	}

	fmt.Fprintln(w, document, ":", this)

//	target := "data/olympiads/" + collection + ".csv"
//	
//	file, err := os.Open(target)
//    if err != nil {
//        fmt.Fprintln(w, "Error:", err)
//        return
//    }
//    defer file.Close()
//    
//    reader := csv.NewReader(file)
//    records, err := reader.ReadAll()
//    olympiads := make ([]IntOlympiad, len(records))
//    
//    var this Olympiad 
//    
//    for i, record := range records {
//        _this := Olympiad {
//        	Year: record[0],
//        	City: record[1],
//        	Country: record[2],
//        }
//        
//        olympiads[i] = _this
//        
//        if _this.Year == document {
//        	this = _this
//        }
//    }
//    
//    fmt.Fprintln(w, olympiads)
//    fmt.Fprintln(w, this)
}

