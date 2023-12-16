module github.com/sixwaaaay/token/rpc

go 1.20

require (
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/sixwaaaay/token v0.0.0-20231216070301-25cf4b3b334f
	github.com/stretchr/testify v1.8.4
	google.golang.org/grpc v1.60.0
)

replace github.com/sixwaaaay/token => ../

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
