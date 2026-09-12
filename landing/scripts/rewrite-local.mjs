#!/usr/bin/env node
import fs from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const ASSETS = path.join(ROOT, "assets");
const PAGES = path.join(ROOT, "pages");

const SITE_HOSTS = new Set([
  "sociable-nonogon-193816-ab6b6f8bb.framer.app",
  "sociable-nonogon-193816.framer.app",
]);

function destForUrl(urlStr) {
  let u;
  try {
    u = new URL(urlStr.replace(/&amp;/g, "&"));
  } catch {
    return null;
  }
  const host = u.hostname;
  const pathname = decodeURIComponent(u.pathname);
  const base = path.basename(pathname);
  if (!base || base === "/" || pathname === "/") {
    if (SITE_HOSTS.has(host)) return path.join(PAGES, "index.html");
    return null;
  }

  if (host === "framerusercontent.com") {
    if (pathname.startsWith("/images/")) return path.join(ASSETS, "images", base);
    if (pathname.startsWith("/modules/") || pathname.startsWith("/cms/")) {
      return path.join(ASSETS, "misc", base);
    }
    if (pathname.startsWith("/sites/") && base.endsWith(".json")) return path.join(ASSETS, "json", base);
    if (pathname.startsWith("/sites/") && /\.(png|jpe?g|gif|webp|svg|ico)$/i.test(base)) {
      return path.join(ASSETS, "images", base);
    }
    if (pathname.startsWith("/sites/")) return path.join(ASSETS, "js", base);
    if (pathname.startsWith("/assets/") && /\.woff2?$/i.test(base)) return path.join(ASSETS, "fonts", base);
    if (pathname.startsWith("/assets/") && /\.json$/i.test(base)) return path.join(ASSETS, "json", base);
    if (pathname.startsWith("/assets/") && /\.(mp4|webm|mov)$/i.test(base)) return path.join(ASSETS, "videos", base);
    if (pathname.startsWith("/assets/") && /\.(mjs|js)$/i.test(base)) return path.join(ASSETS, "js", base);
    if (pathname.startsWith("/assets/")) return path.join(ASSETS, "images", base);
    if (pathname.startsWith("/third-party-assets/")) return path.join(ASSETS, "fonts", base);
    return path.join(ASSETS, "misc", base);
  }
  if (host === "fonts.gstatic.com" || host === "fonts.googleapis.com") {
    return path.join(ASSETS, "fonts", base);
  }
  if (SITE_HOSTS.has(host)) {
    const name = pathname.replace(/\/+$/, "").replace(/^\//, "") || "index";
    return path.join(PAGES, `${name}.html`);
  }
  return null;
}

function webPath(toFile) {
  return "/" + path.relative(ROOT, toFile).split(path.sep).join("/");
}

async function walk(dir) {
  const out = [];
  try {
    for (const ent of await fs.readdir(dir, { withFileTypes: true })) {
      const p = path.join(dir, ent.name);
      if (ent.isDirectory()) out.push(...(await walk(p)));
      else out.push(p);
    }
  } catch {
    /* missing dir */
  }
  return out;
}

const URL_RE =
  /https:\\?\/\\?\/(?:framerusercontent\.com|fonts\.gstatic\.com|fonts\.googleapis\.com|sociable-nonogon-193816(?:-ab6b6f8bb)?\.framer\.app)(?:\\?\/[A-Za-z0-9._~\-?=&%,+@]+)+/g;

function unescapeUrl(raw) {
  return raw.replace(/\\\//g, "/").replace(/&amp;/g, "&");
}

async function collect() {
  const files = [
    ...(await walk(PAGES)),
    ...(await walk(ASSETS)),
  ].filter((f) => /\.(html|mjs|js|css|json)$/i.test(f));
  const urls = new Set();
  for (const f of files) {
    if (f.endsWith("MANIFEST.json")) continue;
    const text = await fs.readFile(f, "utf8");
    for (const m of text.match(URL_RE) || []) urls.add(unescapeUrl(m));
  }
  return { files, urls: [...urls] };
}

async function downloadMissing(urls) {
  const missing = [];
  for (const url of urls) {
    const dest = destForUrl(url);
    if (!dest) continue;
    try {
      await fs.access(dest);
    } catch {
      missing.push({ url, dest });
    }
  }
  console.log(`missing assets: ${missing.length}`);
  for (const { url, dest } of missing) {
    await fs.mkdir(path.dirname(dest), { recursive: true });
    try {
      const res = await fetch(url, {
        headers: {
          "User-Agent":
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
        },
        redirect: "follow",
      });
      if (!res.ok) {
        console.log(`fail ${res.status} ${url}`);
        continue;
      }
      await fs.writeFile(dest, Buffer.from(await res.arrayBuffer()));
      console.log(`got  ${res.status} ${path.relative(ROOT, dest)}`);
    } catch (err) {
      console.log(`err  ${url} ${err.message}`);
    }
  }
}

function rewritePageHrefs(text) {
  return text
    .replace(/href="\.\/"/g, 'href="/pages/index.html"')
    .replace(/href="\.\/([a-z0-9\-\/%\u2014]+)"/gi, (all, p) => {
      if (p.includes(".")) return all;
      return `href="/pages/${p}.html"`;
    });
}

function stripQueryLeftovers(text) {
  return text.replace(
    /((?:\/|\.\.\/)+(?:assets\/)?(?:images|fonts|js|json|css|videos|misc)\/[A-Za-z0-9._-]+\.(?:png|webp|svg|jpe?g|gif|ico|woff2|mjs|json|css|mp4|webm|js|framercms))(?:[?;](?:width|height|scale-down-to|amp)[^"'\s,]*)+/gi,
    "$1"
  );
}

function rewrite(text) {
  let out = text.replace(URL_RE, (raw) => {
    const url = unescapeUrl(raw);
    const dest = destForUrl(url);
    if (!dest) return raw;
    return webPath(dest);
  });
  out = rewritePageHrefs(out);
  out = stripQueryLeftovers(out);
  out = out.replace(/(?:https?:)?\/\/framerusercontent\.com\/images\//g, "/assets/images/");
  out = out.replace(/\.\.\/images\//g, "/assets/images/");
  out = out.replace(/\.\.\/fonts\//g, "/assets/fonts/");
  out = out.replace(/\.\.\/json\//g, "/assets/json/");
  out = out.replace(/\.\.\/videos\//g, "/assets/videos/");
  // new URL("./x", "/assets/...") is an invalid base; keep it absolute.
  out = out.replace("`,`/assets/", "`,location.origin+`/assets/");
  return out;
}

const { files, urls } = await collect();
console.log(`scanned ${files.length} files, ${urls.length} unique CDN urls`);
await downloadMissing(urls);

let changedFiles = 0;
for (const f of files) {
  if (f.endsWith("MANIFEST.json")) continue;
  const before = await fs.readFile(f, "utf8");
  const after = rewrite(before);
  if (after !== before) {
    await fs.writeFile(f, after);
    changedFiles++;
    console.log(`rewrote ${path.relative(ROOT, f)}`);
  }
}

const leftover = [];
for (const f of files) {
  if (f.endsWith("MANIFEST.json")) continue;
  const text = await fs.readFile(f, "utf8");
  leftover.push(...(text.match(URL_RE) || []).map((u) => unescapeUrl(u)));
}

console.log(
  JSON.stringify(
    {
      changedFiles,
      leftoverCdn: [...new Set(leftover)].slice(0, 30),
      leftoverCount: new Set(leftover).size,
    },
    null,
    2
  )
);
