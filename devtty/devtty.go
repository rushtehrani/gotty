package devtty

import (
	"io"
	"os"

	"github.com/kr/pty"
	"github.com/pkg/errors"
)

type DevTTY struct {
	connection io.ReadCloser
	pty        *os.File
	tty        *os.File
}

func New(connection io.ReadWriteCloser) (*DevTTY, error) {
	pty, tty, err := pty.Open()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open pty")
	}

	return &DevTTY{
		connection: connection,
		pty:        pty,
		tty:        tty,
	}, nil
}

func (dt *DevTTY) Out() *os.File {
	return dt.tty
}

func (dt *DevTTY) In() io.Reader {
	return dt.tty
}

func (dt *DevTTY) Start() {
	go func() {
		for {

		}
	}()

	go func() {
		for {

		}
	}()
}

func processReceive(reader io.Reader, buffer []byte) (int, err) {
	n, err := reader.Read(buffer)
	if err != nil {
		return errors.Wrap(err, "failed to read")
	}

	if n == 0 {
		return errors.New("unexpected zero length read")
	}

	switch biffer[0] {
	case Input:
		if !context.app.options.PermitWrite {
			break
		}

		_, err := context.InputWriter().Write(data[1:])
		if err != nil {
			return
		}

	case Ping:
		if err := context.write([]byte{Pong}); err != nil {
			log.Print(err.Error())
			return
		}
	case ResizeTerminal:
		var args argResizeTerminal
		err = json.Unmarshal(data[1:], &args)
		if err != nil {
			log.Print("Malformed remote command")
			return
		}
		rows := uint16(context.app.options.Height)
		if rows == 0 {
			rows = uint16(args.Rows)
		}

		columns := uint16(context.app.options.Width)
		if columns == 0 {
			columns = uint16(args.Columns)
		}

		err = context.ResizeTerminal(columns, rows)
		if err != nil {
			log.Printf("failed to resize terminal %v", err)
			return
		}
	default:
		log.Print("Unknown message type")
		return
	}
}
