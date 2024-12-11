module keybinding

go 1.20

require (
    github.com/atotto/clipboard v0.1.4 // indirect
    github.com/micmonay/keybd_event v1.1.2 // indirect
    github.com/go-vgo/robotgo v0.100.0 // Add robotgo dependency
)

replace github.com/go-vgo/robotgo => ./vendor/github.com/go-vgo/robotgo
