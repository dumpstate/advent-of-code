INPUT := ../.inputs/$(year)/$(day)/$(or $(input), input)
TMP_DIR := /tmp

run:
ifndef year
	$(error year is undefined)
endif
ifndef day
	$(error day is undefined)
endif
	[ -f "./.inputs/$(year)/$(day)/input" ] &&\
		echo "Input already downloaded." ||\
		./fetch_input.sh $(year) $(day)
ifeq ($(year), 2019)
	cd 2019 && guile -e main -l common.scm -s day$(day).scm $(INPUT) && cd ..
endif
ifeq ($(year), 2020)
	cd 2020 && python3 day$(day).py $(INPUT) && cd ..
endif
ifeq ($(year), 2021)
	cd 2021 && go run day$(day)/main.go $(INPUT) && cd ..
endif
ifeq ($(year), 2022)
	cd 2022 && scala ./day$(day).scala $(INPUT) && cd ..
endif
ifeq ($(year), 2023)
	cd 2023 && npm run build && node ./build/day$(day).js $(INPUT) && cd ..
endif
ifeq ($(year), 2024)
	cd 2024 && \
	rm -rf $(TMP_DIR)/day$(day) && \
	rustc day$(day).rs -o $(TMP_DIR)/day$(day) && \
	$(TMP_DIR)/day$(day) $(INPUT) && \
	cd ..
endif
