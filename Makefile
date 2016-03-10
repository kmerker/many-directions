build:
	go build .

image:
	docker build -t mnoorali/many-directions:0.0.3 .

push:
	docker push mnoorali/many-directions:0.0.3

