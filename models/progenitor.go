package models

import (
	db "manlogin/dgclient"
	"manlogin/hfuncs"
	"manlogin/services/response"

	"github.com/dgraph-io/dgo"

	"github.com/dgraph-io/dgo/protos/api"

	"encoding/json"
	"fmt"
)

// ─── FATHER OF ALL MODELS ───────────────────────────────────────────────────────

//Model : progenitor of all models
type Model struct {
	//Uid string `json:"uid,omitempty"`
}

const Child = "child"

//
// ──────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: P U B L I C   F U N C T I O N S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────
//

/* Sync : Sync the node
v := Person{...}
Sync(v) : اطلاعات نود را در دیتا بیس دخیره میکند

v.email=""
Sync(v) : اطلاعات نود را در دیتا بیس ذخیره میکند و مقدار email را پاک میکند
*/
func Sync(model interface{}, dgc *dgo.Dgraph) (handler *api.Response, err error) {

	q, err := json.Marshal(model)

	if err != nil {
		return nil, err
	}

	myg := db.NewDgraphTrasn(dgc)
	handler, _, err = myg.Mutate(q)
	if err != nil {
		return handler, err
	}

	return handler, nil
}

//Unlink : unlink node from a node
func Unlink(suid string, Predicate string, uid string, dgc *dgo.Dgraph) error {

	Uid, err := hfuncs.UIDStrX(uid)
	if err != nil {
		return err
	}
	SUid, err := hfuncs.UIDStrX(suid)
	if err != nil {
		return err
	}

	myg := db.NewDgraphTrasn(dgc)

	str := fmt.Sprintf(`<%s> <%s> <%s> . `, SUid, Predicate, Uid)
	q := []byte(str)

	_, err = myg.MutateRDF(q, "delete")

	return err
}

//Link : link one node to other
func Link(suid string, Predicate string, uid string, dgc *dgo.Dgraph) error {
	var res response.Response
	Uid, err := hfuncs.UIDStrX(uid)
	if err != nil {
		return err
	}
	SUid, err := hfuncs.UIDStrX(suid)
	if err != nil {
		return err
	}

	myg := db.NewDgraphTrasn(dgc)

	str := fmt.Sprintf(`<%s> <%s> <%s> . `, SUid, Predicate, Uid)
	q := []byte(str)

	res.Data, err = myg.MutateRDF(q, "set")
	if res.HandleErr(err) {
		return err
	}
	res.Status = response.NULL
	res.Code = response.OK

	return err
}

//LinkFacets : linkFacets one node to other
func LinkFacets(suid string, Predicate string, uid string, facets map[string]string, dgc *dgo.Dgraph) error {
	var res response.Response
	Uid, err := hfuncs.UIDStrX(uid)
	if err != nil {
		return err
	}
	SUid, err := hfuncs.UIDStrX(suid)
	if err != nil {
		return err
	}

	myg := db.NewDgraphTrasn(dgc)
	f := ``
	for key, val := range facets {
		f += fmt.Sprintf(`%s="%s" ,`, key, val)
	}
	if len(f) > 0 {
		f = f[:len(f)-1]
	}
	str := fmt.Sprintf(`<%s> <%s> <%s> (%s) . `, SUid, Predicate, Uid, f)
	q := []byte(str)

	res.Data, err = myg.MutateRDF(q, "set")

	res.Status = response.NULL
	res.Code = response.OK

	return err
}

//set : link one node to other
func Set(uid string, Predicate string, value string, dgc *dgo.Dgraph) error {
	var res response.Response

	SUid, err := hfuncs.UIDStrX(uid)
	if err != nil {
		return err
	}

	myg := db.NewDgraphTrasn(dgc)

	str := fmt.Sprintf(`<%s> <%s> "%s" . `, SUid, Predicate, value)
	q := []byte(str)

	res.Data, err = myg.MutateRDF(q, "set")
	if res.HandleErr(err) {
		return err
	}
	res.Status = response.NULL
	res.Code = response.OK

	return err
}

//
// ──────────────────────────────────────────────────────────────────────────── II ──────────
//   :::::: I N H E R T E D   F U N C T I O N S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────────
//
