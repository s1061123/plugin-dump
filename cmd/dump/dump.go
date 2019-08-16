package main

import (
	"fmt"
	"io"
	"os"
	"encoding/json"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"	
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

type NetConf struct {
	types.NetConf
	File string `json:"file,omitempty"`
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, bv.BuildString("dump"))
}

func dumpCmdArgs(fp io.Writer, args *skel.CmdArgs) {
	fmt.Fprintf(fp, `ContainerID: %s
Netns: %s
IfName: %s
Args: %s
Path: %s
StdinData: %s
----------------------
`,
		args.ContainerID,
		args.Netns,
		args.IfName,
		args.Args,
		args.Path,
		string(args.StdinData))
}

func parseConf(data []byte) (*NetConf, error) {
	conf := &NetConf{}
	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse")
	}

	if conf.File == "" {
		conf.File = "/tmp/cni_dump"
	}
	return conf, nil
}

func getResult(netConf *NetConf) (*current.Result) {

	if netConf.RawPrevResult == nil {
		return &current.Result{}
	}

	version.ParsePrevResult(&netConf.NetConf)
	result, _ := current.NewResultFromResult(netConf.PrevResult)
	return result
}

func cmdAdd(args *skel.CmdArgs) error {
	netConf, _ := parseConf(args.StdinData)
	fp, _ := os.OpenFile(netConf.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fp.Close()

	fmt.Fprintf(fp, "CmdAdd\n")
	dumpCmdArgs(fp, args)
	return types.PrintResult(getResult(netConf), netConf.CNIVersion)
}

func cmdDel(args *skel.CmdArgs) error {
	netConf, _ := parseConf(args.StdinData)
	fp, _ := os.OpenFile(netConf.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fp.Close()
	fmt.Fprintf(fp, "CmdDel\n")
	dumpCmdArgs(fp, args)
	return types.PrintResult(&current.Result{}, netConf.CNIVersion)
}

func cmdCheck(args *skel.CmdArgs) error {
	netConf, _ := parseConf(args.StdinData)
	fp, _ := os.OpenFile(netConf.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fp.Close()
	fmt.Fprintf(fp, "CmdCheck\n")
	dumpCmdArgs(fp, args)
	return types.PrintResult(&current.Result{}, netConf.CNIVersion)
}
