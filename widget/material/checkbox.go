// SPDX-License-Identifier: Unlicense OR MIT

package material

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

type CheckBoxStyle struct {
	checkable
}

func CheckBox(th *Theme, label string) CheckBoxStyle {
	return CheckBoxStyle{
		checkable{
			Label:              label,
			Color:              th.Color.Text,
			IconColor:          th.Color.Primary,
			TextSize:           th.TextSize.Scale(14.0 / 16.0),
			Size:               unit.Dp(26),
			shaper:             th.Shaper,
			checkedStateIcon:   th.checkBoxCheckedIcon,
			uncheckedStateIcon: th.checkBoxUncheckedIcon,
		},
	}
}

// Layout updates the checkBox and displays it.
func (c CheckBoxStyle) Layout(gtx layout.Context, checkBox *widget.Bool) layout.Dimensions {
	checkBox.Update(gtx)
	dims := c.layout(gtx, checkBox.Value)
	checkBox.Layout(gtx)
	return dims
}
