**how to connect the aplication and db running in different containers**

#create a docker network

 `docker network create tinyurl_network`

#run postgres docker container on this network

 `docker run --name postgress-container -e POSTGRES_PASSWORD=admin -p 5432:5432 --network tinyurl_network -v postgress_data:/var/lib/postgresql/data -d postgres`

#update the db_host in .env file to 

 `host.docker.internal`

#run application docker container on this network

 `docker run -itd -p 3000:3000 -v ./.env:/app/.env --network tinyurl_network tinyurl:latest`