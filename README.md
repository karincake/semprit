# Semprit - Struct Filler
Package to help in filling struct with data from various sources. 

## Installation and Usage
Get the package

`go get github.com/karincake/semprit`

Import the pakcage

`import "github.com/karincake/semprit"`

Call the function according to the need

```go
myData := mySruct{}
err := semprit.HttpFormData(&myData, r) // assuming r is http.Request pointer
if err != nil {
    // do something with error
}
```

## Available Functions
There are 3 functions available for filling struct
- `HttpFormData(&struct, *http.Request)`, to fill struct with data from http.Request form-data.
- `UrlQueryParam(&struct, url.URL)`, to fill struct with data from url query parameter
- `IOReaderJson(&struct, io.Reader)`, to fill struct with data from io.Reader that having json encoding content
