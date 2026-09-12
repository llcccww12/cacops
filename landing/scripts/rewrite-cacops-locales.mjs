#!/usr/bin/env node
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const LOCALES = path.join(ROOT, "locales");

const OVERRIDES = {
  "Solutions": "解决方案",
  "Pricing": "定价",
  "Blog": "博客",
  "About": "关于我们",
  "Work": "案例",
  "Get Started": "立即开始",
  "Navigation": "导航",
  "Results": "成果",
  "Process": "流程",
  "FAQs": "常见问题",
  "FAQ": "常见问题",
  "Legal": "法律信息",
  "Privacy policy": "隐私政策",
  "Terms of service": "服务条款",
  "404 Page": "404 页面",
  "Socials": "社交媒体",
  "©2026 Scalar. All rights reserved.": "©2026 CacOps. 保留所有权利。",
  "Built by Chayan": "由 CacOps 打造",
  "Use for Free": "免费使用",
  "High-quality Framer template crafted for AI automation agencies":
    "CacOps 智算服务平台 — 算力、数据、模型与 Token 一站式交付",
  "Create a free website with Framer, the website builder loved by startups, designers and agencies.":
    "CacOps 智算服务平台：从底层算力到 Token 交付的一站式能力。",
  "Book a free 30-min audit. We'll map your biggest automation opportunity.":
    "预约 30 分钟免费评估，我们将为你梳理算力、数据与模型落地路径。",
  "Scale your growth, not your workload.": "用智算放大业务，而不是堆砌成本。",
  "Frequently Asked Questions": "常见问题",
  "What does Scalar automate?": "CacOps 提供哪些服务？",
  "Scalar automates repetitive business tasks like lead management, onboarding, reporting, support workflows, and internal operations using AI systems.":
    "CacOps 提供一站式智算服务：底层算力资源、数据生成与标注、模型训练与推理，以及 Token 按量售卖。",
  "Does Scalar build custom AI workflows?": "CacOps 能定制模型与推理方案吗？",
  "Yes. Scalar creates custom automation systems tailored to your business, tools, and operational needs.":
    "可以。我们按行业场景定制训练、微调与推理方案，并对接你现有的业务系统。",
  "Can Scalar integrate with our current tools?": "CacOps 能与现有系统集成吗？",
  "Absolutely. Scalar works with platforms like Slack, Notion, HubSpot, Airtable, Google Workspace, and other third-party tools.":
    "当然。支持对接企业现有云平台、数据仓库、业务中台与 API 网关等环境。",
  "How long does a Scalar project take?": "CacOps 项目通常需要多久？",
  "Most Scalar projects are completed within 1 to 4 weeks depending on the scope and complexity.":
    "根据范围与复杂度，多数项目可在 1 到 4 周完成首期上线。",
  "Is Scalar only for large companies?": "CacOps 只服务大公司吗？",
  "No. Scalar works with startups, agencies, SaaS companies, and growing teams looking to scale operations efficiently.":
    "不是。我们同时服务政企、成长型科技公司与需要弹性算力 / Token 的 AI 团队。",
  "What makes Scalar different?": "CacOps 有什么不同？",
  "Scalar focuses on practical AI automation systems that improve real business workflows, not just experimental AI ideas.":
    "我们不是单点卖卡或卖接口，而是打通算力、数据、模型与 Token 的全链路智算平台。",
  "About — Scalar": "关于我们 — CacOps",
  "Pricing — Scalar": "定价 — CacOps",
  "We are a small team that builds automation systems operators actually keep using.":
    "我们是一支专注智算落地的团队，为客户交付可运营的算力、数据与模型能力。",
  "We build systems operators keep using": "我们打造真正能持续运营的智算体系",
  "Scalar is a small team of operators and engineers. We have run the messy version of your process ourselves, which is why we build for the handover rather than the demo.":
    "CacOps 由智算架构、数据工程与模型服务专家组成。我们深知碎片化采购的痛点，因此按可交接、可扩容的方式交付，而不是只做演示。",
  "We started Scalar after watching the same project fail three times":
    "我们在多次看到「只买算力、无法落地」后创立了 CacOps",
  "Between us we spent a decade inside operations teams at companies that grew faster than their processes. Every one of them ran the same play: hire a contractor, build an impressive prototype, demo it to leadership, and watch it quietly rot once the contractor left.":
    "过去十年，我们见证大量团队只采购 GPU 或接口，却缺少数据、训练与计费闭环。结果是：演示很漂亮，上线后无人维护、成本不可控、效果无法复现。",
  "The pattern was never a tooling problem. It was that nobody had designed for the day after launch — no owner, no monitoring, no written rule for what the automation was allowed to decide on its own.":
    "问题从来不是缺一张卡。而是没有人为上线之后设计完整链路——没有负责人、没有监控、没有 Token 计量与 SLA。",
  "So we built the company around the boring half. We ship the documentation, the alerting and the handover session as part of the build, not as an upsell. It is why our work is still running years later, and why most of our projects arrive by referral.":
    "因此 CacOps 把「可运营」作为默认交付：文档、监控、计费与交接培训都是方案一部分，而不是加价项。这也是为什么多数客户来自转介绍。",
  "Principles": "原则",
  "How we work": "我们的工作方式",
  "Scope narrow, finish": "范围收窄，完成交付",
  "One workflow owned end to end beats six half-built. We would rather ship less and have it survive contact with your team.":
    "一条打通的智算链路，胜过六个半成品。我们宁愿少交付，也要让它经得起真实业务考验。",
  "Design for the handover": "为交接而设计",
  "Documentation, alerting and a named owner ship with the build. If it only works while we are in the room, it does not count.":
    "文档、监控告警与指定负责人随交付一并移交。如果只有我们在场才能运转，那就不算数。",
  "Measure before building": "先评估，再建设",
  "Every engagement opens with an audit. We will tell you when a process should be deleted rather than automated.":
    "每次合作都从智算评估开始。我们会告诉你该先补数据、扩算力，还是直接上推理与 Token。",
  "No lock-in, ever": "绝不锁定",
  "We build in tools you already own and hand over full access. Leaving should cost you nothing but a goodbye.":
    "我们兼容你现有云与工具栈，并移交完整访问权限。离开不应有额外成本。",
  "Team": "团队",
  "The people you’ll work with": "你将合作的人",
  "Founder": "创始人",
  "Systems Lead": "智算架构负责人",
  "Head of Operations": "交付运营负责人",
  "Automation Engineer": "模型服务工程师",
  "Teams automated": "已服务团队",
  "Hours saved monthly": "每月节省成本",
  "Faster response times": "更快上线响应",
  "Average days to launch": "平均上线天数",
  "Transparent pricing for automation sprints, full system builds and ongoing partnership.":
    "算力起步、全栈智算搭建与持续合作的透明定价。",
  "Pricing that scales with you": "随业务成长的弹性定价",
  "Start with one workflow or hand us the whole operation. Every engagement begins with a free audit — no retainer required.":
    "可以从一类能力起步，也可以一次性打通全栈。每次合作都从免费评估开始——无需预付顾问费。",
  "Delivered in 2–3 weeks": "2–3 周交付",
  "Launch Sprint": "算力起步包",
  "One automation built, tested, and deployed to solve a bottleneck before scaling further.":
    "先交付一类核心能力（算力 / 数据 / 推理），验证价值后再扩展全栈。",
  "One Workflow · One Payment · Done": "一个能力包 · 一次付费 · 完成交付",
  "Full automation roadmap audit": "完整智算路线图评估",
  "1 workflow built in Make or n8n": "落地 1 条核心智算能力链路",
  "Up to 4 tool integrations": "最多对接 4 类系统 / 工具",
  "Live testing with your real data": "使用真实业务数据联调",
  "Handoff + documentation": "交付 + 完整文档",
  "14-day support post-launch": "上线后 14 天支持",
  "Live in under 30 days": "30 天内上线",
  "Full System Build": "全栈智算搭建",
  "Your entire ops stack mapped, automated, and ready to scale without the patchwork.":
    "打通算力、数据、训练推理与 Token 计费，形成可扩展的一站式智算体系。",
  "Full Scope · Fixed Price · No Surprises": "全范围 · 固定价格 · 无隐藏费用",
  "Book a scoping call": "预约需求沟通",
  "Everything in Launch Sprint": "包含算力起步包全部内容",
  "Up to 5 workflows built": "最多落地 5 条智算能力链路",
  "Unlimited tool integrations": "不限系统对接数量",
  "Custom AI model setup": "定制模型配置",
  "Team training session": "团队培训",
  "30-day post-launch monitoring": "上线后 30 天监控",
  "Ongoing · Retainer": "持续合作 · 顾问制",
  "Automation Partner": "智算合作伙伴",
  "For teams needing ongoing optimisation, new workflows, and a partner who knows your stack inside out.":
    "适合需要持续扩容、新模型上线，以及长期智算运营支持的团队。",
  "Min. 3-month commitment · cancel anytime after": "最少 3 个月合作 · 之后可随时取消",
  "Book a call": "预约咨询",
  "Up to 2 new workflows/month": "每月最多新增 2 条能力链路",
  "Ongoing fixes & optimisations": "持续优化与运维",
  "Ongoing fixes &amp; optimisations": "持续优化与运维",
  "Dedicated Slack channel": "专属协作通道",
  "Monthly strategy call": "每月策略会议",
  "Priority turnaround": "优先交付",
  "Pause or cancel anytime": "随时暂停或取消",
  "Compare": "对比",
  "What’s included in each": "各方案包含内容",
  "Feature": "功能",
  "Discovery workshop": "需求发现工作坊",
  "Workflows included": "包含能力链路",
  "Up to 5": "最多 5 个",
  "2 / month": "每月 2 个",
  "Typical delivery": "典型交付周期",
  "2–3 weeks": "2–3 周",
  "Under 30 days": "30 天内",
  "Ongoing": "持续",
  "Custom integrations": "定制集成",
  "Monitoring & alerting": "监控与告警",
  "Questions about pricing": "关于定价的问题",
  "Contact — Scalar": "联系我们 — CacOps",
  "Tell us about the process you want automated.": "告诉我们你希望落地的智算场景。",
  "Let's talk about your automation goals": "聊聊你的智算目标",
  "Tell us what's slowing you down": "告诉我们算力、数据或模型卡在哪里",
  "Work — Scalar": "案例 — CacOps",
  "Case studies": "客户案例",
  "Real results from real teams": "来自真实客户的落地成果",
  "Blog — Scalar": "博客 — CacOps",
  "Insights & updates": "洞察与动态",
  "Insights &amp; updates": "洞察与动态",
  "Latest from the Scalar team": "CacOps 团队最新分享",
  "Challenge": "挑战",
  "Approach": "方案",
  "Outcome": "成果",
  "Client": "客户",
  "Industry": "行业",
  "Engagement": "合作方式",
};

