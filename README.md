go-copy
=======

Copy (http://copy.com) service library for Go lang


Description
===========

Copy is a new service like dropbox, where you can sync your files in the cloud
with his application, web interface or REST API. 

The aim of this library is to support all or almost all the API calls that copy
exposes to the public.

This library is written in Go.

| Status        | Tests                                                                                                                   | Coverage                                                                                                                                 | Docs
| :-------------: |:-----------------------------------------------------------------------------------------------------------------------:| ----------------------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------:|
| Development   | [![Build Status](https://drone.io/github.com/slok/go-copy/status.png)](https://drone.io/github.com/slok/go-copy/latest) | [![Coverage Status](https://coveralls.io/repos/slok/go-copy/badge.png?branch=master)](https://coveralls.io/r/slok/go-copy?branch=master) | [![GoDoc](http://godoc.org/github.com/slok/go-copy/copy?status.png)](http://godoc.org/github.com/slok/go-copy/copy)



API implementation status
-------------------------

* User
    * ~~Get User data~~
    * ~~Update User Profile~~

* Files
    * ~~Get root path meta~~
    * ~~Get path meta~~
    * ~~Get File revisions meta~~
        * Tested in sandbox (Copy API fails for now, can't test it in prod)
    * ~~Get concrete file revision meta~~
        * Tested in sandbox (Copy API fails for now, can't test it in prod)
    * ~~Get file data~~
    * ~~Delete file~~
    * ~~Update file~~
        * Tested in sandbox (Copy API fails for now, can't test it in prod)
    * ~~Rename file~~
    * ~~Move file~~
    * ~~Create dir~~
    * Upload file data
        * ~~At once (Warning, in memory)~~
        * Chunked (Not possible for now, see API docs)
    * ~~Get thumbnail~~
        * Tested in sandbox (Copy API fails for now, can't test it in prod)

* Links
    * Get link information
    * Get all user links
    * Create a link
    * Update a link
    * Delete a link
    * Get meta of files attached to a link

How to use it
-------------

```go
import (
    "github.com/slok/go-copy/copy"
    "io/ioutil"
)

client, _ := copy.NewDefaultClient(appToken, appSecret, accessToken, accessSecret)
fs := copy.NewFileService(client)

filePath := "photos/awesome.png"
r, _ := fs.GetFile(filePath)
fileBytes, _ := ioutil.ReadAll(r)

err = ioutil.WriteFile("awesome.png", fileBytes, 0644)
if err != nil {
    panic(err)
}
```

License
=======

See LICENSE file

[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/slok/go-copy/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

