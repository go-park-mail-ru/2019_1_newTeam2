package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
)

func (server *Server) IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	_, err := server.GetUserIdFromCookie(r)
	return err == nil
}

func (server *Server) GetUserIdFromCookie(r *http.Request) (int, error){
	cookie, err := r.Cookie(server.CookieField)
	if err != nil {
		return 0, err
	}
	ctx := context.Background()
	StrUserId, err := server.AuthClient.GetIdFromCookie(ctx, &authorization.AuthCookie{
		Data: cookie.Value,
	})
	return int(StrUserId.UserId), nil
}

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
