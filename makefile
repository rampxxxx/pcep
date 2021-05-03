LIST = pce/pce_main.go pcc/pcc_main.go
APPS = pce_main pcc_main

.PHONY: $(LIST)
all: $(LIST)

$(LIST):
	@echo $@
	go build $@

.PHONY: $(APPS)
clean: $(APPS)

$(APPS):
	rm $@


