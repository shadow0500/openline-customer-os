
docker build -t api-gateway-lambda .
docker build --platform=linux/amd64 -t api-gateway-lambda .
docker tag api-gateway-lambda:latest 769325097132.dkr.ecr.eu-west-2.amazonaws.com/api-gateway-lambda:latest
docker push 769325097132.dkr.ecr.eu-west-2.amazonaws.com/api-gateway-lambda:latest

docker run -e RDS_HOST=host.docker.internal -e RDS_PORT=5432 -e RDS_USERNAME=postgres -e RDS_PASSWORD=password -e RDS_DATABASE=openline -e TARGET_API_URL=http://host.docker.internal:10000 -p 9000:8080 api-gateway-lambda
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"headers":{"X-OPENLINE-TENANT-KEY": "10fc4ede-75a8-4e68-9e8c-98b2f3bbe2b3"}}'

aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 769325097132.dkr.ecr.eu-west-1.amazonaws.com                
Login Succeeded