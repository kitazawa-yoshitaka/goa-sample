package calcapi

import (
	"context"
	"fmt"
	calc "goa-sample/gen/calc"
	"log"
	"net/http"

	"github.com/pkg/errors"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	if p.A == 1 {
		return 0, calc.MakeDivByZero(fmt.Errorf("完全にテスト用の適当なエラー"))
	}

	if p.A == 2 {
		return 0, calc.MakeTimeout(fmt.Errorf("まさかのタイムアウト"))
	}
	return p.A + p.B, nil
}

type MyErrorResponse struct {
	Type   string
	Code   string
	Title  string
	Detail string
}

func NewMyErrorResponse(err error) goahttp.Statuser {
	switch err := errors.Cause(err).(type) {
	case *goa.ServiceError:
		if err.Name == "DivByZero" {
			return &MyErrorResponse{
				Type:   "http://test.test",
				Code:   "001",
				Title:  err.Message,
				Detail: "１つ目のパラメータに1を指定しないでください。（ひどい）",
			}
		} else if err.Name == "Timeout" {
			return &MyErrorResponse{
				Type:   "http://test.test",
				Code:   "002",
				Title:  err.Message,
				Detail: "1つめのパラメータに2を指定しないでください。（ひどい）",
			}
		} else {
			return &MyErrorResponse{
				Type:   "http://test.test",
				Code:   "000",
				Title:  err.Message,
				Detail: "管理者にお問い合わせください。",
			}
		}

	default:
		return &MyErrorResponse{
			Type:   "http://test.test",
			Code:   "000",
			Title:  err.Error(),
			Detail: "管理者にお問い合わせください。",
		}
	}
}

// StatusCode implements a heuristic that computes a HTTP response status code
// appropriate for the timeout, temporary and fault characteristics of the
// error. This method is used by the generated server code when the error is not
// described explicitly in the design.
func (resp *MyErrorResponse) StatusCode() int {
	return http.StatusBadRequest
}
