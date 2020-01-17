package provider

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_remoteCaller_Call(t *testing.T) {

	tests := []struct {
		name       string
		response   string
		want       interface{}
		statusCode int
		errmsg     string
	}{
		{
			name:       "Success",
			response:   `{"status":"OK","message":"Success"}`,
			statusCode: 200,
			want:       `{"status":"OK","message":"Success"}`,
		},
		{
			name:       "Error",
			response:   `{"status":"FAIL","message":"Error"}`,
			statusCode: 500,
			want:       nil,
			errmsg:     "error from gateway",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			c := NewProvider(mockserver.URL)
			if got, err := c.Call(); got != tt.want && err != nil && err.Error() != tt.errmsg {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}
