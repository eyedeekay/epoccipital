package cli

import (
	"crypto/tls"
	"embed"
	"io/fs"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/eyedeekay/onramp"
	"github.com/eyedeekay/sam3"
	"github.com/juanfont/headscale"
)

//go:embed web/*
//go:embed web/_app
//go:embed web/_app/immutable/chunks/_layout-da46b06b.js
//go:embed web/_app/immutable/modules/pages/_layout.js-9cbb603b.js
//go:embed web/_app/immutable/assets/_layout-65c5e8f3.css
//go:embed web/_app/immutable/components/pages/groups.html/_page.svelte-2d3e607c.js
//go:embed web/_app/immutable/components/pages/settings.html/_page.svelte-f15c69ed.js
//go:embed web/_app/immutable/components/pages/users.html/_page.svelte-988a5fbf.js
//go:embed web/_app/immutable/components/pages/_page.svelte-42d9e92d.js
//go:embed web/_app/immutable/components/pages/_layout.svelte-df0707af.js
//go:embed web/_app/immutable/components/pages/devices.html/_page.svelte-552d0c89.js
var f embed.FS

func init() {
	rootCmd.AddCommand(hiddenServeCmd)
}

type mimeHandler struct {
	//FS http.FileSystem
}

func (m *mimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/web/index.html"
	}
	if strings.HasSuffix(path, "favicon.ico") {
		path = "/web/favicon.png"
	}
	if path == "/web/" {
		path = "/web/index.html"
	}
	path = strings.TrimPrefix(path, "/")
	f, err := f.Open(path)
	if err != nil {
		log.Error().Caller().Err(err).Msg("Server Error '" + path + "'")
		return
	}
	ext := filepath.Ext(path)
	mt := mime.TypeByExtension(ext)
	if ext == ".js" {
		log.Debug().Caller().Msg("'" + path + "''" + ext + "'")
		mt = "application/javascript"
	}
	w.Header().Add("content-type", mt)
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error().Caller().Err(err).Msg("Server Error '" + path + "'")
		return
	}
	w.Write(bytes)
}

func MimeHandler(fs http.FileSystem) http.Handler {
	var mimeTest http.Handler = &mimeHandler{}
	return mimeTest
}

func run() error {
	return fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		//fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		return nil
	})
}

func UIServer(cmd *cobra.Command, app *headscale.Headscale) {
	samaddr, err := cmd.Flags().GetString("samaddr")
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("Error initializing SAMv3 Address")
	}
	tries := 0
	mime.AddExtensionType(".js", "application/javascript")
	for tries < 10 {
		garlic, err := onramp.NewGarlic(hs+"-webui", samaddr, sam3.Options_Wide)
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error initializing SAMv3 API for UI server")
		}
		ln, err := garlic.ListenTLS()
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error initializing SAMv3 API for UI server")
		}
		app.SetAllowOrigin("https://" + ln.Addr().String())
		mux.Unlock("Access-Control-Allow-Origin Header is set")
		fileServer := MimeHandler(http.FS(f)) //http.FileServer(http.FS(f))
		http.Handle("/web/", fileServer)

		if err := http.Serve(ln, fileServer); err != nil {
			log.Warn().Caller().Err(err).Msg("Error running WebUI server, retrying in " + strconv.Itoa(tries) + " minutes")
			time.Sleep(time.Duration(tries) * time.Second)
		}
		ln.Close()
	}
}

type mutex struct {
	sync.Mutex
	locked bool
}

func (m *mutex) Lock(msg ...string) {
	for _, m := range msg {
		log.Printf("Lock Message: %s\n", m)
	}
	if !m.locked {
		return
	}
	m.locked = true
	m.Mutex.Lock()
}

func (m *mutex) Unlock(msg ...string) {
	for _, m := range msg {
		log.Printf("Unlock Message: %s\n", m)
	}
	if !m.locked {
		return
	}
	m.locked = false
	m.Mutex.Unlock()
}

var mux mutex

var hiddenServeCmd = &cobra.Command{
	Use:   "hiddenserve",
	Short: "Launches the " + hs + " server as a hidden(I2P) service",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		run()
		app, err := getHeadscaleApp()
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error initializing")
		}
		go UIServer(cmd, app)
		mux.Lock("locking while we wait for the UI server to start")
		samaddr, err := cmd.Flags().GetString("samaddr")
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("Error initializing SAMv3 Address")
		}
		garlic, err := onramp.NewGarlic(hs, samaddr, sam3.Options_Wide)
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
