package main

import (
	"log"
	"github.com/go-pg/pg/orm"
)

type StudentFilter struct {
	orm.Pager 
}

func (f *StudentFilter) Filter(q *orm.Query) (*orm.Query, error) {
	f.Pager.MaxLimit = 50 // default max limit is 1000
	f.Pager.MaxOffset = 100 // default max offset is 1000000
	f.Pager.Offset = 5
	f.Pager.Limit = 2

	q = q.Apply(f.Pager.Paginate)

	return q, nil
}

func LimitOffset() {
	var filter StudentFilter
	var students []Student
	err := Db.Model(&students).Apply(filter.Filter).OrderExpr("id DESC").Select()

	if err != nil {
		panic(err)
	}

	for _,res := range students {
		log.Println(res)
	}
}