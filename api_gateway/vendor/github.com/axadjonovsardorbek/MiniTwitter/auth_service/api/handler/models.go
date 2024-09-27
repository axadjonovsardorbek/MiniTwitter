package handler

type UserUpdateReq struct {
	Name      string
	Username  string
	Email     string
	Bio       string
	Image_url string
}

type UserCreateReq struct {
	Name      string
	Username  string
	Bio       string
	Image_url string
	Password  string
	Email     string
}

type LoginReq struct {
	Username string
	Password string
}

type ChangePasswordReq struct {
	OldPassword string
	NewPassword string
}
