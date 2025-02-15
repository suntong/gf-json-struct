type AutoGenerated struct {
	AutoResponder struct {
		From    string `json:"from,omitempty"`
		Message string `json:"message,omitempty"`
		Subject string `json:"subject,omitempty"`
		ToField string `json:"toField,omitempty"`
	} `json:"autoResponder,omitempty"`
	Button struct {
		ConditionalLogic     any    `json:"conditionalLogic,omitempty"`
		ID                   string `json:"id,omitempty"`
		ImageURL             string `json:"imageUrl,omitempty"`
		LayoutGridColumnSpan int    `json:"layoutGridColumnSpan,omitempty"`
		Location             string `json:"location,omitempty"`
		Text                 string `json:"text,omitempty"`
		Type                 string `json:"type,omitempty"`
		Width                string `json:"width,omitempty"`
	} `json:"button,omitempty"`
	CSSClass                string `json:"cssClass,omitempty"`
	CustomRequiredIndicator string `json:"customRequiredIndicator,omitempty"`
	DateCreated             string `json:"date_created,omitempty"`
	DeleteEntry             string `json:"delete_entry,omitempty"`
	DeleteEntryCondition    string `json:"delete_entry_condition,omitempty"`
	DeleteEntryPeriod       string `json:"delete_entry_period,omitempty"`
	DeleteEntryUnits        string `json:"delete_entry_units,omitempty"`
	Deprecated              string `json:"deprecated,omitempty"`
	Description             string `json:"description,omitempty"`
	DescriptionPlacement    string `json:"descriptionPlacement,omitempty"`
	EnableAnimation         bool   `json:"enableAnimation,omitempty"`
	EnableHoneypot          bool   `json:"enableHoneypot,omitempty"`
	Feeds                   struct {
		Gravityformsadvancedpostcreation []any `json:"gravityformsadvancedpostcreation,omitempty"`
	} `json:"feeds,omitempty"`
	Fields []struct {
		AdminLabel               string   `json:"adminLabel,omitempty"`
		AllowsPrepopulate        bool     `json:"allowsPrepopulate,omitempty"`
		AutocompleteAttribute    string   `json:"autocompleteAttribute,omitempty"`
		CalculationFormula       string   `json:"calculationFormula,omitempty"`
		CalculationRounding      string   `json:"calculationRounding,omitempty"`
		Choices                  string   `json:"choices,omitempty"`
		ConditionalLogic         string   `json:"conditionalLogic,omitempty"`
		CSSClass                 string   `json:"cssClass,omitempty"`
		DefaultValue             string   `json:"defaultValue,omitempty"`
		Description              string   `json:"description,omitempty"`
		DescriptionPlacement     string   `json:"descriptionPlacement,omitempty"`
		DisableQuantity          bool     `json:"disableQuantity,omitempty"`
		DisplayAllCategories     bool     `json:"displayAllCategories,omitempty"`
		DisplayOnly              string   `json:"displayOnly,omitempty"`
		EnableAutocomplete       bool     `json:"enableAutocomplete,omitempty"`
		EnableCalculation        string   `json:"enableCalculation,omitempty"`
		EnableEnhancedUI         int      `json:"enableEnhancedUI,omitempty"`
		EnablePasswordInput      string   `json:"enablePasswordInput,omitempty"`
		ErrorMessage             string   `json:"errorMessage,omitempty"`
		Errors                   []any    `json:"errors,omitempty"`
		Fields                   string   `json:"fields,omitempty"`
		FormID                   int      `json:"formId,omitempty"`
		GkCustomMapIconEnabled   bool     `json:"gk_custom_map_icon_enabled,omitempty"`
		GpDateTimeCalculatorUnit bool     `json:"gp-date-time-calculator_unit,omitempty"`
		GwwordcountMaxWordCount  string   `json:"gwwordcount_max_word_count,omitempty"`
		GwwordcountMinWordCount  string   `json:"gwwordcount_min_word_count,omitempty"`
		ID                       int      `json:"id,omitempty"`
		InputMask                bool     `json:"inputMask,omitempty"`
		InputMaskIsCustom        bool     `json:"inputMaskIsCustom,omitempty"`
		InputMaskValue           string   `json:"inputMaskValue,omitempty"`
		InputName                string   `json:"inputName,omitempty"`
		InputType                string   `json:"inputType,omitempty"`
		Inputs                   any      `json:"inputs,omitempty"`
		IsRequired               bool     `json:"isRequired,omitempty"`
		Label                    string   `json:"label,omitempty"`
		LabelPlacement           string   `json:"labelPlacement,omitempty"`
		LayoutGridColumnSpan     int      `json:"layoutGridColumnSpan,omitempty"`
		LayoutGroupID            string   `json:"layoutGroupId,omitempty"`
		MaxFiles                 string   `json:"maxFiles,omitempty"`
		MaxLength                string   `json:"maxLength,omitempty"`
		MultipleFiles            bool     `json:"multipleFiles,omitempty"`
		NoDuplicates             bool     `json:"noDuplicates,omitempty"`
		PageNumber               int      `json:"pageNumber,omitempty"`
		Placeholder              string   `json:"placeholder,omitempty"`
		ProductField             string   `json:"productField,omitempty"`
		Size                     string   `json:"size,omitempty"`
		SubLabelPlacement        string   `json:"subLabelPlacement,omitempty"`
		Type                     string   `json:"type,omitempty"`
		UseRichTextEditor        bool     `json:"useRichTextEditor,omitempty"`
		Visibility               string   `json:"visibility,omitempty"`
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
	} `json:"fields,omitempty"`
	FirstPageCSSClass     any    `json:"firstPageCssClass,omitempty"`
	FormRevisionsEnabled  string `json:"formRevisionsEnabled,omitempty"`
	GfpdfFormSettings     []any  `json:"gfpdf_form_settings,omitempty"`
	Gravityformsrecaptcha struct {
		DisableRecaptchav3 string `json:"disable-recaptchav3,omitempty"`
	} `json:"gravityformsrecaptcha,omitempty"`
	GravityviewEntryModeration string `json:"gravityview_entry_moderation,omitempty"`
	GvInlineEditEnable         string `json:"gv_inline_edit_enable,omitempty"`
	HoneypotAction             string `json:"honeypotAction,omitempty"`
	ID                         int    `json:"id,omitempty"`
	IsActive                   string `json:"is_active,omitempty"`
	IsTrash                    string `json:"is_trash,omitempty"`
	LabelPlacement             string `json:"labelPlacement,omitempty"`
	LastPageButton             any    `json:"lastPageButton,omitempty"`
	LimitEntries               bool   `json:"limitEntries,omitempty"`
	LimitEntriesCount          string `json:"limitEntriesCount,omitempty"`
	LimitEntriesMessage        string `json:"limitEntriesMessage,omitempty"`
	LimitEntriesPeriod         string `json:"limitEntriesPeriod,omitempty"`
	MarkupVersion              int    `json:"markupVersion,omitempty"`
	MaxEntriesAllowed          string `json:"maxEntriesAllowed,omitempty"`
	NextFieldID                int    `json:"nextFieldId,omitempty"`
	Pagination                 any    `json:"pagination,omitempty"`
	PostAuthor                 string `json:"postAuthor,omitempty"`
	PostCategory               string `json:"postCategory,omitempty"`
	PostContentTemplate        string `json:"postContentTemplate,omitempty"`
	PostContentTemplateEnabled bool   `json:"postContentTemplateEnabled,omitempty"`
	PostStatus                 string `json:"postStatus,omitempty"`
	PostTitleTemplate          string `json:"postTitleTemplate,omitempty"`
	PostTitleTemplateEnabled   bool   `json:"postTitleTemplateEnabled,omitempty"`
	RequireLogin               bool   `json:"requireLogin,omitempty"`
	RequireLoginMessage        string `json:"requireLoginMessage,omitempty"`
	RequiredIndicator          string `json:"requiredIndicator,omitempty"`
	Save                       struct {
		Button struct {
			Text string `json:"text,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"button,omitempty"`
		Enabled bool `json:"enabled,omitempty"`
	} `json:"save,omitempty"`
	SaveButtonText         string `json:"saveButtonText,omitempty"`
	SaveEnabled            string `json:"saveEnabled,omitempty"`
	ScheduleEnd            string `json:"scheduleEnd,omitempty"`
	ScheduleEndAmpm        string `json:"scheduleEndAmpm,omitempty"`
	ScheduleEndHour        string `json:"scheduleEndHour,omitempty"`
	ScheduleEndMinute      string `json:"scheduleEndMinute,omitempty"`
	ScheduleForm           bool   `json:"scheduleForm,omitempty"`
	ScheduleMessage        string `json:"scheduleMessage,omitempty"`
	SchedulePendingMessage string `json:"schedulePendingMessage,omitempty"`
	ScheduleStart          string `json:"scheduleStart,omitempty"`
	ScheduleStartAmpm      string `json:"scheduleStartAmpm,omitempty"`
	ScheduleStartHour      string `json:"scheduleStartHour,omitempty"`
	ScheduleStartMinute    string `json:"scheduleStartMinute,omitempty"`
	SubLabelPlacement      string `json:"subLabelPlacement,omitempty"`
	TemplateID             string `json:"template_id,omitempty"`
	Title                  string `json:"title,omitempty"`
	UseCurrentUserAsAuthor string `json:"useCurrentUserAsAuthor,omitempty"`
	ValidationPlacement    string `json:"validationPlacement,omitempty"`
	ValidationSummary      string `json:"validationSummary,omitempty"`
	Version                string `json:"version,omitempty"`
}