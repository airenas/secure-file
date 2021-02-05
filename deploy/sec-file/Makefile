#####################################################################################
dist_dir=$(CURDIR)/data
tools=encrypt-file check-decrypt-file
dist_tools_files=$(patsubst %, $(dist_dir)/%, $(tools))
main_dir=../../cmd
service=airenas/secure-file
version=0.1
go_build_cmd=CGO_ENABLED=0 go build -installsuffix -cgo
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

dbuild: build
	docker build -t $(tag) ./

dpush: dbuild
	docker push $(tag)
#####################################################################################
clean:
	rm -rf $(dist_dir)

.PHONY:
	clean build dbuild dpush
