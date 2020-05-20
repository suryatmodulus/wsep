package proto

const (
	TypeStart      = "start"
	TypeResize     = "resize_header"
	TypeStdin      = "stdin"
	TypeCloseStdin = "close_stdin"
)

type ClientStartHeader struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
	TTY     bool     `json:"tty"`
}

type ClientResizeHeader struct {
	Rows uint16 `json:"rows"`
	Cols uint16 `json:"cols"`
}
