package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/lynxbat/gospel"
)

/*

gospel
	--truth key1,key2,key3
	--tags tag1,tag2,tag3
	--witness-conf <FILE>
	--ext-truths <DIR>

		gather
			--speak table|json|yaml|line
			--concurrent 1

		list


*/

func main() {
	app := cli.NewApp()
	app.Name = "gospel"
	app.Version = "0.0.1" // TODO replace with git version
	app.Usage = "A tool used to gather truths about a system"

	app.Run(os.Args)
	g := gospel.New()
	l := g.ListTruths()
	for k, x := range l {
		fmt.Printf("%s %s\n", k, x.Description())
	}
	gt := g.GatherTruths()
	for k, v := range gt {
		fmt.Printf("%s=%s\n", k, v)
	}
}
