#!/bin/bash
rm -rf /tmp/kaspid-temp

NUM_CLIENTS=128
kaspid --devnet --appdir=/tmp/kaspid-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
KASPID_PID=$!
KASPID_KILLED=0
function killKaspidIfNotKilled() {
  if [ $KASPID_KILLED -eq 0 ]; then
    kill $KASPID_PID
  fi
}
trap "killKaspidIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $KASPID_PID

wait $KASPID_PID
KASPID_EXIT_CODE=$?
KASPID_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Kaspid exit code: $KASPID_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPID_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
