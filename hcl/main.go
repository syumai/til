package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	filename := "./example.hcl"
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	writer, diag := hclwrite.ParseConfig(f, filename, hcl.Pos{Line: 1, Column: 1})
	if diag != nil {
		return diag
	}
	body := writer.Body()

	attr := body.GetAttribute("fruits")
	fmt.Printf("%#v\n", attr)
	fmt.Printf("%#v\n", attr.Expr())
	for tok := range attr.BuildTokens(nil) {
		fmt.Printf("%#v\n", tok)
	}
	fmt.Printf("%#v\n", attr.Expr().Variables())
	body.SetAttributeValue("hige", cty.ListVal([]cty.Value{cty.StringVal("fugi")}))
	body.SetAttributeValue("hoge", cty.ListVal([]cty.Value{cty.StringVal("fuga")}))

	formattedSrc := hclwrite.Format(writer.Bytes())
	buf := bytes.NewBuffer(formattedSrc)
	if _, err := buf.WriteTo(os.Stdout); err != nil {
		return err
	}
	return nil
}
