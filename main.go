package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {

	//db, err = sql.Open("mysql", "romdlnyq_master:}-j+ML08O[Ba@tcp(199.188.200.231:3306)/romdlnyq_master")

	db, err = sql.Open("mysql", "root:master123@tcp(127.0.0.1:3306)/master")

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	// insert, err := db.Query(sql)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

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
	name := r.FormValue("name")
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	repass := r.FormValue("repass")

	qs := "INSERT INTO `request` (`id`, `name`, `email`, `pass`, `repass`, `status`) VALUES (NULL, '%s', '%s', '%s', '%s',  '1');"
	sql := fmt.Sprintf(qs, name, email, pass, repass)
	//fmt.Println(sql)
	insert, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Fprintf(w, `ok`)

	// by loop-----
	// r.ParseForm()

	// for key, val := range r.Form { // slice string

	// 	fmt.Println(key, val)
	// }

	// fmt.Fprint(w, `Recived`)
}
