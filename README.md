**Creating docker image**

docker build -t rpsslweb .

docker run  -d --name rpsslWebContainer -p 8080:8080 rpsslweb 

**Dependencies**

I used only mux for router

**Endpoint**

https://rpssl-web.herokuapp.com

You can use above url at https://codechallenge.boohma.com/

**CHOICES**

Request:

curl --location --request GET 'https://rpssl-web.herokuapp.com/choices' \
--header 'Content-Type: application/json' \
--data-raw '{
"player": 1
}'

Response:

[
{
"id": 1,
"name": "rock"
},
{
"id": 2,
"name": "paper"
},
{
"id": 3,
"name": "scissors"
},
{
"id": 4,
"name": "lizard"
},
{
"id": 5,
"name": "spock"
}
]

**CHOICE**

Request:

curl --location --request GET 'https://rpssl-web.herokuapp.com/choice'

Response:

{
"id": 2,
"name": "paper"
}

**PLAY**

Request:

curl --location --request POST 'https://rpssl-web.herokuapp.com/play' \
--header 'Content-Type: application/json' \
--data-raw '{
"player": 1
}'

Response:

{
"results": "lose",
"player": 1,
"computer": 5
}
