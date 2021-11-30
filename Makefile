run:
ifndef year
	$(error year is undefined)
endif
ifndef day
	$(error day is undefined)
endif
ifeq ($(year), 2021)
	[ -f "./.inputs/$(year)/$(day)/input" ] &&\
		echo "Input already downloaded." ||\
		./fetch_input.sh $(year) $(day)

	go run $(year)/day$(day)/main.go
endif
