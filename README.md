It is mainly to quickly start an HTTP service for file transmission, including file upload and download and video file playback, to get rid of the embarrassment of no tools available

![](https://github.com/lflxp/lflxp-static/blob/master/asset/httpstatic.png)

# Requirements

* go get -u github.com/jteeuwen/go-bindata/...
* go get -u github.com/elazarl/go-bindata-assetfs/...
* go get -u github.com/swaggo/swag/cmd/swag

# Install

> make

There were two step in makefile

* cp cmd/main.go ..
* go build

# Usage

`For Coder Demo`

> cmd/main.go

```go
package main

import (
	"flag"

	"github.com/lflxp/lflxp-static/pkg"
)

var (
	portHttpStatic string
	pathHttpStatic string
	isVideo        bool
	pagesize       int
	types          string
)

func init() {
	flag.StringVar(&portHttpStatic, "p", "9090", "service port")
	flag.StringVar(&pathHttpStatic, "f", "./", "load catalogs")
	flag.StringVar(&types, "t", ".avi,.wma,.rmvb,.rm,.mp4,.mov,.3gp,.mpeg,.mpg,.mpe,.m4v,.mkv,.flv,.vob,.wmv,.asf,.asx", "Filter video types, multiple separated by commas")
	flag.BoolVar(&isVideo, "v", false, "Switch to video mode or not")
	flag.IntVar(&pagesize, "c", 20, "Videos per page")
	flag.Parse()
}

func main() {
	api := pkg.Apis{
		Port:     portHttpStatic,
		Path:     pathHttpStatic,
		Types:    types,
		IsVideo:  isVideo,
		PageSize: pagesize,
	}

	err := api.Check()
	if err != nil {
		panic(err)
	}

	api.Execute()
}
```

# Running

`Format`

```bash
➜  lflxp-static git:(master) ✗ lflxp-static -h
Usage of lflxp-static:
  -c int
        Videos per page (default 20)
  -f string
        load catalogs (default "./")
  -p string
        service port (default "9090")
  -t string
        Filter video types, multiple separated by commas (default ".avi,.wma,.rmvb,.rm,.mp4,.mov,.3gp,.mpeg,.mpg,.mpe,.m4v,.mkv,.flv,.vob,.wmv,.asf,.asx")
  -v    Switch to video mode or not
```

## The parameters of the parameters

> lflxp-static -v -p 9091 -c -f /tmp

```
Simple file transfer and file presentation through local HTTP service
Usage:
showme static [flags]
Flags:
-h, --help help for static
-c. -- PageSize int number of videos per page (default 20)
-f. -- path string load directory (default ". /")
-p. -- port string service port (default "9090")
-t. -- Types string to filter video types, multiple separated by commas (default ". Avi,. Wma,. RMVB,. RM,. MP4,. MOV,. 3gp,. MPEG,. Mpg,. MPE,. M4v,. MKV,. Flv,. VOB,. WMV,. ASF,. ASX")
-v. -- whether video is switched to video mode
Global Flags:

--config string config file (default is $HOME/.showme.yaml)
```

## Function points

* Web interface operations
* Static file download
* Bulk file upload
* Video file playback
*Prometheus monitoring

## Optimization of Chinese Medicine

* Full function operation of web page
* Web page for file download
* Upload web page (no curl command operation, convenient and fast)
* View the monitoring indicators on the web page, which can be connected to the Prometheus server for monitoring
* Web video file loading and viewing function