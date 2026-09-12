package pipeline

import (
	"code.gitea.io/gitea/modules/log"
	"fmt"
	"net/http"
	"net/http/httputil"
)

//func DumpRequest(request *http.Request) {
//	command, _ := http2curl.GetCurlCommand(request)
//	log.Info("Redirect REQUEST:\n%s", command)
//}

func DumpRequest(request *http.Request) {
	dump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		fmt.Printf("DumpRequestOut error: %+v", err)
		return
	}
	log.Info("Redirect REQUEST: %s", string(dump))
}

func DumpResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("DumpRequestOut error: %+v", err)
		return
	}
	if len(dump) > 256 {
		dump = dump[:256]
	}
	log.Info("Redirect RESPONSE: %s", string(dump))
}
