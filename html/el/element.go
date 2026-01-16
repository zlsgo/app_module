package el

/*
# Doctype 文档类型声明

在 HTML 中，doctype 是所有文档顶部必需的 <!doctype html> 前导声明。其唯一目的是防止浏览器在渲染文档时切换到所谓的"怪异模式"；也就是说，<!doctype html> doctype 确保浏览器尽最大努力遵循相关规范，而不是使用与某些规范不兼容的不同渲染模式。

doctype 不区分大小写。MDN 代码示例的惯例是使用小写，但将其写为 <!DOCTYPE html> 也很常见。

https://developer.mozilla.org/en-US/docs/Glossary/Doctype
*/
var Doctype = VoidEl("!DOCTYPE", Attr("html", ""))

/*
# 锚点元素

<a> HTML 元素（或锚点元素）通过其 href 属性创建指向网页、文件、电子邮件地址、同一页面中的位置或 URL 可以寻址的任何其他内容的超链接。

每个 <a> 中的内容应指示链接的目标。如果存在 href 属性，则在聚焦于 <a> 元素时按 Enter 键将激活它。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a
*/
func A(items ...Item) *Element { return El("a", items...) }

/*
# 缩写元素

<abbr> HTML 元素表示缩写或首字母缩略词。

在包含缩写或首字母缩略词时，首次使用时应以纯文本形式提供术语的完整展开，并使用 <abbr> 标记缩写。这会告知用户缩写或首字母缩略词的含义。

可选的 title 属性可以在没有完整展开的情况下为缩写或首字母缩略词提供展开。这为用户代理提供了如何宣布/显示内容的提示，同时告知所有用户缩写的含义。如果存在，title 必须包含此完整描述，不能包含其他内容。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/abbr
*/
func ABBR(items ...Item) *Element { return El("abbr", items...) }

/*
# 联系地址元素

<address> HTML 元素表示所包含的 HTML 提供个人、人员或组织的联系信息。

<address> 元素内容提供的联系信息可以采用适合上下文的任何形式，并且可以包括所需的任何类型的联系信息，例如实际地址、URL、电子邮件地址、电话号码、社交媒体账号、地理坐标等。<address> 元素应包括联系信息所指的人员、人员或组织的名称。

<address> 可以在各种上下文中使用，例如在页面标题中提供企业的联系信息，或通过在 <article> 中包含 <address> 元素来指示文章的作者。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/address
*/
func ADDRESS(items ...Item) *Element { return El("address", items...) }

/*
# 图像映射区域元素

<area> HTML 元素定义图像映射内具有预定义可点击区域的区域。图像映射允许将图像上的几何区域与超文本链接关联。

此元素仅在 <map> 元素内使用。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/area
*/
func AREA(items ...Item) *Element { return VoidEl("area", items...) }

/*
# 文章内容元素

<article> HTML 元素表示文档、页面、应用程序或站点中的独立组成部分，旨在可独立分发或重用（例如，在联合发布中）。示例包括：论坛帖子、杂志或报纸文章、博客条目、产品卡片、用户提交的评论、交互式小部件或小工具，或任何其他独立的内容项。

给定文档可以包含多篇文章；例如，在显示读者滚动时逐个显示每篇文章文本的博客上，每个帖子都将包含在一个 <article> 元素中，可能在其中包含一个或多个 <section>。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/article
*/
func ARTICLE(items ...Item) *Element { return El("article", items...) }

/*
# 侧边栏元素

<aside> HTML 元素表示文档的一部分，其内容仅与文档的主要内容间接相关。侧边栏通常呈现为侧边栏或标注框。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/aside
*/
func ASIDE(items ...Item) *Element { return El("aside", items...) }

/*
# 嵌入音频元素

<audio> HTML 元素用于在文档中嵌入声音内容。它可以包含一个或多个音频源,使用 src 属性或 <source> 元素表示:浏览器将选择最合适的一个。它也可以作为流媒体的目标,使用 MediaStream (https://developer.mozilla.org/en-US/docs/Web/API/MediaStream)。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/audio
*/
func AUDIO(items ...Item) *Element { return El("audio", items...) }

/*
# 引起注意元素

<b> HTML 元素用于吸引读者对元素内容的注意,但这些内容并不具有特殊的重要性。该元素以前被称为粗体元素,大多数浏览器仍然以粗体显示文本。但是,你不应该使用 <b> 来设置文本样式或赋予重要性。如果你希望创建粗体文本,应该使用 CSS font-weight 属性。如果你希望表示某个元素具有特殊重要性,应该使用 <strong> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/b
*/
func B(items ...Item) *Element { return El("b", items...) }

/*
# 文档基础 URL 元素

<base> HTML 元素指定用于文档中所有相对 URL 的基础 URL。一个文档中只能有一个 <base> 元素。

文档使用的基础 URL 可以通过脚本使用 Node.baseURI 访问。如果文档没有 <base> 元素,则 baseURI 默认为 location.href。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base
*/
func BASE(items ...Item) *Element { return VoidEl("base", items...) }

/*
# 双向隔离元素

<bdi> HTML 元素告诉浏览器的双向算法将其包含的文本与周围文本隔离处理。当网站动态插入某些文本且不知道所插入文本的方向性时,它特别有用。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/bdi
*/
func BDI(items ...Item) *Element { return El("bdi", items...) }

/*
# 双向文本覆盖元素

<bdo> HTML 元素覆盖当前文本的方向性,使其中的文本以不同的方向呈现。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/bdo
*/
func BDO(items ...Item) *Element { return El("bdo", items...) }

/*
# 块引用元素

<blockquote> HTML 元素表示所包含的文本是一段扩展引用。通常,这在视觉上通过缩进呈现(有关如何更改它,请参阅注释)。引用来源的 URL 可以使用 cite 属性给出,而来源的文本表示可以使用 <cite> 元素给出。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote
*/
func BLOCKQUOTE(items ...Item) *Element { return El("blockquote", items...) }

/*
# 文档主体元素

<body> HTML 元素表示 HTML 文档的内容。一个文档中只能有一个 <body> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
*/
func BODY(items ...Item) *Element { return El("body", items...) }

/*
# 换行元素

<br> HTML 元素在文本中产生换行（回车）。它对于编写诗歌或地址等行的划分很重要的内容很有用。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/br
*/
func BR(items ...Item) *Element { return VoidEl("br", items...) }

/*
# 按钮元素

<button> HTML 元素是一个交互式元素，由用户通过鼠标、键盘、手指、语音命令或其他辅助技术激活。激活后，它会执行一个操作，例如提交表单或打开对话框。

默认情况下，HTML 按钮以类似于用户代理运行平台的样式呈现，但你可以使用 CSS 更改按钮的外观。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/button
*/
func BUTTON(items ...Item) *Element { return El("button", items...) }

/*
# 图形画布元素

使用 HTML <canvas> 元素配合 canvas 脚本 API 或 WebGL API 来绘制图形和动画。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/canvas
*/
func CANVAS(items ...Item) *Element { return El("canvas", items...) }

/*
# 表格标题元素

<caption> HTML 元素指定表格的标题（或标题），为表格提供可访问的名称或可访问的描述。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/caption
*/
func CAPTION(items ...Item) *Element { return El("caption", items...) }

/*
# 引用元素

<cite> HTML 元素用于标记创意作品的标题。引用可以根据与引用元数据相关的上下文适当约定采用缩写形式。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/cite
*/
func CITE(items ...Item) *Element { return El("cite", items...) }

