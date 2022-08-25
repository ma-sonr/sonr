package commands

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/cmd/sonrd/internal/tui"
	"github.com/sonr-io/sonr/cmd/sonrd/internal/utils"
	"github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	"github.com/spf13/cobra"
)

var (
	logger = golog.Default.Child("sonrd")
)

func RootMotorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "motor",
		Short: "Setup a local Motor instance",
		Run: func(cmd *cobra.Command, args []string) {
			// Create a new TUI model which will be rendered in Bubbletea.
			state := tui.NewModel()
			// tea.NewProgram starts the Bubbletea framework which will render our
			// application using our state.
			if err := tea.NewProgram(state).Start(); err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.AddCommand(loginCmd(), registerCmd(), listCmd())
	return cmd
}

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to an existing sonr account on disk",
		Run: func(cmd *cobra.Command, args []string) {

			if ok := utils.PromptConfirm("Continue Login with System Keychain"); !ok {
				logger.Infof("Aborting login.")
				return
			}

			ual, err := utils.GetUserAuthList()
			if err != nil {
				logger.Errorf("Failed to fetch UserAuthList %e", err)
				return
			}
			addr, ua, err := utils.PromptAccSelect(ual, "Select an account for login")
			if err != nil {
				logger.Errorf("Failed to select account %e", err)
				return
			}
			req := mt.LoginRequest{
				Did:       addr,
				Password:  ua.Password,
				AesPskKey: ua.AesPSKKey,
				AesDscKey: ua.AesDSCKey,
			}
			m := setupMotor()
			_, err = m.Login(req)
			if err != nil {
				logger.Errorf("Failed to login with UserAuth %e", err)
				return
			}
			utils.DisplayMotorTable(m, "Logged In")
			if err := utils.Set([]byte("currentAccount"), []byte(addr)); err != nil {
				logger.Errorf("Failed to set currentAccount %e", err)
				return
			}
		},
	}
	return cmd
}

func registerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Create a new Sonr Account",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if user already has an account
			if exists := utils.HasUserAuth(); exists {
				if yes := utils.PromptConfirm("You already have an account. Do you want to register a new one"); !yes {
					logger.Infof("Aborting registration.")
					return
				}
			}
			logger.Infof("Registering new account...")
			passwd, err := utils.PromptNewPassword()
			if err != nil {
				logger.Errorf("Failed to create account %e", err)
			}
			ua, err := utils.NewUserAuth(passwd)
			if err != nil {
				logger.Errorf("Error creating new AES Key %e", err)
				return
			}
			req, err := ua.GenAccountCreateRequest()
			if err != nil {
				logger.Errorf("Error creating new AES Key %e", err)
				return
			}
			m := setupMotor()
			res, err := m.CreateAccount(*req)
			if err != nil {
				logger.Errorf("CreateAccount Error: %e", err)
				return
			}
			if yes := utils.PromptConfirm("Would you like to store AuthInfo in the system keychain"); yes {
				if err := ua.StoreAuth(res.Address, res.GetAesPsk()); err != nil {
					logger.Errorf("Failed to save UserAuth to Keychain %e", err)
					return
				}
			}
			utils.DisplayMotorTable(m, "Account Registered")
		},
	}
	return cmd
}

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all accounts on User Keychain",
		Run: func(cmd *cobra.Command, args []string) {
			if ok := utils.PromptConfirm("Logging in requires authorization over system Keychain."); !ok {
				logger.Infof("Aborting list.")
				return
			}
			ual, err := utils.GetUserAuthList()
			if err != nil {
				logger.Errorf("Failed to fetch UserAuthList %e", err)
				return
			}
			utils.DisplayAccListTable(ual)
		},
	}
	return cmd
}

func setupMotor() motor.MotorNode {
	initreq := &mt.InitializeRequest{
		DeviceId: utils.DesktopID(),
	}
	m, err := motor.EmptyMotor(initreq, common.DefaultCallback())
	if err != nil {
		logger.Fatalf("Error initializing motor node, %e", err)
	}
	return m
}
