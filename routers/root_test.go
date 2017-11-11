package routers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"log"
	"bytes"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal("not equal")
}


func TestServer(t *testing.T) {
	var jsonString string = "{\"message\":\"ok\",\"data\":{\"name\":\"thoreau\",\"email\":\"zz.thoreau@gmail.com\",\"ext\":1,\"privileged\":false}}"
	server := httptest.NewServer(UserResource())
	defer server.Close()
	resp, err := http.Get(server.URL+"/api/v1/user")
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	assertEqual(t,jsonString,buf.String())

}
