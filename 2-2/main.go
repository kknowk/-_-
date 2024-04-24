package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>String Concatenation</title>
</head>
<body>
    <form method="POST" action="/concatenate">
		文字列A: <input type="text" name="文字列A" />
        文字列B: <input type="text" name="文字列B" />
        <input type="submit" value="Concatenate!" />
    </form>
    {{if .}}
    <p>Result: {{ . }}</p>
    {{end}}
</body>
</html>
`))

func concatenateStrings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	stringA := r.FormValue("文字列A")
	stringB := r.FormValue("文字列B")
	result := stringA + stringB

	tmpl.Execute(w, result)
}

func main() {
	http.HandleFunc("/concatenate", concatenateStrings)

	fmt.Println("Server is running at http://localhost:8080/concatenate")
	http.ListenAndServe(":8080", nil)
}
