package storage

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ohsu-comp-bio/funnel/config"
)

const (
	protocol       = "htsget://"
	protocolBearer = protocol + "bearer:"
	privateKeyFile = ".private.key"
	publicKeyFile  = ".public.key"
)

// HTSGET provides read access to public URLs.
//
// Note that it relies on following programs to be installed and available in
// the system PATH:
//
//   - "htsget" (client implementation of the protocol)
//   - "crypt4gh" (to support "*.c4gh" encrypted resources)
//   - "crypt4gh-keygen" (to generate private and public keys)
//
// For more info about the programs:
//   - https://htsget.readthedocs.io/en/latest/
//   - https://crypt4gh.readthedocs.io/en/latest/
type HTSGET struct {
	conf config.HTSGETStorage
}

// NewHTSGET creates a new HTSGET instance.
func NewHTSGET(conf config.HTSGETStorage) (*HTSGET, error) {
	return &HTSGET{conf: conf}, nil
}

// Join a directory URL with a subpath.
func (b *HTSGET) Join(url, path string) (string, error) {
	return "", nil
}

// Stat returns information about the object at the given storage URL.
func (b *HTSGET) Stat(ctx context.Context, url string) (*Object, error) {
	return nil, nil
}

// List a directory. Calling List on a File is an error.
func (b *HTSGET) List(ctx context.Context, url string) ([]*Object, error) {
	return nil, nil
}

func (b *HTSGET) Put(ctx context.Context, url, path string) (*Object, error) {
	return nil, nil
}

// Get copies a file from a given URL to the host path.
//
// Supports fetching "*.c4gh" encrypted files. If the `path“ does not end with
// "*.c4gh", the content will be also decrypted. Otherwise, the private key for
// decrypting, will be located in `.private.key`.
func (b *HTSGET) Get(ctx context.Context, url, path string) (*Object, error) {
	httpsUrl := strings.Replace(url, protocol, "https://", 1)
	cmdArgs := make([]string, 0)
	tmpPath := path

	if strings.HasPrefix(url, protocolBearer) {
		bearerStart := len(protocolBearer)
		bearerStop := strings.Index(url, "@")

		if bearerStop < 1 {
			return nil, fmt.Errorf("Bearer token not terminated by @")
		}

		httpsUrl = "https://" + url[bearerStop+1:]
		cmdArgs = append(cmdArgs, "--bearer-token", url[bearerStart:bearerStop])
	}

	if strings.HasSuffix(url, ".c4gh") {
		cmdArgs = append(cmdArgs, "--headers", createHtsgetHeader())
		if !strings.HasSuffix(path, ".c4gh") {
			tmpPath = path + ".tmp"
		}
	}

	cmdArgs = append(cmdArgs, "--output", path)
	cmdArgs = append(cmdArgs, httpsUrl)

	err := runCmd("htsget", cmdArgs...)
	if err != nil {
		return nil, err
	}

	if tmpPath != path {
		err = decrypt(tmpPath, path)
		if err != nil {
			return nil, err
		}
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &Object{
		URL:          url,
		Size:         info.Size(),
		LastModified: info.ModTime(),
		Name:         path,
	}, nil
}

// UnsupportedOperations describes which operations (Get, Put, etc) are not
// supported for the given URL.
func (b *HTSGET) UnsupportedOperations(url string) UnsupportedOperations {
	if err := b.supportsPrefix(url); err != nil {
		return AllUnsupported(err)
	}

	ops := UnsupportedOperations{
		List: fmt.Errorf("htsgetStorage: List operation is not supported"),
		Put:  fmt.Errorf("htsgetStorage: Put operation is not supported"),
		Join: fmt.Errorf("htsgetStorage: Join operation is not supported"),
		Stat: fmt.Errorf("htsgetStorage: Stat operation is not supported"),
	}
	return ops
}

func (b *HTSGET) supportsPrefix(url string) error {
	if !strings.HasPrefix(url, protocol) {
		return &ErrUnsupportedProtocol{"htsgetStorage"}
	}
	return nil
}

func createHtsgetHeader() string {
	ensureKeyFiles()

	file, err := os.Open(publicKeyFile)
	if err != nil {
		fmt.Println("Could not read", publicKeyFile, "file, which should exist:", err)
		panic(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Text()              // Skip header.
	publicKey := scanner.Text() // The key is on the second line.
	file.Close()

	// HTTP headers to be encoded as JSON:
	headers := make(map[string]string)
	headers["client-public-key"] = publicKey

	headersJson, err := json.Marshal(&headers)
	if err != nil {
		fmt.Println("Failed to format JSON-header for passing client-public-key:", err)
		panic(1)
	}

	return string(headersJson)
}

func ensureKeyFiles() {
	files := []string{publicKeyFile, privateKeyFile}
	filesExist := true

	for i := range files {
		if file, err := os.Open(files[i]); err == nil {
			file.Close()
		} else {
			filesExist = false
			break
		}
	}

	if !filesExist {
		err := runCmd("crypt4gh-keygen", "-f", "--nocrypt",
			"--sk", privateKeyFile, "--pk", publicKeyFile)
		if err != nil {
			fmt.Println("Could not generate crypt4gh key-files:", err)
			panic(1)
		}
	}
}

func runCmd(commandName string, commandArgs ...string) error {
	cmd := exec.Command(commandName, commandArgs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("Error running command %s: %v\nSTDOUT: %s\nSTDERR: %s",
			commandName, err, stdout.String(), stderr.String())
	}
	return err
}

func decrypt(encryptedFile, decryptedFile string) error {
	encryptedReader, err := os.Open(encryptedFile)
	if err != nil {
		return err
	}
	defer encryptedReader.Close()
	defer os.Remove(encryptedFile)

	var stderr bytes.Buffer

	cmd := exec.Command("crypt4gh", "--sk", privateKeyFile)
	cmd.Stdin = encryptedReader
	cmd.Stderr = &stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("[FATAL] Failed to obtain STDOUT pipe for crypt4gh (decryption):", err)
		panic(1)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("[ERROR] Failed to execute crypt4gh (decryption):", err)
		return err
	}

	writer, err := os.Open(decryptedFile)
	if err != nil {
		fmt.Println("[ERROR] Failed to open file for decryption:", decryptedFile, err)
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, stdout)
	if err != nil {
		fmt.Println("[ERROR] Failed to copy decrypted content to", decryptedFile, ":", err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Error while decrypting the content (crypt4gh): %v\nSTDERR: %s",
			err, stderr.String())
	}

	return nil
}
