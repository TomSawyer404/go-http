package entity

const (
	NO_DATA = iota
	JSON_DATA
	FORM_DATA
)

type (
	Client struct {
		DataType    int
		HeaderTable Headers
		BodyTable   Bodies
	}

	Headers map[string][]string
	Bodies  map[string][]string
)

func NewClient() *Client {
	client := Client{
		DataType:    NO_DATA,
		HeaderTable: Headers{},
		BodyTable:   Bodies{},
	}

	return &client
}
