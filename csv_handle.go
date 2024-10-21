package data_storage

import (
	"encoding/json"
	"time"

	"github.com/gocarina/gocsv"
	rotatefile "github.com/lianyun0502/data_storage/rotate_file"
	"github.com/sirupsen/logrus"
)

func WithCsvJsonHandle[T any](log *logrus.Logger, fileName string, opts ...rotatefile.Option) func([]byte) {
	writer, err := rotatefile.New(
		fileName + "_%Y%m%d%H%M.csv",
		rotatefile.WithMaxAge(time.Duration(24*7)*time.Hour),
		rotatefile.WithRotationTime(time.Duration(1)*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return func(rawData []byte) {
		obj := new(T)
		json.Unmarshal(rawData, obj)
		if data ,err := gocsv.MarshalStringWithoutHeaders([]T{*obj}); err != nil {
			log.Error(err)
		}else{
			_, err :=writer.Write([]byte(data))
			if err != nil {
				log.Error(err)
			}
		}
	}

}

func WithCsvHandle[T any](log *logrus.Logger, fileName string, opts ...rotatefile.Option) func(T) {
	writer, err := rotatefile.New(
		fileName + "_%Y%m%d%H%M.csv",
		rotatefile.WithMaxAge(time.Duration(24*7)*time.Hour),
		rotatefile.WithRotationTime(time.Duration(1)*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return func(obj T) {
		if data ,err := gocsv.MarshalStringWithoutHeaders([]T{obj}); err != nil {
			log.Error(err)
		}else{
			_, err :=writer.Write([]byte(data))
			if err != nil {
				log.Error(err)
			}
		}
	}

}