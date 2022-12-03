// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token string `json:"token"`
}

type UserDetailRequest struct {
	Identity string `json:"indetity"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendReply struct {
}
