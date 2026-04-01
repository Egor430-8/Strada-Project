package validation

import "testing"

func TestIsValidTitle(t *testing.T) {
	title := "lol"
	ok := IsValidTitle(title)
	if !ok {
		t.Errorf("Ожидалось true, получено %v", ok)
	}

	title = "12"
	ok = IsValidTitle(title)
	if ok {
		t.Errorf("Ожидалось false, получено %v", ok)
	}
}
