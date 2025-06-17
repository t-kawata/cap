# CAP (ShyMe Inc. Toshimi Kawata)
Claude Code が、自立型のコーディング TUI エージェントを目指すことで、ジュニアクラス ~ ミドルクラスに対する能力拡張を行っています。一方、シニアクラス以上にとって、Claude Code の親切さは、巨大なソースコード内において、じわじわとアンコントローラブル領域を増加させる遅効性毒となる可能性があり、2025年6月現在において Claude Code や Claude 4 シリーズよりも高いエンジニアリング能力を持つ人間のシニアエンジニアにとっては、副作用が大きくなりすぎる可能性があります。そこで、Claude Code よりも「人間が手綱を離さない」という性質を高めた、類似でありながらシニア向けのコーディング TUI エージェントを企画します。「帽子でもかぶるように、シニア専用モビルスーツを着る」という意味合いを込めて `CAP` とします。尚、第一に自分用の開発であり、日本語ユーザーに最適化します。また、美しいコードであることは考慮しません。自分が必要な時に必要なものを雑に拡張していきます。（OpenCodeを日本語及びシニア用にカスタマイズしていくものです。）

# 利用できるプロバイダー
- OpenAI
- Anthropic Claude
- Google Gemini
- AWS Bedrock
- Groq
- Azure OpenAI
- OpenRouter
- Ollama
- lamma.cpp

# Features
- TUI で動きます
- コマンドラインとしても使えます
- セッション管理内蔵です
- エージェントは、
    * 自分でコマンドを実行したり、
    * ファイルを探したり、
    * ファイルを編集したり、
    * デバッグしたり、
    * 編集したコードにエラーがないか LSP で確認したり、
    * 作業が長くなってきたら自分でそこまでを要約して作業を続けたり、
    * 作業指示書の指示に沿って、
    * 一つずつタスクを進めてチェックボックスにチェックしていったり、
    * 外付けのカスタムコマンドを通じて、任意のプログラムを実行したり、
    * APIにリクエストしたり、
    * MCPを使ったり、
    * ・・・などなど色々やってくれますが、
    * 基本的には、シニアエンジニアが「何を正しいとするか」を全て把握できていることが前提です。
- 詳細な指示書を書けば、連続的なタスクを自動的に行わせることができますが、
- 2025.06.17 現在、使用するモデルは `GPT-4.1-mini`, `Gemini-2.5-Flash` を想定して調整していますので、
- あまりに丸投げなことを指示しても期待する成果は出してくれません。
- シニアとして明確に定義・指示することで高度な関数やクラスは作成してくれます。
- つまり、何を、どのように、どのような順番で、どのようなアルゴリズムで
- 作っていけば、プロダクトクラスの開発物が完成するかが、
- ほぼ全て頭の中にあるという状態のシニアが、その実装を自動化することで
- 生産量を倍化する、というのがスコープです。
- もちろん、GPT-4.1-mini, Gemini-2.5-Flash よりも上位のモデルも CAP で使用することができます。
- Claude 4 モデルも使用可能です。
- が、速度と精度のバランスは、シニアにとって GPT-4.1-mini, Gemini-2.5-Flash が最も適していると思います。
- Gemini 2.5 Flash も安定的に動きますが、GPT-4.1-mini の方がクールな結果になるかもしれません。

# Kawata 用 Makefile （on Mac Book Pro）
- Apple Silicon Mac でしたら、
- 以下の make コマンドでインストールできるはずです。
- go v1.24.0 以上がインストールされていることが前提です。
## build & install
```
make build-install
```
## build (Optional)
```
make build-kawata
```
## install (Optional)
```
make install
```
## install 確認
```
$ cap -v
v1.0.2
```
## LSP
```
brew install ripgrep
brew install fzf
npm install -g vscode-html-languageserver-bin
npm install -g vscode-css-languageserver-bin
npm install -g vscode-json-languageserver
npm install -g typescript typescript-language-server
go install golang.org/x/tools/gopls@latest
```

# Apple Silicon Mac 以外の方のインストール
- ./main.go を go build を使用してビルドして下さい。
- できあがったバイナリひとつで動作します。
- PATH の通った場所に、適宜 cap という名前で移動、またはシンボリックリンクして下さい。
- LSP は、上記 `「LSP」` のものたちや自分に必要なものを環境に合わせてインストールして下さい。

