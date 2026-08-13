#!/usr/bin/env bash
# Run every example and write TEST_REPORT.md in this directory.

cd "$(dirname "$0")" || exit 1

OUT="TEST_REPORT.md"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

# Longer than the 5 minute HTTP client timeout in auth/login.go, so a slow
# agent stream reports its own error instead of being killed here.
TIMEOUT=360

if ! command -v go > /dev/null; then
  echo "go is not on PATH."
  exit 1
fi

if [ ! -f ".env" ]; then
  echo "No .env in this directory. Copy .env.example to .env first."
  exit 1
fi

# Compile everything once so the per-example timeout measures run time,
# not build time. A build failure here stops the run.
if ! go build ./... 2> "$TMP/build"; then
  echo "go build ./... failed:"
  cat "$TMP/build"
  exit 1
fi

# Every runnable example is a package with its own main.go.
PKGS=$(find . -name 'main.go' | sed 's|^\./||;s|/main.go$||' | sort)

# agent/update and agent/delete take an agent key as a second argument.
# Give each of them a throwaway agent so they never touch PIPESHUB_AGENT_KEY.
scratch_agent_key() {
  timeout "$TIMEOUT" go run ./agent/create .env 2> /dev/null \
    | sed -n 's/^created agent key: //p' | head -1
}

PASS=0
FAIL=0
GROUP=""

echo "# Example test report" > "$OUT"
echo "" >> "$OUT"
echo "$(date '+%Y-%m-%d %H:%M:%S')" >> "$OUT"
echo "" >> "$OUT"

for p in $PKGS; do
  printf '%-46s ' "$p"

  ARGS=".env"
  CLEANUP=""
  NOKEY=""
  case "$p" in
    agent/update|agent/delete)
      key=$(scratch_agent_key)
      [ -z "$key" ] && NOKEY=1
      ARGS=".env $key"
      # agent/delete removes its own agent; agent/update leaves it behind.
      [ "$p" = "agent/update" ] && CLEANUP="$key"
      ;;
  esac

  if [ -n "$NOKEY" ]; then
    echo "could not create a scratch agent (go run ./agent/create failed)" > "$TMP/err"
    code=1
  else
    timeout "$TIMEOUT" go run "./$p" $ARGS > "$TMP/out" 2> "$TMP/err"
    code=$?
  fi

  if [ -n "$CLEANUP" ]; then
    timeout "$TIMEOUT" go run ./agent/delete .env "$CLEANUP" > /dev/null 2>&1
  fi

  # Every example ends with log.Fatal on failure, so the exit code is the
  # verdict.
  if [ $code -eq 0 ]; then
    status="PASS"
    PASS=$((PASS + 1))
  else
    status="FAIL"
    FAIL=$((FAIL + 1))
  fi
  echo "$status"

  g=$(dirname "$p")
  if [ "$g" != "$GROUP" ]; then
    echo "" >> "$OUT"
    echo "## $g" >> "$OUT"
    echo "" >> "$OUT"
    GROUP="$g"
  fi

  if [ "$status" = "PASS" ]; then
    echo "- [x] $(basename "$p") — PASS" >> "$OUT"
  else
    echo "- [ ] $(basename "$p") — FAIL" >> "$OUT"
    echo "" >> "$OUT"
    echo '  ```' >> "$OUT"
    if [ $code -eq 124 ]; then
      echo "  timed out after ${TIMEOUT}s" >> "$OUT"
    fi
    grep -v '^[[:space:]]*$' "$TMP/err" | tail -n 20 | sed 's/^/  /' >> "$OUT"
    echo '  ```' >> "$OUT"
    echo "" >> "$OUT"
  fi
done

echo "" >> "$OUT"
echo "**$PASS passed, $FAIL failed**" >> "$OUT"

echo ""
echo "$PASS passed, $FAIL failed"
echo "report: $OUT"

# Non-zero exit when anything failed, so scripted callers see the result.
[ "$FAIL" -eq 0 ] || exit 1