/*
# 内联代码元素

<code> HTML 元素以一种旨在表明文本是计算机代码短片段的方式显示其内容。默认情况下，内容文本使用用户代理的默认等宽字体显示。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/code
*/
func CODE(items ...Item) *Element { return El("code", items...) }

/*
# 表格列元素

<col> HTML 元素定义由其父 <colgroup> 元素表示的列组中的一列或多列。<col> 元素仅在没有定义 span 属性的 <colgroup> 元素的子元素中有效。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/col
*/
func COL(items ...Item) *Element { return VoidEl("col", items...) }

/*
# 表格列组元素

<colgroup> HTML 元素定义表格中的一组列。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/colgroup
*/
func COLGROUP(items ...Item) *Element { return El("colgroup", items...) }

/*
# 数据元素

<data> HTML 元素将给定的内容与机器可读的翻译链接起来。如果内容与时间或日期相关，则必须使用 <time> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/data
*/
func DATAEL(items ...Item) *Element { return El("data", items...) }

/*
# HTML 数据列表元素

> 有限可用性

<datalist> HTML 元素包含一组 <option> 元素，这些元素表示在其他控件中可供选择的允许或推荐选项。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/datalist
*/
func DATALIST(items ...Item) *Element { return El("datalist", items...) }

/*
# 描述详情元素

<dd> HTML 元素为描述列表 (<dl>) 中前面的术语 (<dt>) 提供描述、定义或值。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dd
*/
func DD(items ...Item) *Element { return El("dd", items...) }

/*
# 已删除文本元素

<del> HTML 元素表示已从文档中删除的文本范围。例如，这可以在呈现"跟踪更改"或源代码差异信息时使用。<ins> 元素可用于相反的目的：指示已添加到文档的文本。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/del
*/
func DEL(items ...Item) *Element { return El("del", items...) }

/*
# 详情披露元素

<details> HTML 元素创建一个披露小部件，其中的信息仅在小部件切换到打开状态时可见。必须使用 <summary> 元素提供摘要或标签。

披露小部件通常在屏幕上使用一个旋转（或扭转）的小三角形来指示打开/关闭状态，三角形旁边有一个标签。<summary> 元素的内容用作披露小部件的标签。<details> 的内容为 <summary> 提供可访问的描述。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/details
*/
func DETAILS(items ...Item) *Element { return El("details", items...) }

/*
# 定义元素

<dfn> HTML 元素指示要定义的术语。<dfn> 元素应在完整的定义语句中使用，其中术语的完整定义可以是以下之一：
  - 祖先段落（文本块，有时由 <p> 元素标记）
  - <dt>/<dd> 配对
  - <dfn> 元素的最近节祖先

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dfn
*/
func DFN(items ...Item) *Element { return El("dfn", items...) }

/*
# 对话框元素

<dialog> HTML 元素表示模态或非模态对话框或其他交互式组件，例如可关闭的警报、检查器或子窗口。

HTML <dialog> 元素用于创建模态和非模态对话框。模态对话框会中断与页面其余部分的交互，使其处于惰性状态，而非模态对话框允许与页面其余部分交互。

应使用 JavaScript 显示 <dialog> 元素。使用 .showModal() 方法显示模态对话框，使用 .show() 方法显示非模态对话框。可以使用 .close() 方法关闭对话框，或在提交嵌套在 <dialog> 元素中的 <form> 时使用 dialog 方法。模态对话框也可以通过按 Esc 键关闭。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dialog
*/
func DIALOG(items ...Item) *Element { return El("dialog", items...) }

/*
# 内容分区元素

<div> HTML 元素是流内容的通用容器。在使用 CSS 以某种方式设置样式之前，它对内容或布局没有影响（例如，直接对其应用样式，或对其父元素应用某种布局模型，如 Flexbox）。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/div
*/
func DIV(items ...Item) *Element { return El("div", items...) }

/*
# 描述列表元素

<dl> HTML 元素表示描述列表。该元素包含一组术语（使用 <dt> 元素指定）和描述（由 <dd> 元素提供）的列表。此元素的常见用途是实现词汇表或显示元数据（键值对列表）。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dl
*/
func DL(items ...Item) *Element { return El("dl", items...) }

/*
# 描述术语元素

<dt> HTML 元素指定描述或定义列表中的术语，因此必须在 <dl> 元素内使用。它通常后跟一个 <dd> 元素；但是，连续的多个 <dt> 元素表示多个术语，这些术语都由紧接着的下一个 <dd> 元素定义。

后续的 <dd>（描述详情）元素提供与使用 <dt> 指定的术语相关的定义或其他相关文本。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/dt
*/
func DT(items ...Item) *Element { return El("dt", items...) }

/*
# 强调元素

<em> HTML 元素标记具有强调重音的文本。<em> 元素可以嵌套，每个嵌套级别表示更大程度的强调。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/em
*/
func EM(items ...Item) *Element { return El("em", items...) }

/*
# 嵌入外部内容元素

<embed> HTML 元素在文档中的指定点嵌入外部内容。此内容由外部应用程序或其他交互式内容源（例如浏览器插件）提供。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/embed
*/
func EMBED(items ...Item) *Element { return VoidEl("embed", items...) }

/*
# 字段集元素

<fieldset> HTML 元素用于在 Web 表单中对多个控件以及标签 (<label>) 进行分组。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/fieldset
*/
func FIELDSET(items ...Item) *Element { return El("fieldset", items...) }

/*
# 图形标题元素

<figcaption> HTML 元素表示描述其父 <figure> 元素其余内容的标题或图例，为 <figure> 提供可访问的名称。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/figcaption
*/
func FIGCAPTION(items ...Item) *Element { return El("figcaption", items...) }

/*
# 带可选标题的图形元素

<figure> HTML 元素表示独立的内容，可能带有可选标题，使用 <figcaption> 元素指定。图形、其标题及其内容作为单个单元引用。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/figure
*/
func FIGURE(items ...Item) *Element { return El("figure", items...) }

/*
# 页脚元素

<footer> HTML 元素表示其最近的祖先分节内容或分节根元素的页脚。<footer> 通常包含有关该节作者的信息、版权数据或相关文档的链接。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/footer
*/
func FOOTER(items ...Item) *Element { return El("footer", items...) }

/*
# 表单元素

<form> HTML 元素表示包含用于提交信息的交互式控件的文档部分。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form
*/
func FORM(items ...Item) *Element { return El("form", items...) }

/*
# HTML 节标题元素

<h1> 到 <h6> HTML 元素表示六个级别的节标题。<h1> 是最高的节级别，<h6> 是最低的。默认情况下，所有标题元素在布局中创建块级框，从新行开始并占用其包含块中的全部可用宽度。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
*/
func H1(items ...Item) *Element { return El("h1", items...) }

/*
# HTML 节标题元素

<h1> 到 <h6> HTML 元素表示六个级别的节标题。<h1> 是最高的节级别，<h6> 是最低的。默认情况下，所有标题元素在布局中创建块级框，从新行开始并占用其包含块中的全部可用宽度。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
*/
func H2(items ...Item) *Element { return El("h2", items...) }

/*
# HTML 节标题元素

<h1> 到 <h6> HTML 元素表示六个级别的节标题。<h1> 是最高的节级别，<h6> 是最低的。默认情况下，所有标题元素在布局中创建块级框，从新行开始并占用其包含块中的全部可用宽度。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
*/
func H3(items ...Item) *Element { return El("h3", items...) }

/*
# HTML 节标题元素

<h1> 到 <h6> HTML 元素表示六个级别的节标题。<h1> 是最高的节级别，<h6> 是最低的。默认情况下，所有标题元素在布局中创建块级框，从新行开始并占用其包含块中的全部可用宽度。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
*/
func H4(items ...Item) *Element { return El("h4", items...) }

