package api

type ApiServiceAdapter interface {
	DoRequestWithOptionalAuth(method, url string, withAuth bool, headers map[string]string, body interface{}) ([]byte, error)
	Get(url string, headers map[string]string) ([]byte, error)
	Post(url string, headers map[string]string, body interface{}) ([]byte, error)
	Put(url string, headers map[string]string, body interface{}) ([]byte, error)
	Delete(url string, headers map[string]string) ([]byte, error)
	GetWithAuth(url string, headers map[string]string) ([]byte, error)
	PostWithAuth(url string, headers map[string]string, body interface{}) ([]byte, error)
	PutWithAuth(url string, headers map[string]string, body interface{}) ([]byte, error)
	DeleteWithAuth(url string, headers map[string]string) ([]byte, error)
}
