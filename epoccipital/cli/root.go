package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/juanfont/headscale"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-latest"
)

var cfgFile string = ""

func init() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "mockoidc" || os.Args[1] == "completion") {
		return
	}

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVarP(&cfgFile, "config", "c", "", "config file (default is /etc/headscale/config.yaml)")
	rootCmd.PersistentFlags().
		StringP("output", "o", "", "Output format. Empty for human-readable, 'json', 'json-line' or 'yaml'")
	rootCmd.PersistentFlags().
		StringP("samaddr", "s", "127.0.0.1:7656", "Address of the SAMv3 API for I2P")
	rootCmd.PersistentFlags().
		Bool("force", false, "Disable prompts and forces the execution")
	rootCmd.PersistentFlags().StringP("tunname", "n", hs, "Name to use for the I2P tunnel")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = os.Getenv("HEADSCALE_CONFIG")
	}
	if cfgFile != "" {
		err := headscale.LoadConfig(cfgFile, true)
		if err != nil {
			log.Fatal().Caller().Err(err).Msgf("Error loading config file %s", cfgFile)
		}
	} else {
		err := headscale.LoadConfig("", false)
		if err != nil {
			log.Fatal().Caller().Err(err).Msgf("Error loading config")
		}
	}

	cfg, err := headscale.GetHeadscaleConfig()
	if err != nil {
		log.Fatal().Caller().Err(err)
	}

	machineOutput := HasMachineOutputFlag()

	zerolog.SetGlobalLevel(cfg.Log.Level)

	// If the user has requested a "machine" readable format,
	// then disable login so the output remains valid.
	if machineOutput {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	if cfg.Log.Format == headscale.JSONLogFormat {
		log.Logger = log.Output(os.Stdout)
	}

	if !cfg.DisableUpdateCheck && !machineOutput {
		if (runtime.GOOS == "linux" || runtime.GOOS == "darwin") &&
			Version != "dev" {
			githubTag := &latest.GithubTag{
				Owner:      strings.Split(RepositoryURL, ",")[len(strings.Split(RepositoryURL, "/"))-2],
				Repository: strings.Split(RepositoryURL, ",")[len(strings.Split(RepositoryURL, "/"))-1],
			}
			res, err := latest.Check(githubTag, Version)
			if err == nil && res.Outdated {
				//nolint
				fmt.Printf(
					"An updated version of Headscale has been found (%s vs. your current %s). Check it out https://github.com/"+githubTag.Owner+"/"+githubTag.Repository+"/releases\n",
					res.Current,
					Version,
				)
			}
		}
	}
}

func Executable() (string, error) {
	hs, exeerr := os.Executable()
	return filepath.Base(hs), exeerr
}

var hs, exeerr = Executable()

var HelpMessage string = "is an I2P-Hosted Tailscale control server based on the Open-Source Headscale server"
var RepositoryURL string = "https://github.com/eyedeekay/epoccipital"
var rootCmd = &cobra.Command{
	Use:   hs,
	Short: hs + " - a Tailscale control server",
	Long: `
` + hs + ` is ` + HelpMessage + `
` + strings.Split(RepositoryURL, ",")[len(strings.Split(RepositoryURL, "/"))-2] + `
` + strings.Split(RepositoryURL, ",")[len(strings.Split(RepositoryURL, "/"))-1] + `
` + RepositoryURL,
}

func Execute() {
	if exeerr != nil {
		fmt.Fprintln(os.Stderr, exeerr)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
