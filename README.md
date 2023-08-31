# Semprit
Is a struct filler from various sources. Due to its nature as data parser, each process stops the moment it meets an error, and immediately return the error.

For now Semprit is able to import data from 3 sources:
- Form-Data from Net/HTTP
- Query Param from Net/Url
- IoReader with json content format (this is just wrapper around)

## Installation and Usage
Just use go get command

`go get github.com/karincake/semprit`

Import in the pakcage

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