# CAP の使い方
## 初期化
- CAP を使用するプロジェクトのルートディレクトリに `.cap.json` を作成することが初期化です。
- 以下のコマンドで雛形を作ります。
```
$ cd /path/to/project/root
$ cap init
```
## 設定
- `providers` は、使用するプロバイダーの設定だけにして、
- あとは削除して構いません。
- ただ、設定だけしておけば、
- CAP の中でモデルを切り替えたい時に `ctrl + o` で切り替えられるので便利です。
- `title`, `summarizer`, `translater` agent は、
- ollama や llama.cpp で起動した gemma-3-4b や gemma-3-8b あたりでも十分だったりします。
- OpenRouter の無料エンドポイントでこれらを使用することでも十分実用的に動作しますが、
- やはりちょっともっさりしているので、全てのエージェントタイプをメインモデルと同じにするか、
- 軽量なローカルモデルを使用する方が快適です。
- LSP は、使用する言語のソレだけを残せばOKです。
- LSP本体は、好みのものをインストールして下さい（VSCode系が成熟しているので良いです）。
```
{
    "providers": {
        "openai": {
            "apiKey": "<OPENAI_API_KEY>",
            "disabled": false
        },
        "gemini": {
            "apiKey": "<GEMINI_API_KEY>",
            "disabled": false
        },
        "anthropic": {
            "apiKey": "<ANTHROPIC_API_KEY>",
            "disabled": false
        },
        "groq": {
            "apiKey": "<GROQ_API_KEY>",
            "disabled": false
        },
        "openrouter": {
            "apiKey": "<OPENROUTER_API_KEY>",
            "disabled": false
        },
        "local": {
            "apiKey": "dummy",
            "disabled": false,
            "endpoint": "<e.g. http://localhost:11434/v1>"
        }
    },
    "agents": {
        "coder": {
            "model": "gpt-4.1-mini",
            "maxTokens": 30000
        },
        "task": {
            "model": "gpt-4.1-mini",
            "maxTokens": 5000
        },
        "title": {
            "model": "gpt-4.1-mini",
            "maxTokens": 80
        },
        "summarizer": {
            "model": "gpt-4.1-mini",
            "maxTokens": 2000
        },
        "translater": {
            "model": "gpt-4.1-mini",
            "maxTokens": 5000
        }
    },
    "lsp": {
        "go": {
            "disabled": false,
            "command": "gopls"
        },
        "typescript": {
            "disabled": false,
            "command": "typescript-language-server",
            "args": ["--stdio"]
        },
        "html": {
            "disabled": false,
            "command": "html-languageserver",
            "args": ["--stdio"]
        },
        "css": {    
            "disabled": false,    
            "command": "css-languageserver",    
            "args": ["--stdio"]    
        },
        "json": {
            "disabled": false,
            "command": "vscode-json-languageserver",
            "args": ["--stdio"]
        }
    }
}
```
- 使用可能なモデルのリストは `cap models` コマンドで確認できます。
- `.cap.json` の `agents` にて指定する `model` は、このいずれかです。
- 使用可能なモデルは、必要に応じて追加していきます。
```
$ cap models
azure.gpt-4.1
azure.gpt-4.1-mini
azure.gpt-4.1-nano
azure.gpt-4.5-preview
azure.gpt-4o
azure.gpt-4o-mini
azure.o1
azure.o1-mini
azure.o3
azure.o3-mini
azure.o4-mini
bedrock.claude-3.7-sonnet
claude-3-haiku
claude-3-opus
claude-3.5-haiku
claude-3.5-sonnet
claude-3.7-sonnet
claude-4-opus
claude-4-sonnet
deepseek-r1-distill-llama-70b
gemini-2.0-flash
gemini-2.0-flash-lite
gemini-2.5
gemini-2.5-flash
gpt-4.1
gpt-4.1-mini
gpt-4.1-nano
gpt-4.5-preview
gpt-4o
gpt-4o-mini
grok-3-beta
grok-3-fast-beta
grok-3-mini-beta
grok-3-mini-fast-beta
llama-3.3-70b-versatile
meta-llama/llama-4-maverick-17b-128e-instruct
meta-llama/llama-4-scout-17b-16e-instruct
o1
o1-mini
o1-pro
o3
o3-mini
o4-mini
openrouter.claude-3-haiku
openrouter.claude-3-opus
openrouter.claude-3.5-haiku
openrouter.claude-3.5-sonnet
openrouter.claude-3.7-sonnet
openrouter.deepseek-r1-free
openrouter.devstral-small-free
openrouter.gemini-2.5
openrouter.gemini-2.5-flash
openrouter.gemma-3-12b-it-free
openrouter.gemma-3-4b-it-free
openrouter.gpt-4.1
openrouter.gpt-4.1-mini
openrouter.gpt-4.1-nano
openrouter.gpt-4.5-preview
openrouter.gpt-4o
openrouter.gpt-4o-mini
openrouter.o1
openrouter.o1-mini
openrouter.o1-pro
openrouter.o3
openrouter.o3-mini
openrouter.o4-mini
qwen-qwq
vertexai.gemini-2.5
vertexai.gemini-2.5-flash
```
## 起動
- 起動は、プロジェクトルートディレクトリで `cap` を実行するだけです。
```
$ cd /path/to/project/root
$ cap
```
- 事前に `cap init` コマンドで `.cap.json` を作成して、
- 適切に設定しておくことを忘れないでください。

