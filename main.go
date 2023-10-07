package main

import (
	
	"fmt"
	// "time"

	// "github.com/google/uuid"

	// rdata "github.com/reactioon/rdata-go/rdata"

)

var (
	_ = fmt.Println
)

func main() {

	// host := "192.168.1.210"
	// host := "192.168.1.202"
	// port := "60162"

	// client := rdata.CLIENT{}
	// loader := client.Load(host, port)

	// ListMultiples(loader)

	fmt.Println("[alert] to see more details how to run, check the tests.")

}

// func InsertMultiples(loader rdata.CLIENT) {

// 	dateNow := time.Now()

// 	conn,_ := loader.Connect()
// 	defer conn.Close()

// 	// Qty := 100000
// 	Qty := 10000
// 	// Qty := 1000
// 	// Qty := 100
// 	// Qty := 10
	
// 	var i int
// 	for i=0; i < Qty; i++ {
		
// 		key := uuid.New().String()
// 		conn.Send(`route=collection.books.documents.insert&collection=test&book=users&key=`+key+`&value={"teste":"teste-123"}`)

// 	}

// 	elapsed := fmt.Sprintf("%s", time.Since(dateNow))

// 	fmt.Println("Executed INSERT", i, "times")
// 	fmt.Println("Time elapsed: ", elapsed)
// 	fmt.Println("Done!")

// }

// func GetMultiples(loader rdata.CLIENT) {

// 	dateNow := time.Now()

// 	conn,_ := loader.Connect()
// 	defer conn.Close()
	
// 	// Qty := 100000
// 	// Qty := 10000
// 	Qty := 1000
// 	// Qty := 100
// 	// Qty := 10

// 	var i int
// 	for i=0; i < Qty; i++ {
		
// 		key := "4ec735ef-5574-445f-9f83-70aeccbe4d8e"
// 		conn.Send(`route=collection.books.documents.get&collection=test&book=test&key=`+key+`&meta=1`) 

// 	}

// 	elapsed := fmt.Sprintf("%s", time.Since(dateNow))

// 	fmt.Println("Executed GET", i, "times")
// 	fmt.Println("Time elapsed: ", elapsed)
// 	fmt.Println("Done!")

// }

// func ListMultiples(loader rdata.CLIENT) {

// 	// dateNow := time.Now()

// 	conn,_ := loader.Connect()
// 	defer conn.Close()
	
// 	// Qty := 100000
// 	// Qty := 10000
// 	// Qty := 1000
// 	// Qty := 100
// 	// Qty := 50

// 	// var i int
// 	// for i=0; i < Qty; i++ {

// 		dateNow := time.Now()
// 		_ = conn.Send(`route=collection.books.documents.list&collection=test&book=users&limit=50&meta=1`)
// 		// _ = conn.Send(`route=core.metrics`)
// 		// _ = conn.Send(`route=collection.metrics&collection=test`)
// 		// _ = conn.Send(`route=collection.books.metrics&collection=test&book=users`)

// 		// _ = loader.Connect().Send(`route=core.metrics`)
// 		// _ = loader.Connect().Send(`route=about`)
// 		// go loader.Connect().Send(`route=core.metrics`)

// 		elapsed := fmt.Sprintf("%s", time.Since(dateNow))
// 		fmt.Println(elapsed)

// 		// time.Sleep(1 * time.Second)

// 	// }

// 	fmt.Println("Done!")

// }