package handler_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	handler "github.com/aimzeter/wuts/2_handler"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		name   string
		method string
		body   string
		url    string

		wantResp string
		wantCode int
	}{
		{
			name:   "valid body",
			method: "POST",
			body: `
				{
					"nik": "1234567",
					"agree": true
				}`,
			url:      "/participants",
			wantResp: "Participant with NIK 1234567 successfully registered",
			wantCode: 201,
		},
		{
			name:   "invalid method",
			method: "GET",
			body: `
				{
					"nik": "1234567",
					"agree": true
				}`,
			url:      "/participants",
			wantCode: 405,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := handler.NewServer()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(tc.method, tc.url, bytes.NewBufferString(tc.body))
			srv.ServeHTTP(w, r)

			assert.Equal(t, tc.wantCode, w.Code)
			if tc.wantResp != "" {
				assert.Equal(t, tc.wantResp, w.Body.String())
			}
		})
	}
}
