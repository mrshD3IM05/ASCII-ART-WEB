package ASCII

import (
	"log"
	"net/http"
	"strings"
)

func CreateASCIIArtTable(input, font string) (string, int) {
	font = strings.ToLower(font)

	asciiArt, err := loadCharacterArt(font)
	if err != nil {
		log.Printf("failed to load font %s: %v", font, err)
		return "", http.StatusNotFound
	}

	input1 := strings.ReplaceAll(input, "\r", "")
	if strings.Trim(input1, "\n") == "" {
		return input, http.StatusOK
	}

	var result string
	var temp string
	split := strings.Split(input1, "\n")
	for _, line := range split {
		if line == "" {
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "", http.StatusBadRequest
				}
				index := int(char - 32)
				if index*8+i >= len(asciiArt) {
					log.Printf("index out of bounds: %d", index*8+i)
					return "", http.StatusInternalServerError
				}
				temp += asciiArt[index*8+i]
			}
			temp = (strings.TrimRight(temp, " "))
			result += temp + "\n"
			temp = ""
		}
	}
	result = (strings.TrimRight(result, "\n"))
	return result, http.StatusOK
}