/*
# HTML 节标题元素

<h1> 到 <h6> HTML 元素表示六个级别的节标题。<h1> 是最高的节级别，<h6> 是最低的。默认情况下，所有标题元素在布局中创建块级框，从新行开始并占用其包含块中的全部可用宽度。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
*/
func H5(items ...Item) *Element { return El("h5", items...) }

/*
# HTML 节标题元素

<h1> 到 <h6> HTML 元素表示六个级别的节标题。<h1> 是最高的节级别，<h6> 是最低的。默认情况下，所有标题元素在布局中创建块级框，从新行开始并占用其包含块中的全部可用宽度。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
*/
func H6(items ...Item) *Element { return El("h6", items...) }

/*
# 文档元数据（头部）元素

<head> HTML 元素包含有关文档的机器可读信息（元数据），例如其标题、脚本和样式表。HTML 文档中只能有一个 <head> 元素。

	> <head> 主要保存用于机器处理的信息，而不是人类可读性。对于人类可见的信息，例如顶级标题和列出的作者，请参阅 <header> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
*/
func HEAD(items ...Item) *Element { return El("head", items...) }

/*
# 页眉元素

<header> HTML 元素表示介绍性内容，通常是一组介绍性或导航辅助工具。它可能包含一些标题元素，但也可能包含徽标、搜索表单、作者姓名和其他元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/header
*/
func HEADER(items ...Item) *Element { return El("header", items...) }

/*
# 标题组元素

<hgroup> HTML 元素表示标题和相关内容。它将单个 <h1>–<h6> 元素与一个或多个 <p> 分组。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hgroup
*/
func HGROUP(items ...Item) *Element { return El("hgroup", items...) }

/*
# 主题分隔（水平线）元素

<hr> HTML 元素表示段落级元素之间的主题分隔：例如，故事中的场景变化，或节内的主题转换。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hr
*/
func HR(items ...Item) *Element { return VoidEl("hr", items...) }

/*
# HTML 文档/根元素

<html> HTML 元素表示 HTML 文档的根（顶级元素），因此也称为根元素。所有其他元素都必须是此元素的后代。文档中只能有一个 <html> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html
*/
func HTML(items ...Item) *Element { return El("html", items...) }

/*
# 惯用文本元素

<i> HTML 元素表示由于某种原因与正常文本区分开的文本范围，例如惯用文本、技术术语、分类名称等。从历史上看，这些文本使用斜体呈现，这是此元素 <i> 命名的原始来源。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/i
*/
func I(items ...Item) *Element { return El("i", items...) }

/*
# 内联框架元素

<iframe> HTML 元素表示嵌套的浏览上下文，将另一个 HTML 页面嵌入到当前页面中。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/iframe
*/
func IFRAME(items ...Item) *Element { return El("iframe", items...) }

/*
# 图像嵌入元素

<img> HTML 元素将图像嵌入到文档中。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img
*/
func IMG(items ...Item) *Element { return VoidEl("img", items...) }

/*
# HTML 输入元素

<input> HTML 元素用于为基于 Web 的表单创建交互式控件，以便接受来自用户的数据；根据设备和用户代理，可以使用各种类型的输入数据和控件小部件。由于输入类型和属性的组合数量众多，<input> 元素是所有 HTML 中最强大和最复杂的元素之一。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input
*/
func INPUT(items ...Item) *Element { return VoidEl("input", items...) }

/*
# 插入文本元素

<ins> HTML 元素表示已添加到文档的文本范围。你可以使用 <del> 元素类似地表示已从文档中删除的文本范围。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ins
*/
func INS(items ...Item) *Element { return El("ins", items...) }

/*
# 键盘输入元素

<kbd> HTML 元素表示表示来自键盘、语音输入或任何其他文本输入设备的文本用户输入的内联文本范围。按照惯例，用户代理默认使用其默认等宽字体呈现 <kbd> 元素的内容，尽管 HTML 标准并未强制要求这样做。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/kbd
*/
func KBD(items ...Item) *Element { return El("kbd", items...) }

/*
# 标签元素

<label> HTML 元素表示用户界面中项目的标题。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
*/
func LABEL(items ...Item) *Element { return El("label", items...) }

/*
# 字段集图例元素

<legend> HTML 元素表示其父 <fieldset> 内容的标题。
在可自定义的 <select> 元素中，<legend> 元素允许作为 <optgroup> 的子元素，以提供易于定位和设置样式的标签。这将替换 <optgroup> 元素的 label 属性中设置的任何文本，并且具有相同的语义。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/legend
*/
func LEGEND(items ...Item) *Element { return El("legend", items...) }

/*
# 列表项元素

<li> HTML 元素用于表示列表中的项目。它必须包含在父元素中：有序列表 (<ol>)、无序列表 (<ul>) 或菜单 (<menu>)。在菜单和无序列表中，列表项通常使用项目符号显示。在有序列表中，它们通常在左侧显示升序计数器，例如数字或字母。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
*/
func LI(items ...Item) *Element { return El("li", items...) }

/*
# 外部资源链接元素

<link> HTML 元素指定当前文档与外部资源之间的关系。
此元素最常用于链接到样式表，但也用于建立站点图标（包括"favicon"样式图标以及移动设备上主屏幕和应用程序的图标）等。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
*/
func LINK(items ...Item) *Element { return VoidEl("link", items...) }

/*
# 主要内容元素

<main> HTML 元素表示文档 <body> 的主要内容。主要内容区域由与文档的中心主题直接相关或扩展的内容，或应用程序的中心功能组成。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/main
*/
func MAIN(items ...Item) *Element { return El("main", items...) }

/*
# 图像映射元素

<map> HTML 元素与 <area> 元素一起使用以定义图像映射（可点击的链接区域）。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/map
*/
func MAPEL(items ...Item) *Element { return El("map", items...) }

/*
# 标记文本元素

<mark> HTML 元素表示由于标记段落在封闭上下文中的相关性而为参考或注释目的标记或突出显示的文本。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/mark
*/
func MARK(items ...Item) *Element { return El("mark", items...) }

/*
# 菜单元素

<menu> HTML 元素在 HTML 规范中被描述为 <ul> 的语义替代品，但浏览器（并通过可访问性树公开）将其视为与 <ul> 没有区别。它表示项目的无序列表（由 <li> 元素表示）。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/menu
*/
func MENU(items ...Item) *Element { return El("menu", items...) }

/*
# 元数据元素

<meta> HTML 元素表示无法由其他元相关元素（例如 <base>、<link>、<script>、<style> 或 <title>）表示的元数据。
<meta> 元素提供的元数据类型可以是以下之一：

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta
*/
func META(items ...Item) *Element { return VoidEl("meta", items...) }

/*
# HTML 仪表元素

<meter> HTML 元素表示已知范围内的标量值或分数值。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meter
*/
func METER(items ...Item) *Element { return El("meter", items...) }

/*
# 导航节元素

<nav> HTML 元素表示页面的一个部分，其目的是提供导航链接，无论是在当前文档内还是到其他文档。导航节的常见示例包括菜单、目录和索引。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/nav
*/
func NAV(items ...Item) *Element { return El("nav", items...) }

/*
# 无脚本元素

<noscript> HTML 元素定义当页面上的脚本类型不受支持或浏览器当前关闭脚本时要插入的 HTML 部分。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/noscript
*/
func NOSCRIPT(items ...Item) *Element { return El("noscript", items...) }

