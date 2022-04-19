package bar

import (
	"fmt"
	"io"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

// Mpb 该实现用于在实现 copy 的同时在终端显示进度
type Mpb struct {
	p *mpb.Progress
}

// NewMpb .
func NewMpb() *Mpb {
	return &Mpb{
		p: mpb.New(
			mpb.WithWidth(120),
			mpb.WithRefreshRate(180*time.Millisecond),
		),
	}
}

// Wait .
func (m *Mpb) Wait() {
	m.p.Wait()
}

// Copy .
func (m *Mpb) Copy(fileName string, fileSize int64, dst io.Writer, src io.Reader) (written int64, err error) {
	bar := m.p.New(fileSize, mpb.BarStyle().Rbound("|"),
		mpb.PrependDecorators(
			decor.CountersKibiByte("% .1f / % .1f", decor.WCSyncWidthR),
		),
		mpb.AppendDecorators(
			decor.OnComplete(
				decor.EwmaETA(decor.ET_STYLE_GO, 60), "done",
			),
			decor.Name(" ] "),
			decor.EwmaSpeed(decor.UnitKiB, "% 2.1f", 60, decor.WCSyncWidthR),
			decor.Name(fmt.Sprintf(" | %-20s", fileName), decor.WCSyncWidthR),
		))
	proxyReader := bar.ProxyReader(src)
	defer func() {
		_ = proxyReader.Close()
	}()
	return io.Copy(dst, proxyReader)
}
