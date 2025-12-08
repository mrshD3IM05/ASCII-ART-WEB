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
    *   CHECKING IF THE INPUT STRING IS EMPTY OR IF IT IS JUST WHITESPACES. IF IT IS, RETURNING THE INPUT STRING AS IS.
    *   The input string is split by newlines to handle multi-line input.
    *   For each line of input, the program iterates through the 8 vertical "slices" of the font characters EXCEPT FOR EMPTY LINES. WE JUST PRINT THE EMPTY LINE AND CONTINUE.
    *   It builds the output line by line: for the first line of the output, it concatenates the 1st line of the ASCII representation for each character in the input string. Then it does the same for the 2nd line, and so on.
4.  **Rendering**: The constructed ASCII string is sent back to the client and displayed in a `<pre>` tag to preserve formatting.

### Handlers

The backend uses standard HTTP handlers to manage requests and responses.

| Function | Endpoint | Method | Description |
| :--- | :--- | :--- | :--- |
| **`indexHandler`** | `/` | `GET` | Serves the home page. It renders `index.html` within the `App.html` layout, displaying the input form. Returns `404 Not Found` for any unknown paths. |
| **`submitHandler`** | `/ascii-art` | `POST` | Processes form submissions. It validates the request, generates the ASCII art, and renders `result.html` with the output. Errors (e.g., bad request) are handled via `serveError`. |
| **`serveError`** | N/A | N/A | A utility function used by other handlers to render standardized error pages using `error.html`. |
| **`renderTemplate`** | N/A | N/A | A helper function defined in `rendertemplate.go` to standardize template execution and error handling. |

### Templates

The application utilizes Go's `html/template` package with a modular structure.

*   **`frontend/App.html` (Layout)**:
    *   Acts as the master template containing the common structure (HTML boilerplate, header).
    *   Includes inline CSS styles for the application theme (`.card`, `.btn`, etc.).
    *   Defines a `{{block "content" .}}{{end}}` placeholder where other templates inject their specific content.

*   **`frontend/pages/index.html` (Home)**:
    *   Renders the main input interface.
    *   Includes the text area for user input and radio buttons for font selection.
    *   Fills the "content" block of `App.html`.

*   **`frontend/pages/result.html` (Output)**:
    *   Displays the generated ASCII art inside a `<pre>` tag.
    *   Provides a "Go Back" button to return to the home page.
    *   Fills the "content" block of `App.html`.

*   **`frontend/pages/error.html` (Error)**:
    *   Displays user-friendly error messages (e.g., "404 Not Found", "500 Internal Server Error").
    *   Fills the "content" block of `App.html`.