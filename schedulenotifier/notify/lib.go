package notify

import "github.com/gen2brain/beeep"

func Do(title, message, appIcon string) error {
	return beeep.Notify(title, message, appIcon)
}
