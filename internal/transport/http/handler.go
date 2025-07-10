package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvsdev/auth-service-go/internal/domain"
	"github.com/matheusvsdev/auth-service-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	UserRepo *repository.UserRepository
}

func (h *Handler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Plan     string `json:"plan"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash"})
	}

	user := domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Provider:     domain.LocalProvider,
		Plan:         domain.PlanType(req.Plan),
	}

	if err := h.UserRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucesso"})
}
