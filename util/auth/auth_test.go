package auth

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/stretchr/testify/require"
)

func TestAuthMw_CheckAuth(t *testing.T) {
	middleware := New(util.Config{
		SecretKey: "secret",
	})
	type wantStruct struct {
		code    int
		message string
		user    model.User
	}
	tests := []struct {
		name     string
		mockNext func() http.HandlerFunc
		args     func() *http.Request
		want     wantStruct
	}{
		{
			name: "Normal",
			mockNext: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write([]byte(`OK`))
					require.NoError(t, err)
				}
			},
			args: func() *http.Request {
				req, err := http.NewRequest("GET", "/get/data", nil)
				require.NoError(t, err)
				req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6IjkyYmZiOTgwLTI0ZGUtNDI1Zi1hYjY0LTI3N2I2MWE3NGQ1ZCIsImlhdCI6MTYwMDYxODQ0OCwiZXhwIjoxNjAwNjIyMTgzLCJ1c2VyIjp7ImlkIjoxLCJ1c2VybmFtZSI6ImFsbWFuYWxmYXJ1cSIsImZ1bGxfbmFtZSI6IkFsbWFudGVyYSIsInJvbGVfaWQiOjF9fQ.pgVk8sYEWOJVkLbPDS9tX_D9NryFPi41jdpcCr43TP4")
				return req
			},
			want: wantStruct{
				code:    http.StatusOK,
				message: "OK",
				user: model.User{
					Template: model.Template{
						ID: 1,
					},
					Username: "almanalfaruq",
					FullName: "Almantera",
					RoleID:   1,
				},
			},
		},
		{
			name: "Error-CannotParse",
			mockNext: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			},
			args: func() *http.Request {
				req, err := http.NewRequest("GET", "/get/data", nil)
				require.NoError(t, err)
				req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MDAwMDMyMTMsImV4cCI6MzI1MjU3NjAwMTMsImF1ZCI6Ind3dy5leGFtcGxlLmNvbSIsInN1YiI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJHaXZlbk5hbWUiOiJKb2hubnkiLCJTdXJuYW1lIjoiUm9ja2V0IiwiRW1haWwiOiJqcm9ja2V0QGV4YW1wbGUuY29tIiwiUm9sZSI6WyJNYW5hZ2VyIiwiUHJvamVjdCBBZG1pbmlzdHJhdG9yIl19.jwwkkb6B1O8dZeQL9s2zKuUAX9B1jzc6eMZxxSrC8MQ")
				return req
			},
			want: wantStruct{
				code:    http.StatusUnauthorized,
				message: "{\"code\":401,\"data\":null,\"message\":\"User data is null\"}\n",
			},
		},
		{
			name: "Error Token Expired",
			mockNext: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			},
			args: func() *http.Request {
				req, err := http.NewRequest("GET", "/get/data", nil)
				require.NoError(t, err)
				req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1OTk5OTU0NDksImV4cCI6MTU5OTkwOTA0OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.WCozN6eF10ODuE8eAP7o6fHhDGY6VgW5Q7X61Phhsd0")
				return req
			},
			want: wantStruct{
				code:    http.StatusUnauthorized,
				message: "{\"code\":401,\"data\":null,\"message\":\"token is expired by 197h15m34s\"}\n",
			},
		},
		{
			name: "Error Parse Token Unsigned",
			mockNext: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			},
			args: func() *http.Request {
				req, err := http.NewRequest("GET", "/get/data", nil)
				require.NoError(t, err)
				req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE1OTk5OTU0NDksImV4cCI6MTYzMTUzMTQ0OSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.Zy0u8wMPipBGNdtoqT-SsH9hFKGPNr-gf5RViwV9UyA")
				return req
			},
			want: wantStruct{
				code:    http.StatusUnauthorized,
				message: "{\"code\":401,\"data\":null,\"message\":\"signature is invalid\"}\n",
			},
		},
		{
			name: "Error Not Bearer Token",
			mockNext: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			},
			args: func() *http.Request {
				req, err := http.NewRequest("GET", "/get/data", nil)
				require.NoError(t, err)
				req.Header.Add("Authorization", "NotBearer secret")
				return req
			},
			want: wantStruct{
				code:    http.StatusUnauthorized,
				message: "{\"code\":401,\"data\":null,\"message\":\"token contains an invalid number of segments\"}\n",
			},
		},
		{
			name: "Error Missing Header",
			mockNext: func() http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
				}
			},
			args: func() *http.Request {
				req, err := http.NewRequest("GET", "/get/data", nil)
				require.NoError(t, err)
				return req
			},
			want: wantStruct{
				code:    http.StatusUnauthorized,
				message: "{\"code\":401,\"data\":null,\"message\":\"Missing authorization header\"}\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := middleware.CheckAuth(tt.mockNext())
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, tt.args())
			require.Equal(t, tt.want.code, rr.Code)
			byt, err := ioutil.ReadAll(rr.Result().Body)
			require.NoError(t, err)
			defer rr.Result().Body.Close()
			if tt.name != "Error Token Expired" {
				require.Equal(t, tt.want.message, string(byt))
			}
			if user, ok := tt.args().Context().Value(model.CTX_USER).(model.User); ok {
				require.Equal(t, tt.want.user, user)
			}
		})
	}
}
