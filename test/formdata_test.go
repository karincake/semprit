package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	s "github.com/karincake/semprit"
)

func TestMediumDataNormal(t *testing.T) {
	// mock request
	r, _ := http.NewRequest("POST", "/", nil)
	r.PostForm = make(url.Values)
	r.PostForm.Add("name", "Santo Sembodo")
	r.PostForm.Add("married", "true")
	r.PostForm.Add("score", "20")
	r.PostForm.Add("creditScore", "-10")
	r.PostForm.Add("age", "80")
	r.PostForm.Add("hoursActive", "181500")
	r.PostForm.Add("income", "800")
	r.PostForm.Add("netWorth", "200000")

	// data and wanted
	data := DataMedium{}
	want := DataMedium{
		Name:        "Santo Sembodo",
		Married:     true,
		Score:       20,
		CreditScore: -10,
		Age:         80,
		HoursActive: 181500,
		Income:      800,
		NetWorth:    200000,
	}

	if err := s.HttpFormData(&data, r); err != nil {
		t.Error("failed to parse request: ", err)
	} else {
		if data != want {
			fmt.Println(data)
			fmt.Println(want)
			t.Error("failed to parse ")
		}
	}
}
