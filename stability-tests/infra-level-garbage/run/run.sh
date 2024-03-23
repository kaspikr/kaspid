#!/bin/bash
rm -rf /tmp/kaspid-temp

kaspid --devnet --appdir=/tmp/kaspid-temp --profile=6061 &
KASPID_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASPID_PID

wait $KASPID_PID
KASPID_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kaspid exit code: $KASPID_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPID_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
