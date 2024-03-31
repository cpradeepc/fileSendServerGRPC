


gen :
	protoc --go-grpc_out=. --go_out=. ./proto/*.proto
cls :
	rm proto/upload_grpc.pb.go proto/upload.pb.go
run_s :

run_c : 