#!/usr/bin/env bash

set -euo pipefail

yd=$(date +"%Y %d" | sed -e 's/ 0/ /')
year=$(awk '{print $1}' <<< "$yd")
day=$(awk '{print $2}' <<< "$yd")

BASEDIR=$(dirname "${BASH_SOURCE[0]}")
sed -e "s/ay0/ay${day}/g" "${BASEDIR}/aoc${year}/day0.txt" > "${BASEDIR}/aoc${year}/day${day}.go"
sed -e "s/ay0/ay${day}/g" "${BASEDIR}/aoc${year}/day0_test.txt" > "${BASEDIR}/aoc${year}/day${day}_test.go"

if [[ ! -v AOC_SESSION_TOKEN ]]; then
  echo "Environment variable AOC_SESSION_TOKEN must be set to download input"
  exit 0
fi

curl -b "session=${AOC_SESSION_TOKEN}" -H "User-Agent: ghiloni@gmail.com/$(hostname)" \
  "https://adventofcode.com/${year}/day/${day}/input" > "${BASEDIR}/aoc${year}/testdata/day${day}.txt"
