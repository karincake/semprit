package test

import (
	"net/http"
	"testing"

	s "github.com/karincake/semprit"
)

func BenchmarkQueryParamMediumDataNormal(b *testing.B) {
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
	r, _ := http.NewRequest("GET", "/", nil)
	// mock url encoded like : ?name=Santo%20Sembodo&address=JL%20Localhost%202023&married=true&score=20 ...
	q := r.URL.Query()
	q.Add("name", "Santo Sembodo")
	q.Add("address", "JL Localhost 2023")
	q.Add("married", "true")
	q.Add("score", "20")
	q.Add("creditScore", "-50")
	q.Add("socialScore", "25000")
	q.Add("age", "80")
	q.Add("bloodPressure", "150")
	q.Add("hoursActive", "22400")
	q.Add("income", "1100")
	q.Add("netWorth", "412300")
	r.URL.RawQuery = q.Encode()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.UrlQueryParam(&data, *r.URL)
	}
}

func BenchmarkQueryParamMediumDataCustom(b *testing.B) {
	// data
	data := DataLargeCT{}

	// mock request
	r, _ := http.NewRequest("GET", "/", nil)
	// mock url encoded like : ?nameValidity=valid&marriedStatus=true&scoreClass=1 ...
	q := r.URL.Query()
	q.Add("nameValidity", "valid")
	q.Add("marriedStatus", "true")
	q.Add("scoreClass", "1")
	q.Add("creditScoreClass", "1")
	q.Add("ageClass", "1")
	q.Add("hoursActiveClass", "1")
	q.Add("incomeClass", "1")
	q.Add("netWorthClass", "1")
	r.URL.RawQuery = q.Encode()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.UrlQueryParam(&data, *r.URL)
	}
}
