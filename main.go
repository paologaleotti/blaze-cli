package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/paologaleotti/blaze-cli/cli"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	color.BgYellow.Println(" blaze ")
	color.Yellow.Println(":: fast and simple web services in Go ::\n")

	fmt.Print("Enter project name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	fmt.Println("1. Cloning template...")
	err := cli.CloneRepository(projectName)
	if err != nil {
		log.Fatal("error cloning repository:", err)
		return
	}

	fmt.Println("2. Setting up project...")
	err = cli.RemoveIgnoredFiles(projectName)
	if err != nil {
		log.Fatal("error removing ignored files:", err)
		return
	}

	err = cli.ReplaceProjectName(projectName)
	if err != nil {
		log.Fatal("error replacing project name in files:", err)
		return
	}

	fmt.Println("3. Installing dependencies...")
	err = cli.InstallDependencies(projectName)
	if err != nil {
		log.Fatal("error installing dependencies:", err)
		return
	}

	color.Green.Println("\nYour new blaze project is ready!")

	color.Cyan.Println("To build and test the project:")
	color.Cyan.Println("> cd", projectName)
	color.Cyan.Println("> make")
}
