package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/sahilm/fuzzy"
)

type Element struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

//go:embed static
var static embed.FS

var candidates []string
var questions []Element

func serve_static() http.Handler {
	sub, err := fs.Sub(static, "static/public")
	if err != nil {
		panic(err)
	}
	file_server := http.FileServer(http.FS(sub))
	return file_server
}

func SearchAnswer(query string) (string, error) {

	var answer string
	matches := fuzzy.Find(query, candidates)

	if len(matches) != 1 {
		return answer, errors.New("Couldn't find specific answer")
	}

	answer = questions[matches[0].Index].Answer
	return answer, nil

}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	query := r.FormValue("search")

	trim := strings.TrimSpace(query)

	if len(trim) == 0 {
		w.Write([]byte{})
		return
	}
	answer, err := SearchAnswer(trim)

	if err != nil {
		w.Write([]byte("<b>Nie znaleziono odpowiedzi w bazie</b>"))
		return
	}
	w.Write([]byte(answer))
}

func main() {

	json_data, _ := static.ReadFile("static/out.json")

	json.Unmarshal(json_data, &questions)

	candidates = make([]string, len(questions))
	for i, q := range questions {
		candidates[i] = q.Question
	}

	fmt.Println("Starting server")
	http.Handle("/", serve_static())
	http.HandleFunc("/search", HandleSearch)

	err := http.ListenAndServe(":3342", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
