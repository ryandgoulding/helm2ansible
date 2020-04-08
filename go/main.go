package main

import (
	"fmt"
	"text/template"
	"text/template/parse"
)

func main() {
	t, err := template.ParseFiles("./preparsed_bitnami_helm_charts/deployment.yml")
	if err != nil { panic(err) }
	fmt.Println(ListTemplFields(t))
}

func ListTemplFields(t *template.Template) []string {
	return listNodeFields(t.Tree.Root, nil)
}

func listNodeFields(node parse.Node, res []string) []string {
	if node.Type() == parse.NodeIf {
		res = append(res, node.String() + "\n=============\n")
	}

	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			res = listNodeFields(n, res)
		}
	}
	return res
}
