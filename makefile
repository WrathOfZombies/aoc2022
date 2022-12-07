perms: 
	@echo "Updating executable permissions on scripts"
	@find $(CURDIR)/scripts -name "*.sh" -exec chmod u+x {} \;
	@echo "Done."

pretty: perms
	@echo "Formatting files"
	@"$(CURDIR)/scripts/gofmt_check.sh"
	@echo "Done."

start:
	@go run "$(CURDIR)/cmd/haze/main.go"