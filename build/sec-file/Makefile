#####################################################################################
dist_dir=$(CURDIR)/data
tools=encrypt-file check-decrypt-file
dist_tools_files=$(patsubst %, $(dist_dir)/%, $(tools))
main_dir=../../cmd
service=airenas/secure-file
version=0.2
go_build_cmd=CGO_ENABLED=0 go build -ldflags "-X main.version=$(version).$(commit_count)"
commit_count=$(shell git rev-list --count HEAD)
#####################################################################################
$(dist_dir):
	mkdir -p $@
$(dist_dir)/%: $(main_dir)/% | $(dist_dir)
	cd $< && $(go_build_cmd) -o $@	
#####################################################################################
build: $(dist_tools_files)
########### DOCKER ##################################################################
tag=$(service):$(version).$(commit_count)
up-version:
	rm -f $(dist_dir)/.version
$(dist_dir)/.version:
	echo $(version).$(commit_count) > $@
Dockerfile: Dockerfile.in up-version $(dist_dir)/.version Makefile
	cat $< | VERSION=$(version).$(commit_count) envsubst > $@
#####################################################################################
dbuild: build Dockerfile
	docker build -t $(tag) ./

dpush: dbuild
	docker push $(tag)
#####################################################################################
clean:
	rm -rf $(dist_dir)

.PHONY:
	clean build dbuild dpush
