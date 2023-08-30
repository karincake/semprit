package test

type DataSmallString struct {
	Name string `json:"name"`
}

type DataSmallBool struct {
	Married bool `json:"married"`
}

type DataSmallInt8 struct {
	Score int64 `json:"score"`
}

type DataSmallInt struct {
	CreditScore int `json:"creditScore"`
}

type DataSmallUint8 struct {
	Age uint8 `json:"age"`
}

type DataSmallUint struct {
	HoursActive int `json:"hoursActive"`
}

type DataSmallFloat32 struct {
	Income float32 `json:"income"`
}

type DataSmallFloat64 struct {
	NetWorth float64 `json:"netWorth"`
}

type DataMedium struct {
	Name        string  `json:"name"`
	Married     bool    `json:"married"`
	Score       int64   `json:"score"`
	CreditScore int     `json:"creditScore"`
	Age         uint8   `json:"age"`
	HoursActive int     `json:"hoursActive"`
	Income      float32 `json:"income"`
	NetWorth    float64 `json:"netWorth"`
}
