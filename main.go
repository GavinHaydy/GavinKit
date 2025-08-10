//package main
//
//import (
//	"GavinKit/tiktok"
//	"gioui.org/app"
//	"gioui.org/layout"
//	"gioui.org/op"
//	"gioui.org/text"
//	"gioui.org/unit"
//	"gioui.org/widget/material"
//	"image/color"
//	"log"
//	"os"
//	"time"
//)
//
//func run(window *app.Window) error {
//	theme := material.NewTheme()
//	//data := tiktok.GetStreamAddress("line")
//	var ops op.Ops
//	data := &[]string{"正在监听流地址...（请稍等）"}
//
//	go func() {
//		// 模拟等待或网卡监听延迟
//		time.Sleep(1 * time.Second)
//
//		res := tiktok.GetStreamAddress("line")
//		if len(res) == 0 {
//			res = []string{"未获取到流地址"}
//		}
//		*data = res
//
//		// 请求窗口重新绘制
//		window.Invalidate()
//	}()
//
//	for {
//		switch e := window.Event().(type) {
//		case app.DestroyEvent:
//			return e.Err
//
//		case app.FrameEvent:
//			gtx := app.NewContext(&ops, e)
//
//			// 使用 layout.Flex 垂直布局所有行
//			layout.Flex{
//				Axis: layout.Vertical,
//			}.Layout(gtx,
//
//				// 标题
//				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//					title := material.H3(theme, "抖音推流地址")
//					title.Alignment = text.Middle
//					title.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
//					return title.Layout(gtx)
//				}),
//
//				// 内容
//				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
//					dims := layout.Dimensions{}
//					for _, line := range *data {
//						lbl := material.Body1(theme, line)
//						lbl.Color = color.NRGBA{R: 50, G: 50, B: 150, A: 255}
//						lbl.Alignment = text.Start
//						gtx.Constraints.Min.X = gtx.Dp(unit.Dp(20))
//						dims = lbl.Layout(gtx)
//					}
//					return dims
//				}),
//			)
//
//			e.Frame(gtx.Ops)
//		}
//	}
//}
//
//func main() {
//	go func() {
//		window := new(app.Window)
//		err := run(window)
//		if err != nil {
//			log.Fatal(err)
//		}
//		os.Exit(0)
//	}()
//	app.Main()
//}

package main

import (
	"GavinKit/tiktok"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"log"
	"os"
)

// }
var (
	editors []*widget.Editor // 👈 持久化 editor 实例
)

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	// 初始内容
	data := &[]string{"正在监听流地址...（请稍等）"}

	// 初始化 editors（与 data 长度对应）
	editors = make([]*widget.Editor, len(*data))
	for i, line := range *data {
		ed := new(widget.Editor)
		ed.SetText(line)
		ed.ReadOnly = true
		editors[i] = ed
	}

	// 开启监听协程
	go func() {
		res := tiktok.GetStreamAddress("line")
		log.Printf("监听完成，结果如下：%v\n", res)

		*data = res

		// 更新编辑器内容
		editors = make([]*widget.Editor, len(res))
		for i, line := range res {
			ed := new(widget.Editor)
			ed.SetText(line)
			ed.ReadOnly = true
			editors[i] = ed
		}

		window.Invalidate()
	}()

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			var children []layout.FlexChild

			// 标题
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				title := material.H3(theme, "抖音推流地址")
				title.Alignment = text.Middle
				title.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
				return title.Layout(gtx)
			}))

			// 每一行流地址使用可复制的 Editor
			for _, ed := range editors {
				editorStyle := material.Editor(theme, ed, "")
				editorStyle.TextSize = unit.Sp(14)
				editorStyle.Color = color.NRGBA{R: 50, G: 50, B: 150, A: 255}

				children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return editorStyle.Layout(gtx)
				}))
			}

			// 布局
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx, children...)

			e.Frame(gtx.Ops)
		}
	}
}

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
