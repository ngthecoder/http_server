package middleware

import (
	"bytes"
	"compress/gzip"
)

func GzipString(message string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(message)); err != nil {
		return ""
	}
	if err := gz.Close(); err != nil {
		return ""
	}
	return b.String()
}
