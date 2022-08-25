package main

import (
	"github.com/sonr-io/sonr/cmd/sonrd/internal/state"
	"github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	"github.com/spf13/cobra"
)

func RootMotorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "motor",
		Short: "Setup a local Motor instance",
	}
	cmd.AddCommand(loginCmd(), registerCmd(), listCmd())
	return cmd
}

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to an existing sonr account on disk",
		Run: func(cmd *cobra.Command, args []string) {

			if ok := state.PromptConfirm("Continue Login with System Keychain"); !ok {
				logger.Infof("Aborting login.")
				return
			}

			ual, err := state.GetUserAuthList()
			if err != nil {
				logger.Errorf("Failed to fetch UserAuthList %e", err)
				return
			}
			addr, ua, err := state.PromptAccSelect(ual, "Select an account for login")
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
			cb := state.GetCallback()
			cb.StartSpinner()
			m := setupMotor(cb)
			_, err = m.Login(req)
			if err != nil {
				logger.Errorf("Failed to login with UserAuth %e", err)
				return
			}
			cb.StopSpinner("User Authorized")
			state.DisplayMotorTable(m)
			if err := state.Set([]byte("currentAccount"), []byte(addr)); err != nil {
				logger.Errorf("Failed to set currentAccount %e", err)
				return
			}
		},
	}
	return cmd
}

func registerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register",
		Short:   "Create a new Sonr Account",
		Aliases: []string{"new", "create"},
		Run: func(cmd *cobra.Command, args []string) {
			// Check if user already has an account
			passwd, err := state.PromptNewPassword()
			if err != nil {
				logger.Errorf("Failed to create account %e", err)
			}
			ua, err := state.NewUserAuth(passwd)
			if err != nil {
				logger.Errorf("Error creating new AES Key %e", err)
				return
			}
			req, err := ua.GenAccountCreateRequest()
			if err != nil {
				logger.Errorf("Error creating new AES Key %e", err)
				return
			}
			cb := state.GetCallback()
			cb.StartSpinner()
			m := setupMotor(cb)
			res, err := m.CreateAccount(*req)
			if err != nil {
				logger.Errorf("CreateAccount Error: %e", err)
				return
			}

			cb.StopSpinner("Account Registered")
			if yes := state.PromptConfirm("Sonr requires system keychain access to store your Account Info. Continue?"); yes {
				ual, err := ua.StoreAuth(res.Address, res.GetAesPsk())
				if err != nil {
					logger.Errorf("Failed to save UserAuth to Keychain %e", err)
					return
				}
				state.DisplayAccListTable(ual)
			}
		},
	}
	return cmd
}

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all accounts on User Keychain",
		Run: func(cmd *cobra.Command, args []string) {
			if ok := state.PromptConfirm("Logging in requires authorization over system Keychain. Continue?"); !ok {
				logger.Infof("Aborting list.")
				return
			}
			ual, err := state.GetUserAuthList()
			if err != nil {
				logger.Errorf("Failed to fetch UserAuthList %e", err)
				return
			}
			state.DisplayAccListTable(ual)
		},
	}
	return cmd
}

func setupMotor(cb common.MotorCallback) motor.MotorNode {
	initreq := &mt.InitializeRequest{
		DeviceId: state.DesktopID(),
	}
	m, err := motor.EmptyMotor(initreq, cb)
	if err != nil {
		logger.Fatalf("Error initializing motor node, %e", err)
	}
	return m
}
