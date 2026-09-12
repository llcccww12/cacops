
import { initClipboard } from '~/utils';
import SparkMD5 from "spark-md5";

try {
	this === window;
    initClipboard('.markdown-it-code-copy');
    
}
catch (_err) {
}



function renderCode(origRule, options) {
	return (...args) => {
		const [tokens, idx] = args;
		const content = tokens[idx].content
			.replaceAll('"', '&quot;')
			.replaceAll("'", "&apos;");
		const origRendered = origRule(...args);

		if (content.length === 0)
			return origRendered;

		return `
<div style="position: relative">
	${origRendered}
	<button class="ui poping inline up markdown-it-code-copy " id="clipboard-${sparkMD5Hash(content)}" 
        data-clipboard-text="${content}" style="${options.buttonStyle}" title="Copy"
        data-position="top center" data-variation="inverted tiny" data-success="复制成功"
        data-content="复制" data-original="复制">
		<span style="${options.iconStyle}" class="${options.iconClass}">${options.element}</span>
	</button>
</div>
`;
	};
}
function sparkMD5Hash(str = '') {
    return SparkMD5.hash(str) + Math.random().toString().replace('0.', '');
}
const markdownCopy = (md, options) => {
	
	md.renderer.rules.code_block = renderCode(md.renderer.rules.code_block, options);
	md.renderer.rules.fence = renderCode(md.renderer.rules.fence, options);
};

export default markdownCopy