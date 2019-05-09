package server_test

/*import (
	"github.com/user/2019_1_newTeam2/pkg/apps/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TypeRequestTestCase struct {
	url     string
	string1 string
	string2 string
}

func TestTypeRequest(t *testing.T) {
	cases := []TypeRequestTestCase{
		{"/localhost/users/",
			"localhost",
			"/users/",
		},
		{
			" url",
			"url",
			"/",
		},
	}
	for _, it := range cases {
		str1, str2 := server.TypeRequest(it.url)
		if str1 != it.string1 || str2 != it.string2 {
			log.Println(str1, it.string1)
			log.Println(str2, it.string2)
			t.Error("fail, trim doesn't work")
		}
	}

}

type ParseParamsCase struct {
	rows   string
	page   string
	ifRows bool
	ifPage bool
	ifErr  bool
	res    int
}

func TestParseParams(t *testing.T) {
	cases := []ParseParamsCase{
		{
			rows:   "5",
			ifRows: true,
			ifPage: false,
			ifErr:  true,
			res:    http.StatusBadRequest,
		},
		{
			page:   "6",
			ifPage: true,
			ifRows: false,
			ifErr:  true,
			res:    http.StatusBadRequest,
		},
		{
			page:   "7",
			rows:   "jncj",
			ifPage: true,
			ifRows: true,
			ifErr:  true,
			res:    http.StatusBadRequest,
		},
		{
			page:   "dmdmdm",
			rows:   "7",
			ifPage: true,
			ifRows: true,
			ifErr:  true,
			res:    http.StatusBadRequest,
		},
	}
	for _, it := range cases {
		r, _ := http.NewRequest(http.MethodGet, "/users", nil)
		q := r.URL.Query()
		if it.ifRows {
			q.Add("rows", it.rows)
		}
		if it.ifPage {
			q.Add("page", it.page)
		}
		r.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		page := 0
		rows := 0
		err := server.ParseParams(w, r, &page, &rows)
		if err != nil {
			if !it.ifErr {
				t.Error("oops, unexpected error")
			} else {
				if w.Result().StatusCode != it.res {
					t.Errorf("Statuses don't match expected %v and got %v", it.res, w.Result().StatusCode)
				}
			}
		}
	}

}
*/
