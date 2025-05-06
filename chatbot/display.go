package main

import (
	"fmt"

	"github.com/fatih/color"
)

const welcomeArt = `
              )\       _
     .--._  ,'  \_.-~~/'
     \    \'_    __ ( _ _
       \   (_)  |__|/~ ~~=~\
         )_____.---~~ \>\~-./'
        /'    //===  /==(
       (     /' __\ ( __\)
       ( /~\(   o   |_o_(
       (( (         _____)
        \\_/     ,      )
         \ \   ~-.._  /
        ._)/ \        /
         \/   \     ./
          /     \~/~'
        /       /
       '~~-.__./
`

const goodbyeArt = `
  _____                 _ _                _ 
 / ____|               | | |              | |
| |  __  ___   ___   __| | |__  _   _  ___| |
| | |_ |/ _ \ / _ \ / _' | '_ \| | | |/ _ \ |
| |__| | (_) | (_) | (_| | |_) | |_| |  __/_|
 \_____|\___/ \___/ \__,_|_.__/ \__, |\___(_)
                                 __/ |       
                                |___/        
`

// printWelcomeArt prints the welcome ASCII art with colors
func printWelcomeArt() {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(cyan(welcomeArt))
	fmt.Println(yellow("Welcome to Chatty Charles!"))
	fmt.Println(green("Your friendly AI companion"))
	fmt.Println("--------------------------------------------")
}

// printGoodbyeArt prints the goodbye ASCII art with colors
func printGoodbyeArt() {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(cyan(goodbyeArt))
	fmt.Println(yellow("Thanks for chatting! Come back soon!"))
} 