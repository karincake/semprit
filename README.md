# Semprit
Is a struct filler from various sources. 

For now Semprit is able to import data from 3 sources:
- Form-Data from Net/HTTP
- Query Param from Net/Url
- IoReader with json content format, this is just wrapper around for json.Unmarshal functtion. Due to json.Unmarshal limitations, any errors error occured will only show the first error.

## Installation and Usage
Get the package

`go get github.com/karincake/semprit`

Import the pakcage

`import "github.com/karincake/semprit"`

Call the function according to the need

```
myData := mySruct{}
err := semprit.HttpFormData(&myData, r) // assuming r is http.Request pointer
if err != nil {
    // do something with error
}
```

The other function that can be used are :
- `UrlQueryParam(&struct, url.URL)`
- `IOReaderJson(&struct, io.Reader)`