// import 'highlight.js/styles/tomorrow-night.css';
import MarkdownIt from 'markdown-it';
// import 'katex/dist/katex.min.css';
import markdownItKatexGpt from 'markdown-it-katex-gpt'
import markdownCopy from './markdown-copy'
function createRenderer() {
	const renderer = new MarkdownIt({
		})
		// .use(require('markdown-it-texmath'), {
		// 	engine: katex,
		// 	...config.texmath
		// })
		// .use(require('markdown-it-mermaid-plugin'))
		.use(require('markdown-it-highlightjs'), {
			auto: true,
		})
		.use(markdownCopy, {
            iconClass: "",
            buttonStyle: 'position: absolute; top: 1px; right: 6px; cursor: pointer; outline: none;',
            iconStyle: 'font-size: 16px; opacity: 0.4;',
            element: '<i class="ri-file-copy-line"></i>',
        }).use(markdownItKatexGpt, {
            delimiters: [
                { left: '\\[', right: '\\]', display: true },
                { left: '\\(', right: '\\)', display: false },
                { left: '$$', right: '$$', display: true },
                { left: '$', right: '$', display: false },
            ]
        })
	return renderer;
}

export default createRenderer;