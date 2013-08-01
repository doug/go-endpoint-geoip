GeoIP Endpoints Basic Example

```

go get github.com/doug/go-endpoint-geoip/geoip

```

```yaml

application: your-app-id
version: master
runtime: go
api_version: go1

handlers:
  - url: /_ah/spi/.*
    script: _go_app

  # this links everything else in your client/app folder to be served
  # Change this to something else if your public folder is not client/app
  - url: (/?.*)/
    static_files: client/app\1/index.html
    upload: client/app(.*)/index.html
  - url: /
    static_dir: client/app

```

```go

package app

import (
        "github.com/crhym3/go-endpoints/endpoints"
        _ "github.com/doug/go-endpoint-geoip/geoip"
)

func init() {
        endpoints.HandleHttp()
}


```
