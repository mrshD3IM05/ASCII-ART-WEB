# ASCII Art Web

## Description
ASCII Art Web is a web application that converts normal text into ASCII art using various artistic fonts. Users can type text, choose a font style (Standard, Shadow, or Thinkertoy), and view the generated ASCII art.

## Authors
- ACHENT - [GitHub](https://github.com/chentaymane)
- AHADDOU - [GitHub](https://github.com/mrshD3IM05)

## Usage: how to run
To run the application locally, follow these steps:

1.  **Prerequisites**: Ensure you have [Go](https://go.dev/dl/) installed on your machine.
2.  **Clone/Navigate**: Open your terminal and navigate to the project root directory.
3.  **Run**: Execute the following command:
    ```bash
    go run main.go
    ```
4.  **Access**: Open your web browser and go to:
    [http://localhost:8080](http://localhost:8080)

## Implementation details: algorithm
The core logic resides in reduced complexity steps:

1.  **Input Handling**: The server accepts the user's text and selected font via a POST request.
2.  **Font Loading**:
    *   The backend reads the corresponding ASCII font file (e.g., `standard.txt`).
    *   REPLACE all \r with empty string.
    *   TRIM the input string to remove leading and trailing newlines.
    *   REPLACE all \n\n with \n.
    *   ALL TO REMOVE EXTRA SPACES just to make sure the input is clean AND each character is EXACTLY 8 lines. Each character in these files is typically represented by a grid (e.g., 8 lines high). 
    *   CHECKING IF WE HAVE 8 LINES PER CHARACTER FOR ALL 95 CHARACTERS AND IF NOT, RETURNING ERROR AND STATUS CODE 400.
3.  **Processing**:
    *   REPLACE ALL \r WITH EMPTY STRING.
    *   TRIM THE INPUT STRING TO REMOVE LEADING AND TRAILING NEWLINES.
    *   CHECKING IF THE INPUT STRING IS EMPTY OR IF IT IS JUST WHITESPACES. IF IT IS, RETURNING THE INPUT STRING AS IS. 
    *   CHECKING IF WE HAVE 8 LINES PER CHARACTER FOR ALL 95 CHARACTERS AND IF NOT, RETURNING ERROR AND STATUS CODE 400.
    *   The input string is split by newlines to handle multi-line input.
    *   For each line of input, the program iterates through the 8 vertical "slices" of the font characters.
    *   It builds the output line by line: for the first line of the output, it concatenates the 1st line of the ASCII representation for each character in the input string. Then it does the same for the 2nd line, and so on.
4.  **Rendering**: The constructed ASCII string is sent back to the client and displayed in a `<pre>` tag to preserve formatting.