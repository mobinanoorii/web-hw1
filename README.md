# web-hw1
run these commands on the root project directory
## Docker initial setup
create the network and build the images required
```bash
docker network create --driver bridge project-network
docker build -t gateway-server ./gateway
```
## Golang gateway server
the golang server will log the incoming requests.
to see the logs of the server u can use the docker command :
```bash
docker run -d --name gateway-server gateway-server
docker network connect project-network gateway-server --alias gateway-server
```
to see the logs of the server u can use the following command 
```bash
docker logs -f gateway-server
```
## Nginx inital setup 
start the Nginx container and connect it to the network created
```bash
docker run --name nginx -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx
```