## 日本語で入力した内容を自動で英語翻訳してエージェントに渡す
- モデルの多くは、圧倒的に英語データで学習されています。
- よって、日本語で入力してもエージェントが成功させられなかった指示が、
- 英語で入力したら成功するということは珍しくありません。
- 特に、ローカルでSLMを使用する場合は尚更です。
- 入力するプロンプトの先頭に `/tl` + 半角スペースをつけてから送信すると、
- 自動的に英語翻訳した上でエージェントに送信してくれます。
- `tl` は `Trans-Late` の頭文字です。
- TUI の場合
```
/tl こんにちは。今日のご機嫌はいかが？
```
- と送信すると、エージェントが
```
Hello. How are you doing today?
```
- という英語で解釈されていることが画面上で見て取れます。
- 下記の「TUI ではなく、コマンド一発でエージェントを動かす」の際にも、
```
$ cap -p "/tl こんにちは。今日のご機嫌はいかがですか？"
Hello! I’m feeling ready to help you with your coding tasks today. How can I assist you with your project? 😊
```
- このように、入力は日本語でも自動的に英語で解釈されて実行されます。
- コマンド一発で実行した時は、どのように翻訳されたのかは見れません。

## TUI ではなく、コマンド一発でエージェントを動かす
- TUIでエージェントと会話しながら、複雑なミッションを進めていくことも便利ですが、
- README.mdをザッと書いて欲しいなど、
- 「サッとできるに決まってる系」は、
- コマンド一発でやってもらうとかが便利です。
- イメージとしては Makefile や シェル で細々とやっていた系のことを、
- そもそもそういうシェル自体も自分で書いて実行してくれる系エージェントに任せちゃう、
- みたいな使い方を Kawata はしています。
- あとは、TUI を起動するまでもないけど、ちょっとお願いしたい系を
```
cap -p "./src/notify/slack.go には、受信のEventをSubscribeする関数は定義されてる？"

cap -p "./src/drizzle/schema.ts に、id,name,email,passwordだけのusersの雛形書いて。他の定義を参考に。"
```
- とかは、コマンド一発でかなり正確にやってくれます。
- 以下、例です。
```
# シンプルなプロンプトで実行
cap -p "Explain the use of context in Go"

# 実行結果をJSONで受け取る（好みのプログラミング言語内で呼ぶとき便利）
cap -p "Explain the use of context in Go" -f json

# 実行ローダーが鬱陶しい時は -q オプションで消せます
cap -p "Explain the use of context in Go" -q
```

## キーボードショートカット
```
ctrl+?: ヘルプ表示
ctrl+l: ログ表示
ctrl+c: システム終了
ctrl+o: モデル切り替え
ctrl+t: テーマ切り替え
ctrl+e: エディタを開く
ctrl+s: メッセージ送信
esc: 閉じたり戻ったり
ctrl+n: 新しいセッション
ctrl+j: セッション移動
ctrl+k: コマンド一覧

ctrl+u: ⬆︎ ページを上へ
ctrl+d: ⬇︎ ページを下へ

@: ファイルパス探索
```

