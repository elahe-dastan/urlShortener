package middleware

import (
	"github.com/elahe-dastan/urlShortener_KGS/config"
	"github.com/felixge/httpsnoop"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var configuration config.LogFile

func SetConfig(constants config.Constants)  {
	configuration = constants.Log
}

// LogReqInfo describes info about HTTP request
type HTTPReqInfo struct {
	// GET etc.
	method string
	uri string
	referer string
	ipaddr string
	// response code, like 200, 404
	code int
	// number of bytes of the response sent
	size int64
	// how long did it take to
	duration time.Duration
	userAgent string
	err	string
}

func LogRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ri := &HTTPReqInfo{
			method: r.Method,
			uri: r.URL.String(),
			referer: r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
		}

		ri.ipaddr = requestGetRemoteAddress(r)

		// this runs handler h and captures information about
		// HTTP request
		m := httpsnoop.CaptureMetrics(h, w, r)

		ri.code = m.Code
		ri.size = m.Written
		ri.duration = m.Duration
		ri.err = w.Header().Get("err")
		write(ri)
	}
	return http.HandlerFunc(fn)
}

func write(ri *HTTPReqInfo)  {
	f, err := os.OpenFile(configuration.Address, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "request", log.LstdFlags)
	logger.Println(ri)
}

// requestGetRemoteAddress returns ip address of the client making the request,
// taking into account http proxies
func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIP
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}