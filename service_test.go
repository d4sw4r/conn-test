package main

import (
	"testing"
)

func TestIcmpDNS(t *testing.T) {
	err := IcmpTest("google.com")
	if err != nil {
		t.Error(err)
	}

}

func TestIcmpIP(t *testing.T) {
	err := IcmpTest("8.8.8.8")
	if err != nil {
		t.Error(err)
	}

}

func TestHttpURL(t *testing.T) {
	err := HttpTest("https://google.com")
	if err != nil {
		t.Error(err)
	}

}

func TestHttpFail(t *testing.T) {
	err := HttpTest("google.com")
	if err != nil {
		return
	}
	t.Error("should be missing protocol")
}

func TestPing(t *testing.T) {
	err := IcmpTest("8.8.8.8")
	if err != nil {
		t.Error(err)
	}
}
