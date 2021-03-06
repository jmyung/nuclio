# Copyright 2017 The Nuclio Authors.
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

#
# Build stage: builds the autoscaler binary
#

FROM golang:1.10-alpine3.7 as builder

RUN apk --update --no-cache add \
    git \
    gcc \
    musl-dev

WORKDIR /go/src/github.com/nuclio/nuclio

COPY . .

RUN go get github.com/v3io/scaler-types \
    && go get github.com/v3io/v3io-go-http \
    && go get github.com/nuclio/logger \
    && go get github.com/nuclio/nuclio-sdk-go

RUN cd $GOPATH/src/github.com/v3io/scaler-types \
    && git checkout v1.3.0 \
    && cd /go/src/github.com/nuclio/nuclio

RUN mkdir -p /home/nuclio/bin \
            && GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -installsuffix cgo -ldflags="-s -w" \
            -o /home/nuclio/plugins/plugin.so pkg/platform/kube/resourcescaler/resourcescaler.go \
            && rm -rf /go/src/github.com/nuclio/nuclio

FROM alpine:3.7

WORKDIR /home/nuclio

COPY --from=builder /home/nuclio/plugins/plugin.so /home/nuclio/plugins/plugin.so
COPY --from=quay.io/v3io/autoscaler:v0.1.2 /home/v3io/bin/autoscaler /home/nuclio/autoscaler

CMD ["/home/nuclio/autoscaler"]
