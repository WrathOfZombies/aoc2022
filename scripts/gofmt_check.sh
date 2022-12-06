#!/usr/bin/env bash

# Lifted off of the Terraform repo (https://github.com/hashicorp/terraform/blob/main/scripts/gofmtcheck.sh)

# Check go fmt
echo "==> Checking that code complies with go fmt requirements..."
gofmt_files=$(go fmt ./...)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`go fmt\` to reformat code."
    exit 1
fi

exit 0