package user

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
	crypto  *CrytoMock
	service *service
}

func (s *ServiceTestSuite) SetupTest() {
	s.repo = new(RepositoryMock)
	s.crypto = new(CrytoMock)
	s.service = &service{s.repo, s.crypto}
}

func (s *ServiceTestSuite) TestRegisterUser_Success() {
	a := s.Assert()
	ctx, in := mockRegisterUserInput()

	expectedId := uint(1)
	s.repo.On("Exists", ctx, in.Document, in.Email).Return(false, nil).Once()
	s.crypto.On("Hash", &in.Password).Return(nil).Once()
	s.repo.On("Register", ctx, in).Return(expectedId, nil).Once()

	id, err := s.service.RegisterUser(ctx, in)

	a.Nil(err, "should not return error")
	a.Equal(expectedId, id, "should return expected id")
	s.repo.AssertExpectations(s.T())
	s.crypto.AssertExpectations(s.T())
}

func (s *ServiceTestSuite) TestRegisterUser_UserAlreadyExists() {
	a := s.Assert()
	ctx, in := mockRegisterUserInput()

	s.repo.On("Exists", ctx, in.Document, in.Email).Return(true, nil).Once()

	_, err := s.service.RegisterUser(ctx, in)

	a.Equal(ErrUserAlreadyExists, err)
	s.repo.AssertExpectations(s.T())
}

func (s *ServiceTestSuite) TestCheckPassword_Success() {
	a := s.Assert()
	hash := "anything"
	password := "12345678"
	s.crypto.On("Compare", hash, password).Return(true).Once()

	ok := s.service.CheckPassword(hash, password)

	a.True(ok)
	s.crypto.AssertExpectations(s.T())
}

func (s *ServiceTestSuite) TestCheckPassword_Failed() {
	a := s.Assert()
	hash := "anything"
	password := "12345678"
	s.crypto.On("Compare", hash, password).Return(false).Once()

	ok := s.service.CheckPassword(hash, password)

	a.False(ok)
	s.crypto.AssertExpectations(s.T())
}

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Exists(ctx context.Context, document string, email string) (bool, error) {
	args := r.Called(ctx, document, email)
	return args.Bool(0), args.Error(1)

}
func (r *RepositoryMock) Register(ctx context.Context, in RegisterUserInput) (uint, error) {
	args := r.Called(ctx, in)
	return args.Get(0).(uint), args.Error(1)
}

type CrytoMock struct {
	mock.Mock
}

func (c *CrytoMock) Hash(password *string) error {
	args := c.Called(password)
	return args.Error(0)
}

func (c *CrytoMock) Compare(hash, password string) bool {
	args := c.Called(hash, password)
	return args.Bool(0)
}

func mockRegisterUserInput() (context.Context, RegisterUserInput) {
	return context.Background(), RegisterUserInput{
		Kind:     Seller,
		Document: "57872973006",
		Name:     "Jonh Doe",
		Email:    "jonh@doe.com",
		Password: "12345678",
	}
}
