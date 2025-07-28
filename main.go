package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"

	"github.com/reiver/anthropic-test-20250728/cfg"
	"github.com/reiver/anthropic-test-20250728/srv/anthropic"
)

func main() {

	fmt.Println("Hello world!")

	const prompt string = "What is a quaternion?"

	{
		fmt.Printf("PROMPT:\n%q\n\n", prompt)
		fmt.Printf("MAX-TOKENS:\n%d\n\n", cfg.MaxTokens)
	}

	messageParams := anthropic.MessageNewParams{
		MaxTokens: cfg.MaxTokens,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
		Model: cfg.Model,
	}

	var message *anthropic.Message
	{
		var err error

		message, err = anthropicsrv.Client.Messages.New(context.TODO(), messageParams)
		if nil != err {
			fmt.Printf("ERROR: problem making request to Anthropic API: %s\n", err)
			fmt.Printf("prompt was:\n%q", prompt)
			os.Exit(1)
			return
		}
		if nil == message {
			fmt.Print("ERROR: problem making request to Anthropic API — nil message\n")
			fmt.Printf("prompt was:\n%q", prompt)
			os.Exit(1)
			return
		}
	}

	{
		fmt.Println("=========================")

		fmt.Printf("MESSAGE-ID: %q\n", message.ID)
		fmt.Printf("MESSAGE-MODEL: %q\n", message.Model)
		fmt.Printf("MESSAGE-ROLE: %q\n", message.Role)
		fmt.Printf("MESSAGE-STOP-REASON: %q\n", message.StopReason)
		fmt.Printf("MESSAGE-STOP-SEQUENCE: %q\n", message.StopSequence)
		fmt.Printf("MESSAGE-TYPE: %q\n", message.Type)


		fmt.Println("-------------------------")

		fmt.Printf("%+v\n", message.Content)

	}
}
