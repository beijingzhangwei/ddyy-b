package main

import (
	"context"
	"errors"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"io"
)

func main() {
	// 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
	// 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
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 4000,
		Prompt:    "请用vue3.0实现提交用户输入到服务器，并输出结果，完整的代码实现",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}

func main2() {
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}
