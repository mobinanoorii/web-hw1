# web-hw1
## Nginx inital setup 
run these commands on the root project directory
```bash
docker network create --driver bridge Nginx
docker run --name docker-nginx -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx
```