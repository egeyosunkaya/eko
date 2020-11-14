module github.com/egeyosunkaya/EKO

go 1.15

require (
	github.com/egeyosunkaya/EKO/netmodules v1.0.0
	github.com/gizak/termui/v3 v3.1.0
	github.com/marcusolsson/tui-go v0.4.0
)

replace github.com/egeyosunkaya/EKO/netmodules => ./cmd/netmodules
