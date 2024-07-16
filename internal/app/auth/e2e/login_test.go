package e2e

import (
	"chat/internal/app/auth/env"
	"chat/internal/app/auth/httpserver"
	"chat/internal/app/auth/login"
	"chat/internal/pkg/test"
	"flag"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		env.Init()
		router := httpserver.GetRouter()
		testServer = httptest.NewServer(router)
		defer testServer.Close()
		m.Run()
	}
}

func TestLogin(t *testing.T) {
	reqUrl := testServer.URL + login.Path
	res, err := http.Post(reqUrl, gin.MIMEJSON, nil)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	test.AssertEqual(res.StatusCode, http.StatusOK, t)
	test.AssertEqual("Logged in", string(body), t)
}
