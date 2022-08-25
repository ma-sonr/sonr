package state

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/jedib0t/go-pretty/table"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/pkg/motor"
)

func DisplayMotorTable(m motor.MotorNode) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Account Information")
	t.AppendHeader(table.Row{"Address", "DID", "Balance"})
	t.AppendRows([]table.Row{
		{m.GetAddress(), m.GetDID().String(), m.GetBalance()},
	})
	t.SetStyle(table.StyleLight)
	t.Render()
}

func DisplayAccListTable(ul UserAuthList) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Available Accounts")
	t.AppendHeader(table.Row{"#", "Address", "Created"})
	var rows []table.Row
	idx := 1
	for addr, ua := range ul.Auths {
		// Parse time RFC3339
		createdAt, _ := time.Parse(time.RFC3339, ua.CreatedAt)

		// Convert created at to D/M/Y HH:MM format
		createdAtStr := createdAt.Format("02/01/2006 15:04")
		rows = append(rows, table.Row{idx, addr, createdAtStr})
		idx++
	}
	t.AppendRows(rows)
	t.SetStyle(table.StyleLight)
	t.Render()
}

func PromptAccSelect(ul UserAuthList, label string) (string, UserAuth, error) {
	// Build options for promptui
	var items []*selection.Choice
	index := 0
	for addr, ua := range ul.Auths {
		items = append(items, &selection.Choice{
			String: addr,
			Value:  ua,
			Index:  index,
		})
		index++
	}
	sp := selection.New(label, items)
	sp.PageSize = 5
	sp.KeyMap.Select = append(sp.KeyMap.Select, "Continue")
	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	ua := choice.Value.(UserAuth)
	addr := choice.String

	// Prompt for password
	if ok, _ := PromptExistingPassword(ua.Password); !ok {
		return "", UserAuth{}, errors.New("Incorrect password entered")
	}
	return addr, ua, nil
}

func PromptConfirm(msg string) bool {
	input := confirmation.New(msg, confirmation.Yes)
	input.Template = confirmation.TemplateYN
	input.ResultTemplate = confirmation.ResultTemplateYN
	input.KeyMap.SelectYes = append(input.KeyMap.SelectYes, "✅")
	input.KeyMap.SelectNo = append(input.KeyMap.SelectNo, "❌")
	ready, err := input.RunPrompt()
	if err != nil {
		golog.Default.Error(err)
		return false
	}
	return ready
}

func PromptNewPassword() (string, error) {
	minCharacters := 8
	input := textinput.New("Choose a passphrase:")
	input.Placeholder = "make it strong!"
	input.Validate = func(s string) error {
		if len(s) < minCharacters {
			return fmt.Errorf("at least %d more characters", minCharacters-len(s))
		}
		return nil
	}
	input.Hidden = true
	input.Template += `
	{{- if .ValidationError -}}
		{{- print " " (Foreground "1" .ValidationError.Error) -}}
	{{- end -}}`

	pwd, err := input.RunPrompt()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func PromptExistingPassword(knownPwd string) (bool, error) {
	input := textinput.New("Confirm with your Password:")
	input.Placeholder = "password"
	input.Validate = func(s string) error {
		if knownPwd != s {
			return errors.New("Incorrect password entered")
		}
		return nil
	}
	input.Hidden = true
	input.Template += `
	{{- if .ValidationError -}}
		{{- print " " (Foreground "1" .ValidationError.Error) -}}
	{{- end -}}`

	_, err := input.RunPrompt()
	if err != nil {
		return false, err
	}
	return true, nil
}
