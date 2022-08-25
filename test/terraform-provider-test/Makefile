default: testacc

# Run acceptance tests
.PHONY: testacc

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

install:
	go build
	mkdir -p ../infra-test/terraform.d/plugins/local/local/test/1.0.0/linux_amd64
	mv ./terraform-provider-test ../infra-test/terraform.d/plugins/local/local/test/1.0.0/linux_amd64/terraform-provider-test_v1.0.0
	cd ../infra-test && rm -f .terraform.lock.hcl && terraform init
