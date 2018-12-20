package handlers

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/vision8tech/goings/common"
)

// TemplateBundleHandler returns the handler for POST "/template-bundle" request.
func TemplateBundleHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var templateContentItemsBuffer bytes.Buffer
		enc := gob.NewEncoder(&templateContentItemsBuffer)
		items := env.TemplateSet.Bundle().Items()
		err := enc.Encode(&items)
		if err != nil {
			log.Print("(handlers) TemplateBundleHandler > Encoding error: ", err)
		}
		w.Header().Set("content-type", "application/octet-stream")
		w.Write(templateContentItemsBuffer.Bytes())

	})

}
