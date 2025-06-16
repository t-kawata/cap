package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/cap-ai/cap/cmd"
	"github.com/cap-ai/cap/internal/llm/models"
	"github.com/cap-ai/cap/internal/logging"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		firstArg := args[0]
		switch firstArg {
		case "init":
			initialize(false)
			return
		case "conf":
			initialize(true)
			return
		case "models":
			showModels()
			return
		}
	}

	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}

const configSample = `{
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
}`

func initialize(print bool) {
	if print {
		fmt.Println(configSample)
		return
	}
	err := os.WriteFile("./.cap.json", []byte(configSample), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	lists := []string{
		"brew install ripgrep",
		"brew install fzf",
		"npm install -g vscode-html-languageserver-bin",
		"npm install -g vscode-css-languageserver-bin",
		"npm install -g vscode-json-languageserver",
		"npm install -g typescript typescript-language-server",
		"go install golang.org/x/tools/gopls@latest",
	}
	fmt.Print("\nYou may need the following dependencies.\n")
	fmt.Print("The followings assumes you're using an Apple Silicon Mac.\n")
	fmt.Print("If you are not using an Apple Silicon Mac, \nplease install the followings as appropriate for your environment.\n\n")
	for _, list := range lists {
		fmt.Printf("- %s\n", list)
	}
	fmt.Print("\n")
}

func showModels() {
	modelList := []string{}
	for modelID := range models.SupportedModels {
		modelList = append(modelList, string(modelID))
	}
	sort.Strings(modelList)
	for _, m := range modelList {
		fmt.Println(m)
	}
}