/*
# 外部对象元素

<object> HTML 元素表示外部资源，可以将其视为图像、嵌套浏览上下文或由插件处理的资源。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/object
*/
func OBJECT(items ...Item) *Element { return El("object", items...) }

/*
# 有序列表元素

<ol> HTML 元素表示项目的有序列表 — 通常呈现为编号列表。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ol
*/
func OL(items ...Item) *Element { return El("ol", items...) }

/*
# 选项组元素

<optgroup> HTML 元素在 <select> 元素中创建选项分组。
在可自定义的 <select> 元素中，<legend> 元素允许作为 <optgroup> 的子元素，以提供易于定位和设置样式的标签。这将替换 <optgroup> 元素的 label 属性中设置的任何文本，并且具有相同的语义。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/optgroup
*/
func OPTGROUP(items ...Item) *Element { return El("optgroup", items...) }

/*
# HTML 选项元素

<option> HTML 元素用于定义包含在 <select>、<optgroup> 或 <datalist> 元素中的项目。因此，<option> 可以表示 HTML 文档中弹出窗口和其他项目列表中的菜单项。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/option
*/
func OPTION(items ...Item) *Element { return El("option", items...) }

/*
# 输出元素

<output> HTML 元素是一个容器元素，站点或应用程序可以将计算结果或用户操作的结果注入其中。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/output
*/
func OUTPUT(items ...Item) *Element { return El("output", items...) }

/*
# 段落元素

<p> HTML 元素表示段落。段落通常在视觉媒体中表示为由空行和/或首行缩进与相邻块分隔的文本块，但 HTML 段落可以是任何相关内容的结构分组，例如图像或表单字段。
段落是块级元素，值得注意的是，如果在结束 </p> 标记之前解析另一个块级元素，它们将自动关闭。请参阅下面的"标记省略"。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
*/
func P(items ...Item) *Element { return El("p", items...) }

/*
# 图片元素

<picture> HTML 元素包含零个或多个 <source> 元素和一个 <img> 元素，以为不同的显示/设备场景提供图像的替代版本。
浏览器将考虑每个子 <source> 元素并在其中选择最佳匹配。如果未找到匹配项 — 或浏览器不支持 <picture> 元素 — 则选择 <img> 元素的 src 属性的 URL。然后在 <img> 元素占用的空间中呈现所选图像。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/picture
*/
func PICTURE(items ...Item) *Element { return El("picture", items...) }

/*
# 预格式化文本元素

<pre> HTML 元素表示预格式化文本，将完全按照 HTML 文件中编写的方式呈现。文本通常使用非比例或等宽字体呈现。

此元素内的空白按编写的方式显示，但有一个例外。如果在开始 <pre> 标记之后立即包含一个或多个前导换行符，则第一个换行符将被删除。

<pre> 元素的文本内容被解析为 HTML，因此如果你想确保文本内容保持为纯文本，某些语法字符（例如 <）可能需要使用其各自的字符引用进行转义。有关更多信息，请参阅转义歧义字符。

<pre> 元素通常包含 <code>、<samp> 和 <kbd> 元素，分别表示计算机代码、计算机输出和用户输入。

默认情况下，<pre> 是块级元素，即其默认显示值为 block。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
*/
func PRE(items ...Item) *Element { return El("pre", items...) }

/*
# 进度指示器元素

<progress> HTML 元素显示指示任务完成进度的指示器，通常显示为进度条。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/progress
*/
func PROGRESS(items ...Item) *Element { return El("progress", items...) }

/*
# 内联引用元素

<q> HTML 元素表示所包含的文本是一个简短的内联引用。大多数现代浏览器通过在文本周围加上引号来实现这一点。此元素用于不需要段落分隔的简短引用；对于长引用，请使用 <blockquote> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/q
*/
func Q(items ...Item) *Element { return El("q", items...) }

/*
# Ruby 后备括号元素

<rp> HTML 元素用于为不支持使用 <ruby> 元素显示 ruby 注释的浏览器提供后备括号。一个 <rp> 元素应包含包裹包含注释文本的 <rt> 元素的开括号和闭括号中的每一个。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/rp
*/
func RP(items ...Item) *Element { return El("rp", items...) }

/*
# Ruby 文本元素

<rt> HTML 元素指定 ruby 注释的 ruby 文本组件，用于为东亚排版提供发音、翻译或音译信息。<rt> 元素必须始终包含在 <ruby> 元素中。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/rt
*/
func RT(items ...Item) *Element { return El("rt", items...) }

/*
# Ruby 注释元素

<ruby> HTML 元素表示在基本文本上方、下方或旁边呈现的小注释，通常用于显示东亚字符的发音。它也可以用于注释其他类型的文本，但这种用法不太常见。

术语 ruby 起源于排字工使用的测量单位，表示文本可以在新闻纸上打印的最小尺寸，同时保持清晰可读。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ruby
*/
func RUBY(items ...Item) *Element { return El("ruby", items...) }

/*
# 删除线元素

<s> HTML 元素使用删除线或贯穿线呈现文本。使用 <s> 元素表示不再相关或不再准确的内容。但是，<s> 不适合指示文档编辑；为此，请酌情使用 <del> 和 <ins> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/s
*/
func S(items ...Item) *Element { return El("s", items...) }

/*
# 示例输出元素

<samp> HTML 元素用于包含表示计算机程序的示例（或引用）输出的内联文本。其内容通常使用浏览器的默认等宽字体（例如 Courier 或 Lucida Console）呈现。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/samp
*/
func SAMP(items ...Item) *Element { return El("samp", items...) }

/*
# 脚本元素

<script> HTML 元素用于嵌入可执行代码或数据；这通常用于嵌入或引用 JavaScript 代码。<script> 元素也可以与其他语言一起使用，例如 WebGL 的 GLSL 着色器编程语言和 JSON。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
*/
func SCRIPT(items ...Item) *Element { return El("script", items...) }

/*
# 通用搜索元素

<search> HTML 元素是一个容器，表示文档或应用程序中与执行搜索或过滤操作相关的表单控件或其他内容的部分。<search> 元素在语义上将元素内容的目的标识为具有搜索或过滤功能。搜索或过滤功能可以用于网站或应用程序、当前网页或文档，或整个互联网或其子部分。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/search
*/
func SEARCH(items ...Item) *Element { return El("search", items...) }

/*
# 通用节元素

<section> HTML 元素表示文档的通用独立节，没有更具体的语义元素来表示它。节应始终有标题，极少数例外。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/section
*/
func SECTION(items ...Item) *Element { return El("section", items...) }

/*
# HTML 选择元素

<select> HTML 元素表示提供选项菜单的控件。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/select
*/
func SELECT(items ...Item) *Element { return El("select", items...) }

/*
# Web 组件插槽元素

<slot> HTML 元素 — Web 组件技术套件的一部分 — 是 Web 组件内的占位符，你可以用自己的标记填充它，这使你可以创建单独的 DOM 树并将它们一起呈现。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/slot
*/
func SLOT(items ...Item) *Element { return El("slot", items...) }

/*
# 旁注元素

<small> HTML 元素表示旁注和小字，例如版权和法律文本，独立于其样式呈现。默认情况下，它将其中的文本呈现为小一号字体，例如从 small 到 x-small。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/small
*/
func SMALL(items ...Item) *Element { return El("small", items...) }

/*
# 媒体或图像源元素

<source> HTML 元素为 <picture>、<audio> 和 <video> 元素指定一个或多个媒体资源。它是一个空元素，这意味着它没有内容并且不需要结束标记。此元素通常用于以多种文件格式提供相同的媒体内容，以便在浏览器对图像文件格式和媒体文件格式的支持不同的情况下提供兼容性。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/source
*/
func SOURCE(items ...Item) *Element { return VoidEl("source", items...) }

