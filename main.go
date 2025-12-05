package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}
func Run(input string, banner string) (string, error) {
	font_path := "banners/" + banner + ".txt"

	linesInput := strings.Split(input, "\\n")
	if linesInput[0] == "" && linesInput[1] == "" {
		linesInput = linesInput[1:] // remove extra element from strings.Split
	}

	content, err := os.ReadFile(font_path)
	if err != nil {
		return "", fmt.Errorf("banner not found")
	}

	text := strings.ReplaceAll(string(content), "\r\n", "\n")

	fontLines := strings.Split(text, "\n")
	final := ""
	for _, line := range linesInput {
		if line == "" {
			final += "\n"
			continue
		}
		runes := []rune(line)
		chars := make([][]string, len(runes))

		for i, char := range runes {
			if char < ' ' || char > '~' {
				return "", fmt.Errorf("invalid character")
			}

			index := int(((char - ' ') * 9) + 1)
			chars[i] = fontLines[index : index+8]
		}

		for height := 0; height < 8; height++ {
			for i := range chars {
				final += chars[i][height]
			}
			final += "\n"
		}
	}
	return final, nil
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	ascii, err := Run(text, banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := struct{ Result string }{Result: ascii}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, data)
}
