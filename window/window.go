package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Window 普通窗口.
type Window struct {
	windowBase
}

// NewWindow 窗口_创建.
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口标题.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式: xcc.Window_Style_.
//	@return *Window
//
func NewWindow(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// NewWindowEx 窗口_创建扩展.
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口名.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式, xcc.Window_Style_.
//	@return *Window
//
func NewWindowEx(dwExStyle int, dwStyle int, lpClassName string, x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_CreateEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// Attach 窗口_附加窗口, 返回窗口对象.
//	@param hWnd 要附加的外部窗口句柄.
//	@param XCStyle 窗口样式: xcc.Window_Style_.
//	@return *Window
//
func Attach(hWnd int, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Attach(hWnd, XCStyle))
	return p
}

// NewWindowByHandle 从句柄创建对象.
//	@param hWindow 窗口资源句柄.
//	@return *Window
//
func NewWindowByHandle(hWindow int) *Window {
	p := &Window{}
	p.SetHandle(hWindow)
	return p
}

// NewWindowByLayout 从布局文件创建对象, 失败返回nil.
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func NewWindowByLayout(pFileName string, hParent, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewWindowByLayoutZip 从压缩包中的布局文件创建对象, 失败返回nil.
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func NewWindowByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewWindowByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回nil.
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func NewWindowByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewWindowByLayoutStringW 从布局文件字符串W创建对象, 失败返回nil.
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
//
func NewWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd int) *Window {
	handle := xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewWindowByName 从name创建对象, 失败返回nil.
//	@param name
//	@return *Window
//
func NewWindowByName(name string) *Window {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewWindowByUID 从UID创建对象, 失败返回nil.
//	@param nUID
//	@return *Window
//
func NewWindowByUID(nUID int) *Window {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewWindowByUIDName 从UID名称创建对象, 失败返回nil.
//	@param name
//	@return *Window
//
func NewWindowByUIDName(name string) *Window {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//	@param bEnable 是否启用.
//	@return int
//
func (w *Window) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(w.Handle, bEnable)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//	@param bEnable 是否启用.
//	@return int
//
func (w *Window) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(w.Handle, bEnable)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//	@param bEnable 是否启用.
//	@return int
//
func (w *Window) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(w.Handle, bEnable)
}

// SetAlignH 布局盒子_置水平对齐.
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
//
func (w *Window) SetAlignH(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignH(w.Handle, nAlign)
}

// SetAlignV 布局盒子_置垂直对齐.
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
//
func (w *Window) SetAlignV(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignV(w.Handle, nAlign)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
//
func (w *Window) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {
	return xc.XLayoutBox_SetAlignBaseline(w.Handle, nAlign)
}

// SetSpace 布局盒子_置间距.
//	@param nSpace 项间距大小.
//	@return int
//
func (w *Window) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(w.Handle, nSpace)
}

// SetSpaceRow 布局盒子_置行距.
//	@param nSpace 行间距大小.
//	@return int
//
func (w *Window) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(w.Handle, nSpace)
}
