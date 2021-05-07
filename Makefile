.PHONEY: protos protos_chord_node protos_chord_admin

protos: protos_chord_node protos_chord_admin

protos_chord_node:
		protoc -I. models/proto/chordNode/chordNode.proto --go_out=plugins=grpc:./models/proto/chordNode

protos_chord_admin:
		protoc -I. models/proto/chordAdmin/chordAdmin.proto --go_out=plugins=grpc:./models/proto/chordAdmin
