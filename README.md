**Creating docker image**

docker build -t rpsslweb .

docker run  -d --name rpsslWebContainer -p 8080:8080 rpsslweb 

**Dependencies**

I used only mux for router

**Endpoint**

https://rpssl-web.herokuapp.com

You can use above url at https://codechallenge.boohma.com/
