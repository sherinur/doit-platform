package handler

type UserHandler struct {
	uc UserUsecase
}

func NewUserHandler(uc UserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}
