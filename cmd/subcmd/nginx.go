package subcmd

import (
	"fmt"
	gf "github.com/fenger/gomoon/file"
	su "github.com/fenger/gomoon/stringutil"
	"github.com/spf13/cobra"
	"path"
	"runtime"
)

var nginxCmd = &cobra.Command{
	Use:   "nginx",
	Short: "nginx",
	Run: func(cmd *cobra.Command, args []string) {
		runInstall(cmd)
	},
}

func init() {

	nginxCmd.PersistentFlags().StringP("version", "v", "1.18.0", "nginx version")

}

func GetNginxCmd() *cobra.Command {
	return nginxCmd
}

var nginxUnixUrl = "http://nginx.org/download/nginx-{{.Version}}.tar.gz"
var nginxWindowsUrl = "http://nginx.org/download/nginx-{{.Version}}.zip"

func runInstall(cmd *cobra.Command) {
	v := cmd.PersistentFlags().Lookup("version").Value
	fmt.Println("version: ", v)

	osType := runtime.GOOS
	if osType == "linux" {
		fmt.Println("linux")
	} else if osType == "windows" {
		fmt.Println("windows")
	} else if osType == "darwin" {
		fmt.Println("mac")
	}

	vars := make(map[string]interface{})
	vars["Version"] = v

	url, err := su.ParseString(nginxUnixUrl, vars)
	if err != nil {
		fmt.Errorf("url: %s parsed error", nginxUnixUrl)
		fmt.Println(err)
		return
	}
	fmt.Println(url)

	p := path.Base(".")
	fmt.Println(p)
	_, fn := path.Split(url)
	fmt.Println("Downloading", url)
	err = gf.DownloadFile("./"+fn, url)
	if err != nil {
		fmt.Println("file download failed", err)
	}
	fmt.Println("Downloaded", url)
}
