package rdata

import (
	
	"testing"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"

)

var wTestHost = "192.168.1.202"
var wTestPort = "60162"
var wTestCollection = "test"
var wTestBook = "users"
var wTestBookRecordsListLimit = "50"
var wTestBookRecordsMetaTags = "1"
var wTestKeyID = "4ec735ef-5574-445f-9f83-70aeccbe4d8e"
var wTestContentJSONdoc = `{"teste":"teste-123"}`

type TableDrivenRecords struct {
	collection string
	book       string
	key        string
	value      string
	expected   string
}

var arrRecordsTD = []TableDrivenRecords{
	{"test", "users", "123", "xpto", "success"},
	{"test", "users", "1234", "xpto", "success"},
	{"test", "users", "12345", "xpto", "success"},
	{"test", "users", "123456", "xpto", "success"},
}

// get

func GetN(n int) {

	for i := 0; i < n; i++ {

		GetWithLoadAndClose(wTestCollection, wTestBook, wTestKeyID, wTestBookRecordsMetaTags)

	}

}

func GetWithoutLoadAndClose(conn CONN, collection string, book string, key string, meta string) string {

	r := conn.Send(`route=collection.books.documents.get&collection=` + collection + `&book=` + book + `&key=` + key + `&meta=` + meta)
	return r

}

func GetWithLoadAndClose(collection string, book string, key string, meta string) {

	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()
	GetWithoutLoadAndClose(conn, collection, book, key, meta)
	conn.Close()

}

// insert

func InsertN(n int) {

	for i := 0; i < n; i++ {

		wKeyRecord := uuid.New().String()
		InsertWithLoadAndClose(wTestCollection, wTestBook, wKeyRecord, wTestContentJSONdoc)

	}

}

func InsertWithoutLoadAndClose(conn CONN, collection string, book string, key string, value string) string {

	r := conn.Send(`route=collection.books.documents.insert&collection=` + collection + `&book=` + book + `&key=` + key + `&value=` + value)
	return r

}

func InsertWithLoadAndClose(collection string, book string, key string, value string) {

	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()
	InsertWithoutLoadAndClose(conn, collection, book, key, value)
	conn.Close()

}

// list

func ListN(n int) {

	for i := 0; i < n; i++ {

		ListWithLoadAndClose(wTestCollection, wTestBook, wTestBookRecordsListLimit, wTestBookRecordsMetaTags)

	}

}

func ListWithLoadAndClose(collection string, book string, limit string, meta string) {

	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()

	ListWithoutLoadAndClose(conn, collection, book, limit, meta)

	conn.Close()

}

func ListWithoutLoadAndClose(conn CONN, collection string, book string, limit string, meta string) string {

	r := conn.Send(`route=collection.books.documents.list&collection=` + collection + `&book=` + book + `&limit=` + limit + `&meta=` + meta)
	return r

}

//
// Tests
//

// test / basic ops

func TestConnect(t *testing.T) {
	
	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()
	t.Log(conn)
}

// test / insert

func TestInsertUnique(t *testing.T) {
	
	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()
	r := InsertWithoutLoadAndClose(conn, "test", "users", "12345", "4567")
	got := gjson.Parse(r).Get("type").String()
	expected := "success"

	if got != "success" {
		t.Errorf("got %q, expected %q", got, expected)
	}

	conn.Close()

}

func TestInsertTableDriven(t *testing.T) {

	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()

	for _, test := range arrRecordsTD {

		r := InsertWithoutLoadAndClose(conn, test.collection, test.book, test.key, test.value)
		got := gjson.Parse(r).Get("type").String()

		// t.Log(gjson.Parse(r).String())

		if got != test.expected {
			t.Errorf("got %q, expected %q", got, test.expected)
		}

	}

	conn.Close()

}

// test / get

func TestGetDocumentNoExists(t *testing.T) {

	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()

	requestReturn := GetWithoutLoadAndClose(conn, wTestCollection, wTestBook, "1234.456-789-10.11.12", wTestBookRecordsMetaTags)

	got := gjson.Parse(requestReturn).Get("document").String()
	expected := ""

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}

	conn.Close()

}

//
// Benchmarks
//

// bench / list

func BenchmarkListWithLoadAndClose100(b *testing.B) {
	
	ListN(100)

}

func BenchmarkListWithLoadAndClose1000(b *testing.B) {
	
	ListN(1000)

}

func BenchmarkListWithLoadAndClose10000(b *testing.B) {
	
	ListN(10000)

}

func BenchmarkListWithLoadAndClose(b *testing.B) {

	for i := 0; i < b.N; i++ {

		ListWithLoadAndClose(wTestCollection, wTestBook, wTestBookRecordsListLimit, wTestBookRecordsMetaTags)

	}

}

func BenchmarkListWithoutLoadAndClose(b *testing.B) {

	loader := Load(wTestHost, wTestPort)
	conn, _ := loader.Connect()

	for i := 0; i < b.N; i++ {

		ListWithoutLoadAndClose(conn, wTestCollection, wTestBook, wTestBookRecordsListLimit, wTestBookRecordsMetaTags)

	}

	conn.Close()

}

// bench / get

func BenchmarkGetWithLoadAndClose100(b *testing.B) {
	
	GetN(100)

}

func BenchmarkGetWithLoadAndClose1000(b *testing.B) {
	
	GetN(1000)

}

func BenchmarkGetWithLoadAndClose10000(b *testing.B) {
	
	GetN(10000)

}

func BenchmarkGetWithLoadAndClose(b *testing.B) {

	for i := 0; i < b.N; i++ {

		GetWithLoadAndClose(wTestCollection, wTestBook, wTestKeyID, wTestBookRecordsMetaTags)

	}

}

// bench / insert

func BenchmarkInsert100(b *testing.B) {
	InsertN(100)
}

func BenchmarkInsert1000(b *testing.B) {
	InsertN(1000)
}

func BenchmarkInsert10000(b *testing.B) {
	InsertN(10000)
}

func BenchmarkInsertWithLoadAndClose(b *testing.B) {

	for i := 0; i < b.N; i++ {

		wKeyRecord := uuid.New().String()
		InsertWithLoadAndClose(wTestCollection, wTestBook, wKeyRecord, wTestContentJSONdoc)

	}

}
