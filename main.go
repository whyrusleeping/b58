package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jbenet/go-base58"
	"github.com/spf13/cobra"
)

var b58Cmd = &cobra.Command{
	Use:   "b58",
	Short: "",
	Long:  "",
	Run:   b58CmdFunc,
}

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "",
	Long:  "",
	Run:   encodeFunc,
}

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "",
	Long:  "",
	Run:   decodeFunc,
}

func b58CmdFunc(c *cobra.Command, inp []string) {
	c.Help()
}

func encodeFunc(c *cobra.Command, inp []string) {
	i, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Print(base58.Encode(i))
}

func decodeFunc(c *cobra.Command, inp []string) {
	if len(inp) > 0 {
		for _, s := range inp {
			fmt.Println(string(base58.Decode(s)))
		}
	} else {
		i, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if i[len(i)-1] == '\n' {
			i = i[:len(i)-1]
		}
		decoded := base58.Decode(string(i))
		os.Stdout.Write(decoded)
	}
}

func main() {
	b58Cmd.AddCommand(encCmd, decCmd)
	b58Cmd.Execute()
}
