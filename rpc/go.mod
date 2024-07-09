module github.com/sixwaaaay/token/rpc

go 1.21

toolchain go1.22.5

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/sixwaaaay/token v0.1.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.65.0
)

replace github.com/sixwaaaay/token => ../

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
