package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	publiccode "github.com/italia/publiccode-parser-go"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [ OPTIONS ] publiccode.yml\n", os.Args[0])
		flag.PrintDefaults()
	}
	remoteBaseURLPtr := flag.String("remote-base-url", "", "The URL pointing to the directory where the publiccode.yml file is located.")
	localBasePathPtr := flag.String("path", "", "An absolute or relative path pointing to a locally cloned repository where the publiccode.yml is located.")
	disableNetworkPtr := flag.Bool("no-network", false, "Disables checks that require network connections (URL existence and oEmbed). This makes validation much faster.")
	exportPtr := flag.String("export", "", "Export the normalized publiccode.yml file to the given path.")
	helpPtr := flag.Bool("help", false, "Display command line usage.")
	flag.Parse()

	if *helpPtr || len(flag.Args()) < 1 {
		flag.Usage()
		return
	}

	p := publiccode.NewParser()
	p.RemoteBaseURL = *remoteBaseURLPtr
	p.LocalBasePath = *localBasePathPtr
	p.DisableNetwork = *disableNetworkPtr
	err := p.ParseFile(flag.Args()[0])
	if err != nil {
		fmt.Printf("validation ko:\n%v\n", err)
		return
	}
	fmt.Println("validation ok")

	if *exportPtr != "" {
		yaml, err := p.ToYAML()
		err = ioutil.WriteFile(*exportPtr, yaml, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("publiccode written to %s\n", *exportPtr)
	}
}
