all: help

FILENAME := ip.txt

run: $(FILENAME)
	@echo "Running the test..."
	@/usr/bin/time -f "Test took time    : %E" go run cmd/test/main.go

$(FILENAME):
	@echo "Generating test file..."
	@/usr/bin/time -f "Time to generate  : %E" go run cmd/test-file-generator/main.go 1000000000 && \
	echo "File              : \"$(FILENAME)\" generated" && \
	echo "Size              : $$(du -h $(FILENAME) | awk '{print $$1}')" && \
	echo "Lines             : $$(wc -l $(FILENAME) | awk '{print $$1}')"

clean:
	@echo "Cleaning artifacts..."
	@rm -rf $(FILENAME)

help:
	@echo "Makefile commands:"
	@echo "  make               - Show this help message"
	@echo "  make run           - Run the test after generating the file if necessary"
	@echo "  make clean         - Clean the artifacts"
	@echo "  make help          - Show this help message"

.PHONY: all run clean help
