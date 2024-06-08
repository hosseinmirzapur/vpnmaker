build:
	@go build -o vpnmaker

run: build
	@./vpnmaker

clean:
	@rm created.json
	@rm endip.sh
	@rm result.csv
	@rm vpnmaker
	@rm warp*.txt
	@rm yg_update
