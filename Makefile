# global options
Q=@

# go options
TESTFLAGS  :=
GOFLAGS    :=

# changelog options
NUM_COMMITS := 2


# ------------------------------------------------------------------------------
#  test
.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	@echo
	$Qgo test $(GOFLAGS) $(TESTFLAGS)

# ------------------------------------------------------------------------------
#  changelog
.PHONY: changelog
changelog:
	@./$@.sh $(NUM_COMMITS)
	@cat $@.txt

# ------------------------------------------------------------------------------
#  install
.PHONY: install
install:
	$Qgo install

# ------------------------------------------------------------------------------
#  build
.PHONY: build
build:
	$Qgo build
