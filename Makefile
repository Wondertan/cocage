protos:
	cd proto && buf generate --template buf.gen.pulsar.yaml -o ../modules/
.PHONY: protos