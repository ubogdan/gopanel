package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/ubogdan/gopanel/model"
	"github.com/ubogdan/gopanel/repository/sqlite"
)

const (
	capSetGID  = 6
	capSetUID  = 7
	daemonFlag = "d"
	debugFlag  = "debug"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-panel-service"
	app.Version = model.Version().String()
	app.Action = startService

	daemonCliFlag := cli.BoolFlag{ //nolint:exhaustivestruct
		Name: daemonFlag,
	}

	debugCliFlag := cli.BoolFlag{ //nolint:exhaustivestruct
		Name: debugFlag,
	}

	app.Flags = []cli.Flag{
		&daemonCliFlag,
		&debugCliFlag,
	}

	app.Run(os.Args)
}

func startService(c *cli.Context) error {
	if os.Getuid() == 0 {
		self, err := os.Readlink("/proc/self/exe")
		if err != nil {
			return err
		}

		daemonize := c.Bool(daemonFlag)
		procName := filepath.Base(self)
		procArgs := []string{procName}

		if !daemonize && c.Bool(debugFlag) {
			procArgs = append(procArgs, "-"+debugFlag)
		}

		proc, err := os.StartProcess(
			procName,
			procArgs,
			&os.ProcAttr{
				Dir: filepath.Dir(self),
				Files: []*os.File{
					nil,
					os.Stdout,
					os.Stderr,
				},
				Sys: &syscall.SysProcAttr{
					Credential: &syscall.Credential{
						Uid: model.DaemonizeUID,
						Gid: model.DaemonizeGID,
					},
					Setsid: daemonize,
					AmbientCaps: []uintptr{
						capSetUID,
						capSetGID,
					},
				},
			},
		)
		if err != nil {
			return err
		}

		if !daemonize {
			_, err = proc.Wait()

			return err
		}

		return nil
	}

	db, err := sqlite.Database(model.DatabaseFilePath)
	if err != nil {
		return errors.Wrap(err, "open database file")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	r := mux.NewRouter()

	// Internal API
	api := r.PathPrefix("/v1").Subrouter()

	// corsOptions && Authorization goes here
	api.Methods(http.MethodOptions)

	// ----------------
	httpd := http.Server{
		Addr:           ":3380",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	var g run.Group
	ctx, cancel := context.WithCancel(context.Background())

	g.Add(
		func() error {
			return httpd.ListenAndServe()
		},
		func(error) {
			httpd.Shutdown(ctx)
			cancel()
		},
	)

	return g.Run()
}
