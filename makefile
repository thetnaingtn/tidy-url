build:
	docker build -f zarf/docker/Dockerfile -t tidyurl:1.0 .

run:
	docker container run --name tidyurl-app -p 8000:5000 -d tidyurl:1.0

fly-deploy:
	flyctl deploy --dockerfile zarf/docker/Dockerfile
