/****************************************************************************************************
*
* Redirect module, special for Meta CPA, Ltd.
* by Michael S. Merzlyakov AFKA predator_pc@02122018
* version v1.0a
*
* created at 19122018
* last edit: 20122018
*
* usage: $ ./redirect
*
*****************************************************************************************************/

package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	General struct {
		Name string
		Host string
		Port int
	}
	Redirect struct {
		Code int
		From string
		To string
	}
	Debug struct {
		Level int
	}
}

var Cfg Config 

func InitConfig() {
	var actionArg string
	_ = actionArg

	err := gcfg.FatalOnly(gcfg.ReadFileInto(&Cfg, "settings.ini"))

	if len(os.Args) > 1 {
		actionArg = os.Args[1]
		if Cfg.Debug.Level > 0 {
			fmt.Println("Current command", "actionArg", "redirect.go")
		}
	}

	if err != nil {
		fmt.Println("Config error", err, "redirect.go")
		os.Exit(3) // exit anyway
	} else {

		switch actionArg {

		case "config":
			{
				if Cfg.General.Name != "" || Cfg.General.Port != 0 {
					fmt.Println("[ General ]")
					if Cfg.General.Name != "" {
						fmt.Println("[ -- Name ]", Cfg.General.Name)
					}
					if Cfg.General.Host != "" {
						fmt.Println("[ -- Host ]", Cfg.General.Host)
					}
					if Cfg.General.Port != 0 {
						fmt.Println("[ -- Port ]", Cfg.General.Port)
					} else {
						fmt.Println("[ -- Empty ]")
					}
				}

				if Cfg.Redirect.From != "" || Cfg.Redirect.To != "" || Cfg.Redirect.Code != 0 {
					fmt.Println("[ Redirect ]")
					if Cfg.Redirect.Code != 0 {
						fmt.Println("[ -- Code ]", Cfg.Redirect.Code)
					}
					if Cfg.Redirect.From != "" {
						fmt.Println("[ -- From ]", Cfg.Redirect.From)
					}
					if Cfg.Redirect.To != "" {
						fmt.Println("[ -- To ]", Cfg.Redirect.To)
					} else {
						fmt.Println("[ -- Empty ]")
					}
				}

				if Cfg.Debug.Level != 0 {
					fmt.Println("[ Debug ]")
					if Cfg.Debug.Level != 0 {
						fmt.Println("[ -- Debug level ]", Cfg.Debug.Level)
					} else {
						fmt.Println("[ -- Empty ]")
					}
				} else {
					fmt.Println("[ Empty configuration ]")
				}
				os.Exit(3) // exit anyway
			}
		case "run":
			break
		default:
			{
				fmt.Println("Usage: [this-file] command options")
				fmt.Println("Commands: --run - start API service")
				fmt.Println("          --config - show usable ini file settings")
				fmt.Println("          --help /none - show this message")
				os.Exit(3)
			}
		}
	}
}

func init() {
	InitConfig() // loading configuration globally
}

func handler(c echo.Context) error{
	return c.Redirect(Cfg.Redirect.Code,Cfg.Redirect.To)
}

func main() {
	// Echo instance
	router := echo.New()
	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	//avoid chrome to request favicon
	router.GET("/favicon.ico", func(c echo.Context) error {
		return c.String(404, "not found") //nothing
	})
	router.HideBanner = true
	// Routes
	router.GET(Cfg.Redirect.From, handler)
	router.GET(Cfg.Redirect.From+"/", handler)
	// run router
	if Cfg.General.Port != 0 {
		router.Logger.Fatal(router.Start(":" + strconv.Itoa(Cfg.General.Port)))
	} else {
		//exit if not
		panic("[ERROR] Failed to obtain server port from settings.ini")
	}
}
