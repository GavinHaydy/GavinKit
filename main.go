package main

import (
	"GavinKit/tiktok"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func makeTablePage() fyne.CanvasObject {
	data := [][]string{
		{"ID", "Name", "Age"},
		{"1", "Alice", "23"},
		{"2", "Bob", "30"},
		{"3", "Cathy", "28"},
		{"4", "David", "35"},
		{"5", "Eva", "22"},
		{"6", "Frank", "29"},
		{"7", "Grace", "26"},
		{"8", "Henry", "31"},
		{"9", "Ivy", "27"},
		{"10", "Jack", "24"},
		{"11", "Karen", "33"},
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			entry := widget.NewEntry()
			entry.Wrapping = fyne.TextWrapOff
			return entry
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Entry).SetText(data[tci.Row][tci.Col])
		},
	)

	for i := range len(data[0]) {
		table.SetColumnWidth(i, 200)
	}

	return table
}
func makeStream() fyne.CanvasObject {

	// 用于显示地址和流信息的标签
	urlLabel := widget.NewEntry()
	urlLabel.SetText("")
	streamLabel := widget.NewEntry()
	streamLabel.SetText("")

	// 定义一个刷新函数
	updateStream := func(types string) {
		go func() {
			data := tiktok.GetStreamAddress(types)
			if len(data) == 0 {
				log.Println(fmt.Errorf("%s is not a stream", types))
				urlLabel.SetText("无数据")
				streamLabel.SetText("")
				return
			}
			fyne.CurrentApp().Driver().DoFromGoroutine(func() {
				urlLabel.SetText(data[0])
				streamLabel.SetText(data[1])
			}, false)
		}()

	}

	// 两个按钮
	btnLine := widget.NewButton("Line", func() {
		updateStream("line")
	})
	btnWifi := widget.NewButton("Wi-Fi", func() {
		updateStream("Wi-Fi")
	})

	// 先初始化一次默认数据
	updateStream("line")

	// 布局：按钮在上，数据在下
	return container.NewVBox(
		container.NewHBox(btnLine, btnWifi),
		container.NewHBox(urlLabel, streamLabel),
	)
}

func main() {

	a := app.New()
	w := a.NewWindow("Fyne 左侧菜单示例")

	// 右侧内容容器，初始空
	content := container.NewStack()

	// 定义几个“页面”内容

	aboutPage := widget.NewLabel("这是关于页")

	// 切换页面函数
	showPage := func(page fyne.CanvasObject) {
		content.Objects = []fyne.CanvasObject{page}
		content.Refresh()
	}

	// 菜单项
	items := []string{"首页", "抖音推流获取", "关于"}

	// 菜单列表
	list := widget.NewList(
		func() int {
			return len(items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(items[i])
		},
	)

	// 监听选中事件
	list.OnSelected = func(id widget.ListItemID) {
		switch id {
		case 0:
			showPage(makeTablePage())
		case 1:
			showPage(makeStream())
		case 2:
			showPage(aboutPage)
		}
	}

	// 默认显示首页页面，默认选中首页菜单项
	list.Select(0)
	showPage(makeTablePage())

	// 左侧菜单 + 右侧内容布局
	layout := container.NewBorder(nil, nil, list, nil, content)

	w.SetContent(layout)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
