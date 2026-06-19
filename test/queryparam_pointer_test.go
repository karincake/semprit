package test

import (
	"net/http"
	"testing"

	s "github.com/karincake/semprit"
)

// DataPointerCT exercises pointer fields whose element type is a named/custom
// type, plus a pointer to a base float (the two paths that previously panicked).
type DataPointerCT struct {
	NameValidity *CustomTypeString `json:"nameValidity"`
	ScoreClass   *CustomTypeInt8   `json:"scoreClass"`
	AgeClass     *CustomTypeUint8  `json:"ageClass"`
	Income       *float64          `json:"income"`
}

func TestQueryParamPointerCustomType(t *testing.T) {
	data := DataPointerCT{}

	r, _ := http.NewRequest("GET", "/", nil)
	q := r.URL.Query()
	q.Add("nameValidity", "valid")
	q.Add("scoreClass", "1")
	q.Add("ageClass", "1")
	q.Add("income", "1.5")
	r.URL.RawQuery = q.Encode()

	if err := s.UrlQueryParam(&data, *r.URL); err != nil {
		t.Fatal("failed to parse request: ", err)
	}

	if data.NameValidity == nil || *data.NameValidity != CTSValid {
		t.Errorf("nameValidity not filled: %v", data.NameValidity)
	}
	if data.ScoreClass == nil || *data.ScoreClass != CTI8First {
		t.Errorf("scoreClass not filled: %v", data.ScoreClass)
	}
	if data.AgeClass == nil || *data.AgeClass != CTU8First {
		t.Errorf("ageClass not filled: %v", data.AgeClass)
	}
	if data.Income == nil || *data.Income != 1.5 {
		t.Errorf("income not filled: %v", data.Income)
	}
}
