package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func main() {

	// Check for help
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println("Usage: ./ascii-table [space seperated list of things to search for]")
		fmt.Println("Example: ./ascii-table \\? p \\\" A 3F")
		os.Exit(0)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Binary", "Octal", "Hex", "Decimal", "Character"})

	outputText := ""
	for i := 0; i < 128; i++ {
		// Check for invisible control chars
		switch i {
		case 0:
			outputText = "Null"
			break
		case 1:
			outputText = "Start of heading"
			break
		case 2:
			outputText = "Start of text"
			break
		case 3:
			outputText = "End of text"
			break
		case 4:
			outputText = "End of transmission"
			break
		case 5:
			outputText = "Enquiry"
			break
		case 6:
			outputText = "Acknowledge"
			break
		case 7:
			outputText = "Bell"
			break
		case 8:
			outputText = "Backspace"
			break
		case 9:
			outputText = "Horizontal Tab"
			break
		case 10:
			outputText = "Line Feed"
			break
		case 11:
			outputText = "Vertical tab"
			break
		case 12:
			outputText = "Form feed"
			break
		case 13:
			outputText = "Carriage return"
			break
		case 14:
			outputText = "Shift out"
			break
		case 15:
			outputText = "Shift in"
			break
		case 16:
			outputText = "Data link escape"
			break
		case 17:
			outputText = "Device control 1"
			break
		case 18:
			outputText = "Device control 2"
			break
		case 19:
			outputText = "Device control 3"
			break
		case 20:
			outputText = "Device control 4"
			break
		case 21:
			outputText = "Negative acknowledge"
			break
		case 22:
			outputText = "Synchronous idle"
			break
		case 23:
			outputText = "End of trans, block"
			break
		case 24:
			outputText = "Cancel"
			break
		case 25:
			outputText = "End of medium"
			break
		case 26:
			outputText = "Substitute"
			break
		case 27:
			outputText = "Escape"
			break
		case 28:
			outputText = "File seperator"
			break
		case 29:
			outputText = "Group seperator"
			break
		case 30:
			outputText = "Record seperator"
			break
		case 31:
			outputText = "Unit seperator"
			break
		case 32:
			outputText = "Space"
			break
		default:
			outputText = ""
			break
		}

		// Check output text
		if outputText == "" {
			outputText = fmt.Sprintf("%c", rune(i))
		} else {
			outputText = fmt.Sprintf("[%s]", strings.ToUpper(outputText))
		}

		// If a any input chars are provided try and match them
		var tableRow = []string{
			fmt.Sprintf("%b", i),
			fmt.Sprintf("%#o", i),
			fmt.Sprintf("%X", i),
			fmt.Sprintf("%d", i),
			outputText,
		}

		// Get search terms if they exist in the table output and filter
		if len(os.Args) > 1 {

			// Function for setting green
			green := color.New(color.FgGreen).SprintfFunc()

			// Used to identify if there was any match
			foundRow := false // Used to identify if there was any match

			colorTableRow := make([]string, 0)
			for _, item := range tableRow {
				found := false
				for _, arg := range os.Args[1:] {

					// Case insensitive comparison
					if strings.ToLower(item) == strings.ToLower(arg) {
						found = true
						foundRow = true
						break
					}
				}

				// If an arg matched the current item highlight it green before
				// adding it. Otherwise add it back as is
				if found {
					colorTableRow = append(colorTableRow, green(item))
				} else {
					colorTableRow = append(colorTableRow, item)
				}
			}

			// We only care about matches, so only add the colorTableRow to the
			// table if this was set to true
			if foundRow {
				table.Append(colorTableRow)
			}

			// If no search terms provide just add row to table
		} else {
			table.Append(tableRow)
		}

	}

	table.Render()
}