/*
# 内容跨度元素

<span> HTML 元素是用于短语内容的通用内联容器，它本身不代表任何内容。它可以用于出于样式目的对元素进行分组（使用 class 或 id 属性），或者因为它们共享属性值，例如 lang。仅当没有其他语义元素合适时才应使用它。<span> 非常类似于 <div> 元素，但 <div> 是块级元素，而 <span> 是内联级元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/span
*/
func SPAN(items ...Item) *Element { return El("span", items...) }

/*
# 强重要性元素

<strong> HTML 元素表示其内容具有强重要性、严肃性或紧迫性。浏览器通常以粗体呈现内容。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/strong
*/
func STRONG(items ...Item) *Element { return El("strong", items...) }

/*
# 样式信息元素

<style> HTML 元素包含文档或文档部分的样式信息。它包含 CSS，应用于包含 <style> 元素的文档内容。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
*/
func STYLE(items ...Item) *Element { return El("style", items...) }

/*
# 下标元素

<sub> HTML 元素指定纯粹出于排版原因应显示为下标的内联文本。下标通常使用较小的文本以降低的基线呈现。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/sub
*/
func SUB(items ...Item) *Element { return El("sub", items...) }

/*
# 披露摘要元素

<summary> HTML 元素为 <details> 元素的披露框指定摘要、标题或图例。单击 <summary> 元素会切换父 <details> 元素的打开和关闭状态。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/summary
*/
func SUMMARY(items ...Item) *Element { return El("summary", items...) }

/*
# 上标元素

<sup> HTML 元素指定纯粹出于排版原因应显示为上标的内联文本。上标通常使用较小的文本以升高的基线呈现。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/sup
*/
func SUP(items ...Item) *Element { return El("sup", items...) }

/*
# 表格元素

<table> HTML 元素表示表格数据 — 即以包含数据的单元格的行和列组成的二维表格呈现的信息。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/table
*/
func TABLE(items ...Item) *Element { return El("table", items...) }

/*
# 表格主体元素

<tbody> HTML 元素封装一组表格行（<tr> 元素），表示它们构成表格（主要）数据的主体。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tbody
*/
func TBODY(items ...Item) *Element { return El("tbody", items...) }

/*
# 表格数据单元格元素

<td> HTML 元素定义包含数据的表格单元格，可以用作 <tr> 元素的子元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/td
*/
func TD(items ...Item) *Element { return El("td", items...) }

/*
# 内容模板元素

<template> HTML 元素用作保存 HTML 片段的机制，这些片段可以稍后通过 JavaScript 使用或立即生成到 shadow DOM 中。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template
*/
func TEMPLATE(items ...Item) *Element { return El("template", items...) }

/*
# 文本区域元素

<textarea> HTML 元素表示多行纯文本编辑控件，当你想允许用户输入大量自由格式文本时很有用，例如对评论或反馈表单的评论。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/textarea
*/
func TEXTAREA(items ...Item) *Element { return El("textarea", items...) }

/*
# 表格页脚元素

<tfoot> HTML 元素封装一组表格行（<tr> 元素），表示它们构成表格的页脚，其中包含有关表格列的信息。这通常是列的摘要，例如，列中给定数字的总和。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tfoot
*/
func TFOOT(items ...Item) *Element { return El("tfoot", items...) }

/*
# 表格标题元素

<th> HTML 元素将单元格定义为一组表格单元格的标题，可以用作 <tr> 元素的子元素。此组的确切性质由 scope 和 headers 属性定义。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/th
*/
func TH(items ...Item) *Element { return El("th", items...) }

/*
# 表格头部元素

<thead> HTML 元素封装一组表格行（<tr> 元素），表示它们构成表格的头部，其中包含有关表格列的信息。这通常采用列标题（<th> 元素）的形式。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/thead
*/
func THEAD(items ...Item) *Element { return El("thead", items...) }

/*
# （日期）时间元素

<time> HTML 元素表示特定的时间段。它可以包含 datetime 属性以将日期转换为机器可读格式，从而获得更好的搜索引擎结果或自定义功能（例如提醒）。

它可以表示以下之一：
  - 24 小时制的时间。
  - 公历中的精确日期（带有可选的时间和时区信息）。
  - 有效的持续时间。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/time
*/
func TIME(items ...Item) *Element { return El("time", items...) }

/*
# 文档标题元素

<title> HTML 元素定义显示在浏览器标题栏或页面选项卡中的文档标题。它仅包含文本；元素内的 HTML 标记（如果有）也被视为纯文本。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/title
*/
func TITLE(items ...Item) *Element { return El("title", items...) }

/*
# 表格行元素

<tr> HTML 元素定义表格中的单元格行。然后可以使用 <td>（数据单元格）和 <th>（标题单元格）元素的混合来建立行的单元格。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tr
*/
func TR(items ...Item) *Element { return El("tr", items...) }

/*
# 嵌入文本轨道元素

<track> HTML 元素用作媒体元素 <audio> 和 <video> 的子元素。
每个 track 元素允许你指定可以与媒体元素并行显示的定时文本轨道（或基于时间的数据），例如在视频顶部叠加字幕或隐藏式字幕，或与音频轨道一起显示。

可以为媒体元素指定多个轨道，包含不同类型的定时文本数据，或已为不同语言环境翻译的定时文本数据。
使用的数据将是已设置为默认的轨道，或基于用户首选项的类型和翻译。

轨道以 WebVTT 格式（.vtt 文件）格式化 — Web 视频文本轨道。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/track
*/
func TRACK(items ...Item) *Element { return VoidEl("track", items...) }

/*
# 未明确注释（下划线）元素

<u> HTML 元素表示应以指示其具有非文本注释的方式呈现的内联文本范围。默认情况下，这呈现为单个实线下划线，但可以使用 CSS 更改。

> 此元素在旧版本的 HTML 中被称为"下划线"元素，有时仍以这种方式被误用。要为文本加下划线，你应该应用包含设置为 underline 的 CSS text-decoration 属性的样式。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/u
*/
func U(items ...Item) *Element { return El("u", items...) }

/*
# 无序列表元素

<ul> HTML 元素表示项目的无序列表，通常呈现为项目符号列表。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
*/
func UL(items ...Item) *Element { return El("ul", items...) }

/*
# 变量元素

<var> HTML 元素表示数学表达式或编程上下文中变量的名称。它通常使用当前字体的斜体版本呈现，尽管该行为取决于浏览器。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/var
*/
func VAR(items ...Item) *Element { return El("var", items...) }

/*
# 视频嵌入元素

<video> HTML 元素嵌入支持视频播放的媒体播放器到文档中。你也可以将 <video> 用于音频内容，但 <audio> 元素可能提供更合适的用户体验。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/video
*/
func VIDEO(items ...Item) *Element { return El("video", items...) }

/*
# 换行机会元素

<wbr> HTML 元素表示换行机会 — 文本中浏览器可以选择性地换行的位置，尽管其换行规则不会在该位置创建换行。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/wbr
*/
func WBR(items ...Item) *Element { return VoidEl("wbr", items...) }

// Attributes

func As(value string) *Attribute { return Attr("as", value) }

/*
# HTML 属性：accept

accept 属性的值是一个逗号分隔的列表，包含一个或多个文件类型或唯一文件类型说明符，描述允许哪些文件类型。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/accept
*/
func Accept(value string) *Attribute        { return Attr("accept", value) }
func AcceptCharset(value string) *Attribute { return Attr("accept-charset", value) }
func AccessKey(value string) *Attribute     { return Attr("accesskey", value) }
func Action(value string) *Attribute        { return Attr("action", value) }
func Align(value string) *Attribute         { return Attr("align", value) }

