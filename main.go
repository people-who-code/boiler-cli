package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: create-bp <framework> <projectname>")
		// create-bp [0],framework[1],projectname[2]
		os.Exit(1)
	}

	framework := os.Args[1] // store the first command. ie the framework

	projectname := os.Args[2]

	var cmd *exec.Cmd

	switch framework {
	case "express":
		cmd = exec.Command("bash", "-c", fmt.Sprintf("npx express-generator --view=ejs %s", projectname))
	case "nestjs":
		cmd = exec.Command("bash", "-c", fmt.Sprintf("npx @nestjs/cli new %s", projectname))
	case "laravel":
		cmd = exec.Command("bash", "-c", fmt.Sprintf("composer create-project --prefer-dist laravel/laravel %s", projectname))
	case "django":
		cmd = exec.Command("bash", "-c", fmt.Sprintf("django-admin startproject %s .", projectname))
	case "react":
		cmd = exec.Command("bash", "-c", fmt.Sprintf("npx create-react-app %s .", projectname))
	default:
		fmt.Printf("Unsupported framework: %s\n", framework)
		os.Exit(1)
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %s\n", err)
		os.Exit(1)
	}
}
