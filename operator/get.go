package operator

import (
	"go_http/entity"
	"go_http/parser"
	"log"
	"net/http"
)

func addHeaders(req *http.Request, headerTable entity.Headers) {
	req.Header.Add(`User-Agent`, `banana-httpie v0.4`)
	if headerTable != nil {
		for k, v := range headerTable {
			req.Header[k] = v
		}
	}
}

func DoGET(hostAndPort string, headerTable map[string][]string) {
	if `:` == string(hostAndPort[0]) {
		hostAndPort = "localhost" + hostAndPort
	}

	//// Create a new request object and add personal header
	newReq, err := http.NewRequest(`GET`, `http://`+hostAndPort, nil)
	if err != nil {
		log.Fatalln(`http.NewRequest()->`, err)
	}

	//// Add headers
	addHeaders(newReq, headerTable)

	//// Send request
	responsePack, err := http.DefaultClient.Do(newReq)
	if err != nil {
		log.Fatalln(`http.DefaultClient.Do()->`, err)
	}
	defer responsePack.Body.Close()

	/// Display response information
	parser.DisplayResponse(responsePack)
}
