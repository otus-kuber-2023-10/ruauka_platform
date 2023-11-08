build:
	docker build --tag=ruauka/kuberserver:latest .

run:
	docker run -d --rm --name kuberserver -p 8000:8000 ruauka/kuberserver:latest

remove:
	docker stop kuberserver && docker rmi ruauka/kuberserver:latest

push:
	docker push ruauka/kuberserver:latest