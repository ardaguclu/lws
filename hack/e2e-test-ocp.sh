#!/usr/bin/env bash

# Copyright 2025 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

export CWD=$(pwd)

function lws_deploy {
    if [ -z "$RELEASE_IMAGE_LATEST" ]; then
      echo "RELEASE_IMAGE_LATEST is empty"
      exit 1
    fi
    if [ -z "$KUBECONFIG" ]; then
        echo "KUBECONFIG is empty"
        exit 1
    fi
    if [ -z "$NAMESPACE" ]; then
        echo "NAMESPACE is empty"
        exit 1
    fi
    # take the domain name of the cluster
    REGISTRY=$(echo "$RELEASE_IMAGE_LATEST" | awk -F'/' '{print $1}')
    IMAGE_TAG=$REGISTRY/$NAMESPACE/pipeline:lws
    cd $CWD/config/manager && $KUSTOMIZE edit set image controller=$IMAGE_TAG

    # TODO: add and remove some resources. For example, remove internalcert and add certmanager and prometheus
    $KUSTOMIZE build $CWD/test/e2e/config | $KUBECTL apply --server-side -f -
}

# TODO: deploy cert manager
lws_deploy
$GINKGO --junit-report=junit.xml --output-dir=$ARTIFACTS -v $CWD/test/e2e/...