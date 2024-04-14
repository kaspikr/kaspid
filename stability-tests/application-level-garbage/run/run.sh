#!/bin/bash
rm -rf /tmp/kaspid-temp

kaspid --devnet --appdir=/tmp/kaspid-temp --profile=6061 --loglevel=debug &
KASPID_PID=$!
KASPID_KILLED=0
function killKaspidIfNotKilled() {
    if [ $KASPID_KILLED -eq 0 ]; then
      kill $KASPID_PID
    fi
}
trap "killKaspidIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:25611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASPID_PID

wait $KASPID_PID
KASPID_KILLED=1
KASPID_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kaspid exit code: $KASPID_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPID_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
