package test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	s "github.com/karincake/semprit"
)

func TestQueryParamMediumDataNormal(t *testing.T) {
	// data
	dataAddress := ""
	dataSocialScore := 0
	dataHoursActive := uint(0)
	data := DataLarge{
		Address:     &dataAddress,
		SocialScore: &dataSocialScore,
		HoursActive: &dataHoursActive,
	}

	// want
	wantAddress := "JL Localhost 2023"
	wantSocialScore := 25000
	wantHoursActive := uint(30000)
	want := DataLarge{
		Name:          "Santo Sembodo",
		Address:       &wantAddress,
		Married:       true,
		Score:         20,
		CreditScore:   -10,
		SocialScore:   &wantSocialScore,
		Age:           80,
		BloodPressure: 240,
		HoursActive:   &wantHoursActive,
		Income:        800,
		NetWorth:      200000,
	}

	// mock request, make it fast: import from "want" becase what we
	// need is the process
	r, _ := http.NewRequest("GET", "/", nil)
	// mock url encoded like : ?name=Santo%20Sembodo&address=JL%20Localhost%202023&married=true&score=20 ...
	q := r.URL.Query()
	q.Add("name", want.Name)
	q.Add("address", *want.Address)
	q.Add("married", "true")
	q.Add("score", strconv.Itoa(int(want.Score)))
	q.Add("creditScore", strconv.Itoa(int(want.CreditScore)))
	q.Add("socialScore", strconv.Itoa(int(*want.SocialScore)))
	q.Add("age", strconv.Itoa(int(want.Age)))
	q.Add("bloodPressure", strconv.Itoa(int(want.BloodPressure)))
	q.Add("hoursActive", strconv.Itoa(int(*want.HoursActive)))
	q.Add("income", fmt.Sprintf("%.0f", want.Income))
	q.Add("netWorth", fmt.Sprintf("%.0f", want.NetWorth))
	r.URL.RawQuery = q.Encode()

	if err := s.UrlQueryParam(&data, *r.URL); err != nil {
		t.Error("failed to parse request: ", err)
	} else {
		if data.Name != want.Name || *data.Address != *want.Address || data.Married != want.Married ||
			data.Score != want.Score || data.CreditScore != want.CreditScore || *data.SocialScore != *want.SocialScore ||
			data.Age != want.Age || data.BloodPressure != want.BloodPressure || *data.HoursActive != *want.HoursActive ||
			data.Income != want.Income || data.NetWorth != want.NetWorth {
			t.Error("failed to parse ")
			fmt.Println("data:", data)
			fmt.Println("want:", want)
			fmt.Println("social address", *data.Address, *want.Address)
			fmt.Println("social score", *data.SocialScore, *want.SocialScore)
			fmt.Println("social score", *data.HoursActive, *want.HoursActive)
		}
	}
}

func TestQueryParamMediumDataCustom(t *testing.T) {
	// data
	data := DataLargeCT{}

	// want
	want := DataLargeCT{
		NameValidity:      CTSValid,
		MarriedStatus:     CTBValid,
		ScoreClass:        CTI8First,
		CreditScoreClass:  CTIFirst,
		AgeClass:          CTU8First,
		HoursActiveClass:  CTUFirst,
		IncomeRateClass:   CTFFirst,
		NetWorthRateClass: CTF64First,
	}

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

	if err := s.UrlQueryParam(&data, *r.URL); err != nil {
		t.Error("failed to parse request: ", err)
	} else {
		if data != want {
			fmt.Println("data: ", data)
			fmt.Println("want: ", want)
			t.Error("failed to parse ")
		}
	}
}
