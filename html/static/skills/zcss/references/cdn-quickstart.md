# zcss CDN 快速开始

当你需要为 `@sohaha/zcss` 生成可直接复制到浏览器里的代码片段时，读这个文件。

## 推荐 CDN

```html
<script src="https://unpkg.com/@sohaha/zcss/dist/zcss.umd.min.js"></script>
```

全局对象：

```js
window.zcss
```

## 自动初始化

适合能直接在脚本标签里完成配置的普通 HTML 页面。

```html
<!doctype html>
<html lang="zh-CN">
  <body z-block>
    <main class="min-h-screen bg-gray-50 p-6 dark:bg-gray-950">
      <h1 class="text-2xl font-bold text-gray-900 dark:text-white">Hello zcss</h1>
      <button class="mt-4 rounded-lg bg-blue-600 px-4 py-2 text-white hover:bg-blue-700">
        Click
      </button>
    </main>

    <script
      src="https://unpkg.com/@sohaha/zcss/dist/zcss.umd.min.js"
      data-init
      data-config='{
        "setup": {
          "darkMode": "class",
          "preset": true,
          "useVariableSpacing": true
        },
        "observe": {
          "delay": 50,
          "root": "document"
        }
      }'
    ></script>
  </body>
</html>
```

Notes:

- `data-init` 是自动初始化的前提。
- 当 `preset: true` 且你想避免初始化前闪烁时，可以配合 `z-block`。

## 手动初始化

适合需要显式控制顺序或主动调用运行时 API 的页面。

```html
<!doctype html>
<html lang="zh-CN">
  <body>
    <div id="card" class="rounded-xl p-4">Card</div>

    <script src="https://unpkg.com/@sohaha/zcss/dist/zcss.umd.min.js"></script>
    <script>
      zcss.setup({
        darkMode: 'class',
        preset: true,
        useVariableSpacing: true,
        theme: {
          colors: { 'brand-500': '#3b82f6' },
          screens: { md: '768px' },
        },
      })

      zcss.observe()

      document.getElementById('card').className += ' bg-brand-500 text-white md:text-lg'
      zcss.scan(document.getElementById('card'))
    </script>
  </body>
</html>
```

## 运行时 API 速查

- `zcss.setup(config)`：配置 theme 和运行时行为。
- `zcss.observe(root?, opts?)`：先扫描，再持续监听 DOM 和 class 变化。
- `zcss.scan(root?)`：只扫描一次，不创建持续观察器。
- `zcss.disconnect()`：停止当前观察器，并清理延迟模式下的内部轮询。
- `zcss.tw(className)`：编译并注入工具类，返回原始字符串。
- `zcss.css(styleObject)`：返回生成后的类名。
- `zcss.keyframes(frames)`：返回生成后的动画名。

## CSS-in-JS 示例

```html
<script src="https://unpkg.com/@sohaha/zcss/dist/zcss.umd.min.js"></script>
<script>
  const fadeIn = zcss.keyframes({
    '0%': { opacity: '0', transform: 'translateY(8px)' },
    '100%': { opacity: '1', transform: 'translateY(0)' },
  })

  const card = zcss.css({
    backgroundColor: '#111827',
    color: '#fff',
    padding: '16px',
    borderRadius: '16px',
    animation: `${fadeIn} 180ms ease-out`,
    '&:hover': { backgroundColor: '#1f2937' },
  })

  const el = document.createElement('div')
  el.className = card
  el.textContent = '由 zcss.css() 生成'
  document.body.appendChild(el)
  zcss.scan(el)
</script>
```
