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

## Challenges
As I have no experienced in GrapQL previously, so first challenge for me understand GraphQL and its structure. Second challenge is generated graph modules suing
```
github.com/99designs/gqlgen
```
this library in golang.