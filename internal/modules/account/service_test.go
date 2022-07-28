package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite
	repo    *RepositoryMock
	service *Service
}

func (s *ServiceTestSuite) SetupTest() {
	s.repo = new(RepositoryMock)
	s.service = &Service{s.repo}
}

func (s *ServiceTestSuite) TestInsert_Success() {
	a := s.Assert()
	ctx := context.Background()
	userId := uint(1)
	in := OpenAccount(userId)

	expectedId := uint(1)
	s.repo.On("Insert", ctx, in).Return(expectedId, nil).Once()

	id, err := s.service.OpenAccount(ctx, userId)

	a.Nil(err, "should not return error")
	a.Equal(expectedId, id, "should return expected id")
	s.repo.AssertExpectations(s.T())
}

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Insert(ctx context.Context, in NewAccountInput) (uint, error) {
	args := r.Called(ctx, in)
	return args.Get(0).(uint), args.Error(1)
}
