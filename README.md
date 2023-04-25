# web-hw1
## Docker initial setup
create the network and build the images required
```bash
docker network create --driver bridge project-network
```
## Nginx inital setup 
run these commands on the root project directory
```bash
docker run --name nginx -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx
docker network connect project-network docker-nginx --alias nginx
```