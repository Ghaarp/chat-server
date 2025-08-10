package chat

import (
	"context"
	"testing"

	chat "github.com/Ghaarp/chat-server/internal/api"
	"github.com/Ghaarp/chat-server/internal/service"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Ghaarp/chat-server/internal/service/chat/model"
	serviceMocks "github.com/Ghaarp/chat-server/internal/service/mocks"
	generated "github.com/Ghaarp/chat-server/pkg/chat_v1"
)

func TestSendMessage(t *testing.T) {

	t.Parallel()

	type chatServiceMockFunc func(controller *minimock.Controller) service.ChatService

	type args struct {
		ctx     context.Context
		request *generated.SendMessageRequest
	}

	var (
		ctx            = context.Background()
		mockController = minimock.NewController(t)

		from      = gofakeit.Int64()
		chatId    = gofakeit.Int64()
		text      = gofakeit.BeerName()
		timestamp = gofakeit.Date()

		request = &generated.SendMessageRequest{
			From:      from,
			Chatid:    chatId,
			Text:      text,
			Timestamp: timestamppb.New(timestamp),
		}

		model = &model.SendMessageRequest{
			From:      from,
			Chatid:    chatId,
			Text:      text,
			Timestamp: timestamp,
		}

		result = &generated.SendMessageResponse{}
	)

	tests := []struct {
		name            string
		args            args
		want            *generated.SendMessageResponse
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
				mock.SendMessageMock.Expect(ctx, model).Return(nil)
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

			sendResult, err := api.SendMessage(test.args.ctx, test.args.request)
			require.Equal(t, test.err, err)
			require.Equal(t, test.want, sendResult)
		})
	}
}
