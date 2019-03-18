module github.com/xuruiray/go-web-framework

replace (
    cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go latest
    golang.org/x/crypto => github.com/golang/crypto latest
    golang.org/x/net => github.com/golang/net latest
    golang.org/x/sync => github.com/golang/sync latest
    golang.org/x/sys => github.com/golang/sys latest
    golang.org/x/text => github.com/golang/text latest
    google.golang.org/appengine => github.com/golang/appengine latest
)

require (
	github.com/elazarl/goproxy v0.0.0-20181111060418-2ce16c963a8a // indirect
	github.com/gin-gonic/gin v1.3.0 // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/juju/ratelimit v1.0.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/parnurzeal/gorequest v0.2.15
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sirupsen/logrus v1.4.0
	github.com/smartystreets/goconvey v0.0.0-20190306220146-200a235640ff // indirect
	github.com/ugorji/go/codec v0.0.0-20190316192920-e2bddce071ad // indirect
	github.com/xuruiray/binding v0.0.0-20190311123139-d6192edfdcae
	github.com/xuruiray/gosql v0.0.0-20180224093734-b39d511378f3
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	upper.io/db.v3 v3.5.7+incompatible
)
