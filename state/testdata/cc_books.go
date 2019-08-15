package testdata

import (
	"github.com/soodakshay/cckit/extensions/debug"
	"github.com/soodakshay/cckit/extensions/owner"
	"github.com/soodakshay/cckit/router"
	p "github.com/soodakshay/cckit/router/param"
	"github.com/soodakshay/cckit/state/testdata/schema"
)

func NewBooksCC() *router.Chaincode {
	r := router.New(`books`)
	debug.AddHandlers(r, `debug`, owner.Only)

	r.Init(owner.InvokeSetFromCreator).
		Invoke(`bookList`, bookList).
		Invoke(`bookGet`, bookGet, p.String(`id`)).
		Invoke(`bookInsert`, bookInsert, p.Struct(`book`, &schema.Book{})).
		Invoke(`bookUpsert`, bookUpsert, p.Struct(`book`, &schema.Book{})).
		Invoke(`bookDelete`, bookDelete, p.String(`id`))

	return router.NewChaincode(r)
}

func bookList(c router.Context) (interface{}, error) {
	return c.State().List(schema.BookEntity, &schema.Book{})
}

func bookInsert(c router.Context) (interface{}, error) {
	book := c.Param(`book`)
	return book, c.State().Insert(book)
}

func bookUpsert(c router.Context) (interface{}, error) {
	book := c.Param(`book`)
	return book, c.State().Put(book)
}

func bookGet(c router.Context) (interface{}, error) {
	return c.State().Get(schema.Book{Id: c.ParamString(`id`)})
}

func bookDelete(c router.Context) (interface{}, error) {
	return nil, c.State().Delete(schema.Book{Id: c.ParamString(`id`)})
}
