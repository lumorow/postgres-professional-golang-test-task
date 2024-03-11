package command

import (
	"context"
	"pstgrprof/server/internal/entity"
	mock "pstgrprof/server/internal/service/command/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService_CreateCommand(t *testing.T) {
	type fields struct {
		Repository   *mock.MockRepository
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
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
			name: "test_1, error not expected",
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
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          1,
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}, nil)
			},
		}, {
			name: "test_2, error not expected",
			args: args{
				c: context.Background(), req: &entity.CreateCommandReq{
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				},
			},
			want: &entity.CreateCommandRes{
				ID:          "1",
				Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
				Description: "combined script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          1,
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}, nil)
			},
		}, {
			name: "test_3, error not expected",
			args: args{
				c: context.Background(), req: &entity.CreateCommandReq{
					Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				},
			},
			want: &entity.CreateCommandRes{
				ID:          "1",
				Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
				Description: "combined script with sleeping",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          1,
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}, nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c: context.Background(), req: &entity.CreateCommandReq{
					Script:      "ls",
					Description: "simple script foe look files",
				},
			},
			want: &entity.CreateCommandRes{
				ID:          "1",
				Script:      "ls",
				Description: "simple script foe look files",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
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
				Repository:   mock.NewMockRepository(ctrl),
				ScriptsCache: mock.NewMockCache(ctrl),
				ExecCmdCache: mock.NewMockCache(ctrl),
			}
			tt.prepare(tt.args, f)
			service := NewService(f.Repository, f.ScriptsCache, f.ExecCmdCache)
			got, err := service.CreateCommand(tt.args.c, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			service.StopSignal <- struct{}{}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestService_GetCommandById(t *testing.T) {
	type fields struct {
		Repository   *mock.MockRepository
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c  context.Context
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Command
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				c:  context.Background(),
				id: "1",
			},
			want: &entity.Command{
				ID:          1,
				Script:      "echo Hello world!",
				Description: "simple script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().GetCommandById(
					context.Background(),
					int64(1)).Return(&entity.Command{
					ID:          1,
					Script:      "echo Hello world!",
					Description: "simple script",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(1), "echo Hello world!").Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Repository:   mock.NewMockRepository(ctrl),
				ScriptsCache: mock.NewMockCache(ctrl),
				ExecCmdCache: mock.NewMockCache(ctrl),
			}
			tt.prepare(tt.args, f)
			service := NewService(f.Repository, f.ScriptsCache, f.ExecCmdCache)
			got, err := service.GetCommandById(tt.args.c, tt.args.id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopSignal <- struct{}{}
			require.Equal(t, tt.want, got)
		})
	}
}
