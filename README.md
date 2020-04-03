[![Build Status](https://travis-ci.com/AccelByte/logger-go.svg?branch=master)](https://travis-ci.com/AccelByte/logger-go)

# logger-go
Accelbyte logger library

# usage
```go
	logger := InitLogger(serviceName, realm)

    logger.Errorf("err")
    // output: time="2020-04-03T16:37:45+07:00" level=error msg=err file="logger_test.go:21" func=logger-go.TestServiceNameLogged realm=def service=abc

```