## カスタムコマンド
- `cap` コマンドで TUI CAP を起動すると、自動的に `.cap` ディレクトリが作成されます。
- `.cap` ディレクトリ内には `.cap/commands` ディレクトリがあり、
- この中に自由な Markdown ファイルをコマンドとして登録できます。
- 例えば、
```
cat <<EOF > ./.cap/commands/hello.md
こんにちは、ご機嫌はいかが？
EOF
```
- のように Markdown を作成すると、TUI 内で `ctrl + k` から hello コマンドを実行できます。
- Markdown 内の内容がメッセージとしてエージェントに送信される形です。
- また、例えば
```
cat <<EOF > ./.cap/commands/tasks.md
# 作業指示書
\$TASKS を最初に読み取り、\$TASKS に記載のタスク群を順番に連続的に実行していって下さい。一つのタスクが完了したら、チェックボックスにチェックを書き込んでから、現在の進捗状況を毎回報告して下さい。タスクは必ず一つずつ進めなければなりません。複数の項目を一度で進めないよう注意して下さい。
EOF
```
- のような連続作業指示のカスタムコマンドMarkdownを作成しておけば、
- `ctrl + k` から tasks コマンドを実行した直後、ダイアログで
- 「`$TASKS` のところには何を入れます？」みたいに聞いてきてくれるので、
```
- [] 1. index.htmlを作成して、bodyが空の雛形を書き込む
- [] 2. index.html内のbodyにh1タグで見出しを付け、その下に作ったpタグに狐の物語を書き込む
- [] 3. style.css を作って index.html内で読み込み、画面をおしゃれな見た目にする
```
- のような感じで作成しておいた ./test_tasks.md だと伝えると、
```
# 作業指示書
./test_tasks.md を最初に読み取り、./test_tasks.md に記載のタスク群を順番に連続的に実行していって下さい。一つのタスクが完了したら、チェックボックスにチェックを書き込んでから、現在の進捗状況を毎回報告して下さい。タスクは必ず一つずつ進めなければなりません。複数の項目を一度で進めないよう注意して下さい。
```
- という風に解釈して順次進めてくれます。
- さらに、例えば、以下のような ./test.sh を作ったとします。
```
#!/bin/bash
echo "テスト成功です！"
```
- そして、以下のようにカスタムコマンドMarkdownを作ったとします。
```
cat <<EOF > ./.cap/commands/test.md
./test.shを実行して出力結果を教えて下さい。
EOF
```
- すると、`ctrl + k` から test コマンドを実行した時、
- ./test.sh を実行し、実行結果を取得して、
```
./test.shを実行しました。実行した結果、「テスト成功です！」という出力結果を得ましたので、テストには成功したようです。
```
- とか返してくれます。
- つまり、外部で作成したどんな言語のプログラムも、エージェントの道具として渡すことができるということです。
- なお、`./cap/commands` 以下に階層的にディレクトリ分けしてカスタムコマンドMarkdownを配置しても適切に認識されます。
- 下記の通り、MCPも利用できます。

## MCP (Model Context Protocol)

CAP implements the Model Context Protocol (MCP) to extend its capabilities through external tools. MCP provides a standardized way for the AI assistant to interact with external services and tools.

### MCP Features

- **External Tool Integration**: Connect to external tools and services via a standardized protocol
- **Tool Discovery**: Automatically discover available tools from MCP servers
- **Multiple Connection Types**:
  - **Stdio**: Communicate with tools via standard input/output
  - **SSE**: Communicate with tools via Server-Sent Events
- **Security**: Permission system for controlling access to MCP tools

### Configuring MCP Servers

MCP servers are defined in the configuration file under the `mcpServers` section:

```json
{
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": [],
      "args": []
    },
    "web-example": {
      "type": "sse",
      "url": "https://example.com/mcp",
      "headers": {
        "Authorization": "Bearer token"
      }
    }
  }
}
```

### MCP Tool Usage

Once configured, MCP tools are automatically available to the AI assistant alongside built-in tools. They follow the same permission model as other tools, requiring user approval before execution.
