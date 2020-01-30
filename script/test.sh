#!/usr/bin/env bash
set -ue

export ADMIN_PASSWORD="abcdefg"
export TEST_USER_ID="p2qfpb2gvxrzedu2"
export FIREBASE_CREDENTIALS="../firebase-service-key-dev.json"
export FIREBASE_APIKEY="AIzaSyAF9quvSMu9n3LMWBrXw_aO5LYwBzqT4Gw"
export GOOGLE_APPLICATION_CREDENTIALS="../gae-service-key-dev.json"

usage_exit() {
  echo "テストを実行します。GNU parallel が利用可能な環境では GNU parallel を利用します。"
  echo ""
  echo "Usage: $0 [-stv]"
  echo
  echo "# Options"
  echo "   -s: force serial mode"
  echo "   -t: dry run"
  echo "   -v: verbose"
  exit $1
}

go_test() {
  TARGET_PKG=$1
  $DRY_RUN go test -v -tags=testing -cover -coverprofile=$TARGET_PKG.out ./$TARGET_PKG
  OUTPUT_FILES+=(${TARGET_PKG}.out)
}

concat_outputs() {
  cat ${OUTPUT_FILES[@]} > coverage.txt
}

remove_outputs() {
  rm ${OUTPUT_FILES[@]}
}

#
# MAIN
# =================================

GO_TARGETS=(controller)
OUTPUT_FILES=()
USE_PARALLEL=0
MAX_PARALLEL=3
FORCE_SERIAL=0
VERBOSE=0
DRY_RUN=""

while getopts :shtv OPT
do
  case $OPT in
    s) FORCE_SERIAL=1
    ;;
    h) usage_exit 0
    ;;
    t) DRY_RUN=echo
    ;;
    v) VERBOSE=1
    ;;
    \?) usage_exit 1
    ;;
  esac
done

# Remove parsed arguments
shift $((OPTIND - 1))

if [ $VERBOSE -eq 1 ]; then
  set -x
fi

if type -p parallel; then
  echo '           _ __  __       _______   ____  __                          ____     __'
  echo ' _      __(_) /_/ /_     / ____/ | / / / / /  ____  ____ __________ _/ / /__  / /'
  echo '| | /| / / / __/ __ \   / / __/  |/ / / / /  / __ \/ __ `/ ___/ __ `/ / / _ \/ /'
  echo '| |/ |/ / / /_/ / / /  / /_/ / /|  / /_/ /  / /_/ / /_/ / /  / /_/ / / /  __/ /'
  echo '|__/|__/_/\__/_/ /_/   \____/_/ |_/\____/  / .___/\__,_/_/   \__,_/_/_/\___/_/'
  echo ''
  echo 'MAX_PARALLE='$MAX_PARALLEL
  USE_PARALLEL=1
fi

if [ $USE_PARALLEL -eq 1 ] && [ $FORCE_SERIAL -eq 0 ]; then
  echo "parallel mode"

  $DRY_RUN parallel -P $MAX_PARALLEL --halt-on-error 2 go test -v --tags=testing -cover -covermode=atomic -coverprofile={}.out ./{} ::: ${GO_TARGETS[@]}

  for pkg in ${GO_TARGETS[@]}; do
    OUTPUT_FILES+=($pkg.out)
  done
else
 echo "serial mode"

  # Run test with go
  for pkg in ${GO_TARGETS[@]}; do
    go_test $pkg
  done
fi

$DRY_RUN concat_outputs
$DRY_RUN remove_outputs
