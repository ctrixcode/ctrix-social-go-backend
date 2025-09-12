package healthcheck

import "net/http"

// Healthz is a simple handler to check the service status.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ctrix Social Backend!"))
}
