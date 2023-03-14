export PROJECT_DIR=$(PWD)
export SCRIPTS_FOLDER=$(PROJECT_DIR)/scripts

docker-build:
	@ sh $${SCRIPTS_FOLDER}/build.sh


docker-export:
	@ sh $${SCRIPTS_FOLDER}/export.sh


docker-run:
	@ sh $${SCRIPTS_FOLDER}/run.sh

docker-execute-all:
	@ sh $${SCRIPTS_FOLDER}/build.sh && sh $${SCRIPTS_FOLDER}/export.sh && sh $${SCRIPTS_FOLDER}/run.sh


$(protoImage)
proto-gen:
	@ sh $${SCRIPTS_FOLDER}/proto-gen.sh
.PHONY: proto-gen