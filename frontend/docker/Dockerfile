# Frontend Dockerfile for Next.js development

FROM node:24-alpine

# pnpmをグローバルインストール
RUN npm install -g pnpm@8

# 作業ディレクトリを設定
WORKDIR /frontend

# package.jsonとpnpm-lock.yamlをコピーして依存関係をインストール
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --no-frozen-lockfile

# アプリケーションのソースコードをコピー
COPY . .

# ポート3000を公開
EXPOSE 3000
