#!/bin/bash
set -euxo pipefail

SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
source "$SCRIPT_DIR"/common/cmd.sh
source "$SCRIPT_DIR"/internal/accounts.sh

cleanup() {
  cleanup_init_network_sh
  exit
}
trap cleanup INT TERM EXIT

VALIDATORS=1
VAL_ACCOUNTS_DIR="networks/nolus/val-accounts"
POSITIONAL=()
ARTIFACT_BIN=""
ARTIFACT_SCRIPTS=""

NATIVE_CURRENCY="unolus"
VAL_TOKENS="1000000000""$NATIVE_CURRENCY"
VAL_STAKE="1000000""$NATIVE_CURRENCY"
CHAIN_ID="nolus-private"
SUSPEND_ADMIN=""
FAUCET_MNEMONIC=""
FAUCET_TOKENS="1000000""$NATIVE_CURRENCY"

while [[ $# -gt 0 ]]; do
  key="$1"

  case $key in

  -h | --help)
    printf \
    "Usage: %s
    [--artifact-bin <tar_gz_nolusd>]
    [--artifact-scripts <tar_gz_scripts>]
    [--chain_id <string>]
    [-v|--validators <number>]
    [--validator_accounts_dir <validator_accounts_dir>]
    [--currency <native_currency>]
    [--validator-tokens <tokens_for_val_genesis_accounts>]
    [--validator-stake <tokens_val_will_stake>]
    [-ips <ip_addrs>]
    [--suspend-admin <bech32address>]
    [--faucet-mnemonic <mnemonic_phrase>]
    [--faucet-tokens <initial_balance>]"
     "$0"
    exit 0
    ;;

   --artifact-bin)
    ARTIFACT_BIN="$2"
    shift
    shift
    ;;

   --artifact-scripts)
    ARTIFACT_SCRIPTS="$2"
    shift
    shift
    ;;
  
   --chain-id)
    CHAIN_ID="$2"
    shift
    shift
    ;;

   -v | --validators)
    VALIDATORS="$2"
    [ "$VALIDATORS" -gt 0 ] || {
      echo >&2 "validators must be a positive number"
      exit 1
    }
    shift
    shift
    ;;

   --validator_accounts_dir)
    VAL_ACCOUNTS_DIR="$2"
    shift
    shift
    ;;

  --currency)
    NATIVE_CURRENCY="$2"
    shift
    shift
    ;;

  --validator-tokens)
    VAL_TOKENS="$2"
    shift
    shift
    ;;

  --validator-stake)
    VAL_STAKE="$2"
    shift
    shift
    ;;

  --suspend-admin)
    SUSPEND_ADMIN="$2"
    shift
    shift
    ;;

  --faucet-mnemonic)
    FAUCET_MNEMONIC="$2"
    shift
    shift
    ;;
  --faucet-tokens)
    FAUCET_TOKENS="$2"
    shift
    shift
    ;;
  *) # unknown option
    POSITIONAL+=("$1") # save it in an array for later
    shift              # past argument
    ;;

  esac
done

__verify_mandatory() {
  local value="$1"
  local description="$2"

  if [[ -z "$value" ]]; then
    echo >&2 "$description was not set"
    exit 1
  fi
}

__add_faucet_account() {
  local mnemonic="$1"
  local amount="$2"

  local account_name="faucet"
  local tmp_faucet_dir
  tmp_faucet_dir="$(mktemp -d)"
  run_cmd "$tmp_faucet_dir" keys add --recover "$account_name" --keyring-backend test <<< "$mnemonic" 1>/dev/null
  local addr
  addr="$(run_cmd "$tmp_faucet_dir" keys show "$account_name" -a --keyring-backend test)"
  add_account "$addr" "$amount"
}


__verify_mandatory "$ARTIFACT_BIN" "Nolus binary actifact"
__verify_mandatory "$ARTIFACT_SCRIPTS" "Nolus scipts actifact"
__verify_mandatory "$SUSPEND_ADMIN" "Suspend admin"
__verify_mandatory "$FAUCET_MNEMONIC" "Faucet mnemonic"

# TBD open a few sample private investor accounts
# TBD open admin accounts, e.g. a treasury and a suspender
#  and pass them to init_network
accounts_spec=$(echo "[]" | __add_faucet_account "$FAUCET_MNEMONIC" "$FAUCET_TOKENS")

source "$SCRIPT_DIR"/internal/setup-validator-dev.sh
init_setup_validator_dev_sh "$SCRIPT_DIR" "$ARTIFACT_BIN" "$ARTIFACT_SCRIPTS"

source "$SCRIPT_DIR"/internal/init-network.sh
init_network "$VAL_ACCOUNTS_DIR" "$VALIDATORS" "$CHAIN_ID" "$NATIVE_CURRENCY" "$SUSPEND_ADMIN" "$VAL_TOKENS" "$VAL_STAKE" "$accounts_spec"