function polishZh(s) {
  let out = s.replaceAll("Scalar", "CacOps");
  const reps = [
    ["AI 自动化机构", "智算服务平台"],
    ["自动化系统", "智算体系"],
    ["自动化流程", "智算能力"],
    ["自动化冲刺", "算力起步"],
    ["自动化合作伙伴", "智算合作伙伴"],
    ["自动化路线图", "智算路线图"],
    ["自动化项目", "智算项目"],
    ["自动化机会", "智算落地机会"],
    ["自动化什么", "提供哪些服务"],
    ["工作流", "能力链路"],
    ["Make 或 n8n", "算力与模型服务"],
  ];
  for (const [a, b] of reps) out = out.split(a).join(b);
  return out;
}

function apply(obj) {
  if (Array.isArray(obj)) return obj.map(apply);
  if (obj && typeof obj === "object") {
    const out = {};
    for (const [k, v] of Object.entries(obj)) {
      if (typeof v === "string") {
        out[k] = Object.prototype.hasOwnProperty.call(OVERRIDES, k)
          ? OVERRIDES[k]
          : polishZh(v);
      } else {
        out[k] = apply(v);
      }
    }
    return out;
  }
  return obj;
}

for (const name of ["pages-translation.json", "pages-extra.json"]) {
  const file = path.join(LOCALES, name);
  const data = JSON.parse(fs.readFileSync(file, "utf8"));
  fs.writeFileSync(file, JSON.stringify(apply(data), null, 2) + "\n");
  console.log("updated", name);
}

