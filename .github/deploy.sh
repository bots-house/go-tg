#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

IFS=', ' read -r -a deploy_hooks <<< $1

for hook in "${deploy_hooks[@]}"; do
    http_code=$(curl --silent --output /dev/null --write-out '%{http_code}\n' -X POST $hook)
    if [[ $http_code -eq 204 ]]; then
        echo '👌'
    else
        echo "🤬 $http_code"
    fi
done