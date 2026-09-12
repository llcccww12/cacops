export default async function highlight(elementOrNodeList) {
  if (!window.config || !window.config.HighlightJS || !elementOrNodeList) return;
  const nodes = 'length' in elementOrNodeList ? elementOrNodeList : [elementOrNodeList];
  if (!nodes.length) return;

  const worker = new Worker(new URL('./highlight.worker.js', import.meta.url));

  worker.addEventListener('message', ({data}) => {
    const {index, html} = data;
    nodes[index].outerHTML = html;
  });

  for (let index = 0; index < nodes.length; index++) {
    const node = nodes[index];
    if (!node) continue;
    worker.postMessage({index, html: node.outerHTML});
  }
}
