package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//r.PostFormValue("name")

	name := r.PostForm.Get("name")
	age := r.PostForm.Get("age")

	fmt.Fprintf(w, "Form Post Name : %s, Age : %s", name, age)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("name=Mananta&age=20")
	request := httptest.NewRequest(http.MethodPost, "/", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
