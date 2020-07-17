// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"
// )

// // TestScrapeSMI tests if ScrapeSMI scrapes 3 stock market indexes
// func TestScrapeSMI(t *testing.T) {
// 	fmt.Println("Test ScrapeSMI")
// 	indexesCount := 3
// 	result := ScrapeSMI(w http.ResponseWriter, r *http.Request)
// 	if indexesCount != result.indexes.count {
// 		t.Error("Failed")
// 	}
// }