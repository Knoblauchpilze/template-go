
COLOR_HIGHLIGHT_BLUE=\e[1;36m
COLOR_HIGHLIGHT_GREEN=\e[1;32m
COLOR_CLEAR=\e[0m

all: setup install

install:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Building apps...$(COLOR_CLEAR)"
	cd cmd/cli && make install
	@echo "$(COLOR_HIGHLIGHT_GREEN)Success!$(COLOR_CLEAR)"

setup:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Setup project...$(COLOR_CLEAR)"
	@./scripts/setup.sh
	@echo "$(COLOR_HIGHLIGHT_GREEN)Success!$(COLOR_CLEAR)"

clean:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Cleaning apps...$(COLOR_CLEAR)"
	@rm -rf build bin
	@cd cmd/cli && make clean
	@echo "$(COLOR_HIGHLIGHT_GREEN)Success!$(COLOR_CLEAR)"

# https://stackoverflow.com/questions/6273608/how-to-pass-argument-to-makefile-from-command-line
run:
	@echo "$(COLOR_HIGHLIGHT_BLUE)Running $(filter-out $@,$(MAKECMDGOALS))...$(COLOR_CLEAR)"
	@cd bin && run.sh --app-name $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
