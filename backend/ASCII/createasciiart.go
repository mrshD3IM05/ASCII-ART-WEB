package ASCII

import (
	"log"
	"net/http"
	"strings"
)

func CreateASCIIArtTable(input, font string) (string, int, string) {
	font = strings.ToLower(font)

	asciiArt, err := loadCharacterArt(font)
	if err != nil {
		log.Printf("failed to load font %s: %v", font, err)
		return "", 500, "Intenal Error: font file not found"
	}

	input1 := strings.ReplaceAll(input, "\r", "")
	if strings.Trim(input1, "\n") == "" {
		return input, 200, ""
	}

	var result strings.Builder
	split := strings.Split(input1, "\n")
	for _, line := range split {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "", http.StatusBadRequest, "Bad Request: character out of range"
				}
				index := int(char - 32)
				if index*8+i >= len(asciiArt) {
					log.Printf("index out of bounds: %d", index*8+i)
					return "", http.StatusInternalServerError, "Internal Server Error"
				}
				result.WriteString(asciiArt[index*8+i])
			}
			result.WriteString("\n")
		}
	}
	return result.String(), 200, ""
}
