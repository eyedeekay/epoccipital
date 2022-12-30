package cli

import (
	"crypto/tls"
	"net"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/eyedeekay/onramp"
	"github.com/eyedeekay/sam3"
	"github.com/juanfont/headscale"
)

func init() {
	rootCmd.AddCommand(hiddenServeCmd)
}

var hiddenServeCmd = &cobra.Command{
	Use:   "hiddenserve",
	Short: "Launches the headscale server as a hidden(I2P) service",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		app, err := getHeadscaleApp()
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error initializing")
		}
		garlic, err := onramp.NewGarlic("headscale", "127.0.0.1:7656", sam3.Options_Wide)
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error initializing SAMv3 API")
		}
		headscale.UnixSocketListenFunc = net.Listen
		headscale.TCPSocketListenFunc = func(network, address string) (net.Listener, error) {
			return garlic.Listen()
		}
		headscale.UDPSocketListenFunc = func(network, address string) (net.PacketConn, error) {
			return garlic.ListenPacket()
		}
		headscale.TLSSocketListenFunc = func(network, laddr string, config *tls.Config) (net.Listener, error) {
			return garlic.ListenTLS()
		}
		defer garlic.Close()
		err = app.Serve()
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error starting server")
		}
	},
}
