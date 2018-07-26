package service

import (
	"github.com/npkhoa2197/example-go/service/book"
	"github.com/npkhoa2197/example-go/service/category"
	"github.com/npkhoa2197/example-go/service/lendbook"
	"github.com/npkhoa2197/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendBookService lendbook.Service
}
