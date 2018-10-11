package fake

import (
	"testing"
)

func TestUID(t *testing.T) {
	if len(KSUID()) != 27 {
		t.Errorf("KSUID failed. expect length 27 got %d", len(KSUID()))
	}
	if len(UUID()) != 32 {
		t.Errorf("UUID failed. expect length 32 got %d", len(UUID()))
	}
	if len(OpenID()) != 28 {
		t.Errorf("OpenID failed. expect length 28 got %d", len(OpenID()))
	}

}
