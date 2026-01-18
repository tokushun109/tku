#!/bin/bash

set -e

echo "=== 🔧 asdfツールバージョン更新スクリプト ==="
echo ""

echo "🍺 Homebrewを更新中..."
echo "Homebrewパッケージリストを更新しています..."
brew update

echo "🔧 asdfを最新版に更新中..."
echo "asdf本体を更新しています..."
brew upgrade asdf
echo "✅ asdf更新完了"
echo ""

echo "📦 asdfプラグインを更新中..."
echo "すべてのプラグインを最新版に更新しています..."
asdf plugin update --all || echo "⚠️ 一部のプラグイン更新に失敗しました（継続します）"
echo "✅ プラグイン更新完了"
echo ""

# .tool-versionsファイルの存在確認
if [ ! -f .tool-versions ]; then
    echo "❌ エラー: .tool-versionsファイルが見つかりません"
    exit 1
fi

# asdfコマンドの存在確認
if ! command -v asdf &> /dev/null; then
    echo "❌ エラー: asdfコマンドが見つかりません。先にasdfをインストールしてください。"
    exit 1
fi

# 現在インストールされているバージョンを取得する関数
get_current_version() {
    local tool=$1
    # asdf currentの出力形式に対応（表形式の場合はVersionカラムから取得）
    local version_output=$(asdf current $tool 2>/dev/null)
    if [ -z "$version_output" ]; then
        echo "not-installed"
        return
    fi

    # 表形式かどうかをチェック（Name,Version,Sourceヘッダーがある場合）
    if echo "$version_output" | head -1 | grep -q "^Name.*Version.*Source"; then
        # 表形式の場合は2行目のVersionカラム（2番目のフィールド）を取得
        echo "$version_output" | sed -n '2p' | awk '{print $2}'
    else
        # 従来形式の場合は2番目のフィールドを取得
        echo "$version_output" | awk '{print $2}'
    fi
}

# .tool-versionsから指定されているバージョンを取得する関数
get_specified_version() {
    local tool=$1
    grep "^$tool " .tool-versions | awk '{print $2}' || echo ""
}

# 指定バージョンがインストール済みか確認する関数
is_version_installed() {
    local tool=$1
    local version=$2
    asdf where "$tool" "$version" >/dev/null 2>&1
}

# Node.js関連の追加パッケージをインストールする関数
install_node_packages() {
    echo "📦 Node.js追加パッケージをインストール中..."
    npm install -g pnpm
    npm install -g @anthropic-ai/claude-code
    echo "asdfのshimを再構築しています..."
    asdf reshim nodejs
    echo "✅ Node.js追加パッケージのインストール完了"
}

# 各ツールを処理
while IFS= read -r line || [ -n "$line" ]; do
    # コメント行と空行をスキップ
    if [[ $line =~ ^#.* ]] || [[ -z "${line// }" ]]; then
        continue
    fi

    # ツール名とバージョンを抽出
    tool=$(echo $line | awk '{print $1}')
    specified_version=$(echo $line | awk '{print $2}')

    if [ -z "$tool" ] || [ -z "$specified_version" ]; then
        continue
    fi

    echo "🔍 $tool を処理中..."

    # プラグインがインストールされているかチェック
    if ! asdf plugin list | grep -q "^$tool$"; then
        echo "  📥 プラグイン $tool がインストールされていません。プラグインを追加中..."
        asdf plugin add $tool
    fi

    # 現在のバージョンを取得
    current_version=$(get_current_version $tool)
    installed_flag=false
    if is_version_installed "$tool" "$specified_version"; then
        installed_flag=true
    fi

    echo "  📋 指定バージョン: $tool $specified_version"
    echo "  📋 現在バージョン: $tool $current_version"
    echo "  📦 インストール済み: $installed_flag"

    # バージョンが異なる場合のみ更新
    if [ "$current_version" != "$specified_version" ] || [ "$installed_flag" != "true" ]; then
        echo "  🔄 バージョンの差異または未インストールを検出しました。更新中..."

        # 古いバージョンがインストールされている場合はアンインストール
        if [ "$current_version" != "$specified_version" ]; then
            if [ "$current_version" != "not-installed" ] && [ "$current_version" != "No version is set" ]; then
                if is_version_installed "$tool" "$current_version"; then
                    echo "  🗑️  古いバージョンをアンインストール中: $current_version"
                    asdf uninstall $tool $current_version || echo "  ⚠️  警告: $tool $current_version のアンインストールに失敗しました（未インストールの可能性があります）"
                fi
            fi
        fi

        # 新しいバージョンをインストール
        echo "  📦 バージョンをインストール中: $specified_version"
        asdf install $tool $specified_version

        # グローバルバージョンを設定
        echo "  🔧 グローバルバージョンを設定中: $specified_version"
        asdf set -u $tool $specified_version

        # Node.jsの場合は追加パッケージをインストール
        if [ "$tool" == "nodejs" ]; then
            install_node_packages
        fi

        echo "  ✅ $tool を $specified_version に更新完了"
    else
        echo "  ✅ $tool は既に最新です"
    fi

    echo ""
done < .tool-versions

echo "=== 🎉 ツールバージョン更新完了 ==="
echo ""
echo "📊 現在のツールバージョン一覧:"
asdf current
echo ""
echo "🚀 すべてのツールが .tool-versions で指定されたバージョンに更新されました！"
