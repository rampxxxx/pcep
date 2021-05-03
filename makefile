DIRS= pce pcc pceplib
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


.PHONY: $(DIRS)
test: $(DIRS)

$(DIRS):
	cd $@;go test ; cd -
