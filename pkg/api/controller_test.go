package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jalilbengoufa/pixicoreAPI/pkg/config"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"
)

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func TestGetServers(t *testing.T) {
	myConfigFile, err := config.InitConfig()
	if err != nil {
		t.Errorf("Error from InitConfig come normally from a broken configFile. TODO : Test env. must be sandboxed. Then config must be mocked.")
	}

	controller := InitController(myConfigFile)
	r := GetRouter(controller)

	req, _ := http.NewRequest("GET", "/v1/boot/00:00:00:00:00:00", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)

		pageOK := err == nil && strings.Index(string(p), `{file:///home/cedille/coreos_production_pxe.vmlinuz [file:///home/cedille/coreos_production_pxe_image.cpio.gz] coreos.autologin coreos.first_boot=1 coreos.config.url={{ URL \"file:///home/cedille/pxe-config.ign\" }}}"`) > 0

		return statusOK && pageOK
	})

}
