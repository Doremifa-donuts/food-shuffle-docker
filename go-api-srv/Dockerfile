# ベースイメージ
FROM golang:1.23.2-bookworm

# aptでgitをインストール
RUN apt-get update && apt-get install -y git

# DLVをインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest
