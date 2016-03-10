build:
	go build .

image:
	docker build -t mnoorali/many-directions:1.0.0 .

push:
	docker push mnoorali/many-directions:1.0.0

