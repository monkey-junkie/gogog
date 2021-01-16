package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"os"
)

func sqlTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	q := fmt.Sprintf("SELECT * FROM foo where Name = '%s'", os.Args[1])
	rows, err := db.Query(q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func handler(req *http.Request) {
	cmdName := req.URL.Query()["cmd"][0]
	url := fmt.Sprintf("/?a=5&bbbbb=%s", cmdName)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    	param1 := r.URL.Query()["param1"][0]
	value, _ := strconv.Atoi(param1)
	out := 1337 / value
	fmt.Println(out)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Contamination() {
	url := os.Getenv("tainted_url")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}

func main() {
	Contamination()
	sqlTest()
	fmt.Printf("Hello actions world\n")
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}
