#!/usr/bin/env bash

set -eu -o pipefail

if ! hash terraform > /dev/null 2>&1; then
    echo "Terraform is missing from your machine, please install it." >&2
    exit 1
fi

if [ "$#" -ne 1 ]; then
    echo "Don't forget team name" >&2
    exit 1
fi

export TEAM="$1"

mkdir "$TEAM"
cat "geegle.tf_" | envsubst > "$TEAM"/geegle.tf
cd "$TEAM"

terraform init >&2

(echo "yes" | terraform apply >&2) || exit $?

terraform output priv_ip
cd - >&2

