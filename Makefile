http-over-vsock.eif: ./server/server.go Dockerfile go.mod go.sum
	docker build -t http-over-vsock:latest .
	nitro-cli build-enclave --docker-uri http-over-vsock:latest --output-file $@

.PHONY: enclave-run
enclave-run: http-over-vsock.eif
	nitro-cli run-enclave --cpu-count 2 --memory 2048 --eif-path http-over-vsock.eif --debug-mode --enclave-cid 16

.PHONY: enclave-console
enclave-console:
	ENCLAVE_ID=$$(nitro-cli describe-enclaves | jq -r ".[0].EnclaveID") && \
	  [ "$$ENCLAVE_ID" != "null" ] && nitro-cli console --enclave-id $${ENCLAVE_ID}

.PHONY: enclave-terminate
enclave-terminate:
	ENCLAVE_ID=$$(nitro-cli describe-enclaves | jq -r ".[0].EnclaveID") && \
	  [ "$$ENCLAVE_ID" != "null" ] && nitro-cli terminate-enclave --enclave-id $${ENCLAVE_ID}


multiple-forwarder.eif: ./server2/server2.go Dockerfile go.mod go.sum
	docker build -t multiple-forwarder:latest .
	nitro-cli build-enclave --docker-uri multiple-forwarder:latest --output-file $@

.PHONY: enclave2-run
enclave2-run: multiple-forwarder.eif
	nitro-cli run-enclave --cpu-count 2 --memory 2048 --eif-path multiple-forwarder.eif --debug-mode --enclave-cid 16

.PHONY: enclave2-console
enclave2-console:
	ENCLAVE_ID=$$(nitro-cli describe-enclaves | jq -r ".[0].EnclaveID") && \
	  [ "$$ENCLAVE_ID" != "null" ] && nitro-cli console --enclave-id $${ENCLAVE_ID}

.PHONY: enclave2-terminate
enclave2-terminate:
	ENCLAVE_ID=$$(nitro-cli describe-enclaves | jq -r ".[0].EnclaveID") && \
	  [ "$$ENCLAVE_ID" != "null" ] && nitro-cli terminate-enclave --enclave-id $${ENCLAVE_ID}
