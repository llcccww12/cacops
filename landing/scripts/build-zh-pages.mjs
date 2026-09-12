#!/usr/bin/env node
/**
 * Generate Chinese pages under pages/zh/ from English templates.
 * Usage: node scripts/build-zh-pages.mjs
 */
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, "..");
const PAGES = path.join(ROOT, "pages");
const ZH_ROOT = path.join(PAGES, "zh");
const LOCALES = path.join(ROOT, "locales");

function readJson(file) {
  return JSON.parse(fs.readFileSync(path.join(LOCALES, file), "utf8"));
}

const base = readJson("zh.json");
const pagesTranslation = readJson("pages-translation.json");
const pagesExtra = readJson("pages-extra.json");

function buildUnifiedPool() {
  const pool = {};

  const add = (map) => {
    if (!map || typeof map !== "object") return;
    for (const [en, zh] of Object.entries(map)) {
      if (en && zh && en !== zh) pool[en] = zh;
    }
  };

  add(base.global);
  add(pagesTranslation.global_extra);
  add(pagesExtra.global_extra);
  add(base.letterSpans);

  for (const [key, value] of Object.entries(base)) {
    if (key.startsWith("pages/") && typeof value === "object") add(value);
  }
  for (const [key, value] of Object.entries(pagesTranslation)) {
    if (key.startsWith("pages/") && typeof value === "object") add(value);
  }

  delete pool.Work;
  pool[">Work</p>"] = ">案例</p>";
  pool[">Scalar</p>"] = ">CacOps</p>";
  pool["hi@scalar.com"] = "hi@cacops.com";

  return Object.entries(pool).sort((a, b) => b[0].length - a[0].length);
}

const unifiedReplacements = buildUnifiedPool();
const letterSpanTexts = {
  ...(base.letterSpans || {}),
};

function listHtmlFiles(dir, baseDir = dir) {
  const out = [];
  for (const name of fs.readdirSync(dir)) {
    const full = path.join(dir, name);
    const stat = fs.statSync(full);
    if (stat.isDirectory()) {
      if (name === "zh") continue;
      out.push(...listHtmlFiles(full, baseDir));
    } else if (name.endsWith(".html")) {
      out.push(path.relative(baseDir, full));
    }
  }
  return out;
}

function rewritePageLinks(html) {
  return html.replace(/href="\/pages\/(?!zh\/)/g, 'href="/pages/zh/');
}

function applyRawReplacements(html, relPath) {
  const pageKey = `pages/${relPath.replace(/\\/g, "/")}`;
  const pairs = [...(base.raw?.global || []), ...(base.raw?.[pageKey] || [])];
  let out = html;
  for (const [from, to] of pairs) {
    out = out.split(from).join(to);
  }
  return out;
}

function replaceLetterSpanParagraphs(html) {
  return html.replace(
    /<p([^>]*class="framer-text[^"]*"[^>]*)>([\s\S]*?)<\/p>/g,
    (full, attrs, inner) => {
      if (!/<span[^>]*>/.test(inner)) return full;
      const text = [...inner.matchAll(/<span[^>]*>([^<]*)<\/span>/g)]
        .map((m) => m[1])
        .join("");
      const zh = letterSpanTexts[text];
      if (zh) return `<p${attrs}>${zh}</p>`;
      return full;
    }
  );
}

function translateHtml(html, relPath) {
  let out = html;
  out = out.replace(/<html lang="en"/, '<html lang="zh-CN"');
  out = rewritePageLinks(out);
  out = applyRawReplacements(out, relPath);
  out = replaceLetterSpanParagraphs(out);

  for (const [en, zh] of unifiedReplacements) {
    out = out.split(en).join(zh);
  }

  return out;
}

function injectAssets(html, { runtime = false } = {}) {
  const css = '<link rel="stylesheet" href="/assets/css/lang-switch.css">';
  const js = '<script src="/assets/js/lang-switch.js" defer></script>';
  const portalMeta =
    '<meta name="cacops-portal-login" content="http://127.0.0.1:4180/auth/login">';
  const portalCss =
    '<link rel="stylesheet" href="/assets/css/portal-login.css">';
  const portalJs = '<script src="/assets/js/portal-login.js" defer></script>';
  const runtimeJs =
    '<script src="/assets/js/lang-runtime.js" defer></script>';
  let out = html;
  if (!out.includes("lang-switch.css")) {
    out = out.replace("</head>", `  ${css}\n</head>`);
  }
  if (!out.includes("lang-switch.js")) {
    out = out.replace("</head>", `  ${js}\n</head>`);
  }
  if (!out.includes('name="cacops-portal-login"')) {
    out = out.replace("</head>", `  ${portalMeta}\n</head>`);
  }
  if (!out.includes("portal-login.css")) {
    out = out.replace("</head>", `  ${portalCss}\n</head>`);
  }
  if (!out.includes("portal-login.js")) {
    out = out.replace("</head>", `  ${portalJs}\n</head>`);
  }
  if (runtime && !out.includes("lang-runtime.js")) {
    out = out.replace("</head>", `  ${runtimeJs}\n</head>`);
  }
  return out;
}

const files = listHtmlFiles(PAGES);
let zhCount = 0;

for (const rel of files) {
  const src = path.join(PAGES, rel);
  const enHtml = injectAssets(fs.readFileSync(src, "utf8"));
  fs.writeFileSync(src, enHtml);

  const zhDir = path.dirname(path.join(ZH_ROOT, rel));
  fs.mkdirSync(zhDir, { recursive: true });
  const zhHtml = injectAssets(translateHtml(fs.readFileSync(src, "utf8"), rel), {
    runtime: true,
  });
  fs.writeFileSync(path.join(ZH_ROOT, rel), zhHtml);
  zhCount += 1;
}

console.log(`Unified translation entries: ${unifiedReplacements.length}`);
console.log(`Updated ${files.length} English pages with lang switch assets.`);
console.log(`Generated ${zhCount} Chinese pages under pages/zh/.`);
