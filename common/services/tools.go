package services

import "net/http"

func HttpHeaderToServiceHeaders(header http.Header) map[string]*HttpHeader {
	headers := map[string]*HttpHeader{}
	for name, values := range header {
		headers[name] = &HttpHeader{Name: name, Values: values}
	}
	return headers
}


func ServiceHeadersToHttpHeader(headers map[string]*HttpHeader) http.Header {
	header := http.Header{}
	for name, httpHeader := range headers {
		for _, value := range httpHeader.Values {
			header.Add(name, value)
		}
	}
	return header
}
