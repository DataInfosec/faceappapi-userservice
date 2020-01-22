build:
	# protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=.  proto/user/user.proto
	sudo docker build -t faceappapi-useruservice .

run: 
	sudo docker rm -f userservices
	sudo docker run -it -p 8812:8807 -e COUCHBASE_HOST=http://192.168.0.101:30886/ -e COUCHBASE_BUCKET=default -e COUCHBASE_USERNAME=elvis -e COUCHBASE_PASSWORD=password --network="bridge" --name userservices faceappapi-userservices
	# sudo docker run -d -p 1111:8805  --network="bridge" --name userservices2 faceappapi-userservices

go:
	env DB_URL=elvis:0gbunike@cluster0-oqk2c.mongodb.net/test DB_NAME=faceappdb DB_TABLE=locations  SENDGRID_API_KEY=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 go run *.go

network:
	sudo docker network ls

clean:
	go clean -modcache

fixuuid:
	go mod edit -replace=github.com/satori/go.uuid@v1.2.0=github.com/satori/go.uuid@master
	go mod tidy
	go build

creatingPrimaryIndex:
	CREATE PRIMARY INDEX ON default;

#DB_URL := "mongodb://elvis:Ogbunike@ds211259.mlab.com:11259/faceappdb?authMechanism=SCRAM-SHA-1"