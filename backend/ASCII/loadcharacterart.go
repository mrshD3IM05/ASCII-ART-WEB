package ASCII

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func loadCharacterArt(font string) ([]string, error) {
	file, err := os.ReadFile("backend/fonts/" + font + ".txt")
	if err != nil {
		log.Printf("error reading font %s: %v", font, err)
		return nil, err
	}
	asciifile := string(file)
	asciifile = strings.ReplaceAll(asciifile, "\r", "")
	asciifile = strings.Trim(asciifile, "\n")
	asciifile = strings.ReplaceAll(asciifile, "\n\n", "\n")
	spl := strings.Split(asciifile, "\n")
	if len(spl) < 95*8 {
		log.Printf("error reading font %s: %v", font, "font file incomplete")
		return nil, fmt.Errorf("incomplete")
	}
	return spl, nil
}
