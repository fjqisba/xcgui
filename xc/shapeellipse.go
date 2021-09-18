package xc

// 形状圆_创建, 创建圆形形状对象, 返回句柄
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func XShapeEllipse_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xShapeEllipse_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 形状圆_置边框色
// hShape: 形状对象句柄.
// color: RGB颜色值.
// alpha: 透明度.
func XShapeEllipse_SetBorderColor(hShape int, color int, alpha uint8) int {
	r, _, _ := xShapeEllipse_SetBorderColor.Call(uintptr(hShape), uintptr(color), uintptr(alpha))
	return int(r)
}

// 形状圆_置填充色
// hShape: 形状对象句柄.
// color: RGB颜色值.
// alpha: 透明度.
func XShapeEllipse_SetFillColor(hShape int, color int, alpha uint8) int {
	r, _, _ := xShapeEllipse_SetFillColor.Call(uintptr(hShape), uintptr(color), uintptr(alpha))
	return int(r)
}

// 形状圆_启用边框, 启用绘制圆边框
// hShape: 形状对象句柄.
// bEnable: 是否启用.
func XShapeEllipse_EnableBorder(hShape int, bEnable bool) int {
	r, _, _ := xShapeEllipse_EnableBorder.Call(uintptr(hShape), boolPtr(bEnable))
	return int(r)
}

// 形状圆_启用填充, 启用填充圆
// hShape: 形状对象句柄.
// bEnable: 是否启用.
func XShapeEllipse_EnableFill(hShape int, bEnable bool) int {
	r, _, _ := xShapeEllipse_EnableFill.Call(uintptr(hShape), boolPtr(bEnable))
	return int(r)
}