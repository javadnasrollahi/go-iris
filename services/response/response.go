package response

import (
	"fmt"
	"log"
	"manlogin/hfuncs"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
)

const filename = "ServerFiles/errors.log"
const indexFilename = "errorsIndex"

//type response
type Response struct {
	Code     int         `json:"Code,omitempry"`
	Status   int         `json:"Status,omitempry"`
	Message  string      `json:"Message,omitempry"`
	Data     interface{} `json:"Data,omitempry"`
	Messages []Message   `json:"Messages,omitempry"`
}

//error , success ,warn
type Message struct {
	Type  string `json:"Type"` // error , warning, success
	Title string `json:"Title"`
	Text  string `json:"Text"`
}

//HandleErr : handle an error , preapare response if err
func (r *Response) HandleErr(err error) bool {
	if err != nil {
		r.Status = Fail
		r.Code = Fail
		r.Message = hfuncs.ToCamelCase(err.Error())
		r.Data = err
		err := r.log(err, nil)
		if err != nil {
			r.Data = err
		}
		return true
	} else {
		return false
	}

}
func (r *Response) HandleErrCtx(err error, ctx iris.Context) bool {
	if err != nil {
		log.Println(err)
		r.Status = iris.StatusOK
		r.Code = Fail
		r.Message = hfuncs.ToCamelCase(err.Error())
		r.Data = err
		err := r.log(err, ctx)
		if err != nil {
			r.Data = err
		}
		r.Translate(ctx)
		ctx.StatusCode(r.Status)
		return true
	} else {
		return false
	}

}
func (r *Response) Translate(ctx iris.Context) {
	r.Message = i18n.Translate(ctx, r.Message)
}
func (r *Response) HandleStatus(ctx iris.Context) {
	ctx.StatusCode(r.Status)
}

func (r *Response) HandleInternalErr(err error) bool {
	if err != nil {
		r.Status = Fail
		r.Code = Fail
		r.Message = err.Error()
		r.Data = err
		err := r.log(err, nil)
		if err != nil {
			r.Data = err
		}
		return true
	} else {
		return false
	}

}
func (r *Response) OK() {
	r.Code = OK
	r.Status = OK
}
func (r *Response) Fail() {
	r.Code = Fail
	r.Status = BadRequest
}

func (r *Response) log(erro error, ctx iris.Context) error {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0633)
	if err != nil {
		return err
	}

	defer f.Close()
	var logstr string
	if ctx != nil {
		logstr = fmt.Sprintf("%s \t %s \t %s \t %s \t:\t %s \n", time.Now().Format(time.RFC3339), ctx.GetReferrer().Path, ctx.Path(), ctx.RemoteAddr(), erro.Error())
	} else {
		logstr = fmt.Sprintf("%s \t %s \t %s \t %s \t:\t %s \n", time.Now().Format(time.RFC3339), "Progenitor", "Progenitor", "Progenitor", erro.Error())

	}
	_, err = f.WriteString(logstr)
	r.indexErr(erro)
	return err
}
func (r *Response) indexErr(erro error) {
	if _, err := os.Stat(indexFilename); os.IsNotExist(err) {
		os.Create(indexFilename)

	}

	f, _ := os.OpenFile(indexFilename, os.O_APPEND|os.O_WRONLY, 0633)
	defer f.Close()

	var logstr string
	logstr = hfuncs.ToCamelCase(erro.Error())

	f.WriteString(logstr + "\n")

}

//Continue : cheack if continue is ok
func (r *Response) Continue() bool {
	if r.Status > -1 && r.Code > 0 {
		return true
	} else {
		return false
	}
}

//
// ────────────────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: I N F O R M A T I O N   R E S P O N S E S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────────────────
//

const Continue = 100
const SwitchingProtocol = 101
const Processing = 102
const EarlyHints = 103

//
// ──────────────────────────────────────────────────────────────────────────────── II ──────────
//   :::::: S U C C E S S F U L   R E S P O N S E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────
//

const OK = 200
const Created = 201
const Accepted = 202
const NonAuthoritativeInformation = 203
const NoContent = 204
const ResetContent = 205

//
// ──────────────────────────────────────────────────────────────────────────────── III ──────────
//   :::::: R E D I R E C T I O N   M E S S A G E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────
//
const MultipleChoice = 300
const MovedPermanently = 301
const Found = 302
const PermanentRedirect = 308

//
// ──────────────────────────────────────────────────────────────────────────────────── IV ──────────
//   :::::: C L I E N T   E R R O R   R E S P O N S E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────────
//

const BadRequest = 400
const Unauthorized = 401
const PaymentRequired = 402
const Forbidden = 403
const NotFound = 404
const MethodNotAllowed = 405
const NotAcceptable = 406
const RequestTimeout = 408
const TooManyRequests = 429

//
// ──────────────────────────────────────────────────────────────────────────────────── V ──────────
//   :::::: S E R V E R   E R R O R   R E S P O N S E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────────────
//
const InternalServerError = 500
const BadGateway = 502
const ServiceUnavailable = 503
const LoopDetected = 508
const NetworkAuthenticationRequired = 511

//
// ──────────────────────────────────────────────────── VI ──────────
//   :::::: C U S T O N : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//

const Fail = -1
const NULL = 0
