package monitor

import (
        //"log"
        //"net/smtp"
)

type mailConfig struct {
        server string
        user string
        pass string
        address string
}


// EnableEmail enables email notifications
func EnableEmail() {
        notifications["email"] = true
}

// DisableEmail disables email notifications
func DisableEmail() {
        notifications["email"] = false
}

func LoadMailConfig(filename string) error {
        return nil
}