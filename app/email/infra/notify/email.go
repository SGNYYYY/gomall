package notify

import (
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("%v", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
