## Devla go utils
- Devla Go utils is a collection of go utils to help fasten development, it is a work in progress and will be updated as needed.

### Install
```bash
go get github.com\vongkeo\go-lib
```

### Test
```bash
# to update go.mod
go mod tidy
# to run all tests
go test -v
# to run all tests in sub packages
go test -v ./...
# to run all tests in sub packages with coverage
go test -cover ./...
```
### Update to go package
```bash
git add .
git commit -m "update to v0.1.4"
git push origin
git tag v0.1.4
git push origin v0.1.4
GOPROXY=proxy.golang.org go list -m github.com/vongkeo/go-lib@v0.1.4
go list -m
```