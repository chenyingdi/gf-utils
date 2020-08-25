package gfUtils

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func Find(req *ghttp.Request) {
	rsp := NewResp(req)
	q := &query{}

	err := gconv.Struct(req.Get("query"), q)
	if err != nil {
		rsp.FAIL(err.Error())
		return
	}

	prefix := req.GetString("prefix")
	suffix := req.GetString("suffix")

	table := prefix

	if suffix != "" {
		table += fmt.Sprintf("_%s", suffix)
	}

	rep := NewRepo(table)

	res, err := rep.Find(*q)
	if err != nil {
		rsp.FAIL(err.Error())
	}

	rsp.SUCCESS(res)
}

func New(req *ghttp.Request) {
	rsp := NewResp(req)

	prefix := req.GetString("prefix")
	suffix := req.GetString("suffix")
	input := gconv.Map(req.Get("input"))

	table := prefix

	if suffix != "" {
		table += fmt.Sprintf("_%s", suffix)
	}

	rep := NewRepo(table)

	id, err := rep.Insert(input)
	if err != nil {
		rsp.FAIL(err.Error())
		return
	}

	res, err := rep.Find(query{Where: g.Map{
		"id": id,
	}})
	if err != nil {
		rsp.FAIL(err.Error())
		return
	}

	rsp.SUCCESS(res)
}

func Edit(req *ghttp.Request) {
	rsp := NewResp(req)

	where := gconv.Map(req.Get("where"))
	updates := gconv.Map(req.Get("updates"))
	prefix := req.GetString("prefix")
	suffix := req.GetString("suffix")
	table := prefix

	if suffix != "" {
		table += fmt.Sprintf("_%s", suffix)
	}

	rep := NewRepo(table)

	err := rep.Update(where, updates)
	if err != nil {
		rsp.FAIL(err.Error())
		return
	}

	res, err := rep.Find(query{
		Where: where,
	})
	if err != nil {
		rsp.FAIL(err.Error())
		return
	}

	rsp.SUCCESS(res)
}

func Delete(req *ghttp.Request) {
	rsp := NewResp(req)

	where := gconv.Map(req.Get("where"))

	prefix := req.GetString("prefix")
	suffix := req.GetString("suffix")
	table := prefix

	if suffix != "" {
		table += fmt.Sprintf("_%s", suffix)
	}

	rep := NewRepo(table)

	err := rep.Delete(where)
	if err != nil {
		rsp.FAIL(err.Error())
		return
	}

	rsp.SUCCESS(nil)
}
