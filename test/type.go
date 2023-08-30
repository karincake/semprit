package test

type CustomTypeString string
type CustomTypeBool bool
type CustomTypeInt8 int8
type CustomTypeInt int
type CustomTypeUint8 uint8
type CustomTypeUint uint
type CustomTypeFloat32 float32
type CustomTypeFloat64 float64

const CTSValid CustomTypeString = "valid"
const CTSInvalid CustomTypeString = "Invalid"
const CTBSuccess CustomTypeBool = true
const CTBFail CustomTypeBool = false
const CTI8First CustomTypeInt8 = 1
const CTI8Second CustomTypeInt8 = 2
const CTIFirst CustomTypeInt = 1
const CTISecond CustomTypeInt = 2
const CTU8First CustomTypeUint8 = 1
const CTU8Second CustomTypeUint8 = 2
const CTUFirst CustomTypeUint = 1
const CTUSecond CustomTypeUint = 2
const CTFFirst CustomTypeFloat32 = 1
const CTFSecond CustomTypeFloat32 = 2
const CTF6First CustomTypeFloat64 = 1
const CTF6Second CustomTypeFloat64 = 2

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

type DataSmallCustomTypeString struct {
	Name CustomTypeString `json:"name"`
}

type DataSmallCustomTypeBool struct {
	Married bool `json:"married"`
}

type DataSmallCustomTypeInt8 struct {
	Score int64 `json:"score"`
}

type DataSmallCustomTypeInt struct {
	CreditScore int `json:"creditScore"`
}

type DataSmallCustomTypeUint8 struct {
	Age uint8 `json:"age"`
}

type DataSmallCustomTypeUint struct {
	HoursActive int `json:"hoursActive"`
}

type DataSmallCustomTypeFloat32 struct {
	Income float32 `json:"income"`
}

type DataSmallCustomTypeFloat64 struct {
	NetWorth float64 `json:"netWorth"`
}

type DataMedium struct {
	Name        string  `json:"name"`
	Married     bool    `json:"married"`
	Score       int8    `json:"score"`
	CreditScore int     `json:"creditScore"`
	Age         uint8   `json:"age"`
	HoursActive uint    `json:"hoursActive"`
	Income      float32 `json:"income"`
	NetWorth    float64 `json:"netWorth"`
}

type DataMediumCT struct {
	Name         CustomTypeString  `json:"name"`
	Married      CustomTypeBool    `json:"married"`
	Score        CustomTypeInt8    `json:"score"`
	CreditScore  CustomTypeInt     `json:"creditScore"`
	Age          CustomTypeUint8   `json:"age"`
	HoursActive  CustomTypeUint    `json:"hoursActive"`
	IncomeRate   CustomTypeFloat32 `json:"income"`
	NetWorthRate CustomTypeFloat64 `json:"netWorth"`
}
