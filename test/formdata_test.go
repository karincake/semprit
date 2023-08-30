package test

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	s "github.com/karincake/semprit"
)

func TestMediumDataNormal(t *testing.T) {
	// data and wanted
	dataAddress := ""
	dataSocialScore := 0
	dataHoursActive := uint(0)
	data := DataMedium{
		Address:     &dataAddress,
		SocialScore: &dataSocialScore,
		HoursActive: &dataHoursActive,
	}
	wantAddress := "JL Localhost 2023"
	wantSocialScore := 25000
	wantHoursActive := uint(30000)
	want := DataMedium{
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

	// mock request, make it fast import from "want"
	r, _ := http.NewRequest("POST", "/", nil)
	r.PostForm = make(url.Values)
	r.PostForm.Add("name", want.Name)
	r.PostForm.Add("address", *want.Address)
	r.PostForm.Add("married", "true")
	r.PostForm.Add("score", strconv.Itoa(int(want.Score)))
	r.PostForm.Add("creditScore", strconv.Itoa(int(want.CreditScore)))
	r.PostForm.Add("socialScore", strconv.Itoa(int(*want.SocialScore)))
	r.PostForm.Add("age", strconv.Itoa(int(want.Age)))
	r.PostForm.Add("bloodPressure", strconv.Itoa(int(want.BloodPressure)))
	r.PostForm.Add("hoursActive", strconv.Itoa(int(*want.HoursActive)))
	r.PostForm.Add("income", fmt.Sprintf("%.0f", want.Income))
	r.PostForm.Add("netWorth", fmt.Sprintf("%.0f", want.NetWorth))

	if err := s.HttpFormData(&data, r); err != nil {
		t.Error("failed to parse request: ", err)
	} else {
		if data.Name != want.Name || *data.Address != *want.Address || data.Married != want.Married ||
			data.Score != want.Score || data.CreditScore != want.CreditScore || *data.SocialScore != *want.SocialScore ||
			data.Age != want.Age || data.BloodPressure != want.BloodPressure || *data.HoursActive != *want.HoursActive ||
			data.Income != want.Income || data.NetWorth != want.NetWorth {
			t.Error("failed to parse ")
			fmt.Println("data:", data)
			fmt.Println("want:", want)
			fmt.Println("social addess", *data.Address, *want.Address)
			fmt.Println("social score", *data.SocialScore, *want.SocialScore)
			fmt.Println("social score", *data.HoursActive, *want.HoursActive)
		}
	}
}

func TestMediumDataCustom(t *testing.T) {
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
	data := DataMediumCT{}
	want := DataMediumCT{
		NameValidity:      CTSValid,
		MarriedStatus:     CTBValid,
		ScoreClass:        CTI8First,
		CreditScoreClass:  CTIFirst,
		AgeClass:          CTU8First,
		HoursActiveClass:  CTUFirst,
		IncomeRateClass:   CTFFirst,
		NetWorthRateClass: CTF64First,
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
