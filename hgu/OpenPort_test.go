package hgu

import (
	"reflect"
	"testing"
)

func TestOpenPort(t *testing.T) {
	record := "40,test-1,TCP,192.168.1.165,12345,12346,12390,T,ppp0.1"

	parse, err := Parse(record)
	if err != nil {
		t.Error(err)
	}

	expected := OpenPort{
		Id:                40,
		Name:              "test-1",
		Protocol:          TCP,
		Address:           "192.168.1.165",
		ExternalPortStart: 12345,
		ExternalPortEnd:   12346,
		InternalPortStart: 12390,
		Enabled:           true,
		Interface:         "ppp0.1",
	}

	if !reflect.DeepEqual(parse, expected) {
		t.Error("Parse result doesn't match expectations.")
	}
}

func TestDisabledOpenPort(t *testing.T) {
	record := "40,test-1,TCP,192.168.1.165,12345,12346,12390,F,ppp0.1"

	parse, err := Parse(record)
	if err != nil {
		t.Error(err)
	}

	expected := OpenPort{
		Id:                40,
		Name:              "test-1",
		Protocol:          TCP,
		Address:           "192.168.1.165",
		ExternalPortStart: 12345,
		ExternalPortEnd:   12346,
		InternalPortStart: 12390,
		Enabled:           false,
		Interface:         "ppp0.1",
	}

	if !reflect.DeepEqual(parse, expected) {
		t.Error("Parse result doesn't match expectations.")
	}
}

func TestNoRangeOpenPort(t *testing.T) {
	record := "40,test-1,TCP,192.168.1.165,12345,0,12390,T,ppp0.1"

	parse, err := Parse(record)
	if err != nil {
		t.Error(err)
	}

	expected := OpenPort{
		Id:                40,
		Name:              "test-1",
		Protocol:          TCP,
		Address:           "192.168.1.165",
		ExternalPortStart: 12345,
		ExternalPortEnd:   0,
		InternalPortStart: 12390,
		Enabled:           true,
		Interface:         "ppp0.1",
	}

	if !reflect.DeepEqual(parse, expected) {
		t.Error("Parse result doesn't match expectations.")
	}
}
