package model

type UserAuth struct {
	ID   int64
	Info UserAuthInfo
}

type UserAuthInfo struct {
	Email    string
	Name     string
	Password string
}
