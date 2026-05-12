package ipfilter

import "testing"

func TestIsValidSingleIPv4(t *testing.T) {
	if !IsValid("192.168.1.1") {
		t.Fail()
	}
}

func TestIsValidCIDR(t *testing.T) {
	if !IsValid("10.0.0.0/8") {
		t.Fail()
	}
}

func TestIsValidRange(t *testing.T) {
	if !IsValid("192.168.1.1-192.168.1.10") {
		t.Fail()
	}
}

func TestInvalidIP(t *testing.T) {
	if IsValid("bad-ip") {
		t.Fail()
	}
}

func TestInvalidRange(t *testing.T) {
	if IsValid("192.168.1.10-192.168.1.1") {
		t.Fail()
	}
}

func TestMatchSingleIPv4(t *testing.T) {
	if !Match("192.168.1.1", "192.168.1.1") {
		t.Fail()
	}
}

func TestNotMatchSingleIPv4(t *testing.T) {
	if Match("192.168.1.1", "192.168.1.2") {
		t.Fail()
	}
}

func TestMatchCIDR(t *testing.T) {
	if !Match("10.0.0.0/8", "10.1.2.3") {
		t.Fail()
	}
}

func TestNotMatchCIDR(t *testing.T) {
	if Match("10.0.0.0/8", "192.168.1.1") {
		t.Fail()
	}
}

func TestMatchRange(t *testing.T) {
	if !Match("192.168.1.1-192.168.1.10", "192.168.1.5") {
		t.Fail()
	}
}

func TestNotMatchRange(t *testing.T) {
	if Match("192.168.1.1-192.168.1.10", "192.168.1.20") {
		t.Fail()
	}
}
