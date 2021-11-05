package parser

import (
	"fmt"
	"go_http/entity"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func DisplayResponse(responsePack *http.Response) {
	//// Display response status line
	fmt.Printf("\x1b[34m" + responsePack.Proto)
	fmt.Printf(" \x1b[33m" + responsePack.Status + "\n")

	//// Display response headers
	for k, v := range responsePack.Header {
		fmt.Printf("\x1b[35m" + k + ": ")
		for _, v2 := range v {
			fmt.Printf("\x1b[37m" + v2 + " ")
		}
		fmt.Println()
	}
	//// Display response body
	responBody, err := ioutil.ReadAll(responsePack.Body)
	if err != nil {
		log.Fatalln(`ioutil.ReadAll()->`, err)
	}
	fmt.Println("\n\x1b[32m" + string(responBody))
}

// Parse argv[2:] and return (headerTable, bodyTable)
// User probably input some of them:
//  `./go-http :8080 header1:a header2:b name=banana age=12` ; Accepted!
//  `./go-http :8080 header1:a header2:b name="banana, apple, cheery" age=12` ; Accepted!
//  `./go-http :8080 header1:a header2:b name#banana age#12` ; Accepted!
//  `./go-http :8080 header1:a header2:b name=banana age#12` ; Not Allowed!
//  `./go-http :8080 header1:a header2:b name=banana#apple ` ; Not Allowed!
func ParseAgrv(argv []string) (*entity.Client, error) {
	cli := entity.NewClient()

	isJsonDataFound := false
	isFormDataFound := false
	for i := 0; i < len(argv); i += 1 {
		// Found headers
		headerDelimeterIndex := strings.Index(argv[i], `:`)
		if headerDelimeterIndex != -1 {
			key := argv[i][0:headerDelimeterIndex]
			val := argv[i][headerDelimeterIndex+1:]

			cli.HeaderTable[key] = append(cli.HeaderTable[key], val)
			continue
		}

		jsonDelimeterIndex := strings.Index(argv[i], `=`)
		formDelimeterIndex := strings.Index(argv[i], `#`)
		if jsonDelimeterIndex != -1 && formDelimeterIndex != -1 {
			return nil, fmt.Errorf(`Is it a JSON data or a FORM data?`)
		}

		// Found JSON data
		if jsonDelimeterIndex != -1 && formDelimeterIndex == -1 {
			isJsonDataFound = true
			if isFormDataFound {
				return nil, fmt.Errorf(`JSON data and FORM data cannot appear at the same time`)
			}
			key := argv[i][0:jsonDelimeterIndex]
			val := argv[i][jsonDelimeterIndex+1:]
			cli.BodyTable[key] = append(cli.BodyTable[key], val)
			cli.DataType = entity.JSON_DATA
			continue
		}

		// Found FORM data
		if jsonDelimeterIndex == -1 && formDelimeterIndex != -1 {
			isFormDataFound = true
			if isJsonDataFound {
				return nil, fmt.Errorf(`JSON data and FORM data cannot appear at the same time`)
			}
			key := argv[i][0:formDelimeterIndex]
			val := argv[i][formDelimeterIndex+1:]
			cli.BodyTable[key] = append(cli.BodyTable[key], val)
			cli.DataType = entity.FORM_DATA
			continue
		}

		return nil, fmt.Errorf(`Bad data format`)
	}

	return cli, nil
}