/*
# <img> alt 属性

定义可以在页面中替换图像的文本。

> 浏览器并不总是显示图像。浏览器可能不显示图像的情况有很多，例如：

  - 非视觉浏览器（例如视觉障碍人士使用的浏览器）
  - 用户选择不显示图像（节省带宽、隐私原因）
  - 图像无效或不支持的类型

在这些情况下，浏览器可能会用元素的 alt 属性中的文本替换图像。出于这些原因和其他原因，请尽可能为 alt 提供有用的值。

将此属性设置为空字符串 (alt="") 表示此图像不是内容的关键部分（它是装饰或跟踪像素），非视觉浏览器可能会在渲染时省略它。如果 alt 属性为空且图像无法显示，视觉浏览器也会隐藏损坏的图像图标。

此属性还用于将图像复制并粘贴到文本时，或将链接的图像保存到书签时。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img#alt
*/
func Alt(value string) *Attribute        { return Attr("alt", value) }
func Aria(name, value string) *Attribute { return Attr("aria-"+name, value) }
func Async(value string) *Attribute      { return Attr("async", value) }

/*
# HTML 属性：autocomplete

它可用于接受文本或数字值作为输入的 <input> 元素、<textarea> 元素、<select> 元素和 <form> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/autocomplete
*/
func Autocomplete(value string) *Attribute { return Attr("autocomplete", value) }

/*
# HTML 属性：autofocus
*/
func Autofocus(value string) *Attribute { return Attr("autofocus", value) }

/*
# HTML 属性：autoplay
*/
func Autoplay(value string) *Attribute { return Attr("autoplay", value) }

/*
# HTML 属性：charset

此属性声明文档的字符编码。如果存在该属性，其值必须是字符串 "utf-8" 的 ASCII 不区分大小写匹配，因为 UTF-8 是 HTML5 文档的唯一有效编码。声明字符编码的 <meta> 元素必须完全位于文档的前 1024 字节内。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#charset
*/
func Charset(value string) *Attribute { return Attr("charset", value) }

var Checked = Attr("checked", "")

/*
# <blockquote> cite 属性

指定引用信息的源文档或消息的 URL。此属性旨在指向解释引用的上下文或参考的信息。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote#cite
*/
func Cite(value string) *Attribute { return Attr("cite", value) }

/*
# HTML class 全局属性

class 全局属性是元素的类列表，由 ASCII 空格分隔。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/class
*/
func Class(class string) *Attribute   { return Attr("class", class) }
func Color(value string) *Attribute   { return Attr("color", value) }
func Cols(value string) *Attribute    { return Attr("cols", value) }
func Colspan(value string) *Attribute { return Attr("colspan", value) }

/*
# HTML 属性：content

content 属性指定由 <meta> name 属性定义的元数据名称的值。
它以字符串作为其值，预期的语法因使用的 name 值而异。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/content
*/
func Content(value string) *Attribute { return Attr("content", value) }

/*
# HTML contenteditable 全局属性

contenteditable 全局属性是一个枚举属性，指示元素是否应由用户编辑。如果是，浏览器会修改其小部件以允许编辑。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/contenteditable
*/
var ContentEditable = Attr("contenteditable", "")

func Controls(value string) *Attribute { return Attr("controls", value) }
func Coords(value string) *Attribute   { return Attr("coords", value) }

/*
# HTML 属性：crossorigin

crossorigin 属性对 <audio>、<img>、<link>、<script> 和 <video> 元素有效，提供对 CORS 的支持，定义元素如何处理跨源请求，从而启用元素获取数据的 CORS 请求配置。根据元素的不同，该属性可以是 CORS 设置属性。

媒体元素上的 crossorigin 内容属性是 CORS 设置属性。

这些属性是枚举的，具有以下可能的值：

请求使用 CORS 标头，凭据标志设置为 'same-origin'。除非目标是同源，否则不会通过 cookie、客户端 TLS 证书或 HTTP 身份验证交换用户凭据。

请求使用 CORS 标头，凭据标志设置为 'include'，并且始终包含用户凭据。

将属性名称设置为空值，如 crossorigin 或 crossorigin=""，与 anonymous 相同。

无效的关键字和空字符串将作为 anonymous 关键字处理。

默认情况下（即未指定属性时），根本不使用 CORS。用户代理不会请求对资源的完全访问权限，在跨源请求的情况下，将根据相关元素的类型应用某些限制：

	> 基于 Chromium 的浏览器不支持 rel="icon" 的 crossorigin 属性。请参阅未解决的 Chromium 问题。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/crossorigin
*/
func Crossorigin(value string) *Attribute { return Attr("crossorigin", value) }

/*
# HTML data-* 全局属性

data-* 全局属性形成一类称为自定义数据属性的属性，允许脚本在 HTML 及其 DOM 表示之间交换专有信息。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/data-*
*/
func Data[T AttrValue](name string, value T) *Attribute { return Attr("data-"+name, value) }
func Datetime(value string) *Attribute                  { return Attr("datetime", value) }
func Default(value string) *Attribute                   { return Attr("default", value) }

var Defer = Attr("defer", "")

/*
# HTML 属性：dirname

dirname 属性可用于 <textarea> 元素和多个 <input> 类型，并描述表单提交期间元素文本内容的方向性。浏览器使用此属性的值来确定用户输入的文本是从左到右还是从右到左定向的。使用时，元素的文本方向性值与 dirname 属性的值一起作为字段名称包含在表单提交数据中。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/dirname
*/
func Dirname(value string) *Attribute { return Attr("dirname", value) }

/*
# HTML 属性：disabled

布尔 disabled 属性存在时，使元素不可变、不可聚焦，甚至不能与表单一起提交。用户既不能编辑也不能聚焦于控件，也不能聚焦于其表单控件后代。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/disabled
*/
var Disabled = Attr("disabled", "")

func Download(value string) *Attribute { return Attr("download", value) }

/*
# HTML draggable 全局属性

draggable 全局属性是一个枚举属性，指示元素是否可以拖动，无论是使用原生浏览器行为还是 HTML 拖放 API。

draggable 属性可以应用于严格属于 HTML 命名空间的元素，这意味着它不能应用于 SVG。
有关命名空间声明的外观以及它们的作用的更多信息，请参阅命名空间速成课程。

draggable 可以具有以下值：

	> 此属性是枚举的而不是布尔的。true 或 false 的值是强制性的，并且禁止使用像 <img draggable> 这样的简写。正确的用法是 <img draggable="true">。

如果未设置此属性，其默认值为 auto，这意味着拖动行为是默认的浏览器行为：只能拖动文本选择、图像和链接。对于其他元素，必须设置事件 ondragstart 才能使拖放工作，如此综合示例所示。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/draggable
*/
func Draggable(value string) *Attribute    { return Attr("draggable", value) }
func Enctype(value string) *Attribute      { return Attr("enctype", value) }
func EnterKeyHint(value string) *Attribute { return Attr("enterkeyhint", value) }

/*
# HTML 属性：for

for 属性是 <label> 和 <output> 的允许属性。在 <label> 元素上使用时，它指示此标签描述的表单元素。在 <output> 元素上使用时，它允许在表示输出中使用的值的元素之间建立显式关系。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/for
*/
func For(value string) *Attribute { return Attr("for", value) }

