package user

import (
	"context"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/api"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/airo507/GoProjectCore/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type UserService struct {
	repo repository.Userable
}

const (
	// TODO: Что тут значит 1?
	// TODO: Лучше не хранить секрет в константах проекта. Нужно задавать
	// его в конфиг файле или env переменной.
	secretKey = "secretkey1"
)

func NewUserService(userRepository repository.Userable) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) Register(ctx context.Context, userInfo api.ResponseUser) (int64, error) {

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

	// TODO: Не обрабатывается ошибка.
	checkUser, _ := s.repo.Get(ctx, userData.Login)

	if checkUser.Login == userInfo.Login {
		return 0, fmt.Errorf("User already exists")
	}

	userCreated, err := s.repo.Create(ctx, userData)
	if err != nil {
		return 0, fmt.Errorf("Failed to create user: %w", err)
	}

	// TODO: У тебя есть if на обработку ошибки. Тут она всегда будет nil.
	// Стоит это явно показать.
	// TODO: Тип возвращаемого значения не совпадает с типом ID в структуре userEntity.User.
	return userCreated, err
}

func (s *UserService) Login(ctx context.Context, input api.InputUser) (string, error) {
	checkUser, err := s.repo.Get(ctx, input.Login)
	if err != nil {
		return "", fmt.Errorf("User not find: %w", err)
	}
	if !s.CheckPassword(input.Password, checkUser.Password) {
		return "", fmt.Errorf("Invalid password or login")
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
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
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

func (s *UserService) GetUsers(ctx context.Context) ([]userEntity.User, error) {
	// TODO: Потенциально можем доставать очень много записей. Нужна пагинация.
	users, err := s.repo.GetUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}
