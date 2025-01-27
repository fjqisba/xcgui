package font

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Font 炫彩字体.
type Font struct {
	objectbase.ObjectBase
}

// NewFont 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放.
//	@param size 字体大小,单位(pt,磅).
//	@return *Font 返回字体对象.
//
func NewFont(size int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_Create(size))
	return p
}

// NewFontEX 字体_创建扩展. 创建炫彩字体.
//	@param pName 字体名称.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return *Font 返回字体对象.
//
func NewFontEX(pName string, size int, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateEx(pName, size, style))
	return p
}

// NewFontLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体.
//	@param pFontInfo 字体信息.
//	@return *Font 返回字体对象.
//
func NewFontLOGFONTW(pFontInfo *xc.LOGFONTW) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateLOGFONTW(pFontInfo))
	return p
}

// NewFontFromHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体.
//	@param hFont 字体句柄.
//	@return *Font 返回字体对象.
//
func NewFontFromHFONT(hFont int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromHFONT(hFont))
	return p
}

// NewFontFromFont 字体_创建从Font. 创建炫彩字体从GDI+字体(Font).
//	@param pFont GDI+字体指针(Font*).
//	@return *Font 返回字体对象.
//
func NewFontFromFont(pFont int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromFont(pFont))
	return p
}

// NewFontFromFile 字体_创建从文件. 创建字体从文件.
//	@param pFontFile 字体文件名.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return *Font 返回字体对象.
//
func NewFontFromFile(pFontFile string, size int, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromFile(pFontFile, size, style))
	return p
}

// NewFontFromZip 字体_创建从ZIP.
//	@param pZipFileName zip文件名.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return *Font 返回炫彩字体对象.
//
func NewFontFromZip(pZipFileName, pFileName, pPassword string, fontSize int, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromZip(pZipFileName, pFileName, pPassword, fontSize, style))
	return p
}

// NewFontFromZipMem 字体_创建从内存ZIP.
//	@param data zip数据.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return *Font 返回炫彩字体对象.
//
func NewFontFromZipMem(data []byte, pFileName, pPassword string, fontSize int, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromZipMem(data, pFileName, pPassword, fontSize, style))
	return p
}

// NewFontFromMem 字体_创建从内存. 创建炫彩字体从内存.
//	@param data 字体文件数据.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return *Font 返回字体对象.
//
func NewFontFromMem(data []byte, fontSize int, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromMem(data, fontSize, style))
	return p
}

// NewFontFromRes 字体_创建从资源. 创建字体从资源.
//	@param id xx.
//	@param pType xx.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@param hModule xx.
//	@return *Font 返回字体对象.
//
func NewFontFromRes(id int, pType string, fontSize int, style xcc.FontStyle_, hModule int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromRes(id, pType, fontSize, style, hModule))
	return p
}

// NewFontByHandle 从句柄创建对象.
//	@param handle
//	@return *Font
//
func NewFontByHandle(handle int) *Font {
	p := &Font{}
	p.SetHandle(handle)
	return p
}

// NewFontByName 根据资源文件中的name创建对象, 失败返回nil.
//	@param name
//	@return *Font
//
func NewFontByName(name string) *Font {
	handle := res.GetFont(name)
	if handle > 0 {
		p := &Font{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// EnableAutoDestroy 字体_启用自动销毁. 是否自动销毁.
//	@param bEnable 是否启用.
//	@return int
//
func (f *Font) EnableAutoDestroy(bEnable bool) int {
	return xc.XFont_EnableAutoDestroy(f.Handle, bEnable)
}

// GetFont 字体_取Font. 获取字体.
//	@return int 返回GDI+ Font指针
//
func (f *Font) GetFont() int {
	return xc.XFont_GetFont(f.Handle)
}

// GetFontInfo 字体_取信息. 获取字体信息.
//	@param pInfo 接收返回的字体信息.
//	@return int
//
func (f *Font) GetFontInfo(pInfo *xc.Font_Info_) int {
	return xc.XFont_GetFontInfo(f.Handle, pInfo)
}

// GetLOGFONTW 字体_取LOGFONTW. 获取字体LOGFONTW.
//	@param hdc hdc句柄.
//	@param pOut 接收返回信息.
//	@return bool
//
func (f *Font) GetLOGFONTW(hdc int, pOut *xc.LOGFONTW) bool {
	return xc.XFont_GetLOGFONTW(f.Handle, hdc, pOut)
}

//
// Destroy 字体_销毁. 强制销毁炫彩字体, 谨慎使用, 建议使用 Release() 释放.
//	@return int
//
func (f *Font) Destroy() int {
	return xc.XFont_Destroy(f.Handle)
}

// AddRef 字体_增加引用计数.
//	@return int
//
func (f *Font) AddRef() int {
	return xc.XFont_AddRef(f.Handle)
}

// GetRefCount 字体_取引用计数.
//	@return int
//
func (f *Font) GetRefCount() int {
	return xc.XFont_GetRefCount(f.Handle)
}

// Release 字体_释放引用计数. 释放引用计数, 当引用计数为0时自动销毁.
//	@return int
//
func (f *Font) Release() int {
	return xc.XFont_Release(f.Handle)
}
