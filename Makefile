.PHONY: deploy clean get/tag

get/tag:
	$(eval TAG := $(shell ./get_tag.sh))

$(TAG):
	mkdir -p tags/$(TAG)

deploy: get/tag $(TAG) $(GOPATH)/bin/gox $(GOPATH)/bin/ghr
	gox -output "tags/$(TAG)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	ghr $(TAG) tags/$(TAG)

clean:
	rm -rf tags/* main

$(GOPATH)/bin/gox:
	go get github.com/mitchellh/gox

$(GOPATH)/bin/ghr:
	go get github.com/tcnksm/ghr

main:
	go build -o main