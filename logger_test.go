package loggergo

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServiceNameLogged(t *testing.T) {
	serviceName := "abc"
	realm := ""

	bufln := &bytes.Buffer{}
	logrus.SetOutput(bufln)

	logger := InitLogger(serviceName, realm)

	logger.Errorf("err")
	assert.Contains(t, bufln.String(), fmt.Sprintf("service=%s", serviceName))
}

func TestRealmLogged(t *testing.T) {
	serviceName := ""
	realm := "abc"

	bufln := &bytes.Buffer{}
	logrus.SetOutput(bufln)

	logger := InitLogger(serviceName, realm)

	logger.Errorf("err")
	assert.Contains(t, bufln.String(), fmt.Sprintf("realm=%s", realm))
}
