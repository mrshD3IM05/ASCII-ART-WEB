package main

import (
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", submitHandler)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{ Result template.HTML }{}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	font := r.FormValue("font")

	ascii := CreateASCIIArtTable(text, font)
	escaped := html.EscapeString(ascii)

	data := struct{ Result template.HTML }{Result: template.HTML("<pre>" + escaped + "</pre>")}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}

func CreateASCIIArtTable(input, font string) string {
	font = strings.ToLower(font)
	switch font {
	case "standard", "shadow", "thinkertoy":
	default:
		return "FONT NOT FOUND"
	}
	result := ""
	asciiArt, err := loadCharacterArt(font)
	if err != nil {
		return "FONT NOT FOUND"
	}
	input1 := strings.ReplaceAll(input, "\r", "")
	if strings.Trim(input1, "\n") == "" {
		return input
	}
	split := strings.Split(input1, "\n")
	for _, line := range split {
		if line == "" {
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "A CHARACTER IS OUT OF RANGE"
				}
				index := int(char - 32)
				result += asciiArt[index*8+i]
			}
			result += "\n"
		}

	}
	return result
}

func loadCharacterArt(font string) ([]string, error) {
	file, err := os.ReadFile("./backend/fonts/" + font + ".txt")
	if err != nil {
		return nil, err
	}
	asciifile := string(file)
	asciifile = strings.ReplaceAll(asciifile, "\r", "")
	asciifile = strings.Trim(asciifile, "\n")
	asciifile = strings.ReplaceAll(asciifile, "\n\n", "\n")
	spl := strings.Split(asciifile, "\n")
	return spl, nil
}
