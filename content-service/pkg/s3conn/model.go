package s3conn

import "encoding/xml"

type ErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Error   string   `xml:"Error"`
	Message string   `xml:"Message"`
}
