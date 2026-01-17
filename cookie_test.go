package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "pannn"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("pannn")
	if err != nil {
		fmt.Fprint(w, "Get cookie fail")
	} else {
		fmt.Fprint(w, cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/setCookie", SetCookie)
	mux.HandleFunc("/getCookie", GetCookie)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "/?name=mananta", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "/?name=Mananta", nil)
	cookie := new(http.Cookie)
	cookie.Name = "pannn"
	cookie.Value = "Mananta"

	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(cookie.Name, cookie.Value)
	fmt.Println(string(body))
}
