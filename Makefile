.PHONEY: protos

protos:
		protoc -I models/proto/ models/proto/chord.proto --go_out=plugins=grpc:./models/proto/