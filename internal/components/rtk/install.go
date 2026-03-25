package rtk

import (
	"github.com/gentleman-programming/gentle-ai/internal/installcmd"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

// InstallCommand returns the platform-specific command sequence to install RTK.
func InstallCommand(profile system.PlatformProfile) (installcmd.CommandSequence, error) {
	return installcmd.NewResolver().ResolveComponentInstall(profile, "rtk")
}
