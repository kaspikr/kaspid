#!/bin/bash

APPDIR=/tmp/kaspid-temp
KASPID_RPC_PORT=29587

rm -rf "${APPDIR}"

kaspid --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${KASPID_RPC_PORT}" --profile=6061 &
KASPID_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${KASPID_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $KASPID_PID

wait $KASPID_PID
KASPID_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kaspid exit code: $KASPID_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPID_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
