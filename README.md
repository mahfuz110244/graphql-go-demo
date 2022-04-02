# github.com/mahfuz110244/graphql-go-demo
GraphQL demo using Golang and MySQL & PostgreSQL

## Setup
```
go run github.com/99designs/gqlgen gqlgen
go get github.com/99designs/gqlgen@v0.17.2
go get github.com/99designs/gqlgen/internal/imports@v0.17.2
```
## Run
###
Create MySQL graphql_demo using below docker command
```
docker-compose up --build
```

### Run the server
```
go run .
```

### GraphQL playground
```
http://localhost:8080
```

## Query

### Create two authors John and Karim
```
mutation creatAuthor {
  creatAuthor(name: "John", biography: "good writer") {
    id
    name
    biography
  }
}

mutation creatAuthor {
  creatAuthor(name: "Karim", biography: "good writer") {
    id
    name
    biography
  }
}
```

### Create two books with author John and one book with author Karim
```
mutation createBook {
  createBook(title: "Demo Book 1 John", price: 1000, isbn_no: "BK98969886", author:"1"){
    id
    title
    price
    isbn_no
    authors{
      id
      name
    	biography
    }
  }
}

mutation createBook {
  createBook(title: "Demo Book 2 John", price: 500, isbn_no: "BK98969887", author:"1"){
    id
    title
    price
    isbn_no
    authors{
      id
      name
    	biography
    }
  }
}

mutation createBook {
  createBook(title: "Demo Book 1 Karim", price: 500, isbn_no: "BK98969886", author:"2"){
    id
    title
    price
    isbn_no
    authors{
      id
      name
    	biography
    }
  }
}
```

### Get all the books of John
```
query GetAllTheBooksOfJohn{
  authors(name: "John") {
    books{
      id
      title
      price
      isbn_no
    }
  }
}
```

### Get all the books of all authors
```
query GetAllBooksWithAuthors {
  books {
    id
    title
    price
    isbn_no
    authors {
      id
      name
      biography
		} 
  }
}
```


## Challenges
As I have no experience in GraphQL previously, the first challenge for me is understanding GraphQL and its structure. Second challenge is generated graph modules using
```
github.com/99designs/gqlgen
```
this library in golang. Last challenge I have faced is generated the query result for the GetAllTheBooksOfJohn. I need to make a different struct format for this.
