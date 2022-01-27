package goleri

//Version exposes the go-thingsdb version
const Version = "0.1.1"

//Publish module:
//
//  go mod tidy
//  go test ./...
//  git tag {VERSION}
//  git push origin {VERSION}
//  GOPROXY=proxy.golang.org go list -m github.com/cesbit/goleri@{VERSION}
