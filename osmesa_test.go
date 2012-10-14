package osmesa

import "testing"
import "unsafe"

func TestOSMesa(t *testing.T) {

	ctx := CreateContext(RGBA, nil)
	if ctx == nil {
		t.Fatal("unable to create context")
		t.FailNow()
	}

	if GetCurrentContext() != nil {
		t.Fatal("current context is not nil")
		t.FailNow()
	}

	buffer := make([]byte, 640*480*4)

	if !MakeCurrent(ctx, unsafe.Pointer(&buffer[0]), UNSIGNED_BYTE, 640, 480) {
		t.Fatal("MakeCurrent failed")
		t.FailNow()
	}

	if GetCurrentContext() == nil {
		t.Fatal("current context is nil")
		t.FailNow()
	}

	if GetCurrentContext() != ctx {
		t.Fatal("invalid current context")
		t.FailNow()
	}

	if GetIntegerv(WIDTH) != 640 {
		t.Fatal("invalid width")
		t.FailNow()
	}

	if GetIntegerv(WIDTH) != 640 {
		t.Fatal("invalid height")
		t.FailNow()
	}

	DestroyContext(ctx)

}
