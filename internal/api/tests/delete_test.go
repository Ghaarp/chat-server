package chat

import (
	"context"
	"testing"

	chat "github.com/Ghaarp/chat-server/internal/api"
	"github.com/Ghaarp/chat-server/internal/service"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	serviceMocks "github.com/Ghaarp/chat-server/internal/service/mocks"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

func TestDelete(t *testing.T) {

	t.Parallel()

	type chatServiceMockFunc func(controller *minimock.Controller) service.ChatService

	type args struct {
		ctx     context.Context
		request *generated.DeleteRequest
	}

	var (
		ctx            = context.Background()
		mockController = minimock.NewController(t)

		id = gofakeit.Int64()

		request = &generated.DeleteRequest{
			Id: id,
		}

		result = &generated.DeleteResponse{}
	)

	tests := []struct {
		name            string
		args            args
		want            *generated.DeleteResponse
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
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := test.chatServiceMock(mockController)
			api := chat.NewChatImplementation(authServiceMock)

			deleteResult, err := api.Delete(test.args.ctx, test.args.request)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, deleteResult)
		})
	}
}
