Mode ?= DEBUG
PORT_NUM := $(shell lsof -i tcp:8719 | grep space | awk -F ' ' '{print $$2}')
PROD_DIR = $(CODE_DIR)/deploy/space/back

run: FORCE
	go build
	mv atto out
	cp config.yaml out/config.yaml
	mkdir -p out/log
	bash -c "echo -n > out/log/atto.log"
	cd out;./atto

deploy: FORCE
ifneq ($(PORT_NUM),)
	kill -9 $(PORT_NUM)
endif
	go build
	mv space-backend $(PROD_DIR)
	cp config.deploy.yaml $(PROD_DIR)/config.yaml
	export GIN_MODE=release;cd $(PROD_DIR);./space-backend > gin.log &

FORCE: