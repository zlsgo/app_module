//go:build !nostatic
// +build !nostatic

package html

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sohaha/zlsgo/znet"
)

func TestRegisterStaticDefaultPrefix(t *testing.T) {
	e := znet.New("test")
	if err := registerStatic(e, ""); err != nil {
		t.Fatalf("registerStatic: %v", err)
	}
	for _, name := range []string{"zcss.js", "zview.js"} {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/__static_html/"+name, nil)
		e.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("GET %s code=%d", name, w.Code)
		}
		var got []byte
		for i := range staticFiles {
			if staticFiles[i].Name == name {
				got = staticFiles[i].Data
			}
		}
		body, _ := io.ReadAll(w.Result().Body)
		if string(body) != string(got) {
			t.Fatalf("%s content mismatch: got %d bytes want %d", name, len(body), len(got))
		}
	}
}

func TestRegisterStaticCustomPrefix(t *testing.T) {
	e := znet.New("test-custom")
	if err := registerStatic(e, "/assets"); err != nil {
		t.Fatalf("registerStatic: %v", err)
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/assets/zcss.js", nil)
	e.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("custom prefix GET code=%d", w.Code)
	}
}
