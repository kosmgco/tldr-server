package tools

type HttpMethod uint8

const (
	HTTP_METHOD__GET HttpMethod = iota
	HTTP_METHOD__POST
	HTTP_METHOD__PUT
	HTTP_METHOD__DELETE
	HTTP_METHOD__HEAD
	HTTP_METHOD__PATCH
)

func (h HttpMethod) String() string {
	switch h {
	case HTTP_METHOD__DELETE:
		return "DELETE"
	case HTTP_METHOD__GET:
		return "GET"
	case HTTP_METHOD__HEAD:
		return "HEAD"
	case HTTP_METHOD__PATCH:
		return "PATCH"
	case HTTP_METHOD__POST:
		return "POST"
	case HTTP_METHOD__PUT:
		return "PUT"
	default:
		return ""
	}
}
