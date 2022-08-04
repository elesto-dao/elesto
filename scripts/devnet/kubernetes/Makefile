INIT_IMAGE_TAG ?= elestod-init:local
NODE_IMAGE_TAG ?= us-central1-docker.pkg.dev/elesto/elesto/elestod:2.0.0-rc2-snapshots

k3d-start:
	k3d cluster create --config k3d.yaml --wait || k3d cluster list | grep elesto
.PHONY: k3d-start

build-init:
	docker build -t $(INIT_IMAGE_TAG) ./0-init/
.PHONY: build-init

import-images: build-init
	k3d image import $(INIT_IMAGE_TAG) $(NODE_IMAGE_TAG) --cluster elesto
.PHONY: import-init

k3d-deploy: import-images
	kubectl apply -k . --wait=true
.PHONY: k3d-deploy

k3d-stop k3d-delete:
	k3d cluster delete --config k3d.yaml
.PHONY:k3d-stop k3d-delete

port-forward:
	kubectl port-forward --namespace local svc/gensis 26657
.PHONY:port-forward