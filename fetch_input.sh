#!/usr/bin/env bash

set -euo pipefail

BASE_URL="https://adventofcode.com"
OUT_DIR="./.inputs"

usage () {
	printf "Usage:\n"
	printf "\t1. Make sure to place session cookie into .cookie file.\n"
	printf "\t2. fetch_input.sh\t<year>\t<day>\n"
}

main () {
	local year=$1
	local day=$2

	if [ -z "$year" ]; then
		usage
		exit 1
	fi

	if [ -z "$day" ]; then
		usage
		exit 1
	fi

	local session_cookie=$(cat ./.cookie)

	if [ -z "$session_cookie" ]; then
		printf "Empty session cookie (.cookie)"
		exit 1
	fi

	mkdir -p "${OUT_DIR}/${year}/${day}"
	curl -H "Cookie: session=${session_cookie}" \
		"${BASE_URL}/${year}/day/${day}/input" > "${OUT_DIR}/${year}/${day}/input"
}

main $@
