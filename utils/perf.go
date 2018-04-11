package utils

import (
	"net/http"
	"net/http/pprof"

	log "github.com/gonethopper/libs/logs"
)

func Init(profBind []string) {
	pprofServeMux := http.NewServeMux()
	pprofServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	for _, addr := range profBind {

		go func(profaddr string) {
			http.ListenAndServe(profaddr, pprofServeMux)
			if err := http.ListenAndServe(profaddr, pprofServeMux); err != nil {
				log.Error("http.ListenAndServe(\"%s\", pprofServeMux) error(%v)", profaddr, err)
				panic(err)
			} else {
				log.Info("pprof start a http service:(%s)", profaddr)
			}
		}(addr)
	}
}
