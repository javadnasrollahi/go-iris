package auth

/*
TODO: این فایل خام و کلی میباشد. ایده های خود را برای بهبود مطرح کنید.
*/
import (
	"encoding/json"
	"fmt"
	db "manlogin/dgclient"
	"manlogin/services/response"

	"github.com/dgraph-io/dgo"

	"github.com/kataras/iris"
)

func Owner(ctx iris.Context, OwnerUID uint64, TargetUID uint64, dgc *dgo.Dgraph) (response.Response, bool) {
	res := response.Response{}
	myg := db.NewDgraphTrasn(dgc)
	q := fmt.Sprintf(`
		{
			create(func: uid(%#x)) {
				uid
				creator  @filter(uid(%#x)){
					uid
				}
			}
			own(func:uid(%#x)) {
				uid
				owner @filter(uid(%#x)){
					uid
				}
			}
		}
		`, OwnerUID, TargetUID, OwnerUID, TargetUID)
	var dbres struct {
		Create []struct {
			UID     string `json:"uid,omitempty"`
			Creator []struct {
				UID string `json:"uid,omitempty"`
			} `json:"creator,omitempty"`
		} `json:"create"`
		Own []struct {
			UID   string `json:"uid,omitempty"`
			Owner []struct {
				UID string `json:"uid,omitempty"`
			} `json:"owner,omitempty"`
		} `json:"own"`
	}
	dbstrres, _ := myg.Query(q)
	err := json.Unmarshal(dbstrres.Json, &dbres)

	if res.HandleErrCtx(err , ctx) {
		return res, false
	}
	if len(dbres.Create[0].Creator) > 0 || len(dbres.Own[0].Owner) > 0 {
		res.OK()
		return res, true
	} else {
		res.Fail()
		return res, false
	}

}
