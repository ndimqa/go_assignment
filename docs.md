

---
Documentation for Nuran Dinmukamed API in Golang
---
---
## This API takes data from moked static data
---

## Register Request
http://127.0.0.1:8181/register (POST)
## Body
```
{
    "username": qwerty,
    "password": 123456,
    "mail": qwerty@mail.ru
}
```
## Response
```
{
    "message": "Status OK"
}
```
---
## Notes
- [] Need to add hashing method for passwords
- [] Need to add regex to check password vulnerability and mail validation 
---



## Filter Request
http://127.0.0.1:8181/filter?sortBy={}?byAttribute={} (GET)
## Body
```
{}
```
## Response
```
[
    {
        "ID": 2,
        "Title": "Don Quixote",
        "Author": "Miguel de Cervantes",
        "Price": 10.99,
        "Rating": 0,
        "AmountOfRating": 0,
        "ListOfRatings": []
    },
    {
        "ID": 1,
        "Title": "War and Peace",
        "Author": " Leo Tolstoy",
        "Price": 12.99,
        "Rating": 0,
        "AmountOfRating": 0,
        "ListOfRatings": []
    },
    {
        "ID": 0,
        "Title": "The Great Gatsby",
        "Author": "F. Scott Fitzgerald",
        "Price": 20.99,
        "Rating": 0,
        "AmountOfRating": 0,
        "ListOfRatings": []
    }
]
```
---
## Notes
---

## GiveRating Request
http://127.0.0.1:8181/giverating (PUT)
## Body
```
{
    "id": 0,
    "rating": 3
}
```
## Response
```
{
    "message": "Status OK"
}
```
---
## Notes
---

## Login Request
http://127.0.0.1:8181/login (GET)
## Body
```
{
    "username": qwerty,
    "password": 123456,
}
```
## Response
```
{
    "message": "Status OK"
}
```
---
## Notes
- [] Need to change password checking from raw password to hash checking
---

## Search Request
http://127.0.0.1:8181/searh?Srch={} (GET)
## Body
```
{}
```
## Response
```
{
    "ID": 2,
    "Title": "Don Quixote",
    "Author": "Miguel de Cervantes",
    "Price": 10.99,
    "Rating": 0,
    "AmountOfRating": 0,
    "ListOfRatings": []
}
```
---
## Notes
---