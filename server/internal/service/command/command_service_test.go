package command

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"pstgrprof/server/internal/entity"
	mock "pstgrprof/server/internal/service/command/mocks"
	"testing"
	"time"
)

func TestService_CreateCommand(t *testing.T) {
	type fields struct {
		Repository *mock.MockRepository
		timeout    time.Duration
	}
	type args struct {
		c   context.Context
		req *entity.CreateCommandReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.CreateCommandRes
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "OK",
			args: args{
				c: context.Background(), req: &entity.CreateCommandReq{
					Script:      "echo Hello world!",
					Description: "simple script",
				},
			},
			want: &entity.CreateCommandRes{
				ID:          "1",
				Script:      "echo Hello world!",
				Description: "simple script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(args2.c, &entity.Command{
					ID:          1,
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          1,
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Repository: mock.NewMockRepository(ctrl),
				timeout:    1 * time.Second,
			}
			tt.prepare(tt.args, f)
			service := NewService(f.Repository)
			got, err := service.CreateCommand(tt.args.c, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
