package test

import (
	_ "Beego/routers"
	"path/filepath"
	"runtime"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {

}
