package ui

import "net/http"

// ===========================================================
// ==          GopherJS specific requests handlers          ==
// ===========================================================

// GopherjsScriptHandlerExt is returning the handler for requests getting the main GopherJS file.
func GopherjsScriptHandlerExt(webAppRoot string, gopherjsMainScriptFile string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, webAppRoot+gopherjsMainScriptFile)
	})
}

// GopherjsScriptMapHandlerExt is returning the handler for requests getting the main GopherJS script map file.
func GopherjsScriptMapHandlerExt(webAppRoot string, gopherjsMainScriptMapFile string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, webAppRoot+gopherjsMainScriptMapFile)
	})
}
