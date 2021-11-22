package ipify

import (
	"testing"
)

func Test_Ip(t *testing.T) {
	t.Log("Testing both IPV4 and IPV6 resolution within the Ip struct type.")
	dip := DefaultIp()
	err := dip.Retrieve()
	if err != nil {
		t.Fatal(err)
	}
	if dip.Ip4 == "" && dip.Ip6 == "" {
		t.Fatal("The IP address retrieval through the struct was reported as successful, but no IP addresses were retrieved.\nResults:\n", dip)
	}
	t.Log("IP address retrieval within the Ip struct has succeeded. Result below.")
	t.Log(dip)
}
