FROM golang:1.14-alpine

# 作業ディレクトリの作成・設定
WORKDIR /techtrain-CA

# Go Modules を有効化
ENV GO111MODULE=on

COPY go.mod .
# COPY go.sum .

# go.mod 内のパッケージをダウンロード
RUN go mod download

EXPOSE 8080
