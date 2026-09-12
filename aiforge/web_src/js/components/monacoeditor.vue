
<template>
  <div ref="container" class="monaco-editor" style="height: calc(100vh - 114px)" />
</template>

<script>
let monaco;
import configSuggest, { tipTxt } from './monaco-hint';
export default {
  name: "AcMonaco",
  inject: ["reload"],
  props: {
    opts: {
      type: Object,
      default() {
        return {};
      },
    },
    height: {
      type: Number,
      default: 300,
    },

    monacaValue: {
      type: String,
      default: "false",
    },
    onChange: {
      type: Function,
      default: new Function
    },
    errorLine: {
      type: Number,
      default: 3,
    },
    errorContent: {
      type: String,
      default: "",
    },

  },
  data() {
    return {
      isDiff: false,
      // 主要配置
      defaultOpts: {
        lineNumbersMinChars: 0,
        lineDecorationsWidth: 0,
        ignoreTrimWhitespace: false,
        // renderSideBySide:false,  diff单个还是双视图
        glyphMargin: false,
        value: "", // 编辑器的值
        wordWrap: "on",
        theme: "vs-dark", // 编辑器主题：vs, hc-black, or vs-dark，更多选择详见官网
        roundedSelection: false, // 右侧不显示编辑器预览框
        autoIndent: true, // 自动缩进
        quickSuggestions: true, // 快速搜索

        enableSplitViewResizing: false,
        scrollBeyondLastLine: false,
        // quickSuggestions: false, //智能提示
        renderLineHighlight: false, //选中行外部边框
        lineHeight: 22,

        renderWhitespace: 'boundary',
      },

      // 编辑器对象
      monacoEditor: {},
      oldValue: "123123/n\n",
      newValue: "123121311111111",
    };
  },
  watch: {
    opts: {

    },
    errorLine: {
      handler(value) {

        if (this.errorLine) {
          let instance = this.monacoEditor;
          instance.changeViewZones((changeAccessor) => {
            var domNode = document.createElement('div');
            domNode.style.padding = '10px 20px';
            domNode.style.width = 'calc(100% - 20px)';
            domNode.className = 'my-error-line-wrp';
            domNode.innerHTML = this.errorContent || "";
            changeAccessor.addZone({
              afterLineNumber: this.errorLine || 11,
              heightInLines: 3,
              domNode: domNode,
            });
          });
          var overlayWidget = {
            domNode: null,
            getId: function () {
              return 'my.overlay.widget';
            },
            getDomNode: function () {
              if (!this.domNode) {
                this.domNode = document.createElement('div');
                this.domNode.innerHTML = '';
                this.domNode.style.width = '100%';
                this.domNode.style.padding = '20px 100px';
                this.domNode.style.right = '0px';
                this.domNode.style.top = '50px';
                this.domNode.style.position = 'relative';
                this.domNode.style.color = '#333';
              }
              return this.domNode;
            },
            getPosition: function () {
              return null;
            },
          };

          setTimeout(() => {
            this.monacoEditor.deltaDecorations(
              [],
              [
                {
                  range: new monaco.Range(this.errorLine, 1, this.errorLine, 1),
                  options: {
                    isWholeLine: true,
                    className: "myContentClass",
                    glyphMarginClassName: "myGlyphMarginClass",
                  },
                },
              ]
            );
          },500)
          instance.addOverlayWidget(overlayWidget);
          // instance.revealPositionInCenter(11,1);
          instance.revealPositionInCenter({ lineNumber: value, column: 1 });
        }
      }
    },
    monacaValue: {
      handler(value) {
        if (this.isDiff) return;
        const editor = this.monacoEditor
        // const output = isBase64(value)
        //   ? Buffer.from(value, "base64").toString("utf8")
        //   : value;
        const output = value;
        const positions = this.monacoEditor.getPosition()
        this.monacoEditor.setValue(output);
        this.monacoEditor.setPosition({
          lineNumber: positions.lineNumber,
          column: positions.column
        })
      },
      deep: true
    },
    // monacaValue(value) {
    //   alert(value)
    //   console.log(isBase64);
    //   const output = isBase64(value)
    //     ? Buffer.from(value, "base64").toString("utf8")
    //     : value;
    //   this.monacoEditor.setValue(output);
    // },
  },
  beforeDestroy() {
    document.removeEventListener("keyup", this.onSaveHandler);
    window.removeEventListener("resize", this.onResize, false);
  },
  mounted() {
    document.addEventListener("keyup", this.onSaveHandler, false);
    window.addEventListener("resize", this.onResize, false);

    this.init();
  },
  methods: {
    onResize() {
      if (this.monacoEditor) {
        this.monacoEditor.layout()
      }
    },
    async init() {
      monaco = await import( 'monaco-editor')
      try {
        // 初始化container的内容，销毁之前生成的编辑器
        this.$refs.container.innerHTML = "";
        // 生成编辑器配置
        const editorOptions = Object.assign(this.defaultOpts, this.opts);
        if (!this.isDiff) {
          // 初始化编辑器实例
          this.monacoEditor = monaco.editor.create(
            this.$refs.container,
            editorOptions
          );
          this.monacoEditor.onDidChangeModelContent(
            event => {
              this.$emit("onChange", this.monacoEditor.getValue());
            },
          );


          this.monacoEditor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS, () => console.log("saved"))

          this.monacoEditor.addAction({
            id: 'd123123',
            label: 'Paste',
            contextMenuGroupId: '9_cutcopypaste',
            run: (editor) => {
              navigator.clipboard.readText().then(text => {
                let selection = editor.getSelection()
                let id = { major: 1, minor: 1 }
                let textFromClipboard = text
                let op = { identifier: id, range: selection, text: textFromClipboard, forceMoveMarkers: true }
                editor.executeEdits("my-source", [op])
              })
              // exportRaw(filename || 'educoder.txt', instance.getValue());
            },
          });
          if (this.monacaValue) {
            // const output = isBase64(this.monacaValue)
            //   ? Buffer.from(this.monacaValue, "base64").toString("utf8")
            //   : this.monacaValue;
            this.monacoEditor.setValue(this.monacaValue);
          }

          configSuggest(monaco)
          // 编辑器内容发生改变时触发
        } else {
          editorOptions.readOnly = true;
          editorOptions.language = "javascript";
          // editorOptions.inlineHints = true;
          // 初始化编辑器实例
          this.monacoDiffInstance = monaco.editor.createDiffEditor(
            this.$refs.container,
            editorOptions
          );
          this.monacoDiffInstance.setModel({
            original: monaco.editor.createModel(
              this.oldValue,
              editorOptions.language
            ),
            modified: monaco.editor.createModel(
              this.newValue,
              editorOptions.language
            ),
          });
        }

      } catch (error) {
        console.error("error-init:", error)
      }
    },
    upDateDiff(val) {
      this.monacoDiffInstance.updateOptions({
        renderSideBySide: !val,
      });
    },
    onSaveHandler(e) {
      if (this.isDiff) return;
      if (
        (window.navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey) &&
        e.keyCode === 83
      ) {
        e.preventDefault();
      }
      setTimeout(() => {
        this.$emit("handleSave", this.monacoEditor.getValue());
      }, 100)
    },
    // 供父组件调用手动获取值
    getVal() {
      return this.monacoEditor.getValue();
    },
    setLanguage(lang) {
      setTimeout(() => {
        if (monaco.editor.setModelLanguage) {
          monaco.editor.setModelLanguage(this.monacoEditor.getModel(), lang)
        }
      }), 300;
      // alert(JSON.stringify(this.monacoEditor.getModel()))
    },
    updateOptions(opts) {
      const editorOptions = Object.assign(this.defaultOpts, opts);
      this.monacoEditor.updateOptions(editorOptions);
    },
    setTheme(theme) {
      monaco.editor.setTheme(theme);
    },
    setDiff(oldContent, newContent) {
      // this.oldValue = isBase64(oldContent) ? Buffer.from(oldContent, "base64").toString("utf8") : oldContent;
      // this.newValue = isBase64(newContent) ? Buffer.from(newContent, "base64").toString("utf8") : newContent;
      this.oldValue = oldContent;
      this.newValue = newContent;
      this.isDiff = true;
      this.init();
    },
    setCloseDiff() {
      if (this.isDiff) {
        this.isDiff = false;
        this.init();
      }
    },
  },
};
</script>

<style>
.margin-view-overlays {
  background: #f5f5f5;
}

.my-error-line-wrp {
  width: calc(100% - 20px) !important;
  background: rgba(245, 0, 0, 0.7) !important;
  height: auto !important;
  color: #FFF;
}

.myGlyphMarginClass {
  background: red;
}

/* .view-lines {
  padding-left: 10px;
} */
</style>
