# Intention
Visualizes basic rules of staticcheck linter, such as `U1000`, `SA4018`, `SA1028`, etc.

# Prerequisities
```bash
go get -u honnef.co/go/tools/cmd/staticcheck
```


# Execution
Makefile contains only one `target`:
```bash
make static-check
```

# References

+ [Official Doc](https://staticcheck.io/docs)

+ [Test Data](https://github.com/dominikh/go-tools/tree/master/staticcheck/testdata/src)

+ [Medium Story](https://medium.com/wesionary-team/introduction-to-linters-in-go-and-know-about-golangci-lint-6486fb2864b3)

+ [Prometheus](https://chromium.googlesource.com/external/github.com/prometheus/client_golang/+/cb063c2bf24a8bbb45f88164ea26fb027563d36b/Makefile.common)

+ [Withnic](https://blog.withnic.net/2018/09/golang%E3%81%A7%E3%81%AE%E9%96%8B%E7%99%BA%E3%81%AB%E4%BD%BF%E3%81%86%E8%87%AA%E5%88%86%E3%83%86%E3%83%B3%E3%83%97%E3%83%AC%E3%83%BC%E3%83%88makefile2018/)