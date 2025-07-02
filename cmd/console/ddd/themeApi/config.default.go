package themeApi

var defaultConfig = themeConfig{
	// 是否开启布局配置抽屉
	IsDrawer: false,

	/**
	 * 全局主题
	 */
	// 默认 primary 主题颜色
	Primary: "#409eff",
	// 是否开启深色模式
	IsIsDark: false,

	/**
	 * 顶栏设置
	 */
	// 默认顶栏导航背景颜色
	TopBar: "#ffffff",
	// 默认顶栏导航字体颜色
	TopBarColor: "#606266",
	// 是否开启顶栏背景颜色渐变
	IsTopBarColorGradual: false,

	/**
	 * 菜单设置
	 */
	// 默认菜单导航背景颜色
	MenuBar: "#545c64",
	// 默认菜单导航字体颜色
	MenuBarColor: "#eaeaea",
	// 默认菜单高亮背景色
	MenuBarActiveColor: "rgba(0, 0, 0, 0.2)",
	// 是否开启菜单背景颜色渐变
	IsMenuBarColorGradual: false,

	/**
	 * 分栏设置
	 */
	// 默认分栏菜单背景颜色
	ColumnsMenuBar: "#545c64",
	// 默认分栏菜单字体颜色
	ColumnsMenuBarColor: "#e6e6e6",
	// 是否开启分栏菜单背景颜色渐变
	IsColumnsMenuBarColorGradual: false,
	// 是否开启分栏菜单鼠标悬停预加载(预览菜单)
	IsColumnsMenuHoverPreload: false,

	/**
	 * 界面设置
	 */
	// 是否开启菜单水平折叠效果
	IsCollapse: false,
	// 是否开启菜单手风琴效果
	IsUniqueOpened: true,
	// 是否开启固定 Header
	IsFixedHeader: false,
	// 初始化变量，用于更新菜单 el-scrollbar 的高度，请勿删除
	IsFixedHeaderChange: false,
	// 是否开启经典布局分割菜单（仅经典布局生效）
	IsClassicSplitMenu: false,
	// 是否开启自动锁屏
	IsLockScreen: false,
	// 开启自动锁屏倒计时(s/秒)
	LockScreenTime: 30,

	/**
	 * 界面显示
	 */
	// 是否开启侧边栏 Logo
	IsShowLogo: false,
	// 初始化变量，用于 el-scrollbar 的高度更新，请勿删除
	IsShowLogoChange: false,
	// 是否开启 Breadcrumb，强制经典、横向布局不显示
	IsBreadcrumb: true,
	// 是否开启 Tagsview
	IsTagsview: true,
	// 是否开启 Breadcrumb 图标
	IsBreadcrumbIcon: false,
	// 是否开启 Tagsview 图标
	IsTagsviewIcon: false,
	// 是否开启 TagsView 缓存
	IsCacheTagsView: false,
	// 是否开启 TagsView 拖拽
	IsSortableTagsView: true,
	// 是否开启 TagsView 共用
	IsShareTagsView: false,
	// 是否开启 Footer 底部版权信息
	IsFooter: false,
	// 是否开启灰色模式
	IsGrayscale: false,
	// 是否开启色弱模式
	IsInvert: false,
	// 是否开启水印
	IsWartermark: false,
	// 水印文案
	WartermarkText: "owl-messager",

	/**
	 * 其它设置
	 */
	// Tagsview 风格：可选值"<tags-style-one|tags-style-four|tags-style-five>"，默认 tags-style-five
	// 定义的值与 `/src/layout/navBars/tagsView/tagsView.vue` 中的 class 同名
	TagsStyle: "tags-style-five",
	// 主页面切换动画：可选值"<slide-right|slide-left|opacitys>"，默认 slide-right
	Animation: "slide-right",
	// 分栏高亮风格：可选值"<columns-round|columns-card>"，默认 columns-round
	ColumnsAsideStyle: "columns-round",
	// 分栏布局风格：可选值"<columns-horizontal|columns-vertical>"，默认 columns-horizontal
	ColumnsAsideLayout: "columns-vertical",

	/**
	 * 布局切换
	 * 注意：为了演示，切换布局时，颜色会被还原成默认，代码位置：/@/layout/navBars/breadcrumb/setings.vue
	 * 中的 `initSetLayoutChange(设置布局切换，重置主题样式)` 方法
	 */
	// 布局切换：可选值"<defaults|classic|transverse|columns>"，默认 defaults
	Layout: "defaults",

	/**
	 * 后端控制路由
	 */
	// 是否开启后端控制路由
	IsRequestRoutes: false,

	/**
	 * 全局网站标题 / 副标题
	 */
	// 网站主标题（菜单导航、浏览器当前网页标题）
	GlobalTitle: "owl messager",
	// 网站副标题（登录页顶部文字）
	GlobalViceTitle: "OwlMessager",
	// 网站副标题（登录页顶部文字）
	GlobalViceTitleMsg: "专注消息推送",
	// 默认初始语言，可选值"<zh-cn|en|zh-tw>"，默认 zh-cn
	GlobalI18N: "zh-cn",
	// 默认全局组件大小，可选值"<large|"default"|small>"，默认 "large"
	GlobalComponentSize: "large",
}
