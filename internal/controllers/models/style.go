package models

import (
	pb "protos/maps"
	"time"
)

type Style struct {
	ID string `json:"id"`

	StyleName                 string  `json:"style_name"`
	StyleType                 string  `json:"style_type"`
	StyleSourceLayer          string  `json:"style_source_layer"`
	StyleFilter_1             bool    `json:"style_filter_1"`
	StyleFilterField_1        string  `json:"style_filter_field_1"`
	StyleFilter_2             bool    `json:"style_filter_2"`
	StyleFilterField_2        string  `json:"style_filter_field_2"`
	StyleFilterValues         string  `json:"style_filter_values"`
	StyleMaxZoom              float32 `json:"style_max_zoom"`
	StyleMinZoom              float32 `json:"style_min_zoom"`
	StyleLabel                bool    `json:"style_label"`
	LabelTextColor            string  `json:"label_text_color"`
	LabelTextHaloWidth        float32 `json:"label_text_halo_width"`
	LabelTextHaloBlur         float32 `json:"label_text_halo_blur"`
	LabelTextHaloColor        string  `json:"label_text_halo_color"`
	LabelTextField            string  `json:"label_text_field"`
	LabelTextFont             string  `json:"label_text_font"`
	LabelTextOffset           string  `json:"label_text_offset"`
	LabelTextOpacity          string  `json:"label_text_opacity"`
	LabelTextJustify          string  `json:"label_text_justify"`
	LabelTextLineHeight       float32 `json:"label_text_line_height"`
	LabelTextIgnorePlacement  bool    `json:"label_text_ignore_placement"`
	LabelTextPadding          int32   `json:"label_text_padding"`
	LabelTextRotate           float32 `json:"label_text_rotate"`
	LabelTextSize             float32 `json:"label_text_size"`
	LabelTextTransform        string  `json:"label_text_transform"`
	FillAntialias             string  `json:"fill_antialias"`
	FillColor                 string  `json:"fill_color"`
	FillOpacity               string  `json:"fill_opacity"`
	FillOutlineColor          string  `json:"fill_outline_color"`
	FillPattern               string  `json:"fill_pattern"`
	FillVisibility            string  `json:"fill_visibility"`
	LineBlur                  string  `json:"line_blur"`
	LineColor                 string  `json:"line_color"`
	LineGapWidth              string  `json:"line_gap_width"`
	LineOpacity               string  `json:"line_opacity"`
	LineWidth                 string  `json:"line_width"`
	LinePattern               string  `json:"line_pattern"`
	LineDasharray             string  `json:"line_dasharray"`
	LineCap                   string  `json:"line_cap"`
	LineJoin                  string  `json:"line_join"`
	LineVisibility            string  `json:"line_visibility"`
	SymbolTextAllowOverlap    string  `json:"symbol_text_allow_overlap"`
	SymbolTextColor           string  `json:"symbol_text_color"`
	SymbolTextField           string  `json:"symbol_text_field"`
	SymbolTextFont            string  `json:"symbol_text_font"`
	SymbolTextHaloBlur        string  `json:"symbol_text_halo_blur"`
	SymbolTextHaloColor       string  `json:"symbol_text_halo_color"`
	SymbolTextHaloWidth       string  `json:"symbol_text_halo_width"`
	SymbolTextIgnorePlacement string  `json:"symbol_text_ignore_placement"`
	SymbolTextJustify         string  `json:"symbol_text_justify"`
	SymbolTextRotate          string  `json:"symbol_text_rotate"`
	SymbolTextSize            string  `json:"symbol_text_size"`
	SymbolTextOffset          string  `json:"symbol_text_offset"`
	SymbolTextOpacity         string  `json:"symbol_text_opacity"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (s *Style) ToMStyle() *pb.MStyle {
	return &pb.MStyle{
		Id:                        s.ID,
		StyleName:                 s.StyleName,
		StyleType:                 s.StyleType,
		StyleSourceLayer:          s.StyleSourceLayer,
		StyleFilter_1:             s.StyleFilter_1,
		StyleFilterField_1:        s.StyleFilterField_1,
		StyleFilter_2:             s.StyleFilter_2,
		StyleFilterField_2:        s.StyleFilterField_2,
		StyleFilterValues:         s.StyleFilterValues,
		StyleMaxZoom:              s.StyleMaxZoom,
		StyleMinZoom:              s.StyleMinZoom,
		StyleLabel:                s.StyleLabel,
		LabelTextColor:            s.LabelTextColor,
		LabelTextHaloWidth:        s.LabelTextHaloWidth,
		LabelTextHaloBlur:         s.LabelTextHaloBlur,
		LabelTextHaloColor:        s.LabelTextHaloColor,
		LabelTextField:            s.LabelTextField,
		LabelTextFont:             s.LabelTextFont,
		LabelTextOffset:           s.LabelTextOffset,
		LabelTextOpacity:          s.LabelTextOpacity,
		LabelTextJustify:          s.LabelTextJustify,
		LabelTextLineHeight:       s.LabelTextLineHeight,
		LabelTextIgnorePlacement:  s.LabelTextIgnorePlacement,
		LabelTextPadding:          s.LabelTextPadding,
		LabelTextRotate:           s.LabelTextRotate,
		LabelTextSize:             s.LabelTextSize,
		LabelTextTransform:        s.LabelTextTransform,
		FillAntialias:             s.FillAntialias,
		FillColor:                 s.FillColor,
		FillOpacity:               s.FillOpacity,
		FillOutlineColor:          s.FillOutlineColor,
		FillPattern:               s.FillPattern,
		FillVisibility:            s.FillVisibility,
		LineBlur:                  s.LineBlur,
		LineColor:                 s.LineColor,
		LineGapWidth:              s.LineGapWidth,
		LineOpacity:               s.LineOpacity,
		LineWidth:                 s.LineWidth,
		LinePattern:               s.LinePattern,
		LineDasharray:             s.LineDasharray,
		LineCap:                   s.LineCap,
		LineJoin:                  s.LineJoin,
		LineVisibility:            s.LineVisibility,
		SymbolTextAllowOverlap:    s.SymbolTextAllowOverlap,
		SymbolTextColor:           s.SymbolTextColor,
		SymbolTextField:           s.SymbolTextField,
		SymbolTextFont:            s.SymbolTextFont,
		SymbolTextHaloBlur:        s.SymbolTextHaloBlur,
		SymbolTextHaloColor:       s.SymbolTextHaloColor,
		SymbolTextHaloWidth:       s.SymbolTextHaloWidth,
		SymbolTextIgnorePlacement: s.SymbolTextIgnorePlacement,
		SymbolTextJustify:         s.SymbolTextJustify,
		SymbolTextRotate:          s.SymbolTextRotate,
		SymbolTextSize:            s.SymbolTextSize,
		SymbolTextOffset:          s.SymbolTextOffset,
		SymbolTextOpacity:         s.SymbolTextOpacity,
		CreateUserIp:              s.CreateUserIP,
		UpdateUserIp:              s.UpdateUserIP,
		CreateUserId:              s.CreateUserID,
		UpdateUserId:              s.UpdateUserID,
		CreatedAt:                 s.CreatedAt.String(),
		UpdatedAt:                 s.UpdatedAt.String(),
	}
}
