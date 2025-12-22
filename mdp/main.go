package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Document</title>
	<style>
		body { max-width: 700px; margin: 40px auto; font-family: sans-serif; line-height: 1.6; }
		pre { background: #f4f4f4; padding: 10px; border-radius: 4px; }
		code { background: #f4f4f4; padding: 2px 4px; border-radius: 4px; }
	</style>
</head>
<body>
`
	footer = `
</body>
</html>`
)

func main() {
	skipPreview := flag.Bool("s", false, "Skip the preview option")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Provide filename as argument")
		os.Exit(1)
	}

	filename := flag.Args()[0]


	if filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(filename string, writer io.Writer, skipPreview bool) error {
	content, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	htmlData := parseContent(content)

	outfile, err := os.CreateTemp("", "mdp*.html")

	if err != nil {
		return err
	}

	if err := outfile.Close(); err != nil {
		return err
	}

	fmt.Fprintln(writer, outfile.Name())

	if err := saveHtml(outfile.Name(), htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outfile.Name())

	return preview(outfile.Name())
}

func saveHtml(outName string, htmlData []byte) error {
	return os.WriteFile(outName, htmlData, 0644)
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)

	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer

	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func preview(filename string) error {
	cName := ""
	cParams := []string{}

	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "darwin":
		cName = "open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/", "start"}
	}

	cParams = append(cParams, filename)

	cPath, err := exec.LookPath(cName)

	if err != nil {
		return err
	}

	err = exec.Command(cPath, cParams...).Run()

	time.Sleep(2 * time.Second)

	return err
}
