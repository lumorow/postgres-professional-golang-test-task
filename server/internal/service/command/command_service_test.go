package command

import (
	"context"
	"errors"
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
				ID:          "2",
				Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
				Description: "combined script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          2,
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
				ID:          "3",
				Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
				Description: "combined script with sleeping",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          3,
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}, nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c: context.Background(), req: &entity.CreateCommandReq{
					Script:      "ls",
					Description: "simple script for look files",
				},
			},
			want: &entity.CreateCommandRes{
				ID:          "4",
				Script:      "ls",
				Description: "simple script for look files",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.req.Script,
					Description: args2.req.Description,
				}).Return(&entity.Command{
					ID:          4,
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

			service.StopRunner()
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
		}, {
			name: "test_2, error not expected",
			args: args{
				c:  context.Background(),
				id: "2",
			},
			want: &entity.Command{
				ID:          2,
				Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
				Description: "combined script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().GetCommandById(
					context.Background(),
					int64(2)).Return(&entity.Command{
					ID:          2,
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(2), "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami").Return(nil)
			},
		}, {
			name: "test_3, error not expected",
			args: args{
				c:  context.Background(),
				id: "3",
			},
			want: &entity.Command{
				ID:          3,
				Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
				Description: "combined script with sleeping",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().GetCommandById(
					context.Background(),
					int64(3)).Return(&entity.Command{
					ID:          3,
					Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(3), "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"").Return(nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c:  context.Background(),
				id: "4",
			},
			want: &entity.Command{
				ID:          4,
				Script:      "ls",
				Description: "simple script for look files",
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().GetCommandById(
					context.Background(),
					int64(4)).Return(&entity.Command{
					ID:          4,
					Script:      "ls",
					Description: "simple script for look files",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(4), "ls").Return(nil)
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

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestService_DeleteCommandById(t *testing.T) {
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
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().DeleteCommandById(
					context.Background(),
					int64(1)).Return(nil)
			},
		}, {
			name: "test_2, error not expected",
			args: args{
				c:  context.Background(),
				id: "2",
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().DeleteCommandById(
					context.Background(),
					int64(2)).Return(nil)
			},
		}, {
			name: "test_3, error not expected",
			args: args{
				c:  context.Background(),
				id: "3",
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().DeleteCommandById(
					context.Background(),
					int64(3)).Return(nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c:  context.Background(),
				id: "4",
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().DeleteCommandById(
					context.Background(),
					int64(4)).Return(nil)
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
			err := service.DeleteCommandById(tt.args.c, tt.args.id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
		})
	}
}

func TestService_GetAllCommands(t *testing.T) {
	type fields struct {
		Repository   *mock.MockRepository
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]entity.Command
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				c: context.Background(),
			},
			want: &[]entity.Command{
				{
					ID:          int64(1),
					Script:      "echo Hello world!",
					Description: "simple script",
				},
				{
					ID:          2,
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				},
				{
					ID:          3,
					Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				},
				{
					ID:          4,
					Script:      "ls",
					Description: "simple script for look files",
				},
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().GetAllCommands(
					context.Background()).Return(&[]entity.Command{
					{
						ID:          int64(1),
						Script:      "echo Hello world!",
						Description: "simple script",
					},
					{
						ID:          2,
						Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
						Description: "combined script",
					},
					{
						ID:          3,
						Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
						Description: "combined script with sleeping",
					},
					{
						ID:          4,
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
				Repository:   mock.NewMockRepository(ctrl),
				ScriptsCache: mock.NewMockCache(ctrl),
				ExecCmdCache: mock.NewMockCache(ctrl),
			}
			tt.prepare(tt.args, f)
			service := NewService(f.Repository, f.ScriptsCache, f.ExecCmdCache)
			got, err := service.GetAllCommands(tt.args.c)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestService_GetCommands(t *testing.T) {
	type fields struct {
		Repository   *mock.MockRepository
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c   context.Context
		ids []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]entity.Command
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				c:   context.Background(),
				ids: []string{"1", "2", "3", "4"},
			},
			want: &[]entity.Command{
				{
					ID:          int64(1),
					Script:      "echo Hello world!",
					Description: "simple script",
				},
				{
					ID:          int64(2),
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				},
				{
					ID:          int64(3),
					Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				},
				{
					ID:          int64(4),
					Script:      "ls",
					Description: "simple script for look files",
				},
			},
			prepare: func(args2 args, fields fields) {
				fields.Repository.EXPECT().GetCommands(
					context.Background(), []int64{1, 2, 3, 4}).Return(&[]entity.Command{
					{
						ID:          int64(1),
						Script:      "echo Hello world!",
						Description: "simple script",
					},
					{
						ID:          int64(2),
						Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
						Description: "combined script",
					},
					{
						ID:          int64(3),
						Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
						Description: "combined script with sleeping",
					},
					{
						ID:          int64(4),
						Script:      "ls",
						Description: "simple script for look files",
					},
				}, nil)
				var i int64
				resps := []string{
					"echo Hello world!",
					"echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					"ls",
				}
				for i = 0; i < 4; i++ {
					fields.ScriptsCache.EXPECT().Set(i+1, resps[i]).Return(nil)
				}
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
			got, err := service.GetCommands(tt.args.c, tt.args.ids)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestService_StopCommandById(t *testing.T) {
	type fields struct {
		Repository   *mock.MockRepository
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				id: "1",
			},
			wantErr: true,
			prepare: func(args args, fields fields) {
				fields.ExecCmdCache.EXPECT().Get(int64(1)).Return(nil, errors.New("not found value"))
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
			err := service.StopCommandById(tt.args.id)

			service.StopRunner()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
		})
	}
}
