package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

//go:embed envoy.yaml
var originalEnvoyYaml string

func requireEnvoyYaml(t *testing.T, tmpdir string) (yamlPath string) {
	yamlPath = tmpdir + "/envoy.yaml"
	replacedYaml := strings.ReplaceAll(originalEnvoyYaml, "/tmp/", tmpdir)
	require.NoError(t, os.WriteFile(yamlPath, []byte(replacedYaml), 0644))
	fmt.Println("Envoy config:", replacedYaml)
	return
}

func TestIntegration(t *testing.T) {
	envoyImage := "envoy-dynamic"
	if os.Getenv("ENVOY_IMAGE") != "" {
		envoyImage = os.Getenv("ENVOY_IMAGE")
	}

	cwd, err := os.Getwd()
	require.NoError(t, err)

	tmpdir := t.TempDir()
	// Grant write permission to the tmpdir for the envoy process.
	require.NoError(t, exec.Command("chmod", "777", tmpdir).Run())
	yamlPath := requireEnvoyYaml(t, tmpdir)

	cmd := exec.CommandContext(
		context.TODO(),
		"docker",
		"run",
		"--name", "envoyinit-test",
		// "--network", "host",
		// "--rm",
		"-d",
		"-p", "1062:1062", "-p", "19000:19000", //mac doesnt play nicely with network host :(
		"-e", "RUST_BACKTRACE=1",
		"-v", cwd+":/integration",
		"-v", tmpdir+":"+tmpdir,
		"-w", tmpdir,
		envoyImage,
		"--concurrency", "1",
		"--config-path", yamlPath,
		"--base-id", strconv.Itoa(time.Now().Nanosecond()),
	)
	fmt.Println("Running command:", cmd.String())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = append(os.Environ(), "ENVOY_UID=0")
	require.NoError(t, cmd.Start())
	t.Cleanup(func() {
		require.NoError(t, cmd.Process.Signal(os.Interrupt))
	})

	// Let's wait at least 5 seconds for Envoy to start since it might take a while
	// to pull the image.
	for i := 0; i < 50; i++ {
		resp, err := http.Get("http://localhost:19000/ready")
		if err == nil {
			resp.Body.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	t.Run("http_header_mutations", func(t *testing.T) {
		require.Eventually(t, func() bool {
			req, err := http.NewRequest("GET", "http://localhost:1062/headers", nil)
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Logf("Envoy not ready yet: %v", err)
				return false
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Logf("Envoy not ready yet: %v", err)
				return false
			}

			t.Logf("response: headers=%v, body=%s", resp.Header, string(body))
			require.Equal(t, 200, resp.StatusCode)

			// HttpBin returns a JSON object containing the request headers.
			type httpBinHeadersBody struct {
				Headers map[string]string `json:"headers"`
			}
			var headersBody httpBinHeadersBody
			require.NoError(t, json.Unmarshal(body, &headersBody))

			require.Equal(t, "envoy-header", headersBody.Headers["X-Envoy-Header"])
			require.Equal(t, "envoy-header2", headersBody.Headers["X-Envoy-Header2"])

			// We also need to check that the response headers were mutated.
			require.Equal(t, "bar", resp.Header.Get("Foo"))
			require.Equal(t, "bar2", resp.Header.Get("Foo2"))
			return true
		}, 30*time.Second, 200*time.Millisecond)
	})
}
