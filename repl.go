package main


import (
    "strings"
    "os"
    "fmt"
    "bufio"
    "sort"
    "net/http"
    "encoding/json"
    )



type config struct{
    commands    map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}




	
func startRepl(cfg *config) {    

	scanner := bufio.NewScanner(os.Stdin)
	for {
	    fmt.Print("Pokedex >")
	    scanner.Scan()
	    input := scanner.Text()
	    cleaned := cleanInput(input)
        if len(cleaned) == 0 {
			continue
		}

	    commandName := cleaned[0]
        command, ok := cfg.commands[commandName]
        if !ok {
            fmt.Println("Unknown command")
            continue
        } else {
            command.callback(cfg)
        }
        
        
	}
}




func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(cfg *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println("")

    keys := make([]string, 0, len(cfg.commands))

    for k := range cfg.commands {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    for _, k := range keys {
    v := cfg.commands[k]
    fmt.Printf("%s: %s\n", v.name, v.description)
}
    
    return nil
}

func commandMap(cfg *config) error {
    fullURL := "https://pokeapi.co/api/v2/location-area/"  

    res, err := http.Get(fullURL) 
    if err != nil {
        fmt.Println("Error creating request:", err)
        return err

}
    type nameResult struct {
        Name string `json:"name"`
}
    
    type locationAreaJson struct {
        Next *string `json:"next"`
        Previous *string `json:"previous"`
        Results []nameResult `json:"results"`               
       

    }

    

    

}

