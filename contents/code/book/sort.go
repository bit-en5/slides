package book

import "sort"

//START1 OMIT

func SortByTitle(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})

	return books
}

//END1 OMIT
//START2 OMIT

func SortByYear(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Year < books[j].Year
	})

	return books
}

//END2 OMIT
//START3 OMIT

func SortByPages(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Pages < books[j].Pages
	})

	return books
}

//END3 OMIT
//START4 OMIT

func SortByAuthor(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Author < books[j].Author
	})

	return books
}

//END4 OMIT
