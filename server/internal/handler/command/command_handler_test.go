package command

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"pstgrprof/server/internal/entity"
	mock "pstgrprof/server/internal/handler/command/mocks"
	"strings"
	"testing"
)

// mock gin context
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	return ctx
}

func TestHandler_CreateCommand(t *testing.T) {
	type want struct {
		res    string
		status int
	}
	type fields struct {
		Service *mock.MockService
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    want
		prepare func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context)
	}{
		{
			name:    "test_1, StatusOK",
			wantErr: false,
			want: want{
				res:    "{\"id\":\"1\",\"script\":\"ls\",\"description\":\"look files\"}",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/command", strings.NewReader(`{"script": "ls", "description": "look files"}`))
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().CreateCommand(c.Request.Context(), &entity.CreateCommandReq{
					Script:      "ls",
					Description: "look files",
				}).Return(&entity.CreateCommandRes{
					ID:          "1",
					Script:      "ls",
					Description: "look files",
				}, nil)
			},
		}, {
			name:    "test_2, StatusBadRequest",
			wantErr: false,
			want: want{
				res:    "{\"error\":\"script must not be empty\"}",
				status: http.StatusBadRequest,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/command", strings.NewReader(`{"script": "", "description": "look files"}`))
				c.Request.Header.Set("Content-Type", "application/json")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Service: mock.NewMockService(ctrl),
			}
			handler := NewHandler(f.Service)

			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			tt.prepare(f, w, c)

			handler.CreateCommand(c)

			resp := w.Result()

			body, err := ioutil.ReadAll(resp.Body)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want.status, resp.StatusCode)
			require.Equal(t, tt.want.res, string(body))

		})
	}
}

func TestHandler_DeleteCommandById(t *testing.T) {
	type want struct {
		res    string
		status int
	}
	type fields struct {
		Service *mock.MockService
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    want
		prepare func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context)
	}{
		{
			name:    "test_1, StatusOK",
			wantErr: false,
			want: want{
				res:    "{\"message\":\"command deleted\"}",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/command/{id}"
				c.Request = httptest.NewRequest(http.MethodDelete, route, nil)
				c.Params = []gin.Param{{
					Key:   "id",
					Value: "1",
				}}
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().DeleteCommandById(c.Request.Context(), "1").Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Service: mock.NewMockService(ctrl),
			}
			handler := NewHandler(f.Service)

			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			tt.prepare(f, w, c)
			handler.DeleteCommandById(c)

			resp := w.Result()

			body, err := ioutil.ReadAll(resp.Body)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want.status, resp.StatusCode)
			require.Equal(t, tt.want.res, string(body))

		})
	}
}

func TestHandler_GetAllCommands(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.GetAllCommands(tt.args.c)
		})
	}
}

func TestHandler_GetCommandById(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.GetCommandById(tt.args.c)
		})
	}
}

func TestHandler_GetCommands(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.GetCommands(tt.args.c)
		})
	}
}

func TestHandler_StopCommandById(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.StopCommandById(tt.args.c)
		})
	}
}

func TestHandler_validateReqCommand(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		req entity.CreateCommandReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			if err := h.validateReqCommand(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateReqCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