/*
# HTML 属性：form

form HTML 属性将表单关联元素与同一文档中的 <form> 元素关联。此属性适用于 <button>、<fieldset>、<input>、<object>、<output>、<select> 和 <textarea> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/form
*/
func Form(value string) *Attribute { return Attr("form", value) }

/*
# <form> action 属性

处理表单提交的 URL。此值可以被 <button>、<input type="submit"> 或 <input type="image"> 元素上的 formaction 属性覆盖。当设置 method="dialog" 时，此属性将被忽略。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form#action
*/
func FormAction(value string) *Attribute { return Attr("formaction", value) }
func Headers(value string) *Attribute    { return Attr("headers", value) }

/*
# SVG 属性：height

height 属性定义用户坐标系统中元素的垂直长度。

https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Attribute/height
*/
func Height(value string) *Attribute { return Attr("height", value) }

/*
# HTML hidden 全局属性

hidden 全局属性是一个枚举属性，指示浏览器不应呈现元素的内容。例如，它可用于隐藏在登录过程完成之前无法使用的页面元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/hidden
*/
func Hidden(value string) *Attribute { return Attr("hidden", value) }
func High(value string) *Attribute   { return Attr("high", value) }

/*
# <a> href 属性

超链接指向的 URL。链接不限于基于 HTTP 的 URL — 它们可以使用浏览器支持的任何 URL 方案：

  - 使用 tel: URL 的电话号码

  - 使用 mailto: URL 的电子邮件地址

  - 使用 sms: URL 的短信文本消息

  - 使用 javascript: URL 的可执行代码

  - 虽然 Web 浏览器可能不支持其他 URL 方案，但网站可以使用 registerProtocolHandler()

此外，其他 URL 功能可以定位资源的特定部分，包括：

  - 使用文档片段的页面部分

  - 使用文本片段的特定文本部分

  - 使用媒体片段的媒体文件片段

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#href
*/
func Href(value string) *Attribute { return Attr("href", value) }

/*
# <a> hreflang 属性

提示链接 URL 的人类语言。没有内置功能。允许的值与全局 lang 属性相同。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#hreflang
*/
func HrefLang(value string) *Attribute { return Attr("hreflang", value) }

/*
# <meta> http-equiv 属性

<meta> 元素的 http-equiv 属性允许你为浏览器提供处理指令，就好像返回文档的响应包含某些 HTTP 标头一样。
元数据是适用于整个页面的文档级元数据。

当 <meta> 元素具有 http-equiv 属性时，content 属性定义相应的 http-equiv 值。
例如，以下 <meta> 标记告诉浏览器在 5 分钟后刷新页面：

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta/http-equiv
*/
func HttpEquiv(value string) *Attribute { return Attr("http-equiv", value) }

/*
# HTML id 全局属性

id 全局属性定义一个标识符 (ID)，该标识符在整个文档中必须是唯一的。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/id
*/
func ID(value string) *Attribute        { return Attr("id", value) }
func Inert(value string) *Attribute     { return Attr("inert", value) }
func InputMode(value string) *Attribute { return Attr("inputmode", value) }
func IsMap(value string) *Attribute     { return Attr("ismap", value) }
func Kind(value string) *Attribute      { return Attr("kind", value) }
func Label(value string) *Attribute     { return Attr("label", value) }
func Src(value string) *Attribute       { return Attr("src", value) }
func Role(value string) *Attribute      { return Attr("role", value) }

/*
# HTML lang 全局属性

lang 全局属性帮助定义元素的语言：不可编辑元素所用的语言，或可编辑元素应由用户编写的语言。该属性包含单个 BCP 47 语言标记。

	> lang 的默认值是空字符串，这意味着语言未知。因此，建议始终为此属性指定适当的值。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/lang
*/
func Lang(value string) *Attribute { return Attr("lang", value) }
func List(value string) *Attribute { return Attr("list", value) }
func Loop(value string) *Attribute { return Attr("loop", value) }
func Low(value string) *Attribute  { return Attr("low", value) }

/*
# HTML 属性：max

max 属性定义包含该属性的输入可接受和有效的最大值。如果元素的值大于此值，则元素验证失败。此值必须大于或等于 min 属性的值。如果存在 max 属性但未指定或无效，则不应用最大值。如果 max 属性有效且非空值大于 max 属性允许的最大值，则约束验证将阻止表单提交。

max 属性对数字输入类型有效，包括 date、month、week、time、datetime-local、number 和 range 类型，以及 <progress> 和 <meter> 元素。它是一个数字，指定表单控件被视为有效的最大正值。

如果值超过允许的最大值，validityState.rangeOverflow 将为 true，并且控件将与 :out-of-range 和 :invalid 伪类匹配。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/max
*/
func Max(value string) *Attribute { return Attr("max", value) }

/*
# HTML 属性：maxlength

maxlength 属性定义用户可以输入到 <input> 或 <textarea> 中的最大字符串长度。该属性必须具有 0 或更高的整数值。

长度以 UTF-16 代码单元测量，通常但并非总是等于字符数。如果未指定 maxlength 或指定了无效值，则输入没有最大长度。

任何 maxlength 值必须大于或等于 minlength 的值（如果存在且有效）。如果字段的文本值的长度大于 maxlength UTF-16 代码单元长，则输入将无法通过约束验证。约束验证仅在用户更改值时应用。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/maxlength
*/
func MaxLength(value string) *Attribute { return Attr("maxlength", value) }

/*
# HTML 属性：media

media 属性定义应将 content 属性中定义的主题颜色应用于哪些媒体。其值是媒体查询，如果缺少该属性，则默认为 all。此属性仅在元素的 name 属性设置为 theme-color 时才相关。否则，它没有效果，不应包含。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#media
*/
func Media(value string) *Attribute { return Attr("media", value) }

/*
# HTML 属性：method

提交表单的 HTTP 方法。唯一允许的方法/值是（不区分大小写）：

post：POST 方法；表单数据作为请求正文发送。
get（默认）：GET；表单数据附加到带有 ? 分隔符的 action URL。当表单没有副作用时使用此方法。
dialog：当表单位于 <dialog> 内时，关闭对话框并在提交时触发 submit 事件，而不提交数据或清除表单。
此值被 <button>、<input type="submit"> 或 <input type="image"> 元素上的 formmethod 属性覆盖。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form#method
*/
func Method(value string) *Attribute { return Attr("method", value) }

/*
# HTML 属性：min

min 属性定义包含该属性的输入可接受和有效的最小值。如果元素的值小于此值，则元素验证失败。此值必须小于或等于 max 属性的值。

某些输入类型具有默认最小值。如果输入没有默认最小值，并且为 min 指定的值无法转换为有效数字（或未设置最小值），则输入没有最小值。

它对输入类型有效，包括：date、month、week、time、datetime-local、number 和 range 类型，以及 <meter> 元素。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/min
*/
func Min(value string) *Attribute { return Attr("min", value) }

/*
# HTML 属性：minlength

minlength 属性定义用户可以输入到 <input> 或 <textarea> 中的最小字符串长度。该属性必须具有 0 或更高的整数值。

长度以 UTF-16 代码单元测量，通常但并非总是等于字符数。如果未指定 minlength 或指定了无效值，则输入没有最小长度。此值必须小于或等于 maxlength 的值，否则该值将永远无效，因为不可能同时满足两个条件。

如果字段的文本值的长度小于 minlength UTF-16 代码单元长，则输入将无法通过约束验证，validityState.tooShort 返回 true。约束验证仅在用户更改值时应用。一旦提交失败，某些浏览器将显示错误消息，指示所需的最小长度和当前长度。

minlength 不意味着 required；仅当用户输入了值时，输入才违反 minlength 约束。如果输入不是必需的，即使设置了 minlength，也可以提交空字符串。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/minlength
*/
func MinLength(value string) *Attribute { return Attr("minlength", value) }

