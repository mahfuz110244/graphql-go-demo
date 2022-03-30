# github.com/mahfuz110244/graphql-go-demo
GraphQL demo using Golang and MySQL & PostgreSQL

## Query
```
mutation creatAuthor {
  creatAuthor(name: "John", biography: "good writer") {
    id
    name
    biography
  }
}

mutation createBook {
  createBook(title: "Demo Book 1", price: 1000, isbn_no: "BK99y86886", author:"1"){
    id
    title
    price
    isbn_no
    author{
      id
      name
    	biography
    }
  }
}

query GetAllBooksByAuthorName{
  authors(name: "John") {
    books{
      id
      title
      price
      isbn_no
      author{
        id
        name
        biography
      }
    }
  }
}


query GetAllAuthors {
  allAuthors {
    id
    name
    biography
  }
}

query GetAllBooks {
  allBooks {
   	id
    title
    price
    isbn_no
    author{
      id
      name
      biography
    }
  }
}
```