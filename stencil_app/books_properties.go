package stencil_app

import "github.com/beauxarts/fedorov/data"

var BooksProperties = []string{
	//data.TitleProperty,
	data.BookTypeProperty,
	data.BookCompletedProperty,
	data.AuthorsProperty,
	data.DateCreatedProperty,
}

var BooksLabels = []string{
	data.BookTypeProperty,
	data.BookCompletedProperty,
}
