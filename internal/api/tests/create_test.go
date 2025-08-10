package chat

import (
	"context"
	"testing"

	chat "github.com/Ghaarp/chat-server/internal/api"
	"github.com/Ghaarp/chat-server/internal/service"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Ghaarp/chat-server/internal/service/chat/model"
	serviceMocks "github.com/Ghaarp/chat-server/internal/service/mocks"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {

	t.Parallel()

	type chatServiceMockFunc func(controller *minimock.Controller) service.ChatService

	type args struct {
		ctx     context.Context
		request *generated.CreateRequest
	}

	var (
		ctx            = context.Background()
		mockController = minimock.NewController(t)

		id       = gofakeit.Int64()
		author   = gofakeit.Int64()
		chatName = gofakeit.Name()
		users    = []int64{
			gofakeit.Int64(),
		}

		request = &generated.CreateRequest{
			Author:   author,
			ChatName: chatName,
			Users:    users,
		}

		result = &generated.CreateResponse{
			Id: id,
		}

		model = &model.CreateRequest{
			Author:   author,
			ChatName: chatName,
			Users:    users,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *generated.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "Case 1",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: result,
			err:  nil,
			chatServiceMock: func(controller *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(t)
				mock.CreateMock.Expect(ctx, model).Return(id, nil)
				return mock
			},
		},
	}

	for _, test := range tests {

		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := test.chatServiceMock(mockController)
			api := chat.NewChatImplementation(authServiceMock)

			createResult, err := api.Create(test.args.ctx, test.args.request)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, createResult)
		})
	}
}
