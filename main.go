package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// db, err := sql.Open("mysql", "romdlnyq_master:}-j+ML08O[Ba@tcp(localhost:3306)/romdlnyq_master")
	connString := "romdlnyq_master:}-j+ML08O[Ba@tcp(localhost:3306)/romdlnyq_master"
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO `request` (`id`, `name`, `email`, `pass`, `repass`, `status`) VALUES (NULL, 'mostain', 'mostain@email.com', 'mostain123', 'mostain123', '1');")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("db connection successfully")

}
func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/register", register)
	http.HandleFunc("/forgot", forgot)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/default.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func forgot(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/default.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/forgot.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/default.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/register.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

// request to ...
func request(w http.ResponseWriter, r *http.Request) {

	// by loop
	r.ParseForm()

	for key, val := range r.Form { // slice string

		fmt.Println(key, val)
	}

	fmt.Fprint(w, `Recived`)
}
