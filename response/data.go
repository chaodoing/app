package response

import (
	`encoding/xml`
)

type Data struct {
	XMLName xml.Name    `xml:"root" json:"-"`
	Status  uint32      `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}
