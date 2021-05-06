.PHONEY: protos

protos:
		protoc -I models/proto/chordNode models/proto/chordNode/chordNode.proto --go_out=plugins=grpc:./models/proto/chordNode
		protoc -I models/proto/chordAdmin models/proto/chordAdmin/chordAdmin.proto --go_out=plugins=grpc:./models/proto/chordAdmin