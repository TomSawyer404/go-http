package operator

import (
	"bytes"
	"encoding/json"
	"go_http/entity"
	"go_http/parser"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func wrapData(bodyTable map[string][]string, dataType int) *bytes.Buffer {
	if dataType == entity.JSON_DATA {
		var jsonData []byte
		var err error

		if 1 == len(bodyTable) {
			tmp := make(map[string]string)
			for k, v := range bodyTable {
				tmp[k] = v[0]
			}
			jsonData, err = json.Marshal(tmp)
		} else {
			jsonData, err = json.Marshal(bodyTable)
			if err != nil {
				log.Fatalln(`json.Marshal()->`, err)
			}
		}

		return bytes.NewBuffer([]byte(jsonData))
	}

	if dataType == entity.FORM_DATA {
		formData := make(url.Values)
		for k, v := range bodyTable {
			//formData[k] = append(formData[k], v)
			formData[k] = v
		}
		return bytes.NewBufferString(formData.Encode())
	}

	return nil
}

func setHeaders(newReq *http.Request, cli *entity.Client, bodyData *bytes.Buffer) {
	if entity.JSON_DATA == cli.DataType {
		newReq.Header.Set(`Content-Type`, `application/json`)
	} else if entity.FORM_DATA == cli.DataType {
		newReq.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
	}

	newReq.Header.Add(`User-Agent`, `banana-httpie v0.4`)
	newReq.Header.Set(`Content-Length`, strconv.Itoa(bodyData.Len()))
	if cli.HeaderTable != nil {
		for k, v := range cli.HeaderTable {
			newReq.Header[k] = v
		}
	}
}

func DoPOST(hostAndPort string, cli *entity.Client) {
	//// Serialized data
	bodyData := wrapData(cli.BodyTable, cli.DataType)
	if `:` == string(hostAndPort[0]) {
		hostAndPort = "localhost" + hostAndPort
	}

	//// Create a new request object and add my personal `User-Agent`
	newReq, err := http.NewRequest(`POST`, `http://`+hostAndPort, bodyData)
	if err != nil {
		log.Fatalln(`http.NewRequest()->`, err)
	}

	//// Set header information
	setHeaders(newReq, cli, bodyData)

	//// Send request
	responsePack, err := http.DefaultClient.Do(newReq)
	if err != nil {
		log.Fatalln(`http.DefaultClient.Do()->`, err)
	}
	defer responsePack.Body.Close()

	/// Display response information
	parser.DisplayResponse(responsePack)
}
