
import ClipboardJS from 'clipboard';

export default async function initClipboard(elements) {
  const els = elements || document.querySelectorAll(".clipboard");
  if (!els || !els.length) return;
  const clipboard = new ClipboardJS(els);
  window.$ && $().popup && $(els).popup();
  clipboard.on("success", (e) => {
    e.clearSelection();
    const popUpEl =  $(e.trigger);
    popUpEl.popup("destroy");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-success")
    );
    popUpEl.popup("show");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-original")
    );
  });

  clipboard.on("error", (e) => {
    const popUpEl =  $(e.trigger);
    popUpEl.popup("destroy");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-error")
    );
    popUpEl.popup("show");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-original")
    );
  });
}
