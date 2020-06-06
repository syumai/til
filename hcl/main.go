package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	rd, err := patch("./example.hcl", "./patch.hcl")
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, rd)
	return err
}

func parseFile(fileName string) (*hclwrite.File, error) {
	src, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	f, diags := hclwrite.ParseConfig(src, fileName, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, diags
	}
	return f, nil
}

func patch(baseFileName, patchFileName string) (io.Reader, error) {
	baseFile, err := parseFile(baseFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base file: %w", err)
	}
	baseBody := baseFile.Body()

	patchFile, err := parseFile(patchFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to parse patch file: %w", err)
	}
	patchBody := patchFile.Body()

	var addBody, removeBody *hclwrite.Body
	for _, b := range patchBody.Blocks() {
		if b.Type() == "add" {
			addBody = b.Body()
			continue
		}
		if b.Type() == "remove" {
			removeBody = b.Body()
		}
	}

	err = patchAdd(baseBody, addBody)
	if err != nil {
		return nil, err
	}

	err = patchRemove(baseBody, removeBody)
	if err != nil {
		return nil, err
	}

	//
	//baseBody.SetAttributeValue("hoge", cty.ListVal([]cty.Value{cty.StringVal("fuga")}))
	//baseBody.SetAttributeValue("hige", cty.ListVal([]cty.Value{cty.StringVal("fugi")}))
	//fmt.Println(baseBody.Attributes())
	//baseBody.AppendNewBlock("block", []string{"blockA"})
	//baseBody.AppendNewBlock("block", []string{"blockB"})
	//for _, b := range baseBody.Blocks() {
	//	if b.Type() == "block" && reflect.DeepEqual(b.Labels(), []string{"blockA"}) {
	//		b.Body().SetAttributeValue("hoho", cty.StringVal("hehe"))
	//	}
	//}

	return bytes.NewBuffer(hclwrite.Format(baseFile.Bytes())), nil
}

func patchAdd(baseBody, addBody *hclwrite.Body) error {
	for name, patchAttr := range addBody.Attributes() {
		baseAttr := baseBody.GetAttribute(name)
		if baseAttr == nil {
			fmt.Println(name)
			fmt.Printf("%#v\n", patchAttr)
			tokens := patchAttr.BuildTokens(nil)
			for _, t := range tokens {
				fmt.Printf("%#v\n", t)
			}
			fmt.Printf("%#v\n", patchAttr.Expr())
			continue
		}
		//baseTokens := baseAttr.BuildTokens(nil)
		//patchTokens := patchAttr.BuildTokens(nil)

	}
	return nil
}

func patchRemove(baseBody, removeBody *hclwrite.Body) error {
	return nil
}
