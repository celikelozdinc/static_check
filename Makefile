GP := $(shell go env GOPATH)
STATICCHECK := $(GP)/bin/staticcheck
ERR := "No staticcheck found, exiting..."
PKGS := ./...
LOG := $(shell mktemp -t staticcheck.XXXXX)

.PHONY: static-check
static-check:
	@which $(STATICCHECK) || (echo $(ERR) ; exit 1;)
	@echo "staticcheck : $(shell which $(STATICCHECK))"
	@echo "logs : $(LOG)" 
	@$(STATICCHECK) $(PKGS) > $(LOG) || true
	@[ ! -s "$(LOG)" ] || (echo "staticcheck failed, report is below :" ; cat $(LOG))