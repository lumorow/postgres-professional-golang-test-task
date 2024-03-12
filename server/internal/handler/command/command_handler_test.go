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
				c.AddParam("id", "1")
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
				res:    "[{\"id\":1,\"script\":\"ls\",\"description\":\"look files\"},{\"id\":2,\"script\":\"ls\",\"description\":\"simple script for look files\"}]",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/all-commands"
				c.Request = httptest.NewRequest(http.MethodGet, route, nil)
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().GetAllCommands(c.Request.Context()).Return(&[]entity.Command{
					{
						ID:          1,
						Script:      "ls",
						Description: "look files",
					},
					{
						ID:          2,
						Script:      "ls",
						Description: "simple script for look files",
					},
				}, nil)
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
			handler.GetAllCommands(c)

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

func TestHandler_GetCommandById(t *testing.T) {
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
				res:    "{\"id\":1,\"script\":\"ls\",\"description\":\"look files\"}",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/command/{id}"
				value := "1"
				c.Request = httptest.NewRequest(http.MethodGet, route, nil)
				c.Params = []gin.Param{{
					Key:   "id",
					Value: value,
				}}
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().GetCommandById(c.Request.Context(), value).Return(&entity.Command{
					ID:          1,
					Script:      "ls",
					Description: "look files",
				}, nil)
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
			handler.GetCommandById(c)

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

func TestHandler_GetCommands(t *testing.T) {
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
				res:    "[{\"id\":1,\"script\":\"ls\",\"description\":\"look files\"},{\"id\":2,\"script\":\"ls\",\"description\":\"simple script for look files\"}]",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/commands"
				//c.Keys["id"] = []string{"1", "2"}
				//c.AddParam("id", "1")
				//c.AddParam("id", "2")

				c.Request = httptest.NewRequest(http.MethodGet, route+"?id=1&id=2", nil)
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().GetCommands(c.Request.Context(), []string{"1", "2"}).Return(&[]entity.Command{
					{
						ID:          1,
						Script:      "ls",
						Description: "look files",
					},
					{
						ID:          2,
						Script:      "ls",
						Description: "simple script for look files",
					},
				}, nil)
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
			handler.GetCommands(c)

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
