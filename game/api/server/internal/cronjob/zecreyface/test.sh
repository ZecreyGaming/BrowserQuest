#!/bin/bash

set -e
set -v

export TESTING_ACCOUNT_NAME="amber1"
export TESTING_ACCOUNT_SEED="ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"

go test -v zecreyface_test.go