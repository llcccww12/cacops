#!/usr/bin/env node
/**
 * Inject portal-login assets into all marketing HTML pages.
 */
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const PAGES = path.join(ROOT, "pages");

const META =
  '<meta name="cacops-portal-login" content="http://127.0.0.1:4180/auth/login?from=landing">';
const CSS = '<link rel="stylesheet" href="/assets/css/portal-login.css">';
const JS = '<script src="/assets/js/portal-login.js" defer></script>';

function walk(dir) {
  const out = [];
  for (const ent of fs.readdirSync(dir, { withFileTypes: true })) {
    const p = path.join(dir, ent.name);
    if (ent.isDirectory()) out.push(...walk(p));
    else if (ent.name.endsWith(".html")) out.push(p);
  }
  return out;
}

function inject(html) {
  let out = html;
  if (!out.includes('name="cacops-portal-login"')) {
    out = out.replace("</head>", `  ${META}\n</head>`);
  }
  if (!out.includes("portal-login.css")) {
    out = out.replace("</head>", `  ${CSS}\n</head>`);
  }
  if (!out.includes("portal-login.js")) {
    out = out.replace("</head>", `  ${JS}\n</head>`);
  }
  // Framer 水合会动 head；body 末尾再挂一份，保证点击拦截一定装上
  if (!out.includes("data-cacops-portal-boot")) {
    out = out.replace(
      "</body>",
      `  <script data-cacops-portal-boot src="/assets/js/portal-login.js"></script>\n</body>`,
    );
  }
  return out;
}

let n = 0;
for (const file of walk(PAGES)) {
  const before = fs.readFileSync(file, "utf8");
  const after = inject(before);
  if (after !== before) {
    fs.writeFileSync(file, after);
    n += 1;
  }
}
console.log(`portal-login injected into ${n} html files`);
