package response

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}
