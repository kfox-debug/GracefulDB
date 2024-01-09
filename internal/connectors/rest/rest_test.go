package rest

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Kwynto/GracefulDB/internal/config"
	"github.com/Kwynto/GracefulDB/pkg/lib/closer"
)

func Test_home(t *testing.T) {
	t.Run("home() function testing - positive", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		home(w, r) // calling the tested function
		status := w.Code
		if status != http.StatusNotFound {
			t.Error("home() error.")
		}
	})

	t.Run("home() function testing - negative", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		home(w, r) // calling the tested function
		status := w.Code
		if status == http.StatusOK {
			t.Error("home() error.")
		}
	})

}

func Test_squery(t *testing.T) {
	t.Run("squery() function testing - GET error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		squery(w, r) // calling the tested function
		status := w.Code
		if status != http.StatusMethodNotAllowed {
			t.Error("squery() error. GET error.")
		}
	})

	t.Run("squery() function testing - POST - an empty query", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", nil)

		squery(w, r) // calling the tested function
		status := w.Code
		if status != http.StatusBadRequest {
			t.Error("squery() error. POST - an empty query.")
		}
	})

	t.Run("squery() function testing - POST - positive", func(t *testing.T) {
		w := httptest.NewRecorder()

		form := url.Values{}
		form.Add("instruction", "instruction")
		form.Add("placeholder", "[]")
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		// r.Form = form
		r.PostForm = form

		squery(w, r) // calling the tested function
		status := w.Code
		if status != http.StatusOK {
			t.Errorf("squery() error. POST - an empty query. %v", status)
		}
	})
}

func Test_vquery(t *testing.T) {
	t.Run("vquery() function testing - GET error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		vquery(w, r) // calling the tested function
		status := w.Code
		if status != http.StatusMethodNotAllowed {
			t.Error("vquery() error. GET error.")
		}
	})

	t.Run("vquery() function testing - POST - positive", func(t *testing.T) {
		w := httptest.NewRecorder()

		form := url.Values{}
		form.Add("instruction", "instruction")
		r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		// r.Form = form
		r.PostForm = form

		vquery(w, r) // calling the tested function
		status := w.Code
		if status != http.StatusOK {
			t.Errorf("vquery() error. POST - an empty query. %v", status)
		}
	})
}

func Test_routes(t *testing.T) {
	t.Run("routes() function testing", func(t *testing.T) {
		res := routes() // calling the tested function

		if reflect.TypeOf(res) != reflect.TypeOf(&http.ServeMux{}) {
			t.Error("routes() error = The function returns the wrong type")
		}
	})

}

func Test_Start_and_Shutdown(t *testing.T) {
	t.Run("Start() and Shutdown() function testing", func(t *testing.T) {
		tf := "../../../../config/develop.yaml"
		config.MustLoad(tf)
		go Start(&config.DefaultConfig) // calling the tested function
		closer.AddHandler(Shutdown)
		time.Sleep(2 * time.Second)
		// srvRest.Shutdown(context.Background())
		Shutdown(context.Background(), closer.CloseProcs)

		if reflect.TypeOf(muxRest) != reflect.TypeOf(&http.ServeMux{}) {
			t.Error("Start() error = The function has created an incorrect dependency.")
		}

		if reflect.TypeOf(srvRest) != reflect.TypeOf(&http.Server{}) {
			t.Error("Start() error = The function has created an incorrect dependency.")
		}

		if closer.CloseProcs.Counter != 0 {
			t.Errorf("Shutdown() error: %v.", closer.CloseProcs.Counter)
		}
	})

	t.Run("Shutdown() function testing - negative", func(t *testing.T) {
		Shutdown(context.Background(), closer.CloseProcs)

		if len(closer.CloseProcs.Msgs) > 0 {
			t.Errorf("Shutdown() error.")
		}
	})
}

func Test_Shutdown(t *testing.T) {
	t.Run("Shutdown() function testing - negative", func(t *testing.T) {
		address = fmt.Sprintf("%s:%s", "0.0.0.0", "31337")
		muxRest = routes()
		srvRest = &http.Server{
			Addr:    address,
			Handler: muxRest,
		}

		ctx, cf := context.WithTimeout(context.Background(), 500*time.Millisecond)
		// defer cf()

		go Shutdown(ctx, closer.CloseProcs)
		cf()
		time.Sleep(3 * time.Second)

		if len(closer.CloseProcs.Msgs) > 0 {
			t.Errorf("Shutdown() error.")
		}
	})
}