/*
# HTML 属性：multiple

布尔 multiple 属性（如果设置）意味着表单控件接受一个或多个值。该属性对 email 和 file 输入类型以及 <select> 有效。用户选择多个值的方式取决于表单控件。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/multiple
*/
func Multiple(value string) *Attribute { return Attr("multiple", value) }

/*
# HTML 属性：muted
*/
func Muted(value string) *Attribute { return Attr("muted", value) }

/*
# HTML 属性：name

name 和 content 属性可以一起使用，以名称-值对的形式提供文档元数据，name 属性给出元数据名称，content 属性给出值。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#name
*/
func Name(value string) *Attribute { return Attr("name", value) }

/*
# HTML 属性：novalidate

此布尔属性指示提交时不应验证表单。如果未设置此属性（因此验证表单），则可以被属于表单的 <button>、<input type="submit"> 或 <input type="image"> 元素上的 formnovalidate 属性覆盖。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form#novalidate
*/
var NoValidate = Attr("novalidate", "")

/*
# HTML 属性：placeholder

有效的占位符文本包括提示预期数据类型的单词或短语，而不是解释或提示。占位符不得用于代替 <label>。由于如果表单控件的值不为 null，占位符不可见，因此使用占位符代替 <label> 作为提示会损害可用性和可访问性。

placeholder 属性由以下输入类型支持：text、search、url、tel、email 和 password。<textarea> 元素也支持它。下面的示例显示了使用 placeholder 属性来解释输入字段的预期格式。

	> 除了 <textarea> 元素外，placeholder 属性不能包含任何换行符 (LF) 或回车符 (CR)。如果值中包含任何一个，占位符文本将被剪切。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/placeholder
*/
func Placeholder(value string) *Attribute { return Attr("placeholder", value) }

/*
# HTML style 全局属性

style 全局属性包含将应用于元素的 CSS 样式声明。推荐在单独的样式文件中定义样式；该属性和 `<style>` 元素主要用于快速试验或临时调整样式。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/style
*/
func Style(value string) *Attribute { return Attr("style", value) }

/*
# HTML tabindex 全局属性

tabindex 全局属性让开发者可以使元素可聚焦，允许或禁止它参与顺序聚焦（通常通过 Tab 键）并指定顺序聚焦时的相对顺序。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/tabindex
*/
func Tabindex(value string) *Attribute { return Attr("tabindex", value) }

/*
# HTML title 全局属性

title 全局属性包含与所属元素相关的提示性文本信息。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/title
*/
func Title(value string) *Attribute { return Attr("title", value) }

/*
# <input> type 属性

用于指定要渲染的表单控件类型的字符串。例如，要创建复选框，需将值设为 checkbox。若省略该属性（或指定未知值），则默认使用 text 类型，生成一个纯文本输入框。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#type
*/
func Type(value string) *Attribute { return Attr("type", value) }

/*
# HTML 属性：rel

rel 属性定义当前文档与链接资源之间的关系。它适用于 `<link>`、`<a>`、`<area>` 与 `<form>` 元素，支持的取值取决于所在元素。

关系类型由 rel 属性的取值决定，值为以空格分隔的唯一关键字集合。与仅用于样式的 class 不同，rel 必须表达机器与人类都能理解的语义标记。可使用的关键字收录于 IANA 链接关系注册表、HTML Living Standard 以及 microformats wiki 的 existing-rel-values 页面；若使用不在这些来源中的关键字，部分 HTML 校验器（如 W3C Markup Validation Service）会给出警告。

下表列出常见关键字，每个值中各关键字必须唯一。

rel 属性主要与 `<link>`、`<a>`、`<area>`、`<form>` 元素相关，但部分关键字仅对其中的子集有效。与所有 HTML 关键字属性相同，这些值不区分大小写。

rel 属性没有默认值。当属性被省略或值全部不受支持时，文档与目标资源之间除了存在超链接外没有额外关系。在这种情况下，对于 `<link>` 与 `<form>`，若 rel 属性缺失、为空或值无效，则该元素不会创建链接；`<a>` 与 `<area>` 仍会创建链接，但没有定义关系。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/rel
*/
func Rel(value string) *Attribute { return Attr("rel", value) }

/*
# HTML 属性：required

布尔型 required 属性存在时，表示用户必须在提交所属表单前为该输入提供值。

该属性适用于 text、search、url、tel、email、password、date、month、week、time、datetime-local、number、checkbox、radio、file 等 `<input>` 类型，以及 `<select>` 与 `<textarea>` 控件。设置后，控件会匹配 `:required` 伪类；若未设置，则匹配 `:optional`。

range 与 color 类型具有默认值，因此不支持 required。color 默认 `#000000`，range 默认位于 min 与 max 中点（若未声明，常见默认为 0 与 100）。hidden 类型也不支持 required，因为用户无法填写隐藏字段；所有按钮类型（包括 image）同样不支持。

对于同名的单选按钮组，只要其中任意一个设置了 required，该组就必须至少选中一个选项，建议要么全部设置 required，要么全部不设，以便维护。对于同名复选框组，仅那些带 required 的复选框是必填项。

\t> 将 `aria-required="true"` 设置在元素上会通知屏幕阅读器该元素必填，但不会影响实际校验逻辑。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/required
*/
var Required = Attr("rel", "")

/*
# SVG 属性：width

width 属性定义元素在用户坐标系中的水平长度。

https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Attribute/width
*/
func Width(value string) *Attribute { return Attr("width", value) }

/*
# `<input>` value 属性

表示输入控件的值。在 HTML 中声明时，该值为初始值；之后可通过 JavaScript 访问对应的 `HTMLInputElement.value` 属性进行修改或读取。value 属性可选，但对 checkbox、radio 与 hidden 类型通常视为必填。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#value
*/
func Value(value string) *Attribute { return Attr("value", value) }

/*
# `<template>` shadowrootmode 属性

为父元素创建 shadow root，是 `Element.attachShadow()` 方法的声明式版本，接受相同的枚举值。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template#shadowrootmode
*/
func ShadowRootMode(value string) *Attribute { return Attr("shadowrootmode", value) }

/*
# HTML slot 全局属性

slot 全局属性用于将元素分配到 shadow DOM 的插槽中：带有 slot 属性的元素会映射到具有同名 `name` 属性的 `<slot>` 元素。通过复用相同的插槽名称，可以让多个元素进入同一插槽；未声明 slot 的元素（若存在）会落入未命名插槽。

示例可参考“Using templates and slots”指南。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/slot
*/
func Slot(value string) *Attribute { return Attr("slot", value) }

/*
# `<a>` target 属性

指定链接 URL 的呈现位置，即浏览上下文（标签页、窗口或 `<iframe>`）的名称。以下关键字具有特殊含义：

  - _self：当前浏览上下文（默认）。
  - _blank：通常为新标签页，用户也可配置为新窗口。
  - _parent：当前上下文的父级，无父级时等同 _self。
  - _top：最顶层的浏览上下文，即当前上下文的最高祖先，无祖先时等同 _self。
  - _unfencedTop：允许嵌入的 fenced frame 导航至顶级上下文；在非 fenced frame 中使用也能成功，但不会作为保留关键字处理。

https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#target
*/
func Target(value string) *Attribute   { return Attr("target", value) }
func Property(value string) *Attribute { return Attr("property", value) }
