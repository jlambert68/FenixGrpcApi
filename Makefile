
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

compileServerAdminProto:
	@echo "Compile Fenix Admin proto file..."

 # generate the messages
# protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Fenix/fenixTestDataSyncServerGrpcApi && protoc --go_out=. fenixTestDataSyncServerAdminGrpcApi.proto

# generate the services
# protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Fenix/fenixTestDataSyncServerGrpcApi && protoc --go-grpc_out=. fenixTestDataSyncServerAdminGrpcApi.proto

compileClientEchoProto:
	@echo "Compile Client Echo proto file..."

 # generate the messages
# protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Client/fenixClientTestDataSyncServerGrpcApi/echo && protoc --go_out=. echo.proto

# generate the services
# protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
	cd Client/fenixClientTestDataSyncServerGrpcApi/echo && protoc --go-grpc_out=. echo.proto

compileTestCaseGUIServerProto_all: compileTestCaseGUIServerProto_go compileTestCaseGUIServerProto_csharp

compileTestCaseGUIServerProto_go:
	@echo "Compile TestCaseGUIServer proto file..."

 # generate the messages
# protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
#	cd FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi && protoc --go_out=. fenixTestCaseBuilderServerGrpcApi.proto
	protoc --go_out=FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/*.proto

# generate the services
# protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"
#	cd FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi && protoc --go-grpc_out=. fenixTestCaseBuilderServerGrpcApi.proto
	protoc --go-grpc_out=FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/*.proto

compileTestCaseGUIServerProto_csharp:
	protoc --csharp_out=FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/csharp_grpc_api FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/*.proto

################################################################

compileExecutionGUIServerProto_all: compileExecutionGUIServerProto_go compileExecutionServerProto_go compileExecutionServerClientProto_go

compileExecutionGUIServerProto_go:
	@echo "Compile ExecutionGUIServer proto file..."

 # generate the messages
	protoc --go_out=FenixExecutionServer/fenixExecutionServerClientGrpcApi FenixExecutionServer/fenixExecutionServerClientGrpcApi/*.proto

# generate the services
	protoc --go-grpc_out=FenixExecutionServer/fenixExecutionServerClientGrpcApi FenixExecutionServer/fenixExecutionServerClientGrpcApi/*.proto


compileExecutionServerProto_go:
	@echo "Compile ExecutionServer proto file..."

 # generate the messages
	protoc --go_out=FenixExecutionServer/fenixExecutionServerGrpcApi FenixExecutionServer/fenixExecutionServerGrpcApi/*.proto

# generate the services
	protoc --go-grpc_out=FenixExecutionServer/fenixExecutionServerGrpcApi FenixExecutionServer/fenixExecutionServerGrpcApi/*.proto

compileExecutionServerClientProto_go:
	@echo "Compile ExecutionServerClient proto file..."

 # generate the messages
	protoc --go_out=FenixExecutionServer/fenixExecutionServerGuiGrpcApi FenixExecutionServer/fenixExecutionServerGuiGrpcApi/*.proto

# generate the services
	protoc --go-grpc_out=FenixExecutionServer/fenixExecutionServerGuiGrpcApi FenixExecutionServer/fenixExecutionServerGuiGrpcApi/*.proto



