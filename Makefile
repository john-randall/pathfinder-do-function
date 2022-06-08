invoke:
	doctl serverless functions invoke sample/pathfinder -P params.json |jq

deploy:
	doctl serverless deploy .

test:
	find . -name go.mod -execdir go test ./... \;