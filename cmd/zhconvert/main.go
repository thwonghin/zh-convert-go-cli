package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/thwonghin/zh-convert-go-cli/internal/flagutils"
	"github.com/thwonghin/zh-convert-go-cli/internal/iohandler"
	"github.com/thwonghin/zh-convert-go-cli/internal/zhconvert"
)

func main() {
	req := &zhconvert.ConvertRequest{}
	stringFlags := make(map[string]*string)
	boolFlags := make(map[string]*bool)
	intFlags := make(map[string]*int)
	stringSliceFlags := make(map[string]*string)

	flagutils.BindFlagsFromStruct(req, stringFlags, boolFlags, intFlags, stringSliceFlags)
	// delete(stringFlags, "text")
	flag.Parse()
	flagutils.PopulateStructFromFlags(req, stringFlags, boolFlags, intFlags, stringSliceFlags)

	if err := req.Validate(); err != nil {
		log.Fatalf("invalid flag: %v", err)
	}

	ctx := context.Background()
	processor := func(text *string) (*string, error) {
		client := zhconvert.NewClient()
		req.Text = text

		result, err := client.Convert(ctx, *req)
		if err != nil {
			return nil, err
		}
		return &result.Text, nil
	}
	iohandler.ProcessStreamingBatches(os.Stdin, os.Stdout, processor, 100_000_000, "\n")
}
