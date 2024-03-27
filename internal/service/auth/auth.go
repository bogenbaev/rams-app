package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"rams/internal/repository"
	"rams/pkg/models"
	"strconv"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UId          string `json:"uid"`
	UserFullName string `json:"user_full_name"`
}

type Service struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *Service) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(ctx, username, s.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	return s.Generate(strconv.Itoa(user.ID), user.FullName)
}

func (s *Service) Generate(id string, fullName string) (string, error) {
	fmt.Printf("id: %v\n", id)
	fmt.Printf("fullName: %v\n", fullName)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "RAMS",
		},
		UId:          id,
		UserFullName: fullName,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) CreateUser(ctx context.Context, user models.User) error {
	user.Password = s.GeneratePasswordHash(user.Password)
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetListUser(ctx context.Context) ([]models.User, error) {
	return s.repo.GetListUser(ctx)
}

func (s *Service) GetUserByID(ctx context.Context, user models.User) (models.User, error) {
	return s.repo.GetUserByID(ctx, user)
}

func (s *Service) GetUserByLogin(ctx context.Context, login string) (user models.User, err error) {
	return s.repo.GetUserByLogin(ctx, login)
}

func (s *Service) GetUser(ctx context.Context, username, password string) (user models.User, err error) {
	return s.repo.GetUser(ctx, username, password)
}
