package schdulenotifier

import "github.com/gen2brain/beeep"

func Notify(title, message, appIcon string) error {
	return beeep.Notify(title, message, appIcon)
}
