package adapter

import "net/http"

const contentTypeHeaderKey = "Content-Type"

type contentTypeAdapter struct {
	contentType string
}

func NewContentTypeAdapter(contentType string) Adapter {
	return contentTypeAdapter{
		contentType: contentType,
	}
}

func (a contentTypeAdapter) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set(contentTypeHeaderKey, a.contentType)
		next.ServeHTTP(rw, r)
	})
}
