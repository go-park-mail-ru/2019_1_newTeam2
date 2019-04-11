package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func TypeRequest(url string) (string, string) {
	separatorPosition := strings.Index(url[1:], "/")
	if separatorPosition < 0 {
		return url[1:], "/"
	}
	separatorPosition++
	return url[1:separatorPosition], url[separatorPosition:]
}

func ParseParams(w http.ResponseWriter, r *http.Request,
	page *int, rowsNum *int) error {
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("no pages in query")
	}
	rows, ok := r.URL.Query()["rows"]
	if !ok || len(rows[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("no rows in query")
	}

	pg, err := strconv.Atoi(pages[0])
	*page = pg
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("page incorrect")
	}
	*rowsNum, err = strconv.Atoi(rows[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("rows incorrect")
	}
	return nil
}
