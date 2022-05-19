package middleware

import (
	"fmt"
	"github.com/llatrelle/toolkit/api/app"
	"net/http"
	"os"
	"runtime/debug"
)

// The original work was derived from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middleware

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func RecovererMDW(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
				debug.PrintStack()
				app.Error(w, rvr.(error).Error())
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
