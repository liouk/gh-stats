.PHONY: build
build:
	go build github.com/liouk/gh-stats/cmd/gh-stats

.PHONY: clean
clean:
	rm gh-stats
