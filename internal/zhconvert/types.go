package zhconvert

import "fmt"

type ConvertRequest struct {
	Text      *string          `json:"text"`
	Converter ConvertConverter `json:"converter"`

	IgnoreTextStyles          *string                        `json:"ignoreTextStyle,omitempty"`
	JpTextStyle               *string                        `json:"jpTextStyles,omitempty"`
	JpStyleConversionStrategy *string                        `json:"jpStyleConversionStrategy,omitempty"`
	JpTextConversionStrategy  *string                        `json:"jpTextConversionStrategy,omitempty"`
	Modules                   *map[string]ConvertModuleState `json:"modules,omitempty"`
	UserPostReplace           *string                        `json:"userPostReplace,omitempty"`
	UserPreReplace            *string                        `json:"userPreReplace,omitempty"`
	UserProtectReplace        *string                        `json:"userProtectReplace,omitempty"`

	DiffCharLevel         *bool                `json:"diffCharLevel,omitempty"`
	DiffContextLines      *int                 `json:"diffContextLines,omitempty"`
	DiffEnable            *bool                `json:"diffEnable,omitempty"`
	DiffIgnoreCase        *bool                `json:"diffIgnoreCase,omitempty"`
	DiffIgnoreWhiteSpaces *bool                `json:"diffIgnoreWhiteSpaces,omitempty"`
	DiffTemplate          *ConvertDiffTemplate `json:"diffTemplate,omitempty"`

	CleanUpText             *bool `json:"cleanUpText,omitempty"`
	EnsureNewlineAtEof      *bool `json:"ensureNewlineAtEof,omitempty"`
	TranslateTabsToSpaces   *int  `json:"translateTabsToSpaces,omitempty"`
	TrimTrailingWhiteSpaces *bool `json:"trimTrailingWhiteSpaces,omitempty"`
	UnifyLeadingHyphen      *bool `json:"unifyLeadingHyphen,omitempty"`
}

func (c ConvertRequest) Validate() error {
	if err := c.Converter.Validate(); err != nil {
		return err
	}

	if c.Modules != nil {
		for _, value := range *c.Modules {
			if err := value.Validate(); err != nil {
				return err
			}
		}
	}

	if c.DiffTemplate != nil {
		if err := c.DiffTemplate.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type ConvertApiResponse struct {
	Code      int             `json:"code"`
	Data      ConvertResponse `json:"data"`
	Msg       string          `json:"msg"`
	Revisions struct {
		Build string `json:"build"`
		Msg   string `json:"msg"`
		Time  int    `json:"time"`
	} `json:"revisions"`
	ExecTime float64 `json:"execTime"`
}

type ConvertResponse struct {
	Converter    ConvertConverter `json:"converter"`
	Text         string           `json:"text"`
	Diff         *string          `json:"diff"`
	UsedModules  []string         `json:"usedModules"`
	JpTextStyles []string         `json:"jpTextStyles"`
	TextFormat   string           `json:"textFormat"`
}

type ConvertConverter string

// Supported converter list:
// https://docs.zhconvert.org/api/convert/#%E5%BF%85%E5%A1%AB
const (
	ConvertConverterSimplified      ConvertConverter = "Simplified"
	ConvertConverterTraditional     ConvertConverter = "Traditional"
	ConvertConverterChina           ConvertConverter = "China"
	ConvertConverterHongkong        ConvertConverter = "Hongkong"
	ConvertConverterTaiwan          ConvertConverter = "Taiwan"
	ConvertConverterPinyin          ConvertConverter = "Pinyin"
	ConvertConverterBopomofo        ConvertConverter = "Bopomofo"
	ConvertConverterMars            ConvertConverter = "Mars"
	ConvertConverterWikiSimplified  ConvertConverter = "WikiSimplified"
	ConvertConverterWikiTraditional ConvertConverter = "WikiTraditional"
)

func (c ConvertConverter) Validate() error {
	switch c {
	case ConvertConverterSimplified:
	case ConvertConverterTraditional:
	case ConvertConverterChina:
	case ConvertConverterHongkong:
	case ConvertConverterTaiwan:
	case ConvertConverterPinyin:
	case ConvertConverterBopomofo:
	case ConvertConverterMars:
	case ConvertConverterWikiSimplified:
	case ConvertConverterWikiTraditional:
	default:
		return fmt.Errorf("Invalid Converter %s", c)
	}

	return nil
}

type ConvertModuleState int8

const (
	ConvertModuleStateAuto     ConvertModuleState = -1
	ConvertModuleStateDisabled ConvertModuleState = 0
	ConvertModuleStateEnabled  ConvertModuleState = 1
)

func (c ConvertModuleState) Validate() error {
	switch c {
	case ConvertModuleStateAuto:
	case ConvertModuleStateDisabled:
	case ConvertModuleStateEnabled:
	default:
		return fmt.Errorf("Invalid ModuleState %d", c)
	}

	return nil
}

type ConvertDiffTemplate string

const (
	ConvertDiffTemplateInline     ConvertDiffTemplate = "Inline"
	ConvertDiffTemplateSideBySide ConvertDiffTemplate = "SideBySide"
	ConvertDiffTemplateUnified    ConvertDiffTemplate = "Unified"
	ConvertDiffTemplateJsonHtml   ConvertDiffTemplate = "JsonHtml"
	ConvertDiffTemplateJsonText   ConvertDiffTemplate = "JsonText"
)

func (c ConvertDiffTemplate) Validate() error {
	switch c {
	case ConvertDiffTemplateInline:
	case ConvertDiffTemplateSideBySide:
	case ConvertDiffTemplateUnified:
	case ConvertDiffTemplateJsonHtml:
	case ConvertDiffTemplateJsonText:
	default:
		return fmt.Errorf("Invalid DiffTemplate %s", c)
	}

	return nil
}
