package logic

import "testing"

func TestVerifyEmailFormat(t *testing.T) {
	if ok := VerifyEmailFormat("144098@qq.com"); !ok {
		t.Error("error")
	}

	if ok := VerifyEmailFormat("1234@w@w.com"); ok {
		t.Error("error")
	}

	if ok := VerifyEmailFormat("1@q.c"); ok {
		t.Error("error")
	}
}

func TestVerifyMobileFormat(t *testing.T) {
	if ok := VerifyMobileFormat("17864437289"); !ok {
		t.Error("error")
	}

	if ok := VerifyMobileFormat("1786447289"); ok {
		t.Error("error")
	}

	if ok := VerifyMobileFormat("98767743567"); ok {
		t.Error("error")
	}

	if ok := VerifyMobileFormat("28767743567"); ok {
		t.Error("error")
	}
}