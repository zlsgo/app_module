# zview CDN 快速开始

当你需要为 `@sohaha/zview` 生成可直接运行的浏览器示例时，读这个文件。

## 推荐 CDN

```html
<script src="https://unpkg.com/@sohaha/zview/dist/zview.umd.min.js"></script>
```

初始化后的全局对象：

```js
window.Zview
window.$z
```

## 自动初始化

```html
<!doctype html>
<html lang="zh-CN">
  <body>
    <section z-signals='{"count":0,"open":false}'>
      <button z-on-click="$count++">+1</button>
      <button z-on-click="$open = !$open">toggle</button>
      <p z-text="$count"></p>
      <div z-show="$open">panel</div>
    </section>

    <script
      src="https://unpkg.com/@sohaha/zview/dist/zview.umd.min.js"
      data-init
      data-config='{
        "setup": {
          "debug": true,
          "timeout": 8000,
          "formSerializationFormat": "json"
        }
      }'
    ></script>
  </body>
</html>
```

## 手动初始化

```html
<!doctype html>
<html lang="zh-CN">
  <body>
    <section id="app" z-signals='{"count":0}'>
      <button z-on-click="$count++">+1</button>
      <span z-text="$count"></span>
    </section>

    <script src="https://unpkg.com/@sohaha/zview/dist/zview.umd.min.js"></script>
    <script>
      Zview.init({
        debug: true,
        timeout: 8000,
        formSerializationFormat: 'json',
      })
    </script>
  </body>
</html>
```

## 核心 API 速查

- `Zview.init(config?)`：初始化全局引擎，并注入 `window.$z`。
- `Zview.activate(element)`：激活一棵子树。
- `Zview.deactivate(element)`：停用一棵子树并清理副作用。
- `Zview.register(directive)`：在初始化后注册自定义指令。
- `Zview.getEngine()`：返回全局引擎实例或 `null`。

## Signal 示例

```html
<section z-signals='{"price":12,"qty":3}'>
  <div z-computed-total="$price * $qty"></div>
  <output z-text="$total"></output>
  <input type="number" z-bind="qty" />
</section>
```

## 请求示例

```html
<section>
  <button
    z-req="GET /partials/card"
    z-trigger="click"
    z-target="#result"
    z-swap="inner"
  >
    加载卡片
  </button>

  <div id="result"></div>
</section>
```

## 事件示例

```html
<input z-on-input.debounce.300="$query = event.target.value" />
<button z-on-click.prevent.stop="submitSearch()">搜索</button>
```

## morph / preserve 示例

局部刷新时用 `z-swap="morph"`，已复用节点会保留用户输入、焦点和选区；第三方组件等不可重建节点用 `z-preserve`（响应里也必须返回同 id 的占位节点）：

```html
<section>
  <form>
    <input type="text" z-bind="query" />
    <button
      z-req="GET /partials/results"
      z-trigger="submit"
      z-swap="morph"
      z-target="#results"
    >
      搜索
    </button>
  </form>

  <div id="results">
    <div id="player" z-preserve><!-- 第三方播放器实例，不重建 --></div>
    <ul><!-- 结果列表 --></ul>
  </div>
</section>
```

服务端响应示例（播放器占位仍带 marker + 同 id）：

```html
<div id="results">
  <div id="player" z-preserve></div>
  <ul><li>新结果</li></ul>
</div>
```
