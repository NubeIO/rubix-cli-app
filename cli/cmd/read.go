package cmd

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gthub.com/NubeIO/rubix-cli-app/modbus"
)

// Flags
var startRange int
var endRange int

// whoIsCmd represents the whoIs command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "BACnet device discovery",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: read,
}

func Title(ans interface{}) interface{} {

	mbClient := &modbus.Client{
		HostIP:   modbusIp,
		HostPort: modbusPort,
	}
	mbClient, err := modbus.SetClient(mbClient)
	if err != nil {
		fmt.Println("HERE")
		return errors.New("failed to connect to modbus")
	}
	mbClient.TCPClientHandler.Address = fmt.Sprintf("%s:%d", modbusIp, modbusPort)
	mbClient.TCPClientHandler.SlaveID = byte(1)
	coils, err := mbClient.Client.ReadCoils(uint16(startRange), uint16(endRange))
	if err != nil {
		fmt.Println("coils", err)
	}
	fmt.Println("coils", coils)
	return "res"
}

var qs = []*survey.Question{
	{
		Name:   "name",
		Prompt: &survey.Input{Message: "What is your name?"},
		//Validate:  survey.Required,
		Transform: Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
			Default: "red",
		},
	},
	{
		Name:   "age",
		Prompt: &survey.Input{Message: "How old are you?"},
	},
}

func read(cmd *cobra.Command, args []string) {

	answers := struct {
		Name          string // survey will match the question and field names
		FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
		Age           int    // if the types don't match, survey will convert it
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&startRange, "start", "s", 1, "Start range of discovery")
	readCmd.Flags().IntVarP(&endRange, "end", "e", 1, "End range of discovery")
}
