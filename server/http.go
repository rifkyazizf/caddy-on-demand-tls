package server

import (
	"caddy-on-demand-tls/pkg/routy"
	"net/http"
	"os"
	"slices"
	"strings"
)

var allowedDomains []string

func init() {
	allowedDomainEnv := os.Getenv("ALLOWED_DOMAINS")

	allowedDomains = strings.Split(allowedDomainEnv, ",")
}

func RunHttp() *http.Server {
	r := routy.NewRouter()

	r.Get("/", handleOnDemandTls)
	server := &http.Server{
		Addr:    ":5555",
		Handler: r,
	}

	return server
}

func handleOnDemandTls(w http.ResponseWriter, r *http.Request) {
	if len(allowedDomains) == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

	domain := r.URL.Query().Get("domain")
	if !slices.Contains(allowedDomains, domain) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("NOT OK"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
