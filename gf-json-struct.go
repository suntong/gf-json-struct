type AutoGenerated struct {
	Button struct {
		ConditionalLogic     string `json:"conditionalLogic"`
		ID                   string `json:"id"`
		ImageURL             string `json:"imageUrl"`
		LayoutGridColumnSpan int    `json:"layoutGridColumnSpan"`
		Location             string `json:"location"`
		Text                 string `json:"text"`
		Type                 string `json:"type"`
		Width                string `json:"width"`
	} `json:"button"`
	CSSClass                string `json:"cssClass"`
	CustomRequiredIndicator string `json:"customRequiredIndicator"`
	DateCreated             string `json:"date_created"`
	Deprecated              string `json:"deprecated"`
	Description             string `json:"description"`
	DescriptionPlacement    string `json:"descriptionPlacement"`
	EnableAnimation         bool   `json:"enableAnimation"`
	EnableHoneypot          bool   `json:"enableHoneypot"`
	Feeds                   []any  `json:"feeds"`
	Fields                  []struct {
		AdminLabel               string `json:"adminLabel"`
		AllowsPrepopulate        bool   `json:"allowsPrepopulate"`
		AutocompleteAttribute    string `json:"autocompleteAttribute"`
		CalculationFormula       string `json:"calculationFormula"`
		CalculationRounding      string `json:"calculationRounding"`
		Choices                  string `json:"choices"`
		ConditionalLogic         string `json:"conditionalLogic"`
		CSSClass                 string `json:"cssClass"`
		DefaultValue             string `json:"defaultValue"`
		Description              string `json:"description"`
		DescriptionPlacement     string `json:"descriptionPlacement"`
		DisableQuantity          bool   `json:"disableQuantity"`
		DisplayAllCategories     bool   `json:"displayAllCategories"`
		DisplayOnly              string `json:"displayOnly"`
		EnableAutocomplete       bool   `json:"enableAutocomplete"`
		EnableCalculation        string `json:"enableCalculation"`
		EnableEnhancedUI         int    `json:"enableEnhancedUI"`
		EnablePasswordInput      string `json:"enablePasswordInput,omitempty"`
		ErrorMessage             string `json:"errorMessage"`
		Errors                   []any  `json:"errors"`
		Fields                   string `json:"fields"`
		FormID                   int    `json:"formId"`
		GkCustomMapIconEnabled   bool   `json:"gk_custom_map_icon_enabled"`
		GpDateTimeCalculatorUnit bool   `json:"gp-date-time-calculator_unit"`
		GppaChoicesFilterGroups  []any  `json:"gppa-choices-filter-groups,omitempty"`
		GppaChoicesTemplates     []any  `json:"gppa-choices-templates,omitempty"`
		GppaValuesFilterGroups   []any  `json:"gppa-values-filter-groups,omitempty"`
		GppaValuesObjectType     string `json:"gppa-values-object-type,omitempty"`
		GppaValuesTemplates      []any  `json:"gppa-values-templates,omitempty"`
		GwwordcountMaxWordCount  string `json:"gwwordcount_max_word_count"`
		GwwordcountMinWordCount  string `json:"gwwordcount_min_word_count"`
		ID                       int    `json:"id"`
		InputMask                bool   `json:"inputMask"`
		InputMaskIsCustom        bool   `json:"inputMaskIsCustom"`
		InputMaskValue           string `json:"inputMaskValue"`
		InputName                string `json:"inputName"`
		InputType                string `json:"inputType"`
		Inputs                   any    `json:"inputs"`
		IsRequired               bool   `json:"isRequired"`
		Label                    string `json:"label"`
		LabelPlacement           string `json:"labelPlacement"`
		LayoutGridColumnSpan     int    `json:"layoutGridColumnSpan"`
		LayoutGroupID            string `json:"layoutGroupId"`
		MaxFiles                 string `json:"maxFiles"`
		MaxLength                string `json:"maxLength"`
		MultipleFiles            bool   `json:"multipleFiles"`
		NoDuplicates             bool   `json:"noDuplicates"`
		PageNumber               int    `json:"pageNumber"`
		Placeholder              string `json:"placeholder"`
		ProductField             string `json:"productField"`
		Size                     string `json:"size"`
		SubLabelPlacement        string `json:"subLabelPlacement"`
		Type                     string `json:"type"`
		UseRichTextEditor        bool   `json:"useRichTextEditor"`
		Visibility               string `json:"visibility"`
		CheckboxLabel            string `json:"checkboxLabel,omitempty"`
		EmailConfirmEnabled      string `json:"emailConfirmEnabled,omitempty"`
	} `json:"fields"`
	FirstPageCSSClass any `json:"firstPageCssClass"`
	GfStylespro       struct {
		FooterStyle string `json:"footer_style"`
		IcnEt       string `json:"icn_et"`
		IcnEtLine   string `json:"icn_et_line"`
		IcnFa       string `json:"icn_fa"`
		IcnMd       string `json:"icn_md"`
		Iconsets    string `json:"iconsets"`
		Theme       string `json:"theme"`
		VMessage    string `json:"v_message"`
		VPopup      string `json:"v_popup"`
		VScroll     string `json:"v_scroll"`
	} `json:"gf_stylespro"`
	GfpdfFormSettings struct {
		Six385Ac7Ae1A86 struct {
			Active                      bool   `json:"active"`
			BackgroundColor             string `json:"background_color"`
			BackgroundImage             string `json:"background_image"`
			Conditional                 string `json:"conditional"`
			ConditionalLogic            any    `json:"conditionalLogic"`
			CustomPdfSize               []any  `json:"custom_pdf_size"`
			EnableConditional           string `json:"enable_conditional"`
			Filename                    string `json:"filename"`
			FirstFooter                 string `json:"first_footer"`
			FirstHeader                 string `json:"first_header"`
			FocusgravityAccentColour    string `json:"focusgravity_accent_colour"`
			FocusgravityLabelFormat     string `json:"focusgravity_label_format"`
			FocusgravitySecondaryColour string `json:"focusgravity_secondary_colour"`
			Font                        string `json:"font"`
			FontColour                  string `json:"font_colour"`
			FontSize                    int    `json:"font_size"`
			Footer                      string `json:"footer"`
			Format                      string `json:"format"`
			Header                      string `json:"header"`
			ID                          string `json:"id"`
			ImageDpi                    int    `json:"image_dpi"`
			MasterPassword              string `json:"master_password"`
			Name                        string `json:"name"`
			Notification                []any  `json:"notification"`
			Orientation                 string `json:"orientation"`
			Password                    string `json:"password"`
			PdfSize                     string `json:"pdf_size"`
			Privileges                  struct {
				AnnotForms   string `json:"annot-forms"`
				Assemble     string `json:"assemble"`
				Copy         string `json:"copy"`
				Extract      string `json:"extract"`
				FillForms    string `json:"fill-forms"`
				Modify       string `json:"modify"`
				Print        string `json:"print"`
				PrintHighres string `json:"print-highres"`
			} `json:"privileges"`
			PublicAccess       string `json:"public_access"`
			RestrictOwner      string `json:"restrict_owner"`
			Rtl                string `json:"rtl"`
			Save               string `json:"save"`
			Security           string `json:"security"`
			ShowEmpty          string `json:"show_empty"`
			ShowFormTitle      string `json:"show_form_title"`
			ShowHTML           string `json:"show_html"`
			ShowPageNames      string `json:"show_page_names"`
			ShowSectionContent string `json:"show_section_content"`
			Template           string `json:"template"`
		} `json:"6385ac7ae1a86"`
	} `json:"gfpdf_form_settings"`
	GravityviewEntryModeration string `json:"gravityview_entry_moderation"`
	GvInlineEditEnable         string `json:"gv_inline_edit_enable"`
	HoneypotAction             string `json:"honeypotAction"`
	ID                         int    `json:"id"`
	IsActive                   string `json:"is_active"`
	IsTrash                    string `json:"is_trash"`
	LabelPlacement             string `json:"labelPlacement"`
	LastPageButton             any    `json:"lastPageButton"`
	Legacy                     string `json:"legacy"`
	LimitEntries               bool   `json:"limitEntries"`
	LimitEntriesCount          string `json:"limitEntriesCount"`
	LimitEntriesMessage        string `json:"limitEntriesMessage"`
	LimitEntriesPeriod         string `json:"limitEntriesPeriod"`
	MarkupVersion              int    `json:"markupVersion"`
	NextFieldID                int    `json:"nextFieldId"`
	Pagination                 any    `json:"pagination"`
	PostContentTemplate        string `json:"postContentTemplate"`
	PostContentTemplateEnabled bool   `json:"postContentTemplateEnabled"`
	PostTitleTemplate          string `json:"postTitleTemplate"`
	PostTitleTemplateEnabled   bool   `json:"postTitleTemplateEnabled"`
	RequireLogin               bool   `json:"requireLogin"`
	RequireLoginMessage        string `json:"requireLoginMessage"`
	RequiredIndicator          string `json:"requiredIndicator"`
	Save                       struct {
		Button struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"button"`
		Enabled bool `json:"enabled"`
	} `json:"save"`
	SaveButtonText         string `json:"saveButtonText"`
	SaveEnabled            string `json:"saveEnabled"`
	ScheduleEnd            string `json:"scheduleEnd"`
	ScheduleEndAmpm        string `json:"scheduleEndAmpm"`
	ScheduleEndHour        string `json:"scheduleEndHour"`
	ScheduleEndMinute      string `json:"scheduleEndMinute"`
	ScheduleForm           bool   `json:"scheduleForm"`
	ScheduleMessage        string `json:"scheduleMessage"`
	SchedulePendingMessage string `json:"schedulePendingMessage"`
	ScheduleStart          string `json:"scheduleStart"`
	ScheduleStartAmpm      string `json:"scheduleStartAmpm"`
	ScheduleStartHour      string `json:"scheduleStartHour"`
	ScheduleStartMinute    string `json:"scheduleStartMinute"`
	SubLabelPlacement      string `json:"subLabelPlacement"`
	Title                  string `json:"title"`
	UseCurrentUserAsAuthor bool   `json:"useCurrentUserAsAuthor"`
	ValidationSummary      string `json:"validationSummary"`
	Version                string `json:"version"`
}