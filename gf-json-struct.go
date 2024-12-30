type AutoGenerated struct {
	AutoResponder struct {
		From    string `json:"from"`
		Message string `json:"message"`
		Subject string `json:"subject"`
		ToField string `json:"toField"`
	} `json:"autoResponder"`
	Button struct {
		ConditionalLogic     any    `json:"conditionalLogic"`
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
	DeleteEntry             string `json:"delete_entry"`
	DeleteEntryCondition    string `json:"delete_entry_condition"`
	DeleteEntryPeriod       string `json:"delete_entry_period"`
	DeleteEntryUnits        string `json:"delete_entry_units"`
	Deprecated              string `json:"deprecated"`
	Description             string `json:"description"`
	DescriptionPlacement    string `json:"descriptionPlacement"`
	EnableAnimation         bool   `json:"enableAnimation"`
	EnableHoneypot          bool   `json:"enableHoneypot"`
	Feeds                   struct {
		Gravityformsadvancedpostcreation []any `json:"gravityformsadvancedpostcreation"`
	} `json:"feeds"`
	Fields []struct {
		AdminLabel               string   `json:"adminLabel"`
		AllowsPrepopulate        bool     `json:"allowsPrepopulate"`
		AutocompleteAttribute    string   `json:"autocompleteAttribute"`
		CalculationFormula       string   `json:"calculationFormula"`
		CalculationRounding      string   `json:"calculationRounding"`
		Choices                  string   `json:"choices"`
		ConditionalLogic         string   `json:"conditionalLogic"`
		CSSClass                 string   `json:"cssClass"`
		DefaultValue             string   `json:"defaultValue"`
		Description              string   `json:"description"`
		DescriptionPlacement     string   `json:"descriptionPlacement"`
		DisableQuantity          bool     `json:"disableQuantity"`
		DisplayAllCategories     bool     `json:"displayAllCategories"`
		DisplayOnly              string   `json:"displayOnly"`
		EnableAutocomplete       bool     `json:"enableAutocomplete"`
		EnableCalculation        string   `json:"enableCalculation"`
		EnableEnhancedUI         int      `json:"enableEnhancedUI,omitempty"`
		EnablePasswordInput      string   `json:"enablePasswordInput,omitempty"`
		ErrorMessage             string   `json:"errorMessage"`
		Errors                   []any    `json:"errors,omitempty"`
		Fields                   string   `json:"fields"`
		FormID                   int      `json:"formId"`
		GkCustomMapIconEnabled   bool     `json:"gk_custom_map_icon_enabled"`
		GpDateTimeCalculatorUnit bool     `json:"gp-date-time-calculator_unit"`
		GwwordcountMaxWordCount  string   `json:"gwwordcount_max_word_count"`
		GwwordcountMinWordCount  string   `json:"gwwordcount_min_word_count"`
		ID                       int      `json:"id"`
		InputMask                bool     `json:"inputMask"`
		InputMaskIsCustom        bool     `json:"inputMaskIsCustom"`
		InputMaskValue           string   `json:"inputMaskValue"`
		InputName                string   `json:"inputName"`
		InputType                string   `json:"inputType"`
		Inputs                   any      `json:"inputs"`
		IsRequired               bool     `json:"isRequired"`
		Label                    string   `json:"label"`
		LabelPlacement           string   `json:"labelPlacement"`
		LayoutGridColumnSpan     int      `json:"layoutGridColumnSpan,omitempty"`
		LayoutGroupID            string   `json:"layoutGroupId"`
		MaxFiles                 string   `json:"maxFiles"`
		MaxLength                string   `json:"maxLength"`
		MultipleFiles            bool     `json:"multipleFiles"`
		NoDuplicates             bool     `json:"noDuplicates"`
		PageNumber               int      `json:"pageNumber"`
		Placeholder              string   `json:"placeholder"`
		ProductField             string   `json:"productField,omitempty"`
		Size                     string   `json:"size"`
		SubLabelPlacement        string   `json:"subLabelPlacement"`
		Type                     string   `json:"type"`
		UseRichTextEditor        bool     `json:"useRichTextEditor"`
		Visibility               string   `json:"visibility"`
		GwreadonlyEnable         bool     `json:"gwreadonly_enable,omitempty"`
		EmailConfirmEnabled      bool     `json:"emailConfirmEnabled,omitempty"`
		GpNestedFormsFields      string   `json:"gp-nested-forms_fields,omitempty"`
		GpNestedFormsForm        string   `json:"gp-nested-forms_form,omitempty"`
		GpnfEntryLabelPlural     string   `json:"gpnfEntryLabelPlural,omitempty"`
		GpnfEntryLabelSingular   string   `json:"gpnfEntryLabelSingular,omitempty"`
		GpnfEntryLimitMin        string   `json:"gpnfEntryLimitMin,omitempty"`
		GpnfFields               []string `json:"gpnfFields,omitempty"`
		GpnfForm                 string   `json:"gpnfForm,omitempty"`
		GpnfFormTitle            string   `json:"gpnfFormTitle,omitempty"`
		GppaChoicesFilterGroups  []any    `json:"gppa-choices-filter-groups,omitempty"`
		GppaChoicesTemplates     []any    `json:"gppa-choices-templates,omitempty"`
		GppaValuesFilterGroups   []any    `json:"gppa-values-filter-groups,omitempty"`
		GppaValuesObjectType     string   `json:"gppa-values-object-type,omitempty"`
		GppaValuesTemplates      []any    `json:"gppa-values-templates,omitempty"`
		CheckboxLabel            string   `json:"checkboxLabel,omitempty"`
		Content                  string   `json:"content,omitempty"`
		DisableMargins           string   `json:"disableMargins,omitempty"`
	} `json:"fields"`
	FirstPageCSSClass     any    `json:"firstPageCssClass"`
	FormRevisionsEnabled  string `json:"formRevisionsEnabled"`
	GfpdfFormSettings     []any  `json:"gfpdf_form_settings"`
	Gravityformsrecaptcha struct {
		DisableRecaptchav3 string `json:"disable-recaptchav3"`
	} `json:"gravityformsrecaptcha"`
	GravityviewEntryModeration string `json:"gravityview_entry_moderation"`
	GvInlineEditEnable         string `json:"gv_inline_edit_enable"`
	HoneypotAction             string `json:"honeypotAction"`
	ID                         int    `json:"id"`
	IsActive                   string `json:"is_active"`
	IsTrash                    string `json:"is_trash"`
	LabelPlacement             string `json:"labelPlacement"`
	LastPageButton             any    `json:"lastPageButton"`
	LimitEntries               bool   `json:"limitEntries"`
	LimitEntriesCount          string `json:"limitEntriesCount"`
	LimitEntriesMessage        string `json:"limitEntriesMessage"`
	LimitEntriesPeriod         string `json:"limitEntriesPeriod"`
	MarkupVersion              int    `json:"markupVersion"`
	MaxEntriesAllowed          string `json:"maxEntriesAllowed"`
	NextFieldID                int    `json:"nextFieldId"`
	Pagination                 any    `json:"pagination"`
	PostAuthor                 string `json:"postAuthor"`
	PostCategory               string `json:"postCategory"`
	PostContentTemplate        string `json:"postContentTemplate"`
	PostContentTemplateEnabled bool   `json:"postContentTemplateEnabled"`
	PostStatus                 string `json:"postStatus"`
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
	TemplateID             string `json:"template_id"`
	Title                  string `json:"title"`
	UseCurrentUserAsAuthor string `json:"useCurrentUserAsAuthor"`
	ValidationPlacement    string `json:"validationPlacement"`
	ValidationSummary      string `json:"validationSummary"`
	Version                string `json:"version"`
}