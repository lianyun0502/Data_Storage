package data_storage_test

import (
	"github.com/lianyun0502/data_storage"
	"github.com/sirupsen/logrus"

	"testing"
)

var logger = logrus.New()

type TestStruct struct {
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Gender string `json Gender`
}

func TestWithCsvHandle(t *testing.T) {
	handle := data_storage.WithCsvHandle[TestStruct](logger, "test")
	handle([]byte(`{"Name":"test","Age":1, "Gender":"man"}`))
	handle([]byte(`{"Name":"test","Age":2}`))
	handle([]byte(`{"Name":"test","Age":3}`))

	handle = data_storage.WithCsvHandle[TestStruct](logger, "test2")
	handle([]byte(`{"Name":"test2","Age":2}`))
}
