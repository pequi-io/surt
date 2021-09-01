package app

import (
	"github.com/surt-io/surt/pkg/apis"
)

func RunApp() {
	r := apis.SetupRouter()
	r.Run(":8080")
}
