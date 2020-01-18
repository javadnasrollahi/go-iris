package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	db "manlogin/dgclient"
	"manlogin/hfuncs"
	"manlogin/models"
	"manlogin/services/acl"
	"manlogin/services/response"

	"github.com/dgraph-io/dgo"

	"github.com/kataras/iris/middleware/i18n"

	"github.com/kataras/iris"
)

func Authentication(ctx iris.Context, ACL acl.AclVal, dismissAcl bool, dgc *dgo.Dgraph) response.Response {

	var res response.Response
	err, uacl := BasicOuth(ctx, dgc)
	if res.HandleErrCtx(err, ctx) {
		res.Message = i18n.Translate(ctx, "authenticationFailed")
		return res
	}
	if dismissAcl {
		return res
	}

	if !acl.Allow(uacl, ACL) {
		err := errors.New("UserNotAllowed")
		res.HandleErrCtx(err, ctx)
		return res
	}
	return res
}
func AuthenticationMultiACL(ctx iris.Context, ACL []acl.AclVal, dismissAcl bool, dgc *dgo.Dgraph) (response.Response, []bool) {
	//* Uid := ctx.GetHeader("X-USER")
	//* Token := ctx.GetHeader("Authorization-Token")

	var res response.Response
	alist := make([]bool, len(ACL))
	err, uacl := BasicOuth(ctx, dgc)
	if res.HandleErrCtx(err, ctx) {
		return res, alist
	}
	if dismissAcl {
		return res, alist
	}
	some := false
	for i, a := range ACL {
		if acl.Allow(uacl, a) {
			alist[i] = true
			some = true
		} else {
			alist[i] = false
		}
	}
	if !some {
		res.HandleErrCtx(errors.New("AccessDeniend"), ctx)
		return res, alist
	}
	return res, alist
}

func CustomAuth(UserUid string, Token string, dgc *dgo.Dgraph) error {

	Uid, err := hfuncs.UIDStrX(UserUid)
	if err != nil {
		return err
	}
	token := Token

	dbresstr, _ := db.NewDgraphTrasn(dgc).Query(fmt.Sprintf(`
	{
		user(func: uid(%s)) @filter(eq(token,"%s")) @cascade {
		   uid
		   acl
		   
		}
	}
	`, Uid, token))
	var dbres struct {
		User []models.Person
	}
	//ctx.Write(dbresstr)
	if err := json.Unmarshal(dbresstr.Json, &dbres); err != nil {
		return err
	}
	if len(dbres.User) < 1 {
		return errors.New("invalidUserData")
	} else if len(dbres.User) > 1 {
		return errors.New("NonUnicUser")
	}
	return nil

}

func GetAcl(uid string, dgc *dgo.Dgraph) map[string]acl.AclVal {
	myg := db.NewDgraphTrasn(dgc)
	q := fmt.Sprintf(`
		{
			user(func: uid(%s)) {
				acl
			}
		}
		`, uid)

	resb, _ := myg.Query(q)
	var resstrc struct {
		User []models.Person `json:"user"`
	}
	json.Unmarshal(resb.Json, &resstrc)
	useraclval := resstrc.User[0].Acl
	return acl.Privileges(useraclval)

}

//BasicOuth : outhenticate the user with uid and token
func BasicOuth(ctx iris.Context, dgc *dgo.Dgraph) (error, uint64) {
	uid := ctx.GetHeader("X-User")
	Uid, err := hfuncs.UIDStrX(uid)
	if err != nil {
		return err, 0
	}
	token := ctx.GetHeader("Authorization-Token")
	devSerial := ctx.GetCookie("deviceSerial")
	dbresstr, _ := db.NewDgraphTrasn(dgc).Query(fmt.Sprintf(`
	{
		User(func: uid(%s))  @filter(eq(deviceSerial,"%s") AND eq(token,"%s")) @cascade {
		   uid
		   acl
		  
		}
	}
	`, Uid, devSerial, token))
	var dbres struct {
		User []models.Person
	}
	//ctx.Write(dbresstr)
	//log.Println(dbresstr.GetJson())
	err = json.Unmarshal(dbresstr.Json, &dbres)
	if err != nil {
		return err, 0
	}
	if len(dbres.User) < 1 {

		return errors.New("NOUSER"), 0
	}

	return nil, dbres.User[0].Acl

}
