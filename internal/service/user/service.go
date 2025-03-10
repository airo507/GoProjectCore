package user

import (
	"context"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/app/user"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	userRepository "github.com/airo507/GoProjectCore/internal/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Authorization interface {
	Register(ctx context.Context, userInfo user.ResponseUser) (int64, error)
	Login(ctx context.Context, userData user.InputUser) error
}

type UserService struct {
	repo userRepository.Userable
}

const (
	secretKey = "asdfasfil241234nklasdf"
)

//var SecretKey = []byte("secretkey1")

func NewUserService(userRepository userRepository.Userable) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) Register(ctx context.Context, userInfo user.ResponseUser) (int64, error) {

	hashPassword, err := HashPassword(userInfo.Password)
	if err != nil {
		return 0, fmt.Errorf("Error hashing password: %v", err)
	}
	userData := userEntity.User{
		Id:        userInfo.UserId,
		Login:     userInfo.Login,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     userInfo.Email,
		Password:  hashPassword,
	}

	userCreated, err := s.repo.Create(ctx, userData)
	if err != nil {
		return 0, fmt.Errorf("Failed to create user: %w", err)
	}

	return userCreated, err
}

func (s *UserService) Login(ctx context.Context, input user.InputUser) error {
	checkUser, err := s.repo.Get(ctx, input.Login)
	if err != nil {
		return fmt.Errorf("User not find: %w", err)
	}
	if CheckPassword(input.Password, checkUser.Password) {
		return nil
	}
	return fmt.Errorf("User not find")
}

func HashPassword(password string) (string, error) {
	bytePass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) GenerateJwt(login string) (string, error) {

	claims := jwt.MapClaims{
		"username": login,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Error signing token: %v", err)
	}

	return tokenString, nil
}
