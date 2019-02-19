VEGETA ?= `which vegeta`

.PHONY: default
default: benchmark

.PHONY: benchmark
benchmark:
	@cat Vegetafile | $(VEGETA) attack -duration=1m  -rate=500 | $(VEGETA) report -reporter=text
