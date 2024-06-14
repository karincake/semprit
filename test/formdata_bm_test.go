package test

import (
	"net/http"
	"net/url"
	"testing"

	s "github.com/karincake/semprit"
)

func BenchmarkFormDataLargeDataNormal(b *testing.B) {
	// data
	dataAddress := ""
	dataSocialScore := 0
	dataHoursActive := uint(0)
	data := DataLarge{
		Address:     &dataAddress,
		SocialScore: &dataSocialScore,
		HoursActive: &dataHoursActive,
	}

	// mock request, make it fast: import from "want" becase what we
	// need is the process
	r, _ := http.NewRequest("POST", "/", nil)
	r.PostForm = make(url.Values)
	r.PostForm.Add("name", "Santo Sembodo")
	r.PostForm.Add("address", "JL Localhost 2023")
	r.PostForm.Add("married", "true")
	r.PostForm.Add("score", "20")
	r.PostForm.Add("creditScore", "-50")
	r.PostForm.Add("socialScore", "25000")
	r.PostForm.Add("age", "80")
	r.PostForm.Add("bloodPressure", "150")
	r.PostForm.Add("hoursActive", "22400")
	r.PostForm.Add("income", "1100")
	r.PostForm.Add("netWorth", "412300")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.HttpFormData(&data, r)
	}
}

func BenchmarkFormDataLargeDataCustom(b *testing.B) {
	// data
	data := DataLargeCT{}

	// mock request
	r, _ := http.NewRequest("POST", "/", nil)
	r.PostForm = make(url.Values)
	r.PostForm.Add("nameValidity", "valid")
	r.PostForm.Add("marriedStatus", "true")
	r.PostForm.Add("scoreClass", "1")
	r.PostForm.Add("creditScoreClass", "1")
	r.PostForm.Add("ageClass", "1")
	r.PostForm.Add("hoursActiveClass", "1")
	r.PostForm.Add("incomeClass", "1")
	r.PostForm.Add("netWorthClass", "1")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.HttpFormData(&data, r)
	}
}
