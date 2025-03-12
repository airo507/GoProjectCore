package user

import (
	"context"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/app/user"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	userRepository "github.com/airo507/GoProjectCore/internal/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type Authorization interface {
	Register(ctx context.Context, userInfo user.ResponseUser) (int64, error)
	Login(ctx context.Context, userData user.InputUser) (string, error)
}

type UserService struct {
	repo userRepository.Userable
}

const (
	secretKey = "secretkey1"
)

func NewUserService(userRepository userRepository.Userable) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) Register(ctx context.Context, userInfo user.ResponseUser) (int64, error) {

	hashPassword, err := s.HashPassword(userInfo.Password)
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

func (s *UserService) Login(ctx context.Context, input user.InputUser) (string, error) {
	checkUser, err := s.repo.Get(ctx, input.Login)
	if err != nil {
		return "", fmt.Errorf("User not find: %w", err)
	}
	if !s.CheckPassword(input.Password, checkUser.Password) {
		return "", fmt.Errorf("Invalid password or login")
	}
	token, err := s.GenerateJwt(input.Login)

	if err != nil {
		return "", fmt.Errorf("Error generating token: %v", err)
	}

	return token, nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytePass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *UserService) CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) GenerateJwt(login string) (string, error) {

	claims := jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		slog.Error("Error signing token: %v", err)
		return "", fmt.Errorf("Error signing token: %v", err)
	}

	return tokenString, nil
}

func (s *UserService) CheckToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", fmt.Errorf("Token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("Token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Token claims are invalid")
	}
	return claims["login"].(string), nil
}
