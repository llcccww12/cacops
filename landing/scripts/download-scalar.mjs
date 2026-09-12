#!/usr/bin/env node
/**
 * Mirror the public Scalar Framer site into a local dump.
 * Source: https://sociable-nonogon-193816-ab6b6f8bb.framer.app/
 */

import fs from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const OUT = ROOT;
const CONCURRENCY = 8;
const SITE_HOSTS = [
  "sociable-nonogon-193816-ab6b6f8bb.framer.app",
  "sociable-nonogon-193816.framer.app",
];
const ORIGIN = `https://${SITE_HOSTS[0]}`;
const SITE_ID = "2In5ROfZwFkDIVZgZtLHuX";

const PAGES = [
  `${ORIGIN}/`,
  `${ORIGIN}/404`,
  `${ORIGIN}/about`,
  `${ORIGIN}/blog`,
  `${ORIGIN}/contact`,
  `${ORIGIN}/pricing`,
  `${ORIGIN}/work`,
  `${ORIGIN}/legals/privacy-policy`,
  `${ORIGIN}/legals/terms-of-service`,
  `${ORIGIN}/blog/why-most-automation-projects-stall-at-the-pilot-stage`,
  `${ORIGIN}/blog/designing-an-ai-agent-that-actually-ships-work`,
  `${ORIGIN}/blog/the-operations-audit-that-saves-twenty-hours-a-week`,
  `${ORIGIN}/blog/from-spreadsheet-chaos-to-a-single-source-of-truth`,
  `${ORIGIN}/blog/how-kepler-cut-lead-response-time-by-65`,
  `${ORIGIN}/blog/what-to-automate-first-%E2%80%94-a-decision-framework`,
  `${ORIGIN}/work/how-kepler-cut-lead-response-time-by-65`,
  `${ORIGIN}/work/how-altura-increased-qualified-leads-by-48`,
  `${ORIGIN}/work/how-novex-improved-conversion-rate-by-2.3x`,
  `${ORIGIN}/work/how-meridian-automated-dispatch-across-three-depots`,
  `${ORIGIN}/work/how-halcyon-cleared-a-three-week-intake-backlog`,
  `${ORIGIN}/work/how-brightfold-tripled-content-output-without-hiring`,
];

const SEED_ASSETS = [
  `https://framerusercontent.com/sites/${SITE_ID}/script_main.DAR_FhIi.mjs`,
  `https://framerusercontent.com/sites/${SITE_ID}/searchIndex-MWWLCicVHnle.json`,
];

const ALLOW_HOSTS = new Set([
  ...SITE_HOSTS,
  "framerusercontent.com",
  "fonts.gstatic.com",
  "fonts.googleapis.com",
]);

const SKIP_HOSTS = new Set([
  "events.framer.com",
  "app.framerstatic.com",
  "api.framer.com",
]);

