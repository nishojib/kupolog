package handlers

import (
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/data/models"
	"github.com/nishojib/ffxivdailies/internal/validator"
	"github.com/uptrace/bun"
)

// AddUserInput represents the input for the add user endpoint.
type AddUserInput struct {
	UserID string `json:"user_id" example:"1234567890"`
	Name   string `json:"name"    example:"John Doe"`
	Email  string `json:"email"   example:"johndoe@example.com"`
	Image  string `json:"image"   example:"https://example.com/image.png"`
}

// AddUser godoc
//
//	@Summary		Add an user
//	@Description	add by json user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		AddUserInput	true	"Add user"
//	@Success		201		{object}	object{user=models.User}
//	@Failure		400		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		422		{object}	object{errors=object{email=string,user_id=string,name=string,image=string},status=int,title=string,type=string}
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/users [post]
func AddUser(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input AddUserInput

		err := api.ReadJSON(w, r, &input)
		if err != nil {
			api.BadRequestResponse(w, r, err)
			return
		}

		user := &models.User{
			Name:   input.Name,
			Email:  input.Email,
			Image:  input.Image,
			UserID: input.UserID,
		}

		v := validator.New()
		if user.Validate(v); !v.Valid() {
			api.FailedValidationResponse(w, r, v.Errors)
			return
		}

		// err = repository.NewUserRepository(db).Insert(user)
		// if err != nil {
		// 	api.ServerErrorResponse(w, r, err)
		// 	return
		// }

		err = api.WriteJSON(w, http.StatusCreated, api.Envelope[models.User]{"user": *user}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	}
}
