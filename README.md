**Creating docker image**

docker build -t rpsslweb .

docker run  -d --name rpsslWebContainer -p 8080:8080 rpsslweb 

**Dependencies**

I used only mux for router
