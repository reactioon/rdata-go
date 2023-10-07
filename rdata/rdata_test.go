package rdata

import (

	"fmt"
	"testing"

	"github.com/google/uuid"

)

var _ = fmt.Println

var wTestHost = "192.168.1.202"
var wTestPort = "60162"
var wTestCollection = "test"
var wTestBook = "users"
var wTestBookRecordsListLimit = "50"
var wTestBookRecordsMetaTags = "1"
var wTestKeyID = "4ec735ef-5574-445f-9f83-70aeccbe4d8e"
var wTestContentJSONdoc = `{"teste":"teste-123"}`

func BenchmarkListMultiples(b *testing.B) {

    for i := 0; i < b.N; i++ {
		
		client := CLIENT{}
		loader := client.Load(wTestHost, wTestPort)
		conn,_ := loader.Connect()
		conn.Send(`route=collection.books.documents.list&collection=`+wTestCollection+`&book=`+wTestBook+`&limit=`+wTestBookRecordsListLimit+`&meta=` + wTestBookRecordsMetaTags)
		conn.Close()

    }

}

func BenchmarkGetMultiples(b *testing.B) {

    for i := 0; i < b.N; i++ {
		
		client := CLIENT{}
		loader := client.Load(wTestHost, wTestPort)
		conn,_ := loader.Connect()
		conn.Send(`route=collection.books.documents.get&collection=`+wTestCollection+`&book=`+wTestBook+`&key=`+wTestKeyID+`&meta=` + wTestBookRecordsMetaTags)
		conn.Close()

    }

}

func BenchmarkInsertMultiples(b *testing.B) {
	
    for i := 0; i < b.N; i++ {
		
		client := CLIENT{}
		loader := client.Load(wTestHost, wTestPort)
		conn,_ := loader.Connect()
		wKeyRecord := uuid.New().String()
		conn.Send(`route=collection.books.documents.insert&collection=`+wTestCollection+`&book=`+wTestBook+`&key=`+wKeyRecord+`&value=` + wTestContentJSONdoc)
		conn.Close()

    }

}