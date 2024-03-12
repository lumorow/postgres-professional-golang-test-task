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
	w := httptest.NewRecorder()
	type want struct {
		res    string
		status int
	}
	type fields struct {
		Service *mock.MockService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    want
		args    args
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1, error not expected",
			wantErr: false,
			want: want{
				res:    "{\"id\":\"1\",\"script\":\"ls\",\"description\":\"look files\"}",
				status: http.StatusOK,
			},
			args: args{
				c: GetTestGinContext(w),
			},
			prepare: func(args args, fields fields) {
				args.c.Request = httptest.NewRequest(http.MethodPost, "/command", strings.NewReader(`{"script": "ls", "description": "look files"}`))
				args.c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().CreateCommand(args.c.Request.Context(), &entity.CreateCommandReq{
					Script:      "ls",
					Description: "look files",
				}).Return(&entity.CreateCommandRes{
					ID:          "1",
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
			tt.prepare(tt.args, f)
			handler.CreateCommand(tt.args.c)

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
			h.DeleteCommandById(tt.args.c)
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
