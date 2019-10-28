package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func main() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Binary", "Octal", "Hex", "Decimal", "Character"})

	outputText := ""
	for i := 0; i < 128; i++ {
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
			outputText = "File separator"
			break
		case 29:
			outputText = "Group separator"
			break
		case 30:
			outputText = "Record separator"
			break
		case 31:
			outputText = "Unit separator"
			break
		case 32:
			outputText = "Space"
			break
		default:
			outputText = ""
			break
		}
		if outputText == "" {
			table.Append(output(i, fmt.Sprintf("%c", rune(i))))
		} else {
			table.Append(output(i, fmt.Sprintf("[%s]", strings.ToUpper(outputText))))
		}
	}

	table.Render()
}

func output(input int, stringText string) []string {
	return []string{
		fmt.Sprintf("%#b", input),
		fmt.Sprintf("%#o", input),
		fmt.Sprintf("%X", input),
		fmt.Sprintf("%d", input),
		fmt.Sprintf("%s", stringText),
	}
}
