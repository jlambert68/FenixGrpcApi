
# cat -e -t -v Makefile

#.DEFAULT_GOAL := buildAndRun
# Migration steps for new Golang gRPC installation
# https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable


compileAllProto: compileClientProto compileServerProto


compileClientProto:
	@echo "Compile Client proto file..."

 # generate the messages
# protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Client/fenixClientTestDataSyncServerGrpcApi && protoc --go_out=. fenixClientTestDataSyncServerGrpcApi.proto

# generate the services
# protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Client/fenixClientTestDataSyncServerGrpcApi && protoc --go-grpc_out=. fenixClientTestDataSyncServerGrpcApi.proto

compileServerProto:
	@echo "Compile Fenix proto file..."

 # generate the messages
# protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Fenix/fenixTestDataSyncServerGrpcApi && protoc --go_out=. fenixTestDataSyncServerGrpcApi.proto

# generate the services
# protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Fenix/fenixTestDataSyncServerGrpcApi && protoc --go-grpc_out=. fenixTestDataSyncServerGrpcApi.proto

compileClientEchoProto:
	@echo "Compile Client Echo proto file..."

 # generate the messages
# protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Client/fenixClientTestDataSyncServerGrpcApi && protoc --go_out=. echo.proto

# generate the services
# protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Client/fenixClientTestDataSyncServerGrpcApi && protoc --go-grpc_out=. echo.proto