const pool = {};
function add(map) {
  if (!map || typeof map !== "object") return;
  for (const [k, v] of Object.entries(map)) {
    if (typeof v === "string" && k && v && k !== v && k.length >= 2) pool[k] = v;
  }
}

for (const name of ["zh.json", "pages-translation.json", "pages-extra.json"]) {
  const data = JSON.parse(fs.readFileSync(path.join(LOCALES, name), "utf8"));
  add(data.global);
  add(data.global_extra);
  add(data.letterSpans);
  for (const [k, v] of Object.entries(data)) {
    if (k.startsWith("pages/") && typeof v === "object") add(v);
  }
}

Object.assign(pool, {
  Solutions: "解决方案",
  Pricing: "定价",
  Blog: "博客",
  About: "关于我们",
  "Get Started": "立即开始",
  "Book your free audit": "预约免费评估",
  "Trusted by": "合作伙伴",
  "See how it works": "了解服务体系",
  "Start scaling with AI-powered systems": "一站式智算服务平台",
  "Scale your growth, not your workload.": "用智算放大业务，而不是堆砌成本",
});
for (const bad of ["Work", "All", "FAQ", "Legal"]) delete pool[bad];

fs.writeFileSync(
  path.join(LOCALES, "zh-runtime.json"),
  JSON.stringify(pool, null, 2) + "\n"
);
console.log("runtime entries", Object.keys(pool).length);
