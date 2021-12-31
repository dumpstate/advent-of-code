INPUT := ../.inputs/$(year)/$(day)/$(or $(input), input)

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
ifeq ($(year), 2020)
	cd 2020 && python3 day$(day).py $(INPUT) && cd ..
endif
ifeq ($(year), 2021)
	cd 2021 && go run day$(day)/main.go $(INPUT) && cd ..
endif
