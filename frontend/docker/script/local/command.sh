#!/bin/sh

# Frontend local development command

echo "Starting frontend development server..."

# 依存関係を最新化
pnpm install --no-frozen-lockfile

# Next.js開発サーバーを起動
pnpm dev