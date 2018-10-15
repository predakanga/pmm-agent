package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/percona/pmm-agent/api"
	"github.com/percona/pmm-agent/app"
)

func TestList(t *testing.T) {
	t.Parallel()

	serveCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run app.
	{
		go func() {
			rootCmd := New(serveCtx, app.App{})
			rootCmd.SetArgs([]string{
				"--config", "testdata/.pmm-agent.yml",
				"serve",
			})
			buf := &bytes.Buffer{}
			rootCmd.SetOutput(buf)
			assert.NoError(t, rootCmd.Execute())
			assert.Equal(t, "", buf.String())
		}()
	}

	var buf *bytes.Buffer

	// Initial list should be empty.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"list",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// Add new program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"add", "mysql-1", "--env", "DATA_SOURCE_NAME=root@/", "--", "mysqld_exporter", "--collect.all",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// List now should contain new program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"list", "--json",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		resp := &api.ListResponse{}
		err := json.Unmarshal(buf.Bytes(), &resp)
		assert.NoError(t, err)
		expected := &api.ListResponse{
			Statuses: map[string]*api.Status{
				"mysql-1": {
					Program: &api.Program{
						Name: "mysqld_exporter",
						Arg: []string{
							"--collect.all",
						},
						Env: []string{
							"DATA_SOURCE_NAME=root@/",
						},
					},
					Running: true,
				},
			},
		}
		// PID is dynamic so we can't test it but we can ensure it's not empty.
		// OUT is dynamic so we can't test it but we can ensure it's not empty.
		for i := range resp.Statuses {
			assert.NotEmpty(t, resp.Statuses[i].Pid)
			resp.Statuses[i].Pid = 0
			assert.NotEmpty(t, resp.Statuses[i].Out)
			resp.Statuses[i].Out = ""
		}
		assert.Equal(t, expected, resp)
	}

	// Stop program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"stop", "mysql-1",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// List now should contain stopped program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"list", "--json",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		resp := &api.ListResponse{}
		err := json.Unmarshal(buf.Bytes(), &resp)
		assert.NoError(t, err)
		expected := &api.ListResponse{
			Statuses: map[string]*api.Status{
				"mysql-1": {
					Program: &api.Program{
						Name: "mysqld_exporter",
						Arg: []string{
							"--collect.all",
						},
						Env: []string{
							"DATA_SOURCE_NAME=root@/",
						},
					},
					Running: false,
					Err:     "signal: killed",
				},
			},
		}
		// PID is dynamic so we can't test it but we can ensure it's empty.
		// OUT is dynamic so we can't test it but we can ensure it's not empty.
		for i := range resp.Statuses {
			assert.Empty(t, resp.Statuses[i].Pid)
			resp.Statuses[i].Pid = 0
			assert.NotEmpty(t, resp.Statuses[i].Out)
			resp.Statuses[i].Out = ""
		}
		assert.Equal(t, expected, resp)
	}

	// Start program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"start", "mysql-1",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// List now should contain started program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"list", "--json",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		resp := &api.ListResponse{}
		err := json.Unmarshal(buf.Bytes(), &resp)
		assert.NoError(t, err)
		expected := &api.ListResponse{
			Statuses: map[string]*api.Status{
				"mysql-1": {
					Program: &api.Program{
						Name: "mysqld_exporter",
						Arg: []string{
							"--collect.all",
						},
						Env: []string{
							"DATA_SOURCE_NAME=root@/",
						},
					},
					Running: true,
				},
			},
		}
		// PID is dynamic so we can't test it but we can ensure it's empty.
		// OUT is dynamic so we can't test it but we can ensure it's not empty.
		for i := range resp.Statuses {
			assert.NotEmpty(t, resp.Statuses[i].Pid)
			resp.Statuses[i].Pid = 0
			assert.NotEmpty(t, resp.Statuses[i].Out)
			resp.Statuses[i].Out = ""
		}
		assert.Equal(t, expected, resp)
	}

	// Add another new program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"add", "mysql-2", "--env", "DATA_SOURCE_NAME=root@/", "--",
			"mysqld_exporter",
			"--collect.all",
			"--web.listen-address", ":9204",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// List now should contain started programs.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"list", "--json",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		resp := &api.ListResponse{}
		err := json.Unmarshal(buf.Bytes(), &resp)
		assert.NoError(t, err)
		expected := &api.ListResponse{
			Statuses: map[string]*api.Status{
				"mysql-1": {
					Program: &api.Program{
						Name: "mysqld_exporter",
						Arg: []string{
							"--collect.all",
						},
						Env: []string{
							"DATA_SOURCE_NAME=root@/",
						},
					},
					Running: true,
				},
				"mysql-2": {
					Program: &api.Program{
						Name: "mysqld_exporter",
						Arg: []string{
							"--collect.all",
							"--web.listen-address", ":9204",
						},
						Env: []string{
							"DATA_SOURCE_NAME=root@/",
						},
					},
					Running: true,
				},
			},
		}
		// PID is dynamic so we can't test it but we can ensure it's empty.
		// OUT is dynamic so we can't test it but we can ensure it's not empty.
		for i := range resp.Statuses {
			assert.NotEmpty(t, resp.Statuses[i].Pid)
			resp.Statuses[i].Pid = 0
			assert.NotEmpty(t, resp.Statuses[i].Out)
			resp.Statuses[i].Out = ""
		}
		assert.Equal(t, expected, resp)
	}

	// Stop programs.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"stop", "mysql-1", "mysql-2",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// Start programs.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"start", "mysql-1", "mysql-2",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// Stop all programs.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"stop",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// Start all programs.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"start",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// Remove program.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"remove", "mysql-1",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// Remove all programs.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"remove",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}

	// List should be empty again.
	{
		rootCmd := New(serveCtx, app.App{})
		rootCmd.SetArgs([]string{
			"--config", "testdata/.pmm-agent.yml",
			"list",
		})
		buf = &bytes.Buffer{}
		rootCmd.SetOutput(buf)
		assert.NoError(t, rootCmd.Execute())
		assert.Equal(t, "", buf.String())
	}
}
