# book store api

> A simple REST API built

* Mux package/dependencie
```
go get -u github.com/gorilla/mux
```

## Requirements

To be able to show the desired features of curl this REST API must match a few
requirements:

* [x] `GET /api/books` returns list of all books as JSON
* [x] `GET /api/books/{id}` returns details of specific book as JSON
* [x] `POST /api/books` accepts a new book to be added
* [x] `PUT /api/books/{id}` updates current book
* [x] `DELETE /api/books/{id}` deletes book

### Data Types

A book object should look like this:
```json
{
  "id": "someid",
  "isbn": "isbnCOde",
  "title": "titleOfTheBook",
  "Author": {
    "firstnam":"Petros",
    "lastname":"Trak"
  }
}
```

### Persistence

There is no persistence, a temporary in-mem story is fine.
