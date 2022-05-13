package handlers

import (
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
	"log" //"github.com/golang/glog"
)

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit, release string) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/version", version(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/whoami", whoami()).Methods("GET")
	r.HandleFunc("/info", info(release, time.now())).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
