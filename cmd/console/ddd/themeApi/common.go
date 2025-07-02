package themeApi

const defaultTheme = "default_theme"

type themeConfig struct {
	Animation                    string `json:"animation"`
	ColumnsAsideLayout           string `json:"columnsAsideLayout"`
	ColumnsAsideStyle            string `json:"columnsAsideStyle"`
	ColumnsMenuBar               string `json:"columnsMenuBar"`
	ColumnsMenuBarColor          string `json:"columnsMenuBarColor"`
	GlobalComponentSize          string `json:"globalComponentSize"`
	GlobalI18N                   string `json:"globalI18n"`
	GlobalTitle                  string `json:"globalTitle"`
	GlobalViceTitle              string `json:"globalViceTitle"`
	GlobalViceTitleMsg           string `json:"globalViceTitleMsg"`
	IsBreadcrumb                 bool   `json:"isBreadcrumb"`
	IsBreadcrumbIcon             bool   `json:"isBreadcrumbIcon"`
	IsCacheTagsView              bool   `json:"isCacheTagsView"`
	IsClassicSplitMenu           bool   `json:"isClassicSplitMenu"`
	IsCollapse                   bool   `json:"isCollapse"`
	IsColumnsMenuBarColorGradual bool   `json:"isColumnsMenuBarColorGradual"`
	IsColumnsMenuHoverPreload    bool   `json:"isColumnsMenuHoverPreload"`
	IsDrawer                     bool   `json:"isDrawer"`
	IsFixedHeader                bool   `json:"isFixedHeader"`
	IsFixedHeaderChange          bool   `json:"isFixedHeaderChange"`
	IsFooter                     bool   `json:"isFooter"`
	IsGrayscale                  bool   `json:"isGrayscale"`
	IsInvert                     bool   `json:"isInvert"`
	IsIsDark                     bool   `json:"isIsDark"`
	IsLockScreen                 bool   `json:"isLockScreen"`
	IsMenuBarColorGradual        bool   `json:"isMenuBarColorGradual"`
	IsRequestRoutes              bool   `json:"isRequestRoutes"`
	IsShareTagsView              bool   `json:"isShareTagsView"`
	IsShowLogo                   bool   `json:"isShowLogo"`
	IsShowLogoChange             bool   `json:"isShowLogoChange"`
	IsSortableTagsView           bool   `json:"isSortableTagsView"`
	IsTagsview                   bool   `json:"isTagsview"`
	IsTagsviewIcon               bool   `json:"isTagsviewIcon"`
	IsTopBarColorGradual         bool   `json:"isTopBarColorGradual"`
	IsUniqueOpened               bool   `json:"isUniqueOpened"`
	IsWartermark                 bool   `json:"isWartermark"`
	Layout                       string `json:"layout"`
	LockScreenTime               int    `json:"lockScreenTime"`
	MenuBar                      string `json:"menuBar"`
	MenuBarActiveColor           string `json:"menuBarActiveColor"`
	MenuBarColor                 string `json:"menuBarColor"`
	Primary                      string `json:"primary"`
	TagsStyle                    string `json:"tagsStyle"`
	TopBar                       string `json:"topBar"`
	TopBarColor                  string `json:"topBarColor"`
	WartermarkText               string `json:"wartermarkText"`
}
