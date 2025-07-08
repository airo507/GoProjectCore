package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/airo507/GoProjectCore/internal/api"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	userRepository "github.com/airo507/GoProjectCore/internal/repository/user"
	"github.com/golang-jwt/jwt/v5"
)

type UserServiceInterface interface {
	Register(ctx context.Context, userInfo api.ResponseUser) (int64, error)
	Login(ctx context.Context, userData api.InputUser) (string, error)
	CheckToken(tokenString string) (string, error)
	GetUsers(ctx context.Context) ([]userEntity.User, error)
}

type UserService struct {
	repo      userRepository.UserRepository
	secretKey string
}

func NewUserService(userRepository userRepository.UserRepository, secretKey string) *UserService {
	return &UserService{
		repo:      userRepository,
		secretKey: secretKey,
	}
}

func (s *UserService) Register(ctx context.Context, userInfo api.ResponseUser) (int64, error) {
	hashPassword, err := s.HashPassword(userInfo.Password)
	if err != nil {
		return 0, err
	}

	userData := userEntity.User{
		Id:        userInfo.UserId,
		Login:     userInfo.Login,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     userInfo.Email,
		Password:  hashPassword,
	}

	checkUser, err := s.repo.Get(ctx, userData.Login)
	if err != nil {
		return 0, err
	}

	if checkUser.Login == userInfo.Login {
		return 0, err
	}

	userCreated, err := s.repo.Create(ctx, userData)
	if err != nil {
		return 0, err
	}

	return userCreated, nil
}

func (s *UserService) Login(ctx context.Context, input api.InputUser) (string, error) {
	checkUser, err := s.repo.Get(ctx, input.Login)
	if err != nil {
		return "", err
	}

	if !s.CheckPassword(input.Password, checkUser.Password) {
		return "", err
	}

	if checkUser.Login != input.Login {
		return "", fmt.Errorf("Invalid login")
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
		"exp":   time.Now().UTC().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) CheckToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", fmt.Errorf("Token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Token claims are invalid")
	}

	login, ok := claims["login"].(string)
	if !ok {
		return "", fmt.Errorf("login claim missing")
	}

	return login, nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]userEntity.User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}
