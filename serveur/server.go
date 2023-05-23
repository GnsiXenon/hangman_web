package main

import (
	"fmt"
	"hangmanWeb"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("../template/index.html"))
	tmplt.Execute(w, nil)
}
func lose(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("../template/lose.html"))
	tmplt.Execute(w, nil)
}
func getdata(w http.ResponseWriter, r *http.Request) {
	var letter = r.FormValue("name")
	data, site := hangmanWeb.GameStatus(letter)
	var tmplt = template.Must(template.ParseFiles(site))
	tmplt.Execute(w, data)
}

func menu(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	var tmplt = template.Must(template.ParseFiles("../template/game.html"))
	tmplt.Execute(w, nil)
}
func rules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	var tmplt = template.Must(template.ParseFiles("../template/rules.html"))
	tmplt.Execute(w, nil)
}
func GetDifficulty(w http.ResponseWriter, r *http.Request) {
	var difficulty = r.FormValue("difficulty")
	file := ""
	fmt.Println(difficulty)
	switch difficulty {
	case "Facile":
		file = "words.txt"
	case "Moyenne":
		file = "words2.txt"
	case "Difficile":
		file = "words3.txt"
	}
	data := hangmanWeb.Initial(file)
	var tmplt = template.Must(template.ParseFiles("../template/game.html"))
	tmplt.Execute(w, data)
}
func win(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	var tmplt = template.Must(template.ParseFiles("../template/win.html"))
	tmplt.Execute(w, nil)
}

func addword(w http.ResponseWriter, r *http.Request) {
	var word = r.FormValue("word")
	data := hangmanWeb.AddWord(word)
	var tmplt = template.Must(template.ParseFiles("../template/win.html"))
	tmplt.Execute(w, data)
}
func serverOn() {
	styles := http.FileServer(http.Dir("../template/assets"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	http.HandleFunc("/", index)
	http.HandleFunc("/lose", lose)
	http.HandleFunc("/getdata", getdata)
	http.HandleFunc("/menu", menu)
	http.HandleFunc("/rules", rules)
	http.HandleFunc("/GetDifficulty", GetDifficulty)
	http.HandleFunc("/win", win)
	http.HandleFunc("/AddWord", addword)
	http.ListenAndServe(":8080", nil)
}
func main() {
	serverOn()
}
