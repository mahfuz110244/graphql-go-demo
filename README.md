# github.com/mahfuz110244/graphql-go-demo
GraphQL demo using Golang and MySQL & PostgreSQL

## Query
```
mutation creatAuthor{
  creatAuthor(input: {firstName:"John",lastName:"Doe"}) {
  id 
  firstName
  lastName
	}
}

mutation creatAuthor{
  creatAuthor(input: {firstName:"John",lastName:"Doe"}) {
  id 
  firstName
  lastName
	}
}

query GetAllAuthors{
  allAuthors{
    id
    firstName
    lastName
  }
}

query GetAllBooks{
  allBooks{
    id
    title
    Author{
      id
    	firstName
    	lastName
    }
  }
}


type Book{
  id: ID!
  title: String!
  Author: Author!
}

input newBook{
  title: String!
  Author: ID!
}

type Author{
  id: ID!
  firstName:String!
  lastName:String!
}

input newAuthor{
  firstName:String!
  lastName:String!
}

type Query{
  bookByID(id: ID):Book
  allBooks:[Book]
  authorByID(id:ID):Author!
  allAuthors:[Author]!
}

type Mutation{
  createBook(input: newBook): Book!
  creatAuthor(input: newAuthor):Author!
}

```