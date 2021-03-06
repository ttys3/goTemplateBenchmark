// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"github.com/SlinSo/goTemplateBenchmark/model"
	pool "github.com/valyala/bytebufferpool"
)

const (
	index__0  = `<!DOCTYPE html><html><body><header><title>`
	index__1  = `'s Home Page</title><div class="header">Page Header</div></header><nav><ul class="navigation">`
	index__2  = `</ul></nav><section><div class="content"><div class="welcome"><h4>Hello `
	index__3  = `</h4><div class="raw">`
	index__4  = `</div><div class="enc">`
	index__5  = `</div></div>`
	index__6  = `</div></section><footer><div class="footer">copyright 2016</div></footer></body></html>`
	index__7  = `<li><a href="`
	index__8  = `">`
	index__9  = `</a></li>`
	index__10 = `<p>`
	index__11 = ` has `
	index__12 = ` message</p>`
	index__15 = ` messages</p>`
)

func Index(u *model.User, nav []*model.Navigation, title string, buffer *pool.ByteBuffer) {

	buffer.WriteString(index__0)
	WriteEscString(title, buffer)
	buffer.WriteString(index__1)

	for _, item := range nav {
		buffer.WriteString(index__7)
		WriteEscString(item.Link, buffer)
		buffer.WriteString(index__8)
		WriteEscString(item.Item, buffer)
		buffer.WriteString(index__9)

	}
	buffer.WriteString(index__2)
	WriteEscString(u.FirstName, buffer)
	buffer.WriteString(index__3)
	buffer.WriteString(u.RawContent)
	buffer.WriteString(index__4)
	WriteEscString(u.EscapedContent, buffer)
	buffer.WriteString(index__5)

	for i := 1; i <= 5; i++ {
		if i == 1 {
			buffer.WriteString(index__10)
			WriteEscString(u.FirstName, buffer)
			buffer.WriteString(index__11)
			WriteInt(int64(i), buffer)
			buffer.WriteString(index__12)

		} else {
			buffer.WriteString(index__10)
			WriteEscString(u.FirstName, buffer)
			buffer.WriteString(index__11)
			WriteInt(int64(i), buffer)
			buffer.WriteString(index__15)

		}
	}
	buffer.WriteString(index__6)

}
