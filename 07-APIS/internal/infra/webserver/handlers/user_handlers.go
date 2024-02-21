package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/dto"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/entity"
	"github.com/obrunogonzaga/pos-go-expert/07-APIS/07-APIS/internal/infra/database"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB        database.UserInterface
	JWT           *jwtauth.JWTAuth
	JwtExpiriedIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, JwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        userDB,
		JWT:           jwt,
		JwtExpiriedIn: JwtExperiesIn,
	}
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.JWT.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiriedIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary 	Create a user
// @Description Create a user
// @Tags 		users
// @Accept		json
// @Produce		json
// @Param		request		body 	dto.CreateUserInput		true	"user request"
// @Success		201
// @Failure		500	{object} Error
// @Router		/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
