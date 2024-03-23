#!/bin/bash
rm -rf /tmp/kaspid-temp

kaspid --devnet --appdir=/tmp/kaspid-temp --profile=6061 --loglevel=debug &
KASPID_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $KASPID_PID

wait $KASPID_PID
KASPID_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kaspid exit code: $KASPID_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPID_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
