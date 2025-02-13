package handler

import (
	"golang.org/x/net/context"
)

type UaaHandler interface {
	Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) error
}

type Uaa struct {
}

func NewUaaHandler() UaaHandler {
	return &Uaa{}
}

type LoginReq struct{}

type LoginRsp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (u *Uaa) Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) error {
	rsp.Id, rsp.Name = "1", "小明"
	return nil
}