const URL_RE =
  /https?:\/\/(?:framerusercontent\.com|fonts\.gstatic\.com|fonts\.googleapis\.com|sociable-nonogon-193816(?:-ab6b6f8bb)?\.framer\.app)[^\s"'`<>)\\]]+/g;
const IMPORT_RE =
  /(?:from|import)\s*["']([^"']+)["']|import\s*\(\s*["']([^"']+)["']\s*\)/g;

const seen = new Set();
const queue = [];
const results = [];

function normalizeUrl(raw, base) {
  if (!raw) return null;
  let u = raw.trim().replace(/\\u002F/g, "/").replace(/\\\//g, "/");
  u = u.replace(/[),;]+$/, "");
  u = u.replace(/&quot;.*$/, "");
  u = u.replace(/\\+$/, "");
  try {
    const url = new URL(u, base);
    if (!["http:", "https:"].includes(url.protocol)) return null;
    if (SKIP_HOSTS.has(url.hostname)) return null;
    if (!ALLOW_HOSTS.has(url.hostname)) return null;
    url.hash = "";
    return url.toString();
  } catch {
    return null;
  }
}

function enqueue(raw, base) {
  const url = normalizeUrl(raw, base);
  if (!url || seen.has(url)) return;
  seen.add(url);
  queue.push(url);
}

function localPathFor(urlStr) {
  const url = new URL(urlStr);
  let pathname = decodeURIComponent(url.pathname);
  if (SITE_HOSTS.includes(url.hostname)) {
    if (pathname === "/" || pathname === "") return path.join(OUT, "pages", "index.html");
    const name = pathname.replace(/\/+$/, "").replace(/^\//, "") || "index";
    return path.join(OUT, "pages", `${name}.html`);
  }
  if (url.hostname === "framerusercontent.com") {
    const base = path.basename(pathname.split("?")[0]);
    if (pathname.startsWith("/images/")) {
      return path.join(OUT, "assets", "images", base);
    }
    if (pathname.startsWith("/modules/") || pathname.startsWith("/cms/")) {
      const ext = path.extname(pathname).toLowerCase();
      if (ext === ".framercms" || pathname.includes("framercms")) {
        return path.join(OUT, "assets", "misc", base);
      }
      if (ext === ".js" || ext === ".mjs") {
        return path.join(OUT, "assets", "misc", base);
      }
      return path.join(OUT, "assets", "misc", base);
    }
    if (pathname.startsWith("/assets/")) {
      const ext = path.extname(pathname).toLowerCase();
      const folder =
        ext === ".woff2" || ext === ".woff" || ext === ".ttf"
          ? "fonts"
          : ext === ".json"
            ? "json"
            : ext === ".mp4" || ext === ".webm" || ext === ".mov"
              ? "videos"
              : ext === ".mjs" || ext === ".js"
                ? "js"
                : "misc";
      return path.join(OUT, "assets", folder, base);
    }
    if (pathname.startsWith("/sites/")) {
      const ext = path.extname(pathname).toLowerCase();
      const folder =
        ext === ".json"
          ? "json"
          : ext === ".png" || ext === ".ico" || ext === ".svg" || ext === ".webp" || ext === ".jpg" || ext === ".jpeg" || ext === ".gif"
            ? "images"
            : "js";
      return path.join(OUT, "assets", folder, base);
    }
    if (pathname.startsWith("/third-party-assets/")) {
      return path.join(OUT, "assets", "fonts", base);
    }
    return path.join(OUT, "assets", "misc", pathname.replace(/\//g, "_").replace(/^_/, ""));
  }
  if (url.hostname === "fonts.gstatic.com" || url.hostname === "fonts.googleapis.com") {
    return path.join(OUT, "assets", "fonts", path.basename(pathname) || "google-fonts.css");
  }
  return path.join(OUT, "assets", "misc", `${url.hostname}${pathname}`.replace(/\//g, "_"));
}

function discover(text, baseUrl) {
  if (!text) return;
  for (const m of text.match(URL_RE) || []) enqueue(m, baseUrl);
  if (/\.mjs(\?|$)/.test(baseUrl) || /\.js(\?|$)/.test(baseUrl) || baseUrl.includes("script_main")) {
    let m;
    IMPORT_RE.lastIndex = 0;
    while ((m = IMPORT_RE.exec(text))) {
      const spec = m[1] || m[2];
      if (spec && !spec.startsWith("data:")) enqueue(spec, baseUrl);
    }
    const relative = text.matchAll(/["'](\.\/[^"']+\.mjs)["']/g);
    for (const r of relative) enqueue(r[1], baseUrl);
    // CMS chunks relative to module URL
    const cms = text.matchAll(/["'`](\.\/[^"'`]+\.framercms)["'`]/g);
    for (const r of cms) {
      try {
        const moduleUrl = new URL(baseUrl);
        if (moduleUrl.pathname.includes("/modules/")) {
          const cmsBase = moduleUrl.href.replace("/modules/", "/cms/");
          enqueue(r[1], cmsBase);
        } else {
          enqueue(r[1], baseUrl);
        }
        // also try pairing from absolute module base in new URL(`./x`, `https://.../modules/.../file.js`)
      } catch {}
    }
    // new URL(`./foo.framercms`, `https://framerusercontent.com/modules/.../x.js`)
    const pairs = text.matchAll(/new URL\(\s*[`'"](\.\/[^`'"]+\.framercms)[`'"]\s*,\s*[`'"](https:\/\/framerusercontent\.com\/modules\/[^`'"]+)[`'"]/g);
    for (const r of pairs) {
      const abs = new URL(r[1], r[2].replace("/modules/", "/cms/")).toString();
      enqueue(abs, baseUrl);
    }
  }
  const cssUrls = text.matchAll(/url\(\s*['"]?([^'")]+)['"]?\s*\)/g);
  for (const c of cssUrls) {
    const v = c[1];
    if (v.startsWith("http") || v.startsWith("//") || v.startsWith("./") || v.startsWith("/")) {
      enqueue(v, baseUrl);
    }
  }
  const srcs = text.matchAll(/(?:src|href)=["']([^"']+)["']/g);
  for (const s of srcs) enqueue(s[1], baseUrl);

  if (SITE_HOSTS.includes(new URL(baseUrl).hostname) || /html/i.test(baseUrl)) {
    const relPages = text.matchAll(/href=["']\.\/([a-z0-9\-\/%\u2014]+)["']/gi);
    for (const r of relPages) {
      if (r[1].includes(".")) continue;
      enqueue(`${ORIGIN}/${r[1]}`, baseUrl);
    }
  }
}

async function fetchOne(url) {
  const dest = localPathFor(url);
  await fs.mkdir(path.dirname(dest), { recursive: true });
  const res = await fetch(url, {
    headers: {
      "User-Agent":
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
      Accept: "*/*",
    },
    redirect: "follow",
  });
  const isSitePage = SITE_HOSTS.includes(new URL(url).hostname);
  if (!res.ok && !(isSitePage && res.status === 404)) {
    results.push({ url, dest: path.relative(ROOT, dest), status: res.status, ok: false });
    return;
  }
  const ctype = res.headers.get("content-type") || "";
  const buf = Buffer.from(await res.arrayBuffer());
  await fs.writeFile(dest, buf);
  results.push({
    url,
    dest: path.relative(ROOT, dest),
    status: res.status,
    bytes: buf.length,
    type: ctype.split(";")[0],
    ok: true,
  });
  const isText =
    /html|javascript|json|css|svg|xml|text\//.test(ctype) ||
    /\.(html|js|mjs|json|css|svg|txt|xml)$/i.test(dest);
  if (isText) {
    discover(buf.toString("utf8"), url);
  }
}

async function runPool() {
  const workers = Array.from({ length: CONCURRENCY }, async () => {
    while (queue.length) {
      const url = queue.shift();
      if (!url) return;
      try {
        await fetchOne(url);
        const last = results.at(-1);
        process.stdout.write(`${last?.ok ? "ok " : "fail"} ${last?.status ?? ""} ${url}\n`);
      } catch (err) {
        results.push({ url, ok: false, error: String(err) });
        process.stdout.write(`err ${url} ${err}\n`);
      }
    }
  });
  await Promise.all(workers);
}

for (const p of PAGES) enqueue(p);
for (const a of SEED_ASSETS) enqueue(a);

await fs.mkdir(OUT, { recursive: true });
while (queue.length) {
  await runPool();
}

const ok = results.filter((r) => r.ok);
const fail = results.filter((r) => !r.ok);
const manifest = {
  source: ORIGIN + "/",
  downloadedAt: new Date().toISOString(),
  counts: {
    total: results.length,
    ok: ok.length,
    failed: fail.length,
    images: ok.filter((r) => /\/images\//.test(r.url) || /^image\//.test(r.type || "")).length,
    js: ok.filter((r) => /\.mjs($|\?)/.test(r.url)).length,
    fonts: ok.filter((r) => /woff2|\/fonts\//.test(r.url)).length,
    videos: ok.filter((r) => /\.(mp4|webm|mov)($|\?)/i.test(r.url)).length,
    pages: ok.filter((r) => r.dest.startsWith("pages/")).length,
  },
  failed: fail,
  files: ok.sort((a, b) => a.dest.localeCompare(b.dest)),
};

await fs.writeFile(path.join(OUT, "MANIFEST.json"), JSON.stringify(manifest, null, 2));
console.log("\nDone", manifest.counts);
if (fail.length) {
  console.log("Failed:");
  for (const f of fail) console.log(" ", f.url, f.status || f.error);
}
