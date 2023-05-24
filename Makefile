.PHONY: build
build:
	go build github.com/liouk/gh-stats/cmd/gh-stats

.PHONY: examples
examples: build
	./gh-stats all --template examples/basic.tmpl --output examples/basic.md --template-extras examples/basic.json

.PHONY: install
install: build
	cp gh-stats ${HOME}/go/bin/

.PHONY: clean
clean:
	rm gh-stats
