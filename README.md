# Enigma Go

エニグママシンをGoで実装するプロジェクト

## 概要

第二次世界大戦中にドイツ軍が使用した暗号機「エニグマ」の動作をGoで再現したプロジェクトです。

## 特徴

- エニグマの暗号化アルゴリズムを実装
- ローター、リフレクター、プラグボードの再現
- CLIによる操作
- JSON形式での設定管理

## 使用予定の技術

- **言語**: Go
- **設定**: JSON
- **CLI**: コマンドライン引数とフラグ
- **対象**: 標準的なエニグマ（3ローター構成）

## 使用例（予定）

```bash
# 暗号化
./enigma-go encrypt -input "HELLO WORLD" -config config.json

# 復号化  
./enigma-go decrypt -input "ILBDA AMTAZ" -config config.json
