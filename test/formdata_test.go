package test

import (
	"net/http"
	"net/url"
	"testing"

	s "github.com/karincake/semprit"
)

func TestFormDataOnSmallString(t *testing.T) {
	// mock request
	r, _ := http.NewRequest("POST", "/", nil)
	r.PostForm = make(url.Values)
	r.PostForm.Add("name", "Santo Sembodo")

	// data and wanted
	data := DataSmallString{}
	want := DataSmallString{
		Name: "Santo Sembodo",
	}

	if err := s.FormDataFromHttp(&data, r); err != nil {
		t.Error("failed to parse request: ", err)
	} else {
		if data != want {
			t.Error("failed to parse")
		}
	}
}

func TestFormDataOnSmallBool(t *testing.T) {
	// mock request
	r, _ := http.NewRequest("POST", "/", nil)
	r.PostForm = make(url.Values)
	r.PostForm.Add("married", "true")

	// data and wanted
	data := DataSmallBool{}
	want := DataSmallBool{
		Married: true,
	}

	if err := s.FormDataFromHttp(&data, r); err != nil {
		t.Error("failed to parse request: ", err)
	} else {
		if data != want {
			t.Error("failed to parse")
		}
	}
}
