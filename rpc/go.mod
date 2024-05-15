module github.com/sixwaaaay/token/rpc

go 1.20

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/sixwaaaay/token v0.1.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.64.0
)

replace github.com/sixwaaaay/token => ../

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
