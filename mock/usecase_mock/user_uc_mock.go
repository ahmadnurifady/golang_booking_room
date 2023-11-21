package usecasemock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type MockEmailService struct {
	mock.Mock
}
type UserUseCaseMock struct {
	mock.Mock
}

func (b *UserUseCaseMock) FindById(id string) (model.User, error) {
	args := b.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

// func (m *MockEmailService) SendEmail(to, subject, body string) error {
// 	args := m.Called(to, subject, body)
// 	return args.Get(0), errors
// }
