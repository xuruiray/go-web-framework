module github.com/xuruiray/go-web-framework

go 1.14

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.59.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net => github.com/golang/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200622214017-ed371f2e16b4
	golang.org/x/text => github.com/golang/text v0.3.3
	google.golang.org/appengine => github.com/golang/appengine v1.6.6
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
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	upper.io/db.v3 v3.5.7+incompatible
)
