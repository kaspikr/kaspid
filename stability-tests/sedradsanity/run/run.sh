#!/bin/bash
kaspidsanity --command-list-file ./commands-list --profile=7000
TEST_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ]; then
  echo "kaspidsanity test: PASSED"
  exit 0
fi
echo "kaspidsanity test: FAILED"
exit 1
