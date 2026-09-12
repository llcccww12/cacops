"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["monaco"],{

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/abap/abap.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/abap/abap.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/abap/abap.ts
var conf = {
  comments: {
    lineComment: "*"
  },
  brackets: [
    ["[", "]"],
    ["(", ")"]
  ]
};
var language = {
  defaultToken: "invalid",
  ignoreCase: true,
  tokenPostfix: ".abap",
  keywords: [
    "abap-source",
    "abbreviated",
    "abstract",
    "accept",
    "accepting",
    "according",
    "activation",
    "actual",
    "add",
    "add-corresponding",
    "adjacent",
    "after",
    "alias",
    "aliases",
    "align",
    "all",
    "allocate",
    "alpha",
    "analysis",
    "analyzer",
    "and",
    "append",
    "appendage",
    "appending",
    "application",
    "archive",
    "area",
    "arithmetic",
    "as",
    "ascending",
    "aspect",
    "assert",
    "assign",
    "assigned",
    "assigning",
    "association",
    "asynchronous",
    "at",
    "attributes",
    "authority",
    "authority-check",
    "avg",
    "back",
    "background",
    "backup",
    "backward",
    "badi",
    "base",
    "before",
    "begin",
    "between",
    "big",
    "binary",
    "bintohex",
    "bit",
    "black",
    "blank",
    "blanks",
    "blob",
    "block",
    "blocks",
    "blue",
    "bound",
    "boundaries",
    "bounds",
    "boxed",
    "break-point",
    "buffer",
    "by",
    "bypassing",
    "byte",
    "byte-order",
    "call",
    "calling",
    "case",
    "cast",
    "casting",
    "catch",
    "center",
    "centered",
    "chain",
    "chain-input",
    "chain-request",
    "change",
    "changing",
    "channels",
    "character",
    "char-to-hex",
    "check",
    "checkbox",
    "ci_",
    "circular",
    "class",
    "class-coding",
    "class-data",
    "class-events",
    "class-methods",
    "class-pool",
    "cleanup",
    "clear",
    "client",
    "clob",
    "clock",
    "close",
    "coalesce",
    "code",
    "coding",
    "col_background",
    "col_group",
    "col_heading",
    "col_key",
    "col_negative",
    "col_normal",
    "col_positive",
    "col_total",
    "collect",
    "color",
    "column",
    "columns",
    "comment",
    "comments",
    "commit",
    "common",
    "communication",
    "comparing",
    "component",
    "components",
    "compression",
    "compute",
    "concat",
    "concat_with_space",
    "concatenate",
    "cond",
    "condense",
    "condition",
    "connect",
    "connection",
    "constants",
    "context",
    "contexts",
    "continue",
    "control",
    "controls",
    "conv",
    "conversion",
    "convert",
    "copies",
    "copy",
    "corresponding",
    "country",
    "cover",
    "cpi",
    "create",
    "creating",
    "critical",
    "currency",
    "currency_conversion",
    "current",
    "cursor",
    "cursor-selection",
    "customer",
    "customer-function",
    "dangerous",
    "data",
    "database",
    "datainfo",
    "dataset",
    "date",
    "dats_add_days",
    "dats_add_months",
    "dats_days_between",
    "dats_is_valid",
    "daylight",
    "dd/mm/yy",
    "dd/mm/yyyy",
    "ddmmyy",
    "deallocate",
    "decimal_shift",
    "decimals",
    "declarations",
    "deep",
    "default",
    "deferred",
    "define",
    "defining",
    "definition",
    "delete",
    "deleting",
    "demand",
    "department",
    "descending",
    "describe",
    "destination",
    "detail",
    "dialog",
    "directory",
    "disconnect",
    "display",
    "display-mode",
    "distinct",
    "divide",
    "divide-corresponding",
    "division",
    "do",
    "dummy",
    "duplicate",
    "duplicates",
    "duration",
    "during",
    "dynamic",
    "dynpro",
    "edit",
    "editor-call",
    "else",
    "elseif",
    "empty",
    "enabled",
    "enabling",
    "encoding",
    "end",
    "endat",
    "endcase",
    "endcatch",
    "endchain",
    "endclass",
    "enddo",
    "endenhancement",
    "end-enhancement-section",
    "endexec",
    "endform",
    "endfunction",
    "endian",
    "endif",
    "ending",
    "endinterface",
    "end-lines",
    "endloop",
    "endmethod",
    "endmodule",
    "end-of-definition",
    "end-of-editing",
    "end-of-file",
    "end-of-page",
    "end-of-selection",
    "endon",
    "endprovide",
    "endselect",
    "end-test-injection",
    "end-test-seam",
    "endtry",
    "endwhile",
    "endwith",
    "engineering",
    "enhancement",
    "enhancement-point",
    "enhancements",
    "enhancement-section",
    "entries",
    "entry",
    "enum",
    "environment",
    "equiv",
    "errormessage",
    "errors",
    "escaping",
    "event",
    "events",
    "exact",
    "except",
    "exception",
    "exceptions",
    "exception-table",
    "exclude",
    "excluding",
    "exec",
    "execute",
    "exists",
    "exit",
    "exit-command",
    "expand",
    "expanding",
    "expiration",
    "explicit",
    "exponent",
    "export",
    "exporting",
    "extend",
    "extended",
    "extension",
    "extract",
    "fail",
    "fetch",
    "field",
    "field-groups",
    "fields",
    "field-symbol",
    "field-symbols",
    "file",
    "filter",
    "filters",
    "filter-table",
    "final",
    "find",
    "first",
    "first-line",
    "fixed-point",
    "fkeq",
    "fkge",
    "flush",
    "font",
    "for",
    "form",
    "format",
    "forward",
    "found",
    "frame",
    "frames",
    "free",
    "friends",
    "from",
    "function",
    "functionality",
    "function-pool",
    "further",
    "gaps",
    "generate",
    "get",
    "giving",
    "gkeq",
    "gkge",
    "global",
    "grant",
    "green",
    "group",
    "groups",
    "handle",
    "handler",
    "harmless",
    "hashed",
    "having",
    "hdb",
    "header",
    "headers",
    "heading",
    "head-lines",
    "help-id",
    "help-request",
    "hextobin",
    "hide",
    "high",
    "hint",
    "hold",
    "hotspot",
    "icon",
    "id",
    "identification",
    "identifier",
    "ids",
    "if",
    "ignore",
    "ignoring",
    "immediately",
    "implementation",
    "implementations",
    "implemented",
    "implicit",
    "import",
    "importing",
    "in",
    "inactive",
    "incl",
    "include",
    "includes",
    "including",
    "increment",
    "index",
    "index-line",
    "infotypes",
    "inheriting",
    "init",
    "initial",
    "initialization",
    "inner",
    "inout",
    "input",
    "insert",
    "instance",
    "instances",
    "instr",
    "intensified",
    "interface",
    "interface-pool",
    "interfaces",
    "internal",
    "intervals",
    "into",
    "inverse",
    "inverted-date",
    "is",
    "iso",
    "job",
    "join",
    "keep",
    "keeping",
    "kernel",
    "key",
    "keys",
    "keywords",
    "kind",
    "language",
    "last",
    "late",
    "layout",
    "leading",
    "leave",
    "left",
    "left-justified",
    "leftplus",
    "leftspace",
    "legacy",
    "length",
    "let",
    "level",
    "levels",
    "like",
    "line",
    "lines",
    "line-count",
    "linefeed",
    "line-selection",
    "line-size",
    "list",
    "listbox",
    "list-processing",
    "little",
    "llang",
    "load",
    "load-of-program",
    "lob",
    "local",
    "locale",
    "locator",
    "logfile",
    "logical",
    "log-point",
    "long",
    "loop",
    "low",
    "lower",
    "lpad",
    "lpi",
    "ltrim",
    "mail",
    "main",
    "major-id",
    "mapping",
    "margin",
    "mark",
    "mask",
    "match",
    "matchcode",
    "max",
    "maximum",
    "medium",
    "members",
    "memory",
    "mesh",
    "message",
    "message-id",
    "messages",
    "messaging",
    "method",
    "methods",
    "min",
    "minimum",
    "minor-id",
    "mm/dd/yy",
    "mm/dd/yyyy",
    "mmddyy",
    "mode",
    "modif",
    "modifier",
    "modify",
    "module",
    "move",
    "move-corresponding",
    "multiply",
    "multiply-corresponding",
    "name",
    "nametab",
    "native",
    "nested",
    "nesting",
    "new",
    "new-line",
    "new-page",
    "new-section",
    "next",
    "no",
    "no-display",
    "no-extension",
    "no-gap",
    "no-gaps",
    "no-grouping",
    "no-heading",
    "no-scrolling",
    "no-sign",
    "no-title",
    "no-topofpage",
    "no-zero",
    "node",
    "nodes",
    "non-unicode",
    "non-unique",
    "not",
    "null",
    "number",
    "object",
    "objects",
    "obligatory",
    "occurrence",
    "occurrences",
    "occurs",
    "of",
    "off",
    "offset",
    "ole",
    "on",
    "only",
    "open",
    "option",
    "optional",
    "options",
    "or",
    "order",
    "other",
    "others",
    "out",
    "outer",
    "output",
    "output-length",
    "overflow",
    "overlay",
    "pack",
    "package",
    "pad",
    "padding",
    "page",
    "pages",
    "parameter",
    "parameters",
    "parameter-table",
    "part",
    "partially",
    "pattern",
    "percentage",
    "perform",
    "performing",
    "person",
    "pf1",
    "pf10",
    "pf11",
    "pf12",
    "pf13",
    "pf14",
    "pf15",
    "pf2",
    "pf3",
    "pf4",
    "pf5",
    "pf6",
    "pf7",
    "pf8",
    "pf9",
    "pf-status",
    "pink",
    "places",
    "pool",
    "pos_high",
    "pos_low",
    "position",
    "pragmas",
    "precompiled",
    "preferred",
    "preserving",
    "primary",
    "print",
    "print-control",
    "priority",
    "private",
    "procedure",
    "process",
    "program",
    "property",
    "protected",
    "provide",
    "public",
    "push",
    "pushbutton",
    "put",
    "queue-only",
    "quickinfo",
    "radiobutton",
    "raise",
    "raising",
    "range",
    "ranges",
    "read",
    "reader",
    "read-only",
    "receive",
    "received",
    "receiver",
    "receiving",
    "red",
    "redefinition",
    "reduce",
    "reduced",
    "ref",
    "reference",
    "refresh",
    "regex",
    "reject",
    "remote",
    "renaming",
    "replace",
    "replacement",
    "replacing",
    "report",
    "request",
    "requested",
    "reserve",
    "reset",
    "resolution",
    "respecting",
    "responsible",
    "result",
    "results",
    "resumable",
    "resume",
    "retry",
    "return",
    "returncode",
    "returning",
    "returns",
    "right",
    "right-justified",
    "rightplus",
    "rightspace",
    "risk",
    "rmc_communication_failure",
    "rmc_invalid_status",
    "rmc_system_failure",
    "role",
    "rollback",
    "rows",
    "rpad",
    "rtrim",
    "run",
    "sap",
    "sap-spool",
    "saving",
    "scale_preserving",
    "scale_preserving_scientific",
    "scan",
    "scientific",
    "scientific_with_leading_zero",
    "scroll",
    "scroll-boundary",
    "scrolling",
    "search",
    "secondary",
    "seconds",
    "section",
    "select",
    "selection",
    "selections",
    "selection-screen",
    "selection-set",
    "selection-sets",
    "selection-table",
    "select-options",
    "send",
    "separate",
    "separated",
    "set",
    "shared",
    "shift",
    "short",
    "shortdump-id",
    "sign_as_postfix",
    "single",
    "size",
    "skip",
    "skipping",
    "smart",
    "some",
    "sort",
    "sortable",
    "sorted",
    "source",
    "specified",
    "split",
    "spool",
    "spots",
    "sql",
    "sqlscript",
    "stable",
    "stamp",
    "standard",
    "starting",
    "start-of-editing",
    "start-of-selection",
    "state",
    "statement",
    "statements",
    "static",
    "statics",
    "statusinfo",
    "step-loop",
    "stop",
    "structure",
    "structures",
    "style",
    "subkey",
    "submatches",
    "submit",
    "subroutine",
    "subscreen",
    "subtract",
    "subtract-corresponding",
    "suffix",
    "sum",
    "summary",
    "summing",
    "supplied",
    "supply",
    "suppress",
    "switch",
    "switchstates",
    "symbol",
    "syncpoints",
    "syntax",
    "syntax-check",
    "syntax-trace",
    "system-call",
    "system-exceptions",
    "system-exit",
    "tab",
    "tabbed",
    "table",
    "tables",
    "tableview",
    "tabstrip",
    "target",
    "task",
    "tasks",
    "test",
    "testing",
    "test-injection",
    "test-seam",
    "text",
    "textpool",
    "then",
    "throw",
    "time",
    "times",
    "timestamp",
    "timezone",
    "tims_is_valid",
    "title",
    "titlebar",
    "title-lines",
    "to",
    "tokenization",
    "tokens",
    "top-lines",
    "top-of-page",
    "trace-file",
    "trace-table",
    "trailing",
    "transaction",
    "transfer",
    "transformation",
    "translate",
    "transporting",
    "trmac",
    "truncate",
    "truncation",
    "try",
    "tstmp_add_seconds",
    "tstmp_current_utctimestamp",
    "tstmp_is_valid",
    "tstmp_seconds_between",
    "type",
    "type-pool",
    "type-pools",
    "types",
    "uline",
    "unassign",
    "under",
    "unicode",
    "union",
    "unique",
    "unit_conversion",
    "unix",
    "unpack",
    "until",
    "unwind",
    "up",
    "update",
    "upper",
    "user",
    "user-command",
    "using",
    "utf-8",
    "valid",
    "value",
    "value-request",
    "values",
    "vary",
    "varying",
    "verification-message",
    "version",
    "via",
    "view",
    "visible",
    "wait",
    "warning",
    "when",
    "whenever",
    "where",
    "while",
    "width",
    "window",
    "windows",
    "with",
    "with-heading",
    "without",
    "with-title",
    "word",
    "work",
    "write",
    "writer",
    "xml",
    "xsd",
    "yellow",
    "yes",
    "yymmdd",
    "zero",
    "zone",
    "abap_system_timezone",
    "abap_user_timezone",
    "access",
    "action",
    "adabas",
    "adjust_numbers",
    "allow_precision_loss",
    "allowed",
    "amdp",
    "applicationuser",
    "as_geo_json",
    "as400",
    "associations",
    "balance",
    "behavior",
    "breakup",
    "bulk",
    "cds",
    "cds_client",
    "check_before_save",
    "child",
    "clients",
    "corr",
    "corr_spearman",
    "cross",
    "cycles",
    "datn_add_days",
    "datn_add_months",
    "datn_days_between",
    "dats_from_datn",
    "dats_tims_to_tstmp",
    "dats_to_datn",
    "db2",
    "db6",
    "ddl",
    "dense_rank",
    "depth",
    "deterministic",
    "discarding",
    "entities",
    "entity",
    "error",
    "failed",
    "finalize",
    "first_value",
    "fltp_to_dec",
    "following",
    "fractional",
    "full",
    "graph",
    "grouping",
    "hierarchy",
    "hierarchy_ancestors",
    "hierarchy_ancestors_aggregate",
    "hierarchy_descendants",
    "hierarchy_descendants_aggregate",
    "hierarchy_siblings",
    "incremental",
    "indicators",
    "lag",
    "last_value",
    "lead",
    "leaves",
    "like_regexpr",
    "link",
    "locale_sap",
    "lock",
    "locks",
    "many",
    "mapped",
    "matched",
    "measures",
    "median",
    "mssqlnt",
    "multiple",
    "nodetype",
    "ntile",
    "nulls",
    "occurrences_regexpr",
    "one",
    "operations",
    "oracle",
    "orphans",
    "over",
    "parent",
    "parents",
    "partition",
    "pcre",
    "period",
    "pfcg_mapping",
    "preceding",
    "privileged",
    "product",
    "projection",
    "rank",
    "redirected",
    "replace_regexpr",
    "reported",
    "response",
    "responses",
    "root",
    "row",
    "row_number",
    "sap_system_date",
    "save",
    "schema",
    "session",
    "sets",
    "shortdump",
    "siblings",
    "spantree",
    "start",
    "stddev",
    "string_agg",
    "subtotal",
    "sybase",
    "tims_from_timn",
    "tims_to_timn",
    "to_blob",
    "to_clob",
    "total",
    "trace-entry",
    "tstmp_to_dats",
    "tstmp_to_dst",
    "tstmp_to_tims",
    "tstmpl_from_utcl",
    "tstmpl_to_utcl",
    "unbounded",
    "utcl_add_seconds",
    "utcl_current",
    "utcl_seconds_between",
    "uuid",
    "var",
    "verbatim"
  ],
  builtinFunctions: [
    "abs",
    "acos",
    "asin",
    "atan",
    "bit-set",
    "boolc",
    "boolx",
    "ceil",
    "char_off",
    "charlen",
    "cmax",
    "cmin",
    "concat_lines_of",
    "contains",
    "contains_any_not_of",
    "contains_any_of",
    "cos",
    "cosh",
    "count",
    "count_any_not_of",
    "count_any_of",
    "dbmaxlen",
    "distance",
    "escape",
    "exp",
    "find_any_not_of",
    "find_any_of",
    "find_end",
    "floor",
    "frac",
    "from_mixed",
    "ipow",
    "line_exists",
    "line_index",
    "log",
    "log10",
    "matches",
    "nmax",
    "nmin",
    "numofchar",
    "repeat",
    "rescale",
    "reverse",
    "round",
    "segment",
    "shift_left",
    "shift_right",
    "sign",
    "sin",
    "sinh",
    "sqrt",
    "strlen",
    "substring",
    "substring_after",
    "substring_before",
    "substring_from",
    "substring_to",
    "tan",
    "tanh",
    "to_lower",
    "to_mixed",
    "to_upper",
    "trunc",
    "utclong_add",
    "utclong_current",
    "utclong_diff",
    "xsdbool",
    "xstrlen"
  ],
  typeKeywords: [
    "b",
    "c",
    "d",
    "decfloat16",
    "decfloat34",
    "f",
    "i",
    "int8",
    "n",
    "p",
    "s",
    "string",
    "t",
    "utclong",
    "x",
    "xstring",
    "any",
    "clike",
    "csequence",
    "decfloat",
    "numeric",
    "simple",
    "xsequence",
    "accp",
    "char",
    "clnt",
    "cuky",
    "curr",
    "datn",
    "dats",
    "d16d",
    "d16n",
    "d16r",
    "d34d",
    "d34n",
    "d34r",
    "dec",
    "df16_dec",
    "df16_raw",
    "df34_dec",
    "df34_raw",
    "fltp",
    "geom_ewkb",
    "int1",
    "int2",
    "int4",
    "lang",
    "lchr",
    "lraw",
    "numc",
    "quan",
    "raw",
    "rawstring",
    "sstring",
    "timn",
    "tims",
    "unit",
    "utcl",
    "df16_scl",
    "df34_scl",
    "prec",
    "varc",
    "abap_bool",
    "abap_false",
    "abap_true",
    "abap_undefined",
    "me",
    "screen",
    "space",
    "super",
    "sy",
    "syst",
    "table_line",
    "*sys*"
  ],
  builtinMethods: ["class_constructor", "constructor"],
  derivedTypes: [
    "%CID",
    "%CID_REF",
    "%CONTROL",
    "%DATA",
    "%ELEMENT",
    "%FAIL",
    "%KEY",
    "%MSG",
    "%PARAM",
    "%PID",
    "%PID_ASSOC",
    "%PID_PARENT",
    "%_HINTS"
  ],
  cdsLanguage: [
    "@AbapAnnotation",
    "@AbapCatalog",
    "@AccessControl",
    "@API",
    "@ClientDependent",
    "@ClientHandling",
    "@CompatibilityContract",
    "@DataAging",
    "@EndUserText",
    "@Environment",
    "@LanguageDependency",
    "@MappingRole",
    "@Metadata",
    "@MetadataExtension",
    "@ObjectModel",
    "@Scope",
    "@Semantics",
    "$EXTENSION",
    "$SELF"
  ],
  selectors: ["->", "->*", "=>", "~", "~*"],
  operators: [
    " +",
    " -",
    "/",
    "*",
    "**",
    "div",
    "mod",
    "=",
    "#",
    "@",
    "+=",
    "-=",
    "*=",
    "/=",
    "**=",
    "&&=",
    "?=",
    "&",
    "&&",
    "bit-and",
    "bit-not",
    "bit-or",
    "bit-xor",
    "m",
    "o",
    "z",
    "<",
    " >",
    "<=",
    ">=",
    "<>",
    "><",
    "=<",
    "=>",
    "bt",
    "byte-ca",
    "byte-cn",
    "byte-co",
    "byte-cs",
    "byte-na",
    "byte-ns",
    "ca",
    "cn",
    "co",
    "cp",
    "cs",
    "eq",
    "ge",
    "gt",
    "le",
    "lt",
    "na",
    "nb",
    "ne",
    "np",
    "ns",
    "*/",
    "*:",
    "--",
    "/*",
    "//"
  ],
  symbols: /[=><!~?&+\-*\/\^%#@]+/,
  tokenizer: {
    root: [
      [
        /[a-z_\/$%@]([\w\/$%]|-(?!>))*/,
        {
          cases: {
            "@typeKeywords": "type",
            "@keywords": "keyword",
            "@cdsLanguage": "annotation",
            "@derivedTypes": "type",
            "@builtinFunctions": "type",
            "@builtinMethods": "type",
            "@operators": "key",
            "@default": "identifier"
          }
        }
      ],
      [/<[\w]+>/, "identifier"],
      [/##[\w|_]+/, "comment"],
      { include: "@whitespace" },
      [/[:,.]/, "delimiter"],
      [/[{}()\[\]]/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@selectors": "tag",
            "@operators": "key",
            "@default": ""
          }
        }
      ],
      [/'/, { token: "string", bracket: "@open", next: "@stringquote" }],
      [/`/, { token: "string", bracket: "@open", next: "@stringping" }],
      [/\|/, { token: "string", bracket: "@open", next: "@stringtemplate" }],
      [/\d+/, "number"]
    ],
    stringtemplate: [
      [/[^\\\|]+/, "string"],
      [/\\\|/, "string"],
      [/\|/, { token: "string", bracket: "@close", next: "@pop" }]
    ],
    stringping: [
      [/[^\\`]+/, "string"],
      [/`/, { token: "string", bracket: "@close", next: "@pop" }]
    ],
    stringquote: [
      [/[^\\']+/, "string"],
      [/'/, { token: "string", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/^\*.*$/, "comment"],
      [/\".*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/apex/apex.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/apex/apex.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/apex/apex.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\#\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "<", close: ">" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*//\\s*(?:(?:#?region\\b)|(?:<editor-fold\\b))"),
      end: new RegExp("^\\s*//\\s*(?:(?:#?endregion\\b)|(?:</editor-fold>))")
    }
  }
};
var keywords = [
  "abstract",
  "activate",
  "and",
  "any",
  "array",
  "as",
  "asc",
  "assert",
  "autonomous",
  "begin",
  "bigdecimal",
  "blob",
  "boolean",
  "break",
  "bulk",
  "by",
  "case",
  "cast",
  "catch",
  "char",
  "class",
  "collect",
  "commit",
  "const",
  "continue",
  "convertcurrency",
  "decimal",
  "default",
  "delete",
  "desc",
  "do",
  "double",
  "else",
  "end",
  "enum",
  "exception",
  "exit",
  "export",
  "extends",
  "false",
  "final",
  "finally",
  "float",
  "for",
  "from",
  "future",
  "get",
  "global",
  "goto",
  "group",
  "having",
  "hint",
  "if",
  "implements",
  "import",
  "in",
  "inner",
  "insert",
  "instanceof",
  "int",
  "interface",
  "into",
  "join",
  "last_90_days",
  "last_month",
  "last_n_days",
  "last_week",
  "like",
  "limit",
  "list",
  "long",
  "loop",
  "map",
  "merge",
  "native",
  "new",
  "next_90_days",
  "next_month",
  "next_n_days",
  "next_week",
  "not",
  "null",
  "nulls",
  "number",
  "object",
  "of",
  "on",
  "or",
  "outer",
  "override",
  "package",
  "parallel",
  "pragma",
  "private",
  "protected",
  "public",
  "retrieve",
  "return",
  "returning",
  "rollback",
  "savepoint",
  "search",
  "select",
  "set",
  "short",
  "sort",
  "stat",
  "static",
  "strictfp",
  "super",
  "switch",
  "synchronized",
  "system",
  "testmethod",
  "then",
  "this",
  "this_month",
  "this_week",
  "throw",
  "throws",
  "today",
  "tolabel",
  "tomorrow",
  "transaction",
  "transient",
  "trigger",
  "true",
  "try",
  "type",
  "undelete",
  "update",
  "upsert",
  "using",
  "virtual",
  "void",
  "volatile",
  "webservice",
  "when",
  "where",
  "while",
  "yesterday"
];
var uppercaseFirstLetter = (lowercase) => lowercase.charAt(0).toUpperCase() + lowercase.substr(1);
var keywordsWithCaseVariations = [];
keywords.forEach((lowercase) => {
  keywordsWithCaseVariations.push(lowercase);
  keywordsWithCaseVariations.push(lowercase.toUpperCase());
  keywordsWithCaseVariations.push(uppercaseFirstLetter(lowercase));
});
var language = {
  defaultToken: "",
  tokenPostfix: ".apex",
  keywords: keywordsWithCaseVariations,
  operators: [
    "=",
    ">",
    "<",
    "!",
    "~",
    "?",
    ":",
    "==",
    "<=",
    ">=",
    "!=",
    "&&",
    "||",
    "++",
    "--",
    "+",
    "-",
    "*",
    "/",
    "&",
    "|",
    "^",
    "%",
    "<<",
    ">>",
    ">>>",
    "+=",
    "-=",
    "*=",
    "/=",
    "&=",
    "|=",
    "^=",
    "%=",
    "<<=",
    ">>=",
    ">>>="
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  digits: /\d+(_+\d+)*/,
  octaldigits: /[0-7]+(_+[0-7]+)*/,
  binarydigits: /[0-1]+(_+[0-1]+)*/,
  hexdigits: /[[0-9a-fA-F]+(_+[0-9a-fA-F]+)*/,
  tokenizer: {
    root: [
      [
        /[a-z_$][\w$]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      [
        /[A-Z][\w\$]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "type.identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/@\s*[a-zA-Z_\$][\w\$]*/, "annotation"],
      [/(@digits)[eE]([\-+]?(@digits))?[fFdD]?/, "number.float"],
      [/(@digits)\.(@digits)([eE][\-+]?(@digits))?[fFdD]?/, "number.float"],
      [/(@digits)[fFdD]/, "number.float"],
      [/(@digits)[lL]?/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", '@string."'],
      [/'/, "string", "@string.'"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@apexdoc"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    apexdoc: [
      [/[^\/*]+/, "comment.doc"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"']+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/azcli/azcli.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/azcli/azcli.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/azcli/azcli.ts
var conf = {
  comments: {
    lineComment: "#"
  }
};
var language = {
  defaultToken: "keyword",
  ignoreCase: true,
  tokenPostfix: ".azcli",
  str: /[^#\s]/,
  tokenizer: {
    root: [
      { include: "@comment" },
      [
        /\s-+@str*\s*/,
        {
          cases: {
            "@eos": { token: "key.identifier", next: "@popall" },
            "@default": { token: "key.identifier", next: "@type" }
          }
        }
      ],
      [
        /^-+@str*\s*/,
        {
          cases: {
            "@eos": { token: "key.identifier", next: "@popall" },
            "@default": { token: "key.identifier", next: "@type" }
          }
        }
      ]
    ],
    type: [
      { include: "@comment" },
      [
        /-+@str*\s*/,
        {
          cases: {
            "@eos": { token: "key.identifier", next: "@popall" },
            "@default": "key.identifier"
          }
        }
      ],
      [
        /@str+\s*/,
        {
          cases: {
            "@eos": { token: "string", next: "@popall" },
            "@default": "string"
          }
        }
      ]
    ],
    comment: [
      [
        /#.*$/,
        {
          cases: {
            "@eos": { token: "comment", next: "@popall" }
          }
        }
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/bat/bat.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/bat/bat.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/bat/bat.ts
var conf = {
  comments: {
    lineComment: "REM"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ],
  surroundingPairs: [
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*(::\\s*|REM\\s+)#region"),
      end: new RegExp("^\\s*(::\\s*|REM\\s+)#endregion")
    }
  }
};
var language = {
  defaultToken: "",
  ignoreCase: true,
  tokenPostfix: ".bat",
  brackets: [
    { token: "delimiter.bracket", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" }
  ],
  keywords: /call|defined|echo|errorlevel|exist|for|goto|if|pause|set|shift|start|title|not|pushd|popd/,
  symbols: /[=><!~?&|+\-*\/\^;\.,]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/^(\s*)(rem(?:\s.*|))$/, ["", "comment"]],
      [/(\@?)(@keywords)(?!\w)/, [{ token: "keyword" }, { token: "keyword.$2" }]],
      [/[ \t\r\n]+/, ""],
      [/setlocal(?!\w)/, "keyword.tag-setlocal"],
      [/endlocal(?!\w)/, "keyword.tag-setlocal"],
      [/[a-zA-Z_]\w*/, ""],
      [/:\w*/, "metatag"],
      [/%[^%]+%/, "variable"],
      [/%%[\w]+(?!\w)/, "variable"],
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F_]*[0-9a-fA-F]/, "number.hex"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"],
      [/"/, "string", '@string."'],
      [/'/, "string", "@string.'"]
    ],
    string: [
      [
        /[^\\"'%]+/,
        {
          cases: {
            "@eos": { token: "string", next: "@popall" },
            "@default": "string"
          }
        }
      ],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/%[\w ]+%/, "variable"],
      [/%%[\w]+(?!\w)/, "variable"],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      [/$/, "string", "@popall"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/bicep/bicep.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/bicep/bicep.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/bicep/bicep.ts
var bounded = (text) => `\\b${text}\\b`;
var identifierStart = "[_a-zA-Z]";
var identifierContinue = "[_a-zA-Z0-9]";
var identifier = bounded(`${identifierStart}${identifierContinue}*`);
var keywords = [
  "targetScope",
  "resource",
  "module",
  "param",
  "var",
  "output",
  "for",
  "in",
  "if",
  "existing"
];
var namedLiterals = ["true", "false", "null"];
var nonCommentWs = `[ \\t\\r\\n]`;
var numericLiteral = `[0-9]+`;
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "'", close: "'" },
    { open: "'''", close: "'''" }
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: "'''", close: "'''", notIn: ["string", "comment"] }
  ],
  autoCloseBefore: ":.,=}])' \n	",
  indentationRules: {
    increaseIndentPattern: new RegExp("^((?!\\/\\/).)*(\\{[^}\"'`]*|\\([^)\"'`]*|\\[[^\\]\"'`]*)$"),
    decreaseIndentPattern: new RegExp("^((?!.*?\\/\\*).*\\*/)?\\s*[\\}\\]].*$")
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".bicep",
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  symbols: /[=><!~?:&|+\-*/^%]+/,
  keywords,
  namedLiterals,
  escapes: `\\\\(u{[0-9A-Fa-f]+}|n|r|t|\\\\|'|\\\${)`,
  tokenizer: {
    root: [{ include: "@expression" }, { include: "@whitespace" }],
    stringVerbatim: [
      { regex: `(|'|'')[^']`, action: { token: "string" } },
      { regex: `'''`, action: { token: "string.quote", next: "@pop" } }
    ],
    stringLiteral: [
      { regex: `\\\${`, action: { token: "delimiter.bracket", next: "@bracketCounting" } },
      { regex: `[^\\\\'$]+`, action: { token: "string" } },
      { regex: "@escapes", action: { token: "string.escape" } },
      { regex: `\\\\.`, action: { token: "string.escape.invalid" } },
      { regex: `'`, action: { token: "string", next: "@pop" } }
    ],
    bracketCounting: [
      { regex: `{`, action: { token: "delimiter.bracket", next: "@bracketCounting" } },
      { regex: `}`, action: { token: "delimiter.bracket", next: "@pop" } },
      { include: "expression" }
    ],
    comment: [
      { regex: `[^\\*]+`, action: { token: "comment" } },
      { regex: `\\*\\/`, action: { token: "comment", next: "@pop" } },
      { regex: `[\\/*]`, action: { token: "comment" } }
    ],
    whitespace: [
      { regex: nonCommentWs },
      { regex: `\\/\\*`, action: { token: "comment", next: "@comment" } },
      { regex: `\\/\\/.*$`, action: { token: "comment" } }
    ],
    expression: [
      { regex: `'''`, action: { token: "string.quote", next: "@stringVerbatim" } },
      { regex: `'`, action: { token: "string.quote", next: "@stringLiteral" } },
      { regex: numericLiteral, action: { token: "number" } },
      {
        regex: identifier,
        action: {
          cases: {
            "@keywords": { token: "keyword" },
            "@namedLiterals": { token: "keyword" },
            "@default": { token: "identifier" }
          }
        }
      }
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/cameligo/cameligo.js":
/*!********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/cameligo/cameligo.js ***!
  \********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/cameligo/cameligo.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["(*", "*)"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: '"', close: '"' },
    { open: "(*", close: "*)" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: '"', close: '"' },
    { open: "(*", close: "*)" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".cameligo",
  ignoreCase: true,
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  keywords: [
    "abs",
    "assert",
    "block",
    "Bytes",
    "case",
    "Crypto",
    "Current",
    "else",
    "failwith",
    "false",
    "for",
    "fun",
    "if",
    "in",
    "let",
    "let%entry",
    "let%init",
    "List",
    "list",
    "Map",
    "map",
    "match",
    "match%nat",
    "mod",
    "not",
    "operation",
    "Operation",
    "of",
    "record",
    "Set",
    "set",
    "sender",
    "skip",
    "source",
    "String",
    "then",
    "to",
    "true",
    "type",
    "with"
  ],
  typeKeywords: ["int", "unit", "string", "tz", "nat", "bool"],
  operators: [
    "=",
    ">",
    "<",
    "<=",
    ">=",
    "<>",
    ":",
    ":=",
    "and",
    "mod",
    "or",
    "+",
    "-",
    "*",
    "/",
    "@",
    "&",
    "^",
    "%",
    "->",
    "<-",
    "&&",
    "||"
  ],
  symbols: /[=><:@\^&|+\-*\/\^%]+/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_][\w]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/\$[0-9a-fA-F]{1,16}/, "number.hex"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/'/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/'/, "string.invalid"],
      [/\#\d+/, "string"]
    ],
    comment: [
      [/[^\(\*]+/, "comment"],
      [/\*\)/, "comment", "@pop"],
      [/\(\*/, "comment"]
    ],
    string: [
      [/[^\\']+/, "string"],
      [/\\./, "string.escape.invalid"],
      [/'/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\(\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/clojure/clojure.js":
/*!******************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/clojure/clojure.js ***!
  \******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/clojure/clojure.ts
var conf = {
  comments: {
    lineComment: ";;"
  },
  brackets: [
    ["[", "]"],
    ["(", ")"],
    ["{", "}"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: '"', close: '"' },
    { open: "(", close: ")" },
    { open: "{", close: "}" }
  ],
  surroundingPairs: [
    { open: "[", close: "]" },
    { open: '"', close: '"' },
    { open: "(", close: ")" },
    { open: "{", close: "}" }
  ]
};
var language = {
  defaultToken: "",
  ignoreCase: true,
  tokenPostfix: ".clj",
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "{", close: "}", token: "delimiter.curly" }
  ],
  constants: ["true", "false", "nil"],
  numbers: /^(?:[+\-]?\d+(?:(?:N|(?:[eE][+\-]?\d+))|(?:\.?\d*(?:M|(?:[eE][+\-]?\d+))?)|\/\d+|[xX][0-9a-fA-F]+|r[0-9a-zA-Z]+)?(?=[\\\[\]\s"#'(),;@^`{}~]|$))/,
  characters: /^(?:\\(?:backspace|formfeed|newline|return|space|tab|o[0-7]{3}|u[0-9A-Fa-f]{4}|x[0-9A-Fa-f]{4}|.)?(?=[\\\[\]\s"(),;@^`{}~]|$))/,
  escapes: /^\\(?:["'\\bfnrt]|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  qualifiedSymbols: /^(?:(?:[^\\\/\[\]\d\s"#'(),;@^`{}~][^\\\[\]\s"(),;@^`{}~]*(?:\.[^\\\/\[\]\d\s"#'(),;@^`{}~][^\\\[\]\s"(),;@^`{}~]*)*\/)?(?:\/|[^\\\/\[\]\d\s"#'(),;@^`{}~][^\\\[\]\s"(),;@^`{}~]*)*(?=[\\\[\]\s"(),;@^`{}~]|$))/,
  specialForms: [
    ".",
    "catch",
    "def",
    "do",
    "if",
    "monitor-enter",
    "monitor-exit",
    "new",
    "quote",
    "recur",
    "set!",
    "throw",
    "try",
    "var"
  ],
  coreSymbols: [
    "*",
    "*'",
    "*1",
    "*2",
    "*3",
    "*agent*",
    "*allow-unresolved-vars*",
    "*assert*",
    "*clojure-version*",
    "*command-line-args*",
    "*compile-files*",
    "*compile-path*",
    "*compiler-options*",
    "*data-readers*",
    "*default-data-reader-fn*",
    "*e",
    "*err*",
    "*file*",
    "*flush-on-newline*",
    "*fn-loader*",
    "*in*",
    "*math-context*",
    "*ns*",
    "*out*",
    "*print-dup*",
    "*print-length*",
    "*print-level*",
    "*print-meta*",
    "*print-namespace-maps*",
    "*print-readably*",
    "*read-eval*",
    "*reader-resolver*",
    "*source-path*",
    "*suppress-read*",
    "*unchecked-math*",
    "*use-context-classloader*",
    "*verbose-defrecords*",
    "*warn-on-reflection*",
    "+",
    "+'",
    "-",
    "-'",
    "->",
    "->>",
    "->ArrayChunk",
    "->Eduction",
    "->Vec",
    "->VecNode",
    "->VecSeq",
    "-cache-protocol-fn",
    "-reset-methods",
    "..",
    "/",
    "<",
    "<=",
    "=",
    "==",
    ">",
    ">=",
    "EMPTY-NODE",
    "Inst",
    "StackTraceElement->vec",
    "Throwable->map",
    "accessor",
    "aclone",
    "add-classpath",
    "add-watch",
    "agent",
    "agent-error",
    "agent-errors",
    "aget",
    "alength",
    "alias",
    "all-ns",
    "alter",
    "alter-meta!",
    "alter-var-root",
    "amap",
    "ancestors",
    "and",
    "any?",
    "apply",
    "areduce",
    "array-map",
    "as->",
    "aset",
    "aset-boolean",
    "aset-byte",
    "aset-char",
    "aset-double",
    "aset-float",
    "aset-int",
    "aset-long",
    "aset-short",
    "assert",
    "assoc",
    "assoc!",
    "assoc-in",
    "associative?",
    "atom",
    "await",
    "await-for",
    "await1",
    "bases",
    "bean",
    "bigdec",
    "bigint",
    "biginteger",
    "binding",
    "bit-and",
    "bit-and-not",
    "bit-clear",
    "bit-flip",
    "bit-not",
    "bit-or",
    "bit-set",
    "bit-shift-left",
    "bit-shift-right",
    "bit-test",
    "bit-xor",
    "boolean",
    "boolean-array",
    "boolean?",
    "booleans",
    "bound-fn",
    "bound-fn*",
    "bound?",
    "bounded-count",
    "butlast",
    "byte",
    "byte-array",
    "bytes",
    "bytes?",
    "case",
    "cast",
    "cat",
    "char",
    "char-array",
    "char-escape-string",
    "char-name-string",
    "char?",
    "chars",
    "chunk",
    "chunk-append",
    "chunk-buffer",
    "chunk-cons",
    "chunk-first",
    "chunk-next",
    "chunk-rest",
    "chunked-seq?",
    "class",
    "class?",
    "clear-agent-errors",
    "clojure-version",
    "coll?",
    "comment",
    "commute",
    "comp",
    "comparator",
    "compare",
    "compare-and-set!",
    "compile",
    "complement",
    "completing",
    "concat",
    "cond",
    "cond->",
    "cond->>",
    "condp",
    "conj",
    "conj!",
    "cons",
    "constantly",
    "construct-proxy",
    "contains?",
    "count",
    "counted?",
    "create-ns",
    "create-struct",
    "cycle",
    "dec",
    "dec'",
    "decimal?",
    "declare",
    "dedupe",
    "default-data-readers",
    "definline",
    "definterface",
    "defmacro",
    "defmethod",
    "defmulti",
    "defn",
    "defn-",
    "defonce",
    "defprotocol",
    "defrecord",
    "defstruct",
    "deftype",
    "delay",
    "delay?",
    "deliver",
    "denominator",
    "deref",
    "derive",
    "descendants",
    "destructure",
    "disj",
    "disj!",
    "dissoc",
    "dissoc!",
    "distinct",
    "distinct?",
    "doall",
    "dorun",
    "doseq",
    "dosync",
    "dotimes",
    "doto",
    "double",
    "double-array",
    "double?",
    "doubles",
    "drop",
    "drop-last",
    "drop-while",
    "eduction",
    "empty",
    "empty?",
    "ensure",
    "ensure-reduced",
    "enumeration-seq",
    "error-handler",
    "error-mode",
    "eval",
    "even?",
    "every-pred",
    "every?",
    "ex-data",
    "ex-info",
    "extend",
    "extend-protocol",
    "extend-type",
    "extenders",
    "extends?",
    "false?",
    "ffirst",
    "file-seq",
    "filter",
    "filterv",
    "find",
    "find-keyword",
    "find-ns",
    "find-protocol-impl",
    "find-protocol-method",
    "find-var",
    "first",
    "flatten",
    "float",
    "float-array",
    "float?",
    "floats",
    "flush",
    "fn",
    "fn?",
    "fnext",
    "fnil",
    "for",
    "force",
    "format",
    "frequencies",
    "future",
    "future-call",
    "future-cancel",
    "future-cancelled?",
    "future-done?",
    "future?",
    "gen-class",
    "gen-interface",
    "gensym",
    "get",
    "get-in",
    "get-method",
    "get-proxy-class",
    "get-thread-bindings",
    "get-validator",
    "group-by",
    "halt-when",
    "hash",
    "hash-combine",
    "hash-map",
    "hash-ordered-coll",
    "hash-set",
    "hash-unordered-coll",
    "ident?",
    "identical?",
    "identity",
    "if-let",
    "if-not",
    "if-some",
    "ifn?",
    "import",
    "in-ns",
    "inc",
    "inc'",
    "indexed?",
    "init-proxy",
    "inst-ms",
    "inst-ms*",
    "inst?",
    "instance?",
    "int",
    "int-array",
    "int?",
    "integer?",
    "interleave",
    "intern",
    "interpose",
    "into",
    "into-array",
    "ints",
    "io!",
    "isa?",
    "iterate",
    "iterator-seq",
    "juxt",
    "keep",
    "keep-indexed",
    "key",
    "keys",
    "keyword",
    "keyword?",
    "last",
    "lazy-cat",
    "lazy-seq",
    "let",
    "letfn",
    "line-seq",
    "list",
    "list*",
    "list?",
    "load",
    "load-file",
    "load-reader",
    "load-string",
    "loaded-libs",
    "locking",
    "long",
    "long-array",
    "longs",
    "loop",
    "macroexpand",
    "macroexpand-1",
    "make-array",
    "make-hierarchy",
    "map",
    "map-entry?",
    "map-indexed",
    "map?",
    "mapcat",
    "mapv",
    "max",
    "max-key",
    "memfn",
    "memoize",
    "merge",
    "merge-with",
    "meta",
    "method-sig",
    "methods",
    "min",
    "min-key",
    "mix-collection-hash",
    "mod",
    "munge",
    "name",
    "namespace",
    "namespace-munge",
    "nat-int?",
    "neg-int?",
    "neg?",
    "newline",
    "next",
    "nfirst",
    "nil?",
    "nnext",
    "not",
    "not-any?",
    "not-empty",
    "not-every?",
    "not=",
    "ns",
    "ns-aliases",
    "ns-imports",
    "ns-interns",
    "ns-map",
    "ns-name",
    "ns-publics",
    "ns-refers",
    "ns-resolve",
    "ns-unalias",
    "ns-unmap",
    "nth",
    "nthnext",
    "nthrest",
    "num",
    "number?",
    "numerator",
    "object-array",
    "odd?",
    "or",
    "parents",
    "partial",
    "partition",
    "partition-all",
    "partition-by",
    "pcalls",
    "peek",
    "persistent!",
    "pmap",
    "pop",
    "pop!",
    "pop-thread-bindings",
    "pos-int?",
    "pos?",
    "pr",
    "pr-str",
    "prefer-method",
    "prefers",
    "primitives-classnames",
    "print",
    "print-ctor",
    "print-dup",
    "print-method",
    "print-simple",
    "print-str",
    "printf",
    "println",
    "println-str",
    "prn",
    "prn-str",
    "promise",
    "proxy",
    "proxy-call-with-super",
    "proxy-mappings",
    "proxy-name",
    "proxy-super",
    "push-thread-bindings",
    "pvalues",
    "qualified-ident?",
    "qualified-keyword?",
    "qualified-symbol?",
    "quot",
    "rand",
    "rand-int",
    "rand-nth",
    "random-sample",
    "range",
    "ratio?",
    "rational?",
    "rationalize",
    "re-find",
    "re-groups",
    "re-matcher",
    "re-matches",
    "re-pattern",
    "re-seq",
    "read",
    "read-line",
    "read-string",
    "reader-conditional",
    "reader-conditional?",
    "realized?",
    "record?",
    "reduce",
    "reduce-kv",
    "reduced",
    "reduced?",
    "reductions",
    "ref",
    "ref-history-count",
    "ref-max-history",
    "ref-min-history",
    "ref-set",
    "refer",
    "refer-clojure",
    "reify",
    "release-pending-sends",
    "rem",
    "remove",
    "remove-all-methods",
    "remove-method",
    "remove-ns",
    "remove-watch",
    "repeat",
    "repeatedly",
    "replace",
    "replicate",
    "require",
    "reset!",
    "reset-meta!",
    "reset-vals!",
    "resolve",
    "rest",
    "restart-agent",
    "resultset-seq",
    "reverse",
    "reversible?",
    "rseq",
    "rsubseq",
    "run!",
    "satisfies?",
    "second",
    "select-keys",
    "send",
    "send-off",
    "send-via",
    "seq",
    "seq?",
    "seqable?",
    "seque",
    "sequence",
    "sequential?",
    "set",
    "set-agent-send-executor!",
    "set-agent-send-off-executor!",
    "set-error-handler!",
    "set-error-mode!",
    "set-validator!",
    "set?",
    "short",
    "short-array",
    "shorts",
    "shuffle",
    "shutdown-agents",
    "simple-ident?",
    "simple-keyword?",
    "simple-symbol?",
    "slurp",
    "some",
    "some->",
    "some->>",
    "some-fn",
    "some?",
    "sort",
    "sort-by",
    "sorted-map",
    "sorted-map-by",
    "sorted-set",
    "sorted-set-by",
    "sorted?",
    "special-symbol?",
    "spit",
    "split-at",
    "split-with",
    "str",
    "string?",
    "struct",
    "struct-map",
    "subs",
    "subseq",
    "subvec",
    "supers",
    "swap!",
    "swap-vals!",
    "symbol",
    "symbol?",
    "sync",
    "tagged-literal",
    "tagged-literal?",
    "take",
    "take-last",
    "take-nth",
    "take-while",
    "test",
    "the-ns",
    "thread-bound?",
    "time",
    "to-array",
    "to-array-2d",
    "trampoline",
    "transduce",
    "transient",
    "tree-seq",
    "true?",
    "type",
    "unchecked-add",
    "unchecked-add-int",
    "unchecked-byte",
    "unchecked-char",
    "unchecked-dec",
    "unchecked-dec-int",
    "unchecked-divide-int",
    "unchecked-double",
    "unchecked-float",
    "unchecked-inc",
    "unchecked-inc-int",
    "unchecked-int",
    "unchecked-long",
    "unchecked-multiply",
    "unchecked-multiply-int",
    "unchecked-negate",
    "unchecked-negate-int",
    "unchecked-remainder-int",
    "unchecked-short",
    "unchecked-subtract",
    "unchecked-subtract-int",
    "underive",
    "unquote",
    "unquote-splicing",
    "unreduced",
    "unsigned-bit-shift-right",
    "update",
    "update-in",
    "update-proxy",
    "uri?",
    "use",
    "uuid?",
    "val",
    "vals",
    "var-get",
    "var-set",
    "var?",
    "vary-meta",
    "vec",
    "vector",
    "vector-of",
    "vector?",
    "volatile!",
    "volatile?",
    "vreset!",
    "vswap!",
    "when",
    "when-first",
    "when-let",
    "when-not",
    "when-some",
    "while",
    "with-bindings",
    "with-bindings*",
    "with-in-str",
    "with-loading-context",
    "with-local-vars",
    "with-meta",
    "with-open",
    "with-out-str",
    "with-precision",
    "with-redefs",
    "with-redefs-fn",
    "xml-seq",
    "zero?",
    "zipmap"
  ],
  tokenizer: {
    root: [
      { include: "@whitespace" },
      [/@numbers/, "number"],
      [/@characters/, "string"],
      { include: "@string" },
      [/[()\[\]{}]/, "@brackets"],
      [/\/#"(?:\.|(?:")|[^"\n])*"\/g/, "regexp"],
      [/[#'@^`~]/, "meta"],
      [
        /@qualifiedSymbols/,
        {
          cases: {
            "^:.+$": "constant",
            "@specialForms": "keyword",
            "@coreSymbols": "keyword",
            "@constants": "constant",
            "@default": "identifier"
          }
        }
      ]
    ],
    whitespace: [
      [/[\s,]+/, "white"],
      [/;.*$/, "comment"],
      [/\(comment\b/, "comment", "@comment"]
    ],
    comment: [
      [/\(/, "comment", "@push"],
      [/\)/, "comment", "@pop"],
      [/[^()]/, "comment"]
    ],
    string: [[/"/, "string", "@multiLineString"]],
    multiLineString: [
      [/"/, "string", "@popall"],
      [/@escapes/, "string.escape"],
      [/./, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/coffee/coffee.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/coffee/coffee.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/coffee/coffee.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\#%\^\&\*\(\)\=\$\-\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    blockComment: ["###", "###"],
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#region\\b"),
      end: new RegExp("^\\s*#endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  ignoreCase: true,
  tokenPostfix: ".coffee",
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  regEx: /\/(?!\/\/)(?:[^\/\\]|\\.)*\/[igm]*/,
  keywords: [
    "and",
    "or",
    "is",
    "isnt",
    "not",
    "on",
    "yes",
    "@",
    "no",
    "off",
    "true",
    "false",
    "null",
    "this",
    "new",
    "delete",
    "typeof",
    "in",
    "instanceof",
    "return",
    "throw",
    "break",
    "continue",
    "debugger",
    "if",
    "else",
    "switch",
    "for",
    "while",
    "do",
    "try",
    "catch",
    "finally",
    "class",
    "extends",
    "super",
    "undefined",
    "then",
    "unless",
    "until",
    "loop",
    "of",
    "by",
    "when"
  ],
  symbols: /[=><!~?&%|+\-*\/\^\.,\:]+/,
  escapes: /\\(?:[abfnrtv\\"'$]|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/\@[a-zA-Z_]\w*/, "variable.predefined"],
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            this: "variable.predefined",
            "@keywords": { token: "keyword.$0" },
            "@default": ""
          }
        }
      ],
      [/[ \t\r\n]+/, ""],
      [/###/, "comment", "@comment"],
      [/#.*$/, "comment"],
      ["///", { token: "regexp", next: "@hereregexp" }],
      [/^(\s*)(@regEx)/, ["", "regexp"]],
      [/(\()(\s*)(@regEx)/, ["@brackets", "", "regexp"]],
      [/(\,)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\=)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\:)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\[)(\s*)(@regEx)/, ["@brackets", "", "regexp"]],
      [/(\!)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\&)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\|)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\?)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\{)(\s*)(@regEx)/, ["@brackets", "", "regexp"]],
      [/(\;)(\s*)(@regEx)/, ["", "", "regexp"]],
      [
        /}/,
        {
          cases: {
            "$S2==interpolatedstring": {
              token: "string",
              next: "@pop"
            },
            "@default": "@brackets"
          }
        }
      ],
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/\d+[eE]([\-+]?\d+)?/, "number.float"],
      [/\d+\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F]+/, "number.hex"],
      [/0[0-7]+(?!\d)/, "number.octal"],
      [/\d+/, "number"],
      [/[,.]/, "delimiter"],
      [/"""/, "string", '@herestring."""'],
      [/'''/, "string", "@herestring.'''"],
      [
        /"/,
        {
          cases: {
            "@eos": "string",
            "@default": { token: "string", next: '@string."' }
          }
        }
      ],
      [
        /'/,
        {
          cases: {
            "@eos": "string",
            "@default": { token: "string", next: "@string.'" }
          }
        }
      ]
    ],
    string: [
      [/[^"'\#\\]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\./, "string.escape.invalid"],
      [/\./, "string.escape.invalid"],
      [
        /#{/,
        {
          cases: {
            '$S2=="': {
              token: "string",
              next: "root.interpolatedstring"
            },
            "@default": "string"
          }
        }
      ],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      [/#/, "string"]
    ],
    herestring: [
      [
        /("""|''')/,
        {
          cases: {
            "$1==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      [/[^#\\'"]+/, "string"],
      [/['"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\./, "string.escape.invalid"],
      [/#{/, { token: "string.quote", next: "root.interpolatedstring" }],
      [/#/, "string"]
    ],
    comment: [
      [/[^#]+/, "comment"],
      [/###/, "comment", "@pop"],
      [/#/, "comment"]
    ],
    hereregexp: [
      [/[^\\\/#]+/, "regexp"],
      [/\\./, "regexp"],
      [/#.*$/, "comment"],
      ["///[igm]*", { token: "regexp", next: "@pop" }],
      [/\//, "regexp"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/cpp/cpp.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/cpp/cpp.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/cpp/cpp.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "(", close: ")" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#pragma\\s+region\\b"),
      end: new RegExp("^\\s*#pragma\\s+endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".cpp",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" },
    { token: "delimiter.angle", open: "<", close: ">" }
  ],
  keywords: [
    "abstract",
    "amp",
    "array",
    "auto",
    "bool",
    "break",
    "case",
    "catch",
    "char",
    "class",
    "const",
    "constexpr",
    "const_cast",
    "continue",
    "cpu",
    "decltype",
    "default",
    "delegate",
    "delete",
    "do",
    "double",
    "dynamic_cast",
    "each",
    "else",
    "enum",
    "event",
    "explicit",
    "export",
    "extern",
    "false",
    "final",
    "finally",
    "float",
    "for",
    "friend",
    "gcnew",
    "generic",
    "goto",
    "if",
    "in",
    "initonly",
    "inline",
    "int",
    "interface",
    "interior_ptr",
    "internal",
    "literal",
    "long",
    "mutable",
    "namespace",
    "new",
    "noexcept",
    "nullptr",
    "__nullptr",
    "operator",
    "override",
    "partial",
    "pascal",
    "pin_ptr",
    "private",
    "property",
    "protected",
    "public",
    "ref",
    "register",
    "reinterpret_cast",
    "restrict",
    "return",
    "safe_cast",
    "sealed",
    "short",
    "signed",
    "sizeof",
    "static",
    "static_assert",
    "static_cast",
    "struct",
    "switch",
    "template",
    "this",
    "thread_local",
    "throw",
    "tile_static",
    "true",
    "try",
    "typedef",
    "typeid",
    "typename",
    "union",
    "unsigned",
    "using",
    "virtual",
    "void",
    "volatile",
    "wchar_t",
    "where",
    "while",
    "_asm",
    "_based",
    "_cdecl",
    "_declspec",
    "_fastcall",
    "_if_exists",
    "_if_not_exists",
    "_inline",
    "_multiple_inheritance",
    "_pascal",
    "_single_inheritance",
    "_stdcall",
    "_virtual_inheritance",
    "_w64",
    "__abstract",
    "__alignof",
    "__asm",
    "__assume",
    "__based",
    "__box",
    "__builtin_alignof",
    "__cdecl",
    "__clrcall",
    "__declspec",
    "__delegate",
    "__event",
    "__except",
    "__fastcall",
    "__finally",
    "__forceinline",
    "__gc",
    "__hook",
    "__identifier",
    "__if_exists",
    "__if_not_exists",
    "__inline",
    "__int128",
    "__int16",
    "__int32",
    "__int64",
    "__int8",
    "__interface",
    "__leave",
    "__m128",
    "__m128d",
    "__m128i",
    "__m256",
    "__m256d",
    "__m256i",
    "__m512",
    "__m512d",
    "__m512i",
    "__m64",
    "__multiple_inheritance",
    "__newslot",
    "__nogc",
    "__noop",
    "__nounwind",
    "__novtordisp",
    "__pascal",
    "__pin",
    "__pragma",
    "__property",
    "__ptr32",
    "__ptr64",
    "__raise",
    "__restrict",
    "__resume",
    "__sealed",
    "__single_inheritance",
    "__stdcall",
    "__super",
    "__thiscall",
    "__try",
    "__try_cast",
    "__typeof",
    "__unaligned",
    "__unhook",
    "__uuidof",
    "__value",
    "__virtual_inheritance",
    "__w64",
    "__wchar_t"
  ],
  operators: [
    "=",
    ">",
    "<",
    "!",
    "~",
    "?",
    ":",
    "==",
    "<=",
    ">=",
    "!=",
    "&&",
    "||",
    "++",
    "--",
    "+",
    "-",
    "*",
    "/",
    "&",
    "|",
    "^",
    "%",
    "<<",
    ">>",
    ">>>",
    "+=",
    "-=",
    "*=",
    "/=",
    "&=",
    "|=",
    "^=",
    "%=",
    "<<=",
    ">>=",
    ">>>="
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[0abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  integersuffix: /([uU](ll|LL|l|L)|(ll|LL|l|L)?[uU]?)/,
  floatsuffix: /[fFlL]?/,
  encoding: /u|u8|U|L/,
  tokenizer: {
    root: [
      [/@encoding?R\"(?:([^ ()\\\t]*))\(/, { token: "string.raw.begin", next: "@raw.$1" }],
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      [/^\s*#\s*include/, { token: "keyword.directive.include", next: "@include" }],
      [/^\s*#\s*\w+/, "keyword.directive"],
      { include: "@whitespace" },
      [/\[\s*\[/, { token: "annotation", next: "@annotation" }],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\d+[eE]([\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/\d*\.\d+([eE][\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/0[xX][0-9a-fA-F']*[0-9a-fA-F](@integersuffix)/, "number.hex"],
      [/0[0-7']*[0-7](@integersuffix)/, "number.octal"],
      [/0[bB][0-1']*[0-1](@integersuffix)/, "number.binary"],
      [/\d[\d']*\d(@integersuffix)/, "number"],
      [/\d(@integersuffix)/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@doccomment"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*\\$/, "comment", "@linecomment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    linecomment: [
      [/.*[^\\]$/, "comment", "@pop"],
      [/[^]+/, "comment"]
    ],
    doccomment: [
      [/[^\/*]+/, "comment.doc"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    raw: [
      [
        /(.*)(\))(?:([^ ()\\\t"]*))(\")/,
        {
          cases: {
            "$3==$S2": [
              "string.raw",
              "string.raw.end",
              "string.raw.end",
              { token: "string.raw.end", next: "@pop" }
            ],
            "@default": ["string.raw", "string.raw", "string.raw", "string.raw"]
          }
        }
      ],
      [/.*/, "string.raw"]
    ],
    annotation: [
      { include: "@whitespace" },
      [/using|alignas/, "keyword"],
      [/[a-zA-Z0-9_]+/, "annotation"],
      [/[,:]/, "delimiter"],
      [/[()]/, "@brackets"],
      [/\]\s*\]/, { token: "annotation", next: "@pop" }]
    ],
    include: [
      [
        /(\s*)(<)([^<>]*)(>)/,
        [
          "",
          "keyword.directive.include.begin",
          "string.include.identifier",
          { token: "keyword.directive.include.end", next: "@pop" }
        ]
      ],
      [
        /(\s*)(")([^"]*)(")/,
        [
          "",
          "keyword.directive.include.begin",
          "string.include.identifier",
          { token: "keyword.directive.include.end", next: "@pop" }
        ]
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/csharp/csharp.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/csharp/csharp.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/csharp/csharp.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\#\$\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: '"', close: '"' }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#region\\b"),
      end: new RegExp("^\\s*#endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".cs",
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  keywords: [
    "extern",
    "alias",
    "using",
    "bool",
    "decimal",
    "sbyte",
    "byte",
    "short",
    "ushort",
    "int",
    "uint",
    "long",
    "ulong",
    "char",
    "float",
    "double",
    "object",
    "dynamic",
    "string",
    "assembly",
    "is",
    "as",
    "ref",
    "out",
    "this",
    "base",
    "new",
    "typeof",
    "void",
    "checked",
    "unchecked",
    "default",
    "delegate",
    "var",
    "const",
    "if",
    "else",
    "switch",
    "case",
    "while",
    "do",
    "for",
    "foreach",
    "in",
    "break",
    "continue",
    "goto",
    "return",
    "throw",
    "try",
    "catch",
    "finally",
    "lock",
    "yield",
    "from",
    "let",
    "where",
    "join",
    "on",
    "equals",
    "into",
    "orderby",
    "ascending",
    "descending",
    "select",
    "group",
    "by",
    "namespace",
    "partial",
    "class",
    "field",
    "event",
    "method",
    "param",
    "public",
    "protected",
    "internal",
    "private",
    "abstract",
    "sealed",
    "static",
    "struct",
    "readonly",
    "volatile",
    "virtual",
    "override",
    "params",
    "get",
    "set",
    "add",
    "remove",
    "operator",
    "true",
    "false",
    "implicit",
    "explicit",
    "interface",
    "enum",
    "null",
    "async",
    "await",
    "fixed",
    "sizeof",
    "stackalloc",
    "unsafe",
    "nameof",
    "when"
  ],
  namespaceFollows: ["namespace", "using"],
  parenFollows: ["if", "for", "while", "switch", "foreach", "using", "catch", "when"],
  operators: [
    "=",
    "??",
    "||",
    "&&",
    "|",
    "^",
    "&",
    "==",
    "!=",
    "<=",
    ">=",
    "<<",
    "+",
    "-",
    "*",
    "/",
    "%",
    "!",
    "~",
    "++",
    "--",
    "+=",
    "-=",
    "*=",
    "/=",
    "%=",
    "&=",
    "|=",
    "^=",
    "<<=",
    ">>=",
    ">>",
    "=>"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [
        /\@?[a-zA-Z_]\w*/,
        {
          cases: {
            "@namespaceFollows": {
              token: "keyword.$0",
              next: "@namespace"
            },
            "@keywords": {
              token: "keyword.$0",
              next: "@qualified"
            },
            "@default": { token: "identifier", next: "@qualified" }
          }
        }
      ],
      { include: "@whitespace" },
      [
        /}/,
        {
          cases: {
            "$S2==interpolatedstring": {
              token: "string.quote",
              next: "@pop"
            },
            "$S2==litinterpstring": {
              token: "string.quote",
              next: "@pop"
            },
            "@default": "@brackets"
          }
        }
      ],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/[0-9_]*\.[0-9_]+([eE][\-+]?\d+)?[fFdD]?/, "number.float"],
      [/0[xX][0-9a-fA-F_]+/, "number.hex"],
      [/0[bB][01_]+/, "number.hex"],
      [/[0-9_]+/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, { token: "string.quote", next: "@string" }],
      [/\$\@"/, { token: "string.quote", next: "@litinterpstring" }],
      [/\@"/, { token: "string.quote", next: "@litstring" }],
      [/\$"/, { token: "string.quote", next: "@interpolatedstring" }],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    qualified: [
      [
        /[a-zA-Z_][\w]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      [/\./, "delimiter"],
      ["", "", "@pop"]
    ],
    namespace: [
      { include: "@whitespace" },
      [/[A-Z]\w*/, "namespace"],
      [/[\.=]/, "delimiter"],
      ["", "", "@pop"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      ["\\*/", "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, { token: "string.quote", next: "@pop" }]
    ],
    litstring: [
      [/[^"]+/, "string"],
      [/""/, "string.escape"],
      [/"/, { token: "string.quote", next: "@pop" }]
    ],
    litinterpstring: [
      [/[^"{]+/, "string"],
      [/""/, "string.escape"],
      [/{{/, "string.escape"],
      [/}}/, "string.escape"],
      [/{/, { token: "string.quote", next: "root.litinterpstring" }],
      [/"/, { token: "string.quote", next: "@pop" }]
    ],
    interpolatedstring: [
      [/[^\\"{]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/{{/, "string.escape"],
      [/}}/, "string.escape"],
      [/{/, { token: "string.quote", next: "root.interpolatedstring" }],
      [/"/, { token: "string.quote", next: "@pop" }]
    ],
    whitespace: [
      [/^[ \t\v\f]*#((r)|(load))(?=\s)/, "directive.csx"],
      [/^[ \t\v\f]*#\w.*$/, "namespace.cpp"],
      [/[ \t\v\f\r\n]+/, ""],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/csp/csp.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/csp/csp.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/csp/csp.ts
var conf = {
  brackets: [],
  autoClosingPairs: [],
  surroundingPairs: []
};
var language = {
  keywords: [],
  typeKeywords: [],
  tokenPostfix: ".csp",
  operators: [],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/child-src/, "string.quote"],
      [/connect-src/, "string.quote"],
      [/default-src/, "string.quote"],
      [/font-src/, "string.quote"],
      [/frame-src/, "string.quote"],
      [/img-src/, "string.quote"],
      [/manifest-src/, "string.quote"],
      [/media-src/, "string.quote"],
      [/object-src/, "string.quote"],
      [/script-src/, "string.quote"],
      [/style-src/, "string.quote"],
      [/worker-src/, "string.quote"],
      [/base-uri/, "string.quote"],
      [/plugin-types/, "string.quote"],
      [/sandbox/, "string.quote"],
      [/disown-opener/, "string.quote"],
      [/form-action/, "string.quote"],
      [/frame-ancestors/, "string.quote"],
      [/report-uri/, "string.quote"],
      [/report-to/, "string.quote"],
      [/upgrade-insecure-requests/, "string.quote"],
      [/block-all-mixed-content/, "string.quote"],
      [/require-sri-for/, "string.quote"],
      [/reflected-xss/, "string.quote"],
      [/referrer/, "string.quote"],
      [/policy-uri/, "string.quote"],
      [/'self'/, "string.quote"],
      [/'unsafe-inline'/, "string.quote"],
      [/'unsafe-eval'/, "string.quote"],
      [/'strict-dynamic'/, "string.quote"],
      [/'unsafe-hashed-attributes'/, "string.quote"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/css/css.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/css/css.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/css/css.ts
var conf = {
  wordPattern: /(#?-?\d*\.\d\w*%?)|((::|[@#.!:])?[\w-?]+%?)|::|[@#.!:]/g,
  comments: {
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*\\/\\*\\s*#region\\b\\s*(.*?)\\s*\\*\\/"),
      end: new RegExp("^\\s*\\/\\*\\s*#endregion\\b.*\\*\\/")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".css",
  ws: "[ 	\n\r\f]*",
  identifier: "-?-?([a-zA-Z]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))([\\w\\-]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))*",
  brackets: [
    { open: "{", close: "}", token: "delimiter.bracket" },
    { open: "[", close: "]", token: "delimiter.bracket" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  tokenizer: {
    root: [{ include: "@selector" }],
    selector: [
      { include: "@comments" },
      { include: "@import" },
      { include: "@strings" },
      [
        "[@](keyframes|-webkit-keyframes|-moz-keyframes|-o-keyframes)",
        { token: "keyword", next: "@keyframedeclaration" }
      ],
      ["[@](page|content|font-face|-moz-document)", { token: "keyword" }],
      ["[@](charset|namespace)", { token: "keyword", next: "@declarationbody" }],
      [
        "(url-prefix)(\\()",
        ["attribute.value", { token: "delimiter.parenthesis", next: "@urldeclaration" }]
      ],
      [
        "(url)(\\()",
        ["attribute.value", { token: "delimiter.parenthesis", next: "@urldeclaration" }]
      ],
      { include: "@selectorname" },
      ["[\\*]", "tag"],
      ["[>\\+,]", "delimiter"],
      ["\\[", { token: "delimiter.bracket", next: "@selectorattribute" }],
      ["{", { token: "delimiter.bracket", next: "@selectorbody" }]
    ],
    selectorbody: [
      { include: "@comments" },
      ["[*_]?@identifier@ws:(?=(\\s|\\d|[^{;}]*[;}]))", "attribute.name", "@rulevalue"],
      ["}", { token: "delimiter.bracket", next: "@pop" }]
    ],
    selectorname: [
      ["(\\.|#(?=[^{])|%|(@identifier)|:)+", "tag"]
    ],
    selectorattribute: [{ include: "@term" }, ["]", { token: "delimiter.bracket", next: "@pop" }]],
    term: [
      { include: "@comments" },
      [
        "(url-prefix)(\\()",
        ["attribute.value", { token: "delimiter.parenthesis", next: "@urldeclaration" }]
      ],
      [
        "(url)(\\()",
        ["attribute.value", { token: "delimiter.parenthesis", next: "@urldeclaration" }]
      ],
      { include: "@functioninvocation" },
      { include: "@numbers" },
      { include: "@name" },
      { include: "@strings" },
      ["([<>=\\+\\-\\*\\/\\^\\|\\~,])", "delimiter"],
      [",", "delimiter"]
    ],
    rulevalue: [
      { include: "@comments" },
      { include: "@strings" },
      { include: "@term" },
      ["!important", "keyword"],
      [";", "delimiter", "@pop"],
      ["(?=})", { token: "", next: "@pop" }]
    ],
    warndebug: [["[@](warn|debug)", { token: "keyword", next: "@declarationbody" }]],
    import: [["[@](import)", { token: "keyword", next: "@declarationbody" }]],
    urldeclaration: [
      { include: "@strings" },
      ["[^)\r\n]+", "string"],
      ["\\)", { token: "delimiter.parenthesis", next: "@pop" }]
    ],
    parenthizedterm: [
      { include: "@term" },
      ["\\)", { token: "delimiter.parenthesis", next: "@pop" }]
    ],
    declarationbody: [
      { include: "@term" },
      [";", "delimiter", "@pop"],
      ["(?=})", { token: "", next: "@pop" }]
    ],
    comments: [
      ["\\/\\*", "comment", "@comment"],
      ["\\/\\/+.*", "comment"]
    ],
    comment: [
      ["\\*\\/", "comment", "@pop"],
      [/[^*/]+/, "comment"],
      [/./, "comment"]
    ],
    name: [["@identifier", "attribute.value"]],
    numbers: [
      ["-?(\\d*\\.)?\\d+([eE][\\-+]?\\d+)?", { token: "attribute.value.number", next: "@units" }],
      ["#[0-9a-fA-F_]+(?!\\w)", "attribute.value.hex"]
    ],
    units: [
      [
        "(em|ex|ch|rem|fr|vmin|vmax|vw|vh|vm|cm|mm|in|px|pt|pc|deg|grad|rad|turn|s|ms|Hz|kHz|%)?",
        "attribute.value.unit",
        "@pop"
      ]
    ],
    keyframedeclaration: [
      ["@identifier", "attribute.value"],
      ["{", { token: "delimiter.bracket", switchTo: "@keyframebody" }]
    ],
    keyframebody: [
      { include: "@term" },
      ["{", { token: "delimiter.bracket", next: "@selectorbody" }],
      ["}", { token: "delimiter.bracket", next: "@pop" }]
    ],
    functioninvocation: [
      ["@identifier\\(", { token: "attribute.value", next: "@functionarguments" }]
    ],
    functionarguments: [
      ["\\$@identifier@ws:", "attribute.name"],
      ["[,]", "delimiter"],
      { include: "@term" },
      ["\\)", { token: "attribute.value", next: "@pop" }]
    ],
    strings: [
      ['~?"', { token: "string", next: "@stringenddoublequote" }],
      ["~?'", { token: "string", next: "@stringendquote" }]
    ],
    stringenddoublequote: [
      ["\\\\.", "string"],
      ['"', { token: "string", next: "@pop" }],
      [/[^\\"]+/, "string"],
      [".", "string"]
    ],
    stringendquote: [
      ["\\\\.", "string"],
      ["'", { token: "string", next: "@pop" }],
      [/[^\\']+/, "string"],
      [".", "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/cypher/cypher.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/cypher/cypher.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/cypher/cypher.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: `.cypher`,
  ignoreCase: true,
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.bracket" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "ALL",
    "AND",
    "AS",
    "ASC",
    "ASCENDING",
    "BY",
    "CALL",
    "CASE",
    "CONTAINS",
    "CREATE",
    "DELETE",
    "DESC",
    "DESCENDING",
    "DETACH",
    "DISTINCT",
    "ELSE",
    "END",
    "ENDS",
    "EXISTS",
    "IN",
    "IS",
    "LIMIT",
    "MANDATORY",
    "MATCH",
    "MERGE",
    "NOT",
    "ON",
    "ON",
    "OPTIONAL",
    "OR",
    "ORDER",
    "REMOVE",
    "RETURN",
    "SET",
    "SKIP",
    "STARTS",
    "THEN",
    "UNION",
    "UNWIND",
    "WHEN",
    "WHERE",
    "WITH",
    "XOR",
    "YIELD"
  ],
  builtinLiterals: ["true", "TRUE", "false", "FALSE", "null", "NULL"],
  builtinFunctions: [
    "abs",
    "acos",
    "asin",
    "atan",
    "atan2",
    "avg",
    "ceil",
    "coalesce",
    "collect",
    "cos",
    "cot",
    "count",
    "degrees",
    "e",
    "endNode",
    "exists",
    "exp",
    "floor",
    "head",
    "id",
    "keys",
    "labels",
    "last",
    "left",
    "length",
    "log",
    "log10",
    "lTrim",
    "max",
    "min",
    "nodes",
    "percentileCont",
    "percentileDisc",
    "pi",
    "properties",
    "radians",
    "rand",
    "range",
    "relationships",
    "replace",
    "reverse",
    "right",
    "round",
    "rTrim",
    "sign",
    "sin",
    "size",
    "split",
    "sqrt",
    "startNode",
    "stDev",
    "stDevP",
    "substring",
    "sum",
    "tail",
    "tan",
    "timestamp",
    "toBoolean",
    "toFloat",
    "toInteger",
    "toLower",
    "toString",
    "toUpper",
    "trim",
    "type"
  ],
  operators: [
    "+",
    "-",
    "*",
    "/",
    "%",
    "^",
    "=",
    "<>",
    "<",
    ">",
    "<=",
    ">=",
    "->",
    "<-",
    "-->",
    "<--"
  ],
  escapes: /\\(?:[tbnrf\\"'`]|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  digits: /\d+/,
  octaldigits: /[0-7]+/,
  hexdigits: /[0-9a-fA-F]+/,
  tokenizer: {
    root: [[/[{}[\]()]/, "@brackets"], { include: "common" }],
    common: [
      { include: "@whitespace" },
      { include: "@numbers" },
      { include: "@strings" },
      [/:[a-zA-Z_][\w]*/, "type.identifier"],
      [
        /[a-zA-Z_][\w]*(?=\()/,
        {
          cases: {
            "@builtinFunctions": "predefined.function"
          }
        }
      ],
      [
        /[a-zA-Z_$][\w$]*/,
        {
          cases: {
            "@keywords": "keyword",
            "@builtinLiterals": "predefined.literal",
            "@default": "identifier"
          }
        }
      ],
      [/`/, "identifier.escape", "@identifierBacktick"],
      [/[;,.:|]/, "delimiter"],
      [
        /[<>=%+\-*/^]+/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ]
    ],
    numbers: [
      [/-?(@digits)[eE](-?(@digits))?/, "number.float"],
      [/-?(@digits)?\.(@digits)([eE]-?(@digits))?/, "number.float"],
      [/-?0x(@hexdigits)/, "number.hex"],
      [/-?0(@octaldigits)/, "number.octal"],
      [/-?(@digits)/, "number"]
    ],
    strings: [
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@stringDouble"],
      [/'/, "string", "@stringSingle"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/\/\/.*/, "comment"],
      [/[^/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[/*]/, "comment"]
    ],
    stringDouble: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string"],
      [/\\./, "string.invalid"],
      [/"/, "string", "@pop"]
    ],
    stringSingle: [
      [/[^\\']+/, "string"],
      [/@escapes/, "string"],
      [/\\./, "string.invalid"],
      [/'/, "string", "@pop"]
    ],
    identifierBacktick: [
      [/[^\\`]+/, "identifier.escape"],
      [/@escapes/, "identifier.escape"],
      [/\\./, "identifier.escape.invalid"],
      [/`/, "identifier.escape", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/dart/dart.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/dart/dart.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/dart/dart.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "`", close: "`", notIn: ["string", "comment"] },
    { open: "/**", close: " */", notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "`", close: "`" }
  ],
  folding: {
    markers: {
      start: /^\s*\s*#?region\b/,
      end: /^\s*\s*#?endregion\b/
    }
  }
};
var language = {
  defaultToken: "invalid",
  tokenPostfix: ".dart",
  keywords: [
    "abstract",
    "dynamic",
    "implements",
    "show",
    "as",
    "else",
    "import",
    "static",
    "assert",
    "enum",
    "in",
    "super",
    "async",
    "export",
    "interface",
    "switch",
    "await",
    "extends",
    "is",
    "sync",
    "break",
    "external",
    "library",
    "this",
    "case",
    "factory",
    "mixin",
    "throw",
    "catch",
    "false",
    "new",
    "true",
    "class",
    "final",
    "null",
    "try",
    "const",
    "finally",
    "on",
    "typedef",
    "continue",
    "for",
    "operator",
    "var",
    "covariant",
    "Function",
    "part",
    "void",
    "default",
    "get",
    "rethrow",
    "while",
    "deferred",
    "hide",
    "return",
    "with",
    "do",
    "if",
    "set",
    "yield"
  ],
  typeKeywords: ["int", "double", "String", "bool"],
  operators: [
    "+",
    "-",
    "*",
    "/",
    "~/",
    "%",
    "++",
    "--",
    "==",
    "!=",
    ">",
    "<",
    ">=",
    "<=",
    "=",
    "-=",
    "/=",
    "%=",
    ">>=",
    "^=",
    "+=",
    "*=",
    "~/=",
    "<<=",
    "&=",
    "!=",
    "||",
    "&&",
    "&",
    "|",
    "^",
    "~",
    "<<",
    ">>",
    "!",
    ">>>",
    "??",
    "?",
    ":",
    "|="
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  digits: /\d+(_+\d+)*/,
  octaldigits: /[0-7]+(_+[0-7]+)*/,
  binarydigits: /[0-1]+(_+[0-1]+)*/,
  hexdigits: /[[0-9a-fA-F]+(_+[0-9a-fA-F]+)*/,
  regexpctl: /[(){}\[\]\$\^|\-*+?\.]/,
  regexpesc: /\\(?:[bBdDfnrstvwWn0\\\/]|@regexpctl|c[A-Z]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4})/,
  tokenizer: {
    root: [[/[{}]/, "delimiter.bracket"], { include: "common" }],
    common: [
      [
        /[a-z_$][\w$]*/,
        {
          cases: {
            "@typeKeywords": "type.identifier",
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/[A-Z_$][\w\$]*/, "type.identifier"],
      { include: "@whitespace" },
      [
        /\/(?=([^\\\/]|\\.)+\/([gimsuy]*)(\s*)(\.|;|,|\)|\]|\}|$))/,
        { token: "regexp", bracket: "@open", next: "@regexp" }
      ],
      [/@[a-zA-Z]+/, "annotation"],
      [/[()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [/!(?=([^=]|$))/, "delimiter"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/(@digits)[eE]([\-+]?(@digits))?/, "number.float"],
      [/(@digits)\.(@digits)([eE][\-+]?(@digits))?/, "number.float"],
      [/0[xX](@hexdigits)n?/, "number.hex"],
      [/0[oO]?(@octaldigits)n?/, "number.octal"],
      [/0[bB](@binarydigits)n?/, "number.binary"],
      [/(@digits)n?/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string_double"],
      [/'/, "string", "@string_single"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@jsdoc"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/\/.*$/, "comment.doc"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    jsdoc: [
      [/[^\/*]+/, "comment.doc"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    regexp: [
      [
        /(\{)(\d+(?:,\d*)?)(\})/,
        ["regexp.escape.control", "regexp.escape.control", "regexp.escape.control"]
      ],
      [
        /(\[)(\^?)(?=(?:[^\]\\\/]|\\.)+)/,
        ["regexp.escape.control", { token: "regexp.escape.control", next: "@regexrange" }]
      ],
      [/(\()(\?:|\?=|\?!)/, ["regexp.escape.control", "regexp.escape.control"]],
      [/[()]/, "regexp.escape.control"],
      [/@regexpctl/, "regexp.escape.control"],
      [/[^\\\/]/, "regexp"],
      [/@regexpesc/, "regexp.escape"],
      [/\\\./, "regexp.invalid"],
      [/(\/)([gimsuy]*)/, [{ token: "regexp", bracket: "@close", next: "@pop" }, "keyword.other"]]
    ],
    regexrange: [
      [/-/, "regexp.escape.control"],
      [/\^/, "regexp.invalid"],
      [/@regexpesc/, "regexp.escape"],
      [/[^\]]/, "regexp"],
      [
        /\]/,
        {
          token: "regexp.escape.control",
          next: "@pop",
          bracket: "@close"
        }
      ]
    ],
    string_double: [
      [/[^\\"\$]+/, "string"],
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"],
      [/\$\w+/, "identifier"]
    ],
    string_single: [
      [/[^\\'\$]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/'/, "string", "@pop"],
      [/\$\w+/, "identifier"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/dockerfile/dockerfile.js":
/*!************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/dockerfile/dockerfile.js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/dockerfile/dockerfile.ts
var conf = {
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".dockerfile",
  variable: /\${?[\w]+}?/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      { include: "@comment" },
      [/(ONBUILD)(\s+)/, ["keyword", ""]],
      [/(ENV)(\s+)([\w]+)/, ["keyword", "", { token: "variable", next: "@arguments" }]],
      [
        /(FROM|MAINTAINER|RUN|EXPOSE|ENV|ADD|ARG|VOLUME|LABEL|USER|WORKDIR|COPY|CMD|STOPSIGNAL|SHELL|HEALTHCHECK|ENTRYPOINT)/,
        { token: "keyword", next: "@arguments" }
      ]
    ],
    arguments: [
      { include: "@whitespace" },
      { include: "@strings" },
      [
        /(@variable)/,
        {
          cases: {
            "@eos": { token: "variable", next: "@popall" },
            "@default": "variable"
          }
        }
      ],
      [
        /\\/,
        {
          cases: {
            "@eos": "",
            "@default": ""
          }
        }
      ],
      [
        /./,
        {
          cases: {
            "@eos": { token: "", next: "@popall" },
            "@default": ""
          }
        }
      ]
    ],
    whitespace: [
      [
        /\s+/,
        {
          cases: {
            "@eos": { token: "", next: "@popall" },
            "@default": ""
          }
        }
      ]
    ],
    comment: [[/(^#.*$)/, "comment", "@popall"]],
    strings: [
      [/\\'$/, "", "@popall"],
      [/\\'/, ""],
      [/'$/, "string", "@popall"],
      [/'/, "string", "@stringBody"],
      [/"$/, "string", "@popall"],
      [/"/, "string", "@dblStringBody"]
    ],
    stringBody: [
      [
        /[^\\\$']/,
        {
          cases: {
            "@eos": { token: "string", next: "@popall" },
            "@default": "string"
          }
        }
      ],
      [/\\./, "string.escape"],
      [/'$/, "string", "@popall"],
      [/'/, "string", "@pop"],
      [/(@variable)/, "variable"],
      [/\\$/, "string"],
      [/$/, "string", "@popall"]
    ],
    dblStringBody: [
      [
        /[^\\\$"]/,
        {
          cases: {
            "@eos": { token: "string", next: "@popall" },
            "@default": "string"
          }
        }
      ],
      [/\\./, "string.escape"],
      [/"$/, "string", "@popall"],
      [/"/, "string", "@pop"],
      [/(@variable)/, "variable"],
      [/\\$/, "string"],
      [/$/, "string", "@popall"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/ecl/ecl.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/ecl/ecl.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/ecl/ecl.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: '"', close: '"' }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".ecl",
  ignoreCase: true,
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  pounds: [
    "append",
    "break",
    "declare",
    "demangle",
    "end",
    "for",
    "getdatatype",
    "if",
    "inmodule",
    "loop",
    "mangle",
    "onwarning",
    "option",
    "set",
    "stored",
    "uniquename"
  ].join("|"),
  keywords: [
    "__compressed__",
    "after",
    "all",
    "and",
    "any",
    "as",
    "atmost",
    "before",
    "beginc",
    "best",
    "between",
    "case",
    "cluster",
    "compressed",
    "compression",
    "const",
    "counter",
    "csv",
    "default",
    "descend",
    "embed",
    "encoding",
    "encrypt",
    "end",
    "endc",
    "endembed",
    "endmacro",
    "enum",
    "escape",
    "except",
    "exclusive",
    "expire",
    "export",
    "extend",
    "fail",
    "few",
    "fileposition",
    "first",
    "flat",
    "forward",
    "from",
    "full",
    "function",
    "functionmacro",
    "group",
    "grouped",
    "heading",
    "hole",
    "ifblock",
    "import",
    "in",
    "inner",
    "interface",
    "internal",
    "joined",
    "keep",
    "keyed",
    "last",
    "left",
    "limit",
    "linkcounted",
    "literal",
    "little_endian",
    "load",
    "local",
    "locale",
    "lookup",
    "lzw",
    "macro",
    "many",
    "maxcount",
    "maxlength",
    "min skew",
    "module",
    "mofn",
    "multiple",
    "named",
    "namespace",
    "nocase",
    "noroot",
    "noscan",
    "nosort",
    "not",
    "noxpath",
    "of",
    "onfail",
    "only",
    "opt",
    "or",
    "outer",
    "overwrite",
    "packed",
    "partition",
    "penalty",
    "physicallength",
    "pipe",
    "prefetch",
    "quote",
    "record",
    "repeat",
    "retry",
    "return",
    "right",
    "right1",
    "right2",
    "rows",
    "rowset",
    "scan",
    "scope",
    "self",
    "separator",
    "service",
    "shared",
    "skew",
    "skip",
    "smart",
    "soapaction",
    "sql",
    "stable",
    "store",
    "terminator",
    "thor",
    "threshold",
    "timelimit",
    "timeout",
    "token",
    "transform",
    "trim",
    "type",
    "unicodeorder",
    "unordered",
    "unsorted",
    "unstable",
    "update",
    "use",
    "validate",
    "virtual",
    "whole",
    "width",
    "wild",
    "within",
    "wnotrim",
    "xml",
    "xpath"
  ],
  functions: [
    "abs",
    "acos",
    "aggregate",
    "allnodes",
    "apply",
    "ascii",
    "asin",
    "assert",
    "asstring",
    "atan",
    "atan2",
    "ave",
    "build",
    "buildindex",
    "case",
    "catch",
    "choose",
    "choosen",
    "choosesets",
    "clustersize",
    "combine",
    "correlation",
    "cos",
    "cosh",
    "count",
    "covariance",
    "cron",
    "dataset",
    "dedup",
    "define",
    "denormalize",
    "dictionary",
    "distribute",
    "distributed",
    "distribution",
    "ebcdic",
    "enth",
    "error",
    "evaluate",
    "event",
    "eventextra",
    "eventname",
    "exists",
    "exp",
    "fail",
    "failcode",
    "failmessage",
    "fetch",
    "fromunicode",
    "fromxml",
    "getenv",
    "getisvalid",
    "global",
    "graph",
    "group",
    "hash",
    "hash32",
    "hash64",
    "hashcrc",
    "hashmd5",
    "having",
    "httpcall",
    "httpheader",
    "if",
    "iff",
    "index",
    "intformat",
    "isvalid",
    "iterate",
    "join",
    "keydiff",
    "keypatch",
    "keyunicode",
    "length",
    "library",
    "limit",
    "ln",
    "loadxml",
    "local",
    "log",
    "loop",
    "map",
    "matched",
    "matchlength",
    "matchposition",
    "matchtext",
    "matchunicode",
    "max",
    "merge",
    "mergejoin",
    "min",
    "nofold",
    "nolocal",
    "nonempty",
    "normalize",
    "nothor",
    "notify",
    "output",
    "parallel",
    "parse",
    "pipe",
    "power",
    "preload",
    "process",
    "project",
    "pull",
    "random",
    "range",
    "rank",
    "ranked",
    "realformat",
    "recordof",
    "regexfind",
    "regexreplace",
    "regroup",
    "rejected",
    "rollup",
    "round",
    "roundup",
    "row",
    "rowdiff",
    "sample",
    "sequential",
    "set",
    "sin",
    "sinh",
    "sizeof",
    "soapcall",
    "sort",
    "sorted",
    "sqrt",
    "stepped",
    "stored",
    "sum",
    "table",
    "tan",
    "tanh",
    "thisnode",
    "topn",
    "tounicode",
    "toxml",
    "transfer",
    "transform",
    "trim",
    "truncate",
    "typeof",
    "ungroup",
    "unicodeorder",
    "variance",
    "wait",
    "which",
    "workunit",
    "xmldecode",
    "xmlencode",
    "xmltext",
    "xmlunicode"
  ],
  typesint: ["integer", "unsigned"].join("|"),
  typesnum: ["data", "qstring", "string", "unicode", "utf8", "varstring", "varunicode"],
  typesone: [
    "ascii",
    "big_endian",
    "boolean",
    "data",
    "decimal",
    "ebcdic",
    "grouped",
    "integer",
    "linkcounted",
    "pattern",
    "qstring",
    "real",
    "record",
    "rule",
    "set of",
    "streamed",
    "string",
    "token",
    "udecimal",
    "unicode",
    "unsigned",
    "utf8",
    "varstring",
    "varunicode"
  ].join("|"),
  operators: ["+", "-", "/", ":=", "<", "<>", "=", ">", "\\", "and", "in", "not", "or"],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/@typesint[4|8]/, "type"],
      [/#(@pounds)/, "type"],
      [/@typesone/, "type"],
      [
        /[a-zA-Z_$][\w-$]*/,
        {
          cases: {
            "@functions": "keyword.function",
            "@keywords": "keyword",
            "@operators": "operator"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/[0-9_]*\.[0-9_]+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F_]+/, "number.hex"],
      [/0[bB][01]+/, "number.hex"],
      [/[0-9_]+/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\v\f\r\n]+/, ""],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    string: [
      [/[^\\']+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/'/, "string", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/elixir/elixir.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/elixir/elixir.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/elixir/elixir.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "'", close: "'" },
    { open: '"', close: '"' }
  ],
  autoClosingPairs: [
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["comment"] },
    { open: '"""', close: '"""' },
    { open: "`", close: "`", notIn: ["string", "comment"] },
    { open: "(", close: ")" },
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "<<", close: ">>" }
  ],
  indentationRules: {
    increaseIndentPattern: /^\s*(after|else|catch|rescue|fn|[^#]*(do|<\-|\->|\{|\[|\=))\s*$/,
    decreaseIndentPattern: /^\s*((\}|\])\s*$|(after|else|catch|rescue|end)\b)/
  }
};
var language = {
  defaultToken: "source",
  tokenPostfix: ".elixir",
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "<<", close: ">>", token: "delimiter.angle.special" }
  ],
  declarationKeywords: [
    "def",
    "defp",
    "defn",
    "defnp",
    "defguard",
    "defguardp",
    "defmacro",
    "defmacrop",
    "defdelegate",
    "defcallback",
    "defmacrocallback",
    "defmodule",
    "defprotocol",
    "defexception",
    "defimpl",
    "defstruct"
  ],
  operatorKeywords: ["and", "in", "not", "or", "when"],
  namespaceKeywords: ["alias", "import", "require", "use"],
  otherKeywords: [
    "after",
    "case",
    "catch",
    "cond",
    "do",
    "else",
    "end",
    "fn",
    "for",
    "if",
    "quote",
    "raise",
    "receive",
    "rescue",
    "super",
    "throw",
    "try",
    "unless",
    "unquote_splicing",
    "unquote",
    "with"
  ],
  constants: ["true", "false", "nil"],
  nameBuiltin: ["__MODULE__", "__DIR__", "__ENV__", "__CALLER__", "__STACKTRACE__"],
  operator: /-[->]?|!={0,2}|\*{1,2}|\/|\\\\|&{1,3}|\.\.?|\^(?:\^\^)?|\+\+?|<(?:-|<<|=|>|\|>|~>?)?|=~|={1,3}|>(?:=|>>)?|\|~>|\|>|\|{1,3}|~>>?|~~~|::/,
  variableName: /[a-z_][a-zA-Z0-9_]*[?!]?/,
  atomName: /[a-zA-Z_][a-zA-Z0-9_@]*[?!]?|@specialAtomName|@operator/,
  specialAtomName: /\.\.\.|<<>>|%\{\}|%|\{\}/,
  aliasPart: /[A-Z][a-zA-Z0-9_]*/,
  moduleName: /@aliasPart(?:\.@aliasPart)*/,
  sigilSymmetricDelimiter: /"""|'''|"|'|\/|\|/,
  sigilStartDelimiter: /@sigilSymmetricDelimiter|<|\{|\[|\(/,
  sigilEndDelimiter: /@sigilSymmetricDelimiter|>|\}|\]|\)/,
  sigilModifiers: /[a-zA-Z0-9]*/,
  decimal: /\d(?:_?\d)*/,
  hex: /[0-9a-fA-F](_?[0-9a-fA-F])*/,
  octal: /[0-7](_?[0-7])*/,
  binary: /[01](_?[01])*/,
  escape: /\\u[0-9a-fA-F]{4}|\\x[0-9a-fA-F]{2}|\\./,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      { include: "@comments" },
      { include: "@keywordsShorthand" },
      { include: "@numbers" },
      { include: "@identifiers" },
      { include: "@strings" },
      { include: "@atoms" },
      { include: "@sigils" },
      { include: "@attributes" },
      { include: "@symbols" }
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [[/(#)(.*)/, ["comment.punctuation", "comment"]]],
    keywordsShorthand: [
      [/(@atomName)(:)(\s+)/, ["constant", "constant.punctuation", "white"]],
      [
        /"(?=([^"]|#\{.*?\}|\\")*":)/,
        { token: "constant.delimiter", next: "@doubleQuotedStringKeyword" }
      ],
      [
        /'(?=([^']|#\{.*?\}|\\')*':)/,
        { token: "constant.delimiter", next: "@singleQuotedStringKeyword" }
      ]
    ],
    doubleQuotedStringKeyword: [
      [/":/, { token: "constant.delimiter", next: "@pop" }],
      { include: "@stringConstantContentInterpol" }
    ],
    singleQuotedStringKeyword: [
      [/':/, { token: "constant.delimiter", next: "@pop" }],
      { include: "@stringConstantContentInterpol" }
    ],
    numbers: [
      [/0b@binary/, "number.binary"],
      [/0o@octal/, "number.octal"],
      [/0x@hex/, "number.hex"],
      [/@decimal\.@decimal([eE]-?@decimal)?/, "number.float"],
      [/@decimal/, "number"]
    ],
    identifiers: [
      [
        /\b(defp?|defnp?|defmacrop?|defguardp?|defdelegate)(\s+)(@variableName)(?!\s+@operator)/,
        [
          "keyword.declaration",
          "white",
          {
            cases: {
              unquote: "keyword",
              "@default": "function"
            }
          }
        ]
      ],
      [
        /(@variableName)(?=\s*\.?\s*\()/,
        {
          cases: {
            "@declarationKeywords": "keyword.declaration",
            "@namespaceKeywords": "keyword",
            "@otherKeywords": "keyword",
            "@default": "function.call"
          }
        }
      ],
      [
        /(@moduleName)(\s*)(\.)(\s*)(@variableName)/,
        ["type.identifier", "white", "operator", "white", "function.call"]
      ],
      [
        /(:)(@atomName)(\s*)(\.)(\s*)(@variableName)/,
        ["constant.punctuation", "constant", "white", "operator", "white", "function.call"]
      ],
      [
        /(\|>)(\s*)(@variableName)/,
        [
          "operator",
          "white",
          {
            cases: {
              "@otherKeywords": "keyword",
              "@default": "function.call"
            }
          }
        ]
      ],
      [
        /(&)(\s*)(@variableName)/,
        ["operator", "white", "function.call"]
      ],
      [
        /@variableName/,
        {
          cases: {
            "@declarationKeywords": "keyword.declaration",
            "@operatorKeywords": "keyword.operator",
            "@namespaceKeywords": "keyword",
            "@otherKeywords": "keyword",
            "@constants": "constant.language",
            "@nameBuiltin": "variable.language",
            "_.*": "comment.unused",
            "@default": "identifier"
          }
        }
      ],
      [/@moduleName/, "type.identifier"]
    ],
    strings: [
      [/"""/, { token: "string.delimiter", next: "@doubleQuotedHeredoc" }],
      [/'''/, { token: "string.delimiter", next: "@singleQuotedHeredoc" }],
      [/"/, { token: "string.delimiter", next: "@doubleQuotedString" }],
      [/'/, { token: "string.delimiter", next: "@singleQuotedString" }]
    ],
    doubleQuotedHeredoc: [
      [/"""/, { token: "string.delimiter", next: "@pop" }],
      { include: "@stringContentInterpol" }
    ],
    singleQuotedHeredoc: [
      [/'''/, { token: "string.delimiter", next: "@pop" }],
      { include: "@stringContentInterpol" }
    ],
    doubleQuotedString: [
      [/"/, { token: "string.delimiter", next: "@pop" }],
      { include: "@stringContentInterpol" }
    ],
    singleQuotedString: [
      [/'/, { token: "string.delimiter", next: "@pop" }],
      { include: "@stringContentInterpol" }
    ],
    atoms: [
      [/(:)(@atomName)/, ["constant.punctuation", "constant"]],
      [/:"/, { token: "constant.delimiter", next: "@doubleQuotedStringAtom" }],
      [/:'/, { token: "constant.delimiter", next: "@singleQuotedStringAtom" }]
    ],
    doubleQuotedStringAtom: [
      [/"/, { token: "constant.delimiter", next: "@pop" }],
      { include: "@stringConstantContentInterpol" }
    ],
    singleQuotedStringAtom: [
      [/'/, { token: "constant.delimiter", next: "@pop" }],
      { include: "@stringConstantContentInterpol" }
    ],
    sigils: [
      [/~[a-z]@sigilStartDelimiter/, { token: "@rematch", next: "@sigil.interpol" }],
      [/~([A-Z]+)@sigilStartDelimiter/, { token: "@rematch", next: "@sigil.noInterpol" }]
    ],
    sigil: [
      [/~([a-z]|[A-Z]+)\{/, { token: "@rematch", switchTo: "@sigilStart.$S2.$1.{.}" }],
      [/~([a-z]|[A-Z]+)\[/, { token: "@rematch", switchTo: "@sigilStart.$S2.$1.[.]" }],
      [/~([a-z]|[A-Z]+)\(/, { token: "@rematch", switchTo: "@sigilStart.$S2.$1.(.)" }],
      [/~([a-z]|[A-Z]+)\</, { token: "@rematch", switchTo: "@sigilStart.$S2.$1.<.>" }],
      [
        /~([a-z]|[A-Z]+)(@sigilSymmetricDelimiter)/,
        { token: "@rematch", switchTo: "@sigilStart.$S2.$1.$2.$2" }
      ]
    ],
    "sigilStart.interpol.s": [
      [
        /~s@sigilStartDelimiter/,
        {
          token: "string.delimiter",
          switchTo: "@sigilContinue.$S2.$S3.$S4.$S5"
        }
      ]
    ],
    "sigilContinue.interpol.s": [
      [
        /(@sigilEndDelimiter)@sigilModifiers/,
        {
          cases: {
            "$1==$S5": { token: "string.delimiter", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      { include: "@stringContentInterpol" }
    ],
    "sigilStart.noInterpol.S": [
      [
        /~S@sigilStartDelimiter/,
        {
          token: "string.delimiter",
          switchTo: "@sigilContinue.$S2.$S3.$S4.$S5"
        }
      ]
    ],
    "sigilContinue.noInterpol.S": [
      [/(^|[^\\])\\@sigilEndDelimiter/, "string"],
      [
        /(@sigilEndDelimiter)@sigilModifiers/,
        {
          cases: {
            "$1==$S5": { token: "string.delimiter", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      { include: "@stringContent" }
    ],
    "sigilStart.interpol.r": [
      [
        /~r@sigilStartDelimiter/,
        {
          token: "regexp.delimiter",
          switchTo: "@sigilContinue.$S2.$S3.$S4.$S5"
        }
      ]
    ],
    "sigilContinue.interpol.r": [
      [
        /(@sigilEndDelimiter)@sigilModifiers/,
        {
          cases: {
            "$1==$S5": { token: "regexp.delimiter", next: "@pop" },
            "@default": "regexp"
          }
        }
      ],
      { include: "@regexpContentInterpol" }
    ],
    "sigilStart.noInterpol.R": [
      [
        /~R@sigilStartDelimiter/,
        {
          token: "regexp.delimiter",
          switchTo: "@sigilContinue.$S2.$S3.$S4.$S5"
        }
      ]
    ],
    "sigilContinue.noInterpol.R": [
      [/(^|[^\\])\\@sigilEndDelimiter/, "regexp"],
      [
        /(@sigilEndDelimiter)@sigilModifiers/,
        {
          cases: {
            "$1==$S5": { token: "regexp.delimiter", next: "@pop" },
            "@default": "regexp"
          }
        }
      ],
      { include: "@regexpContent" }
    ],
    "sigilStart.interpol": [
      [
        /~([a-z]|[A-Z]+)@sigilStartDelimiter/,
        {
          token: "sigil.delimiter",
          switchTo: "@sigilContinue.$S2.$S3.$S4.$S5"
        }
      ]
    ],
    "sigilContinue.interpol": [
      [
        /(@sigilEndDelimiter)@sigilModifiers/,
        {
          cases: {
            "$1==$S5": { token: "sigil.delimiter", next: "@pop" },
            "@default": "sigil"
          }
        }
      ],
      { include: "@sigilContentInterpol" }
    ],
    "sigilStart.noInterpol": [
      [
        /~([a-z]|[A-Z]+)@sigilStartDelimiter/,
        {
          token: "sigil.delimiter",
          switchTo: "@sigilContinue.$S2.$S3.$S4.$S5"
        }
      ]
    ],
    "sigilContinue.noInterpol": [
      [/(^|[^\\])\\@sigilEndDelimiter/, "sigil"],
      [
        /(@sigilEndDelimiter)@sigilModifiers/,
        {
          cases: {
            "$1==$S5": { token: "sigil.delimiter", next: "@pop" },
            "@default": "sigil"
          }
        }
      ],
      { include: "@sigilContent" }
    ],
    attributes: [
      [
        /\@(module|type)?doc (~[sS])?"""/,
        {
          token: "comment.block.documentation",
          next: "@doubleQuotedHeredocDocstring"
        }
      ],
      [
        /\@(module|type)?doc (~[sS])?'''/,
        {
          token: "comment.block.documentation",
          next: "@singleQuotedHeredocDocstring"
        }
      ],
      [
        /\@(module|type)?doc (~[sS])?"/,
        {
          token: "comment.block.documentation",
          next: "@doubleQuotedStringDocstring"
        }
      ],
      [
        /\@(module|type)?doc (~[sS])?'/,
        {
          token: "comment.block.documentation",
          next: "@singleQuotedStringDocstring"
        }
      ],
      [/\@(module|type)?doc false/, "comment.block.documentation"],
      [/\@(@variableName)/, "variable"]
    ],
    doubleQuotedHeredocDocstring: [
      [/"""/, { token: "comment.block.documentation", next: "@pop" }],
      { include: "@docstringContent" }
    ],
    singleQuotedHeredocDocstring: [
      [/'''/, { token: "comment.block.documentation", next: "@pop" }],
      { include: "@docstringContent" }
    ],
    doubleQuotedStringDocstring: [
      [/"/, { token: "comment.block.documentation", next: "@pop" }],
      { include: "@docstringContent" }
    ],
    singleQuotedStringDocstring: [
      [/'/, { token: "comment.block.documentation", next: "@pop" }],
      { include: "@docstringContent" }
    ],
    symbols: [
      [/\?(\\.|[^\\\s])/, "number.constant"],
      [/&\d+/, "operator"],
      [/<<<|>>>/, "operator"],
      [/[()\[\]\{\}]|<<|>>/, "@brackets"],
      [/\.\.\./, "identifier"],
      [/=>/, "punctuation"],
      [/@operator/, "operator"],
      [/[:;,.%]/, "punctuation"]
    ],
    stringContentInterpol: [
      { include: "@interpolation" },
      { include: "@escapeChar" },
      { include: "@stringContent" }
    ],
    stringContent: [[/./, "string"]],
    stringConstantContentInterpol: [
      { include: "@interpolation" },
      { include: "@escapeChar" },
      { include: "@stringConstantContent" }
    ],
    stringConstantContent: [[/./, "constant"]],
    regexpContentInterpol: [
      { include: "@interpolation" },
      { include: "@escapeChar" },
      { include: "@regexpContent" }
    ],
    regexpContent: [
      [/(\s)(#)(\s.*)$/, ["white", "comment.punctuation", "comment"]],
      [/./, "regexp"]
    ],
    sigilContentInterpol: [
      { include: "@interpolation" },
      { include: "@escapeChar" },
      { include: "@sigilContent" }
    ],
    sigilContent: [[/./, "sigil"]],
    docstringContent: [[/./, "comment.block.documentation"]],
    escapeChar: [[/@escape/, "constant.character.escape"]],
    interpolation: [[/#{/, { token: "delimiter.bracket.embed", next: "@interpolationContinue" }]],
    interpolationContinue: [
      [/}/, { token: "delimiter.bracket.embed", next: "@pop" }],
      { include: "@root" }
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/flow9/flow9.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/flow9/flow9.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/flow9/flow9.ts
var conf = {
  comments: {
    blockComment: ["/*", "*/"],
    lineComment: "//"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string"] },
    { open: "[", close: "]", notIn: ["string"] },
    { open: "(", close: ")", notIn: ["string"] },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "'", close: "'", notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "<", close: ">" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".flow",
  keywords: [
    "import",
    "require",
    "export",
    "forbid",
    "native",
    "if",
    "else",
    "cast",
    "unsafe",
    "switch",
    "default"
  ],
  types: [
    "io",
    "mutable",
    "bool",
    "int",
    "double",
    "string",
    "flow",
    "void",
    "ref",
    "true",
    "false",
    "with"
  ],
  operators: [
    "=",
    ">",
    "<",
    "<=",
    ">=",
    "==",
    "!",
    "!=",
    ":=",
    "::=",
    "&&",
    "||",
    "+",
    "-",
    "*",
    "/",
    "@",
    "&",
    "%",
    ":",
    "->",
    "\\",
    "$",
    "??",
    "^"
  ],
  symbols: /[@$=><!~?:&|+\-*\\\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": "keyword",
            "@types": "type",
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "delimiter"],
      [/[<>](?!@symbols)/, "delimiter"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/((0(x|X)[0-9a-fA-F]*)|(([0-9]+\.?[0-9]*)|(\.[0-9]+))((e|E)(\+|-)?[0-9]+)?)/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/freemarker2/freemarker2.js":
/*!**************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/freemarker2/freemarker2.js ***!
  \**************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   TagAngleInterpolationBracket: function() { return /* binding */ TagAngleInterpolationBracket; },
/* harmony export */   TagAngleInterpolationDollar: function() { return /* binding */ TagAngleInterpolationDollar; },
/* harmony export */   TagAutoInterpolationBracket: function() { return /* binding */ TagAutoInterpolationBracket; },
/* harmony export */   TagAutoInterpolationDollar: function() { return /* binding */ TagAutoInterpolationDollar; },
/* harmony export */   TagBracketInterpolationBracket: function() { return /* binding */ TagBracketInterpolationBracket; },
/* harmony export */   TagBracketInterpolationDollar: function() { return /* binding */ TagBracketInterpolationDollar; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/freemarker2/freemarker2.ts
var EMPTY_ELEMENTS = [
  "assign",
  "flush",
  "ftl",
  "return",
  "global",
  "import",
  "include",
  "break",
  "continue",
  "local",
  "nested",
  "nt",
  "setting",
  "stop",
  "t",
  "lt",
  "rt",
  "fallback"
];
var BLOCK_ELEMENTS = [
  "attempt",
  "autoesc",
  "autoEsc",
  "compress",
  "comment",
  "escape",
  "noescape",
  "function",
  "if",
  "list",
  "items",
  "sep",
  "macro",
  "noparse",
  "noParse",
  "noautoesc",
  "noAutoEsc",
  "outputformat",
  "switch",
  "visit",
  "recurse"
];
var TagSyntaxAngle = {
  close: ">",
  id: "angle",
  open: "<"
};
var TagSyntaxBracket = {
  close: "\\]",
  id: "bracket",
  open: "\\["
};
var TagSyntaxAuto = {
  close: "[>\\]]",
  id: "auto",
  open: "[<\\[]"
};
var InterpolationSyntaxDollar = {
  close: "\\}",
  id: "dollar",
  open1: "\\$",
  open2: "\\{"
};
var InterpolationSyntaxBracket = {
  close: "\\]",
  id: "bracket",
  open1: "\\[",
  open2: "="
};
function createLangConfiguration(ts) {
  return {
    brackets: [
      ["<", ">"],
      ["[", "]"],
      ["(", ")"],
      ["{", "}"]
    ],
    comments: {
      blockComment: [`${ts.open}--`, `--${ts.close}`]
    },
    autoCloseBefore: "\n\r	 }]),.:;=",
    autoClosingPairs: [
      { open: "{", close: "}" },
      { open: "[", close: "]" },
      { open: "(", close: ")" },
      { open: '"', close: '"', notIn: ["string"] },
      { open: "'", close: "'", notIn: ["string"] }
    ],
    surroundingPairs: [
      { open: '"', close: '"' },
      { open: "'", close: "'" },
      { open: "{", close: "}" },
      { open: "[", close: "]" },
      { open: "(", close: ")" },
      { open: "<", close: ">" }
    ],
    folding: {
      markers: {
        start: new RegExp(`${ts.open}#(?:${BLOCK_ELEMENTS.join("|")})([^/${ts.close}]*(?!/)${ts.close})[^${ts.open}]*$`),
        end: new RegExp(`${ts.open}/#(?:${BLOCK_ELEMENTS.join("|")})[\\r\\n\\t ]*>`)
      }
    },
    onEnterRules: [
      {
        beforeText: new RegExp(`${ts.open}#(?!(?:${EMPTY_ELEMENTS.join("|")}))([a-zA-Z_]+)([^/${ts.close}]*(?!/)${ts.close})[^${ts.open}]*$`),
        afterText: new RegExp(`^${ts.open}/#([a-zA-Z_]+)[\\r\\n\\t ]*${ts.close}$`),
        action: {
          indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
        }
      },
      {
        beforeText: new RegExp(`${ts.open}#(?!(?:${EMPTY_ELEMENTS.join("|")}))([a-zA-Z_]+)([^/${ts.close}]*(?!/)${ts.close})[^${ts.open}]*$`),
        action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
      }
    ]
  };
}
function createLangConfigurationAuto() {
  return {
    brackets: [
      ["<", ">"],
      ["[", "]"],
      ["(", ")"],
      ["{", "}"]
    ],
    autoCloseBefore: "\n\r	 }]),.:;=",
    autoClosingPairs: [
      { open: "{", close: "}" },
      { open: "[", close: "]" },
      { open: "(", close: ")" },
      { open: '"', close: '"', notIn: ["string"] },
      { open: "'", close: "'", notIn: ["string"] }
    ],
    surroundingPairs: [
      { open: '"', close: '"' },
      { open: "'", close: "'" },
      { open: "{", close: "}" },
      { open: "[", close: "]" },
      { open: "(", close: ")" },
      { open: "<", close: ">" }
    ],
    folding: {
      markers: {
        start: new RegExp(`[<\\[]#(?:${BLOCK_ELEMENTS.join("|")})([^/>\\]]*(?!/)[>\\]])[^<\\[]*$`),
        end: new RegExp(`[<\\[]/#(?:${BLOCK_ELEMENTS.join("|")})[\\r\\n\\t ]*>`)
      }
    },
    onEnterRules: [
      {
        beforeText: new RegExp(`[<\\[]#(?!(?:${EMPTY_ELEMENTS.join("|")}))([a-zA-Z_]+)([^/>\\]]*(?!/)[>\\]])[^[<\\[]]*$`),
        afterText: new RegExp(`^[<\\[]/#([a-zA-Z_]+)[\\r\\n\\t ]*[>\\]]$`),
        action: {
          indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
        }
      },
      {
        beforeText: new RegExp(`[<\\[]#(?!(?:${EMPTY_ELEMENTS.join("|")}))([a-zA-Z_]+)([^/>\\]]*(?!/)[>\\]])[^[<\\[]]*$`),
        action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
      }
    ]
  };
}
function createMonarchLanguage(ts, is) {
  const id = `_${ts.id}_${is.id}`;
  const s = (name) => name.replace(/__id__/g, id);
  const r = (regexp) => {
    const source = regexp.source.replace(/__id__/g, id);
    return new RegExp(source, regexp.flags);
  };
  return {
    unicode: true,
    includeLF: false,
    start: s("default__id__"),
    ignoreCase: false,
    defaultToken: "invalid",
    tokenPostfix: `.freemarker2`,
    brackets: [
      { open: "{", close: "}", token: "delimiter.curly" },
      { open: "[", close: "]", token: "delimiter.square" },
      { open: "(", close: ")", token: "delimiter.parenthesis" },
      { open: "<", close: ">", token: "delimiter.angle" }
    ],
    [s("open__id__")]: new RegExp(ts.open),
    [s("close__id__")]: new RegExp(ts.close),
    [s("iOpen1__id__")]: new RegExp(is.open1),
    [s("iOpen2__id__")]: new RegExp(is.open2),
    [s("iClose__id__")]: new RegExp(is.close),
    [s("startTag__id__")]: r(/(@open__id__)(#)/),
    [s("endTag__id__")]: r(/(@open__id__)(\/#)/),
    [s("startOrEndTag__id__")]: r(/(@open__id__)(\/?#)/),
    [s("closeTag1__id__")]: r(/((?:@blank)*)(@close__id__)/),
    [s("closeTag2__id__")]: r(/((?:@blank)*\/?)(@close__id__)/),
    blank: /[ \t\n\r]/,
    keywords: ["false", "true", "in", "as", "using"],
    directiveStartCloseTag1: /attempt|recover|sep|auto[eE]sc|no(?:autoe|AutoE)sc|compress|default|no[eE]scape|comment|no[pP]arse/,
    directiveStartCloseTag2: /else|break|continue|return|stop|flush|t|lt|rt|nt|nested|recurse|fallback|ftl/,
    directiveStartBlank: /if|else[iI]f|list|for[eE]ach|switch|case|assign|global|local|include|import|function|macro|transform|visit|stop|return|call|setting|output[fF]ormat|nested|recurse|escape|ftl|items/,
    directiveEndCloseTag1: /if|list|items|sep|recover|attempt|for[eE]ach|local|global|assign|function|macro|output[fF]ormat|auto[eE]sc|no(?:autoe|AutoE)sc|compress|transform|switch|escape|no[eE]scape/,
    escapedChar: /\\(?:[ntrfbgla\\'"\{=]|(?:x[0-9A-Fa-f]{1,4}))/,
    asciiDigit: /[0-9]/,
    integer: /[0-9]+/,
    nonEscapedIdStartChar: /[\$@-Z_a-z\u00AA\u00B5\u00BA\u00C0-\u00D6\u00D8-\u00F6\u00F8-\u1FFF\u2071\u207F\u2090-\u209C\u2102\u2107\u210A-\u2113\u2115\u2119-\u211D\u2124\u2126\u2128\u212A-\u212D\u212F-\u2139\u213C-\u213F\u2145-\u2149\u214E\u2183-\u2184\u2C00-\u2C2E\u2C30-\u2C5E\u2C60-\u2CE4\u2CEB-\u2CEE\u2CF2-\u2CF3\u2D00-\u2D25\u2D27\u2D2D\u2D30-\u2D67\u2D6F\u2D80-\u2D96\u2DA0-\u2DA6\u2DA8-\u2DAE\u2DB0-\u2DB6\u2DB8-\u2DBE\u2DC0-\u2DC6\u2DC8-\u2DCE\u2DD0-\u2DD6\u2DD8-\u2DDE\u2E2F\u3005-\u3006\u3031-\u3035\u303B-\u303C\u3040-\u318F\u31A0-\u31BA\u31F0-\u31FF\u3300-\u337F\u3400-\u4DB5\u4E00-\uA48C\uA4D0-\uA4FD\uA500-\uA60C\uA610-\uA62B\uA640-\uA66E\uA67F-\uA697\uA6A0-\uA6E5\uA717-\uA71F\uA722-\uA788\uA78B-\uA78E\uA790-\uA793\uA7A0-\uA7AA\uA7F8-\uA801\uA803-\uA805\uA807-\uA80A\uA80C-\uA822\uA840-\uA873\uA882-\uA8B3\uA8D0-\uA8D9\uA8F2-\uA8F7\uA8FB\uA900-\uA925\uA930-\uA946\uA960-\uA97C\uA984-\uA9B2\uA9CF-\uA9D9\uAA00-\uAA28\uAA40-\uAA42\uAA44-\uAA4B\uAA50-\uAA59\uAA60-\uAA76\uAA7A\uAA80-\uAAAF\uAAB1\uAAB5-\uAAB6\uAAB9-\uAABD\uAAC0\uAAC2\uAADB-\uAADD\uAAE0-\uAAEA\uAAF2-\uAAF4\uAB01-\uAB06\uAB09-\uAB0E\uAB11-\uAB16\uAB20-\uAB26\uAB28-\uAB2E\uABC0-\uABE2\uABF0-\uABF9\uAC00-\uD7A3\uD7B0-\uD7C6\uD7CB-\uD7FB\uF900-\uFB06\uFB13-\uFB17\uFB1D\uFB1F-\uFB28\uFB2A-\uFB36\uFB38-\uFB3C\uFB3E\uFB40-\uFB41\uFB43-\uFB44\uFB46-\uFBB1\uFBD3-\uFD3D\uFD50-\uFD8F\uFD92-\uFDC7\uFDF0-\uFDFB\uFE70-\uFE74\uFE76-\uFEFC\uFF10-\uFF19\uFF21-\uFF3A\uFF41-\uFF5A\uFF66-\uFFBE\uFFC2-\uFFC7\uFFCA-\uFFCF\uFFD2-\uFFD7\uFFDA-\uFFDC]/,
    escapedIdChar: /\\[\-\.:#]/,
    idStartChar: /(?:@nonEscapedIdStartChar)|(?:@escapedIdChar)/,
    id: /(?:@idStartChar)(?:(?:@idStartChar)|(?:@asciiDigit))*/,
    specialHashKeys: /\*\*|\*|false|true|in|as|using/,
    namedSymbols: /&lt;=|&gt;=|\\lte|\\lt|&lt;|\\gte|\\gt|&gt;|&amp;&amp;|\\and|-&gt;|->|==|!=|\+=|-=|\*=|\/=|%=|\+\+|--|<=|&&|\|\||:|\.\.\.|\.\.\*|\.\.<|\.\.!|\?\?|=|<|\+|-|\*|\/|%|\||\.\.|\?|!|&|\.|,|;/,
    arrows: ["->", "-&gt;"],
    delimiters: [";", ":", ",", "."],
    stringOperators: ["lte", "lt", "gte", "gt"],
    noParseTags: ["noparse", "noParse", "comment"],
    tokenizer: {
      [s("default__id__")]: [
        { include: s("@directive_token__id__") },
        { include: s("@interpolation_and_text_token__id__") }
      ],
      [s("fmExpression__id__.directive")]: [
        { include: s("@blank_and_expression_comment_token__id__") },
        { include: s("@directive_end_token__id__") },
        { include: s("@expression_token__id__") }
      ],
      [s("fmExpression__id__.interpolation")]: [
        { include: s("@blank_and_expression_comment_token__id__") },
        { include: s("@expression_token__id__") },
        { include: s("@greater_operators_token__id__") }
      ],
      [s("inParen__id__.plain")]: [
        { include: s("@blank_and_expression_comment_token__id__") },
        { include: s("@directive_end_token__id__") },
        { include: s("@expression_token__id__") }
      ],
      [s("inParen__id__.gt")]: [
        { include: s("@blank_and_expression_comment_token__id__") },
        { include: s("@expression_token__id__") },
        { include: s("@greater_operators_token__id__") }
      ],
      [s("noSpaceExpression__id__")]: [
        { include: s("@no_space_expression_end_token__id__") },
        { include: s("@directive_end_token__id__") },
        { include: s("@expression_token__id__") }
      ],
      [s("unifiedCall__id__")]: [{ include: s("@unified_call_token__id__") }],
      [s("singleString__id__")]: [{ include: s("@string_single_token__id__") }],
      [s("doubleString__id__")]: [{ include: s("@string_double_token__id__") }],
      [s("rawSingleString__id__")]: [{ include: s("@string_single_raw_token__id__") }],
      [s("rawDoubleString__id__")]: [{ include: s("@string_double_raw_token__id__") }],
      [s("expressionComment__id__")]: [{ include: s("@expression_comment_token__id__") }],
      [s("noParse__id__")]: [{ include: s("@no_parse_token__id__") }],
      [s("terseComment__id__")]: [{ include: s("@terse_comment_token__id__") }],
      [s("directive_token__id__")]: [
        [
          r(/(?:@startTag__id__)(@directiveStartCloseTag1)(?:@closeTag1__id__)/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : [
            { token: "@brackets.directive" },
            { token: "delimiter.directive" },
            {
              cases: {
                "@noParseTags": { token: "tag", next: s("@noParse__id__.$3") },
                "@default": { token: "tag" }
              }
            },
            { token: "delimiter.directive" },
            { token: "@brackets.directive" }
          ]
        ],
        [
          r(/(?:@startTag__id__)(@directiveStartCloseTag2)(?:@closeTag2__id__)/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : [
            { token: "@brackets.directive" },
            { token: "delimiter.directive" },
            { token: "tag" },
            { token: "delimiter.directive" },
            { token: "@brackets.directive" }
          ]
        ],
        [
          r(/(?:@startTag__id__)(@directiveStartBlank)(@blank)/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : [
            { token: "@brackets.directive" },
            { token: "delimiter.directive" },
            { token: "tag" },
            { token: "", next: s("@fmExpression__id__.directive") }
          ]
        ],
        [
          r(/(?:@endTag__id__)(@directiveEndCloseTag1)(?:@closeTag1__id__)/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : [
            { token: "@brackets.directive" },
            { token: "delimiter.directive" },
            { token: "tag" },
            { token: "delimiter.directive" },
            { token: "@brackets.directive" }
          ]
        ],
        [
          r(/(@open__id__)(@)/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : [
            { token: "@brackets.directive" },
            { token: "delimiter.directive", next: s("@unifiedCall__id__") }
          ]
        ],
        [
          r(/(@open__id__)(\/@)((?:(?:@id)(?:\.(?:@id))*)?)(?:@closeTag1__id__)/),
          [
            { token: "@brackets.directive" },
            { token: "delimiter.directive" },
            { token: "tag" },
            { token: "delimiter.directive" },
            { token: "@brackets.directive" }
          ]
        ],
        [
          r(/(@open__id__)#--/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : { token: "comment", next: s("@terseComment__id__") }
        ],
        [
          r(/(?:@startOrEndTag__id__)([a-zA-Z_]+)/),
          ts.id === "auto" ? {
            cases: {
              "$1==<": { token: "@rematch", switchTo: `@default_angle_${is.id}` },
              "$1==[": { token: "@rematch", switchTo: `@default_bracket_${is.id}` }
            }
          } : [
            { token: "@brackets.directive" },
            { token: "delimiter.directive" },
            { token: "tag.invalid", next: s("@fmExpression__id__.directive") }
          ]
        ]
      ],
      [s("interpolation_and_text_token__id__")]: [
        [
          r(/(@iOpen1__id__)(@iOpen2__id__)/),
          [
            { token: is.id === "bracket" ? "@brackets.interpolation" : "delimiter.interpolation" },
            {
              token: is.id === "bracket" ? "delimiter.interpolation" : "@brackets.interpolation",
              next: s("@fmExpression__id__.interpolation")
            }
          ]
        ],
        [/[\$#<\[\{]|(?:@blank)+|[^\$<#\[\{\n\r\t ]+/, { token: "source" }]
      ],
      [s("string_single_token__id__")]: [
        [/[^'\\]/, { token: "string" }],
        [/@escapedChar/, { token: "string.escape" }],
        [/'/, { token: "string", next: "@pop" }]
      ],
      [s("string_double_token__id__")]: [
        [/[^"\\]/, { token: "string" }],
        [/@escapedChar/, { token: "string.escape" }],
        [/"/, { token: "string", next: "@pop" }]
      ],
      [s("string_single_raw_token__id__")]: [
        [/[^']+/, { token: "string.raw" }],
        [/'/, { token: "string.raw", next: "@pop" }]
      ],
      [s("string_double_raw_token__id__")]: [
        [/[^"]+/, { token: "string.raw" }],
        [/"/, { token: "string.raw", next: "@pop" }]
      ],
      [s("expression_token__id__")]: [
        [
          /(r?)(['"])/,
          {
            cases: {
              "r'": [
                { token: "keyword" },
                { token: "string.raw", next: s("@rawSingleString__id__") }
              ],
              'r"': [
                { token: "keyword" },
                { token: "string.raw", next: s("@rawDoubleString__id__") }
              ],
              "'": [{ token: "source" }, { token: "string", next: s("@singleString__id__") }],
              '"': [{ token: "source" }, { token: "string", next: s("@doubleString__id__") }]
            }
          }
        ],
        [
          /(?:@integer)(?:\.(?:@integer))?/,
          {
            cases: {
              "(?:@integer)": { token: "number" },
              "@default": { token: "number.float" }
            }
          }
        ],
        [
          /(\.)(@blank*)(@specialHashKeys)/,
          [{ token: "delimiter" }, { token: "" }, { token: "identifier" }]
        ],
        [
          /(?:@namedSymbols)/,
          {
            cases: {
              "@arrows": { token: "meta.arrow" },
              "@delimiters": { token: "delimiter" },
              "@default": { token: "operators" }
            }
          }
        ],
        [
          /@id/,
          {
            cases: {
              "@keywords": { token: "keyword.$0" },
              "@stringOperators": { token: "operators" },
              "@default": { token: "identifier" }
            }
          }
        ],
        [
          /[\[\]\(\)\{\}]/,
          {
            cases: {
              "\\[": {
                cases: {
                  "$S2==gt": { token: "@brackets", next: s("@inParen__id__.gt") },
                  "@default": { token: "@brackets", next: s("@inParen__id__.plain") }
                }
              },
              "\\]": {
                cases: {
                  ...is.id === "bracket" ? {
                    "$S2==interpolation": { token: "@brackets.interpolation", next: "@popall" }
                  } : {},
                  ...ts.id === "bracket" ? {
                    "$S2==directive": { token: "@brackets.directive", next: "@popall" }
                  } : {},
                  [s("$S1==inParen__id__")]: { token: "@brackets", next: "@pop" },
                  "@default": { token: "@brackets" }
                }
              },
              "\\(": { token: "@brackets", next: s("@inParen__id__.gt") },
              "\\)": {
                cases: {
                  [s("$S1==inParen__id__")]: { token: "@brackets", next: "@pop" },
                  "@default": { token: "@brackets" }
                }
              },
              "\\{": {
                cases: {
                  "$S2==gt": { token: "@brackets", next: s("@inParen__id__.gt") },
                  "@default": { token: "@brackets", next: s("@inParen__id__.plain") }
                }
              },
              "\\}": {
                cases: {
                  ...is.id === "bracket" ? {} : {
                    "$S2==interpolation": { token: "@brackets.interpolation", next: "@popall" }
                  },
                  [s("$S1==inParen__id__")]: { token: "@brackets", next: "@pop" },
                  "@default": { token: "@brackets" }
                }
              }
            }
          }
        ],
        [/\$\{/, { token: "delimiter.invalid" }]
      ],
      [s("blank_and_expression_comment_token__id__")]: [
        [/(?:@blank)+/, { token: "" }],
        [/[<\[][#!]--/, { token: "comment", next: s("@expressionComment__id__") }]
      ],
      [s("directive_end_token__id__")]: [
        [
          />/,
          ts.id === "bracket" ? { token: "operators" } : { token: "@brackets.directive", next: "@popall" }
        ],
        [
          r(/(\/)(@close__id__)/),
          [{ token: "delimiter.directive" }, { token: "@brackets.directive", next: "@popall" }]
        ]
      ],
      [s("greater_operators_token__id__")]: [
        [/>/, { token: "operators" }],
        [/>=/, { token: "operators" }]
      ],
      [s("no_space_expression_end_token__id__")]: [
        [/(?:@blank)+/, { token: "", switchTo: s("@fmExpression__id__.directive") }]
      ],
      [s("unified_call_token__id__")]: [
        [
          /(@id)((?:@blank)+)/,
          [{ token: "tag" }, { token: "", next: s("@fmExpression__id__.directive") }]
        ],
        [
          r(/(@id)(\/?)(@close__id__)/),
          [
            { token: "tag" },
            { token: "delimiter.directive" },
            { token: "@brackets.directive", next: "@popall" }
          ]
        ],
        [/./, { token: "@rematch", next: s("@noSpaceExpression__id__") }]
      ],
      [s("no_parse_token__id__")]: [
        [
          r(/(@open__id__)(\/#?)([a-zA-Z]+)((?:@blank)*)(@close__id__)/),
          {
            cases: {
              "$S2==$3": [
                { token: "@brackets.directive" },
                { token: "delimiter.directive" },
                { token: "tag" },
                { token: "" },
                { token: "@brackets.directive", next: "@popall" }
              ],
              "$S2==comment": [
                { token: "comment" },
                { token: "comment" },
                { token: "comment" },
                { token: "comment" },
                { token: "comment" }
              ],
              "@default": [
                { token: "source" },
                { token: "source" },
                { token: "source" },
                { token: "source" },
                { token: "source" }
              ]
            }
          }
        ],
        [
          /[^<\[\-]+|[<\[\-]/,
          {
            cases: {
              "$S2==comment": { token: "comment" },
              "@default": { token: "source" }
            }
          }
        ]
      ],
      [s("expression_comment_token__id__")]: [
        [
          /--[>\]]/,
          {
            token: "comment",
            next: "@pop"
          }
        ],
        [/[^\->\]]+|[>\]\-]/, { token: "comment" }]
      ],
      [s("terse_comment_token__id__")]: [
        [r(/--(?:@close__id__)/), { token: "comment", next: "@popall" }],
        [/[^<\[\-]+|[<\[\-]/, { token: "comment" }]
      ]
    }
  };
}
function createMonarchLanguageAuto(is) {
  const angle = createMonarchLanguage(TagSyntaxAngle, is);
  const bracket = createMonarchLanguage(TagSyntaxBracket, is);
  const auto = createMonarchLanguage(TagSyntaxAuto, is);
  return {
    ...angle,
    ...bracket,
    ...auto,
    unicode: true,
    includeLF: false,
    start: `default_auto_${is.id}`,
    ignoreCase: false,
    defaultToken: "invalid",
    tokenPostfix: `.freemarker2`,
    brackets: [
      { open: "{", close: "}", token: "delimiter.curly" },
      { open: "[", close: "]", token: "delimiter.square" },
      { open: "(", close: ")", token: "delimiter.parenthesis" },
      { open: "<", close: ">", token: "delimiter.angle" }
    ],
    tokenizer: {
      ...angle.tokenizer,
      ...bracket.tokenizer,
      ...auto.tokenizer
    }
  };
}
var TagAngleInterpolationDollar = {
  conf: createLangConfiguration(TagSyntaxAngle),
  language: createMonarchLanguage(TagSyntaxAngle, InterpolationSyntaxDollar)
};
var TagBracketInterpolationDollar = {
  conf: createLangConfiguration(TagSyntaxBracket),
  language: createMonarchLanguage(TagSyntaxBracket, InterpolationSyntaxDollar)
};
var TagAngleInterpolationBracket = {
  conf: createLangConfiguration(TagSyntaxAngle),
  language: createMonarchLanguage(TagSyntaxAngle, InterpolationSyntaxBracket)
};
var TagBracketInterpolationBracket = {
  conf: createLangConfiguration(TagSyntaxBracket),
  language: createMonarchLanguage(TagSyntaxBracket, InterpolationSyntaxBracket)
};
var TagAutoInterpolationDollar = {
  conf: createLangConfigurationAuto(),
  language: createMonarchLanguageAuto(InterpolationSyntaxDollar)
};
var TagAutoInterpolationBracket = {
  conf: createLangConfigurationAuto(),
  language: createMonarchLanguageAuto(InterpolationSyntaxBracket)
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/fsharp/fsharp.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/fsharp/fsharp.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/fsharp/fsharp.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["(*", "*)"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*//\\s*#region\\b|^\\s*\\(\\*\\s*#region(.*)\\*\\)"),
      end: new RegExp("^\\s*//\\s*#endregion\\b|^\\s*\\(\\*\\s*#endregion\\s*\\*\\)")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".fs",
  keywords: [
    "abstract",
    "and",
    "atomic",
    "as",
    "assert",
    "asr",
    "base",
    "begin",
    "break",
    "checked",
    "component",
    "const",
    "constraint",
    "constructor",
    "continue",
    "class",
    "default",
    "delegate",
    "do",
    "done",
    "downcast",
    "downto",
    "elif",
    "else",
    "end",
    "exception",
    "eager",
    "event",
    "external",
    "extern",
    "false",
    "finally",
    "for",
    "fun",
    "function",
    "fixed",
    "functor",
    "global",
    "if",
    "in",
    "include",
    "inherit",
    "inline",
    "interface",
    "internal",
    "land",
    "lor",
    "lsl",
    "lsr",
    "lxor",
    "lazy",
    "let",
    "match",
    "member",
    "mod",
    "module",
    "mutable",
    "namespace",
    "method",
    "mixin",
    "new",
    "not",
    "null",
    "of",
    "open",
    "or",
    "object",
    "override",
    "private",
    "parallel",
    "process",
    "protected",
    "pure",
    "public",
    "rec",
    "return",
    "static",
    "sealed",
    "struct",
    "sig",
    "then",
    "to",
    "true",
    "tailcall",
    "trait",
    "try",
    "type",
    "upcast",
    "use",
    "val",
    "void",
    "virtual",
    "volatile",
    "when",
    "while",
    "with",
    "yield"
  ],
  symbols: /[=><!~?:&|+\-*\^%;\.,\/]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  integersuffix: /[uU]?[yslnLI]?/,
  floatsuffix: /[fFmM]?/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/\[<.*>\]/, "annotation"],
      [/^#(if|else|endif)/, "keyword"],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/\d*\d+[eE]([\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/\d*\.\d+([eE][\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/0x[0-9a-fA-F]+LF/, "number.float"],
      [/0x[0-9a-fA-F]+(@integersuffix)/, "number.hex"],
      [/0b[0-1]+(@integersuffix)/, "number.bin"],
      [/\d+(@integersuffix)/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"""/, "string", '@string."""'],
      [/"/, "string", '@string."'],
      [/\@"/, { token: "string.quote", next: "@litstring" }],
      [/'[^\\']'B?/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\(\*(?!\))/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^*(]+/, "comment"],
      [/\*\)/, "comment", "@pop"],
      [/\*/, "comment"],
      [/\(\*\)/, "comment"],
      [/\(/, "comment"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [
        /("""|"B?)/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ]
    ],
    litstring: [
      [/[^"]+/, "string"],
      [/""/, "string.escape"],
      [/"/, { token: "string.quote", next: "@pop" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/go/go.js":
/*!********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/go/go.js ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/go/go.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "`", close: "`", notIn: ["string"] },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "`", close: "`" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".go",
  keywords: [
    "break",
    "case",
    "chan",
    "const",
    "continue",
    "default",
    "defer",
    "else",
    "fallthrough",
    "for",
    "func",
    "go",
    "goto",
    "if",
    "import",
    "interface",
    "map",
    "package",
    "range",
    "return",
    "select",
    "struct",
    "switch",
    "type",
    "var",
    "bool",
    "true",
    "false",
    "uint8",
    "uint16",
    "uint32",
    "uint64",
    "int8",
    "int16",
    "int32",
    "int64",
    "float32",
    "float64",
    "complex64",
    "complex128",
    "byte",
    "rune",
    "uint",
    "int",
    "uintptr",
    "string",
    "nil"
  ],
  operators: [
    "+",
    "-",
    "*",
    "/",
    "%",
    "&",
    "|",
    "^",
    "<<",
    ">>",
    "&^",
    "+=",
    "-=",
    "*=",
    "/=",
    "%=",
    "&=",
    "|=",
    "^=",
    "<<=",
    ">>=",
    "&^=",
    "&&",
    "||",
    "<-",
    "++",
    "--",
    "==",
    "<",
    ">",
    "=",
    "!",
    "!=",
    "<=",
    ">=",
    ":=",
    "...",
    "(",
    ")",
    "",
    "]",
    "{",
    "}",
    ",",
    ";",
    ".",
    ":"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/\[\[.*\]\]/, "annotation"],
      [/^\s*#\w+/, "keyword"],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\d+[eE]([\-+]?\d+)?/, "number.float"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F']*[0-9a-fA-F]/, "number.hex"],
      [/0[0-7']*[0-7]/, "number.octal"],
      [/0[bB][0-1']*[0-1]/, "number.binary"],
      [/\d[\d']*/, "number"],
      [/\d/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"],
      [/`/, "string", "@rawstring"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@doccomment"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    doccomment: [
      [/[^\/*]+/, "comment.doc"],
      [/\/\*/, "comment.doc.invalid"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    rawstring: [
      [/[^\`]/, "string"],
      [/`/, "string", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/graphql/graphql.js":
/*!******************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/graphql/graphql.js ***!
  \******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/graphql/graphql.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"""', close: '"""', notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"""', close: '"""' },
    { open: '"', close: '"' }
  ],
  folding: {
    offSide: true
  }
};
var language = {
  defaultToken: "invalid",
  tokenPostfix: ".gql",
  keywords: [
    "null",
    "true",
    "false",
    "query",
    "mutation",
    "subscription",
    "extend",
    "schema",
    "directive",
    "scalar",
    "type",
    "interface",
    "union",
    "enum",
    "input",
    "implements",
    "fragment",
    "on"
  ],
  typeKeywords: ["Int", "Float", "String", "Boolean", "ID"],
  directiveLocations: [
    "SCHEMA",
    "SCALAR",
    "OBJECT",
    "FIELD_DEFINITION",
    "ARGUMENT_DEFINITION",
    "INTERFACE",
    "UNION",
    "ENUM",
    "ENUM_VALUE",
    "INPUT_OBJECT",
    "INPUT_FIELD_DEFINITION",
    "QUERY",
    "MUTATION",
    "SUBSCRIPTION",
    "FIELD",
    "FRAGMENT_DEFINITION",
    "FRAGMENT_SPREAD",
    "INLINE_FRAGMENT",
    "VARIABLE_DEFINITION"
  ],
  operators: ["=", "!", "?", ":", "&", "|"],
  symbols: /[=!?:&|]+/,
  escapes: /\\(?:["\\\/bfnrt]|u[0-9A-Fa-f]{4})/,
  tokenizer: {
    root: [
      [
        /[a-z_][\w$]*/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "key.identifier"
          }
        }
      ],
      [
        /[$][\w$]*/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "argument.identifier"
          }
        }
      ],
      [
        /[A-Z][\w\$]*/,
        {
          cases: {
            "@typeKeywords": "keyword",
            "@default": "type.identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, { cases: { "@operators": "operator", "@default": "" } }],
      [/@\s*[a-zA-Z_\$][\w\$]*/, { token: "annotation", log: "annotation token: $0" }],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F]+/, "number.hex"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"],
      [/"""/, { token: "string", next: "@mlstring", nextEmbedded: "markdown" }],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, { token: "string.quote", bracket: "@open", next: "@string" }]
    ],
    mlstring: [
      [/[^"]+/, "string"],
      ['"""', { token: "string", next: "@pop", nextEmbedded: "@pop" }]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/#.*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/handlebars/handlebars.js":
/*!************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/handlebars/handlebars.js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/handlebars/handlebars.ts
var EMPTY_ELEMENTS = [
  "area",
  "base",
  "br",
  "col",
  "embed",
  "hr",
  "img",
  "input",
  "keygen",
  "link",
  "menuitem",
  "meta",
  "param",
  "source",
  "track",
  "wbr"
];
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\$\^\&\*\(\)\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\s]+)/g,
  comments: {
    blockComment: ["{{!--", "--}}"]
  },
  brackets: [
    ["<!--", "-->"],
    ["<", ">"],
    ["{{", "}}"],
    ["{", "}"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "<", close: ">" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  onEnterRules: [
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      afterText: /^<\/(\w[\w\d]*)\s*>$/i,
      action: {
        indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
      }
    },
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: "",
  tokenizer: {
    root: [
      [/\{\{!--/, "comment.block.start.handlebars", "@commentBlock"],
      [/\{\{!/, "comment.start.handlebars", "@comment"],
      [/\{\{/, { token: "@rematch", switchTo: "@handlebarsInSimpleState.root" }],
      [/<!DOCTYPE/, "metatag.html", "@doctype"],
      [/<!--/, "comment.html", "@commentHtml"],
      [/(<)(\w+)(\/>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/(<)(script)/, ["delimiter.html", { token: "tag.html", next: "@script" }]],
      [/(<)(style)/, ["delimiter.html", { token: "tag.html", next: "@style" }]],
      [/(<)([:\w]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/(<\/)(\w+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/</, "delimiter.html"],
      [/\{/, "delimiter.html"],
      [/[^<{]+/]
    ],
    doctype: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.comment"
        }
      ],
      [/[^>]+/, "metatag.content.html"],
      [/>/, "metatag.html", "@pop"]
    ],
    comment: [
      [/\}\}/, "comment.end.handlebars", "@pop"],
      [/./, "comment.content.handlebars"]
    ],
    commentBlock: [
      [/--\}\}/, "comment.block.end.handlebars", "@pop"],
      [/./, "comment.content.handlebars"]
    ],
    commentHtml: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.comment"
        }
      ],
      [/-->/, "comment.html", "@pop"],
      [/[^-]+/, "comment.content.html"],
      [/./, "comment.content.html"]
    ],
    otherTag: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.otherTag"
        }
      ],
      [/\/?>/, "delimiter.html", "@pop"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/]
    ],
    script: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.script"
        }
      ],
      [/type/, "attribute.name", "@scriptAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(script\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    scriptAfterType: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.scriptAfterType"
        }
      ],
      [/=/, "delimiter", "@scriptAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptAfterTypeEquals: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.scriptAfterTypeEquals"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptWithCustomType: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.scriptWithCustomType.$S2"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptEmbedded: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInEmbeddedState.scriptEmbedded.$S2",
          nextEmbedded: "@pop"
        }
      ],
      [/<\/script/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }]
    ],
    style: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.style"
        }
      ],
      [/type/, "attribute.name", "@styleAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(style\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    styleAfterType: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.styleAfterType"
        }
      ],
      [/=/, "delimiter", "@styleAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleAfterTypeEquals: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.styleAfterTypeEquals"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleWithCustomType: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInSimpleState.styleWithCustomType.$S2"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleEmbedded: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@handlebarsInEmbeddedState.styleEmbedded.$S2",
          nextEmbedded: "@pop"
        }
      ],
      [/<\/style/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }]
    ],
    handlebarsInSimpleState: [
      [/\{\{\{?/, "delimiter.handlebars"],
      [/\}\}\}?/, { token: "delimiter.handlebars", switchTo: "@$S2.$S3" }],
      { include: "handlebarsRoot" }
    ],
    handlebarsInEmbeddedState: [
      [/\{\{\{?/, "delimiter.handlebars"],
      [
        /\}\}\}?/,
        {
          token: "delimiter.handlebars",
          switchTo: "@$S2.$S3",
          nextEmbedded: "$S3"
        }
      ],
      { include: "handlebarsRoot" }
    ],
    handlebarsRoot: [
      [/"[^"]*"/, "string.handlebars"],
      [/[#/][^\s}]+/, "keyword.helper.handlebars"],
      [/else\b/, "keyword.helper.handlebars"],
      [/[\s]+/],
      [/[^}]/, "variable.parameter.handlebars"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/hcl/hcl.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/hcl/hcl.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/hcl/hcl.ts
var conf = {
  comments: {
    lineComment: "#",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"', notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".hcl",
  keywords: [
    "var",
    "local",
    "path",
    "for_each",
    "any",
    "string",
    "number",
    "bool",
    "true",
    "false",
    "null",
    "if ",
    "else ",
    "endif ",
    "for ",
    "in",
    "endfor"
  ],
  operators: [
    "=",
    ">=",
    "<=",
    "==",
    "!=",
    "+",
    "-",
    "*",
    "/",
    "%",
    "&&",
    "||",
    "!",
    "<",
    ">",
    "?",
    "...",
    ":"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  terraformFunctions: /(abs|ceil|floor|log|max|min|pow|signum|chomp|format|formatlist|indent|join|lower|regex|regexall|replace|split|strrev|substr|title|trimspace|upper|chunklist|coalesce|coalescelist|compact|concat|contains|distinct|element|flatten|index|keys|length|list|lookup|map|matchkeys|merge|range|reverse|setintersection|setproduct|setunion|slice|sort|transpose|values|zipmap|base64decode|base64encode|base64gzip|csvdecode|jsondecode|jsonencode|urlencode|yamldecode|yamlencode|abspath|dirname|pathexpand|basename|file|fileexists|fileset|filebase64|templatefile|formatdate|timeadd|timestamp|base64sha256|base64sha512|bcrypt|filebase64sha256|filebase64sha512|filemd5|filemd1|filesha256|filesha512|md5|rsadecrypt|sha1|sha256|sha512|uuid|uuidv5|cidrhost|cidrnetmask|cidrsubnet|tobool|tolist|tomap|tonumber|toset|tostring)/,
  terraformMainBlocks: /(module|data|terraform|resource|provider|variable|output|locals)/,
  tokenizer: {
    root: [
      [
        /^@terraformMainBlocks([ \t]*)([\w-]+|"[\w-]+"|)([ \t]*)([\w-]+|"[\w-]+"|)([ \t]*)(\{)/,
        ["type", "", "string", "", "string", "", "@brackets"]
      ],
      [
        /(\w+[ \t]+)([ \t]*)([\w-]+|"[\w-]+"|)([ \t]*)([\w-]+|"[\w-]+"|)([ \t]*)(\{)/,
        ["identifier", "", "string", "", "string", "", "@brackets"]
      ],
      [
        /(\w+[ \t]+)([ \t]*)([\w-]+|"[\w-]+"|)([ \t]*)([\w-]+|"[\w-]+"|)(=)(\{)/,
        ["identifier", "", "string", "", "operator", "", "@brackets"]
      ],
      { include: "@terraform" }
    ],
    terraform: [
      [/@terraformFunctions(\()/, ["type", "@brackets"]],
      [
        /[a-zA-Z_]\w*-*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "variable"
          }
        }
      ],
      { include: "@whitespace" },
      { include: "@heredoc" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "operator",
            "@default": ""
          }
        }
      ],
      [/\d*\d+[eE]([\-+]?\d+)?/, "number.float"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/\d[\d']*/, "number"],
      [/\d/, "number"],
      [/[;,.]/, "delimiter"],
      [/"/, "string", "@string"],
      [/'/, "invalid"]
    ],
    heredoc: [
      [/<<[-]*\s*["]?([\w\-]+)["]?/, { token: "string.heredoc.delimiter", next: "@heredocBody.$1" }]
    ],
    heredocBody: [
      [
        /([\w\-]+)$/,
        {
          cases: {
            "$1==$S2": [
              {
                token: "string.heredoc.delimiter",
                next: "@popall"
              }
            ],
            "@default": "string.heredoc"
          }
        }
      ],
      [/./, "string.heredoc"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"],
      [/#.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    string: [
      [/\$\{/, { token: "delimiter", next: "@stringExpression" }],
      [/[^\\"\$]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@popall"]
    ],
    stringInsideExpression: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    stringExpression: [
      [/\}/, { token: "delimiter", next: "@pop" }],
      [/"/, "string", "@stringInsideExpression"],
      { include: "@terraform" }
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/html/html.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/html/html.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/html/html.ts
var EMPTY_ELEMENTS = [
  "area",
  "base",
  "br",
  "col",
  "embed",
  "hr",
  "img",
  "input",
  "keygen",
  "link",
  "menuitem",
  "meta",
  "param",
  "source",
  "track",
  "wbr"
];
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\$\^\&\*\(\)\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\s]+)/g,
  comments: {
    blockComment: ["<!--", "-->"]
  },
  brackets: [
    ["<!--", "-->"],
    ["<", ">"],
    ["{", "}"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" }
  ],
  onEnterRules: [
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))([_:\\w][_:\\w-.\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      afterText: /^<\/([_:\w][_:\w-.\d]*)\s*>$/i,
      action: {
        indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
      }
    },
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*<!--\\s*#region\\b.*-->"),
      end: new RegExp("^\\s*<!--\\s*#endregion\\b.*-->")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".html",
  ignoreCase: true,
  tokenizer: {
    root: [
      [/<!DOCTYPE/, "metatag", "@doctype"],
      [/<!--/, "comment", "@comment"],
      [/(<)((?:[\w\-]+:)?[\w\-]+)(\s*)(\/>)/, ["delimiter", "tag", "", "delimiter"]],
      [/(<)(script)/, ["delimiter", { token: "tag", next: "@script" }]],
      [/(<)(style)/, ["delimiter", { token: "tag", next: "@style" }]],
      [/(<)((?:[\w\-]+:)?[\w\-]+)/, ["delimiter", { token: "tag", next: "@otherTag" }]],
      [/(<\/)((?:[\w\-]+:)?[\w\-]+)/, ["delimiter", { token: "tag", next: "@otherTag" }]],
      [/</, "delimiter"],
      [/[^<]+/]
    ],
    doctype: [
      [/[^>]+/, "metatag.content"],
      [/>/, "metatag", "@pop"]
    ],
    comment: [
      [/-->/, "comment", "@pop"],
      [/[^-]+/, "comment.content"],
      [/./, "comment.content"]
    ],
    otherTag: [
      [/\/?>/, "delimiter", "@pop"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/]
    ],
    script: [
      [/type/, "attribute.name", "@scriptAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter",
          next: "@scriptEmbedded",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/(<\/)(script\s*)(>)/, ["delimiter", "tag", { token: "delimiter", next: "@pop" }]]
    ],
    scriptAfterType: [
      [/=/, "delimiter", "@scriptAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter",
          next: "@scriptEmbedded",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptAfterTypeEquals: [
      [
        /"module"/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.text/javascript"
        }
      ],
      [
        /'module'/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.text/javascript"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter",
          next: "@scriptEmbedded",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptWithCustomType: [
      [
        />/,
        {
          token: "delimiter",
          next: "@scriptEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptEmbedded: [
      [/<\/script/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }],
      [/[^<]+/, ""]
    ],
    style: [
      [/type/, "attribute.name", "@styleAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter",
          next: "@styleEmbedded",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/(<\/)(style\s*)(>)/, ["delimiter", "tag", { token: "delimiter", next: "@pop" }]]
    ],
    styleAfterType: [
      [/=/, "delimiter", "@styleAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter",
          next: "@styleEmbedded",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleAfterTypeEquals: [
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter",
          next: "@styleEmbedded",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleWithCustomType: [
      [
        />/,
        {
          token: "delimiter",
          next: "@styleEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleEmbedded: [
      [/<\/style/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }],
      [/[^<]+/, ""]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/ini/ini.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/ini/ini.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/ini/ini.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".ini",
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/^\[[^\]]*\]/, "metatag"],
      [/(^\w+)(\s*)(\=)/, ["key", "", "delimiter"]],
      { include: "@whitespace" },
      [/\d+/, "number"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", '@string."'],
      [/'/, "string", "@string.'"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/^\s*[#;].*$/, "comment"]
    ],
    string: [
      [/[^\\"']+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/java/java.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/java/java.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/java/java.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\#\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "<", close: ">" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*//\\s*(?:(?:#?region\\b)|(?:<editor-fold\\b))"),
      end: new RegExp("^\\s*//\\s*(?:(?:#?endregion\\b)|(?:</editor-fold>))")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".java",
  keywords: [
    "abstract",
    "continue",
    "for",
    "new",
    "switch",
    "assert",
    "default",
    "goto",
    "package",
    "synchronized",
    "boolean",
    "do",
    "if",
    "private",
    "this",
    "break",
    "double",
    "implements",
    "protected",
    "throw",
    "byte",
    "else",
    "import",
    "public",
    "throws",
    "case",
    "enum",
    "instanceof",
    "return",
    "transient",
    "catch",
    "extends",
    "int",
    "short",
    "try",
    "char",
    "final",
    "interface",
    "static",
    "void",
    "class",
    "finally",
    "long",
    "strictfp",
    "volatile",
    "const",
    "float",
    "native",
    "super",
    "while",
    "true",
    "false",
    "yield",
    "record",
    "sealed",
    "non-sealed",
    "permits"
  ],
  operators: [
    "=",
    ">",
    "<",
    "!",
    "~",
    "?",
    ":",
    "==",
    "<=",
    ">=",
    "!=",
    "&&",
    "||",
    "++",
    "--",
    "+",
    "-",
    "*",
    "/",
    "&",
    "|",
    "^",
    "%",
    "<<",
    ">>",
    ">>>",
    "+=",
    "-=",
    "*=",
    "/=",
    "&=",
    "|=",
    "^=",
    "%=",
    "<<=",
    ">>=",
    ">>>="
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  digits: /\d+(_+\d+)*/,
  octaldigits: /[0-7]+(_+[0-7]+)*/,
  binarydigits: /[0-1]+(_+[0-1]+)*/,
  hexdigits: /[[0-9a-fA-F]+(_+[0-9a-fA-F]+)*/,
  tokenizer: {
    root: [
      ["non-sealed", "keyword.non-sealed"],
      [
        /[a-zA-Z_$][\w$]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/@\s*[a-zA-Z_\$][\w\$]*/, "annotation"],
      [/(@digits)[eE]([\-+]?(@digits))?[fFdD]?/, "number.float"],
      [/(@digits)\.(@digits)([eE][\-+]?(@digits))?[fFdD]?/, "number.float"],
      [/0[xX](@hexdigits)[Ll]?/, "number.hex"],
      [/0(@octaldigits)[Ll]?/, "number.octal"],
      [/0[bB](@binarydigits)[Ll]?/, "number.binary"],
      [/(@digits)[fFdD]/, "number.float"],
      [/(@digits)[lL]?/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"""/, "string", "@multistring"],
      [/"/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@javadoc"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    javadoc: [
      [/[^\/*]+/, "comment.doc"],
      [/\/\*/, "comment.doc.invalid"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    multistring: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"""/, "string", "@pop"],
      [/./, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/javascript/javascript.js":
/*!************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/javascript/javascript.js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../typescript/typescript.js */ "./node_modules/monaco-editor/esm/vs/basic-languages/typescript/typescript.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/javascript/javascript.ts

var conf = _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.conf;
var language = {
  defaultToken: "invalid",
  tokenPostfix: ".js",
  keywords: [
    "break",
    "case",
    "catch",
    "class",
    "continue",
    "const",
    "constructor",
    "debugger",
    "default",
    "delete",
    "do",
    "else",
    "export",
    "extends",
    "false",
    "finally",
    "for",
    "from",
    "function",
    "get",
    "if",
    "import",
    "in",
    "instanceof",
    "let",
    "new",
    "null",
    "return",
    "set",
    "static",
    "super",
    "switch",
    "symbol",
    "this",
    "throw",
    "true",
    "try",
    "typeof",
    "undefined",
    "var",
    "void",
    "while",
    "with",
    "yield",
    "async",
    "await",
    "of"
  ],
  typeKeywords: [],
  operators: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.operators,
  symbols: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.symbols,
  escapes: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.escapes,
  digits: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.digits,
  octaldigits: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.octaldigits,
  binarydigits: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.binarydigits,
  hexdigits: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.hexdigits,
  regexpctl: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.regexpctl,
  regexpesc: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.regexpesc,
  tokenizer: _typescript_typescript_js__WEBPACK_IMPORTED_MODULE_0__.language.tokenizer
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/julia/julia.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/julia/julia.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/julia/julia.ts
var conf = {
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  tokenPostfix: ".julia",
  keywords: [
    "begin",
    "while",
    "if",
    "for",
    "try",
    "return",
    "break",
    "continue",
    "function",
    "macro",
    "quote",
    "let",
    "local",
    "global",
    "const",
    "do",
    "struct",
    "module",
    "baremodule",
    "using",
    "import",
    "export",
    "end",
    "else",
    "elseif",
    "catch",
    "finally",
    "mutable",
    "primitive",
    "abstract",
    "type",
    "in",
    "isa",
    "where",
    "new"
  ],
  types: [
    "LinRange",
    "LineNumberNode",
    "LinearIndices",
    "LoadError",
    "MIME",
    "Matrix",
    "Method",
    "MethodError",
    "Missing",
    "MissingException",
    "Module",
    "NTuple",
    "NamedTuple",
    "Nothing",
    "Number",
    "OrdinalRange",
    "OutOfMemoryError",
    "OverflowError",
    "Pair",
    "PartialQuickSort",
    "PermutedDimsArray",
    "Pipe",
    "Ptr",
    "QuoteNode",
    "Rational",
    "RawFD",
    "ReadOnlyMemoryError",
    "Real",
    "ReentrantLock",
    "Ref",
    "Regex",
    "RegexMatch",
    "RoundingMode",
    "SegmentationFault",
    "Set",
    "Signed",
    "Some",
    "StackOverflowError",
    "StepRange",
    "StepRangeLen",
    "StridedArray",
    "StridedMatrix",
    "StridedVecOrMat",
    "StridedVector",
    "String",
    "StringIndexError",
    "SubArray",
    "SubString",
    "SubstitutionString",
    "Symbol",
    "SystemError",
    "Task",
    "Text",
    "TextDisplay",
    "Timer",
    "Tuple",
    "Type",
    "TypeError",
    "TypeVar",
    "UInt",
    "UInt128",
    "UInt16",
    "UInt32",
    "UInt64",
    "UInt8",
    "UndefInitializer",
    "AbstractArray",
    "UndefKeywordError",
    "AbstractChannel",
    "UndefRefError",
    "AbstractChar",
    "UndefVarError",
    "AbstractDict",
    "Union",
    "AbstractDisplay",
    "UnionAll",
    "AbstractFloat",
    "UnitRange",
    "AbstractIrrational",
    "Unsigned",
    "AbstractMatrix",
    "AbstractRange",
    "Val",
    "AbstractSet",
    "Vararg",
    "AbstractString",
    "VecElement",
    "AbstractUnitRange",
    "VecOrMat",
    "AbstractVecOrMat",
    "Vector",
    "AbstractVector",
    "VersionNumber",
    "Any",
    "WeakKeyDict",
    "ArgumentError",
    "WeakRef",
    "Array",
    "AssertionError",
    "BigFloat",
    "BigInt",
    "BitArray",
    "BitMatrix",
    "BitSet",
    "BitVector",
    "Bool",
    "BoundsError",
    "CapturedException",
    "CartesianIndex",
    "CartesianIndices",
    "Cchar",
    "Cdouble",
    "Cfloat",
    "Channel",
    "Char",
    "Cint",
    "Cintmax_t",
    "Clong",
    "Clonglong",
    "Cmd",
    "Colon",
    "Complex",
    "ComplexF16",
    "ComplexF32",
    "ComplexF64",
    "CompositeException",
    "Condition",
    "Cptrdiff_t",
    "Cshort",
    "Csize_t",
    "Cssize_t",
    "Cstring",
    "Cuchar",
    "Cuint",
    "Cuintmax_t",
    "Culong",
    "Culonglong",
    "Cushort",
    "Cvoid",
    "Cwchar_t",
    "Cwstring",
    "DataType",
    "DenseArray",
    "DenseMatrix",
    "DenseVecOrMat",
    "DenseVector",
    "Dict",
    "DimensionMismatch",
    "Dims",
    "DivideError",
    "DomainError",
    "EOFError",
    "Enum",
    "ErrorException",
    "Exception",
    "ExponentialBackOff",
    "Expr",
    "Float16",
    "Float32",
    "Float64",
    "Function",
    "GlobalRef",
    "HTML",
    "IO",
    "IOBuffer",
    "IOContext",
    "IOStream",
    "IdDict",
    "IndexCartesian",
    "IndexLinear",
    "IndexStyle",
    "InexactError",
    "InitError",
    "Int",
    "Int128",
    "Int16",
    "Int32",
    "Int64",
    "Int8",
    "Integer",
    "InterruptException",
    "InvalidStateException",
    "Irrational",
    "KeyError"
  ],
  keywordops: ["<:", ">:", ":", "=>", "...", ".", "->", "?"],
  allops: /[^\w\d\s()\[\]{}"'#]+/,
  constants: [
    "true",
    "false",
    "nothing",
    "missing",
    "undef",
    "Inf",
    "pi",
    "NaN",
    "\u03C0",
    "\u212F",
    "ans",
    "PROGRAM_FILE",
    "ARGS",
    "C_NULL",
    "VERSION",
    "DEPOT_PATH",
    "LOAD_PATH"
  ],
  operators: [
    "!",
    "!=",
    "!==",
    "%",
    "&",
    "*",
    "+",
    "-",
    "/",
    "//",
    "<",
    "<<",
    "<=",
    "==",
    "===",
    "=>",
    ">",
    ">=",
    ">>",
    ">>>",
    "\\",
    "^",
    "|",
    "|>",
    "~",
    "\xF7",
    "\u2208",
    "\u2209",
    "\u220B",
    "\u220C",
    "\u2218",
    "\u221A",
    "\u221B",
    "\u2229",
    "\u222A",
    "\u2248",
    "\u2249",
    "\u2260",
    "\u2261",
    "\u2262",
    "\u2264",
    "\u2265",
    "\u2286",
    "\u2287",
    "\u2288",
    "\u2289",
    "\u228A",
    "\u228B",
    "\u22BB"
  ],
  brackets: [
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" }
  ],
  ident: /π|ℯ|\b(?!\d)\w+\b/,
  escape: /(?:[abefnrstv\\"'\n\r]|[0-7]{1,3}|x[0-9A-Fa-f]{1,2}|u[0-9A-Fa-f]{4})/,
  escapes: /\\(?:C\-(@escape|.)|c(@escape|.)|@escape)/,
  tokenizer: {
    root: [
      [/(::)\s*|\b(isa)\s+/, "keyword", "@typeanno"],
      [/\b(isa)(\s*\(@ident\s*,\s*)/, ["keyword", { token: "", next: "@typeanno" }]],
      [/\b(type|struct)[ \t]+/, "keyword", "@typeanno"],
      [/^\s*:@ident[!?]?/, "metatag"],
      [/(return)(\s*:@ident[!?]?)/, ["keyword", "metatag"]],
      [/(\(|\[|\{|@allops)(\s*:@ident[!?]?)/, ["", "metatag"]],
      [/:\(/, "metatag", "@quote"],
      [/r"""/, "regexp.delim", "@tregexp"],
      [/r"/, "regexp.delim", "@sregexp"],
      [/raw"""/, "string.delim", "@rtstring"],
      [/[bv]?"""/, "string.delim", "@dtstring"],
      [/raw"/, "string.delim", "@rsstring"],
      [/[bv]?"/, "string.delim", "@dsstring"],
      [
        /(@ident)\{/,
        {
          cases: {
            "$1@types": { token: "type", next: "@gen" },
            "@default": { token: "type", next: "@gen" }
          }
        }
      ],
      [
        /@ident[!?'']?(?=\.?\()/,
        {
          cases: {
            "@types": "type",
            "@keywords": "keyword",
            "@constants": "variable",
            "@default": "keyword.flow"
          }
        }
      ],
      [
        /@ident[!?']?/,
        {
          cases: {
            "@types": "type",
            "@keywords": "keyword",
            "@constants": "variable",
            "@default": "identifier"
          }
        }
      ],
      [/\$\w+/, "key"],
      [/\$\(/, "key", "@paste"],
      [/@@@ident/, "annotation"],
      { include: "@whitespace" },
      [/'(?:@escapes|.)'/, "string.character"],
      [/[()\[\]{}]/, "@brackets"],
      [
        /@allops/,
        {
          cases: {
            "@keywordops": "keyword",
            "@operators": "operator"
          }
        }
      ],
      [/[;,]/, "delimiter"],
      [/0[xX][0-9a-fA-F](_?[0-9a-fA-F])*/, "number.hex"],
      [/0[_oO][0-7](_?[0-7])*/, "number.octal"],
      [/0[bB][01](_?[01])*/, "number.binary"],
      [/[+\-]?\d+(\.\d+)?(im?|[eE][+\-]?\d+(\.\d+)?)?/, "number"]
    ],
    typeanno: [
      [/[a-zA-Z_]\w*(?:\.[a-zA-Z_]\w*)*\{/, "type", "@gen"],
      [/([a-zA-Z_]\w*(?:\.[a-zA-Z_]\w*)*)(\s*<:\s*)/, ["type", "keyword"]],
      [/[a-zA-Z_]\w*(?:\.[a-zA-Z_]\w*)*/, "type", "@pop"],
      ["", "", "@pop"]
    ],
    gen: [
      [/[a-zA-Z_]\w*(?:\.[a-zA-Z_]\w*)*\{/, "type", "@push"],
      [/[a-zA-Z_]\w*(?:\.[a-zA-Z_]\w*)*/, "type"],
      [/<:/, "keyword"],
      [/(\})(\s*<:\s*)/, ["type", { token: "keyword", next: "@pop" }]],
      [/\}/, "type", "@pop"],
      { include: "@root" }
    ],
    quote: [
      [/\$\(/, "key", "@paste"],
      [/\(/, "@brackets", "@paren"],
      [/\)/, "metatag", "@pop"],
      { include: "@root" }
    ],
    paste: [
      [/:\(/, "metatag", "@quote"],
      [/\(/, "@brackets", "@paren"],
      [/\)/, "key", "@pop"],
      { include: "@root" }
    ],
    paren: [
      [/\$\(/, "key", "@paste"],
      [/:\(/, "metatag", "@quote"],
      [/\(/, "@brackets", "@push"],
      [/\)/, "@brackets", "@pop"],
      { include: "@root" }
    ],
    sregexp: [
      [/^.*/, "invalid"],
      [/[^\\"()\[\]{}]/, "regexp"],
      [/[()\[\]{}]/, "@brackets"],
      [/\\./, "operator.scss"],
      [/"[imsx]*/, "regexp.delim", "@pop"]
    ],
    tregexp: [
      [/[^\\"()\[\]{}]/, "regexp"],
      [/[()\[\]{}]/, "@brackets"],
      [/\\./, "operator.scss"],
      [/"(?!"")/, "string"],
      [/"""[imsx]*/, "regexp.delim", "@pop"]
    ],
    rsstring: [
      [/^.*/, "invalid"],
      [/[^\\"]/, "string"],
      [/\\./, "string.escape"],
      [/"/, "string.delim", "@pop"]
    ],
    rtstring: [
      [/[^\\"]/, "string"],
      [/\\./, "string.escape"],
      [/"(?!"")/, "string"],
      [/"""/, "string.delim", "@pop"]
    ],
    dsstring: [
      [/^.*/, "invalid"],
      [/[^\\"\$]/, "string"],
      [/\$/, "", "@interpolated"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string.delim", "@pop"]
    ],
    dtstring: [
      [/[^\\"\$]/, "string"],
      [/\$/, "", "@interpolated"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"(?!"")/, "string"],
      [/"""/, "string.delim", "@pop"]
    ],
    interpolated: [
      [/\(/, { token: "", switchTo: "@interpolated_compound" }],
      [/[a-zA-Z_]\w*/, "identifier"],
      ["", "", "@pop"]
    ],
    interpolated_compound: [[/\)/, "", "@pop"], { include: "@root" }],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/#=/, "comment", "@multi_comment"],
      [/#.*$/, "comment"]
    ],
    multi_comment: [
      [/#=/, "comment", "@push"],
      [/=#/, "comment", "@pop"],
      [/=(?!#)|#(?!=)/, "comment"],
      [/[^#=]+/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/kotlin/kotlin.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/kotlin/kotlin.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/kotlin/kotlin.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\#\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "<", close: ">" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*//\\s*(?:(?:#?region\\b)|(?:<editor-fold\\b))"),
      end: new RegExp("^\\s*//\\s*(?:(?:#?endregion\\b)|(?:</editor-fold>))")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".kt",
  keywords: [
    "as",
    "as?",
    "break",
    "class",
    "continue",
    "do",
    "else",
    "false",
    "for",
    "fun",
    "if",
    "in",
    "!in",
    "interface",
    "is",
    "!is",
    "null",
    "object",
    "package",
    "return",
    "super",
    "this",
    "throw",
    "true",
    "try",
    "typealias",
    "val",
    "var",
    "when",
    "while",
    "by",
    "catch",
    "constructor",
    "delegate",
    "dynamic",
    "field",
    "file",
    "finally",
    "get",
    "import",
    "init",
    "param",
    "property",
    "receiver",
    "set",
    "setparam",
    "where",
    "actual",
    "abstract",
    "annotation",
    "companion",
    "const",
    "crossinline",
    "data",
    "enum",
    "expect",
    "external",
    "final",
    "infix",
    "inline",
    "inner",
    "internal",
    "lateinit",
    "noinline",
    "open",
    "operator",
    "out",
    "override",
    "private",
    "protected",
    "public",
    "reified",
    "sealed",
    "suspend",
    "tailrec",
    "vararg",
    "field",
    "it"
  ],
  operators: [
    "+",
    "-",
    "*",
    "/",
    "%",
    "=",
    "+=",
    "-=",
    "*=",
    "/=",
    "%=",
    "++",
    "--",
    "&&",
    "||",
    "!",
    "==",
    "!=",
    "===",
    "!==",
    ">",
    "<",
    "<=",
    ">=",
    "[",
    "]",
    "!!",
    "?.",
    "?:",
    "::",
    "..",
    ":",
    "?",
    "->",
    "@",
    ";",
    "$",
    "_"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  digits: /\d+(_+\d+)*/,
  octaldigits: /[0-7]+(_+[0-7]+)*/,
  binarydigits: /[0-1]+(_+[0-1]+)*/,
  hexdigits: /[[0-9a-fA-F]+(_+[0-9a-fA-F]+)*/,
  tokenizer: {
    root: [
      [/[A-Z][\w\$]*/, "type.identifier"],
      [
        /[a-zA-Z_$][\w$]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/@\s*[a-zA-Z_\$][\w\$]*/, "annotation"],
      [/(@digits)[eE]([\-+]?(@digits))?[fFdD]?/, "number.float"],
      [/(@digits)\.(@digits)([eE][\-+]?(@digits))?[fFdD]?/, "number.float"],
      [/0[xX](@hexdigits)[Ll]?/, "number.hex"],
      [/0(@octaldigits)[Ll]?/, "number.octal"],
      [/0[bB](@binarydigits)[Ll]?/, "number.binary"],
      [/(@digits)[fFdD]/, "number.float"],
      [/(@digits)[lL]?/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"""/, "string", "@multistring"],
      [/"/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@javadoc"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\/\*/, "comment", "@comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    javadoc: [
      [/[^\/*]+/, "comment.doc"],
      [/\/\*/, "comment.doc", "@push"],
      [/\/\*/, "comment.doc.invalid"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    multistring: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"""/, "string", "@pop"],
      [/./, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/less/less.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/less/less.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/less/less.ts
var conf = {
  wordPattern: /(#?-?\d*\.\d\w*%?)|([@#!.:]?[\w-?]+%?)|[@#!.]/g,
  comments: {
    blockComment: ["/*", "*/"],
    lineComment: "//"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*\\/\\*\\s*#region\\b\\s*(.*?)\\s*\\*\\/"),
      end: new RegExp("^\\s*\\/\\*\\s*#endregion\\b.*\\*\\/")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".less",
  identifier: "-?-?([a-zA-Z]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))([\\w\\-]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))*",
  identifierPlus: "-?-?([a-zA-Z:.]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))([\\w\\-:.]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))*",
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.bracket" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  tokenizer: {
    root: [
      { include: "@nestedJSBegin" },
      ["[ \\t\\r\\n]+", ""],
      { include: "@comments" },
      { include: "@keyword" },
      { include: "@strings" },
      { include: "@numbers" },
      ["[*_]?[a-zA-Z\\-\\s]+(?=:.*(;|(\\\\$)))", "attribute.name", "@attribute"],
      ["url(\\-prefix)?\\(", { token: "tag", next: "@urldeclaration" }],
      ["[{}()\\[\\]]", "@brackets"],
      ["[,:;]", "delimiter"],
      ["#@identifierPlus", "tag.id"],
      ["&", "tag"],
      ["\\.@identifierPlus(?=\\()", "tag.class", "@attribute"],
      ["\\.@identifierPlus", "tag.class"],
      ["@identifierPlus", "tag"],
      { include: "@operators" },
      ["@(@identifier(?=[:,\\)]))", "variable", "@attribute"],
      ["@(@identifier)", "variable"],
      ["@", "key", "@atRules"]
    ],
    nestedJSBegin: [
      ["``", "delimiter.backtick"],
      [
        "`",
        {
          token: "delimiter.backtick",
          next: "@nestedJSEnd",
          nextEmbedded: "text/javascript"
        }
      ]
    ],
    nestedJSEnd: [
      [
        "`",
        {
          token: "delimiter.backtick",
          next: "@pop",
          nextEmbedded: "@pop"
        }
      ]
    ],
    operators: [["[<>=\\+\\-\\*\\/\\^\\|\\~]", "operator"]],
    keyword: [
      [
        "(@[\\s]*import|![\\s]*important|true|false|when|iscolor|isnumber|isstring|iskeyword|isurl|ispixel|ispercentage|isem|hue|saturation|lightness|alpha|lighten|darken|saturate|desaturate|fadein|fadeout|fade|spin|mix|round|ceil|floor|percentage)\\b",
        "keyword"
      ]
    ],
    urldeclaration: [
      { include: "@strings" },
      ["[^)\r\n]+", "string"],
      ["\\)", { token: "tag", next: "@pop" }]
    ],
    attribute: [
      { include: "@nestedJSBegin" },
      { include: "@comments" },
      { include: "@strings" },
      { include: "@numbers" },
      { include: "@keyword" },
      ["[a-zA-Z\\-]+(?=\\()", "attribute.value", "@attribute"],
      [">", "operator", "@pop"],
      ["@identifier", "attribute.value"],
      { include: "@operators" },
      ["@(@identifier)", "variable"],
      ["[)\\}]", "@brackets", "@pop"],
      ["[{}()\\[\\]>]", "@brackets"],
      ["[;]", "delimiter", "@pop"],
      ["[,=:]", "delimiter"],
      ["\\s", ""],
      [".", "attribute.value"]
    ],
    comments: [
      ["\\/\\*", "comment", "@comment"],
      ["\\/\\/+.*", "comment"]
    ],
    comment: [
      ["\\*\\/", "comment", "@pop"],
      [".", "comment"]
    ],
    numbers: [
      ["(\\d*\\.)?\\d+([eE][\\-+]?\\d+)?", { token: "attribute.value.number", next: "@units" }],
      ["#[0-9a-fA-F_]+(?!\\w)", "attribute.value.hex"]
    ],
    units: [
      [
        "(em|ex|ch|rem|fr|vmin|vmax|vw|vh|vm|cm|mm|in|px|pt|pc|deg|grad|rad|turn|s|ms|Hz|kHz|%)?",
        "attribute.value.unit",
        "@pop"
      ]
    ],
    strings: [
      ['~?"', { token: "string.delimiter", next: "@stringsEndDoubleQuote" }],
      ["~?'", { token: "string.delimiter", next: "@stringsEndQuote" }]
    ],
    stringsEndDoubleQuote: [
      ['\\\\"', "string"],
      ['"', { token: "string.delimiter", next: "@popall" }],
      [".", "string"]
    ],
    stringsEndQuote: [
      ["\\\\'", "string"],
      ["'", { token: "string.delimiter", next: "@popall" }],
      [".", "string"]
    ],
    atRules: [
      { include: "@comments" },
      { include: "@strings" },
      ["[()]", "delimiter"],
      ["[\\{;]", "delimiter", "@pop"],
      [".", "key"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/lexon/lexon.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/lexon/lexon.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/lexon/lexon.ts
var conf = {
  comments: {
    lineComment: "COMMENT"
  },
  brackets: [["(", ")"]],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: ":", close: "." }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "`", close: "`" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: ":", close: "." }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*(::\\s*|COMMENT\\s+)#region"),
      end: new RegExp("^\\s*(::\\s*|COMMENT\\s+)#endregion")
    }
  }
};
var language = {
  tokenPostfix: ".lexon",
  ignoreCase: true,
  keywords: [
    "lexon",
    "lex",
    "clause",
    "terms",
    "contracts",
    "may",
    "pay",
    "pays",
    "appoints",
    "into",
    "to"
  ],
  typeKeywords: ["amount", "person", "key", "time", "date", "asset", "text"],
  operators: [
    "less",
    "greater",
    "equal",
    "le",
    "gt",
    "or",
    "and",
    "add",
    "added",
    "subtract",
    "subtracted",
    "multiply",
    "multiplied",
    "times",
    "divide",
    "divided",
    "is",
    "be",
    "certified"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  tokenizer: {
    root: [
      [/^(\s*)(comment:?(?:\s.*|))$/, ["", "comment"]],
      [
        /"/,
        {
          token: "identifier.quote",
          bracket: "@open",
          next: "@quoted_identifier"
        }
      ],
      [
        "LEX$",
        {
          token: "keyword",
          bracket: "@open",
          next: "@identifier_until_period"
        }
      ],
      ["LEXON", { token: "keyword", bracket: "@open", next: "@semver" }],
      [
        ":",
        {
          token: "delimiter",
          bracket: "@open",
          next: "@identifier_until_period"
        }
      ],
      [
        /[a-z_$][\w$]*/,
        {
          cases: {
            "@operators": "operator",
            "@typeKeywords": "keyword.type",
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/\d*\.\d*\.\d*/, "number.semver"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F]+/, "number.hex"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"]
    ],
    quoted_identifier: [
      [/[^\\"]+/, "identifier"],
      [/"/, { token: "identifier.quote", bracket: "@close", next: "@pop" }]
    ],
    space_identifier_until_period: [
      [":", "delimiter"],
      [" ", { token: "white", next: "@identifier_rest" }]
    ],
    identifier_until_period: [
      { include: "@whitespace" },
      [":", { token: "delimiter", next: "@identifier_rest" }],
      [/[^\\.]+/, "identifier"],
      [/\./, { token: "delimiter", bracket: "@close", next: "@pop" }]
    ],
    identifier_rest: [
      [/[^\\.]+/, "identifier"],
      [/\./, { token: "delimiter", bracket: "@close", next: "@pop" }]
    ],
    semver: [
      { include: "@whitespace" },
      [":", "delimiter"],
      [/\d*\.\d*\.\d*/, { token: "number.semver", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [[/[ \t\r\n]+/, "white"]]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/liquid/liquid.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/liquid/liquid.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/liquid/liquid.ts
var EMPTY_ELEMENTS = [
  "area",
  "base",
  "br",
  "col",
  "embed",
  "hr",
  "img",
  "input",
  "keygen",
  "link",
  "menuitem",
  "meta",
  "param",
  "source",
  "track",
  "wbr"
];
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\$\^\&\*\(\)\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\s]+)/g,
  brackets: [
    ["<!--", "-->"],
    ["<", ">"],
    ["{{", "}}"],
    ["{%", "%}"],
    ["{", "}"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "%", close: "%" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "<", close: ">" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  onEnterRules: [
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      afterText: /^<\/(\w[\w\d]*)\s*>$/i,
      action: {
        indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
      }
    },
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: "",
  builtinTags: [
    "if",
    "else",
    "elseif",
    "endif",
    "render",
    "assign",
    "capture",
    "endcapture",
    "case",
    "endcase",
    "comment",
    "endcomment",
    "cycle",
    "decrement",
    "for",
    "endfor",
    "include",
    "increment",
    "layout",
    "raw",
    "endraw",
    "render",
    "tablerow",
    "endtablerow",
    "unless",
    "endunless"
  ],
  builtinFilters: [
    "abs",
    "append",
    "at_least",
    "at_most",
    "capitalize",
    "ceil",
    "compact",
    "date",
    "default",
    "divided_by",
    "downcase",
    "escape",
    "escape_once",
    "first",
    "floor",
    "join",
    "json",
    "last",
    "lstrip",
    "map",
    "minus",
    "modulo",
    "newline_to_br",
    "plus",
    "prepend",
    "remove",
    "remove_first",
    "replace",
    "replace_first",
    "reverse",
    "round",
    "rstrip",
    "size",
    "slice",
    "sort",
    "sort_natural",
    "split",
    "strip",
    "strip_html",
    "strip_newlines",
    "times",
    "truncate",
    "truncatewords",
    "uniq",
    "upcase",
    "url_decode",
    "url_encode",
    "where"
  ],
  constants: ["true", "false"],
  operators: ["==", "!=", ">", "<", ">=", "<="],
  symbol: /[=><!]+/,
  identifier: /[a-zA-Z_][\w]*/,
  tokenizer: {
    root: [
      [/\{\%\s*comment\s*\%\}/, "comment.start.liquid", "@comment"],
      [/\{\{/, { token: "@rematch", switchTo: "@liquidState.root" }],
      [/\{\%/, { token: "@rematch", switchTo: "@liquidState.root" }],
      [/(<)([\w\-]+)(\/>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/(<)([:\w]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/(<\/)([\w\-]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/</, "delimiter.html"],
      [/\{/, "delimiter.html"],
      [/[^<{]+/]
    ],
    comment: [
      [/\{\%\s*endcomment\s*\%\}/, "comment.end.liquid", "@pop"],
      [/./, "comment.content.liquid"]
    ],
    otherTag: [
      [
        /\{\{/,
        {
          token: "@rematch",
          switchTo: "@liquidState.otherTag"
        }
      ],
      [
        /\{\%/,
        {
          token: "@rematch",
          switchTo: "@liquidState.otherTag"
        }
      ],
      [/\/?>/, "delimiter.html", "@pop"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/]
    ],
    liquidState: [
      [/\{\{/, "delimiter.output.liquid"],
      [/\}\}/, { token: "delimiter.output.liquid", switchTo: "@$S2.$S3" }],
      [/\{\%/, "delimiter.tag.liquid"],
      [/raw\s*\%\}/, "delimiter.tag.liquid", "@liquidRaw"],
      [/\%\}/, { token: "delimiter.tag.liquid", switchTo: "@$S2.$S3" }],
      { include: "liquidRoot" }
    ],
    liquidRaw: [
      [/^(?!\{\%\s*endraw\s*\%\}).+/],
      [/\{\%/, "delimiter.tag.liquid"],
      [/@identifier/],
      [/\%\}/, { token: "delimiter.tag.liquid", next: "@root" }]
    ],
    liquidRoot: [
      [/\d+(\.\d+)?/, "number.liquid"],
      [/"[^"]*"/, "string.liquid"],
      [/'[^']*'/, "string.liquid"],
      [/\s+/],
      [
        /@symbol/,
        {
          cases: {
            "@operators": "operator.liquid",
            "@default": ""
          }
        }
      ],
      [/\./],
      [
        /@identifier/,
        {
          cases: {
            "@constants": "keyword.liquid",
            "@builtinFilters": "predefined.liquid",
            "@builtinTags": "predefined.liquid",
            "@default": "variable.liquid"
          }
        }
      ],
      [/[^}|%]/, "variable.liquid"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/lua/lua.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/lua/lua.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/lua/lua.ts
var conf = {
  comments: {
    lineComment: "--",
    blockComment: ["--[[", "]]"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".lua",
  keywords: [
    "and",
    "break",
    "do",
    "else",
    "elseif",
    "end",
    "false",
    "for",
    "function",
    "goto",
    "if",
    "in",
    "local",
    "nil",
    "not",
    "or",
    "repeat",
    "return",
    "then",
    "true",
    "until",
    "while"
  ],
  brackets: [
    { token: "delimiter.bracket", open: "{", close: "}" },
    { token: "delimiter.array", open: "[", close: "]" },
    { token: "delimiter.parenthesis", open: "(", close: ")" }
  ],
  operators: [
    "+",
    "-",
    "*",
    "/",
    "%",
    "^",
    "#",
    "==",
    "~=",
    "<=",
    ">=",
    "<",
    ">",
    "=",
    ";",
    ":",
    ",",
    ".",
    "..",
    "..."
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/(,)(\s*)([a-zA-Z_]\w*)(\s*)(:)(?!:)/, ["delimiter", "", "key", "", "delimiter"]],
      [/({)(\s*)([a-zA-Z_]\w*)(\s*)(:)(?!:)/, ["@brackets", "", "key", "", "delimiter"]],
      [/[{}()\[\]]/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F_]*[0-9a-fA-F]/, "number.hex"],
      [/\d+?/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", '@string."'],
      [/'/, "string", "@string.'"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/--\[([=]*)\[/, "comment", "@comment.$1"],
      [/--.*$/, "comment"]
    ],
    comment: [
      [/[^\]]+/, "comment"],
      [
        /\]([=]*)\]/,
        {
          cases: {
            "$1==$S2": { token: "comment", next: "@pop" },
            "@default": "comment"
          }
        }
      ],
      [/./, "comment"]
    ],
    string: [
      [/[^\\"']+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/m3/m3.js":
/*!********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/m3/m3.js ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/m3/m3.ts
var conf = {
  comments: {
    blockComment: ["(*", "*)"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "(", close: ")" },
    { open: "(*", close: "*)" },
    { open: "<*", close: "*>" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".m3",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" }
  ],
  keywords: [
    "AND",
    "ANY",
    "ARRAY",
    "AS",
    "BEGIN",
    "BITS",
    "BRANDED",
    "BY",
    "CASE",
    "CONST",
    "DIV",
    "DO",
    "ELSE",
    "ELSIF",
    "END",
    "EVAL",
    "EXCEPT",
    "EXCEPTION",
    "EXIT",
    "EXPORTS",
    "FINALLY",
    "FOR",
    "FROM",
    "GENERIC",
    "IF",
    "IMPORT",
    "IN",
    "INTERFACE",
    "LOCK",
    "LOOP",
    "METHODS",
    "MOD",
    "MODULE",
    "NOT",
    "OBJECT",
    "OF",
    "OR",
    "OVERRIDES",
    "PROCEDURE",
    "RAISE",
    "RAISES",
    "READONLY",
    "RECORD",
    "REF",
    "REPEAT",
    "RETURN",
    "REVEAL",
    "SET",
    "THEN",
    "TO",
    "TRY",
    "TYPE",
    "TYPECASE",
    "UNSAFE",
    "UNTIL",
    "UNTRACED",
    "VALUE",
    "VAR",
    "WHILE",
    "WITH"
  ],
  reservedConstNames: [
    "ABS",
    "ADR",
    "ADRSIZE",
    "BITSIZE",
    "BYTESIZE",
    "CEILING",
    "DEC",
    "DISPOSE",
    "FALSE",
    "FIRST",
    "FLOAT",
    "FLOOR",
    "INC",
    "ISTYPE",
    "LAST",
    "LOOPHOLE",
    "MAX",
    "MIN",
    "NARROW",
    "NEW",
    "NIL",
    "NUMBER",
    "ORD",
    "ROUND",
    "SUBARRAY",
    "TRUE",
    "TRUNC",
    "TYPECODE",
    "VAL"
  ],
  reservedTypeNames: [
    "ADDRESS",
    "ANY",
    "BOOLEAN",
    "CARDINAL",
    "CHAR",
    "EXTENDED",
    "INTEGER",
    "LONGCARD",
    "LONGINT",
    "LONGREAL",
    "MUTEX",
    "NULL",
    "REAL",
    "REFANY",
    "ROOT",
    "TEXT"
  ],
  operators: ["+", "-", "*", "/", "&", "^", "."],
  relations: ["=", "#", "<", "<=", ">", ">=", "<:", ":"],
  delimiters: ["|", "..", "=>", ",", ";", ":="],
  symbols: /[>=<#.,:;+\-*/&^]+/,
  escapes: /\\(?:[\\fnrt"']|[0-7]{3})/,
  tokenizer: {
    root: [
      [/_\w*/, "invalid"],
      [
        /[a-zA-Z][a-zA-Z0-9_]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@reservedConstNames": { token: "constant.reserved.$0" },
            "@reservedTypeNames": { token: "type.reserved.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[0-9]+\.[0-9]+(?:[DdEeXx][\+\-]?[0-9]+)?/, "number.float"],
      [/[0-9]+(?:\_[0-9a-fA-F]+)?L?/, "number"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "operators",
            "@relations": "operators",
            "@delimiters": "delimiter",
            "@default": "invalid"
          }
        }
      ],
      [/'[^\\']'/, "string.char"],
      [/(')(@escapes)(')/, ["string.char", "string.escape", "string.char"]],
      [/'/, "invalid"],
      [/"([^"\\]|\\.)*$/, "invalid"],
      [/"/, "string.text", "@text"]
    ],
    text: [
      [/[^\\"]+/, "string.text"],
      [/@escapes/, "string.escape"],
      [/\\./, "invalid"],
      [/"/, "string.text", "@pop"]
    ],
    comment: [
      [/\(\*/, "comment", "@push"],
      [/\*\)/, "comment", "@pop"],
      [/./, "comment"]
    ],
    pragma: [
      [/<\*/, "keyword.pragma", "@push"],
      [/\*>/, "keyword.pragma", "@pop"],
      [/./, "keyword.pragma"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\(\*/, "comment", "@comment"],
      [/<\*/, "keyword.pragma", "@pragma"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/markdown/markdown.js":
/*!********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/markdown/markdown.js ***!
  \********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/markdown/markdown.ts
var conf = {
  comments: {
    blockComment: ["<!--", "-->"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">", notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "(", close: ")" },
    { open: "[", close: "]" },
    { open: "`", close: "`" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*<!--\\s*#?region\\b.*-->"),
      end: new RegExp("^\\s*<!--\\s*#?endregion\\b.*-->")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".md",
  control: /[\\`*_\[\]{}()#+\-\.!]/,
  noncontrol: /[^\\`*_\[\]{}()#+\-\.!]/,
  escapes: /\\(?:@control)/,
  jsescapes: /\\(?:[btnfr\\"']|[0-7][0-7]?|[0-3][0-7]{2})/,
  empty: [
    "area",
    "base",
    "basefont",
    "br",
    "col",
    "frame",
    "hr",
    "img",
    "input",
    "isindex",
    "link",
    "meta",
    "param"
  ],
  tokenizer: {
    root: [
      [/^\s*\|/, "@rematch", "@table_header"],
      [/^(\s{0,3})(#+)((?:[^\\#]|@escapes)+)((?:#+)?)/, ["white", "keyword", "keyword", "keyword"]],
      [/^\s*(=+|\-+)\s*$/, "keyword"],
      [/^\s*((\*[ ]?)+)\s*$/, "meta.separator"],
      [/^\s*>+/, "comment"],
      [/^\s*([\*\-+:]|\d+\.)\s/, "keyword"],
      [/^(\t|[ ]{4})[^ ].*$/, "string"],
      [/^\s*~~~\s*((?:\w|[\/\-#])+)?\s*$/, { token: "string", next: "@codeblock" }],
      [
        /^\s*```\s*((?:\w|[\/\-#])+).*$/,
        { token: "string", next: "@codeblockgh", nextEmbedded: "$1" }
      ],
      [/^\s*```\s*$/, { token: "string", next: "@codeblock" }],
      { include: "@linecontent" }
    ],
    table_header: [
      { include: "@table_common" },
      [/[^\|]+/, "keyword.table.header"]
    ],
    table_body: [{ include: "@table_common" }, { include: "@linecontent" }],
    table_common: [
      [/\s*[\-:]+\s*/, { token: "keyword", switchTo: "table_body" }],
      [/^\s*\|/, "keyword.table.left"],
      [/^\s*[^\|]/, "@rematch", "@pop"],
      [/^\s*$/, "@rematch", "@pop"],
      [
        /\|/,
        {
          cases: {
            "@eos": "keyword.table.right",
            "@default": "keyword.table.middle"
          }
        }
      ]
    ],
    codeblock: [
      [/^\s*~~~\s*$/, { token: "string", next: "@pop" }],
      [/^\s*```\s*$/, { token: "string", next: "@pop" }],
      [/.*$/, "variable.source"]
    ],
    codeblockgh: [
      [/```\s*$/, { token: "string", next: "@pop", nextEmbedded: "@pop" }],
      [/[^`]+/, "variable.source"]
    ],
    linecontent: [
      [/&\w+;/, "string.escape"],
      [/@escapes/, "escape"],
      [/\b__([^\\_]|@escapes|_(?!_))+__\b/, "strong"],
      [/\*\*([^\\*]|@escapes|\*(?!\*))+\*\*/, "strong"],
      [/\b_[^_]+_\b/, "emphasis"],
      [/\*([^\\*]|@escapes)+\*/, "emphasis"],
      [/`([^\\`]|@escapes)+`/, "variable"],
      [/\{+[^}]+\}+/, "string.target"],
      [/(!?\[)((?:[^\]\\]|@escapes)*)(\]\([^\)]+\))/, ["string.link", "", "string.link"]],
      [/(!?\[)((?:[^\]\\]|@escapes)*)(\])/, "string.link"],
      { include: "html" }
    ],
    html: [
      [/<(\w+)\/>/, "tag"],
      [
        /<(\w+)(\-|\w)*/,
        {
          cases: {
            "@empty": { token: "tag", next: "@tag.$1" },
            "@default": { token: "tag", next: "@tag.$1" }
          }
        }
      ],
      [/<\/(\w+)(\-|\w)*\s*>/, { token: "tag" }],
      [/<!--/, "comment", "@comment"]
    ],
    comment: [
      [/[^<\-]+/, "comment.content"],
      [/-->/, "comment", "@pop"],
      [/<!--/, "comment.content.invalid"],
      [/[<\-]/, "comment.content"]
    ],
    tag: [
      [/[ \t\r\n]+/, "white"],
      [
        /(type)(\s*=\s*)(")([^"]+)(")/,
        [
          "attribute.name.html",
          "delimiter.html",
          "string.html",
          { token: "string.html", switchTo: "@tag.$S2.$4" },
          "string.html"
        ]
      ],
      [
        /(type)(\s*=\s*)(')([^']+)(')/,
        [
          "attribute.name.html",
          "delimiter.html",
          "string.html",
          { token: "string.html", switchTo: "@tag.$S2.$4" },
          "string.html"
        ]
      ],
      [/(\w+)(\s*=\s*)("[^"]*"|'[^']*')/, ["attribute.name.html", "delimiter.html", "string.html"]],
      [/\w+/, "attribute.name.html"],
      [/\/>/, "tag", "@pop"],
      [
        />/,
        {
          cases: {
            "$S2==style": {
              token: "tag",
              switchTo: "embeddedStyle",
              nextEmbedded: "text/css"
            },
            "$S2==script": {
              cases: {
                $S3: {
                  token: "tag",
                  switchTo: "embeddedScript",
                  nextEmbedded: "$S3"
                },
                "@default": {
                  token: "tag",
                  switchTo: "embeddedScript",
                  nextEmbedded: "text/javascript"
                }
              }
            },
            "@default": { token: "tag", next: "@pop" }
          }
        }
      ]
    ],
    embeddedStyle: [
      [/[^<]+/, ""],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }],
      [/</, ""]
    ],
    embeddedScript: [
      [/[^<]+/, ""],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }],
      [/</, ""]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/mdx/mdx.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/mdx/mdx.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/mdx/mdx.ts
var conf = {
  comments: {
    blockComment: ["{/*", "*/}"]
  },
  brackets: [["{", "}"]],
  autoClosingPairs: [
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "\u201C", close: "\u201D" },
    { open: "\u2018", close: "\u2019" },
    { open: "`", close: "`" },
    { open: "{", close: "}" },
    { open: "(", close: ")" },
    { open: "_", close: "_" },
    { open: "**", close: "**" },
    { open: "<", close: ">" }
  ],
  onEnterRules: [
    {
      beforeText: /^\s*- .+/,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.None, appendText: "- " }
    },
    {
      beforeText: /^\s*\+ .+/,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.None, appendText: "+ " }
    },
    {
      beforeText: /^\s*\* .+/,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.None, appendText: "* " }
    },
    {
      beforeText: /^> /,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.None, appendText: "> " }
    },
    {
      beforeText: /<\w+/,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    },
    {
      beforeText: /\s+>\s*$/,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    },
    {
      beforeText: /<\/\w+>/,
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Outdent }
    },
    ...Array.from({ length: 100 }, (_, index) => ({
      beforeText: new RegExp(`^${index}\\. .+`),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.None, appendText: `${index + 1}. ` }
    }))
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".mdx",
  control: /[!#()*+.[\\\]_`{}\-]/,
  escapes: /\\@control/,
  tokenizer: {
    root: [
      [/^---$/, { token: "meta.content", next: "@frontmatter", nextEmbedded: "yaml" }],
      [/^\s*import/, { token: "keyword", next: "@import", nextEmbedded: "js" }],
      [/^\s*export/, { token: "keyword", next: "@export", nextEmbedded: "js" }],
      [/<\w+/, { token: "type.identifier", next: "@jsx" }],
      [/<\/?\w+>/, "type.identifier"],
      [
        /^(\s*)(>*\s*)(#{1,6}\s)/,
        [{ token: "white" }, { token: "comment" }, { token: "keyword", next: "@header" }]
      ],
      [/^(\s*)(>*\s*)([*+-])(\s+)/, ["white", "comment", "keyword", "white"]],
      [/^(\s*)(>*\s*)(\d{1,9}\.)(\s+)/, ["white", "comment", "number", "white"]],
      [/^(\s*)(>*\s*)(\d{1,9}\.)(\s+)/, ["white", "comment", "number", "white"]],
      [/^(\s*)(>*\s*)(-{3,}|\*{3,}|_{3,})$/, ["white", "comment", "keyword"]],
      [/`{3,}(\s.*)?$/, { token: "string", next: "@codeblock_backtick" }],
      [/~{3,}(\s.*)?$/, { token: "string", next: "@codeblock_tilde" }],
      [
        /`{3,}(\S+).*$/,
        { token: "string", next: "@codeblock_highlight_backtick", nextEmbedded: "$1" }
      ],
      [
        /~{3,}(\S+).*$/,
        { token: "string", next: "@codeblock_highlight_tilde", nextEmbedded: "$1" }
      ],
      [/^(\s*)(-{4,})$/, ["white", "comment"]],
      [/^(\s*)(>+)/, ["white", "comment"]],
      { include: "content" }
    ],
    content: [
      [
        /(\[)(.+)(]\()(.+)(\s+".*")(\))/,
        ["", "string.link", "", "type.identifier", "string.link", ""]
      ],
      [/(\[)(.+)(]\()(.+)(\))/, ["", "type.identifier", "", "string.link", ""]],
      [/(\[)(.+)(]\[)(.+)(])/, ["", "type.identifier", "", "type.identifier", ""]],
      [/(\[)(.+)(]:\s+)(\S*)/, ["", "type.identifier", "", "string.link"]],
      [/(\[)(.+)(])/, ["", "type.identifier", ""]],
      [/`.*`/, "variable.source"],
      [/_/, { token: "emphasis", next: "@emphasis_underscore" }],
      [/\*(?!\*)/, { token: "emphasis", next: "@emphasis_asterisk" }],
      [/\*\*/, { token: "strong", next: "@strong" }],
      [/{/, { token: "delimiter.bracket", next: "@expression", nextEmbedded: "js" }]
    ],
    import: [[/'\s*(;|$)/, { token: "string", next: "@pop", nextEmbedded: "@pop" }]],
    expression: [
      [/{/, { token: "delimiter.bracket", next: "@expression" }],
      [/}/, { token: "delimiter.bracket", next: "@pop", nextEmbedded: "@pop" }]
    ],
    export: [[/^\s*$/, { token: "delimiter.bracket", next: "@pop", nextEmbedded: "@pop" }]],
    jsx: [
      [/\s+/, ""],
      [/(\w+)(=)("(?:[^"\\]|\\.)*")/, ["attribute.name", "operator", "string"]],
      [/(\w+)(=)('(?:[^'\\]|\\.)*')/, ["attribute.name", "operator", "string"]],
      [/(\w+(?=\s|>|={|$))/, ["attribute.name"]],
      [/={/, { token: "delimiter.bracket", next: "@expression", nextEmbedded: "js" }],
      [/>/, { token: "type.identifier", next: "@pop" }]
    ],
    header: [
      [/.$/, { token: "keyword", next: "@pop" }],
      { include: "content" },
      [/./, { token: "keyword" }]
    ],
    strong: [
      [/\*\*/, { token: "strong", next: "@pop" }],
      { include: "content" },
      [/./, { token: "strong" }]
    ],
    emphasis_underscore: [
      [/_/, { token: "emphasis", next: "@pop" }],
      { include: "content" },
      [/./, { token: "emphasis" }]
    ],
    emphasis_asterisk: [
      [/\*(?!\*)/, { token: "emphasis", next: "@pop" }],
      { include: "content" },
      [/./, { token: "emphasis" }]
    ],
    frontmatter: [[/^---$/, { token: "meta.content", nextEmbedded: "@pop", next: "@pop" }]],
    codeblock_highlight_backtick: [
      [/\s*`{3,}\s*$/, { token: "string", next: "@pop", nextEmbedded: "@pop" }],
      [/.*$/, "variable.source"]
    ],
    codeblock_highlight_tilde: [
      [/\s*~{3,}\s*$/, { token: "string", next: "@pop", nextEmbedded: "@pop" }],
      [/.*$/, "variable.source"]
    ],
    codeblock_backtick: [
      [/\s*`{3,}\s*$/, { token: "string", next: "@pop" }],
      [/.*$/, "variable.source"]
    ],
    codeblock_tilde: [
      [/\s*~{3,}\s*$/, { token: "string", next: "@pop" }],
      [/.*$/, "variable.source"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/mips/mips.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/mips/mips.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/mips/mips.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\#%\^\&\*\(\)\=\$\-\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    blockComment: ["###", "###"],
    lineComment: "#"
  },
  folding: {
    markers: {
      start: new RegExp("^\\s*#region\\b"),
      end: new RegExp("^\\s*#endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  ignoreCase: false,
  tokenPostfix: ".mips",
  regEx: /\/(?!\/\/)(?:[^\/\\]|\\.)*\/[igm]*/,
  keywords: [
    ".data",
    ".text",
    "syscall",
    "trap",
    "add",
    "addu",
    "addi",
    "addiu",
    "and",
    "andi",
    "div",
    "divu",
    "mult",
    "multu",
    "nor",
    "or",
    "ori",
    "sll",
    "slv",
    "sra",
    "srav",
    "srl",
    "srlv",
    "sub",
    "subu",
    "xor",
    "xori",
    "lhi",
    "lho",
    "lhi",
    "llo",
    "slt",
    "slti",
    "sltu",
    "sltiu",
    "beq",
    "bgtz",
    "blez",
    "bne",
    "j",
    "jal",
    "jalr",
    "jr",
    "lb",
    "lbu",
    "lh",
    "lhu",
    "lw",
    "li",
    "la",
    "sb",
    "sh",
    "sw",
    "mfhi",
    "mflo",
    "mthi",
    "mtlo",
    "move"
  ],
  symbols: /[\.,\:]+/,
  escapes: /\\(?:[abfnrtv\\"'$]|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/\$[a-zA-Z_]\w*/, "variable.predefined"],
      [
        /[.a-zA-Z_]\w*/,
        {
          cases: {
            this: "variable.predefined",
            "@keywords": { token: "keyword.$0" },
            "@default": ""
          }
        }
      ],
      [/[ \t\r\n]+/, ""],
      [/#.*$/, "comment"],
      ["///", { token: "regexp", next: "@hereregexp" }],
      [/^(\s*)(@regEx)/, ["", "regexp"]],
      [/(\,)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/(\:)(\s*)(@regEx)/, ["delimiter", "", "regexp"]],
      [/@symbols/, "delimiter"],
      [/\d+[eE]([\-+]?\d+)?/, "number.float"],
      [/\d+\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F]+/, "number.hex"],
      [/0[0-7]+(?!\d)/, "number.octal"],
      [/\d+/, "number"],
      [/[,.]/, "delimiter"],
      [/"""/, "string", '@herestring."""'],
      [/'''/, "string", "@herestring.'''"],
      [
        /"/,
        {
          cases: {
            "@eos": "string",
            "@default": { token: "string", next: '@string."' }
          }
        }
      ],
      [
        /'/,
        {
          cases: {
            "@eos": "string",
            "@default": { token: "string", next: "@string.'" }
          }
        }
      ]
    ],
    string: [
      [/[^"'\#\\]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\./, "string.escape.invalid"],
      [/\./, "string.escape.invalid"],
      [
        /#{/,
        {
          cases: {
            '$S2=="': {
              token: "string",
              next: "root.interpolatedstring"
            },
            "@default": "string"
          }
        }
      ],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      [/#/, "string"]
    ],
    herestring: [
      [
        /("""|''')/,
        {
          cases: {
            "$1==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      [/[^#\\'"]+/, "string"],
      [/['"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\./, "string.escape.invalid"],
      [/#{/, { token: "string.quote", next: "root.interpolatedstring" }],
      [/#/, "string"]
    ],
    comment: [
      [/[^#]+/, "comment"],
      [/#/, "comment"]
    ],
    hereregexp: [
      [/[^\\\/#]+/, "regexp"],
      [/\\./, "regexp"],
      [/#.*$/, "comment"],
      ["///[igm]*", { token: "regexp", next: "@pop" }],
      [/\//, "regexp"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/msdax/msdax.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/msdax/msdax.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/msdax/msdax.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["[", "]"],
    ["(", ")"],
    ["{", "}"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] },
    { open: "{", close: "}", notIn: ["string", "comment"] }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".msdax",
  ignoreCase: true,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "{", close: "}", token: "delimiter.brackets" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "VAR",
    "RETURN",
    "NOT",
    "EVALUATE",
    "DATATABLE",
    "ORDER",
    "BY",
    "START",
    "AT",
    "DEFINE",
    "MEASURE",
    "ASC",
    "DESC",
    "IN",
    "BOOLEAN",
    "DOUBLE",
    "INTEGER",
    "DATETIME",
    "CURRENCY",
    "STRING"
  ],
  functions: [
    "CLOSINGBALANCEMONTH",
    "CLOSINGBALANCEQUARTER",
    "CLOSINGBALANCEYEAR",
    "DATEADD",
    "DATESBETWEEN",
    "DATESINPERIOD",
    "DATESMTD",
    "DATESQTD",
    "DATESYTD",
    "ENDOFMONTH",
    "ENDOFQUARTER",
    "ENDOFYEAR",
    "FIRSTDATE",
    "FIRSTNONBLANK",
    "LASTDATE",
    "LASTNONBLANK",
    "NEXTDAY",
    "NEXTMONTH",
    "NEXTQUARTER",
    "NEXTYEAR",
    "OPENINGBALANCEMONTH",
    "OPENINGBALANCEQUARTER",
    "OPENINGBALANCEYEAR",
    "PARALLELPERIOD",
    "PREVIOUSDAY",
    "PREVIOUSMONTH",
    "PREVIOUSQUARTER",
    "PREVIOUSYEAR",
    "SAMEPERIODLASTYEAR",
    "STARTOFMONTH",
    "STARTOFQUARTER",
    "STARTOFYEAR",
    "TOTALMTD",
    "TOTALQTD",
    "TOTALYTD",
    "ADDCOLUMNS",
    "ADDMISSINGITEMS",
    "ALL",
    "ALLEXCEPT",
    "ALLNOBLANKROW",
    "ALLSELECTED",
    "CALCULATE",
    "CALCULATETABLE",
    "CALENDAR",
    "CALENDARAUTO",
    "CROSSFILTER",
    "CROSSJOIN",
    "CURRENTGROUP",
    "DATATABLE",
    "DETAILROWS",
    "DISTINCT",
    "EARLIER",
    "EARLIEST",
    "EXCEPT",
    "FILTER",
    "FILTERS",
    "GENERATE",
    "GENERATEALL",
    "GROUPBY",
    "IGNORE",
    "INTERSECT",
    "ISONORAFTER",
    "KEEPFILTERS",
    "LOOKUPVALUE",
    "NATURALINNERJOIN",
    "NATURALLEFTOUTERJOIN",
    "RELATED",
    "RELATEDTABLE",
    "ROLLUP",
    "ROLLUPADDISSUBTOTAL",
    "ROLLUPGROUP",
    "ROLLUPISSUBTOTAL",
    "ROW",
    "SAMPLE",
    "SELECTCOLUMNS",
    "SUBSTITUTEWITHINDEX",
    "SUMMARIZE",
    "SUMMARIZECOLUMNS",
    "TOPN",
    "TREATAS",
    "UNION",
    "USERELATIONSHIP",
    "VALUES",
    "SUM",
    "SUMX",
    "PATH",
    "PATHCONTAINS",
    "PATHITEM",
    "PATHITEMREVERSE",
    "PATHLENGTH",
    "AVERAGE",
    "AVERAGEA",
    "AVERAGEX",
    "COUNT",
    "COUNTA",
    "COUNTAX",
    "COUNTBLANK",
    "COUNTROWS",
    "COUNTX",
    "DISTINCTCOUNT",
    "DIVIDE",
    "GEOMEAN",
    "GEOMEANX",
    "MAX",
    "MAXA",
    "MAXX",
    "MEDIAN",
    "MEDIANX",
    "MIN",
    "MINA",
    "MINX",
    "PERCENTILE.EXC",
    "PERCENTILE.INC",
    "PERCENTILEX.EXC",
    "PERCENTILEX.INC",
    "PRODUCT",
    "PRODUCTX",
    "RANK.EQ",
    "RANKX",
    "STDEV.P",
    "STDEV.S",
    "STDEVX.P",
    "STDEVX.S",
    "VAR.P",
    "VAR.S",
    "VARX.P",
    "VARX.S",
    "XIRR",
    "XNPV",
    "DATE",
    "DATEDIFF",
    "DATEVALUE",
    "DAY",
    "EDATE",
    "EOMONTH",
    "HOUR",
    "MINUTE",
    "MONTH",
    "NOW",
    "SECOND",
    "TIME",
    "TIMEVALUE",
    "TODAY",
    "WEEKDAY",
    "WEEKNUM",
    "YEAR",
    "YEARFRAC",
    "CONTAINS",
    "CONTAINSROW",
    "CUSTOMDATA",
    "ERROR",
    "HASONEFILTER",
    "HASONEVALUE",
    "ISBLANK",
    "ISCROSSFILTERED",
    "ISEMPTY",
    "ISERROR",
    "ISEVEN",
    "ISFILTERED",
    "ISLOGICAL",
    "ISNONTEXT",
    "ISNUMBER",
    "ISODD",
    "ISSUBTOTAL",
    "ISTEXT",
    "USERNAME",
    "USERPRINCIPALNAME",
    "AND",
    "FALSE",
    "IF",
    "IFERROR",
    "NOT",
    "OR",
    "SWITCH",
    "TRUE",
    "ABS",
    "ACOS",
    "ACOSH",
    "ACOT",
    "ACOTH",
    "ASIN",
    "ASINH",
    "ATAN",
    "ATANH",
    "BETA.DIST",
    "BETA.INV",
    "CEILING",
    "CHISQ.DIST",
    "CHISQ.DIST.RT",
    "CHISQ.INV",
    "CHISQ.INV.RT",
    "COMBIN",
    "COMBINA",
    "CONFIDENCE.NORM",
    "CONFIDENCE.T",
    "COS",
    "COSH",
    "COT",
    "COTH",
    "CURRENCY",
    "DEGREES",
    "EVEN",
    "EXP",
    "EXPON.DIST",
    "FACT",
    "FLOOR",
    "GCD",
    "INT",
    "ISO.CEILING",
    "LCM",
    "LN",
    "LOG",
    "LOG10",
    "MOD",
    "MROUND",
    "ODD",
    "PERMUT",
    "PI",
    "POISSON.DIST",
    "POWER",
    "QUOTIENT",
    "RADIANS",
    "RAND",
    "RANDBETWEEN",
    "ROUND",
    "ROUNDDOWN",
    "ROUNDUP",
    "SIGN",
    "SIN",
    "SINH",
    "SQRT",
    "SQRTPI",
    "TAN",
    "TANH",
    "TRUNC",
    "BLANK",
    "CONCATENATE",
    "CONCATENATEX",
    "EXACT",
    "FIND",
    "FIXED",
    "FORMAT",
    "LEFT",
    "LEN",
    "LOWER",
    "MID",
    "REPLACE",
    "REPT",
    "RIGHT",
    "SEARCH",
    "SUBSTITUTE",
    "TRIM",
    "UNICHAR",
    "UNICODE",
    "UPPER",
    "VALUE"
  ],
  tokenizer: {
    root: [
      { include: "@comments" },
      { include: "@whitespace" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@complexIdentifiers" },
      [/[;,.]/, "delimiter"],
      [/[({})]/, "@brackets"],
      [
        /[a-z_][a-zA-Z0-9_]*/,
        {
          cases: {
            "@keywords": "keyword",
            "@functions": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=!%&+\-*/|~^]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      [/\/\/+.*/, "comment"],
      [/\/\*/, { token: "comment.quote", next: "@comment" }]
    ],
    comment: [
      [/[^*/]+/, "comment"],
      [/\*\//, { token: "comment.quote", next: "@pop" }],
      [/./, "comment"]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*/, "number"],
      [/[$][+-]*\d*(\.\d*)?/, "number"],
      [/((\d+(\.\d*)?)|(\.\d+))([eE][\-+]?\d+)?/, "number"]
    ],
    strings: [
      [/N"/, { token: "string", next: "@string" }],
      [/"/, { token: "string", next: "@string" }]
    ],
    string: [
      [/[^"]+/, "string"],
      [/""/, "string"],
      [/"/, { token: "string", next: "@pop" }]
    ],
    complexIdentifiers: [
      [/\[/, { token: "identifier.quote", next: "@bracketedIdentifier" }],
      [/'/, { token: "identifier.quote", next: "@quotedIdentifier" }]
    ],
    bracketedIdentifier: [
      [/[^\]]+/, "identifier"],
      [/]]/, "identifier"],
      [/]/, { token: "identifier.quote", next: "@pop" }]
    ],
    quotedIdentifier: [
      [/[^']+/, "identifier"],
      [/''/, "identifier"],
      [/'/, { token: "identifier.quote", next: "@pop" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/mysql/mysql.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/mysql/mysql.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/mysql/mysql.ts
var conf = {
  comments: {
    lineComment: "--",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sql",
  ignoreCase: true,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "ACCESSIBLE",
    "ADD",
    "ALL",
    "ALTER",
    "ANALYZE",
    "AND",
    "AS",
    "ASC",
    "ASENSITIVE",
    "BEFORE",
    "BETWEEN",
    "BIGINT",
    "BINARY",
    "BLOB",
    "BOTH",
    "BY",
    "CALL",
    "CASCADE",
    "CASE",
    "CHANGE",
    "CHAR",
    "CHARACTER",
    "CHECK",
    "COLLATE",
    "COLUMN",
    "CONDITION",
    "CONSTRAINT",
    "CONTINUE",
    "CONVERT",
    "CREATE",
    "CROSS",
    "CUBE",
    "CUME_DIST",
    "CURRENT_DATE",
    "CURRENT_TIME",
    "CURRENT_TIMESTAMP",
    "CURRENT_USER",
    "CURSOR",
    "DATABASE",
    "DATABASES",
    "DAY_HOUR",
    "DAY_MICROSECOND",
    "DAY_MINUTE",
    "DAY_SECOND",
    "DEC",
    "DECIMAL",
    "DECLARE",
    "DEFAULT",
    "DELAYED",
    "DELETE",
    "DENSE_RANK",
    "DESC",
    "DESCRIBE",
    "DETERMINISTIC",
    "DISTINCT",
    "DISTINCTROW",
    "DIV",
    "DOUBLE",
    "DROP",
    "DUAL",
    "EACH",
    "ELSE",
    "ELSEIF",
    "EMPTY",
    "ENCLOSED",
    "ESCAPED",
    "EXCEPT",
    "EXISTS",
    "EXIT",
    "EXPLAIN",
    "FALSE",
    "FETCH",
    "FIRST_VALUE",
    "FLOAT",
    "FLOAT4",
    "FLOAT8",
    "FOR",
    "FORCE",
    "FOREIGN",
    "FROM",
    "FULLTEXT",
    "FUNCTION",
    "GENERATED",
    "GET",
    "GRANT",
    "GROUP",
    "GROUPING",
    "GROUPS",
    "HAVING",
    "HIGH_PRIORITY",
    "HOUR_MICROSECOND",
    "HOUR_MINUTE",
    "HOUR_SECOND",
    "IF",
    "IGNORE",
    "IN",
    "INDEX",
    "INFILE",
    "INNER",
    "INOUT",
    "INSENSITIVE",
    "INSERT",
    "INT",
    "INT1",
    "INT2",
    "INT3",
    "INT4",
    "INT8",
    "INTEGER",
    "INTERVAL",
    "INTO",
    "IO_AFTER_GTIDS",
    "IO_BEFORE_GTIDS",
    "IS",
    "ITERATE",
    "JOIN",
    "JSON_TABLE",
    "KEY",
    "KEYS",
    "KILL",
    "LAG",
    "LAST_VALUE",
    "LATERAL",
    "LEAD",
    "LEADING",
    "LEAVE",
    "LEFT",
    "LIKE",
    "LIMIT",
    "LINEAR",
    "LINES",
    "LOAD",
    "LOCALTIME",
    "LOCALTIMESTAMP",
    "LOCK",
    "LONG",
    "LONGBLOB",
    "LONGTEXT",
    "LOOP",
    "LOW_PRIORITY",
    "MASTER_BIND",
    "MASTER_SSL_VERIFY_SERVER_CERT",
    "MATCH",
    "MAXVALUE",
    "MEDIUMBLOB",
    "MEDIUMINT",
    "MEDIUMTEXT",
    "MIDDLEINT",
    "MINUTE_MICROSECOND",
    "MINUTE_SECOND",
    "MOD",
    "MODIFIES",
    "NATURAL",
    "NOT",
    "NO_WRITE_TO_BINLOG",
    "NTH_VALUE",
    "NTILE",
    "NULL",
    "NUMERIC",
    "OF",
    "ON",
    "OPTIMIZE",
    "OPTIMIZER_COSTS",
    "OPTION",
    "OPTIONALLY",
    "OR",
    "ORDER",
    "OUT",
    "OUTER",
    "OUTFILE",
    "OVER",
    "PARTITION",
    "PERCENT_RANK",
    "PRECISION",
    "PRIMARY",
    "PROCEDURE",
    "PURGE",
    "RANGE",
    "RANK",
    "READ",
    "READS",
    "READ_WRITE",
    "REAL",
    "RECURSIVE",
    "REFERENCES",
    "REGEXP",
    "RELEASE",
    "RENAME",
    "REPEAT",
    "REPLACE",
    "REQUIRE",
    "RESIGNAL",
    "RESTRICT",
    "RETURN",
    "REVOKE",
    "RIGHT",
    "RLIKE",
    "ROW",
    "ROWS",
    "ROW_NUMBER",
    "SCHEMA",
    "SCHEMAS",
    "SECOND_MICROSECOND",
    "SELECT",
    "SENSITIVE",
    "SEPARATOR",
    "SET",
    "SHOW",
    "SIGNAL",
    "SMALLINT",
    "SPATIAL",
    "SPECIFIC",
    "SQL",
    "SQLEXCEPTION",
    "SQLSTATE",
    "SQLWARNING",
    "SQL_BIG_RESULT",
    "SQL_CALC_FOUND_ROWS",
    "SQL_SMALL_RESULT",
    "SSL",
    "STARTING",
    "STORED",
    "STRAIGHT_JOIN",
    "SYSTEM",
    "TABLE",
    "TERMINATED",
    "THEN",
    "TINYBLOB",
    "TINYINT",
    "TINYTEXT",
    "TO",
    "TRAILING",
    "TRIGGER",
    "TRUE",
    "UNDO",
    "UNION",
    "UNIQUE",
    "UNLOCK",
    "UNSIGNED",
    "UPDATE",
    "USAGE",
    "USE",
    "USING",
    "UTC_DATE",
    "UTC_TIME",
    "UTC_TIMESTAMP",
    "VALUES",
    "VARBINARY",
    "VARCHAR",
    "VARCHARACTER",
    "VARYING",
    "VIRTUAL",
    "WHEN",
    "WHERE",
    "WHILE",
    "WINDOW",
    "WITH",
    "WRITE",
    "XOR",
    "YEAR_MONTH",
    "ZEROFILL"
  ],
  operators: [
    "AND",
    "BETWEEN",
    "IN",
    "LIKE",
    "NOT",
    "OR",
    "IS",
    "NULL",
    "INTERSECT",
    "UNION",
    "INNER",
    "JOIN",
    "LEFT",
    "OUTER",
    "RIGHT"
  ],
  builtinFunctions: [
    "ABS",
    "ACOS",
    "ADDDATE",
    "ADDTIME",
    "AES_DECRYPT",
    "AES_ENCRYPT",
    "ANY_VALUE",
    "Area",
    "AsBinary",
    "AsWKB",
    "ASCII",
    "ASIN",
    "AsText",
    "AsWKT",
    "ASYMMETRIC_DECRYPT",
    "ASYMMETRIC_DERIVE",
    "ASYMMETRIC_ENCRYPT",
    "ASYMMETRIC_SIGN",
    "ASYMMETRIC_VERIFY",
    "ATAN",
    "ATAN2",
    "ATAN",
    "AVG",
    "BENCHMARK",
    "BIN",
    "BIT_AND",
    "BIT_COUNT",
    "BIT_LENGTH",
    "BIT_OR",
    "BIT_XOR",
    "Buffer",
    "CAST",
    "CEIL",
    "CEILING",
    "Centroid",
    "CHAR",
    "CHAR_LENGTH",
    "CHARACTER_LENGTH",
    "CHARSET",
    "COALESCE",
    "COERCIBILITY",
    "COLLATION",
    "COMPRESS",
    "CONCAT",
    "CONCAT_WS",
    "CONNECTION_ID",
    "Contains",
    "CONV",
    "CONVERT",
    "CONVERT_TZ",
    "ConvexHull",
    "COS",
    "COT",
    "COUNT",
    "CRC32",
    "CREATE_ASYMMETRIC_PRIV_KEY",
    "CREATE_ASYMMETRIC_PUB_KEY",
    "CREATE_DH_PARAMETERS",
    "CREATE_DIGEST",
    "Crosses",
    "CUME_DIST",
    "CURDATE",
    "CURRENT_DATE",
    "CURRENT_ROLE",
    "CURRENT_TIME",
    "CURRENT_TIMESTAMP",
    "CURRENT_USER",
    "CURTIME",
    "DATABASE",
    "DATE",
    "DATE_ADD",
    "DATE_FORMAT",
    "DATE_SUB",
    "DATEDIFF",
    "DAY",
    "DAYNAME",
    "DAYOFMONTH",
    "DAYOFWEEK",
    "DAYOFYEAR",
    "DECODE",
    "DEFAULT",
    "DEGREES",
    "DES_DECRYPT",
    "DES_ENCRYPT",
    "DENSE_RANK",
    "Dimension",
    "Disjoint",
    "Distance",
    "ELT",
    "ENCODE",
    "ENCRYPT",
    "EndPoint",
    "Envelope",
    "Equals",
    "EXP",
    "EXPORT_SET",
    "ExteriorRing",
    "EXTRACT",
    "ExtractValue",
    "FIELD",
    "FIND_IN_SET",
    "FIRST_VALUE",
    "FLOOR",
    "FORMAT",
    "FORMAT_BYTES",
    "FORMAT_PICO_TIME",
    "FOUND_ROWS",
    "FROM_BASE64",
    "FROM_DAYS",
    "FROM_UNIXTIME",
    "GEN_RANGE",
    "GEN_RND_EMAIL",
    "GEN_RND_PAN",
    "GEN_RND_SSN",
    "GEN_RND_US_PHONE",
    "GeomCollection",
    "GeomCollFromText",
    "GeometryCollectionFromText",
    "GeomCollFromWKB",
    "GeometryCollectionFromWKB",
    "GeometryCollection",
    "GeometryN",
    "GeometryType",
    "GeomFromText",
    "GeometryFromText",
    "GeomFromWKB",
    "GeometryFromWKB",
    "GET_FORMAT",
    "GET_LOCK",
    "GLength",
    "GREATEST",
    "GROUP_CONCAT",
    "GROUPING",
    "GTID_SUBSET",
    "GTID_SUBTRACT",
    "HEX",
    "HOUR",
    "ICU_VERSION",
    "IF",
    "IFNULL",
    "INET_ATON",
    "INET_NTOA",
    "INET6_ATON",
    "INET6_NTOA",
    "INSERT",
    "INSTR",
    "InteriorRingN",
    "Intersects",
    "INTERVAL",
    "IS_FREE_LOCK",
    "IS_IPV4",
    "IS_IPV4_COMPAT",
    "IS_IPV4_MAPPED",
    "IS_IPV6",
    "IS_USED_LOCK",
    "IS_UUID",
    "IsClosed",
    "IsEmpty",
    "ISNULL",
    "IsSimple",
    "JSON_APPEND",
    "JSON_ARRAY",
    "JSON_ARRAY_APPEND",
    "JSON_ARRAY_INSERT",
    "JSON_ARRAYAGG",
    "JSON_CONTAINS",
    "JSON_CONTAINS_PATH",
    "JSON_DEPTH",
    "JSON_EXTRACT",
    "JSON_INSERT",
    "JSON_KEYS",
    "JSON_LENGTH",
    "JSON_MERGE",
    "JSON_MERGE_PATCH",
    "JSON_MERGE_PRESERVE",
    "JSON_OBJECT",
    "JSON_OBJECTAGG",
    "JSON_OVERLAPS",
    "JSON_PRETTY",
    "JSON_QUOTE",
    "JSON_REMOVE",
    "JSON_REPLACE",
    "JSON_SCHEMA_VALID",
    "JSON_SCHEMA_VALIDATION_REPORT",
    "JSON_SEARCH",
    "JSON_SET",
    "JSON_STORAGE_FREE",
    "JSON_STORAGE_SIZE",
    "JSON_TABLE",
    "JSON_TYPE",
    "JSON_UNQUOTE",
    "JSON_VALID",
    "LAG",
    "LAST_DAY",
    "LAST_INSERT_ID",
    "LAST_VALUE",
    "LCASE",
    "LEAD",
    "LEAST",
    "LEFT",
    "LENGTH",
    "LineFromText",
    "LineStringFromText",
    "LineFromWKB",
    "LineStringFromWKB",
    "LineString",
    "LN",
    "LOAD_FILE",
    "LOCALTIME",
    "LOCALTIMESTAMP",
    "LOCATE",
    "LOG",
    "LOG10",
    "LOG2",
    "LOWER",
    "LPAD",
    "LTRIM",
    "MAKE_SET",
    "MAKEDATE",
    "MAKETIME",
    "MASK_INNER",
    "MASK_OUTER",
    "MASK_PAN",
    "MASK_PAN_RELAXED",
    "MASK_SSN",
    "MASTER_POS_WAIT",
    "MAX",
    "MBRContains",
    "MBRCoveredBy",
    "MBRCovers",
    "MBRDisjoint",
    "MBREqual",
    "MBREquals",
    "MBRIntersects",
    "MBROverlaps",
    "MBRTouches",
    "MBRWithin",
    "MD5",
    "MEMBER OF",
    "MICROSECOND",
    "MID",
    "MIN",
    "MINUTE",
    "MLineFromText",
    "MultiLineStringFromText",
    "MLineFromWKB",
    "MultiLineStringFromWKB",
    "MOD",
    "MONTH",
    "MONTHNAME",
    "MPointFromText",
    "MultiPointFromText",
    "MPointFromWKB",
    "MultiPointFromWKB",
    "MPolyFromText",
    "MultiPolygonFromText",
    "MPolyFromWKB",
    "MultiPolygonFromWKB",
    "MultiLineString",
    "MultiPoint",
    "MultiPolygon",
    "NAME_CONST",
    "NOT IN",
    "NOW",
    "NTH_VALUE",
    "NTILE",
    "NULLIF",
    "NumGeometries",
    "NumInteriorRings",
    "NumPoints",
    "OCT",
    "OCTET_LENGTH",
    "OLD_PASSWORD",
    "ORD",
    "Overlaps",
    "PASSWORD",
    "PERCENT_RANK",
    "PERIOD_ADD",
    "PERIOD_DIFF",
    "PI",
    "Point",
    "PointFromText",
    "PointFromWKB",
    "PointN",
    "PolyFromText",
    "PolygonFromText",
    "PolyFromWKB",
    "PolygonFromWKB",
    "Polygon",
    "POSITION",
    "POW",
    "POWER",
    "PS_CURRENT_THREAD_ID",
    "PS_THREAD_ID",
    "PROCEDURE ANALYSE",
    "QUARTER",
    "QUOTE",
    "RADIANS",
    "RAND",
    "RANDOM_BYTES",
    "RANK",
    "REGEXP_INSTR",
    "REGEXP_LIKE",
    "REGEXP_REPLACE",
    "REGEXP_REPLACE",
    "RELEASE_ALL_LOCKS",
    "RELEASE_LOCK",
    "REPEAT",
    "REPLACE",
    "REVERSE",
    "RIGHT",
    "ROLES_GRAPHML",
    "ROUND",
    "ROW_COUNT",
    "ROW_NUMBER",
    "RPAD",
    "RTRIM",
    "SCHEMA",
    "SEC_TO_TIME",
    "SECOND",
    "SESSION_USER",
    "SHA1",
    "SHA",
    "SHA2",
    "SIGN",
    "SIN",
    "SLEEP",
    "SOUNDEX",
    "SOURCE_POS_WAIT",
    "SPACE",
    "SQRT",
    "SRID",
    "ST_Area",
    "ST_AsBinary",
    "ST_AsWKB",
    "ST_AsGeoJSON",
    "ST_AsText",
    "ST_AsWKT",
    "ST_Buffer",
    "ST_Buffer_Strategy",
    "ST_Centroid",
    "ST_Collect",
    "ST_Contains",
    "ST_ConvexHull",
    "ST_Crosses",
    "ST_Difference",
    "ST_Dimension",
    "ST_Disjoint",
    "ST_Distance",
    "ST_Distance_Sphere",
    "ST_EndPoint",
    "ST_Envelope",
    "ST_Equals",
    "ST_ExteriorRing",
    "ST_FrechetDistance",
    "ST_GeoHash",
    "ST_GeomCollFromText",
    "ST_GeometryCollectionFromText",
    "ST_GeomCollFromTxt",
    "ST_GeomCollFromWKB",
    "ST_GeometryCollectionFromWKB",
    "ST_GeometryN",
    "ST_GeometryType",
    "ST_GeomFromGeoJSON",
    "ST_GeomFromText",
    "ST_GeometryFromText",
    "ST_GeomFromWKB",
    "ST_GeometryFromWKB",
    "ST_HausdorffDistance",
    "ST_InteriorRingN",
    "ST_Intersection",
    "ST_Intersects",
    "ST_IsClosed",
    "ST_IsEmpty",
    "ST_IsSimple",
    "ST_IsValid",
    "ST_LatFromGeoHash",
    "ST_Length",
    "ST_LineFromText",
    "ST_LineStringFromText",
    "ST_LineFromWKB",
    "ST_LineStringFromWKB",
    "ST_LineInterpolatePoint",
    "ST_LineInterpolatePoints",
    "ST_LongFromGeoHash",
    "ST_Longitude",
    "ST_MakeEnvelope",
    "ST_MLineFromText",
    "ST_MultiLineStringFromText",
    "ST_MLineFromWKB",
    "ST_MultiLineStringFromWKB",
    "ST_MPointFromText",
    "ST_MultiPointFromText",
    "ST_MPointFromWKB",
    "ST_MultiPointFromWKB",
    "ST_MPolyFromText",
    "ST_MultiPolygonFromText",
    "ST_MPolyFromWKB",
    "ST_MultiPolygonFromWKB",
    "ST_NumGeometries",
    "ST_NumInteriorRing",
    "ST_NumInteriorRings",
    "ST_NumPoints",
    "ST_Overlaps",
    "ST_PointAtDistance",
    "ST_PointFromGeoHash",
    "ST_PointFromText",
    "ST_PointFromWKB",
    "ST_PointN",
    "ST_PolyFromText",
    "ST_PolygonFromText",
    "ST_PolyFromWKB",
    "ST_PolygonFromWKB",
    "ST_Simplify",
    "ST_SRID",
    "ST_StartPoint",
    "ST_SwapXY",
    "ST_SymDifference",
    "ST_Touches",
    "ST_Transform",
    "ST_Union",
    "ST_Validate",
    "ST_Within",
    "ST_X",
    "ST_Y",
    "StartPoint",
    "STATEMENT_DIGEST",
    "STATEMENT_DIGEST_TEXT",
    "STD",
    "STDDEV",
    "STDDEV_POP",
    "STDDEV_SAMP",
    "STR_TO_DATE",
    "STRCMP",
    "SUBDATE",
    "SUBSTR",
    "SUBSTRING",
    "SUBSTRING_INDEX",
    "SUBTIME",
    "SUM",
    "SYSDATE",
    "SYSTEM_USER",
    "TAN",
    "TIME",
    "TIME_FORMAT",
    "TIME_TO_SEC",
    "TIMEDIFF",
    "TIMESTAMP",
    "TIMESTAMPADD",
    "TIMESTAMPDIFF",
    "TO_BASE64",
    "TO_DAYS",
    "TO_SECONDS",
    "Touches",
    "TRIM",
    "TRUNCATE",
    "UCASE",
    "UNCOMPRESS",
    "UNCOMPRESSED_LENGTH",
    "UNHEX",
    "UNIX_TIMESTAMP",
    "UpdateXML",
    "UPPER",
    "USER",
    "UTC_DATE",
    "UTC_TIME",
    "UTC_TIMESTAMP",
    "UUID",
    "UUID_SHORT",
    "UUID_TO_BIN",
    "VALIDATE_PASSWORD_STRENGTH",
    "VALUES",
    "VAR_POP",
    "VAR_SAMP",
    "VARIANCE",
    "VERSION",
    "WAIT_FOR_EXECUTED_GTID_SET",
    "WAIT_UNTIL_SQL_THREAD_AFTER_GTIDS",
    "WEEK",
    "WEEKDAY",
    "WEEKOFYEAR",
    "WEIGHT_STRING",
    "Within",
    "X",
    "Y",
    "YEAR",
    "YEARWEEK"
  ],
  builtinVariables: [],
  tokenizer: {
    root: [
      { include: "@comments" },
      { include: "@whitespace" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@complexIdentifiers" },
      { include: "@scopes" },
      [/[;,.]/, "delimiter"],
      [/[()]/, "@brackets"],
      [
        /[\w@]+/,
        {
          cases: {
            "@operators": "operator",
            "@builtinVariables": "predefined",
            "@builtinFunctions": "predefined",
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=!%&+\-*/|~^]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      [/--+.*/, "comment"],
      [/#+.*/, "comment"],
      [/\/\*/, { token: "comment.quote", next: "@comment" }]
    ],
    comment: [
      [/[^*/]+/, "comment"],
      [/\*\//, { token: "comment.quote", next: "@pop" }],
      [/./, "comment"]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*/, "number"],
      [/[$][+-]*\d*(\.\d*)?/, "number"],
      [/((\d+(\.\d*)?)|(\.\d+))([eE][\-+]?\d+)?/, "number"]
    ],
    strings: [
      [/'/, { token: "string", next: "@string" }],
      [/"/, { token: "string.double", next: "@stringDouble" }]
    ],
    string: [
      [/\\'/, "string"],
      [/[^']+/, "string"],
      [/''/, "string"],
      [/'/, { token: "string", next: "@pop" }]
    ],
    stringDouble: [
      [/[^"]+/, "string.double"],
      [/""/, "string.double"],
      [/"/, { token: "string.double", next: "@pop" }]
    ],
    complexIdentifiers: [[/`/, { token: "identifier.quote", next: "@quotedIdentifier" }]],
    quotedIdentifier: [
      [/[^`]+/, "identifier"],
      [/``/, "identifier"],
      [/`/, { token: "identifier.quote", next: "@pop" }]
    ],
    scopes: []
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/objective-c/objective-c.js":
/*!**************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/objective-c/objective-c.js ***!
  \**************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/objective-c/objective-c.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".objective-c",
  keywords: [
    "#import",
    "#include",
    "#define",
    "#else",
    "#endif",
    "#if",
    "#ifdef",
    "#ifndef",
    "#ident",
    "#undef",
    "@class",
    "@defs",
    "@dynamic",
    "@encode",
    "@end",
    "@implementation",
    "@interface",
    "@package",
    "@private",
    "@protected",
    "@property",
    "@protocol",
    "@public",
    "@selector",
    "@synthesize",
    "__declspec",
    "assign",
    "auto",
    "BOOL",
    "break",
    "bycopy",
    "byref",
    "case",
    "char",
    "Class",
    "const",
    "copy",
    "continue",
    "default",
    "do",
    "double",
    "else",
    "enum",
    "extern",
    "FALSE",
    "false",
    "float",
    "for",
    "goto",
    "if",
    "in",
    "int",
    "id",
    "inout",
    "IMP",
    "long",
    "nil",
    "nonatomic",
    "NULL",
    "oneway",
    "out",
    "private",
    "public",
    "protected",
    "readwrite",
    "readonly",
    "register",
    "return",
    "SEL",
    "self",
    "short",
    "signed",
    "sizeof",
    "static",
    "struct",
    "super",
    "switch",
    "typedef",
    "TRUE",
    "true",
    "union",
    "unsigned",
    "volatile",
    "void",
    "while"
  ],
  decpart: /\d(_?\d)*/,
  decimal: /0|@decpart/,
  tokenizer: {
    root: [
      { include: "@comments" },
      { include: "@whitespace" },
      { include: "@numbers" },
      { include: "@strings" },
      [/[,:;]/, "delimiter"],
      [/[{}\[\]()<>]/, "@brackets"],
      [
        /[a-zA-Z@#]\w*/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=\\+\\-\\*\\/\\^\\|\\~,]|and\\b|or\\b|not\\b]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      ["\\/\\*", "comment", "@comment"],
      ["\\/\\/+.*", "comment"]
    ],
    comment: [
      ["\\*\\/", "comment", "@pop"],
      [".", "comment"]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*(_?[0-9a-fA-F])*/, "number.hex"],
      [
        /@decimal((\.@decpart)?([eE][\-+]?@decpart)?)[fF]*/,
        {
          cases: {
            "(\\d)*": "number",
            $0: "number.float"
          }
        }
      ]
    ],
    strings: [
      [/'$/, "string.escape", "@popall"],
      [/'/, "string.escape", "@stringBody"],
      [/"$/, "string.escape", "@popall"],
      [/"/, "string.escape", "@dblStringBody"]
    ],
    stringBody: [
      [/[^\\']+$/, "string", "@popall"],
      [/[^\\']+/, "string"],
      [/\\./, "string"],
      [/'/, "string.escape", "@popall"],
      [/\\$/, "string"]
    ],
    dblStringBody: [
      [/[^\\"]+$/, "string", "@popall"],
      [/[^\\"]+/, "string"],
      [/\\./, "string"],
      [/"/, "string.escape", "@popall"],
      [/\\$/, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/pascal/pascal.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/pascal/pascal.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/pascal/pascal.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\#\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["{", "}"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*\\{\\$REGION(\\s\\'.*\\')?\\}"),
      end: new RegExp("^\\s*\\{\\$ENDREGION\\}")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".pascal",
  ignoreCase: true,
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  keywords: [
    "absolute",
    "abstract",
    "all",
    "and_then",
    "array",
    "as",
    "asm",
    "attribute",
    "begin",
    "bindable",
    "case",
    "class",
    "const",
    "contains",
    "default",
    "div",
    "else",
    "end",
    "except",
    "exports",
    "external",
    "far",
    "file",
    "finalization",
    "finally",
    "forward",
    "generic",
    "goto",
    "if",
    "implements",
    "import",
    "in",
    "index",
    "inherited",
    "initialization",
    "interrupt",
    "is",
    "label",
    "library",
    "mod",
    "module",
    "name",
    "near",
    "not",
    "object",
    "of",
    "on",
    "only",
    "operator",
    "or_else",
    "otherwise",
    "override",
    "package",
    "packed",
    "pow",
    "private",
    "program",
    "protected",
    "public",
    "published",
    "interface",
    "implementation",
    "qualified",
    "read",
    "record",
    "resident",
    "requires",
    "resourcestring",
    "restricted",
    "segment",
    "set",
    "shl",
    "shr",
    "specialize",
    "stored",
    "strict",
    "then",
    "threadvar",
    "to",
    "try",
    "type",
    "unit",
    "uses",
    "var",
    "view",
    "virtual",
    "dynamic",
    "overload",
    "reintroduce",
    "with",
    "write",
    "xor",
    "true",
    "false",
    "procedure",
    "function",
    "constructor",
    "destructor",
    "property",
    "break",
    "continue",
    "exit",
    "abort",
    "while",
    "do",
    "for",
    "raise",
    "repeat",
    "until"
  ],
  typeKeywords: [
    "boolean",
    "double",
    "byte",
    "integer",
    "shortint",
    "char",
    "longint",
    "float",
    "string"
  ],
  operators: [
    "=",
    ">",
    "<",
    "<=",
    ">=",
    "<>",
    ":",
    ":=",
    "and",
    "or",
    "+",
    "-",
    "*",
    "/",
    "@",
    "&",
    "^",
    "%"
  ],
  symbols: /[=><:@\^&|+\-*\/\^%]+/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_][\w]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/\$[0-9a-fA-F]{1,16}/, "number.hex"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/'/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/'/, "string.invalid"],
      [/\#\d+/, "string"]
    ],
    comment: [
      [/[^\*\}]+/, "comment"],
      [/\}/, "comment", "@pop"],
      [/[\{]/, "comment"]
    ],
    string: [
      [/[^\\']+/, "string"],
      [/\\./, "string.escape.invalid"],
      [/'/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\{/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/pascaligo/pascaligo.js":
/*!**********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/pascaligo/pascaligo.js ***!
  \**********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/pascaligo/pascaligo.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["(*", "*)"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".pascaligo",
  ignoreCase: true,
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  keywords: [
    "begin",
    "block",
    "case",
    "const",
    "else",
    "end",
    "fail",
    "for",
    "from",
    "function",
    "if",
    "is",
    "nil",
    "of",
    "remove",
    "return",
    "skip",
    "then",
    "type",
    "var",
    "while",
    "with",
    "option",
    "None",
    "transaction"
  ],
  typeKeywords: [
    "bool",
    "int",
    "list",
    "map",
    "nat",
    "record",
    "string",
    "unit",
    "address",
    "map",
    "mtz",
    "xtz"
  ],
  operators: [
    "=",
    ">",
    "<",
    "<=",
    ">=",
    "<>",
    ":",
    ":=",
    "and",
    "mod",
    "or",
    "+",
    "-",
    "*",
    "/",
    "@",
    "&",
    "^",
    "%"
  ],
  symbols: /[=><:@\^&|+\-*\/\^%]+/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_][\w]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/\$[0-9a-fA-F]{1,16}/, "number.hex"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/'/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/'/, "string.invalid"],
      [/\#\d+/, "string"]
    ],
    comment: [
      [/[^\(\*]+/, "comment"],
      [/\*\)/, "comment", "@pop"],
      [/\(\*/, "comment"]
    ],
    string: [
      [/[^\\']+/, "string"],
      [/\\./, "string.escape.invalid"],
      [/'/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\(\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/perl/perl.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/perl/perl.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/perl/perl.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".perl",
  brackets: [
    { token: "delimiter.bracket", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" }
  ],
  keywords: [
    "__DATA__",
    "else",
    "lock",
    "__END__",
    "elsif",
    "lt",
    "__FILE__",
    "eq",
    "__LINE__",
    "exp",
    "ne",
    "sub",
    "__PACKAGE__",
    "for",
    "no",
    "and",
    "foreach",
    "or",
    "unless",
    "cmp",
    "ge",
    "package",
    "until",
    "continue",
    "gt",
    "while",
    "CORE",
    "if",
    "xor",
    "do",
    "le",
    "__DIE__",
    "__WARN__"
  ],
  builtinFunctions: [
    "-A",
    "END",
    "length",
    "setpgrp",
    "-B",
    "endgrent",
    "link",
    "setpriority",
    "-b",
    "endhostent",
    "listen",
    "setprotoent",
    "-C",
    "endnetent",
    "local",
    "setpwent",
    "-c",
    "endprotoent",
    "localtime",
    "setservent",
    "-d",
    "endpwent",
    "log",
    "setsockopt",
    "-e",
    "endservent",
    "lstat",
    "shift",
    "-f",
    "eof",
    "map",
    "shmctl",
    "-g",
    "eval",
    "mkdir",
    "shmget",
    "-k",
    "exec",
    "msgctl",
    "shmread",
    "-l",
    "exists",
    "msgget",
    "shmwrite",
    "-M",
    "exit",
    "msgrcv",
    "shutdown",
    "-O",
    "fcntl",
    "msgsnd",
    "sin",
    "-o",
    "fileno",
    "my",
    "sleep",
    "-p",
    "flock",
    "next",
    "socket",
    "-r",
    "fork",
    "not",
    "socketpair",
    "-R",
    "format",
    "oct",
    "sort",
    "-S",
    "formline",
    "open",
    "splice",
    "-s",
    "getc",
    "opendir",
    "split",
    "-T",
    "getgrent",
    "ord",
    "sprintf",
    "-t",
    "getgrgid",
    "our",
    "sqrt",
    "-u",
    "getgrnam",
    "pack",
    "srand",
    "-w",
    "gethostbyaddr",
    "pipe",
    "stat",
    "-W",
    "gethostbyname",
    "pop",
    "state",
    "-X",
    "gethostent",
    "pos",
    "study",
    "-x",
    "getlogin",
    "print",
    "substr",
    "-z",
    "getnetbyaddr",
    "printf",
    "symlink",
    "abs",
    "getnetbyname",
    "prototype",
    "syscall",
    "accept",
    "getnetent",
    "push",
    "sysopen",
    "alarm",
    "getpeername",
    "quotemeta",
    "sysread",
    "atan2",
    "getpgrp",
    "rand",
    "sysseek",
    "AUTOLOAD",
    "getppid",
    "read",
    "system",
    "BEGIN",
    "getpriority",
    "readdir",
    "syswrite",
    "bind",
    "getprotobyname",
    "readline",
    "tell",
    "binmode",
    "getprotobynumber",
    "readlink",
    "telldir",
    "bless",
    "getprotoent",
    "readpipe",
    "tie",
    "break",
    "getpwent",
    "recv",
    "tied",
    "caller",
    "getpwnam",
    "redo",
    "time",
    "chdir",
    "getpwuid",
    "ref",
    "times",
    "CHECK",
    "getservbyname",
    "rename",
    "truncate",
    "chmod",
    "getservbyport",
    "require",
    "uc",
    "chomp",
    "getservent",
    "reset",
    "ucfirst",
    "chop",
    "getsockname",
    "return",
    "umask",
    "chown",
    "getsockopt",
    "reverse",
    "undef",
    "chr",
    "glob",
    "rewinddir",
    "UNITCHECK",
    "chroot",
    "gmtime",
    "rindex",
    "unlink",
    "close",
    "goto",
    "rmdir",
    "unpack",
    "closedir",
    "grep",
    "say",
    "unshift",
    "connect",
    "hex",
    "scalar",
    "untie",
    "cos",
    "index",
    "seek",
    "use",
    "crypt",
    "INIT",
    "seekdir",
    "utime",
    "dbmclose",
    "int",
    "select",
    "values",
    "dbmopen",
    "ioctl",
    "semctl",
    "vec",
    "defined",
    "join",
    "semget",
    "wait",
    "delete",
    "keys",
    "semop",
    "waitpid",
    "DESTROY",
    "kill",
    "send",
    "wantarray",
    "die",
    "last",
    "setgrent",
    "warn",
    "dump",
    "lc",
    "sethostent",
    "write",
    "each",
    "lcfirst",
    "setnetent"
  ],
  builtinFileHandlers: ["ARGV", "STDERR", "STDOUT", "ARGVOUT", "STDIN", "ENV"],
  builtinVariables: [
    "$!",
    "$^RE_TRIE_MAXBUF",
    "$LAST_REGEXP_CODE_RESULT",
    '$"',
    "$^S",
    "$LIST_SEPARATOR",
    "$#",
    "$^T",
    "$MATCH",
    "$$",
    "$^TAINT",
    "$MULTILINE_MATCHING",
    "$%",
    "$^UNICODE",
    "$NR",
    "$&",
    "$^UTF8LOCALE",
    "$OFMT",
    "$'",
    "$^V",
    "$OFS",
    "$(",
    "$^W",
    "$ORS",
    "$)",
    "$^WARNING_BITS",
    "$OS_ERROR",
    "$*",
    "$^WIDE_SYSTEM_CALLS",
    "$OSNAME",
    "$+",
    "$^X",
    "$OUTPUT_AUTO_FLUSH",
    "$,",
    "$_",
    "$OUTPUT_FIELD_SEPARATOR",
    "$-",
    "$`",
    "$OUTPUT_RECORD_SEPARATOR",
    "$.",
    "$a",
    "$PERL_VERSION",
    "$/",
    "$ACCUMULATOR",
    "$PERLDB",
    "$0",
    "$ARG",
    "$PID",
    "$:",
    "$ARGV",
    "$POSTMATCH",
    "$;",
    "$b",
    "$PREMATCH",
    "$<",
    "$BASETIME",
    "$PROCESS_ID",
    "$=",
    "$CHILD_ERROR",
    "$PROGRAM_NAME",
    "$>",
    "$COMPILING",
    "$REAL_GROUP_ID",
    "$?",
    "$DEBUGGING",
    "$REAL_USER_ID",
    "$@",
    "$EFFECTIVE_GROUP_ID",
    "$RS",
    "$[",
    "$EFFECTIVE_USER_ID",
    "$SUBSCRIPT_SEPARATOR",
    "$\\",
    "$EGID",
    "$SUBSEP",
    "$]",
    "$ERRNO",
    "$SYSTEM_FD_MAX",
    "$^",
    "$EUID",
    "$UID",
    "$^A",
    "$EVAL_ERROR",
    "$WARNING",
    "$^C",
    "$EXCEPTIONS_BEING_CAUGHT",
    "$|",
    "$^CHILD_ERROR_NATIVE",
    "$EXECUTABLE_NAME",
    "$~",
    "$^D",
    "$EXTENDED_OS_ERROR",
    "%!",
    "$^E",
    "$FORMAT_FORMFEED",
    "%^H",
    "$^ENCODING",
    "$FORMAT_LINE_BREAK_CHARACTERS",
    "%ENV",
    "$^F",
    "$FORMAT_LINES_LEFT",
    "%INC",
    "$^H",
    "$FORMAT_LINES_PER_PAGE",
    "%OVERLOAD",
    "$^I",
    "$FORMAT_NAME",
    "%SIG",
    "$^L",
    "$FORMAT_PAGE_NUMBER",
    "@+",
    "$^M",
    "$FORMAT_TOP_NAME",
    "@-",
    "$^N",
    "$GID",
    "@_",
    "$^O",
    "$INPLACE_EDIT",
    "@ARGV",
    "$^OPEN",
    "$INPUT_LINE_NUMBER",
    "@INC",
    "$^P",
    "$INPUT_RECORD_SEPARATOR",
    "@LAST_MATCH_START",
    "$^R",
    "$LAST_MATCH_END",
    "$^RE_DEBUG_FLAGS",
    "$LAST_PAREN_MATCH"
  ],
  symbols: /[:+\-\^*$&%@=<>!?|\/~\.]/,
  quoteLikeOps: ["qr", "m", "s", "q", "qq", "qx", "qw", "tr", "y"],
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      [
        /[a-zA-Z\-_][\w\-_]*/,
        {
          cases: {
            "@keywords": "keyword",
            "@builtinFunctions": "type.identifier",
            "@builtinFileHandlers": "variable.predefined",
            "@quoteLikeOps": {
              token: "@rematch",
              next: "quotedConstructs"
            },
            "@default": ""
          }
        }
      ],
      [
        /[\$@%][*@#?\+\-\$!\w\\\^><~:;\.]+/,
        {
          cases: {
            "@builtinVariables": "variable.predefined",
            "@default": "variable"
          }
        }
      ],
      { include: "@strings" },
      { include: "@dblStrings" },
      { include: "@perldoc" },
      { include: "@heredoc" },
      [/[{}\[\]()]/, "@brackets"],
      [/[\/](?:(?:\[(?:\\]|[^\]])+\])|(?:\\\/|[^\]\/]))*[\/]\w*\s*(?=[).,;]|$)/, "regexp"],
      [/@symbols/, "operators"],
      { include: "@numbers" },
      [/[,;]/, "delimiter"]
    ],
    whitespace: [
      [/\s+/, "white"],
      [/(^#!.*$)/, "metatag"],
      [/(^#.*$)/, "comment"]
    ],
    numbers: [
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F_]*[0-9a-fA-F]/, "number.hex"],
      [/\d+/, "number"]
    ],
    strings: [[/'/, "string", "@stringBody"]],
    stringBody: [
      [/'/, "string", "@popall"],
      [/\\'/, "string.escape"],
      [/./, "string"]
    ],
    dblStrings: [[/"/, "string", "@dblStringBody"]],
    dblStringBody: [
      [/"/, "string", "@popall"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      { include: "@variables" },
      [/./, "string"]
    ],
    quotedConstructs: [
      [/(q|qw|tr|y)\s*\(/, { token: "string.delim", switchTo: "@qstring.(.)" }],
      [/(q|qw|tr|y)\s*\[/, { token: "string.delim", switchTo: "@qstring.[.]" }],
      [/(q|qw|tr|y)\s*\{/, { token: "string.delim", switchTo: "@qstring.{.}" }],
      [/(q|qw|tr|y)\s*</, { token: "string.delim", switchTo: "@qstring.<.>" }],
      [/(q|qw|tr|y)#/, { token: "string.delim", switchTo: "@qstring.#.#" }],
      [/(q|qw|tr|y)\s*([^A-Za-z0-9#\s])/, { token: "string.delim", switchTo: "@qstring.$2.$2" }],
      [/(q|qw|tr|y)\s+(\w)/, { token: "string.delim", switchTo: "@qstring.$2.$2" }],
      [/(qr|m|s)\s*\(/, { token: "regexp.delim", switchTo: "@qregexp.(.)" }],
      [/(qr|m|s)\s*\[/, { token: "regexp.delim", switchTo: "@qregexp.[.]" }],
      [/(qr|m|s)\s*\{/, { token: "regexp.delim", switchTo: "@qregexp.{.}" }],
      [/(qr|m|s)\s*</, { token: "regexp.delim", switchTo: "@qregexp.<.>" }],
      [/(qr|m|s)#/, { token: "regexp.delim", switchTo: "@qregexp.#.#" }],
      [/(qr|m|s)\s*([^A-Za-z0-9_#\s])/, { token: "regexp.delim", switchTo: "@qregexp.$2.$2" }],
      [/(qr|m|s)\s+(\w)/, { token: "regexp.delim", switchTo: "@qregexp.$2.$2" }],
      [/(qq|qx)\s*\(/, { token: "string.delim", switchTo: "@qqstring.(.)" }],
      [/(qq|qx)\s*\[/, { token: "string.delim", switchTo: "@qqstring.[.]" }],
      [/(qq|qx)\s*\{/, { token: "string.delim", switchTo: "@qqstring.{.}" }],
      [/(qq|qx)\s*</, { token: "string.delim", switchTo: "@qqstring.<.>" }],
      [/(qq|qx)#/, { token: "string.delim", switchTo: "@qqstring.#.#" }],
      [/(qq|qx)\s*([^A-Za-z0-9#\s])/, { token: "string.delim", switchTo: "@qqstring.$2.$2" }],
      [/(qq|qx)\s+(\w)/, { token: "string.delim", switchTo: "@qqstring.$2.$2" }]
    ],
    qstring: [
      [/\\./, "string.escape"],
      [
        /./,
        {
          cases: {
            "$#==$S3": { token: "string.delim", next: "@pop" },
            "$#==$S2": { token: "string.delim", next: "@push" },
            "@default": "string"
          }
        }
      ]
    ],
    qregexp: [
      { include: "@variables" },
      [/\\./, "regexp.escape"],
      [
        /./,
        {
          cases: {
            "$#==$S3": {
              token: "regexp.delim",
              next: "@regexpModifiers"
            },
            "$#==$S2": { token: "regexp.delim", next: "@push" },
            "@default": "regexp"
          }
        }
      ]
    ],
    regexpModifiers: [[/[msixpodualngcer]+/, { token: "regexp.modifier", next: "@popall" }]],
    qqstring: [{ include: "@variables" }, { include: "@qstring" }],
    heredoc: [
      [/<<\s*['"`]?([\w\-]+)['"`]?/, { token: "string.heredoc.delimiter", next: "@heredocBody.$1" }]
    ],
    heredocBody: [
      [
        /^([\w\-]+)$/,
        {
          cases: {
            "$1==$S2": [
              {
                token: "string.heredoc.delimiter",
                next: "@popall"
              }
            ],
            "@default": "string.heredoc"
          }
        }
      ],
      [/./, "string.heredoc"]
    ],
    perldoc: [[/^=\w/, "comment.doc", "@perldocBody"]],
    perldocBody: [
      [/^=cut\b/, "type.identifier", "@popall"],
      [/./, "comment.doc"]
    ],
    variables: [
      [/\$\w+/, "variable"],
      [/@\w+/, "variable"],
      [/%\w+/, "variable"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/pgsql/pgsql.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/pgsql/pgsql.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/pgsql/pgsql.ts
var conf = {
  comments: {
    lineComment: "--",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sql",
  ignoreCase: true,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "ALL",
    "ANALYSE",
    "ANALYZE",
    "AND",
    "ANY",
    "ARRAY",
    "AS",
    "ASC",
    "ASYMMETRIC",
    "AUTHORIZATION",
    "BINARY",
    "BOTH",
    "CASE",
    "CAST",
    "CHECK",
    "COLLATE",
    "COLLATION",
    "COLUMN",
    "CONCURRENTLY",
    "CONSTRAINT",
    "CREATE",
    "CROSS",
    "CURRENT_CATALOG",
    "CURRENT_DATE",
    "CURRENT_ROLE",
    "CURRENT_SCHEMA",
    "CURRENT_TIME",
    "CURRENT_TIMESTAMP",
    "CURRENT_USER",
    "DEFAULT",
    "DEFERRABLE",
    "DESC",
    "DISTINCT",
    "DO",
    "ELSE",
    "END",
    "EXCEPT",
    "FALSE",
    "FETCH",
    "FOR",
    "FOREIGN",
    "FREEZE",
    "FROM",
    "FULL",
    "GRANT",
    "GROUP",
    "HAVING",
    "ILIKE",
    "IN",
    "INITIALLY",
    "INNER",
    "INTERSECT",
    "INTO",
    "IS",
    "ISNULL",
    "JOIN",
    "LATERAL",
    "LEADING",
    "LEFT",
    "LIKE",
    "LIMIT",
    "LOCALTIME",
    "LOCALTIMESTAMP",
    "NATURAL",
    "NOT",
    "NOTNULL",
    "NULL",
    "OFFSET",
    "ON",
    "ONLY",
    "OR",
    "ORDER",
    "OUTER",
    "OVERLAPS",
    "PLACING",
    "PRIMARY",
    "REFERENCES",
    "RETURNING",
    "RIGHT",
    "SELECT",
    "SESSION_USER",
    "SIMILAR",
    "SOME",
    "SYMMETRIC",
    "TABLE",
    "TABLESAMPLE",
    "THEN",
    "TO",
    "TRAILING",
    "TRUE",
    "UNION",
    "UNIQUE",
    "USER",
    "USING",
    "VARIADIC",
    "VERBOSE",
    "WHEN",
    "WHERE",
    "WINDOW",
    "WITH"
  ],
  operators: [
    "AND",
    "BETWEEN",
    "IN",
    "LIKE",
    "NOT",
    "OR",
    "IS",
    "NULL",
    "INTERSECT",
    "UNION",
    "INNER",
    "JOIN",
    "LEFT",
    "OUTER",
    "RIGHT"
  ],
  builtinFunctions: [
    "abbrev",
    "abs",
    "acldefault",
    "aclexplode",
    "acos",
    "acosd",
    "acosh",
    "age",
    "any",
    "area",
    "array_agg",
    "array_append",
    "array_cat",
    "array_dims",
    "array_fill",
    "array_length",
    "array_lower",
    "array_ndims",
    "array_position",
    "array_positions",
    "array_prepend",
    "array_remove",
    "array_replace",
    "array_to_json",
    "array_to_string",
    "array_to_tsvector",
    "array_upper",
    "ascii",
    "asin",
    "asind",
    "asinh",
    "atan",
    "atan2",
    "atan2d",
    "atand",
    "atanh",
    "avg",
    "bit",
    "bit_and",
    "bit_count",
    "bit_length",
    "bit_or",
    "bit_xor",
    "bool_and",
    "bool_or",
    "bound_box",
    "box",
    "brin_desummarize_range",
    "brin_summarize_new_values",
    "brin_summarize_range",
    "broadcast",
    "btrim",
    "cardinality",
    "cbrt",
    "ceil",
    "ceiling",
    "center",
    "char_length",
    "character_length",
    "chr",
    "circle",
    "clock_timestamp",
    "coalesce",
    "col_description",
    "concat",
    "concat_ws",
    "convert",
    "convert_from",
    "convert_to",
    "corr",
    "cos",
    "cosd",
    "cosh",
    "cot",
    "cotd",
    "count",
    "covar_pop",
    "covar_samp",
    "cume_dist",
    "current_catalog",
    "current_database",
    "current_date",
    "current_query",
    "current_role",
    "current_schema",
    "current_schemas",
    "current_setting",
    "current_time",
    "current_timestamp",
    "current_user",
    "currval",
    "cursor_to_xml",
    "cursor_to_xmlschema",
    "date_bin",
    "date_part",
    "date_trunc",
    "database_to_xml",
    "database_to_xml_and_xmlschema",
    "database_to_xmlschema",
    "decode",
    "degrees",
    "dense_rank",
    "diagonal",
    "diameter",
    "div",
    "encode",
    "enum_first",
    "enum_last",
    "enum_range",
    "every",
    "exp",
    "extract",
    "factorial",
    "family",
    "first_value",
    "floor",
    "format",
    "format_type",
    "gcd",
    "gen_random_uuid",
    "generate_series",
    "generate_subscripts",
    "get_bit",
    "get_byte",
    "get_current_ts_config",
    "gin_clean_pending_list",
    "greatest",
    "grouping",
    "has_any_column_privilege",
    "has_column_privilege",
    "has_database_privilege",
    "has_foreign_data_wrapper_privilege",
    "has_function_privilege",
    "has_language_privilege",
    "has_schema_privilege",
    "has_sequence_privilege",
    "has_server_privilege",
    "has_table_privilege",
    "has_tablespace_privilege",
    "has_type_privilege",
    "height",
    "host",
    "hostmask",
    "inet_client_addr",
    "inet_client_port",
    "inet_merge",
    "inet_same_family",
    "inet_server_addr",
    "inet_server_port",
    "initcap",
    "isclosed",
    "isempty",
    "isfinite",
    "isopen",
    "json_agg",
    "json_array_elements",
    "json_array_elements_text",
    "json_array_length",
    "json_build_array",
    "json_build_object",
    "json_each",
    "json_each_text",
    "json_extract_path",
    "json_extract_path_text",
    "json_object",
    "json_object_agg",
    "json_object_keys",
    "json_populate_record",
    "json_populate_recordset",
    "json_strip_nulls",
    "json_to_record",
    "json_to_recordset",
    "json_to_tsvector",
    "json_typeof",
    "jsonb_agg",
    "jsonb_array_elements",
    "jsonb_array_elements_text",
    "jsonb_array_length",
    "jsonb_build_array",
    "jsonb_build_object",
    "jsonb_each",
    "jsonb_each_text",
    "jsonb_extract_path",
    "jsonb_extract_path_text",
    "jsonb_insert",
    "jsonb_object",
    "jsonb_object_agg",
    "jsonb_object_keys",
    "jsonb_path_exists",
    "jsonb_path_match",
    "jsonb_path_query",
    "jsonb_path_query_array",
    "jsonb_path_exists_tz",
    "jsonb_path_query_first",
    "jsonb_path_query_array_tz",
    "jsonb_path_query_first_tz",
    "jsonb_path_query_tz",
    "jsonb_path_match_tz",
    "jsonb_populate_record",
    "jsonb_populate_recordset",
    "jsonb_pretty",
    "jsonb_set",
    "jsonb_set_lax",
    "jsonb_strip_nulls",
    "jsonb_to_record",
    "jsonb_to_recordset",
    "jsonb_to_tsvector",
    "jsonb_typeof",
    "justify_days",
    "justify_hours",
    "justify_interval",
    "lag",
    "last_value",
    "lastval",
    "lcm",
    "lead",
    "least",
    "left",
    "length",
    "line",
    "ln",
    "localtime",
    "localtimestamp",
    "log",
    "log10",
    "lower",
    "lower_inc",
    "lower_inf",
    "lpad",
    "lseg",
    "ltrim",
    "macaddr8_set7bit",
    "make_date",
    "make_interval",
    "make_time",
    "make_timestamp",
    "make_timestamptz",
    "makeaclitem",
    "masklen",
    "max",
    "md5",
    "min",
    "min_scale",
    "mod",
    "mode",
    "multirange",
    "netmask",
    "network",
    "nextval",
    "normalize",
    "now",
    "npoints",
    "nth_value",
    "ntile",
    "nullif",
    "num_nonnulls",
    "num_nulls",
    "numnode",
    "obj_description",
    "octet_length",
    "overlay",
    "parse_ident",
    "path",
    "pclose",
    "percent_rank",
    "percentile_cont",
    "percentile_disc",
    "pg_advisory_lock",
    "pg_advisory_lock_shared",
    "pg_advisory_unlock",
    "pg_advisory_unlock_all",
    "pg_advisory_unlock_shared",
    "pg_advisory_xact_lock",
    "pg_advisory_xact_lock_shared",
    "pg_backend_pid",
    "pg_backup_start_time",
    "pg_blocking_pids",
    "pg_cancel_backend",
    "pg_client_encoding",
    "pg_collation_actual_version",
    "pg_collation_is_visible",
    "pg_column_compression",
    "pg_column_size",
    "pg_conf_load_time",
    "pg_control_checkpoint",
    "pg_control_init",
    "pg_control_recovery",
    "pg_control_system",
    "pg_conversion_is_visible",
    "pg_copy_logical_replication_slot",
    "pg_copy_physical_replication_slot",
    "pg_create_logical_replication_slot",
    "pg_create_physical_replication_slot",
    "pg_create_restore_point",
    "pg_current_logfile",
    "pg_current_snapshot",
    "pg_current_wal_flush_lsn",
    "pg_current_wal_insert_lsn",
    "pg_current_wal_lsn",
    "pg_current_xact_id",
    "pg_current_xact_id_if_assigned",
    "pg_current_xlog_flush_location",
    "pg_current_xlog_insert_location",
    "pg_current_xlog_location",
    "pg_database_size",
    "pg_describe_object",
    "pg_drop_replication_slot",
    "pg_event_trigger_ddl_commands",
    "pg_event_trigger_dropped_objects",
    "pg_event_trigger_table_rewrite_oid",
    "pg_event_trigger_table_rewrite_reason",
    "pg_export_snapshot",
    "pg_filenode_relation",
    "pg_function_is_visible",
    "pg_get_catalog_foreign_keys",
    "pg_get_constraintdef",
    "pg_get_expr",
    "pg_get_function_arguments",
    "pg_get_function_identity_arguments",
    "pg_get_function_result",
    "pg_get_functiondef",
    "pg_get_indexdef",
    "pg_get_keywords",
    "pg_get_object_address",
    "pg_get_owned_sequence",
    "pg_get_ruledef",
    "pg_get_serial_sequence",
    "pg_get_statisticsobjdef",
    "pg_get_triggerdef",
    "pg_get_userbyid",
    "pg_get_viewdef",
    "pg_get_wal_replay_pause_state",
    "pg_has_role",
    "pg_identify_object",
    "pg_identify_object_as_address",
    "pg_import_system_collations",
    "pg_index_column_has_property",
    "pg_index_has_property",
    "pg_indexam_has_property",
    "pg_indexes_size",
    "pg_is_in_backup",
    "pg_is_in_recovery",
    "pg_is_other_temp_schema",
    "pg_is_wal_replay_paused",
    "pg_is_xlog_replay_paused",
    "pg_jit_available",
    "pg_last_committed_xact",
    "pg_last_wal_receive_lsn",
    "pg_last_wal_replay_lsn",
    "pg_last_xact_replay_timestamp",
    "pg_last_xlog_receive_location",
    "pg_last_xlog_replay_location",
    "pg_listening_channels",
    "pg_log_backend_memory_contexts",
    "pg_logical_emit_message",
    "pg_logical_slot_get_binary_changes",
    "pg_logical_slot_get_changes",
    "pg_logical_slot_peek_binary_changes",
    "pg_logical_slot_peek_changes",
    "pg_ls_archive_statusdir",
    "pg_ls_dir",
    "pg_ls_logdir",
    "pg_ls_tmpdir",
    "pg_ls_waldir",
    "pg_mcv_list_items",
    "pg_my_temp_schema",
    "pg_notification_queue_usage",
    "pg_opclass_is_visible",
    "pg_operator_is_visible",
    "pg_opfamily_is_visible",
    "pg_options_to_table",
    "pg_partition_ancestors",
    "pg_partition_root",
    "pg_partition_tree",
    "pg_postmaster_start_time",
    "pg_promote",
    "pg_read_binary_file",
    "pg_read_file",
    "pg_relation_filenode",
    "pg_relation_filepath",
    "pg_relation_size",
    "pg_reload_conf",
    "pg_replication_origin_advance",
    "pg_replication_origin_create",
    "pg_replication_origin_drop",
    "pg_replication_origin_oid",
    "pg_replication_origin_progress",
    "pg_replication_origin_session_is_setup",
    "pg_replication_origin_session_progress",
    "pg_replication_origin_session_reset",
    "pg_replication_origin_session_setup",
    "pg_replication_origin_xact_reset",
    "pg_replication_origin_xact_setup",
    "pg_replication_slot_advance",
    "pg_rotate_logfile",
    "pg_safe_snapshot_blocking_pids",
    "pg_size_bytes",
    "pg_size_pretty",
    "pg_sleep",
    "pg_sleep_for",
    "pg_sleep_until",
    "pg_snapshot_xip",
    "pg_snapshot_xmax",
    "pg_snapshot_xmin",
    "pg_start_backup",
    "pg_stat_file",
    "pg_statistics_obj_is_visible",
    "pg_stop_backup",
    "pg_switch_wal",
    "pg_switch_xlog",
    "pg_table_is_visible",
    "pg_table_size",
    "pg_tablespace_databases",
    "pg_tablespace_location",
    "pg_tablespace_size",
    "pg_terminate_backend",
    "pg_total_relation_size",
    "pg_trigger_depth",
    "pg_try_advisory_lock",
    "pg_try_advisory_lock_shared",
    "pg_try_advisory_xact_lock",
    "pg_try_advisory_xact_lock_shared",
    "pg_ts_config_is_visible",
    "pg_ts_dict_is_visible",
    "pg_ts_parser_is_visible",
    "pg_ts_template_is_visible",
    "pg_type_is_visible",
    "pg_typeof",
    "pg_visible_in_snapshot",
    "pg_wal_lsn_diff",
    "pg_wal_replay_pause",
    "pg_wal_replay_resume",
    "pg_walfile_name",
    "pg_walfile_name_offset",
    "pg_xact_commit_timestamp",
    "pg_xact_commit_timestamp_origin",
    "pg_xact_status",
    "pg_xlog_location_diff",
    "pg_xlog_replay_pause",
    "pg_xlog_replay_resume",
    "pg_xlogfile_name",
    "pg_xlogfile_name_offset",
    "phraseto_tsquery",
    "pi",
    "plainto_tsquery",
    "point",
    "polygon",
    "popen",
    "position",
    "power",
    "pqserverversion",
    "query_to_xml",
    "query_to_xml_and_xmlschema",
    "query_to_xmlschema",
    "querytree",
    "quote_ident",
    "quote_literal",
    "quote_nullable",
    "radians",
    "radius",
    "random",
    "range_agg",
    "range_intersect_agg",
    "range_merge",
    "rank",
    "regexp_count",
    "regexp_instr",
    "regexp_like",
    "regexp_match",
    "regexp_matches",
    "regexp_replace",
    "regexp_split_to_array",
    "regexp_split_to_table",
    "regexp_substr",
    "regr_avgx",
    "regr_avgy",
    "regr_count",
    "regr_intercept",
    "regr_r2",
    "regr_slope",
    "regr_sxx",
    "regr_sxy",
    "regr_syy",
    "repeat",
    "replace",
    "reverse",
    "right",
    "round",
    "row_number",
    "row_security_active",
    "row_to_json",
    "rpad",
    "rtrim",
    "scale",
    "schema_to_xml",
    "schema_to_xml_and_xmlschema",
    "schema_to_xmlschema",
    "session_user",
    "set_bit",
    "set_byte",
    "set_config",
    "set_masklen",
    "setseed",
    "setval",
    "setweight",
    "sha224",
    "sha256",
    "sha384",
    "sha512",
    "shobj_description",
    "sign",
    "sin",
    "sind",
    "sinh",
    "slope",
    "split_part",
    "sprintf",
    "sqrt",
    "starts_with",
    "statement_timestamp",
    "stddev",
    "stddev_pop",
    "stddev_samp",
    "string_agg",
    "string_to_array",
    "string_to_table",
    "strip",
    "strpos",
    "substr",
    "substring",
    "sum",
    "suppress_redundant_updates_trigger",
    "table_to_xml",
    "table_to_xml_and_xmlschema",
    "table_to_xmlschema",
    "tan",
    "tand",
    "tanh",
    "text",
    "timeofday",
    "timezone",
    "to_ascii",
    "to_char",
    "to_date",
    "to_hex",
    "to_json",
    "to_number",
    "to_regclass",
    "to_regcollation",
    "to_regnamespace",
    "to_regoper",
    "to_regoperator",
    "to_regproc",
    "to_regprocedure",
    "to_regrole",
    "to_regtype",
    "to_timestamp",
    "to_tsquery",
    "to_tsvector",
    "transaction_timestamp",
    "translate",
    "trim",
    "trim_array",
    "trim_scale",
    "trunc",
    "ts_debug",
    "ts_delete",
    "ts_filter",
    "ts_headline",
    "ts_lexize",
    "ts_parse",
    "ts_rank",
    "ts_rank_cd",
    "ts_rewrite",
    "ts_stat",
    "ts_token_type",
    "tsquery_phrase",
    "tsvector_to_array",
    "tsvector_update_trigger",
    "tsvector_update_trigger_column",
    "txid_current",
    "txid_current_if_assigned",
    "txid_current_snapshot",
    "txid_snapshot_xip",
    "txid_snapshot_xmax",
    "txid_snapshot_xmin",
    "txid_status",
    "txid_visible_in_snapshot",
    "unistr",
    "unnest",
    "upper",
    "upper_inc",
    "upper_inf",
    "user",
    "var_pop",
    "var_samp",
    "variance",
    "version",
    "websearch_to_tsquery",
    "width",
    "width_bucket",
    "xml_is_well_formed",
    "xml_is_well_formed_content",
    "xml_is_well_formed_document",
    "xmlagg",
    "xmlcomment",
    "xmlconcat",
    "xmlelement",
    "xmlexists",
    "xmlforest",
    "xmlparse",
    "xmlpi",
    "xmlroot",
    "xmlserialize",
    "xpath",
    "xpath_exists"
  ],
  builtinVariables: [],
  pseudoColumns: [],
  tokenizer: {
    root: [
      { include: "@comments" },
      { include: "@whitespace" },
      { include: "@pseudoColumns" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@complexIdentifiers" },
      { include: "@scopes" },
      [/[;,.]/, "delimiter"],
      [/[()]/, "@brackets"],
      [
        /[\w@#$]+/,
        {
          cases: {
            "@operators": "operator",
            "@builtinVariables": "predefined",
            "@builtinFunctions": "predefined",
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=!%&+\-*/|~^]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      [/--+.*/, "comment"],
      [/\/\*/, { token: "comment.quote", next: "@comment" }]
    ],
    comment: [
      [/[^*/]+/, "comment"],
      [/\*\//, { token: "comment.quote", next: "@pop" }],
      [/./, "comment"]
    ],
    pseudoColumns: [
      [
        /[$][A-Za-z_][\w@#$]*/,
        {
          cases: {
            "@pseudoColumns": "predefined",
            "@default": "identifier"
          }
        }
      ]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*/, "number"],
      [/[$][+-]*\d*(\.\d*)?/, "number"],
      [/((\d+(\.\d*)?)|(\.\d+))([eE][\-+]?\d+)?/, "number"]
    ],
    strings: [[/'/, { token: "string", next: "@string" }]],
    string: [
      [/[^']+/, "string"],
      [/''/, "string"],
      [/'/, { token: "string", next: "@pop" }]
    ],
    complexIdentifiers: [[/"/, { token: "identifier.quote", next: "@quotedIdentifier" }]],
    quotedIdentifier: [
      [/[^"]+/, "identifier"],
      [/""/, "identifier"],
      [/"/, { token: "identifier.quote", next: "@pop" }]
    ],
    scopes: []
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/php/php.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/php/php.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/php/php.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\#\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string"] },
    { open: "[", close: "]", notIn: ["string"] },
    { open: "(", close: ")", notIn: ["string"] },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*(#|//)region\\b"),
      end: new RegExp("^\\s*(#|//)endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: "",
  tokenizer: {
    root: [
      [/<\?((php)|=)?/, { token: "@rematch", switchTo: "@phpInSimpleState.root" }],
      [/<!DOCTYPE/, "metatag.html", "@doctype"],
      [/<!--/, "comment.html", "@comment"],
      [/(<)(\w+)(\/>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/(<)(script)/, ["delimiter.html", { token: "tag.html", next: "@script" }]],
      [/(<)(style)/, ["delimiter.html", { token: "tag.html", next: "@style" }]],
      [/(<)([:\w]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/(<\/)(\w+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/</, "delimiter.html"],
      [/[^<]+/]
    ],
    doctype: [
      [/<\?((php)|=)?/, { token: "@rematch", switchTo: "@phpInSimpleState.comment" }],
      [/[^>]+/, "metatag.content.html"],
      [/>/, "metatag.html", "@pop"]
    ],
    comment: [
      [/<\?((php)|=)?/, { token: "@rematch", switchTo: "@phpInSimpleState.comment" }],
      [/-->/, "comment.html", "@pop"],
      [/[^-]+/, "comment.content.html"],
      [/./, "comment.content.html"]
    ],
    otherTag: [
      [/<\?((php)|=)?/, { token: "@rematch", switchTo: "@phpInSimpleState.otherTag" }],
      [/\/?>/, "delimiter.html", "@pop"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/]
    ],
    script: [
      [/<\?((php)|=)?/, { token: "@rematch", switchTo: "@phpInSimpleState.script" }],
      [/type/, "attribute.name", "@scriptAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(script\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    scriptAfterType: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInSimpleState.scriptAfterType"
        }
      ],
      [/=/, "delimiter", "@scriptAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptAfterTypeEquals: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInSimpleState.scriptAfterTypeEquals"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptWithCustomType: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInSimpleState.scriptWithCustomType.$S2"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptEmbedded: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInEmbeddedState.scriptEmbedded.$S2",
          nextEmbedded: "@pop"
        }
      ],
      [/<\/script/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }]
    ],
    style: [
      [/<\?((php)|=)?/, { token: "@rematch", switchTo: "@phpInSimpleState.style" }],
      [/type/, "attribute.name", "@styleAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(style\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    styleAfterType: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInSimpleState.styleAfterType"
        }
      ],
      [/=/, "delimiter", "@styleAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleAfterTypeEquals: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInSimpleState.styleAfterTypeEquals"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleWithCustomType: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInSimpleState.styleWithCustomType.$S2"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleEmbedded: [
      [
        /<\?((php)|=)?/,
        {
          token: "@rematch",
          switchTo: "@phpInEmbeddedState.styleEmbedded.$S2",
          nextEmbedded: "@pop"
        }
      ],
      [/<\/style/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }]
    ],
    phpInSimpleState: [
      [/<\?((php)|=)?/, "metatag.php"],
      [/\?>/, { token: "metatag.php", switchTo: "@$S2.$S3" }],
      { include: "phpRoot" }
    ],
    phpInEmbeddedState: [
      [/<\?((php)|=)?/, "metatag.php"],
      [
        /\?>/,
        {
          token: "metatag.php",
          switchTo: "@$S2.$S3",
          nextEmbedded: "$S3"
        }
      ],
      { include: "phpRoot" }
    ],
    phpRoot: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@phpKeywords": { token: "keyword.php" },
            "@phpCompileTimeConstants": { token: "constant.php" },
            "@default": "identifier.php"
          }
        }
      ],
      [
        /[$a-zA-Z_]\w*/,
        {
          cases: {
            "@phpPreDefinedVariables": {
              token: "variable.predefined.php"
            },
            "@default": "variable.php"
          }
        }
      ],
      [/[{}]/, "delimiter.bracket.php"],
      [/[\[\]]/, "delimiter.array.php"],
      [/[()]/, "delimiter.parenthesis.php"],
      [/[ \t\r\n]+/],
      [/(#|\/\/)$/, "comment.php"],
      [/(#|\/\/)/, "comment.php", "@phpLineComment"],
      [/\/\*/, "comment.php", "@phpComment"],
      [/"/, "string.php", "@phpDoubleQuoteString"],
      [/'/, "string.php", "@phpSingleQuoteString"],
      [/[\+\-\*\%\&\|\^\~\!\=\<\>\/\?\;\:\.\,\@]/, "delimiter.php"],
      [/\d*\d+[eE]([\-+]?\d+)?/, "number.float.php"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float.php"],
      [/0[xX][0-9a-fA-F']*[0-9a-fA-F]/, "number.hex.php"],
      [/0[0-7']*[0-7]/, "number.octal.php"],
      [/0[bB][0-1']*[0-1]/, "number.binary.php"],
      [/\d[\d']*/, "number.php"],
      [/\d/, "number.php"]
    ],
    phpComment: [
      [/\*\//, "comment.php", "@pop"],
      [/[^*]+/, "comment.php"],
      [/./, "comment.php"]
    ],
    phpLineComment: [
      [/\?>/, { token: "@rematch", next: "@pop" }],
      [/.$/, "comment.php", "@pop"],
      [/[^?]+$/, "comment.php", "@pop"],
      [/[^?]+/, "comment.php"],
      [/./, "comment.php"]
    ],
    phpDoubleQuoteString: [
      [/[^\\"]+/, "string.php"],
      [/@escapes/, "string.escape.php"],
      [/\\./, "string.escape.invalid.php"],
      [/"/, "string.php", "@pop"]
    ],
    phpSingleQuoteString: [
      [/[^\\']+/, "string.php"],
      [/@escapes/, "string.escape.php"],
      [/\\./, "string.escape.invalid.php"],
      [/'/, "string.php", "@pop"]
    ]
  },
  phpKeywords: [
    "abstract",
    "and",
    "array",
    "as",
    "break",
    "callable",
    "case",
    "catch",
    "cfunction",
    "class",
    "clone",
    "const",
    "continue",
    "declare",
    "default",
    "do",
    "else",
    "elseif",
    "enddeclare",
    "endfor",
    "endforeach",
    "endif",
    "endswitch",
    "endwhile",
    "extends",
    "false",
    "final",
    "for",
    "foreach",
    "function",
    "global",
    "goto",
    "if",
    "implements",
    "interface",
    "instanceof",
    "insteadof",
    "namespace",
    "new",
    "null",
    "object",
    "old_function",
    "or",
    "private",
    "protected",
    "public",
    "resource",
    "static",
    "switch",
    "throw",
    "trait",
    "try",
    "true",
    "use",
    "var",
    "while",
    "xor",
    "die",
    "echo",
    "empty",
    "exit",
    "eval",
    "include",
    "include_once",
    "isset",
    "list",
    "require",
    "require_once",
    "return",
    "print",
    "unset",
    "yield",
    "__construct"
  ],
  phpCompileTimeConstants: [
    "__CLASS__",
    "__DIR__",
    "__FILE__",
    "__LINE__",
    "__NAMESPACE__",
    "__METHOD__",
    "__FUNCTION__",
    "__TRAIT__"
  ],
  phpPreDefinedVariables: [
    "$GLOBALS",
    "$_SERVER",
    "$_GET",
    "$_POST",
    "$_FILES",
    "$_REQUEST",
    "$_SESSION",
    "$_ENV",
    "$_COOKIE",
    "$php_errormsg",
    "$HTTP_RAW_POST_DATA",
    "$http_response_header",
    "$argc",
    "$argv"
  ],
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/pla/pla.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/pla/pla.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/pla/pla.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["[", "]"],
    ["<", ">"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "<", close: ">" },
    { open: "(", close: ")" }
  ],
  surroundingPairs: [
    { open: "[", close: "]" },
    { open: "<", close: ">" },
    { open: "(", close: ")" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".pla",
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "<", close: ">", token: "delimiter.angle" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    ".i",
    ".o",
    ".mv",
    ".ilb",
    ".ob",
    ".label",
    ".type",
    ".phase",
    ".pair",
    ".symbolic",
    ".symbolic-output",
    ".kiss",
    ".p",
    ".e",
    ".end"
  ],
  comment: /#.*$/,
  identifier: /[a-zA-Z]+[a-zA-Z0-9_\-]*/,
  plaContent: /[01\-~\|]+/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      [/@comment/, "comment"],
      [
        /\.([a-zA-Z_\-]+)/,
        {
          cases: {
            "@eos": { token: "keyword.$1" },
            "@keywords": {
              cases: {
                ".type": { token: "keyword.$1", next: "@type" },
                "@default": { token: "keyword.$1", next: "@keywordArg" }
              }
            },
            "@default": { token: "keyword.$1" }
          }
        }
      ],
      [/@identifier/, "identifier"],
      [/@plaContent/, "string"]
    ],
    whitespace: [[/[ \t\r\n]+/, ""]],
    type: [{ include: "@whitespace" }, [/\w+/, { token: "type", next: "@pop" }]],
    keywordArg: [
      [
        /[ \t\r\n]+/,
        {
          cases: {
            "@eos": { token: "", next: "@pop" },
            "@default": ""
          }
        }
      ],
      [/@comment/, "comment", "@pop"],
      [
        /[<>()\[\]]/,
        {
          cases: {
            "@eos": { token: "@brackets", next: "@pop" },
            "@default": "@brackets"
          }
        }
      ],
      [
        /\-?\d+/,
        {
          cases: {
            "@eos": { token: "number", next: "@pop" },
            "@default": "number"
          }
        }
      ],
      [
        /@identifier/,
        {
          cases: {
            "@eos": { token: "identifier", next: "@pop" },
            "@default": "identifier"
          }
        }
      ],
      [
        /[;=]/,
        {
          cases: {
            "@eos": { token: "delimiter", next: "@pop" },
            "@default": "delimiter"
          }
        }
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/postiats/postiats.js":
/*!********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/postiats/postiats.js ***!
  \********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/postiats/postiats.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["(*", "*)"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] }
  ]
};
var language = {
  tokenPostfix: ".pats",
  defaultToken: "invalid",
  keywords: [
    "abstype",
    "abst0ype",
    "absprop",
    "absview",
    "absvtype",
    "absviewtype",
    "absvt0ype",
    "absviewt0ype",
    "as",
    "and",
    "assume",
    "begin",
    "classdec",
    "datasort",
    "datatype",
    "dataprop",
    "dataview",
    "datavtype",
    "dataviewtype",
    "do",
    "end",
    "extern",
    "extype",
    "extvar",
    "exception",
    "fn",
    "fnx",
    "fun",
    "prfn",
    "prfun",
    "praxi",
    "castfn",
    "if",
    "then",
    "else",
    "ifcase",
    "in",
    "infix",
    "infixl",
    "infixr",
    "prefix",
    "postfix",
    "implmnt",
    "implement",
    "primplmnt",
    "primplement",
    "import",
    "let",
    "local",
    "macdef",
    "macrodef",
    "nonfix",
    "symelim",
    "symintr",
    "overload",
    "of",
    "op",
    "rec",
    "sif",
    "scase",
    "sortdef",
    "sta",
    "stacst",
    "stadef",
    "static",
    "staload",
    "dynload",
    "try",
    "tkindef",
    "typedef",
    "propdef",
    "viewdef",
    "vtypedef",
    "viewtypedef",
    "prval",
    "var",
    "prvar",
    "when",
    "where",
    "with",
    "withtype",
    "withprop",
    "withview",
    "withvtype",
    "withviewtype"
  ],
  keywords_dlr: [
    "$delay",
    "$ldelay",
    "$arrpsz",
    "$arrptrsize",
    "$d2ctype",
    "$effmask",
    "$effmask_ntm",
    "$effmask_exn",
    "$effmask_ref",
    "$effmask_wrt",
    "$effmask_all",
    "$extern",
    "$extkind",
    "$extype",
    "$extype_struct",
    "$extval",
    "$extfcall",
    "$extmcall",
    "$literal",
    "$myfilename",
    "$mylocation",
    "$myfunction",
    "$lst",
    "$lst_t",
    "$lst_vt",
    "$list",
    "$list_t",
    "$list_vt",
    "$rec",
    "$rec_t",
    "$rec_vt",
    "$record",
    "$record_t",
    "$record_vt",
    "$tup",
    "$tup_t",
    "$tup_vt",
    "$tuple",
    "$tuple_t",
    "$tuple_vt",
    "$break",
    "$continue",
    "$raise",
    "$showtype",
    "$vcopyenv_v",
    "$vcopyenv_vt",
    "$tempenver",
    "$solver_assert",
    "$solver_verify"
  ],
  keywords_srp: [
    "#if",
    "#ifdef",
    "#ifndef",
    "#then",
    "#elif",
    "#elifdef",
    "#elifndef",
    "#else",
    "#endif",
    "#error",
    "#prerr",
    "#print",
    "#assert",
    "#undef",
    "#define",
    "#include",
    "#require",
    "#pragma",
    "#codegen2",
    "#codegen3"
  ],
  irregular_keyword_list: [
    "val+",
    "val-",
    "val",
    "case+",
    "case-",
    "case",
    "addr@",
    "addr",
    "fold@",
    "free@",
    "fix@",
    "fix",
    "lam@",
    "lam",
    "llam@",
    "llam",
    "viewt@ype+",
    "viewt@ype-",
    "viewt@ype",
    "viewtype+",
    "viewtype-",
    "viewtype",
    "view+",
    "view-",
    "view@",
    "view",
    "type+",
    "type-",
    "type",
    "vtype+",
    "vtype-",
    "vtype",
    "vt@ype+",
    "vt@ype-",
    "vt@ype",
    "viewt@ype+",
    "viewt@ype-",
    "viewt@ype",
    "viewtype+",
    "viewtype-",
    "viewtype",
    "prop+",
    "prop-",
    "prop",
    "type+",
    "type-",
    "type",
    "t@ype",
    "t@ype+",
    "t@ype-",
    "abst@ype",
    "abstype",
    "absviewt@ype",
    "absvt@ype",
    "for*",
    "for",
    "while*",
    "while"
  ],
  keywords_types: [
    "bool",
    "double",
    "byte",
    "int",
    "short",
    "char",
    "void",
    "unit",
    "long",
    "float",
    "string",
    "strptr"
  ],
  keywords_effects: [
    "0",
    "fun",
    "clo",
    "prf",
    "funclo",
    "cloptr",
    "cloref",
    "ref",
    "ntm",
    "1"
  ],
  operators: [
    "@",
    "!",
    "|",
    "`",
    ":",
    "$",
    ".",
    "=",
    "#",
    "~",
    "..",
    "...",
    "=>",
    "=<>",
    "=/=>",
    "=>>",
    "=/=>>",
    "<",
    ">",
    "><",
    ".<",
    ">.",
    ".<>.",
    "->",
    "-<>"
  ],
  brackets: [
    { open: ",(", close: ")", token: "delimiter.parenthesis" },
    { open: "`(", close: ")", token: "delimiter.parenthesis" },
    { open: "%(", close: ")", token: "delimiter.parenthesis" },
    { open: "'(", close: ")", token: "delimiter.parenthesis" },
    { open: "'{", close: "}", token: "delimiter.parenthesis" },
    { open: "@(", close: ")", token: "delimiter.parenthesis" },
    { open: "@{", close: "}", token: "delimiter.brace" },
    { open: "@[", close: "]", token: "delimiter.square" },
    { open: "#[", close: "]", token: "delimiter.square" },
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  IDENTFST: /[a-zA-Z_]/,
  IDENTRST: /[a-zA-Z0-9_'$]/,
  symbolic: /[%&+-./:=@~`^|*!$#?<>]/,
  digit: /[0-9]/,
  digitseq0: /@digit*/,
  xdigit: /[0-9A-Za-z]/,
  xdigitseq0: /@xdigit*/,
  INTSP: /[lLuU]/,
  FLOATSP: /[fFlL]/,
  fexponent: /[eE][+-]?[0-9]+/,
  fexponent_bin: /[pP][+-]?[0-9]+/,
  deciexp: /\.[0-9]*@fexponent?/,
  hexiexp: /\.[0-9a-zA-Z]*@fexponent_bin?/,
  irregular_keywords: /val[+-]?|case[+-]?|addr\@?|fold\@|free\@|fix\@?|lam\@?|llam\@?|prop[+-]?|type[+-]?|view[+-@]?|viewt@?ype[+-]?|t@?ype[+-]?|v(iew)?t@?ype[+-]?|abst@?ype|absv(iew)?t@?ype|for\*?|while\*?/,
  ESCHAR: /[ntvbrfa\\\?'"\(\[\{]/,
  start: "root",
  tokenizer: {
    root: [
      { regex: /[ \t\r\n]+/, action: { token: "" } },
      { regex: /\(\*\)/, action: { token: "invalid" } },
      {
        regex: /\(\*/,
        action: { token: "comment", next: "lexing_COMMENT_block_ml" }
      },
      {
        regex: /\(/,
        action: "@brackets"
      },
      {
        regex: /\)/,
        action: "@brackets"
      },
      {
        regex: /\[/,
        action: "@brackets"
      },
      {
        regex: /\]/,
        action: "@brackets"
      },
      {
        regex: /\{/,
        action: "@brackets"
      },
      {
        regex: /\}/,
        action: "@brackets"
      },
      {
        regex: /,\(/,
        action: "@brackets"
      },
      { regex: /,/, action: { token: "delimiter.comma" } },
      { regex: /;/, action: { token: "delimiter.semicolon" } },
      {
        regex: /@\(/,
        action: "@brackets"
      },
      {
        regex: /@\[/,
        action: "@brackets"
      },
      {
        regex: /@\{/,
        action: "@brackets"
      },
      {
        regex: /:</,
        action: { token: "keyword", next: "@lexing_EFFECT_commaseq0" }
      },
      { regex: /\.@symbolic+/, action: { token: "identifier.sym" } },
      {
        regex: /\.@digit*@fexponent@FLOATSP*/,
        action: { token: "number.float" }
      },
      { regex: /\.@digit+/, action: { token: "number.float" } },
      {
        regex: /\$@IDENTFST@IDENTRST*/,
        action: {
          cases: {
            "@keywords_dlr": { token: "keyword.dlr" },
            "@default": { token: "namespace" }
          }
        }
      },
      {
        regex: /\#@IDENTFST@IDENTRST*/,
        action: {
          cases: {
            "@keywords_srp": { token: "keyword.srp" },
            "@default": { token: "identifier" }
          }
        }
      },
      { regex: /%\(/, action: { token: "delimiter.parenthesis" } },
      {
        regex: /^%{(#|\^|\$)?/,
        action: {
          token: "keyword",
          next: "@lexing_EXTCODE",
          nextEmbedded: "text/javascript"
        }
      },
      { regex: /^%}/, action: { token: "keyword" } },
      { regex: /'\(/, action: { token: "delimiter.parenthesis" } },
      { regex: /'\[/, action: { token: "delimiter.bracket" } },
      { regex: /'\{/, action: { token: "delimiter.brace" } },
      [/(')(\\@ESCHAR|\\[xX]@xdigit+|\\@digit+)(')/, ["string", "string.escape", "string"]],
      [/'[^\\']'/, "string"],
      [/"/, "string.quote", "@lexing_DQUOTE"],
      {
        regex: /`\(/,
        action: "@brackets"
      },
      { regex: /\\/, action: { token: "punctuation" } },
      {
        regex: /@irregular_keywords(?!@IDENTRST)/,
        action: { token: "keyword" }
      },
      {
        regex: /@IDENTFST@IDENTRST*[<!\[]?/,
        action: {
          cases: {
            "@keywords": { token: "keyword" },
            "@keywords_types": { token: "type" },
            "@default": { token: "identifier" }
          }
        }
      },
      {
        regex: /\/\/\/\//,
        action: { token: "comment", next: "@lexing_COMMENT_rest" }
      },
      { regex: /\/\/.*$/, action: { token: "comment" } },
      {
        regex: /\/\*/,
        action: { token: "comment", next: "@lexing_COMMENT_block_c" }
      },
      {
        regex: /-<|=</,
        action: { token: "keyword", next: "@lexing_EFFECT_commaseq0" }
      },
      {
        regex: /@symbolic+/,
        action: {
          cases: {
            "@operators": "keyword",
            "@default": "operator"
          }
        }
      },
      {
        regex: /0[xX]@xdigit+(@hexiexp|@fexponent_bin)@FLOATSP*/,
        action: { token: "number.float" }
      },
      { regex: /0[xX]@xdigit+@INTSP*/, action: { token: "number.hex" } },
      {
        regex: /0[0-7]+(?![0-9])@INTSP*/,
        action: { token: "number.octal" }
      },
      {
        regex: /@digit+(@fexponent|@deciexp)@FLOATSP*/,
        action: { token: "number.float" }
      },
      {
        regex: /@digit@digitseq0@INTSP*/,
        action: { token: "number.decimal" }
      },
      { regex: /@digit+@INTSP*/, action: { token: "number" } }
    ],
    lexing_COMMENT_block_ml: [
      [/[^\(\*]+/, "comment"],
      [/\(\*/, "comment", "@push"],
      [/\(\*/, "comment.invalid"],
      [/\*\)/, "comment", "@pop"],
      [/\*/, "comment"]
    ],
    lexing_COMMENT_block_c: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    lexing_COMMENT_rest: [
      [/$/, "comment", "@pop"],
      [/.*/, "comment"]
    ],
    lexing_EFFECT_commaseq0: [
      {
        regex: /@IDENTFST@IDENTRST+|@digit+/,
        action: {
          cases: {
            "@keywords_effects": { token: "type.effect" },
            "@default": { token: "identifier" }
          }
        }
      },
      { regex: /,/, action: { token: "punctuation" } },
      { regex: />/, action: { token: "@rematch", next: "@pop" } }
    ],
    lexing_EXTCODE: [
      {
        regex: /^%}/,
        action: {
          token: "@rematch",
          next: "@pop",
          nextEmbedded: "@pop"
        }
      },
      { regex: /[^%]+/, action: "" }
    ],
    lexing_DQUOTE: [
      { regex: /"/, action: { token: "string.quote", next: "@pop" } },
      {
        regex: /(\{\$)(@IDENTFST@IDENTRST*)(\})/,
        action: [{ token: "string.escape" }, { token: "identifier" }, { token: "string.escape" }]
      },
      { regex: /\\$/, action: { token: "string.escape" } },
      {
        regex: /\\(@ESCHAR|[xX]@xdigit+|@digit+)/,
        action: { token: "string.escape" }
      },
      { regex: /[^\\"]+/, action: { token: "string" } }
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/powerquery/powerquery.js":
/*!************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/powerquery/powerquery.js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/powerquery/powerquery.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["[", "]"],
    ["(", ")"],
    ["{", "}"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment", "identifier"] },
    { open: "[", close: "]", notIn: ["string", "comment", "identifier"] },
    { open: "(", close: ")", notIn: ["string", "comment", "identifier"] },
    { open: "{", close: "}", notIn: ["string", "comment", "identifier"] }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".pq",
  ignoreCase: false,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "{", close: "}", token: "delimiter.brackets" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  operatorKeywords: ["and", "not", "or"],
  keywords: [
    "as",
    "each",
    "else",
    "error",
    "false",
    "if",
    "in",
    "is",
    "let",
    "meta",
    "otherwise",
    "section",
    "shared",
    "then",
    "true",
    "try",
    "type"
  ],
  constructors: ["#binary", "#date", "#datetime", "#datetimezone", "#duration", "#table", "#time"],
  constants: ["#infinity", "#nan", "#sections", "#shared"],
  typeKeywords: [
    "action",
    "any",
    "anynonnull",
    "none",
    "null",
    "logical",
    "number",
    "time",
    "date",
    "datetime",
    "datetimezone",
    "duration",
    "text",
    "binary",
    "list",
    "record",
    "table",
    "function"
  ],
  builtinFunctions: [
    "Access.Database",
    "Action.Return",
    "Action.Sequence",
    "Action.Try",
    "ActiveDirectory.Domains",
    "AdoDotNet.DataSource",
    "AdoDotNet.Query",
    "AdobeAnalytics.Cubes",
    "AnalysisServices.Database",
    "AnalysisServices.Databases",
    "AzureStorage.BlobContents",
    "AzureStorage.Blobs",
    "AzureStorage.Tables",
    "Binary.Buffer",
    "Binary.Combine",
    "Binary.Compress",
    "Binary.Decompress",
    "Binary.End",
    "Binary.From",
    "Binary.FromList",
    "Binary.FromText",
    "Binary.InferContentType",
    "Binary.Length",
    "Binary.ToList",
    "Binary.ToText",
    "BinaryFormat.7BitEncodedSignedInteger",
    "BinaryFormat.7BitEncodedUnsignedInteger",
    "BinaryFormat.Binary",
    "BinaryFormat.Byte",
    "BinaryFormat.ByteOrder",
    "BinaryFormat.Choice",
    "BinaryFormat.Decimal",
    "BinaryFormat.Double",
    "BinaryFormat.Group",
    "BinaryFormat.Length",
    "BinaryFormat.List",
    "BinaryFormat.Null",
    "BinaryFormat.Record",
    "BinaryFormat.SignedInteger16",
    "BinaryFormat.SignedInteger32",
    "BinaryFormat.SignedInteger64",
    "BinaryFormat.Single",
    "BinaryFormat.Text",
    "BinaryFormat.Transform",
    "BinaryFormat.UnsignedInteger16",
    "BinaryFormat.UnsignedInteger32",
    "BinaryFormat.UnsignedInteger64",
    "Byte.From",
    "Character.FromNumber",
    "Character.ToNumber",
    "Combiner.CombineTextByDelimiter",
    "Combiner.CombineTextByEachDelimiter",
    "Combiner.CombineTextByLengths",
    "Combiner.CombineTextByPositions",
    "Combiner.CombineTextByRanges",
    "Comparer.Equals",
    "Comparer.FromCulture",
    "Comparer.Ordinal",
    "Comparer.OrdinalIgnoreCase",
    "Csv.Document",
    "Cube.AddAndExpandDimensionColumn",
    "Cube.AddMeasureColumn",
    "Cube.ApplyParameter",
    "Cube.AttributeMemberId",
    "Cube.AttributeMemberProperty",
    "Cube.CollapseAndRemoveColumns",
    "Cube.Dimensions",
    "Cube.DisplayFolders",
    "Cube.Measures",
    "Cube.Parameters",
    "Cube.Properties",
    "Cube.PropertyKey",
    "Cube.ReplaceDimensions",
    "Cube.Transform",
    "Currency.From",
    "DB2.Database",
    "Date.AddDays",
    "Date.AddMonths",
    "Date.AddQuarters",
    "Date.AddWeeks",
    "Date.AddYears",
    "Date.Day",
    "Date.DayOfWeek",
    "Date.DayOfWeekName",
    "Date.DayOfYear",
    "Date.DaysInMonth",
    "Date.EndOfDay",
    "Date.EndOfMonth",
    "Date.EndOfQuarter",
    "Date.EndOfWeek",
    "Date.EndOfYear",
    "Date.From",
    "Date.FromText",
    "Date.IsInCurrentDay",
    "Date.IsInCurrentMonth",
    "Date.IsInCurrentQuarter",
    "Date.IsInCurrentWeek",
    "Date.IsInCurrentYear",
    "Date.IsInNextDay",
    "Date.IsInNextMonth",
    "Date.IsInNextNDays",
    "Date.IsInNextNMonths",
    "Date.IsInNextNQuarters",
    "Date.IsInNextNWeeks",
    "Date.IsInNextNYears",
    "Date.IsInNextQuarter",
    "Date.IsInNextWeek",
    "Date.IsInNextYear",
    "Date.IsInPreviousDay",
    "Date.IsInPreviousMonth",
    "Date.IsInPreviousNDays",
    "Date.IsInPreviousNMonths",
    "Date.IsInPreviousNQuarters",
    "Date.IsInPreviousNWeeks",
    "Date.IsInPreviousNYears",
    "Date.IsInPreviousQuarter",
    "Date.IsInPreviousWeek",
    "Date.IsInPreviousYear",
    "Date.IsInYearToDate",
    "Date.IsLeapYear",
    "Date.Month",
    "Date.MonthName",
    "Date.QuarterOfYear",
    "Date.StartOfDay",
    "Date.StartOfMonth",
    "Date.StartOfQuarter",
    "Date.StartOfWeek",
    "Date.StartOfYear",
    "Date.ToRecord",
    "Date.ToText",
    "Date.WeekOfMonth",
    "Date.WeekOfYear",
    "Date.Year",
    "DateTime.AddZone",
    "DateTime.Date",
    "DateTime.FixedLocalNow",
    "DateTime.From",
    "DateTime.FromFileTime",
    "DateTime.FromText",
    "DateTime.IsInCurrentHour",
    "DateTime.IsInCurrentMinute",
    "DateTime.IsInCurrentSecond",
    "DateTime.IsInNextHour",
    "DateTime.IsInNextMinute",
    "DateTime.IsInNextNHours",
    "DateTime.IsInNextNMinutes",
    "DateTime.IsInNextNSeconds",
    "DateTime.IsInNextSecond",
    "DateTime.IsInPreviousHour",
    "DateTime.IsInPreviousMinute",
    "DateTime.IsInPreviousNHours",
    "DateTime.IsInPreviousNMinutes",
    "DateTime.IsInPreviousNSeconds",
    "DateTime.IsInPreviousSecond",
    "DateTime.LocalNow",
    "DateTime.Time",
    "DateTime.ToRecord",
    "DateTime.ToText",
    "DateTimeZone.FixedLocalNow",
    "DateTimeZone.FixedUtcNow",
    "DateTimeZone.From",
    "DateTimeZone.FromFileTime",
    "DateTimeZone.FromText",
    "DateTimeZone.LocalNow",
    "DateTimeZone.RemoveZone",
    "DateTimeZone.SwitchZone",
    "DateTimeZone.ToLocal",
    "DateTimeZone.ToRecord",
    "DateTimeZone.ToText",
    "DateTimeZone.ToUtc",
    "DateTimeZone.UtcNow",
    "DateTimeZone.ZoneHours",
    "DateTimeZone.ZoneMinutes",
    "Decimal.From",
    "Diagnostics.ActivityId",
    "Diagnostics.Trace",
    "DirectQueryCapabilities.From",
    "Double.From",
    "Duration.Days",
    "Duration.From",
    "Duration.FromText",
    "Duration.Hours",
    "Duration.Minutes",
    "Duration.Seconds",
    "Duration.ToRecord",
    "Duration.ToText",
    "Duration.TotalDays",
    "Duration.TotalHours",
    "Duration.TotalMinutes",
    "Duration.TotalSeconds",
    "Embedded.Value",
    "Error.Record",
    "Excel.CurrentWorkbook",
    "Excel.Workbook",
    "Exchange.Contents",
    "Expression.Constant",
    "Expression.Evaluate",
    "Expression.Identifier",
    "Facebook.Graph",
    "File.Contents",
    "Folder.Contents",
    "Folder.Files",
    "Function.From",
    "Function.Invoke",
    "Function.InvokeAfter",
    "Function.IsDataSource",
    "GoogleAnalytics.Accounts",
    "Guid.From",
    "HdInsight.Containers",
    "HdInsight.Contents",
    "HdInsight.Files",
    "Hdfs.Contents",
    "Hdfs.Files",
    "Informix.Database",
    "Int16.From",
    "Int32.From",
    "Int64.From",
    "Int8.From",
    "ItemExpression.From",
    "Json.Document",
    "Json.FromValue",
    "Lines.FromBinary",
    "Lines.FromText",
    "Lines.ToBinary",
    "Lines.ToText",
    "List.Accumulate",
    "List.AllTrue",
    "List.Alternate",
    "List.AnyTrue",
    "List.Average",
    "List.Buffer",
    "List.Combine",
    "List.Contains",
    "List.ContainsAll",
    "List.ContainsAny",
    "List.Count",
    "List.Covariance",
    "List.DateTimeZones",
    "List.DateTimes",
    "List.Dates",
    "List.Difference",
    "List.Distinct",
    "List.Durations",
    "List.FindText",
    "List.First",
    "List.FirstN",
    "List.Generate",
    "List.InsertRange",
    "List.Intersect",
    "List.IsDistinct",
    "List.IsEmpty",
    "List.Last",
    "List.LastN",
    "List.MatchesAll",
    "List.MatchesAny",
    "List.Max",
    "List.MaxN",
    "List.Median",
    "List.Min",
    "List.MinN",
    "List.Mode",
    "List.Modes",
    "List.NonNullCount",
    "List.Numbers",
    "List.PositionOf",
    "List.PositionOfAny",
    "List.Positions",
    "List.Product",
    "List.Random",
    "List.Range",
    "List.RemoveFirstN",
    "List.RemoveItems",
    "List.RemoveLastN",
    "List.RemoveMatchingItems",
    "List.RemoveNulls",
    "List.RemoveRange",
    "List.Repeat",
    "List.ReplaceMatchingItems",
    "List.ReplaceRange",
    "List.ReplaceValue",
    "List.Reverse",
    "List.Select",
    "List.Single",
    "List.SingleOrDefault",
    "List.Skip",
    "List.Sort",
    "List.StandardDeviation",
    "List.Sum",
    "List.Times",
    "List.Transform",
    "List.TransformMany",
    "List.Union",
    "List.Zip",
    "Logical.From",
    "Logical.FromText",
    "Logical.ToText",
    "MQ.Queue",
    "MySQL.Database",
    "Number.Abs",
    "Number.Acos",
    "Number.Asin",
    "Number.Atan",
    "Number.Atan2",
    "Number.BitwiseAnd",
    "Number.BitwiseNot",
    "Number.BitwiseOr",
    "Number.BitwiseShiftLeft",
    "Number.BitwiseShiftRight",
    "Number.BitwiseXor",
    "Number.Combinations",
    "Number.Cos",
    "Number.Cosh",
    "Number.Exp",
    "Number.Factorial",
    "Number.From",
    "Number.FromText",
    "Number.IntegerDivide",
    "Number.IsEven",
    "Number.IsNaN",
    "Number.IsOdd",
    "Number.Ln",
    "Number.Log",
    "Number.Log10",
    "Number.Mod",
    "Number.Permutations",
    "Number.Power",
    "Number.Random",
    "Number.RandomBetween",
    "Number.Round",
    "Number.RoundAwayFromZero",
    "Number.RoundDown",
    "Number.RoundTowardZero",
    "Number.RoundUp",
    "Number.Sign",
    "Number.Sin",
    "Number.Sinh",
    "Number.Sqrt",
    "Number.Tan",
    "Number.Tanh",
    "Number.ToText",
    "OData.Feed",
    "Odbc.DataSource",
    "Odbc.Query",
    "OleDb.DataSource",
    "OleDb.Query",
    "Oracle.Database",
    "Percentage.From",
    "PostgreSQL.Database",
    "RData.FromBinary",
    "Record.AddField",
    "Record.Combine",
    "Record.Field",
    "Record.FieldCount",
    "Record.FieldNames",
    "Record.FieldOrDefault",
    "Record.FieldValues",
    "Record.FromList",
    "Record.FromTable",
    "Record.HasFields",
    "Record.RemoveFields",
    "Record.RenameFields",
    "Record.ReorderFields",
    "Record.SelectFields",
    "Record.ToList",
    "Record.ToTable",
    "Record.TransformFields",
    "Replacer.ReplaceText",
    "Replacer.ReplaceValue",
    "RowExpression.Column",
    "RowExpression.From",
    "Salesforce.Data",
    "Salesforce.Reports",
    "SapBusinessWarehouse.Cubes",
    "SapHana.Database",
    "SharePoint.Contents",
    "SharePoint.Files",
    "SharePoint.Tables",
    "Single.From",
    "Soda.Feed",
    "Splitter.SplitByNothing",
    "Splitter.SplitTextByAnyDelimiter",
    "Splitter.SplitTextByDelimiter",
    "Splitter.SplitTextByEachDelimiter",
    "Splitter.SplitTextByLengths",
    "Splitter.SplitTextByPositions",
    "Splitter.SplitTextByRanges",
    "Splitter.SplitTextByRepeatedLengths",
    "Splitter.SplitTextByWhitespace",
    "Sql.Database",
    "Sql.Databases",
    "SqlExpression.SchemaFrom",
    "SqlExpression.ToExpression",
    "Sybase.Database",
    "Table.AddColumn",
    "Table.AddIndexColumn",
    "Table.AddJoinColumn",
    "Table.AddKey",
    "Table.AggregateTableColumn",
    "Table.AlternateRows",
    "Table.Buffer",
    "Table.Column",
    "Table.ColumnCount",
    "Table.ColumnNames",
    "Table.ColumnsOfType",
    "Table.Combine",
    "Table.CombineColumns",
    "Table.Contains",
    "Table.ContainsAll",
    "Table.ContainsAny",
    "Table.DemoteHeaders",
    "Table.Distinct",
    "Table.DuplicateColumn",
    "Table.ExpandListColumn",
    "Table.ExpandRecordColumn",
    "Table.ExpandTableColumn",
    "Table.FillDown",
    "Table.FillUp",
    "Table.FilterWithDataTable",
    "Table.FindText",
    "Table.First",
    "Table.FirstN",
    "Table.FirstValue",
    "Table.FromColumns",
    "Table.FromList",
    "Table.FromPartitions",
    "Table.FromRecords",
    "Table.FromRows",
    "Table.FromValue",
    "Table.Group",
    "Table.HasColumns",
    "Table.InsertRows",
    "Table.IsDistinct",
    "Table.IsEmpty",
    "Table.Join",
    "Table.Keys",
    "Table.Last",
    "Table.LastN",
    "Table.MatchesAllRows",
    "Table.MatchesAnyRows",
    "Table.Max",
    "Table.MaxN",
    "Table.Min",
    "Table.MinN",
    "Table.NestedJoin",
    "Table.Partition",
    "Table.PartitionValues",
    "Table.Pivot",
    "Table.PositionOf",
    "Table.PositionOfAny",
    "Table.PrefixColumns",
    "Table.Profile",
    "Table.PromoteHeaders",
    "Table.Range",
    "Table.RemoveColumns",
    "Table.RemoveFirstN",
    "Table.RemoveLastN",
    "Table.RemoveMatchingRows",
    "Table.RemoveRows",
    "Table.RemoveRowsWithErrors",
    "Table.RenameColumns",
    "Table.ReorderColumns",
    "Table.Repeat",
    "Table.ReplaceErrorValues",
    "Table.ReplaceKeys",
    "Table.ReplaceMatchingRows",
    "Table.ReplaceRelationshipIdentity",
    "Table.ReplaceRows",
    "Table.ReplaceValue",
    "Table.ReverseRows",
    "Table.RowCount",
    "Table.Schema",
    "Table.SelectColumns",
    "Table.SelectRows",
    "Table.SelectRowsWithErrors",
    "Table.SingleRow",
    "Table.Skip",
    "Table.Sort",
    "Table.SplitColumn",
    "Table.ToColumns",
    "Table.ToList",
    "Table.ToRecords",
    "Table.ToRows",
    "Table.TransformColumnNames",
    "Table.TransformColumnTypes",
    "Table.TransformColumns",
    "Table.TransformRows",
    "Table.Transpose",
    "Table.Unpivot",
    "Table.UnpivotOtherColumns",
    "Table.View",
    "Table.ViewFunction",
    "TableAction.DeleteRows",
    "TableAction.InsertRows",
    "TableAction.UpdateRows",
    "Tables.GetRelationships",
    "Teradata.Database",
    "Text.AfterDelimiter",
    "Text.At",
    "Text.BeforeDelimiter",
    "Text.BetweenDelimiters",
    "Text.Clean",
    "Text.Combine",
    "Text.Contains",
    "Text.End",
    "Text.EndsWith",
    "Text.Format",
    "Text.From",
    "Text.FromBinary",
    "Text.Insert",
    "Text.Length",
    "Text.Lower",
    "Text.Middle",
    "Text.NewGuid",
    "Text.PadEnd",
    "Text.PadStart",
    "Text.PositionOf",
    "Text.PositionOfAny",
    "Text.Proper",
    "Text.Range",
    "Text.Remove",
    "Text.RemoveRange",
    "Text.Repeat",
    "Text.Replace",
    "Text.ReplaceRange",
    "Text.Select",
    "Text.Split",
    "Text.SplitAny",
    "Text.Start",
    "Text.StartsWith",
    "Text.ToBinary",
    "Text.ToList",
    "Text.Trim",
    "Text.TrimEnd",
    "Text.TrimStart",
    "Text.Upper",
    "Time.EndOfHour",
    "Time.From",
    "Time.FromText",
    "Time.Hour",
    "Time.Minute",
    "Time.Second",
    "Time.StartOfHour",
    "Time.ToRecord",
    "Time.ToText",
    "Type.AddTableKey",
    "Type.ClosedRecord",
    "Type.Facets",
    "Type.ForFunction",
    "Type.ForRecord",
    "Type.FunctionParameters",
    "Type.FunctionRequiredParameters",
    "Type.FunctionReturn",
    "Type.Is",
    "Type.IsNullable",
    "Type.IsOpenRecord",
    "Type.ListItem",
    "Type.NonNullable",
    "Type.OpenRecord",
    "Type.RecordFields",
    "Type.ReplaceFacets",
    "Type.ReplaceTableKeys",
    "Type.TableColumn",
    "Type.TableKeys",
    "Type.TableRow",
    "Type.TableSchema",
    "Type.Union",
    "Uri.BuildQueryString",
    "Uri.Combine",
    "Uri.EscapeDataString",
    "Uri.Parts",
    "Value.Add",
    "Value.As",
    "Value.Compare",
    "Value.Divide",
    "Value.Equals",
    "Value.Firewall",
    "Value.FromText",
    "Value.Is",
    "Value.Metadata",
    "Value.Multiply",
    "Value.NativeQuery",
    "Value.NullableEquals",
    "Value.RemoveMetadata",
    "Value.ReplaceMetadata",
    "Value.ReplaceType",
    "Value.Subtract",
    "Value.Type",
    "ValueAction.NativeStatement",
    "ValueAction.Replace",
    "Variable.Value",
    "Web.Contents",
    "Web.Page",
    "WebAction.Request",
    "Xml.Document",
    "Xml.Tables"
  ],
  builtinConstants: [
    "BinaryEncoding.Base64",
    "BinaryEncoding.Hex",
    "BinaryOccurrence.Optional",
    "BinaryOccurrence.Repeating",
    "BinaryOccurrence.Required",
    "ByteOrder.BigEndian",
    "ByteOrder.LittleEndian",
    "Compression.Deflate",
    "Compression.GZip",
    "CsvStyle.QuoteAfterDelimiter",
    "CsvStyle.QuoteAlways",
    "Culture.Current",
    "Day.Friday",
    "Day.Monday",
    "Day.Saturday",
    "Day.Sunday",
    "Day.Thursday",
    "Day.Tuesday",
    "Day.Wednesday",
    "ExtraValues.Error",
    "ExtraValues.Ignore",
    "ExtraValues.List",
    "GroupKind.Global",
    "GroupKind.Local",
    "JoinAlgorithm.Dynamic",
    "JoinAlgorithm.LeftHash",
    "JoinAlgorithm.LeftIndex",
    "JoinAlgorithm.PairwiseHash",
    "JoinAlgorithm.RightHash",
    "JoinAlgorithm.RightIndex",
    "JoinAlgorithm.SortMerge",
    "JoinKind.FullOuter",
    "JoinKind.Inner",
    "JoinKind.LeftAnti",
    "JoinKind.LeftOuter",
    "JoinKind.RightAnti",
    "JoinKind.RightOuter",
    "JoinSide.Left",
    "JoinSide.Right",
    "MissingField.Error",
    "MissingField.Ignore",
    "MissingField.UseNull",
    "Number.E",
    "Number.Epsilon",
    "Number.NaN",
    "Number.NegativeInfinity",
    "Number.PI",
    "Number.PositiveInfinity",
    "Occurrence.All",
    "Occurrence.First",
    "Occurrence.Last",
    "Occurrence.Optional",
    "Occurrence.Repeating",
    "Occurrence.Required",
    "Order.Ascending",
    "Order.Descending",
    "Precision.Decimal",
    "Precision.Double",
    "QuoteStyle.Csv",
    "QuoteStyle.None",
    "RelativePosition.FromEnd",
    "RelativePosition.FromStart",
    "RoundingMode.AwayFromZero",
    "RoundingMode.Down",
    "RoundingMode.ToEven",
    "RoundingMode.TowardZero",
    "RoundingMode.Up",
    "SapHanaDistribution.All",
    "SapHanaDistribution.Connection",
    "SapHanaDistribution.Off",
    "SapHanaDistribution.Statement",
    "SapHanaRangeOperator.Equals",
    "SapHanaRangeOperator.GreaterThan",
    "SapHanaRangeOperator.GreaterThanOrEquals",
    "SapHanaRangeOperator.LessThan",
    "SapHanaRangeOperator.LessThanOrEquals",
    "SapHanaRangeOperator.NotEquals",
    "TextEncoding.Ascii",
    "TextEncoding.BigEndianUnicode",
    "TextEncoding.Unicode",
    "TextEncoding.Utf16",
    "TextEncoding.Utf8",
    "TextEncoding.Windows",
    "TraceLevel.Critical",
    "TraceLevel.Error",
    "TraceLevel.Information",
    "TraceLevel.Verbose",
    "TraceLevel.Warning",
    "WebMethod.Delete",
    "WebMethod.Get",
    "WebMethod.Head",
    "WebMethod.Patch",
    "WebMethod.Post",
    "WebMethod.Put"
  ],
  builtinTypes: [
    "Action.Type",
    "Any.Type",
    "Binary.Type",
    "BinaryEncoding.Type",
    "BinaryOccurrence.Type",
    "Byte.Type",
    "ByteOrder.Type",
    "Character.Type",
    "Compression.Type",
    "CsvStyle.Type",
    "Currency.Type",
    "Date.Type",
    "DateTime.Type",
    "DateTimeZone.Type",
    "Day.Type",
    "Decimal.Type",
    "Double.Type",
    "Duration.Type",
    "ExtraValues.Type",
    "Function.Type",
    "GroupKind.Type",
    "Guid.Type",
    "Int16.Type",
    "Int32.Type",
    "Int64.Type",
    "Int8.Type",
    "JoinAlgorithm.Type",
    "JoinKind.Type",
    "JoinSide.Type",
    "List.Type",
    "Logical.Type",
    "MissingField.Type",
    "None.Type",
    "Null.Type",
    "Number.Type",
    "Occurrence.Type",
    "Order.Type",
    "Password.Type",
    "Percentage.Type",
    "Precision.Type",
    "QuoteStyle.Type",
    "Record.Type",
    "RelativePosition.Type",
    "RoundingMode.Type",
    "SapHanaDistribution.Type",
    "SapHanaRangeOperator.Type",
    "Single.Type",
    "Table.Type",
    "Text.Type",
    "TextEncoding.Type",
    "Time.Type",
    "TraceLevel.Type",
    "Type.Type",
    "Uri.Type",
    "WebMethod.Type"
  ],
  tokenizer: {
    root: [
      [/#"[\w \.]+"/, "identifier.quote"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F]+/, "number.hex"],
      [/\d+([eE][\-+]?\d+)?/, "number"],
      [
        /(#?[a-z]+)\b/,
        {
          cases: {
            "@typeKeywords": "type",
            "@keywords": "keyword",
            "@constants": "constant",
            "@constructors": "constructor",
            "@operatorKeywords": "operators",
            "@default": "identifier"
          }
        }
      ],
      [
        /\b([A-Z][a-zA-Z0-9]+\.Type)\b/,
        {
          cases: {
            "@builtinTypes": "type",
            "@default": "identifier"
          }
        }
      ],
      [
        /\b([A-Z][a-zA-Z0-9]+\.[A-Z][a-zA-Z0-9]+)\b/,
        {
          cases: {
            "@builtinFunctions": "keyword.function",
            "@builtinConstants": "constant",
            "@default": "identifier"
          }
        }
      ],
      [/\b([a-zA-Z_][\w\.]*)\b/, "identifier"],
      { include: "@whitespace" },
      { include: "@comments" },
      { include: "@strings" },
      [/[{}()\[\]]/, "@brackets"],
      [/([=\+<>\-\*&@\?\/!])|([<>]=)|(<>)|(=>)|(\.\.\.)|(\.\.)/, "operators"],
      [/[,;]/, "delimiter"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      ["\\/\\*", "comment", "@comment"],
      ["\\/\\/+.*", "comment"]
    ],
    comment: [
      ["\\*\\/", "comment", "@pop"],
      [".", "comment"]
    ],
    strings: [['"', "string", "@string"]],
    string: [
      ['""', "string.escape"],
      ['"', "string", "@pop"],
      [".", "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/powershell/powershell.js":
/*!************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/powershell/powershell.js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/powershell/powershell.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\#%\^\&\*\(\)\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,
  comments: {
    lineComment: "#",
    blockComment: ["<#", "#>"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#region\\b"),
      end: new RegExp("^\\s*#endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  ignoreCase: true,
  tokenPostfix: ".ps1",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.square", open: "[", close: "]" },
    { token: "delimiter.parenthesis", open: "(", close: ")" }
  ],
  keywords: [
    "begin",
    "break",
    "catch",
    "class",
    "continue",
    "data",
    "define",
    "do",
    "dynamicparam",
    "else",
    "elseif",
    "end",
    "exit",
    "filter",
    "finally",
    "for",
    "foreach",
    "from",
    "function",
    "if",
    "in",
    "param",
    "process",
    "return",
    "switch",
    "throw",
    "trap",
    "try",
    "until",
    "using",
    "var",
    "while",
    "workflow",
    "parallel",
    "sequence",
    "inlinescript",
    "configuration"
  ],
  helpKeywords: /SYNOPSIS|DESCRIPTION|PARAMETER|EXAMPLE|INPUTS|OUTPUTS|NOTES|LINK|COMPONENT|ROLE|FUNCTIONALITY|FORWARDHELPTARGETNAME|FORWARDHELPCATEGORY|REMOTEHELPRUNSPACE|EXTERNALHELP/,
  symbols: /[=><!~?&%|+\-*\/\^;\.,]+/,
  escapes: /`(?:[abfnrtv\\"'$]|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_][\w-]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": ""
          }
        }
      ],
      [/[ \t\r\n]+/, ""],
      [/^:\w*/, "metatag"],
      [
        /\$(\{((global|local|private|script|using):)?[\w]+\}|((global|local|private|script|using):)?[\w]+)/,
        "variable"
      ],
      [/<#/, "comment", "@comment"],
      [/#.*$/, "comment"],
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F_]*[0-9a-fA-F]/, "number.hex"],
      [/\d+?/, "number"],
      [/[;,.]/, "delimiter"],
      [/\@"/, "string", '@herestring."'],
      [/\@'/, "string", "@herestring.'"],
      [
        /"/,
        {
          cases: {
            "@eos": "string",
            "@default": { token: "string", next: '@string."' }
          }
        }
      ],
      [
        /'/,
        {
          cases: {
            "@eos": "string",
            "@default": { token: "string", next: "@string.'" }
          }
        }
      ]
    ],
    string: [
      [
        /[^"'\$`]+/,
        {
          cases: {
            "@eos": { token: "string", next: "@popall" },
            "@default": "string"
          }
        }
      ],
      [
        /@escapes/,
        {
          cases: {
            "@eos": { token: "string.escape", next: "@popall" },
            "@default": "string.escape"
          }
        }
      ],
      [
        /`./,
        {
          cases: {
            "@eos": {
              token: "string.escape.invalid",
              next: "@popall"
            },
            "@default": "string.escape.invalid"
          }
        }
      ],
      [
        /\$[\w]+$/,
        {
          cases: {
            '$S2=="': { token: "variable", next: "@popall" },
            "@default": { token: "string", next: "@popall" }
          }
        }
      ],
      [
        /\$[\w]+/,
        {
          cases: {
            '$S2=="': "variable",
            "@default": "string"
          }
        }
      ],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": {
              cases: {
                "@eos": { token: "string", next: "@popall" },
                "@default": "string"
              }
            }
          }
        }
      ]
    ],
    herestring: [
      [
        /^\s*(["'])@/,
        {
          cases: {
            "$1==$S2": { token: "string", next: "@pop" },
            "@default": "string"
          }
        }
      ],
      [/[^\$`]+/, "string"],
      [/@escapes/, "string.escape"],
      [/`./, "string.escape.invalid"],
      [
        /\$[\w]+/,
        {
          cases: {
            '$S2=="': "variable",
            "@default": "string"
          }
        }
      ]
    ],
    comment: [
      [/[^#\.]+/, "comment"],
      [/#>/, "comment", "@pop"],
      [/(\.)(@helpKeywords)(?!\w)/, { token: "comment.keyword.$2" }],
      [/[\.#]/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/protobuf/protobuf.js":
/*!********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/protobuf/protobuf.js ***!
  \********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/protobuf/protobuf.ts
var namedLiterals = ["true", "false"];
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">" },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "'", close: "'", notIn: ["string"] }
  ],
  autoCloseBefore: ".,=}])>' \n	",
  indentationRules: {
    increaseIndentPattern: new RegExp("^((?!\\/\\/).)*(\\{[^}\"'`]*|\\([^)\"'`]*|\\[[^\\]\"'`]*)$"),
    decreaseIndentPattern: new RegExp("^((?!.*?\\/\\*).*\\*/)?\\s*[\\}\\]].*$")
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".proto",
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  symbols: /[=><!~?:&|+\-*/^%]+/,
  keywords: [
    "syntax",
    "import",
    "weak",
    "public",
    "package",
    "option",
    "repeated",
    "oneof",
    "map",
    "reserved",
    "to",
    "max",
    "enum",
    "message",
    "service",
    "rpc",
    "stream",
    "returns",
    "package",
    "optional",
    "true",
    "false"
  ],
  builtinTypes: [
    "double",
    "float",
    "int32",
    "int64",
    "uint32",
    "uint64",
    "sint32",
    "sint64",
    "fixed32",
    "fixed64",
    "sfixed32",
    "sfixed64",
    "bool",
    "string",
    "bytes"
  ],
  operators: ["=", "+", "-"],
  namedLiterals,
  escapes: `\\\\(u{[0-9A-Fa-f]+}|n|r|t|\\\\|'|\\\${)`,
  identifier: /[a-zA-Z]\w*/,
  fullIdentifier: /@identifier(?:\s*\.\s*@identifier)*/,
  optionName: /(?:@identifier|\(\s*@fullIdentifier\s*\))(?:\s*\.\s*@identifier)*/,
  messageName: /@identifier/,
  enumName: /@identifier/,
  messageType: /\.?\s*(?:@identifier\s*\.\s*)*@messageName/,
  enumType: /\.?\s*(?:@identifier\s*\.\s*)*@enumName/,
  floatLit: /[0-9]+\s*\.\s*[0-9]*(?:@exponent)?|[0-9]+@exponent|\.[0-9]+(?:@exponent)?/,
  exponent: /[eE]\s*[+-]?\s*[0-9]+/,
  boolLit: /true\b|false\b/,
  decimalLit: /[1-9][0-9]*/,
  octalLit: /0[0-7]*/,
  hexLit: /0[xX][0-9a-fA-F]+/,
  type: /double|float|int32|int64|uint32|uint64|sint32|sint64|fixed32|fixed64|sfixed32|sfixed64|bool|string|bytes|@messageType|@enumType/,
  keyType: /int32|int64|uint32|uint64|sint32|sint64|fixed32|fixed64|sfixed32|sfixed64|bool|string/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      [/syntax/, "keyword"],
      [/=/, "operators"],
      [/;/, "delimiter"],
      [
        /(")(proto3)(")/,
        ["string.quote", "string", { token: "string.quote", switchTo: "@topLevel.proto3" }]
      ],
      [
        /(")(proto2)(")/,
        ["string.quote", "string", { token: "string.quote", switchTo: "@topLevel.proto2" }]
      ],
      [
        /.*?/,
        { token: "", switchTo: "@topLevel.proto2" }
      ]
    ],
    topLevel: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/=/, "operators"],
      [/[;.]/, "delimiter"],
      [
        /@fullIdentifier/,
        {
          cases: {
            option: { token: "keyword", next: "@option.$S2" },
            enum: { token: "keyword", next: "@enumDecl.$S2" },
            message: { token: "keyword", next: "@messageDecl.$S2" },
            service: { token: "keyword", next: "@serviceDecl.$S2" },
            extend: {
              cases: {
                "$S2==proto2": { token: "keyword", next: "@extendDecl.$S2" }
              }
            },
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ]
    ],
    enumDecl: [
      { include: "@whitespace" },
      [/@identifier/, "type.identifier"],
      [/{/, { token: "@brackets", bracket: "@open", switchTo: "@enumBody.$S2" }]
    ],
    enumBody: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/=/, "operators"],
      [/;/, "delimiter"],
      [/option\b/, "keyword", "@option.$S2"],
      [/@identifier/, "identifier"],
      [/\[/, { token: "@brackets", bracket: "@open", next: "@options.$S2" }],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    messageDecl: [
      { include: "@whitespace" },
      [/@identifier/, "type.identifier"],
      [/{/, { token: "@brackets", bracket: "@open", switchTo: "@messageBody.$S2" }]
    ],
    messageBody: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/=/, "operators"],
      [/;/, "delimiter"],
      [
        "(map)(s*)(<)",
        ["keyword", "white", { token: "@brackets", bracket: "@open", next: "@map.$S2" }]
      ],
      [
        /@identifier/,
        {
          cases: {
            option: { token: "keyword", next: "@option.$S2" },
            enum: { token: "keyword", next: "@enumDecl.$S2" },
            message: { token: "keyword", next: "@messageDecl.$S2" },
            oneof: { token: "keyword", next: "@oneofDecl.$S2" },
            extensions: {
              cases: {
                "$S2==proto2": { token: "keyword", next: "@reserved.$S2" }
              }
            },
            reserved: { token: "keyword", next: "@reserved.$S2" },
            "(?:repeated|optional)": { token: "keyword", next: "@field.$S2" },
            required: {
              cases: {
                "$S2==proto2": { token: "keyword", next: "@field.$S2" }
              }
            },
            "$S2==proto3": { token: "@rematch", next: "@field.$S2" }
          }
        }
      ],
      [/\[/, { token: "@brackets", bracket: "@open", next: "@options.$S2" }],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    extendDecl: [
      { include: "@whitespace" },
      [/@identifier/, "type.identifier"],
      [/{/, { token: "@brackets", bracket: "@open", switchTo: "@extendBody.$S2" }]
    ],
    extendBody: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/;/, "delimiter"],
      [/(?:repeated|optional|required)/, "keyword", "@field.$S2"],
      [/\[/, { token: "@brackets", bracket: "@open", next: "@options.$S2" }],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    options: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/;/, "delimiter"],
      [/@optionName/, "annotation"],
      [/[()]/, "annotation.brackets"],
      [/=/, "operator"],
      [/\]/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    option: [
      { include: "@whitespace" },
      [/@optionName/, "annotation"],
      [/[()]/, "annotation.brackets"],
      [/=/, "operator", "@pop"]
    ],
    oneofDecl: [
      { include: "@whitespace" },
      [/@identifier/, "identifier"],
      [/{/, { token: "@brackets", bracket: "@open", switchTo: "@oneofBody.$S2" }]
    ],
    oneofBody: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/;/, "delimiter"],
      [/(@identifier)(\s*)(=)/, ["identifier", "white", "delimiter"]],
      [
        /@fullIdentifier|\./,
        {
          cases: {
            "@builtinTypes": "keyword",
            "@default": "type.identifier"
          }
        }
      ],
      [/\[/, { token: "@brackets", bracket: "@open", next: "@options.$S2" }],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    reserved: [
      { include: "@whitespace" },
      [/,/, "delimiter"],
      [/;/, "delimiter", "@pop"],
      { include: "@constant" },
      [/to\b|max\b/, "keyword"]
    ],
    map: [
      { include: "@whitespace" },
      [
        /@fullIdentifier|\./,
        {
          cases: {
            "@builtinTypes": "keyword",
            "@default": "type.identifier"
          }
        }
      ],
      [/,/, "delimiter"],
      [/>/, { token: "@brackets", bracket: "@close", switchTo: "identifier" }]
    ],
    field: [
      { include: "@whitespace" },
      [
        "group",
        {
          cases: {
            "$S2==proto2": { token: "keyword", switchTo: "@groupDecl.$S2" }
          }
        }
      ],
      [/(@identifier)(\s*)(=)/, ["identifier", "white", { token: "delimiter", next: "@pop" }]],
      [
        /@fullIdentifier|\./,
        {
          cases: {
            "@builtinTypes": "keyword",
            "@default": "type.identifier"
          }
        }
      ]
    ],
    groupDecl: [
      { include: "@whitespace" },
      [/@identifier/, "identifier"],
      ["=", "operator"],
      [/{/, { token: "@brackets", bracket: "@open", switchTo: "@messageBody.$S2" }],
      { include: "@constant" }
    ],
    type: [
      { include: "@whitespace" },
      [/@identifier/, "type.identifier", "@pop"],
      [/./, "delimiter"]
    ],
    identifier: [{ include: "@whitespace" }, [/@identifier/, "identifier", "@pop"]],
    serviceDecl: [
      { include: "@whitespace" },
      [/@identifier/, "identifier"],
      [/{/, { token: "@brackets", bracket: "@open", switchTo: "@serviceBody.$S2" }]
    ],
    serviceBody: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/;/, "delimiter"],
      [/option\b/, "keyword", "@option.$S2"],
      [/rpc\b/, "keyword", "@rpc.$S2"],
      [/\[/, { token: "@brackets", bracket: "@open", next: "@options.$S2" }],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    rpc: [
      { include: "@whitespace" },
      [/@identifier/, "identifier"],
      [/\(/, { token: "@brackets", bracket: "@open", switchTo: "@request.$S2" }],
      [/{/, { token: "@brackets", bracket: "@open", next: "@methodOptions.$S2" }],
      [/;/, "delimiter", "@pop"]
    ],
    request: [
      { include: "@whitespace" },
      [
        /@messageType/,
        {
          cases: {
            stream: { token: "keyword", next: "@type.$S2" },
            "@default": "type.identifier"
          }
        }
      ],
      [/\)/, { token: "@brackets", bracket: "@close", switchTo: "@returns.$S2" }]
    ],
    returns: [
      { include: "@whitespace" },
      [/returns\b/, "keyword"],
      [/\(/, { token: "@brackets", bracket: "@open", switchTo: "@response.$S2" }]
    ],
    response: [
      { include: "@whitespace" },
      [
        /@messageType/,
        {
          cases: {
            stream: { token: "keyword", next: "@type.$S2" },
            "@default": "type.identifier"
          }
        }
      ],
      [/\)/, { token: "@brackets", bracket: "@close", switchTo: "@rpc.$S2" }]
    ],
    methodOptions: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/;/, "delimiter"],
      ["option", "keyword"],
      [/@optionName/, "annotation"],
      [/[()]/, "annotation.brackets"],
      [/=/, "operator"],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\/\*/, "comment", "@push"],
      ["\\*/", "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    stringSingle: [
      [/[^\\']+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/'/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    constant: [
      ["@boolLit", "keyword.constant"],
      ["@hexLit", "number.hex"],
      ["@octalLit", "number.octal"],
      ["@decimalLit", "number"],
      ["@floatLit", "number.float"],
      [/("([^"\\]|\\.)*|'([^'\\]|\\.)*)$/, "string.invalid"],
      [/"/, { token: "string.quote", bracket: "@open", next: "@string" }],
      [/'/, { token: "string.quote", bracket: "@open", next: "@stringSingle" }],
      [/{/, { token: "@brackets", bracket: "@open", next: "@prototext" }],
      [/identifier/, "identifier"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    prototext: [
      { include: "@whitespace" },
      { include: "@constant" },
      [/@identifier/, "identifier"],
      [/[:;]/, "delimiter"],
      [/}/, { token: "@brackets", bracket: "@close", next: "@pop" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/pug/pug.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/pug/pug.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/pug/pug.ts
var conf = {
  comments: {
    lineComment: "//"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] }
  ],
  folding: {
    offSide: true
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".pug",
  ignoreCase: true,
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.array", open: "[", close: "]" },
    { token: "delimiter.parenthesis", open: "(", close: ")" }
  ],
  keywords: [
    "append",
    "block",
    "case",
    "default",
    "doctype",
    "each",
    "else",
    "extends",
    "for",
    "if",
    "in",
    "include",
    "mixin",
    "typeof",
    "unless",
    "var",
    "when"
  ],
  tags: [
    "a",
    "abbr",
    "acronym",
    "address",
    "area",
    "article",
    "aside",
    "audio",
    "b",
    "base",
    "basefont",
    "bdi",
    "bdo",
    "blockquote",
    "body",
    "br",
    "button",
    "canvas",
    "caption",
    "center",
    "cite",
    "code",
    "col",
    "colgroup",
    "command",
    "datalist",
    "dd",
    "del",
    "details",
    "dfn",
    "div",
    "dl",
    "dt",
    "em",
    "embed",
    "fieldset",
    "figcaption",
    "figure",
    "font",
    "footer",
    "form",
    "frame",
    "frameset",
    "h1",
    "h2",
    "h3",
    "h4",
    "h5",
    "h6",
    "head",
    "header",
    "hgroup",
    "hr",
    "html",
    "i",
    "iframe",
    "img",
    "input",
    "ins",
    "keygen",
    "kbd",
    "label",
    "li",
    "link",
    "map",
    "mark",
    "menu",
    "meta",
    "meter",
    "nav",
    "noframes",
    "noscript",
    "object",
    "ol",
    "optgroup",
    "option",
    "output",
    "p",
    "param",
    "pre",
    "progress",
    "q",
    "rp",
    "rt",
    "ruby",
    "s",
    "samp",
    "script",
    "section",
    "select",
    "small",
    "source",
    "span",
    "strike",
    "strong",
    "style",
    "sub",
    "summary",
    "sup",
    "table",
    "tbody",
    "td",
    "textarea",
    "tfoot",
    "th",
    "thead",
    "time",
    "title",
    "tr",
    "tracks",
    "tt",
    "u",
    "ul",
    "video",
    "wbr"
  ],
  symbols: /[\+\-\*\%\&\|\!\=\/\.\,\:]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [
        /^(\s*)([a-zA-Z_-][\w-]*)/,
        {
          cases: {
            "$2@tags": {
              cases: {
                "@eos": ["", "tag"],
                "@default": ["", { token: "tag", next: "@tag.$1" }]
              }
            },
            "$2@keywords": ["", { token: "keyword.$2" }],
            "@default": ["", ""]
          }
        }
      ],
      [
        /^(\s*)(#[a-zA-Z_-][\w-]*)/,
        {
          cases: {
            "@eos": ["", "tag.id"],
            "@default": ["", { token: "tag.id", next: "@tag.$1" }]
          }
        }
      ],
      [
        /^(\s*)(\.[a-zA-Z_-][\w-]*)/,
        {
          cases: {
            "@eos": ["", "tag.class"],
            "@default": ["", { token: "tag.class", next: "@tag.$1" }]
          }
        }
      ],
      [/^(\s*)(\|.*)$/, ""],
      { include: "@whitespace" },
      [
        /[a-zA-Z_$][\w$]*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": ""
          }
        }
      ],
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/\d+\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/\d+/, "number"],
      [/"/, "string", '@string."'],
      [/'/, "string", "@string.'"]
    ],
    tag: [
      [/(\.)(\s*$)/, [{ token: "delimiter", next: "@blockText.$S2." }, ""]],
      [/\s+/, { token: "", next: "@simpleText" }],
      [
        /#[a-zA-Z_-][\w-]*/,
        {
          cases: {
            "@eos": { token: "tag.id", next: "@pop" },
            "@default": "tag.id"
          }
        }
      ],
      [
        /\.[a-zA-Z_-][\w-]*/,
        {
          cases: {
            "@eos": { token: "tag.class", next: "@pop" },
            "@default": "tag.class"
          }
        }
      ],
      [/\(/, { token: "delimiter.parenthesis", next: "@attributeList" }]
    ],
    simpleText: [
      [/[^#]+$/, { token: "", next: "@popall" }],
      [/[^#]+/, { token: "" }],
      [
        /(#{)([^}]*)(})/,
        {
          cases: {
            "@eos": [
              "interpolation.delimiter",
              "interpolation",
              {
                token: "interpolation.delimiter",
                next: "@popall"
              }
            ],
            "@default": ["interpolation.delimiter", "interpolation", "interpolation.delimiter"]
          }
        }
      ],
      [/#$/, { token: "", next: "@popall" }],
      [/#/, ""]
    ],
    attributeList: [
      [/\s+/, ""],
      [
        /(\w+)(\s*=\s*)("|')/,
        ["attribute.name", "delimiter", { token: "attribute.value", next: "@value.$3" }]
      ],
      [/\w+/, "attribute.name"],
      [
        /,/,
        {
          cases: {
            "@eos": {
              token: "attribute.delimiter",
              next: "@popall"
            },
            "@default": "attribute.delimiter"
          }
        }
      ],
      [/\)$/, { token: "delimiter.parenthesis", next: "@popall" }],
      [/\)/, { token: "delimiter.parenthesis", next: "@pop" }]
    ],
    whitespace: [
      [/^(\s*)(\/\/.*)$/, { token: "comment", next: "@blockText.$1.comment" }],
      [/[ \t\r\n]+/, ""],
      [/<!--/, { token: "comment", next: "@comment" }]
    ],
    blockText: [
      [
        /^\s+.*$/,
        {
          cases: {
            "($S2\\s+.*$)": { token: "$S3" },
            "@default": { token: "@rematch", next: "@popall" }
          }
        }
      ],
      [/./, { token: "@rematch", next: "@popall" }]
    ],
    comment: [
      [/[^<\-]+/, "comment.content"],
      [/-->/, { token: "comment", next: "@pop" }],
      [/<!--/, "comment.content.invalid"],
      [/[<\-]/, "comment.content"]
    ],
    string: [
      [
        /[^\\"'#]+/,
        {
          cases: {
            "@eos": { token: "string", next: "@popall" },
            "@default": "string"
          }
        }
      ],
      [
        /@escapes/,
        {
          cases: {
            "@eos": { token: "string.escape", next: "@popall" },
            "@default": "string.escape"
          }
        }
      ],
      [
        /\\./,
        {
          cases: {
            "@eos": {
              token: "string.escape.invalid",
              next: "@popall"
            },
            "@default": "string.escape.invalid"
          }
        }
      ],
      [/(#{)([^}]*)(})/, ["interpolation.delimiter", "interpolation", "interpolation.delimiter"]],
      [/#/, "string"],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "string", next: "@pop" },
            "@default": { token: "string" }
          }
        }
      ]
    ],
    value: [
      [
        /[^\\"']+/,
        {
          cases: {
            "@eos": { token: "attribute.value", next: "@popall" },
            "@default": "attribute.value"
          }
        }
      ],
      [
        /\\./,
        {
          cases: {
            "@eos": { token: "attribute.value", next: "@popall" },
            "@default": "attribute.value"
          }
        }
      ],
      [
        /["']/,
        {
          cases: {
            "$#==$S2": { token: "attribute.value", next: "@pop" },
            "@default": { token: "attribute.value" }
          }
        }
      ]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/python/python.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/python/python.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/python/python.ts
var conf = {
  comments: {
    lineComment: "#",
    blockComment: ["'''", "'''"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  onEnterRules: [
    {
      beforeText: new RegExp("^\\s*(?:def|class|for|if|elif|else|while|try|with|finally|except|async|match|case).*?:\\s*$"),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    }
  ],
  folding: {
    offSide: true,
    markers: {
      start: new RegExp("^\\s*#region\\b"),
      end: new RegExp("^\\s*#endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".python",
  keywords: [
    "False",
    "None",
    "True",
    "_",
    "and",
    "as",
    "assert",
    "async",
    "await",
    "break",
    "case",
    "class",
    "continue",
    "def",
    "del",
    "elif",
    "else",
    "except",
    "exec",
    "finally",
    "for",
    "from",
    "global",
    "if",
    "import",
    "in",
    "is",
    "lambda",
    "match",
    "nonlocal",
    "not",
    "or",
    "pass",
    "print",
    "raise",
    "return",
    "try",
    "while",
    "with",
    "yield",
    "int",
    "float",
    "long",
    "complex",
    "hex",
    "abs",
    "all",
    "any",
    "apply",
    "basestring",
    "bin",
    "bool",
    "buffer",
    "bytearray",
    "callable",
    "chr",
    "classmethod",
    "cmp",
    "coerce",
    "compile",
    "complex",
    "delattr",
    "dict",
    "dir",
    "divmod",
    "enumerate",
    "eval",
    "execfile",
    "file",
    "filter",
    "format",
    "frozenset",
    "getattr",
    "globals",
    "hasattr",
    "hash",
    "help",
    "id",
    "input",
    "intern",
    "isinstance",
    "issubclass",
    "iter",
    "len",
    "locals",
    "list",
    "map",
    "max",
    "memoryview",
    "min",
    "next",
    "object",
    "oct",
    "open",
    "ord",
    "pow",
    "print",
    "property",
    "reversed",
    "range",
    "raw_input",
    "reduce",
    "reload",
    "repr",
    "reversed",
    "round",
    "self",
    "set",
    "setattr",
    "slice",
    "sorted",
    "staticmethod",
    "str",
    "sum",
    "super",
    "tuple",
    "type",
    "unichr",
    "unicode",
    "vars",
    "xrange",
    "zip",
    "__dict__",
    "__methods__",
    "__members__",
    "__class__",
    "__bases__",
    "__name__",
    "__mro__",
    "__subclasses__",
    "__init__",
    "__import__"
  ],
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.bracket" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  tokenizer: {
    root: [
      { include: "@whitespace" },
      { include: "@numbers" },
      { include: "@strings" },
      [/[,:;]/, "delimiter"],
      [/[{}\[\]()]/, "@brackets"],
      [/@[a-zA-Z_]\w*/, "tag"],
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ]
    ],
    whitespace: [
      [/\s+/, "white"],
      [/(^#.*$)/, "comment"],
      [/'''/, "string", "@endDocString"],
      [/"""/, "string", "@endDblDocString"]
    ],
    endDocString: [
      [/[^']+/, "string"],
      [/\\'/, "string"],
      [/'''/, "string", "@popall"],
      [/'/, "string"]
    ],
    endDblDocString: [
      [/[^"]+/, "string"],
      [/\\"/, "string"],
      [/"""/, "string", "@popall"],
      [/"/, "string"]
    ],
    numbers: [
      [/-?0x([abcdef]|[ABCDEF]|\d)+[lL]?/, "number.hex"],
      [/-?(\d*\.)?\d+([eE][+\-]?\d+)?[jJ]?[lL]?/, "number"]
    ],
    strings: [
      [/'$/, "string.escape", "@popall"],
      [/'/, "string.escape", "@stringBody"],
      [/"$/, "string.escape", "@popall"],
      [/"/, "string.escape", "@dblStringBody"]
    ],
    stringBody: [
      [/[^\\']+$/, "string", "@popall"],
      [/[^\\']+/, "string"],
      [/\\./, "string"],
      [/'/, "string.escape", "@popall"],
      [/\\$/, "string"]
    ],
    dblStringBody: [
      [/[^\\"]+$/, "string", "@popall"],
      [/[^\\"]+/, "string"],
      [/\\./, "string"],
      [/"/, "string.escape", "@popall"],
      [/\\$/, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/qsharp/qsharp.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/qsharp/qsharp.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/qsharp/qsharp.ts
var conf = {
  comments: {
    lineComment: "//"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"', notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ]
};
var language = {
  keywords: [
    "namespace",
    "open",
    "as",
    "operation",
    "function",
    "body",
    "adjoint",
    "newtype",
    "controlled",
    "if",
    "elif",
    "else",
    "repeat",
    "until",
    "fixup",
    "for",
    "in",
    "while",
    "return",
    "fail",
    "within",
    "apply",
    "Adjoint",
    "Controlled",
    "Adj",
    "Ctl",
    "is",
    "self",
    "auto",
    "distribute",
    "invert",
    "intrinsic",
    "let",
    "set",
    "w/",
    "new",
    "not",
    "and",
    "or",
    "use",
    "borrow",
    "using",
    "borrowing",
    "mutable",
    "internal"
  ],
  typeKeywords: [
    "Unit",
    "Int",
    "BigInt",
    "Double",
    "Bool",
    "String",
    "Qubit",
    "Result",
    "Pauli",
    "Range"
  ],
  invalidKeywords: [
    "abstract",
    "base",
    "bool",
    "break",
    "byte",
    "case",
    "catch",
    "char",
    "checked",
    "class",
    "const",
    "continue",
    "decimal",
    "default",
    "delegate",
    "do",
    "double",
    "enum",
    "event",
    "explicit",
    "extern",
    "finally",
    "fixed",
    "float",
    "foreach",
    "goto",
    "implicit",
    "int",
    "interface",
    "lock",
    "long",
    "null",
    "object",
    "operator",
    "out",
    "override",
    "params",
    "private",
    "protected",
    "public",
    "readonly",
    "ref",
    "sbyte",
    "sealed",
    "short",
    "sizeof",
    "stackalloc",
    "static",
    "string",
    "struct",
    "switch",
    "this",
    "throw",
    "try",
    "typeof",
    "unit",
    "ulong",
    "unchecked",
    "unsafe",
    "ushort",
    "virtual",
    "void",
    "volatile"
  ],
  constants: ["true", "false", "PauliI", "PauliX", "PauliY", "PauliZ", "One", "Zero"],
  builtin: [
    "X",
    "Y",
    "Z",
    "H",
    "HY",
    "S",
    "T",
    "SWAP",
    "CNOT",
    "CCNOT",
    "MultiX",
    "R",
    "RFrac",
    "Rx",
    "Ry",
    "Rz",
    "R1",
    "R1Frac",
    "Exp",
    "ExpFrac",
    "Measure",
    "M",
    "MultiM",
    "Message",
    "Length",
    "Assert",
    "AssertProb",
    "AssertEqual"
  ],
  operators: [
    "and=",
    "<-",
    "->",
    "*",
    "*=",
    "@",
    "!",
    "^",
    "^=",
    ":",
    "::",
    "..",
    "==",
    "...",
    "=",
    "=>",
    ">",
    ">=",
    "<",
    "<=",
    "-",
    "-=",
    "!=",
    "or=",
    "%",
    "%=",
    "|",
    "+",
    "+=",
    "?",
    "/",
    "/=",
    "&&&",
    "&&&=",
    "^^^",
    "^^^=",
    ">>>",
    ">>>=",
    "<<<",
    "<<<=",
    "|||",
    "|||=",
    "~~~",
    "_",
    "w/",
    "w/="
  ],
  namespaceFollows: ["namespace", "open"],
  symbols: /[=><!~?:&|+\-*\/\^%@._]+/,
  escapes: /\\[\s\S]/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_$][\w$]*/,
        {
          cases: {
            "@namespaceFollows": {
              token: "keyword.$0",
              next: "@namespace"
            },
            "@typeKeywords": "type",
            "@keywords": "keyword",
            "@constants": "constant",
            "@builtin": "keyword",
            "@invalidKeywords": "invalid",
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, { cases: { "@operators": "operator", "@default": "" } }],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/\d+/, "number"],
      [/[;,.]/, "delimiter"],
      [/"/, { token: "string.quote", bracket: "@open", next: "@string" }]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/"/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    namespace: [
      { include: "@whitespace" },
      [/[A-Za-z]\w*/, "namespace"],
      [/[\.=]/, "delimiter"],
      ["", "", "@pop"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/(\/\/).*/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/r/r.js":
/*!******************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/r/r.js ***!
  \******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/r/r.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".r",
  roxygen: [
    "@alias",
    "@aliases",
    "@assignee",
    "@author",
    "@backref",
    "@callGraph",
    "@callGraphDepth",
    "@callGraphPrimitives",
    "@concept",
    "@describeIn",
    "@description",
    "@details",
    "@docType",
    "@encoding",
    "@evalNamespace",
    "@evalRd",
    "@example",
    "@examples",
    "@export",
    "@exportClass",
    "@exportMethod",
    "@exportPattern",
    "@family",
    "@field",
    "@formals",
    "@format",
    "@import",
    "@importClassesFrom",
    "@importFrom",
    "@importMethodsFrom",
    "@include",
    "@inherit",
    "@inheritDotParams",
    "@inheritParams",
    "@inheritSection",
    "@keywords",
    "@md",
    "@method",
    "@name",
    "@noMd",
    "@noRd",
    "@note",
    "@param",
    "@rawNamespace",
    "@rawRd",
    "@rdname",
    "@references",
    "@return",
    "@S3method",
    "@section",
    "@seealso",
    "@setClass",
    "@slot",
    "@source",
    "@template",
    "@templateVar",
    "@title",
    "@TODO",
    "@usage",
    "@useDynLib"
  ],
  constants: [
    "NULL",
    "FALSE",
    "TRUE",
    "NA",
    "Inf",
    "NaN",
    "NA_integer_",
    "NA_real_",
    "NA_complex_",
    "NA_character_",
    "T",
    "F",
    "LETTERS",
    "letters",
    "month.abb",
    "month.name",
    "pi",
    "R.version.string"
  ],
  keywords: [
    "break",
    "next",
    "return",
    "if",
    "else",
    "for",
    "in",
    "repeat",
    "while",
    "array",
    "category",
    "character",
    "complex",
    "double",
    "function",
    "integer",
    "list",
    "logical",
    "matrix",
    "numeric",
    "vector",
    "data.frame",
    "factor",
    "library",
    "require",
    "attach",
    "detach",
    "source"
  ],
  special: ["\\n", "\\r", "\\t", "\\b", "\\a", "\\f", "\\v", "\\'", '\\"', "\\\\"],
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.bracket" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  tokenizer: {
    root: [
      { include: "@numbers" },
      { include: "@strings" },
      [/[{}\[\]()]/, "@brackets"],
      { include: "@operators" },
      [/#'$/, "comment.doc"],
      [/#'/, "comment.doc", "@roxygen"],
      [/(^#.*$)/, "comment"],
      [/\s+/, "white"],
      [/[,:;]/, "delimiter"],
      [/@[a-zA-Z]\w*/, "tag"],
      [
        /[a-zA-Z]\w*/,
        {
          cases: {
            "@keywords": "keyword",
            "@constants": "constant",
            "@default": "identifier"
          }
        }
      ]
    ],
    roxygen: [
      [
        /@\w+/,
        {
          cases: {
            "@roxygen": "tag",
            "@eos": { token: "comment.doc", next: "@pop" },
            "@default": "comment.doc"
          }
        }
      ],
      [
        /\s+/,
        {
          cases: {
            "@eos": { token: "comment.doc", next: "@pop" },
            "@default": "comment.doc"
          }
        }
      ],
      [/.*/, { token: "comment.doc", next: "@pop" }]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]+/, "number.hex"],
      [/-?(\d*\.)?\d+([eE][+\-]?\d+)?/, "number"]
    ],
    operators: [
      [/<{1,2}-/, "operator"],
      [/->{1,2}/, "operator"],
      [/%[^%\s]+%/, "operator"],
      [/\*\*/, "operator"],
      [/%%/, "operator"],
      [/&&/, "operator"],
      [/\|\|/, "operator"],
      [/<</, "operator"],
      [/>>/, "operator"],
      [/[-+=&|!<>^~*/:$]/, "operator"]
    ],
    strings: [
      [/'/, "string.escape", "@stringBody"],
      [/"/, "string.escape", "@dblStringBody"]
    ],
    stringBody: [
      [
        /\\./,
        {
          cases: {
            "@special": "string",
            "@default": "error-token"
          }
        }
      ],
      [/'/, "string.escape", "@popall"],
      [/./, "string"]
    ],
    dblStringBody: [
      [
        /\\./,
        {
          cases: {
            "@special": "string",
            "@default": "error-token"
          }
        }
      ],
      [/"/, "string.escape", "@popall"],
      [/./, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/razor/razor.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/razor/razor.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/razor/razor.ts
var EMPTY_ELEMENTS = [
  "area",
  "base",
  "br",
  "col",
  "embed",
  "hr",
  "img",
  "input",
  "keygen",
  "link",
  "menuitem",
  "meta",
  "param",
  "source",
  "track",
  "wbr"
];
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\$\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\s]+)/g,
  comments: {
    blockComment: ["<!--", "-->"]
  },
  brackets: [
    ["<!--", "-->"],
    ["<", ">"],
    ["{", "}"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "<", close: ">" }
  ],
  onEnterRules: [
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      afterText: /^<\/(\w[\w\d]*)\s*>$/i,
      action: {
        indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
      }
    },
    {
      beforeText: new RegExp(`<(?!(?:${EMPTY_ELEMENTS.join("|")}))(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: "",
  tokenizer: {
    root: [
      [/@@@@/],
      [/@[^@]/, { token: "@rematch", switchTo: "@razorInSimpleState.root" }],
      [/<!DOCTYPE/, "metatag.html", "@doctype"],
      [/<!--/, "comment.html", "@comment"],
      [/(<)([\w\-]+)(\/>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/(<)(script)/, ["delimiter.html", { token: "tag.html", next: "@script" }]],
      [/(<)(style)/, ["delimiter.html", { token: "tag.html", next: "@style" }]],
      [/(<)([:\w\-]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/(<\/)([\w\-]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/</, "delimiter.html"],
      [/[ \t\r\n]+/],
      [/[^<@]+/]
    ],
    doctype: [
      [/@[^@]/, { token: "@rematch", switchTo: "@razorInSimpleState.comment" }],
      [/[^>]+/, "metatag.content.html"],
      [/>/, "metatag.html", "@pop"]
    ],
    comment: [
      [/@[^@]/, { token: "@rematch", switchTo: "@razorInSimpleState.comment" }],
      [/-->/, "comment.html", "@pop"],
      [/[^-]+/, "comment.content.html"],
      [/./, "comment.content.html"]
    ],
    otherTag: [
      [/@[^@]/, { token: "@rematch", switchTo: "@razorInSimpleState.otherTag" }],
      [/\/?>/, "delimiter.html", "@pop"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/]
    ],
    script: [
      [/@[^@]/, { token: "@rematch", switchTo: "@razorInSimpleState.script" }],
      [/type/, "attribute.name", "@scriptAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(script\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    scriptAfterType: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInSimpleState.scriptAfterType"
        }
      ],
      [/=/, "delimiter", "@scriptAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptAfterTypeEquals: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInSimpleState.scriptAfterTypeEquals"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.text/javascript",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptWithCustomType: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInSimpleState.scriptWithCustomType.$S2"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptEmbedded: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInEmbeddedState.scriptEmbedded.$S2",
          nextEmbedded: "@pop"
        }
      ],
      [/<\/script/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }]
    ],
    style: [
      [/@[^@]/, { token: "@rematch", switchTo: "@razorInSimpleState.style" }],
      [/type/, "attribute.name", "@styleAfterType"],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(style\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    styleAfterType: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInSimpleState.styleAfterType"
        }
      ],
      [/=/, "delimiter", "@styleAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleAfterTypeEquals: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInSimpleState.styleAfterTypeEquals"
        }
      ],
      [
        /"([^"]*)"/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.text/css",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleWithCustomType: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInSimpleState.styleWithCustomType.$S2"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value"],
      [/'([^']*)'/, "attribute.value"],
      [/[\w\-]+/, "attribute.name"],
      [/=/, "delimiter"],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleEmbedded: [
      [
        /@[^@]/,
        {
          token: "@rematch",
          switchTo: "@razorInEmbeddedState.styleEmbedded.$S2",
          nextEmbedded: "@pop"
        }
      ],
      [/<\/style/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }]
    ],
    razorInSimpleState: [
      [/@\*/, "comment.cs", "@razorBlockCommentTopLevel"],
      [/@[{(]/, "metatag.cs", "@razorRootTopLevel"],
      [/(@)(\s*[\w]+)/, ["metatag.cs", { token: "identifier.cs", switchTo: "@$S2.$S3" }]],
      [/[})]/, { token: "metatag.cs", switchTo: "@$S2.$S3" }],
      [/\*@/, { token: "comment.cs", switchTo: "@$S2.$S3" }]
    ],
    razorInEmbeddedState: [
      [/@\*/, "comment.cs", "@razorBlockCommentTopLevel"],
      [/@[{(]/, "metatag.cs", "@razorRootTopLevel"],
      [
        /(@)(\s*[\w]+)/,
        [
          "metatag.cs",
          {
            token: "identifier.cs",
            switchTo: "@$S2.$S3",
            nextEmbedded: "$S3"
          }
        ]
      ],
      [
        /[})]/,
        {
          token: "metatag.cs",
          switchTo: "@$S2.$S3",
          nextEmbedded: "$S3"
        }
      ],
      [
        /\*@/,
        {
          token: "comment.cs",
          switchTo: "@$S2.$S3",
          nextEmbedded: "$S3"
        }
      ]
    ],
    razorBlockCommentTopLevel: [
      [/\*@/, "@rematch", "@pop"],
      [/[^*]+/, "comment.cs"],
      [/./, "comment.cs"]
    ],
    razorBlockComment: [
      [/\*@/, "comment.cs", "@pop"],
      [/[^*]+/, "comment.cs"],
      [/./, "comment.cs"]
    ],
    razorRootTopLevel: [
      [/\{/, "delimiter.bracket.cs", "@razorRoot"],
      [/\(/, "delimiter.parenthesis.cs", "@razorRoot"],
      [/[})]/, "@rematch", "@pop"],
      { include: "razorCommon" }
    ],
    razorRoot: [
      [/\{/, "delimiter.bracket.cs", "@razorRoot"],
      [/\(/, "delimiter.parenthesis.cs", "@razorRoot"],
      [/\}/, "delimiter.bracket.cs", "@pop"],
      [/\)/, "delimiter.parenthesis.cs", "@pop"],
      { include: "razorCommon" }
    ],
    razorCommon: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@razorKeywords": { token: "keyword.cs" },
            "@default": "identifier.cs"
          }
        }
      ],
      [/[\[\]]/, "delimiter.array.cs"],
      [/[ \t\r\n]+/],
      [/\/\/.*$/, "comment.cs"],
      [/@\*/, "comment.cs", "@razorBlockComment"],
      [/"([^"]*)"/, "string.cs"],
      [/'([^']*)'/, "string.cs"],
      [/(<)([\w\-]+)(\/>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/(<)([\w\-]+)(>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/(<\/)([\w\-]+)(>)/, ["delimiter.html", "tag.html", "delimiter.html"]],
      [/[\+\-\*\%\&\|\^\~\!\=\<\>\/\?\;\:\.\,]/, "delimiter.cs"],
      [/\d*\d+[eE]([\-+]?\d+)?/, "number.float.cs"],
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float.cs"],
      [/0[xX][0-9a-fA-F']*[0-9a-fA-F]/, "number.hex.cs"],
      [/0[0-7']*[0-7]/, "number.octal.cs"],
      [/0[bB][0-1']*[0-1]/, "number.binary.cs"],
      [/\d[\d']*/, "number.cs"],
      [/\d/, "number.cs"]
    ]
  },
  razorKeywords: [
    "abstract",
    "as",
    "async",
    "await",
    "base",
    "bool",
    "break",
    "by",
    "byte",
    "case",
    "catch",
    "char",
    "checked",
    "class",
    "const",
    "continue",
    "decimal",
    "default",
    "delegate",
    "do",
    "double",
    "descending",
    "explicit",
    "event",
    "extern",
    "else",
    "enum",
    "false",
    "finally",
    "fixed",
    "float",
    "for",
    "foreach",
    "from",
    "goto",
    "group",
    "if",
    "implicit",
    "in",
    "int",
    "interface",
    "internal",
    "into",
    "is",
    "lock",
    "long",
    "nameof",
    "new",
    "null",
    "namespace",
    "object",
    "operator",
    "out",
    "override",
    "orderby",
    "params",
    "private",
    "protected",
    "public",
    "readonly",
    "ref",
    "return",
    "switch",
    "struct",
    "sbyte",
    "sealed",
    "short",
    "sizeof",
    "stackalloc",
    "static",
    "string",
    "select",
    "this",
    "throw",
    "true",
    "try",
    "typeof",
    "uint",
    "ulong",
    "unchecked",
    "unsafe",
    "ushort",
    "using",
    "var",
    "virtual",
    "volatile",
    "void",
    "when",
    "while",
    "where",
    "yield",
    "model",
    "inject"
  ],
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/redis/redis.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/redis/redis.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/redis/redis.ts
var conf = {
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".redis",
  ignoreCase: true,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "APPEND",
    "AUTH",
    "BGREWRITEAOF",
    "BGSAVE",
    "BITCOUNT",
    "BITFIELD",
    "BITOP",
    "BITPOS",
    "BLPOP",
    "BRPOP",
    "BRPOPLPUSH",
    "CLIENT",
    "KILL",
    "LIST",
    "GETNAME",
    "PAUSE",
    "REPLY",
    "SETNAME",
    "CLUSTER",
    "ADDSLOTS",
    "COUNT-FAILURE-REPORTS",
    "COUNTKEYSINSLOT",
    "DELSLOTS",
    "FAILOVER",
    "FORGET",
    "GETKEYSINSLOT",
    "INFO",
    "KEYSLOT",
    "MEET",
    "NODES",
    "REPLICATE",
    "RESET",
    "SAVECONFIG",
    "SET-CONFIG-EPOCH",
    "SETSLOT",
    "SLAVES",
    "SLOTS",
    "COMMAND",
    "COUNT",
    "GETKEYS",
    "CONFIG",
    "GET",
    "REWRITE",
    "SET",
    "RESETSTAT",
    "DBSIZE",
    "DEBUG",
    "OBJECT",
    "SEGFAULT",
    "DECR",
    "DECRBY",
    "DEL",
    "DISCARD",
    "DUMP",
    "ECHO",
    "EVAL",
    "EVALSHA",
    "EXEC",
    "EXISTS",
    "EXPIRE",
    "EXPIREAT",
    "FLUSHALL",
    "FLUSHDB",
    "GEOADD",
    "GEOHASH",
    "GEOPOS",
    "GEODIST",
    "GEORADIUS",
    "GEORADIUSBYMEMBER",
    "GETBIT",
    "GETRANGE",
    "GETSET",
    "HDEL",
    "HEXISTS",
    "HGET",
    "HGETALL",
    "HINCRBY",
    "HINCRBYFLOAT",
    "HKEYS",
    "HLEN",
    "HMGET",
    "HMSET",
    "HSET",
    "HSETNX",
    "HSTRLEN",
    "HVALS",
    "INCR",
    "INCRBY",
    "INCRBYFLOAT",
    "KEYS",
    "LASTSAVE",
    "LINDEX",
    "LINSERT",
    "LLEN",
    "LPOP",
    "LPUSH",
    "LPUSHX",
    "LRANGE",
    "LREM",
    "LSET",
    "LTRIM",
    "MGET",
    "MIGRATE",
    "MONITOR",
    "MOVE",
    "MSET",
    "MSETNX",
    "MULTI",
    "PERSIST",
    "PEXPIRE",
    "PEXPIREAT",
    "PFADD",
    "PFCOUNT",
    "PFMERGE",
    "PING",
    "PSETEX",
    "PSUBSCRIBE",
    "PUBSUB",
    "PTTL",
    "PUBLISH",
    "PUNSUBSCRIBE",
    "QUIT",
    "RANDOMKEY",
    "READONLY",
    "READWRITE",
    "RENAME",
    "RENAMENX",
    "RESTORE",
    "ROLE",
    "RPOP",
    "RPOPLPUSH",
    "RPUSH",
    "RPUSHX",
    "SADD",
    "SAVE",
    "SCARD",
    "SCRIPT",
    "FLUSH",
    "LOAD",
    "SDIFF",
    "SDIFFSTORE",
    "SELECT",
    "SETBIT",
    "SETEX",
    "SETNX",
    "SETRANGE",
    "SHUTDOWN",
    "SINTER",
    "SINTERSTORE",
    "SISMEMBER",
    "SLAVEOF",
    "SLOWLOG",
    "SMEMBERS",
    "SMOVE",
    "SORT",
    "SPOP",
    "SRANDMEMBER",
    "SREM",
    "STRLEN",
    "SUBSCRIBE",
    "SUNION",
    "SUNIONSTORE",
    "SWAPDB",
    "SYNC",
    "TIME",
    "TOUCH",
    "TTL",
    "TYPE",
    "UNSUBSCRIBE",
    "UNLINK",
    "UNWATCH",
    "WAIT",
    "WATCH",
    "ZADD",
    "ZCARD",
    "ZCOUNT",
    "ZINCRBY",
    "ZINTERSTORE",
    "ZLEXCOUNT",
    "ZRANGE",
    "ZRANGEBYLEX",
    "ZREVRANGEBYLEX",
    "ZRANGEBYSCORE",
    "ZRANK",
    "ZREM",
    "ZREMRANGEBYLEX",
    "ZREMRANGEBYRANK",
    "ZREMRANGEBYSCORE",
    "ZREVRANGE",
    "ZREVRANGEBYSCORE",
    "ZREVRANK",
    "ZSCORE",
    "ZUNIONSTORE",
    "SCAN",
    "SSCAN",
    "HSCAN",
    "ZSCAN"
  ],
  operators: [],
  builtinFunctions: [],
  builtinVariables: [],
  pseudoColumns: [],
  tokenizer: {
    root: [
      { include: "@whitespace" },
      { include: "@pseudoColumns" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@scopes" },
      [/[;,.]/, "delimiter"],
      [/[()]/, "@brackets"],
      [
        /[\w@#$]+/,
        {
          cases: {
            "@keywords": "keyword",
            "@operators": "operator",
            "@builtinVariables": "predefined",
            "@builtinFunctions": "predefined",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=!%&+\-*/|~^]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    pseudoColumns: [
      [
        /[$][A-Za-z_][\w@#$]*/,
        {
          cases: {
            "@pseudoColumns": "predefined",
            "@default": "identifier"
          }
        }
      ]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*/, "number"],
      [/[$][+-]*\d*(\.\d*)?/, "number"],
      [/((\d+(\.\d*)?)|(\.\d+))([eE][\-+]?\d+)?/, "number"]
    ],
    strings: [
      [/'/, { token: "string", next: "@string" }],
      [/"/, { token: "string.double", next: "@stringDouble" }]
    ],
    string: [
      [/[^']+/, "string"],
      [/''/, "string"],
      [/'/, { token: "string", next: "@pop" }]
    ],
    stringDouble: [
      [/[^"]+/, "string.double"],
      [/""/, "string.double"],
      [/"/, { token: "string.double", next: "@pop" }]
    ],
    scopes: []
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/redshift/redshift.js":
/*!********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/redshift/redshift.js ***!
  \********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/redshift/redshift.ts
var conf = {
  comments: {
    lineComment: "--",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sql",
  ignoreCase: true,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "AES128",
    "AES256",
    "ALL",
    "ALLOWOVERWRITE",
    "ANALYSE",
    "ANALYZE",
    "AND",
    "ANY",
    "ARRAY",
    "AS",
    "ASC",
    "AUTHORIZATION",
    "AZ64",
    "BACKUP",
    "BETWEEN",
    "BINARY",
    "BLANKSASNULL",
    "BOTH",
    "BYTEDICT",
    "BZIP2",
    "CASE",
    "CAST",
    "CHECK",
    "COLLATE",
    "COLUMN",
    "CONSTRAINT",
    "CREATE",
    "CREDENTIALS",
    "CROSS",
    "CURRENT_DATE",
    "CURRENT_TIME",
    "CURRENT_TIMESTAMP",
    "CURRENT_USER",
    "CURRENT_USER_ID",
    "DEFAULT",
    "DEFERRABLE",
    "DEFLATE",
    "DEFRAG",
    "DELTA",
    "DELTA32K",
    "DESC",
    "DISABLE",
    "DISTINCT",
    "DO",
    "ELSE",
    "EMPTYASNULL",
    "ENABLE",
    "ENCODE",
    "ENCRYPT",
    "ENCRYPTION",
    "END",
    "EXCEPT",
    "EXPLICIT",
    "FALSE",
    "FOR",
    "FOREIGN",
    "FREEZE",
    "FROM",
    "FULL",
    "GLOBALDICT256",
    "GLOBALDICT64K",
    "GRANT",
    "GROUP",
    "GZIP",
    "HAVING",
    "IDENTITY",
    "IGNORE",
    "ILIKE",
    "IN",
    "INITIALLY",
    "INNER",
    "INTERSECT",
    "INTO",
    "IS",
    "ISNULL",
    "JOIN",
    "LANGUAGE",
    "LEADING",
    "LEFT",
    "LIKE",
    "LIMIT",
    "LOCALTIME",
    "LOCALTIMESTAMP",
    "LUN",
    "LUNS",
    "LZO",
    "LZOP",
    "MINUS",
    "MOSTLY16",
    "MOSTLY32",
    "MOSTLY8",
    "NATURAL",
    "NEW",
    "NOT",
    "NOTNULL",
    "NULL",
    "NULLS",
    "OFF",
    "OFFLINE",
    "OFFSET",
    "OID",
    "OLD",
    "ON",
    "ONLY",
    "OPEN",
    "OR",
    "ORDER",
    "OUTER",
    "OVERLAPS",
    "PARALLEL",
    "PARTITION",
    "PERCENT",
    "PERMISSIONS",
    "PLACING",
    "PRIMARY",
    "RAW",
    "READRATIO",
    "RECOVER",
    "REFERENCES",
    "RESPECT",
    "REJECTLOG",
    "RESORT",
    "RESTORE",
    "RIGHT",
    "SELECT",
    "SESSION_USER",
    "SIMILAR",
    "SNAPSHOT",
    "SOME",
    "SYSDATE",
    "SYSTEM",
    "TABLE",
    "TAG",
    "TDES",
    "TEXT255",
    "TEXT32K",
    "THEN",
    "TIMESTAMP",
    "TO",
    "TOP",
    "TRAILING",
    "TRUE",
    "TRUNCATECOLUMNS",
    "UNION",
    "UNIQUE",
    "USER",
    "USING",
    "VERBOSE",
    "WALLET",
    "WHEN",
    "WHERE",
    "WITH",
    "WITHOUT"
  ],
  operators: [
    "AND",
    "BETWEEN",
    "IN",
    "LIKE",
    "NOT",
    "OR",
    "IS",
    "NULL",
    "INTERSECT",
    "UNION",
    "INNER",
    "JOIN",
    "LEFT",
    "OUTER",
    "RIGHT"
  ],
  builtinFunctions: [
    "current_schema",
    "current_schemas",
    "has_database_privilege",
    "has_schema_privilege",
    "has_table_privilege",
    "age",
    "current_time",
    "current_timestamp",
    "localtime",
    "isfinite",
    "now",
    "ascii",
    "get_bit",
    "get_byte",
    "set_bit",
    "set_byte",
    "to_ascii",
    "approximate percentile_disc",
    "avg",
    "count",
    "listagg",
    "max",
    "median",
    "min",
    "percentile_cont",
    "stddev_samp",
    "stddev_pop",
    "sum",
    "var_samp",
    "var_pop",
    "bit_and",
    "bit_or",
    "bool_and",
    "bool_or",
    "cume_dist",
    "first_value",
    "lag",
    "last_value",
    "lead",
    "nth_value",
    "ratio_to_report",
    "dense_rank",
    "ntile",
    "percent_rank",
    "rank",
    "row_number",
    "case",
    "coalesce",
    "decode",
    "greatest",
    "least",
    "nvl",
    "nvl2",
    "nullif",
    "add_months",
    "at time zone",
    "convert_timezone",
    "current_date",
    "date_cmp",
    "date_cmp_timestamp",
    "date_cmp_timestamptz",
    "date_part_year",
    "dateadd",
    "datediff",
    "date_part",
    "date_trunc",
    "extract",
    "getdate",
    "interval_cmp",
    "last_day",
    "months_between",
    "next_day",
    "sysdate",
    "timeofday",
    "timestamp_cmp",
    "timestamp_cmp_date",
    "timestamp_cmp_timestamptz",
    "timestamptz_cmp",
    "timestamptz_cmp_date",
    "timestamptz_cmp_timestamp",
    "timezone",
    "to_timestamp",
    "trunc",
    "abs",
    "acos",
    "asin",
    "atan",
    "atan2",
    "cbrt",
    "ceil",
    "ceiling",
    "checksum",
    "cos",
    "cot",
    "degrees",
    "dexp",
    "dlog1",
    "dlog10",
    "exp",
    "floor",
    "ln",
    "log",
    "mod",
    "pi",
    "power",
    "radians",
    "random",
    "round",
    "sin",
    "sign",
    "sqrt",
    "tan",
    "to_hex",
    "bpcharcmp",
    "btrim",
    "bttext_pattern_cmp",
    "char_length",
    "character_length",
    "charindex",
    "chr",
    "concat",
    "crc32",
    "func_sha1",
    "initcap",
    "left and rights",
    "len",
    "length",
    "lower",
    "lpad and rpads",
    "ltrim",
    "md5",
    "octet_length",
    "position",
    "quote_ident",
    "quote_literal",
    "regexp_count",
    "regexp_instr",
    "regexp_replace",
    "regexp_substr",
    "repeat",
    "replace",
    "replicate",
    "reverse",
    "rtrim",
    "split_part",
    "strpos",
    "strtol",
    "substring",
    "textlen",
    "translate",
    "trim",
    "upper",
    "cast",
    "convert",
    "to_char",
    "to_date",
    "to_number",
    "json_array_length",
    "json_extract_array_element_text",
    "json_extract_path_text",
    "current_setting",
    "pg_cancel_backend",
    "pg_terminate_backend",
    "set_config",
    "current_database",
    "current_user",
    "current_user_id",
    "pg_backend_pid",
    "pg_last_copy_count",
    "pg_last_copy_id",
    "pg_last_query_id",
    "pg_last_unload_count",
    "session_user",
    "slice_num",
    "user",
    "version",
    "abbrev",
    "acosd",
    "any",
    "area",
    "array_agg",
    "array_append",
    "array_cat",
    "array_dims",
    "array_fill",
    "array_length",
    "array_lower",
    "array_ndims",
    "array_position",
    "array_positions",
    "array_prepend",
    "array_remove",
    "array_replace",
    "array_to_json",
    "array_to_string",
    "array_to_tsvector",
    "array_upper",
    "asind",
    "atan2d",
    "atand",
    "bit",
    "bit_length",
    "bound_box",
    "box",
    "brin_summarize_new_values",
    "broadcast",
    "cardinality",
    "center",
    "circle",
    "clock_timestamp",
    "col_description",
    "concat_ws",
    "convert_from",
    "convert_to",
    "corr",
    "cosd",
    "cotd",
    "covar_pop",
    "covar_samp",
    "current_catalog",
    "current_query",
    "current_role",
    "currval",
    "cursor_to_xml",
    "diameter",
    "div",
    "encode",
    "enum_first",
    "enum_last",
    "enum_range",
    "every",
    "family",
    "format",
    "format_type",
    "generate_series",
    "generate_subscripts",
    "get_current_ts_config",
    "gin_clean_pending_list",
    "grouping",
    "has_any_column_privilege",
    "has_column_privilege",
    "has_foreign_data_wrapper_privilege",
    "has_function_privilege",
    "has_language_privilege",
    "has_sequence_privilege",
    "has_server_privilege",
    "has_tablespace_privilege",
    "has_type_privilege",
    "height",
    "host",
    "hostmask",
    "inet_client_addr",
    "inet_client_port",
    "inet_merge",
    "inet_same_family",
    "inet_server_addr",
    "inet_server_port",
    "isclosed",
    "isempty",
    "isopen",
    "json_agg",
    "json_object",
    "json_object_agg",
    "json_populate_record",
    "json_populate_recordset",
    "json_to_record",
    "json_to_recordset",
    "jsonb_agg",
    "jsonb_object_agg",
    "justify_days",
    "justify_hours",
    "justify_interval",
    "lastval",
    "left",
    "line",
    "localtimestamp",
    "lower_inc",
    "lower_inf",
    "lpad",
    "lseg",
    "make_date",
    "make_interval",
    "make_time",
    "make_timestamp",
    "make_timestamptz",
    "masklen",
    "mode",
    "netmask",
    "network",
    "nextval",
    "npoints",
    "num_nonnulls",
    "num_nulls",
    "numnode",
    "obj_description",
    "overlay",
    "parse_ident",
    "path",
    "pclose",
    "percentile_disc",
    "pg_advisory_lock",
    "pg_advisory_lock_shared",
    "pg_advisory_unlock",
    "pg_advisory_unlock_all",
    "pg_advisory_unlock_shared",
    "pg_advisory_xact_lock",
    "pg_advisory_xact_lock_shared",
    "pg_backup_start_time",
    "pg_blocking_pids",
    "pg_client_encoding",
    "pg_collation_is_visible",
    "pg_column_size",
    "pg_conf_load_time",
    "pg_control_checkpoint",
    "pg_control_init",
    "pg_control_recovery",
    "pg_control_system",
    "pg_conversion_is_visible",
    "pg_create_logical_replication_slot",
    "pg_create_physical_replication_slot",
    "pg_create_restore_point",
    "pg_current_xlog_flush_location",
    "pg_current_xlog_insert_location",
    "pg_current_xlog_location",
    "pg_database_size",
    "pg_describe_object",
    "pg_drop_replication_slot",
    "pg_export_snapshot",
    "pg_filenode_relation",
    "pg_function_is_visible",
    "pg_get_constraintdef",
    "pg_get_expr",
    "pg_get_function_arguments",
    "pg_get_function_identity_arguments",
    "pg_get_function_result",
    "pg_get_functiondef",
    "pg_get_indexdef",
    "pg_get_keywords",
    "pg_get_object_address",
    "pg_get_owned_sequence",
    "pg_get_ruledef",
    "pg_get_serial_sequence",
    "pg_get_triggerdef",
    "pg_get_userbyid",
    "pg_get_viewdef",
    "pg_has_role",
    "pg_identify_object",
    "pg_identify_object_as_address",
    "pg_index_column_has_property",
    "pg_index_has_property",
    "pg_indexam_has_property",
    "pg_indexes_size",
    "pg_is_in_backup",
    "pg_is_in_recovery",
    "pg_is_other_temp_schema",
    "pg_is_xlog_replay_paused",
    "pg_last_committed_xact",
    "pg_last_xact_replay_timestamp",
    "pg_last_xlog_receive_location",
    "pg_last_xlog_replay_location",
    "pg_listening_channels",
    "pg_logical_emit_message",
    "pg_logical_slot_get_binary_changes",
    "pg_logical_slot_get_changes",
    "pg_logical_slot_peek_binary_changes",
    "pg_logical_slot_peek_changes",
    "pg_ls_dir",
    "pg_my_temp_schema",
    "pg_notification_queue_usage",
    "pg_opclass_is_visible",
    "pg_operator_is_visible",
    "pg_opfamily_is_visible",
    "pg_options_to_table",
    "pg_postmaster_start_time",
    "pg_read_binary_file",
    "pg_read_file",
    "pg_relation_filenode",
    "pg_relation_filepath",
    "pg_relation_size",
    "pg_reload_conf",
    "pg_replication_origin_create",
    "pg_replication_origin_drop",
    "pg_replication_origin_oid",
    "pg_replication_origin_progress",
    "pg_replication_origin_session_is_setup",
    "pg_replication_origin_session_progress",
    "pg_replication_origin_session_reset",
    "pg_replication_origin_session_setup",
    "pg_replication_origin_xact_reset",
    "pg_replication_origin_xact_setup",
    "pg_rotate_logfile",
    "pg_size_bytes",
    "pg_size_pretty",
    "pg_sleep",
    "pg_sleep_for",
    "pg_sleep_until",
    "pg_start_backup",
    "pg_stat_file",
    "pg_stop_backup",
    "pg_switch_xlog",
    "pg_table_is_visible",
    "pg_table_size",
    "pg_tablespace_databases",
    "pg_tablespace_location",
    "pg_tablespace_size",
    "pg_total_relation_size",
    "pg_trigger_depth",
    "pg_try_advisory_lock",
    "pg_try_advisory_lock_shared",
    "pg_try_advisory_xact_lock",
    "pg_try_advisory_xact_lock_shared",
    "pg_ts_config_is_visible",
    "pg_ts_dict_is_visible",
    "pg_ts_parser_is_visible",
    "pg_ts_template_is_visible",
    "pg_type_is_visible",
    "pg_typeof",
    "pg_xact_commit_timestamp",
    "pg_xlog_location_diff",
    "pg_xlog_replay_pause",
    "pg_xlog_replay_resume",
    "pg_xlogfile_name",
    "pg_xlogfile_name_offset",
    "phraseto_tsquery",
    "plainto_tsquery",
    "point",
    "polygon",
    "popen",
    "pqserverversion",
    "query_to_xml",
    "querytree",
    "quote_nullable",
    "radius",
    "range_merge",
    "regexp_matches",
    "regexp_split_to_array",
    "regexp_split_to_table",
    "regr_avgx",
    "regr_avgy",
    "regr_count",
    "regr_intercept",
    "regr_r2",
    "regr_slope",
    "regr_sxx",
    "regr_sxy",
    "regr_syy",
    "right",
    "row_security_active",
    "row_to_json",
    "rpad",
    "scale",
    "set_masklen",
    "setseed",
    "setval",
    "setweight",
    "shobj_description",
    "sind",
    "sprintf",
    "statement_timestamp",
    "stddev",
    "string_agg",
    "string_to_array",
    "strip",
    "substr",
    "table_to_xml",
    "table_to_xml_and_xmlschema",
    "tand",
    "text",
    "to_json",
    "to_regclass",
    "to_regnamespace",
    "to_regoper",
    "to_regoperator",
    "to_regproc",
    "to_regprocedure",
    "to_regrole",
    "to_regtype",
    "to_tsquery",
    "to_tsvector",
    "transaction_timestamp",
    "ts_debug",
    "ts_delete",
    "ts_filter",
    "ts_headline",
    "ts_lexize",
    "ts_parse",
    "ts_rank",
    "ts_rank_cd",
    "ts_rewrite",
    "ts_stat",
    "ts_token_type",
    "tsquery_phrase",
    "tsvector_to_array",
    "tsvector_update_trigger",
    "tsvector_update_trigger_column",
    "txid_current",
    "txid_current_snapshot",
    "txid_snapshot_xip",
    "txid_snapshot_xmax",
    "txid_snapshot_xmin",
    "txid_visible_in_snapshot",
    "unnest",
    "upper_inc",
    "upper_inf",
    "variance",
    "width",
    "width_bucket",
    "xml_is_well_formed",
    "xml_is_well_formed_content",
    "xml_is_well_formed_document",
    "xmlagg",
    "xmlcomment",
    "xmlconcat",
    "xmlelement",
    "xmlexists",
    "xmlforest",
    "xmlparse",
    "xmlpi",
    "xmlroot",
    "xmlserialize",
    "xpath",
    "xpath_exists"
  ],
  builtinVariables: [],
  pseudoColumns: [],
  tokenizer: {
    root: [
      { include: "@comments" },
      { include: "@whitespace" },
      { include: "@pseudoColumns" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@complexIdentifiers" },
      { include: "@scopes" },
      [/[;,.]/, "delimiter"],
      [/[()]/, "@brackets"],
      [
        /[\w@#$]+/,
        {
          cases: {
            "@keywords": "keyword",
            "@operators": "operator",
            "@builtinVariables": "predefined",
            "@builtinFunctions": "predefined",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=!%&+\-*/|~^]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      [/--+.*/, "comment"],
      [/\/\*/, { token: "comment.quote", next: "@comment" }]
    ],
    comment: [
      [/[^*/]+/, "comment"],
      [/\*\//, { token: "comment.quote", next: "@pop" }],
      [/./, "comment"]
    ],
    pseudoColumns: [
      [
        /[$][A-Za-z_][\w@#$]*/,
        {
          cases: {
            "@pseudoColumns": "predefined",
            "@default": "identifier"
          }
        }
      ]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*/, "number"],
      [/[$][+-]*\d*(\.\d*)?/, "number"],
      [/((\d+(\.\d*)?)|(\.\d+))([eE][\-+]?\d+)?/, "number"]
    ],
    strings: [[/'/, { token: "string", next: "@string" }]],
    string: [
      [/[^']+/, "string"],
      [/''/, "string"],
      [/'/, { token: "string", next: "@pop" }]
    ],
    complexIdentifiers: [[/"/, { token: "identifier.quote", next: "@quotedIdentifier" }]],
    quotedIdentifier: [
      [/[^"]+/, "identifier"],
      [/""/, "identifier"],
      [/"/, { token: "identifier.quote", next: "@pop" }]
    ],
    scopes: []
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/restructuredtext/restructuredtext.js":
/*!************************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/restructuredtext/restructuredtext.js ***!
  \************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/restructuredtext/restructuredtext.ts
var conf = {
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: "<", close: ">", notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "(", close: ")" },
    { open: "[", close: "]" },
    { open: "`", close: "`" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*<!--\\s*#?region\\b.*-->"),
      end: new RegExp("^\\s*<!--\\s*#?endregion\\b.*-->")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".rst",
  control: /[\\`*_\[\]{}()#+\-\.!]/,
  escapes: /\\(?:@control)/,
  empty: [
    "area",
    "base",
    "basefont",
    "br",
    "col",
    "frame",
    "hr",
    "img",
    "input",
    "isindex",
    "link",
    "meta",
    "param"
  ],
  alphanumerics: /[A-Za-z0-9]/,
  simpleRefNameWithoutBq: /(?:@alphanumerics[-_+:.]*@alphanumerics)+|(?:@alphanumerics+)/,
  simpleRefName: /(?:`@phrase`|@simpleRefNameWithoutBq)/,
  phrase: /@simpleRefNameWithoutBq(?:\s@simpleRefNameWithoutBq)*/,
  citationName: /[A-Za-z][A-Za-z0-9-_.]*/,
  blockLiteralStart: /(?:[!"#$%&'()*+,-./:;<=>?@\[\]^_`{|}~]|[\s])/,
  precedingChars: /(?:[ -:/'"<([{])/,
  followingChars: /(?:[ -.,:;!?/'")\]}>]|$)/,
  punctuation: /(=|-|~|`|#|"|\^|\+|\*|:|\.|'|_|\+)/,
  tokenizer: {
    root: [
      [/^(@punctuation{3,}$){1,1}?/, "keyword"],
      [/^\s*([\*\-+‣•]|[a-zA-Z0-9]+\.|\([a-zA-Z0-9]+\)|[a-zA-Z0-9]+\))\s/, "keyword"],
      [/([ ]::)\s*$/, "keyword", "@blankLineOfLiteralBlocks"],
      [/(::)\s*$/, "keyword", "@blankLineOfLiteralBlocks"],
      { include: "@tables" },
      { include: "@explicitMarkupBlocks" },
      { include: "@inlineMarkup" }
    ],
    explicitMarkupBlocks: [
      { include: "@citations" },
      { include: "@footnotes" },
      [
        /^(\.\.\s)(@simpleRefName)(::\s)(.*)$/,
        [{ token: "", next: "subsequentLines" }, "keyword", "", ""]
      ],
      [
        /^(\.\.)(\s+)(_)(@simpleRefName)(:)(\s+)(.*)/,
        [{ token: "", next: "hyperlinks" }, "", "", "string.link", "", "", "string.link"]
      ],
      [
        /^((?:(?:\.\.)(?:\s+))?)(__)(:)(\s+)(.*)/,
        [{ token: "", next: "subsequentLines" }, "", "", "", "string.link"]
      ],
      [/^(__\s+)(.+)/, ["", "string.link"]],
      [
        /^(\.\.)( \|)([^| ]+[^|]*[^| ]*)(\| )(@simpleRefName)(:: .*)/,
        [{ token: "", next: "subsequentLines" }, "", "string.link", "", "keyword", ""],
        "@rawBlocks"
      ],
      [/(\|)([^| ]+[^|]*[^| ]*)(\|_{0,2})/, ["", "string.link", ""]],
      [/^(\.\.)([ ].*)$/, [{ token: "", next: "@comments" }, "comment"]]
    ],
    inlineMarkup: [
      { include: "@citationsReference" },
      { include: "@footnotesReference" },
      [/(@simpleRefName)(_{1,2})/, ["string.link", ""]],
      [/(`)([^<`]+\s+)(<)(.*)(>)(`)(_)/, ["", "string.link", "", "string.link", "", "", ""]],
      [/\*\*([^\\*]|\*(?!\*))+\*\*/, "strong"],
      [/\*[^*]+\*/, "emphasis"],
      [/(``)((?:[^`]|\`(?!`))+)(``)/, ["", "keyword", ""]],
      [/(__\s+)(.+)/, ["", "keyword"]],
      [/(:)((?:@simpleRefNameWithoutBq)?)(:`)([^`]+)(`)/, ["", "keyword", "", "", ""]],
      [/(`)([^`]+)(`:)((?:@simpleRefNameWithoutBq)?)(:)/, ["", "", "", "keyword", ""]],
      [/(`)([^`]+)(`)/, ""],
      [/(_`)(@phrase)(`)/, ["", "string.link", ""]]
    ],
    citations: [
      [
        /^(\.\.\s+\[)((?:@citationName))(\]\s+)(.*)/,
        [{ token: "", next: "@subsequentLines" }, "string.link", "", ""]
      ]
    ],
    citationsReference: [[/(\[)(@citationName)(\]_)/, ["", "string.link", ""]]],
    footnotes: [
      [
        /^(\.\.\s+\[)((?:[0-9]+))(\]\s+.*)/,
        [{ token: "", next: "@subsequentLines" }, "string.link", ""]
      ],
      [
        /^(\.\.\s+\[)((?:#@simpleRefName?))(\]\s+)(.*)/,
        [{ token: "", next: "@subsequentLines" }, "string.link", "", ""]
      ],
      [
        /^(\.\.\s+\[)((?:\*))(\]\s+)(.*)/,
        [{ token: "", next: "@subsequentLines" }, "string.link", "", ""]
      ]
    ],
    footnotesReference: [
      [/(\[)([0-9]+)(\])(_)/, ["", "string.link", "", ""]],
      [/(\[)(#@simpleRefName?)(\])(_)/, ["", "string.link", "", ""]],
      [/(\[)(\*)(\])(_)/, ["", "string.link", "", ""]]
    ],
    blankLineOfLiteralBlocks: [
      [/^$/, "", "@subsequentLinesOfLiteralBlocks"],
      [/^.*$/, "", "@pop"]
    ],
    subsequentLinesOfLiteralBlocks: [
      [/(@blockLiteralStart+)(.*)/, ["keyword", ""]],
      [/^(?!blockLiteralStart)/, "", "@popall"]
    ],
    subsequentLines: [
      [/^[\s]+.*/, ""],
      [/^(?!\s)/, "", "@pop"]
    ],
    hyperlinks: [
      [/^[\s]+.*/, "string.link"],
      [/^(?!\s)/, "", "@pop"]
    ],
    comments: [
      [/^[\s]+.*/, "comment"],
      [/^(?!\s)/, "", "@pop"]
    ],
    tables: [
      [/\+-[+-]+/, "keyword"],
      [/\+=[+=]+/, "keyword"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/ruby/ruby.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/ruby/ruby.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/ruby/ruby.ts
var conf = {
  comments: {
    lineComment: "#",
    blockComment: ["=begin", "=end"]
  },
  brackets: [
    ["(", ")"],
    ["{", "}"],
    ["[", "]"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  indentationRules: {
    increaseIndentPattern: new RegExp(`^\\s*((begin|class|(private|protected)\\s+def|def|else|elsif|ensure|for|if|module|rescue|unless|until|when|while|case)|([^#]*\\sdo\\b)|([^#]*=\\s*(case|if|unless)))\\b([^#\\{;]|("|'|/).*\\4)*(#.*)?$`),
    decreaseIndentPattern: new RegExp("^\\s*([}\\]]([,)]?\\s*(#|$)|\\.[a-zA-Z_]\\w*\\b)|(end|rescue|ensure|else|elsif|when)\\b)")
  }
};
var language = {
  tokenPostfix: ".ruby",
  keywords: [
    "__LINE__",
    "__ENCODING__",
    "__FILE__",
    "BEGIN",
    "END",
    "alias",
    "and",
    "begin",
    "break",
    "case",
    "class",
    "def",
    "defined?",
    "do",
    "else",
    "elsif",
    "end",
    "ensure",
    "for",
    "false",
    "if",
    "in",
    "module",
    "next",
    "nil",
    "not",
    "or",
    "redo",
    "rescue",
    "retry",
    "return",
    "self",
    "super",
    "then",
    "true",
    "undef",
    "unless",
    "until",
    "when",
    "while",
    "yield"
  ],
  keywordops: ["::", "..", "...", "?", ":", "=>"],
  builtins: [
    "require",
    "public",
    "private",
    "include",
    "extend",
    "attr_reader",
    "protected",
    "private_class_method",
    "protected_class_method",
    "new"
  ],
  declarations: [
    "module",
    "class",
    "def",
    "case",
    "do",
    "begin",
    "for",
    "if",
    "while",
    "until",
    "unless"
  ],
  linedecls: ["def", "case", "do", "begin", "for", "if", "while", "until", "unless"],
  operators: [
    "^",
    "&",
    "|",
    "<=>",
    "==",
    "===",
    "!~",
    "=~",
    ">",
    ">=",
    "<",
    "<=",
    "<<",
    ">>",
    "+",
    "-",
    "*",
    "/",
    "%",
    "**",
    "~",
    "+@",
    "-@",
    "[]",
    "[]=",
    "`",
    "+=",
    "-=",
    "*=",
    "**=",
    "/=",
    "^=",
    "%=",
    "<<=",
    ">>=",
    "&=",
    "&&=",
    "||=",
    "|="
  ],
  brackets: [
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" }
  ],
  symbols: /[=><!~?:&|+\-*\/\^%\.]+/,
  escape: /(?:[abefnrstv\\"'\n\r]|[0-7]{1,3}|x[0-9A-Fa-f]{1,2}|u[0-9A-Fa-f]{4})/,
  escapes: /\\(?:C\-(@escape|.)|c(@escape|.)|@escape)/,
  decpart: /\d(_?\d)*/,
  decimal: /0|@decpart/,
  delim: /[^a-zA-Z0-9\s\n\r]/,
  heredelim: /(?:\w+|'[^']*'|"[^"]*"|`[^`]*`)/,
  regexpctl: /[(){}\[\]\$\^|\-*+?\.]/,
  regexpesc: /\\(?:[AzZbBdDfnrstvwWn0\\\/]|@regexpctl|c[A-Z]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4})?/,
  tokenizer: {
    root: [
      [
        /^(\s*)([a-z_]\w*[!?=]?)/,
        [
          "white",
          {
            cases: {
              "for|until|while": {
                token: "keyword.$2",
                next: "@dodecl.$2"
              },
              "@declarations": {
                token: "keyword.$2",
                next: "@root.$2"
              },
              end: { token: "keyword.$S2", next: "@pop" },
              "@keywords": "keyword",
              "@builtins": "predefined",
              "@default": "identifier"
            }
          }
        ]
      ],
      [
        /[a-z_]\w*[!?=]?/,
        {
          cases: {
            "if|unless|while|until": {
              token: "keyword.$0x",
              next: "@modifier.$0x"
            },
            for: { token: "keyword.$2", next: "@dodecl.$2" },
            "@linedecls": { token: "keyword.$0", next: "@root.$0" },
            end: { token: "keyword.$S2", next: "@pop" },
            "@keywords": "keyword",
            "@builtins": "predefined",
            "@default": "identifier"
          }
        }
      ],
      [/[A-Z][\w]*[!?=]?/, "constructor.identifier"],
      [/\$[\w]*/, "global.constant"],
      [/@[\w]*/, "namespace.instance.identifier"],
      [/@@@[\w]*/, "namespace.class.identifier"],
      [/<<[-~](@heredelim).*/, { token: "string.heredoc.delimiter", next: "@heredoc.$1" }],
      [/[ \t\r\n]+<<(@heredelim).*/, { token: "string.heredoc.delimiter", next: "@heredoc.$1" }],
      [/^<<(@heredelim).*/, { token: "string.heredoc.delimiter", next: "@heredoc.$1" }],
      { include: "@whitespace" },
      [/"/, { token: "string.d.delim", next: '@dstring.d."' }],
      [/'/, { token: "string.sq.delim", next: "@sstring.sq" }],
      [/%([rsqxwW]|Q?)/, { token: "@rematch", next: "pstring" }],
      [/`/, { token: "string.x.delim", next: "@dstring.x.`" }],
      [/:(\w|[$@])\w*[!?=]?/, "string.s"],
      [/:"/, { token: "string.s.delim", next: '@dstring.s."' }],
      [/:'/, { token: "string.s.delim", next: "@sstring.s" }],
      [/\/(?=(\\\/|[^\/\n])+\/)/, { token: "regexp.delim", next: "@regexp" }],
      [/[{}()\[\]]/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@keywordops": "keyword",
            "@operators": "operator",
            "@default": ""
          }
        }
      ],
      [/[;,]/, "delimiter"],
      [/0[xX][0-9a-fA-F](_?[0-9a-fA-F])*/, "number.hex"],
      [/0[_oO][0-7](_?[0-7])*/, "number.octal"],
      [/0[bB][01](_?[01])*/, "number.binary"],
      [/0[dD]@decpart/, "number"],
      [
        /@decimal((\.@decpart)?([eE][\-+]?@decpart)?)/,
        {
          cases: {
            $1: "number.float",
            "@default": "number"
          }
        }
      ]
    ],
    dodecl: [
      [/^/, { token: "", switchTo: "@root.$S2" }],
      [
        /[a-z_]\w*[!?=]?/,
        {
          cases: {
            end: { token: "keyword.$S2", next: "@pop" },
            do: { token: "keyword", switchTo: "@root.$S2" },
            "@linedecls": {
              token: "@rematch",
              switchTo: "@root.$S2"
            },
            "@keywords": "keyword",
            "@builtins": "predefined",
            "@default": "identifier"
          }
        }
      ],
      { include: "@root" }
    ],
    modifier: [
      [/^/, "", "@pop"],
      [
        /[a-z_]\w*[!?=]?/,
        {
          cases: {
            end: { token: "keyword.$S2", next: "@pop" },
            "then|else|elsif|do": {
              token: "keyword",
              switchTo: "@root.$S2"
            },
            "@linedecls": {
              token: "@rematch",
              switchTo: "@root.$S2"
            },
            "@keywords": "keyword",
            "@builtins": "predefined",
            "@default": "identifier"
          }
        }
      ],
      { include: "@root" }
    ],
    sstring: [
      [/[^\\']+/, "string.$S2"],
      [/\\\\|\\'|\\$/, "string.$S2.escape"],
      [/\\./, "string.$S2.invalid"],
      [/'/, { token: "string.$S2.delim", next: "@pop" }]
    ],
    dstring: [
      [/[^\\`"#]+/, "string.$S2"],
      [/#/, "string.$S2.escape", "@interpolated"],
      [/\\$/, "string.$S2.escape"],
      [/@escapes/, "string.$S2.escape"],
      [/\\./, "string.$S2.escape.invalid"],
      [
        /[`"]/,
        {
          cases: {
            "$#==$S3": { token: "string.$S2.delim", next: "@pop" },
            "@default": "string.$S2"
          }
        }
      ]
    ],
    heredoc: [
      [
        /^(\s*)(@heredelim)$/,
        {
          cases: {
            "$2==$S2": ["string.heredoc", { token: "string.heredoc.delimiter", next: "@pop" }],
            "@default": ["string.heredoc", "string.heredoc"]
          }
        }
      ],
      [/.*/, "string.heredoc"]
    ],
    interpolated: [
      [/\$\w*/, "global.constant", "@pop"],
      [/@\w*/, "namespace.class.identifier", "@pop"],
      [/@@@\w*/, "namespace.instance.identifier", "@pop"],
      [
        /[{]/,
        {
          token: "string.escape.curly",
          switchTo: "@interpolated_compound"
        }
      ],
      ["", "", "@pop"]
    ],
    interpolated_compound: [
      [/[}]/, { token: "string.escape.curly", next: "@pop" }],
      { include: "@root" }
    ],
    pregexp: [
      { include: "@whitespace" },
      [
        /[^\(\{\[\\]/,
        {
          cases: {
            "$#==$S3": { token: "regexp.delim", next: "@pop" },
            "$#==$S2": { token: "regexp.delim", next: "@push" },
            "~[)}\\]]": "@brackets.regexp.escape.control",
            "~@regexpctl": "regexp.escape.control",
            "@default": "regexp"
          }
        }
      ],
      { include: "@regexcontrol" }
    ],
    regexp: [
      { include: "@regexcontrol" },
      [/[^\\\/]/, "regexp"],
      ["/[ixmp]*", { token: "regexp.delim" }, "@pop"]
    ],
    regexcontrol: [
      [
        /(\{)(\d+(?:,\d*)?)(\})/,
        [
          "@brackets.regexp.escape.control",
          "regexp.escape.control",
          "@brackets.regexp.escape.control"
        ]
      ],
      [
        /(\[)(\^?)/,
        ["@brackets.regexp.escape.control", { token: "regexp.escape.control", next: "@regexrange" }]
      ],
      [/(\()(\?[:=!])/, ["@brackets.regexp.escape.control", "regexp.escape.control"]],
      [/\(\?#/, { token: "regexp.escape.control", next: "@regexpcomment" }],
      [/[()]/, "@brackets.regexp.escape.control"],
      [/@regexpctl/, "regexp.escape.control"],
      [/\\$/, "regexp.escape"],
      [/@regexpesc/, "regexp.escape"],
      [/\\\./, "regexp.invalid"],
      [/#/, "regexp.escape", "@interpolated"]
    ],
    regexrange: [
      [/-/, "regexp.escape.control"],
      [/\^/, "regexp.invalid"],
      [/\\$/, "regexp.escape"],
      [/@regexpesc/, "regexp.escape"],
      [/[^\]]/, "regexp"],
      [/\]/, "@brackets.regexp.escape.control", "@pop"]
    ],
    regexpcomment: [
      [/[^)]+/, "comment"],
      [/\)/, { token: "regexp.escape.control", next: "@pop" }]
    ],
    pstring: [
      [/%([qws])\(/, { token: "string.$1.delim", switchTo: "@qstring.$1.(.)" }],
      [/%([qws])\[/, { token: "string.$1.delim", switchTo: "@qstring.$1.[.]" }],
      [/%([qws])\{/, { token: "string.$1.delim", switchTo: "@qstring.$1.{.}" }],
      [/%([qws])</, { token: "string.$1.delim", switchTo: "@qstring.$1.<.>" }],
      [/%([qws])(@delim)/, { token: "string.$1.delim", switchTo: "@qstring.$1.$2.$2" }],
      [/%r\(/, { token: "regexp.delim", switchTo: "@pregexp.(.)" }],
      [/%r\[/, { token: "regexp.delim", switchTo: "@pregexp.[.]" }],
      [/%r\{/, { token: "regexp.delim", switchTo: "@pregexp.{.}" }],
      [/%r</, { token: "regexp.delim", switchTo: "@pregexp.<.>" }],
      [/%r(@delim)/, { token: "regexp.delim", switchTo: "@pregexp.$1.$1" }],
      [/%(x|W|Q?)\(/, { token: "string.$1.delim", switchTo: "@qqstring.$1.(.)" }],
      [/%(x|W|Q?)\[/, { token: "string.$1.delim", switchTo: "@qqstring.$1.[.]" }],
      [/%(x|W|Q?)\{/, { token: "string.$1.delim", switchTo: "@qqstring.$1.{.}" }],
      [/%(x|W|Q?)</, { token: "string.$1.delim", switchTo: "@qqstring.$1.<.>" }],
      [/%(x|W|Q?)(@delim)/, { token: "string.$1.delim", switchTo: "@qqstring.$1.$2.$2" }],
      [/%([rqwsxW]|Q?)./, { token: "invalid", next: "@pop" }],
      [/./, { token: "invalid", next: "@pop" }]
    ],
    qstring: [
      [/\\$/, "string.$S2.escape"],
      [/\\./, "string.$S2.escape"],
      [
        /./,
        {
          cases: {
            "$#==$S4": { token: "string.$S2.delim", next: "@pop" },
            "$#==$S3": { token: "string.$S2.delim", next: "@push" },
            "@default": "string.$S2"
          }
        }
      ]
    ],
    qqstring: [[/#/, "string.$S2.escape", "@interpolated"], { include: "@qstring" }],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/^\s*=begin\b/, "comment", "@comment"],
      [/#.*$/, "comment"]
    ],
    comment: [
      [/[^=]+/, "comment"],
      [/^\s*=begin\b/, "comment.invalid"],
      [/^\s*=end\b.*/, "comment", "@pop"],
      [/[=]/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/rust/rust.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/rust/rust.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/rust/rust.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "(", close: ")" },
    { open: '"', close: '"', notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#pragma\\s+region\\b"),
      end: new RegExp("^\\s*#pragma\\s+endregion\\b")
    }
  }
};
var language = {
  tokenPostfix: ".rust",
  defaultToken: "invalid",
  keywords: [
    "as",
    "async",
    "await",
    "box",
    "break",
    "const",
    "continue",
    "crate",
    "dyn",
    "else",
    "enum",
    "extern",
    "false",
    "fn",
    "for",
    "if",
    "impl",
    "in",
    "let",
    "loop",
    "match",
    "mod",
    "move",
    "mut",
    "pub",
    "ref",
    "return",
    "self",
    "static",
    "struct",
    "super",
    "trait",
    "true",
    "try",
    "type",
    "unsafe",
    "use",
    "where",
    "while",
    "catch",
    "default",
    "union",
    "static",
    "abstract",
    "alignof",
    "become",
    "do",
    "final",
    "macro",
    "offsetof",
    "override",
    "priv",
    "proc",
    "pure",
    "sizeof",
    "typeof",
    "unsized",
    "virtual",
    "yield"
  ],
  typeKeywords: [
    "Self",
    "m32",
    "m64",
    "m128",
    "f80",
    "f16",
    "f128",
    "int",
    "uint",
    "float",
    "char",
    "bool",
    "u8",
    "u16",
    "u32",
    "u64",
    "f32",
    "f64",
    "i8",
    "i16",
    "i32",
    "i64",
    "str",
    "Option",
    "Either",
    "c_float",
    "c_double",
    "c_void",
    "FILE",
    "fpos_t",
    "DIR",
    "dirent",
    "c_char",
    "c_schar",
    "c_uchar",
    "c_short",
    "c_ushort",
    "c_int",
    "c_uint",
    "c_long",
    "c_ulong",
    "size_t",
    "ptrdiff_t",
    "clock_t",
    "time_t",
    "c_longlong",
    "c_ulonglong",
    "intptr_t",
    "uintptr_t",
    "off_t",
    "dev_t",
    "ino_t",
    "pid_t",
    "mode_t",
    "ssize_t"
  ],
  constants: ["true", "false", "Some", "None", "Left", "Right", "Ok", "Err"],
  supportConstants: [
    "EXIT_FAILURE",
    "EXIT_SUCCESS",
    "RAND_MAX",
    "EOF",
    "SEEK_SET",
    "SEEK_CUR",
    "SEEK_END",
    "_IOFBF",
    "_IONBF",
    "_IOLBF",
    "BUFSIZ",
    "FOPEN_MAX",
    "FILENAME_MAX",
    "L_tmpnam",
    "TMP_MAX",
    "O_RDONLY",
    "O_WRONLY",
    "O_RDWR",
    "O_APPEND",
    "O_CREAT",
    "O_EXCL",
    "O_TRUNC",
    "S_IFIFO",
    "S_IFCHR",
    "S_IFBLK",
    "S_IFDIR",
    "S_IFREG",
    "S_IFMT",
    "S_IEXEC",
    "S_IWRITE",
    "S_IREAD",
    "S_IRWXU",
    "S_IXUSR",
    "S_IWUSR",
    "S_IRUSR",
    "F_OK",
    "R_OK",
    "W_OK",
    "X_OK",
    "STDIN_FILENO",
    "STDOUT_FILENO",
    "STDERR_FILENO"
  ],
  supportMacros: [
    "format!",
    "print!",
    "println!",
    "panic!",
    "format_args!",
    "unreachable!",
    "write!",
    "writeln!"
  ],
  operators: [
    "!",
    "!=",
    "%",
    "%=",
    "&",
    "&=",
    "&&",
    "*",
    "*=",
    "+",
    "+=",
    "-",
    "-=",
    "->",
    ".",
    "..",
    "...",
    "/",
    "/=",
    ":",
    ";",
    "<<",
    "<<=",
    "<",
    "<=",
    "=",
    "==",
    "=>",
    ">",
    ">=",
    ">>",
    ">>=",
    "@",
    "^",
    "^=",
    "|",
    "|=",
    "||",
    "_",
    "?",
    "#"
  ],
  escapes: /\\([nrt0\"''\\]|x\h{2}|u\{\h{1,6}\})/,
  delimiters: /[,]/,
  symbols: /[\#\!\%\&\*\+\-\.\/\:\;\<\=\>\@\^\|_\?]+/,
  intSuffixes: /[iu](8|16|32|64|128|size)/,
  floatSuffixes: /f(32|64)/,
  tokenizer: {
    root: [
      [/r(#*)"/, { token: "string.quote", bracket: "@open", next: "@stringraw.$1" }],
      [
        /[a-zA-Z][a-zA-Z0-9_]*!?|_[a-zA-Z0-9_]+/,
        {
          cases: {
            "@typeKeywords": "keyword.type",
            "@keywords": "keyword",
            "@supportConstants": "keyword",
            "@supportMacros": "keyword",
            "@constants": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/\$/, "identifier"],
      [/'[a-zA-Z_][a-zA-Z0-9_]*(?=[^\'])/, "identifier"],
      [/'(\S|@escapes)'/, "string.byteliteral"],
      [/"/, { token: "string.quote", bracket: "@open", next: "@string" }],
      { include: "@numbers" },
      { include: "@whitespace" },
      [
        /@delimiters/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "delimiter"
          }
        }
      ],
      [/[{}()\[\]<>]/, "@brackets"],
      [/@symbols/, { cases: { "@operators": "operator", "@default": "" } }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\/\*/, "comment", "@push"],
      ["\\*/", "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    stringraw: [
      [/[^"#]+/, { token: "string" }],
      [
        /"(#*)/,
        {
          cases: {
            "$1==$S2": { token: "string.quote", bracket: "@close", next: "@pop" },
            "@default": { token: "string" }
          }
        }
      ],
      [/["#]/, { token: "string" }]
    ],
    numbers: [
      [/(0o[0-7_]+)(@intSuffixes)?/, { token: "number" }],
      [/(0b[0-1_]+)(@intSuffixes)?/, { token: "number" }],
      [/[\d][\d_]*(\.[\d][\d_]*)?[eE][+-][\d_]+(@floatSuffixes)?/, { token: "number" }],
      [/\b(\d\.?[\d_]*)(@floatSuffixes)?\b/, { token: "number" }],
      [/(0x[\da-fA-F]+)_?(@intSuffixes)?/, { token: "number" }],
      [/[\d][\d_]*(@intSuffixes?)?/, { token: "number" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/sb/sb.js":
/*!********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/sb/sb.js ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/sb/sb.ts
var conf = {
  comments: {
    lineComment: "'"
  },
  brackets: [
    ["(", ")"],
    ["[", "]"],
    ["If", "EndIf"],
    ["While", "EndWhile"],
    ["For", "EndFor"],
    ["Sub", "EndSub"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sb",
  ignoreCase: true,
  brackets: [
    { token: "delimiter.array", open: "[", close: "]" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "keyword.tag-if", open: "If", close: "EndIf" },
    { token: "keyword.tag-while", open: "While", close: "EndWhile" },
    { token: "keyword.tag-for", open: "For", close: "EndFor" },
    { token: "keyword.tag-sub", open: "Sub", close: "EndSub" }
  ],
  keywords: [
    "Else",
    "ElseIf",
    "EndFor",
    "EndIf",
    "EndSub",
    "EndWhile",
    "For",
    "Goto",
    "If",
    "Step",
    "Sub",
    "Then",
    "To",
    "While"
  ],
  tagwords: ["If", "Sub", "While", "For"],
  operators: [">", "<", "<>", "<=", ">=", "And", "Or", "+", "-", "*", "/", "="],
  identifier: /[a-zA-Z_][\w]*/,
  symbols: /[=><:+\-*\/%\.,]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      [/(@identifier)(?=[.])/, "type"],
      [
        /@identifier/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@operators": "operator",
            "@default": "variable.name"
          }
        }
      ],
      [
        /([.])(@identifier)/,
        {
          cases: {
            $2: ["delimiter", "type.member"],
            "@default": ""
          }
        }
      ],
      [/\d*\.\d+/, "number.float"],
      [/\d+/, "number"],
      [/[()\[\]]/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "operator",
            "@default": "delimiter"
          }
        }
      ],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/(\').*$/, "comment"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"C?/, "string", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/scala/scala.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/scala/scala.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/scala/scala.ts
var conf = {
  wordPattern: /(unary_[@~!#%^&*()\-=+\\|:<>\/?]+)|([a-zA-Z_$][\w$]*?_=)|(`[^`]+`)|([a-zA-Z_$][\w$]*)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*//\\s*(?:(?:#?region\\b)|(?:<editor-fold\\b))"),
      end: new RegExp("^\\s*//\\s*(?:(?:#?endregion\\b)|(?:</editor-fold>))")
    }
  }
};
var language = {
  tokenPostfix: ".scala",
  keywords: [
    "asInstanceOf",
    "catch",
    "class",
    "classOf",
    "def",
    "do",
    "else",
    "extends",
    "finally",
    "for",
    "foreach",
    "forSome",
    "if",
    "import",
    "isInstanceOf",
    "macro",
    "match",
    "new",
    "object",
    "package",
    "return",
    "throw",
    "trait",
    "try",
    "type",
    "until",
    "val",
    "var",
    "while",
    "with",
    "yield",
    "given",
    "enum",
    "then"
  ],
  softKeywords: ["as", "export", "extension", "end", "derives", "on"],
  constants: ["true", "false", "null", "this", "super"],
  modifiers: [
    "abstract",
    "final",
    "implicit",
    "lazy",
    "override",
    "private",
    "protected",
    "sealed"
  ],
  softModifiers: ["inline", "opaque", "open", "transparent", "using"],
  name: /(?:[a-z_$][\w$]*|`[^`]+`)/,
  type: /(?:[A-Z][\w$]*)/,
  symbols: /[=><!~?:&|+\-*\/^\\%@#]+/,
  digits: /\d+(_+\d+)*/,
  hexdigits: /[[0-9a-fA-F]+(_+[0-9a-fA-F]+)*/,
  escapes: /\\(?:[btnfr\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  fstring_conv: /[bBhHsScCdoxXeEfgGaAt]|[Tn](?:[HIklMSLNpzZsQ]|[BbhAaCYyjmde]|[RTrDFC])/,
  tokenizer: {
    root: [
      [/\braw"""/, { token: "string.quote", bracket: "@open", next: "@rawstringt" }],
      [/\braw"/, { token: "string.quote", bracket: "@open", next: "@rawstring" }],
      [/\bs"""/, { token: "string.quote", bracket: "@open", next: "@sstringt" }],
      [/\bs"/, { token: "string.quote", bracket: "@open", next: "@sstring" }],
      [/\bf""""/, { token: "string.quote", bracket: "@open", next: "@fstringt" }],
      [/\bf"/, { token: "string.quote", bracket: "@open", next: "@fstring" }],
      [/"""/, { token: "string.quote", bracket: "@open", next: "@stringt" }],
      [/"/, { token: "string.quote", bracket: "@open", next: "@string" }],
      [/(@digits)[eE]([\-+]?(@digits))?[fFdD]?/, "number.float", "@allowMethod"],
      [/(@digits)\.(@digits)([eE][\-+]?(@digits))?[fFdD]?/, "number.float", "@allowMethod"],
      [/0[xX](@hexdigits)[Ll]?/, "number.hex", "@allowMethod"],
      [/(@digits)[fFdD]/, "number.float", "@allowMethod"],
      [/(@digits)[lL]?/, "number", "@allowMethod"],
      [/\b_\*/, "key"],
      [/\b(_)\b/, "keyword", "@allowMethod"],
      [/\bimport\b/, "keyword", "@import"],
      [/\b(case)([ \t]+)(class)\b/, ["keyword.modifier", "white", "keyword"]],
      [/\bcase\b/, "keyword", "@case"],
      [/\bva[lr]\b/, "keyword", "@vardef"],
      [
        /\b(def)([ \t]+)((?:unary_)?@symbols|@name(?:_=)|@name)/,
        ["keyword", "white", "identifier"]
      ],
      [/@name(?=[ \t]*:(?!:))/, "variable"],
      [/(\.)(@name|@symbols)/, ["operator", { token: "@rematch", next: "@allowMethod" }]],
      [/([{(])(\s*)(@name(?=\s*=>))/, ["@brackets", "white", "variable"]],
      [
        /@name/,
        {
          cases: {
            "@keywords": "keyword",
            "@softKeywords": "keyword",
            "@modifiers": "keyword.modifier",
            "@softModifiers": "keyword.modifier",
            "@constants": {
              token: "constant",
              next: "@allowMethod"
            },
            "@default": {
              token: "identifier",
              next: "@allowMethod"
            }
          }
        }
      ],
      [/@type/, "type", "@allowMethod"],
      { include: "@whitespace" },
      [/@[a-zA-Z_$][\w$]*(?:\.[a-zA-Z_$][\w$]*)*/, "annotation"],
      [/[{(]/, "@brackets"],
      [/[})]/, "@brackets", "@allowMethod"],
      [/\[/, "operator.square"],
      [/](?!\s*(?:va[rl]|def|type)\b)/, "operator.square", "@allowMethod"],
      [/]/, "operator.square"],
      [/([=-]>|<-|>:|<:|:>|<%)(?=[\s\w()[\]{},\."'`])/, "keyword"],
      [/@symbols/, "operator"],
      [/[;,\.]/, "delimiter"],
      [/'[a-zA-Z$][\w$]*(?!')/, "attribute.name"],
      [/'[^\\']'/, "string", "@allowMethod"],
      [/(')(@escapes)(')/, ["string", "string.escape", { token: "string", next: "@allowMethod" }]],
      [/'/, "string.invalid"]
    ],
    import: [
      [/;/, "delimiter", "@pop"],
      [/^|$/, "", "@pop"],
      [/[ \t]+/, "white"],
      [/[\n\r]+/, "white", "@pop"],
      [/\/\*/, "comment", "@comment"],
      [/@name|@type/, "type"],
      [/[(){}]/, "@brackets"],
      [/[[\]]/, "operator.square"],
      [/[\.,]/, "delimiter"]
    ],
    allowMethod: [
      [/^|$/, "", "@pop"],
      [/[ \t]+/, "white"],
      [/[\n\r]+/, "white", "@pop"],
      [/\/\*/, "comment", "@comment"],
      [/(?==>[\s\w([{])/, "keyword", "@pop"],
      [
        /(@name|@symbols)(?=[ \t]*[[({"'`]|[ \t]+(?:[+-]?\.?\d|\w))/,
        {
          cases: {
            "@keywords": { token: "keyword", next: "@pop" },
            "->|<-|>:|<:|<%": { token: "keyword", next: "@pop" },
            "@default": { token: "@rematch", next: "@pop" }
          }
        }
      ],
      ["", "", "@pop"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\/\*/, "comment", "@push"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    case: [
      [/\b_\*/, "key"],
      [/\b(_|true|false|null|this|super)\b/, "keyword", "@allowMethod"],
      [/\bif\b|=>/, "keyword", "@pop"],
      [/`[^`]+`/, "identifier", "@allowMethod"],
      [/@name/, "variable", "@allowMethod"],
      [/:::?|\||@(?![a-z_$])/, "keyword"],
      { include: "@root" }
    ],
    vardef: [
      [/\b_\*/, "key"],
      [/\b(_|true|false|null|this|super)\b/, "keyword"],
      [/@name/, "variable"],
      [/:::?|\||@(?![a-z_$])/, "keyword"],
      [/=|:(?!:)/, "operator", "@pop"],
      [/$/, "white", "@pop"],
      { include: "@root" }
    ],
    string: [
      [/[^\\"\n\r]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [
        /"/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ]
    ],
    stringt: [
      [/[^\\"\n\r]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"(?=""")/, "string"],
      [
        /"""/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ],
      [/"/, "string"]
    ],
    fstring: [
      [/@escapes/, "string.escape"],
      [
        /"/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ],
      [/\$\$/, "string"],
      [/(\$)([a-z_]\w*)/, ["operator", "identifier"]],
      [/\$\{/, "operator", "@interp"],
      [/%%/, "string"],
      [
        /(%)([\-#+ 0,(])(\d+|\.\d+|\d+\.\d+)(@fstring_conv)/,
        ["metatag", "keyword.modifier", "number", "metatag"]
      ],
      [/(%)(\d+|\.\d+|\d+\.\d+)(@fstring_conv)/, ["metatag", "number", "metatag"]],
      [/(%)([\-#+ 0,(])(@fstring_conv)/, ["metatag", "keyword.modifier", "metatag"]],
      [/(%)(@fstring_conv)/, ["metatag", "metatag"]],
      [/./, "string"]
    ],
    fstringt: [
      [/@escapes/, "string.escape"],
      [/"(?=""")/, "string"],
      [
        /"""/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ],
      [/\$\$/, "string"],
      [/(\$)([a-z_]\w*)/, ["operator", "identifier"]],
      [/\$\{/, "operator", "@interp"],
      [/%%/, "string"],
      [
        /(%)([\-#+ 0,(])(\d+|\.\d+|\d+\.\d+)(@fstring_conv)/,
        ["metatag", "keyword.modifier", "number", "metatag"]
      ],
      [/(%)(\d+|\.\d+|\d+\.\d+)(@fstring_conv)/, ["metatag", "number", "metatag"]],
      [/(%)([\-#+ 0,(])(@fstring_conv)/, ["metatag", "keyword.modifier", "metatag"]],
      [/(%)(@fstring_conv)/, ["metatag", "metatag"]],
      [/./, "string"]
    ],
    sstring: [
      [/@escapes/, "string.escape"],
      [
        /"/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ],
      [/\$\$/, "string"],
      [/(\$)([a-z_]\w*)/, ["operator", "identifier"]],
      [/\$\{/, "operator", "@interp"],
      [/./, "string"]
    ],
    sstringt: [
      [/@escapes/, "string.escape"],
      [/"(?=""")/, "string"],
      [
        /"""/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ],
      [/\$\$/, "string"],
      [/(\$)([a-z_]\w*)/, ["operator", "identifier"]],
      [/\$\{/, "operator", "@interp"],
      [/./, "string"]
    ],
    interp: [[/{/, "operator", "@push"], [/}/, "operator", "@pop"], { include: "@root" }],
    rawstring: [
      [/[^"]/, "string"],
      [
        /"/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ]
    ],
    rawstringt: [
      [/[^"]/, "string"],
      [/"(?=""")/, "string"],
      [
        /"""/,
        {
          token: "string.quote",
          bracket: "@close",
          switchTo: "@allowMethod"
        }
      ],
      [/"/, "string"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/scheme/scheme.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/scheme/scheme.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/scheme/scheme.ts
var conf = {
  comments: {
    lineComment: ";",
    blockComment: ["#|", "|#"]
  },
  brackets: [
    ["(", ")"],
    ["{", "}"],
    ["[", "]"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' }
  ]
};
var language = {
  defaultToken: "",
  ignoreCase: true,
  tokenPostfix: ".scheme",
  brackets: [
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" }
  ],
  keywords: [
    "case",
    "do",
    "let",
    "loop",
    "if",
    "else",
    "when",
    "cons",
    "car",
    "cdr",
    "cond",
    "lambda",
    "lambda*",
    "syntax-rules",
    "format",
    "set!",
    "quote",
    "eval",
    "append",
    "list",
    "list?",
    "member?",
    "load"
  ],
  constants: ["#t", "#f"],
  operators: ["eq?", "eqv?", "equal?", "and", "or", "not", "null?"],
  tokenizer: {
    root: [
      [/#[xXoObB][0-9a-fA-F]+/, "number.hex"],
      [/[+-]?\d+(?:(?:\.\d*)?(?:[eE][+-]?\d+)?)?/, "number.float"],
      [
        /(?:\b(?:(define|define-syntax|define-macro))\b)(\s+)((?:\w|\-|\!|\?)*)/,
        ["keyword", "white", "variable"]
      ],
      { include: "@whitespace" },
      { include: "@strings" },
      [
        /[a-zA-Z_#][a-zA-Z0-9_\-\?\!\*]*/,
        {
          cases: {
            "@keywords": "keyword",
            "@constants": "constant",
            "@operators": "operators",
            "@default": "identifier"
          }
        }
      ]
    ],
    comment: [
      [/[^\|#]+/, "comment"],
      [/#\|/, "comment", "@push"],
      [/\|#/, "comment", "@pop"],
      [/[\|#]/, "comment"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/#\|/, "comment", "@comment"],
      [/;.*$/, "comment"]
    ],
    strings: [
      [/"$/, "string", "@popall"],
      [/"(?=.)/, "string", "@multiLineString"]
    ],
    multiLineString: [
      [/[^\\"]+$/, "string", "@popall"],
      [/[^\\"]+/, "string"],
      [/\\./, "string.escape"],
      [/"/, "string", "@popall"],
      [/\\$/, "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/scss/scss.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/scss/scss.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/scss/scss.ts
var conf = {
  wordPattern: /(#?-?\d*\.\d\w*%?)|([@$#!.:]?[\w-?]+%?)|[@#!.]/g,
  comments: {
    blockComment: ["/*", "*/"],
    lineComment: "//"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "'", close: "'", notIn: ["string", "comment"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*\\/\\*\\s*#region\\b\\s*(.*?)\\s*\\*\\/"),
      end: new RegExp("^\\s*\\/\\*\\s*#endregion\\b.*\\*\\/")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".scss",
  ws: "[ 	\n\r\f]*",
  identifier: "-?-?([a-zA-Z]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))([\\w\\-]|(\\\\(([0-9a-fA-F]{1,6}\\s?)|[^[0-9a-fA-F])))*",
  brackets: [
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.bracket" },
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "<", close: ">", token: "delimiter.angle" }
  ],
  tokenizer: {
    root: [{ include: "@selector" }],
    selector: [
      { include: "@comments" },
      { include: "@import" },
      { include: "@variabledeclaration" },
      { include: "@warndebug" },
      ["[@](include)", { token: "keyword", next: "@includedeclaration" }],
      [
        "[@](keyframes|-webkit-keyframes|-moz-keyframes|-o-keyframes)",
        { token: "keyword", next: "@keyframedeclaration" }
      ],
      ["[@](page|content|font-face|-moz-document)", { token: "keyword" }],
      ["[@](charset|namespace)", { token: "keyword", next: "@declarationbody" }],
      ["[@](function)", { token: "keyword", next: "@functiondeclaration" }],
      ["[@](mixin)", { token: "keyword", next: "@mixindeclaration" }],
      ["url(\\-prefix)?\\(", { token: "meta", next: "@urldeclaration" }],
      { include: "@controlstatement" },
      { include: "@selectorname" },
      ["[&\\*]", "tag"],
      ["[>\\+,]", "delimiter"],
      ["\\[", { token: "delimiter.bracket", next: "@selectorattribute" }],
      ["{", { token: "delimiter.curly", next: "@selectorbody" }]
    ],
    selectorbody: [
      ["[*_]?@identifier@ws:(?=(\\s|\\d|[^{;}]*[;}]))", "attribute.name", "@rulevalue"],
      { include: "@selector" },
      ["[@](extend)", { token: "keyword", next: "@extendbody" }],
      ["[@](return)", { token: "keyword", next: "@declarationbody" }],
      ["}", { token: "delimiter.curly", next: "@pop" }]
    ],
    selectorname: [
      ["#{", { token: "meta", next: "@variableinterpolation" }],
      ["(\\.|#(?=[^{])|%|(@identifier)|:)+", "tag"]
    ],
    selectorattribute: [{ include: "@term" }, ["]", { token: "delimiter.bracket", next: "@pop" }]],
    term: [
      { include: "@comments" },
      ["url(\\-prefix)?\\(", { token: "meta", next: "@urldeclaration" }],
      { include: "@functioninvocation" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@variablereference" },
      ["(and\\b|or\\b|not\\b)", "operator"],
      { include: "@name" },
      ["([<>=\\+\\-\\*\\/\\^\\|\\~,])", "operator"],
      [",", "delimiter"],
      ["!default", "literal"],
      ["\\(", { token: "delimiter.parenthesis", next: "@parenthizedterm" }]
    ],
    rulevalue: [
      { include: "@term" },
      ["!important", "literal"],
      [";", "delimiter", "@pop"],
      ["{", { token: "delimiter.curly", switchTo: "@nestedproperty" }],
      ["(?=})", { token: "", next: "@pop" }]
    ],
    nestedproperty: [
      ["[*_]?@identifier@ws:", "attribute.name", "@rulevalue"],
      { include: "@comments" },
      ["}", { token: "delimiter.curly", next: "@pop" }]
    ],
    warndebug: [["[@](warn|debug)", { token: "keyword", next: "@declarationbody" }]],
    import: [["[@](import)", { token: "keyword", next: "@declarationbody" }]],
    variabledeclaration: [
      ["\\$@identifier@ws:", "variable.decl", "@declarationbody"]
    ],
    urldeclaration: [
      { include: "@strings" },
      ["[^)\r\n]+", "string"],
      ["\\)", { token: "meta", next: "@pop" }]
    ],
    parenthizedterm: [
      { include: "@term" },
      ["\\)", { token: "delimiter.parenthesis", next: "@pop" }]
    ],
    declarationbody: [
      { include: "@term" },
      [";", "delimiter", "@pop"],
      ["(?=})", { token: "", next: "@pop" }]
    ],
    extendbody: [
      { include: "@selectorname" },
      ["!optional", "literal"],
      [";", "delimiter", "@pop"],
      ["(?=})", { token: "", next: "@pop" }]
    ],
    variablereference: [
      ["\\$@identifier", "variable.ref"],
      ["\\.\\.\\.", "operator"],
      ["#{", { token: "meta", next: "@variableinterpolation" }]
    ],
    variableinterpolation: [
      { include: "@variablereference" },
      ["}", { token: "meta", next: "@pop" }]
    ],
    comments: [
      ["\\/\\*", "comment", "@comment"],
      ["\\/\\/+.*", "comment"]
    ],
    comment: [
      ["\\*\\/", "comment", "@pop"],
      [".", "comment"]
    ],
    name: [["@identifier", "attribute.value"]],
    numbers: [
      ["(\\d*\\.)?\\d+([eE][\\-+]?\\d+)?", { token: "number", next: "@units" }],
      ["#[0-9a-fA-F_]+(?!\\w)", "number.hex"]
    ],
    units: [
      [
        "(em|ex|ch|rem|fr|vmin|vmax|vw|vh|vm|cm|mm|in|px|pt|pc|deg|grad|rad|turn|s|ms|Hz|kHz|%)?",
        "number",
        "@pop"
      ]
    ],
    functiondeclaration: [
      ["@identifier@ws\\(", { token: "meta", next: "@parameterdeclaration" }],
      ["{", { token: "delimiter.curly", switchTo: "@functionbody" }]
    ],
    mixindeclaration: [
      ["@identifier@ws\\(", { token: "meta", next: "@parameterdeclaration" }],
      ["@identifier", "meta"],
      ["{", { token: "delimiter.curly", switchTo: "@selectorbody" }]
    ],
    parameterdeclaration: [
      ["\\$@identifier@ws:", "variable.decl"],
      ["\\.\\.\\.", "operator"],
      [",", "delimiter"],
      { include: "@term" },
      ["\\)", { token: "meta", next: "@pop" }]
    ],
    includedeclaration: [
      { include: "@functioninvocation" },
      ["@identifier", "meta"],
      [";", "delimiter", "@pop"],
      ["(?=})", { token: "", next: "@pop" }],
      ["{", { token: "delimiter.curly", switchTo: "@selectorbody" }]
    ],
    keyframedeclaration: [
      ["@identifier", "meta"],
      ["{", { token: "delimiter.curly", switchTo: "@keyframebody" }]
    ],
    keyframebody: [
      { include: "@term" },
      ["{", { token: "delimiter.curly", next: "@selectorbody" }],
      ["}", { token: "delimiter.curly", next: "@pop" }]
    ],
    controlstatement: [
      [
        "[@](if|else|for|while|each|media)",
        { token: "keyword.flow", next: "@controlstatementdeclaration" }
      ]
    ],
    controlstatementdeclaration: [
      ["(in|from|through|if|to)\\b", { token: "keyword.flow" }],
      { include: "@term" },
      ["{", { token: "delimiter.curly", switchTo: "@selectorbody" }]
    ],
    functionbody: [
      ["[@](return)", { token: "keyword" }],
      { include: "@variabledeclaration" },
      { include: "@term" },
      { include: "@controlstatement" },
      [";", "delimiter"],
      ["}", { token: "delimiter.curly", next: "@pop" }]
    ],
    functioninvocation: [["@identifier\\(", { token: "meta", next: "@functionarguments" }]],
    functionarguments: [
      ["\\$@identifier@ws:", "attribute.name"],
      ["[,]", "delimiter"],
      { include: "@term" },
      ["\\)", { token: "meta", next: "@pop" }]
    ],
    strings: [
      ['~?"', { token: "string.delimiter", next: "@stringenddoublequote" }],
      ["~?'", { token: "string.delimiter", next: "@stringendquote" }]
    ],
    stringenddoublequote: [
      ["\\\\.", "string"],
      ['"', { token: "string.delimiter", next: "@pop" }],
      [".", "string"]
    ],
    stringendquote: [
      ["\\\\.", "string"],
      ["'", { token: "string.delimiter", next: "@pop" }],
      [".", "string"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/shell/shell.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/shell/shell.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/shell/shell.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ]
};
var language = {
  defaultToken: "",
  ignoreCase: true,
  tokenPostfix: ".shell",
  brackets: [
    { token: "delimiter.bracket", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" }
  ],
  keywords: [
    "if",
    "then",
    "do",
    "else",
    "elif",
    "while",
    "until",
    "for",
    "in",
    "esac",
    "fi",
    "fin",
    "fil",
    "done",
    "exit",
    "set",
    "unset",
    "export",
    "function"
  ],
  builtins: [
    "ab",
    "awk",
    "bash",
    "beep",
    "cat",
    "cc",
    "cd",
    "chown",
    "chmod",
    "chroot",
    "clear",
    "cp",
    "curl",
    "cut",
    "diff",
    "echo",
    "find",
    "gawk",
    "gcc",
    "get",
    "git",
    "grep",
    "hg",
    "kill",
    "killall",
    "ln",
    "ls",
    "make",
    "mkdir",
    "openssl",
    "mv",
    "nc",
    "node",
    "npm",
    "ping",
    "ps",
    "restart",
    "rm",
    "rmdir",
    "sed",
    "service",
    "sh",
    "shopt",
    "shred",
    "source",
    "sort",
    "sleep",
    "ssh",
    "start",
    "stop",
    "su",
    "sudo",
    "svn",
    "tee",
    "telnet",
    "top",
    "touch",
    "vi",
    "vim",
    "wall",
    "wc",
    "wget",
    "who",
    "write",
    "yes",
    "zsh"
  ],
  startingWithDash: /\-+\w+/,
  identifiersWithDashes: /[a-zA-Z]\w+(?:@startingWithDash)+/,
  symbols: /[=><!~?&|+\-*\/\^;\.,]+/,
  tokenizer: {
    root: [
      [/@identifiersWithDashes/, ""],
      [/(\s)((?:@startingWithDash)+)/, ["white", "attribute.name"]],
      [
        /[a-zA-Z]\w*/,
        {
          cases: {
            "@keywords": "keyword",
            "@builtins": "type.identifier",
            "@default": ""
          }
        }
      ],
      { include: "@whitespace" },
      { include: "@strings" },
      { include: "@parameters" },
      { include: "@heredoc" },
      [/[{}\[\]()]/, "@brackets"],
      [/@symbols/, "delimiter"],
      { include: "@numbers" },
      [/[,;]/, "delimiter"]
    ],
    whitespace: [
      [/\s+/, "white"],
      [/(^#!.*$)/, "metatag"],
      [/(^#.*$)/, "comment"]
    ],
    numbers: [
      [/\d*\.\d+([eE][\-+]?\d+)?/, "number.float"],
      [/0[xX][0-9a-fA-F_]*[0-9a-fA-F]/, "number.hex"],
      [/\d+/, "number"]
    ],
    strings: [
      [/'/, "string", "@stringBody"],
      [/"/, "string", "@dblStringBody"]
    ],
    stringBody: [
      [/'/, "string", "@popall"],
      [/./, "string"]
    ],
    dblStringBody: [
      [/"/, "string", "@popall"],
      [/./, "string"]
    ],
    heredoc: [
      [
        /(<<[-<]?)(\s*)(['"`]?)([\w\-]+)(['"`]?)/,
        [
          "constants",
          "white",
          "string.heredoc.delimiter",
          "string.heredoc",
          "string.heredoc.delimiter"
        ]
      ]
    ],
    parameters: [
      [/\$\d+/, "variable.predefined"],
      [/\$\w+/, "variable"],
      [/\$[*@#?\-$!0_]/, "variable"],
      [/\$'/, "variable", "@parameterBodyQuote"],
      [/\$"/, "variable", "@parameterBodyDoubleQuote"],
      [/\$\(/, "variable", "@parameterBodyParen"],
      [/\$\{/, "variable", "@parameterBodyCurlyBrace"]
    ],
    parameterBodyQuote: [
      [/[^#:%*@\-!_']+/, "variable"],
      [/[#:%*@\-!_]/, "delimiter"],
      [/[']/, "variable", "@pop"]
    ],
    parameterBodyDoubleQuote: [
      [/[^#:%*@\-!_"]+/, "variable"],
      [/[#:%*@\-!_]/, "delimiter"],
      [/["]/, "variable", "@pop"]
    ],
    parameterBodyParen: [
      [/[^#:%*@\-!_)]+/, "variable"],
      [/[#:%*@\-!_]/, "delimiter"],
      [/[)]/, "variable", "@pop"]
    ],
    parameterBodyCurlyBrace: [
      [/[^#:%*@\-!_}]+/, "variable"],
      [/[#:%*@\-!_]/, "delimiter"],
      [/[}]/, "variable", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/solidity/solidity.js":
/*!********************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/solidity/solidity.js ***!
  \********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/solidity/solidity.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sol",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" },
    { token: "delimiter.angle", open: "<", close: ">" }
  ],
  keywords: [
    "pragma",
    "solidity",
    "contract",
    "library",
    "using",
    "struct",
    "function",
    "modifier",
    "constructor",
    "address",
    "string",
    "bool",
    "Int",
    "Uint",
    "Byte",
    "Fixed",
    "Ufixed",
    "int",
    "int8",
    "int16",
    "int24",
    "int32",
    "int40",
    "int48",
    "int56",
    "int64",
    "int72",
    "int80",
    "int88",
    "int96",
    "int104",
    "int112",
    "int120",
    "int128",
    "int136",
    "int144",
    "int152",
    "int160",
    "int168",
    "int176",
    "int184",
    "int192",
    "int200",
    "int208",
    "int216",
    "int224",
    "int232",
    "int240",
    "int248",
    "int256",
    "uint",
    "uint8",
    "uint16",
    "uint24",
    "uint32",
    "uint40",
    "uint48",
    "uint56",
    "uint64",
    "uint72",
    "uint80",
    "uint88",
    "uint96",
    "uint104",
    "uint112",
    "uint120",
    "uint128",
    "uint136",
    "uint144",
    "uint152",
    "uint160",
    "uint168",
    "uint176",
    "uint184",
    "uint192",
    "uint200",
    "uint208",
    "uint216",
    "uint224",
    "uint232",
    "uint240",
    "uint248",
    "uint256",
    "byte",
    "bytes",
    "bytes1",
    "bytes2",
    "bytes3",
    "bytes4",
    "bytes5",
    "bytes6",
    "bytes7",
    "bytes8",
    "bytes9",
    "bytes10",
    "bytes11",
    "bytes12",
    "bytes13",
    "bytes14",
    "bytes15",
    "bytes16",
    "bytes17",
    "bytes18",
    "bytes19",
    "bytes20",
    "bytes21",
    "bytes22",
    "bytes23",
    "bytes24",
    "bytes25",
    "bytes26",
    "bytes27",
    "bytes28",
    "bytes29",
    "bytes30",
    "bytes31",
    "bytes32",
    "fixed",
    "fixed0x8",
    "fixed0x16",
    "fixed0x24",
    "fixed0x32",
    "fixed0x40",
    "fixed0x48",
    "fixed0x56",
    "fixed0x64",
    "fixed0x72",
    "fixed0x80",
    "fixed0x88",
    "fixed0x96",
    "fixed0x104",
    "fixed0x112",
    "fixed0x120",
    "fixed0x128",
    "fixed0x136",
    "fixed0x144",
    "fixed0x152",
    "fixed0x160",
    "fixed0x168",
    "fixed0x176",
    "fixed0x184",
    "fixed0x192",
    "fixed0x200",
    "fixed0x208",
    "fixed0x216",
    "fixed0x224",
    "fixed0x232",
    "fixed0x240",
    "fixed0x248",
    "fixed0x256",
    "fixed8x8",
    "fixed8x16",
    "fixed8x24",
    "fixed8x32",
    "fixed8x40",
    "fixed8x48",
    "fixed8x56",
    "fixed8x64",
    "fixed8x72",
    "fixed8x80",
    "fixed8x88",
    "fixed8x96",
    "fixed8x104",
    "fixed8x112",
    "fixed8x120",
    "fixed8x128",
    "fixed8x136",
    "fixed8x144",
    "fixed8x152",
    "fixed8x160",
    "fixed8x168",
    "fixed8x176",
    "fixed8x184",
    "fixed8x192",
    "fixed8x200",
    "fixed8x208",
    "fixed8x216",
    "fixed8x224",
    "fixed8x232",
    "fixed8x240",
    "fixed8x248",
    "fixed16x8",
    "fixed16x16",
    "fixed16x24",
    "fixed16x32",
    "fixed16x40",
    "fixed16x48",
    "fixed16x56",
    "fixed16x64",
    "fixed16x72",
    "fixed16x80",
    "fixed16x88",
    "fixed16x96",
    "fixed16x104",
    "fixed16x112",
    "fixed16x120",
    "fixed16x128",
    "fixed16x136",
    "fixed16x144",
    "fixed16x152",
    "fixed16x160",
    "fixed16x168",
    "fixed16x176",
    "fixed16x184",
    "fixed16x192",
    "fixed16x200",
    "fixed16x208",
    "fixed16x216",
    "fixed16x224",
    "fixed16x232",
    "fixed16x240",
    "fixed24x8",
    "fixed24x16",
    "fixed24x24",
    "fixed24x32",
    "fixed24x40",
    "fixed24x48",
    "fixed24x56",
    "fixed24x64",
    "fixed24x72",
    "fixed24x80",
    "fixed24x88",
    "fixed24x96",
    "fixed24x104",
    "fixed24x112",
    "fixed24x120",
    "fixed24x128",
    "fixed24x136",
    "fixed24x144",
    "fixed24x152",
    "fixed24x160",
    "fixed24x168",
    "fixed24x176",
    "fixed24x184",
    "fixed24x192",
    "fixed24x200",
    "fixed24x208",
    "fixed24x216",
    "fixed24x224",
    "fixed24x232",
    "fixed32x8",
    "fixed32x16",
    "fixed32x24",
    "fixed32x32",
    "fixed32x40",
    "fixed32x48",
    "fixed32x56",
    "fixed32x64",
    "fixed32x72",
    "fixed32x80",
    "fixed32x88",
    "fixed32x96",
    "fixed32x104",
    "fixed32x112",
    "fixed32x120",
    "fixed32x128",
    "fixed32x136",
    "fixed32x144",
    "fixed32x152",
    "fixed32x160",
    "fixed32x168",
    "fixed32x176",
    "fixed32x184",
    "fixed32x192",
    "fixed32x200",
    "fixed32x208",
    "fixed32x216",
    "fixed32x224",
    "fixed40x8",
    "fixed40x16",
    "fixed40x24",
    "fixed40x32",
    "fixed40x40",
    "fixed40x48",
    "fixed40x56",
    "fixed40x64",
    "fixed40x72",
    "fixed40x80",
    "fixed40x88",
    "fixed40x96",
    "fixed40x104",
    "fixed40x112",
    "fixed40x120",
    "fixed40x128",
    "fixed40x136",
    "fixed40x144",
    "fixed40x152",
    "fixed40x160",
    "fixed40x168",
    "fixed40x176",
    "fixed40x184",
    "fixed40x192",
    "fixed40x200",
    "fixed40x208",
    "fixed40x216",
    "fixed48x8",
    "fixed48x16",
    "fixed48x24",
    "fixed48x32",
    "fixed48x40",
    "fixed48x48",
    "fixed48x56",
    "fixed48x64",
    "fixed48x72",
    "fixed48x80",
    "fixed48x88",
    "fixed48x96",
    "fixed48x104",
    "fixed48x112",
    "fixed48x120",
    "fixed48x128",
    "fixed48x136",
    "fixed48x144",
    "fixed48x152",
    "fixed48x160",
    "fixed48x168",
    "fixed48x176",
    "fixed48x184",
    "fixed48x192",
    "fixed48x200",
    "fixed48x208",
    "fixed56x8",
    "fixed56x16",
    "fixed56x24",
    "fixed56x32",
    "fixed56x40",
    "fixed56x48",
    "fixed56x56",
    "fixed56x64",
    "fixed56x72",
    "fixed56x80",
    "fixed56x88",
    "fixed56x96",
    "fixed56x104",
    "fixed56x112",
    "fixed56x120",
    "fixed56x128",
    "fixed56x136",
    "fixed56x144",
    "fixed56x152",
    "fixed56x160",
    "fixed56x168",
    "fixed56x176",
    "fixed56x184",
    "fixed56x192",
    "fixed56x200",
    "fixed64x8",
    "fixed64x16",
    "fixed64x24",
    "fixed64x32",
    "fixed64x40",
    "fixed64x48",
    "fixed64x56",
    "fixed64x64",
    "fixed64x72",
    "fixed64x80",
    "fixed64x88",
    "fixed64x96",
    "fixed64x104",
    "fixed64x112",
    "fixed64x120",
    "fixed64x128",
    "fixed64x136",
    "fixed64x144",
    "fixed64x152",
    "fixed64x160",
    "fixed64x168",
    "fixed64x176",
    "fixed64x184",
    "fixed64x192",
    "fixed72x8",
    "fixed72x16",
    "fixed72x24",
    "fixed72x32",
    "fixed72x40",
    "fixed72x48",
    "fixed72x56",
    "fixed72x64",
    "fixed72x72",
    "fixed72x80",
    "fixed72x88",
    "fixed72x96",
    "fixed72x104",
    "fixed72x112",
    "fixed72x120",
    "fixed72x128",
    "fixed72x136",
    "fixed72x144",
    "fixed72x152",
    "fixed72x160",
    "fixed72x168",
    "fixed72x176",
    "fixed72x184",
    "fixed80x8",
    "fixed80x16",
    "fixed80x24",
    "fixed80x32",
    "fixed80x40",
    "fixed80x48",
    "fixed80x56",
    "fixed80x64",
    "fixed80x72",
    "fixed80x80",
    "fixed80x88",
    "fixed80x96",
    "fixed80x104",
    "fixed80x112",
    "fixed80x120",
    "fixed80x128",
    "fixed80x136",
    "fixed80x144",
    "fixed80x152",
    "fixed80x160",
    "fixed80x168",
    "fixed80x176",
    "fixed88x8",
    "fixed88x16",
    "fixed88x24",
    "fixed88x32",
    "fixed88x40",
    "fixed88x48",
    "fixed88x56",
    "fixed88x64",
    "fixed88x72",
    "fixed88x80",
    "fixed88x88",
    "fixed88x96",
    "fixed88x104",
    "fixed88x112",
    "fixed88x120",
    "fixed88x128",
    "fixed88x136",
    "fixed88x144",
    "fixed88x152",
    "fixed88x160",
    "fixed88x168",
    "fixed96x8",
    "fixed96x16",
    "fixed96x24",
    "fixed96x32",
    "fixed96x40",
    "fixed96x48",
    "fixed96x56",
    "fixed96x64",
    "fixed96x72",
    "fixed96x80",
    "fixed96x88",
    "fixed96x96",
    "fixed96x104",
    "fixed96x112",
    "fixed96x120",
    "fixed96x128",
    "fixed96x136",
    "fixed96x144",
    "fixed96x152",
    "fixed96x160",
    "fixed104x8",
    "fixed104x16",
    "fixed104x24",
    "fixed104x32",
    "fixed104x40",
    "fixed104x48",
    "fixed104x56",
    "fixed104x64",
    "fixed104x72",
    "fixed104x80",
    "fixed104x88",
    "fixed104x96",
    "fixed104x104",
    "fixed104x112",
    "fixed104x120",
    "fixed104x128",
    "fixed104x136",
    "fixed104x144",
    "fixed104x152",
    "fixed112x8",
    "fixed112x16",
    "fixed112x24",
    "fixed112x32",
    "fixed112x40",
    "fixed112x48",
    "fixed112x56",
    "fixed112x64",
    "fixed112x72",
    "fixed112x80",
    "fixed112x88",
    "fixed112x96",
    "fixed112x104",
    "fixed112x112",
    "fixed112x120",
    "fixed112x128",
    "fixed112x136",
    "fixed112x144",
    "fixed120x8",
    "fixed120x16",
    "fixed120x24",
    "fixed120x32",
    "fixed120x40",
    "fixed120x48",
    "fixed120x56",
    "fixed120x64",
    "fixed120x72",
    "fixed120x80",
    "fixed120x88",
    "fixed120x96",
    "fixed120x104",
    "fixed120x112",
    "fixed120x120",
    "fixed120x128",
    "fixed120x136",
    "fixed128x8",
    "fixed128x16",
    "fixed128x24",
    "fixed128x32",
    "fixed128x40",
    "fixed128x48",
    "fixed128x56",
    "fixed128x64",
    "fixed128x72",
    "fixed128x80",
    "fixed128x88",
    "fixed128x96",
    "fixed128x104",
    "fixed128x112",
    "fixed128x120",
    "fixed128x128",
    "fixed136x8",
    "fixed136x16",
    "fixed136x24",
    "fixed136x32",
    "fixed136x40",
    "fixed136x48",
    "fixed136x56",
    "fixed136x64",
    "fixed136x72",
    "fixed136x80",
    "fixed136x88",
    "fixed136x96",
    "fixed136x104",
    "fixed136x112",
    "fixed136x120",
    "fixed144x8",
    "fixed144x16",
    "fixed144x24",
    "fixed144x32",
    "fixed144x40",
    "fixed144x48",
    "fixed144x56",
    "fixed144x64",
    "fixed144x72",
    "fixed144x80",
    "fixed144x88",
    "fixed144x96",
    "fixed144x104",
    "fixed144x112",
    "fixed152x8",
    "fixed152x16",
    "fixed152x24",
    "fixed152x32",
    "fixed152x40",
    "fixed152x48",
    "fixed152x56",
    "fixed152x64",
    "fixed152x72",
    "fixed152x80",
    "fixed152x88",
    "fixed152x96",
    "fixed152x104",
    "fixed160x8",
    "fixed160x16",
    "fixed160x24",
    "fixed160x32",
    "fixed160x40",
    "fixed160x48",
    "fixed160x56",
    "fixed160x64",
    "fixed160x72",
    "fixed160x80",
    "fixed160x88",
    "fixed160x96",
    "fixed168x8",
    "fixed168x16",
    "fixed168x24",
    "fixed168x32",
    "fixed168x40",
    "fixed168x48",
    "fixed168x56",
    "fixed168x64",
    "fixed168x72",
    "fixed168x80",
    "fixed168x88",
    "fixed176x8",
    "fixed176x16",
    "fixed176x24",
    "fixed176x32",
    "fixed176x40",
    "fixed176x48",
    "fixed176x56",
    "fixed176x64",
    "fixed176x72",
    "fixed176x80",
    "fixed184x8",
    "fixed184x16",
    "fixed184x24",
    "fixed184x32",
    "fixed184x40",
    "fixed184x48",
    "fixed184x56",
    "fixed184x64",
    "fixed184x72",
    "fixed192x8",
    "fixed192x16",
    "fixed192x24",
    "fixed192x32",
    "fixed192x40",
    "fixed192x48",
    "fixed192x56",
    "fixed192x64",
    "fixed200x8",
    "fixed200x16",
    "fixed200x24",
    "fixed200x32",
    "fixed200x40",
    "fixed200x48",
    "fixed200x56",
    "fixed208x8",
    "fixed208x16",
    "fixed208x24",
    "fixed208x32",
    "fixed208x40",
    "fixed208x48",
    "fixed216x8",
    "fixed216x16",
    "fixed216x24",
    "fixed216x32",
    "fixed216x40",
    "fixed224x8",
    "fixed224x16",
    "fixed224x24",
    "fixed224x32",
    "fixed232x8",
    "fixed232x16",
    "fixed232x24",
    "fixed240x8",
    "fixed240x16",
    "fixed248x8",
    "ufixed",
    "ufixed0x8",
    "ufixed0x16",
    "ufixed0x24",
    "ufixed0x32",
    "ufixed0x40",
    "ufixed0x48",
    "ufixed0x56",
    "ufixed0x64",
    "ufixed0x72",
    "ufixed0x80",
    "ufixed0x88",
    "ufixed0x96",
    "ufixed0x104",
    "ufixed0x112",
    "ufixed0x120",
    "ufixed0x128",
    "ufixed0x136",
    "ufixed0x144",
    "ufixed0x152",
    "ufixed0x160",
    "ufixed0x168",
    "ufixed0x176",
    "ufixed0x184",
    "ufixed0x192",
    "ufixed0x200",
    "ufixed0x208",
    "ufixed0x216",
    "ufixed0x224",
    "ufixed0x232",
    "ufixed0x240",
    "ufixed0x248",
    "ufixed0x256",
    "ufixed8x8",
    "ufixed8x16",
    "ufixed8x24",
    "ufixed8x32",
    "ufixed8x40",
    "ufixed8x48",
    "ufixed8x56",
    "ufixed8x64",
    "ufixed8x72",
    "ufixed8x80",
    "ufixed8x88",
    "ufixed8x96",
    "ufixed8x104",
    "ufixed8x112",
    "ufixed8x120",
    "ufixed8x128",
    "ufixed8x136",
    "ufixed8x144",
    "ufixed8x152",
    "ufixed8x160",
    "ufixed8x168",
    "ufixed8x176",
    "ufixed8x184",
    "ufixed8x192",
    "ufixed8x200",
    "ufixed8x208",
    "ufixed8x216",
    "ufixed8x224",
    "ufixed8x232",
    "ufixed8x240",
    "ufixed8x248",
    "ufixed16x8",
    "ufixed16x16",
    "ufixed16x24",
    "ufixed16x32",
    "ufixed16x40",
    "ufixed16x48",
    "ufixed16x56",
    "ufixed16x64",
    "ufixed16x72",
    "ufixed16x80",
    "ufixed16x88",
    "ufixed16x96",
    "ufixed16x104",
    "ufixed16x112",
    "ufixed16x120",
    "ufixed16x128",
    "ufixed16x136",
    "ufixed16x144",
    "ufixed16x152",
    "ufixed16x160",
    "ufixed16x168",
    "ufixed16x176",
    "ufixed16x184",
    "ufixed16x192",
    "ufixed16x200",
    "ufixed16x208",
    "ufixed16x216",
    "ufixed16x224",
    "ufixed16x232",
    "ufixed16x240",
    "ufixed24x8",
    "ufixed24x16",
    "ufixed24x24",
    "ufixed24x32",
    "ufixed24x40",
    "ufixed24x48",
    "ufixed24x56",
    "ufixed24x64",
    "ufixed24x72",
    "ufixed24x80",
    "ufixed24x88",
    "ufixed24x96",
    "ufixed24x104",
    "ufixed24x112",
    "ufixed24x120",
    "ufixed24x128",
    "ufixed24x136",
    "ufixed24x144",
    "ufixed24x152",
    "ufixed24x160",
    "ufixed24x168",
    "ufixed24x176",
    "ufixed24x184",
    "ufixed24x192",
    "ufixed24x200",
    "ufixed24x208",
    "ufixed24x216",
    "ufixed24x224",
    "ufixed24x232",
    "ufixed32x8",
    "ufixed32x16",
    "ufixed32x24",
    "ufixed32x32",
    "ufixed32x40",
    "ufixed32x48",
    "ufixed32x56",
    "ufixed32x64",
    "ufixed32x72",
    "ufixed32x80",
    "ufixed32x88",
    "ufixed32x96",
    "ufixed32x104",
    "ufixed32x112",
    "ufixed32x120",
    "ufixed32x128",
    "ufixed32x136",
    "ufixed32x144",
    "ufixed32x152",
    "ufixed32x160",
    "ufixed32x168",
    "ufixed32x176",
    "ufixed32x184",
    "ufixed32x192",
    "ufixed32x200",
    "ufixed32x208",
    "ufixed32x216",
    "ufixed32x224",
    "ufixed40x8",
    "ufixed40x16",
    "ufixed40x24",
    "ufixed40x32",
    "ufixed40x40",
    "ufixed40x48",
    "ufixed40x56",
    "ufixed40x64",
    "ufixed40x72",
    "ufixed40x80",
    "ufixed40x88",
    "ufixed40x96",
    "ufixed40x104",
    "ufixed40x112",
    "ufixed40x120",
    "ufixed40x128",
    "ufixed40x136",
    "ufixed40x144",
    "ufixed40x152",
    "ufixed40x160",
    "ufixed40x168",
    "ufixed40x176",
    "ufixed40x184",
    "ufixed40x192",
    "ufixed40x200",
    "ufixed40x208",
    "ufixed40x216",
    "ufixed48x8",
    "ufixed48x16",
    "ufixed48x24",
    "ufixed48x32",
    "ufixed48x40",
    "ufixed48x48",
    "ufixed48x56",
    "ufixed48x64",
    "ufixed48x72",
    "ufixed48x80",
    "ufixed48x88",
    "ufixed48x96",
    "ufixed48x104",
    "ufixed48x112",
    "ufixed48x120",
    "ufixed48x128",
    "ufixed48x136",
    "ufixed48x144",
    "ufixed48x152",
    "ufixed48x160",
    "ufixed48x168",
    "ufixed48x176",
    "ufixed48x184",
    "ufixed48x192",
    "ufixed48x200",
    "ufixed48x208",
    "ufixed56x8",
    "ufixed56x16",
    "ufixed56x24",
    "ufixed56x32",
    "ufixed56x40",
    "ufixed56x48",
    "ufixed56x56",
    "ufixed56x64",
    "ufixed56x72",
    "ufixed56x80",
    "ufixed56x88",
    "ufixed56x96",
    "ufixed56x104",
    "ufixed56x112",
    "ufixed56x120",
    "ufixed56x128",
    "ufixed56x136",
    "ufixed56x144",
    "ufixed56x152",
    "ufixed56x160",
    "ufixed56x168",
    "ufixed56x176",
    "ufixed56x184",
    "ufixed56x192",
    "ufixed56x200",
    "ufixed64x8",
    "ufixed64x16",
    "ufixed64x24",
    "ufixed64x32",
    "ufixed64x40",
    "ufixed64x48",
    "ufixed64x56",
    "ufixed64x64",
    "ufixed64x72",
    "ufixed64x80",
    "ufixed64x88",
    "ufixed64x96",
    "ufixed64x104",
    "ufixed64x112",
    "ufixed64x120",
    "ufixed64x128",
    "ufixed64x136",
    "ufixed64x144",
    "ufixed64x152",
    "ufixed64x160",
    "ufixed64x168",
    "ufixed64x176",
    "ufixed64x184",
    "ufixed64x192",
    "ufixed72x8",
    "ufixed72x16",
    "ufixed72x24",
    "ufixed72x32",
    "ufixed72x40",
    "ufixed72x48",
    "ufixed72x56",
    "ufixed72x64",
    "ufixed72x72",
    "ufixed72x80",
    "ufixed72x88",
    "ufixed72x96",
    "ufixed72x104",
    "ufixed72x112",
    "ufixed72x120",
    "ufixed72x128",
    "ufixed72x136",
    "ufixed72x144",
    "ufixed72x152",
    "ufixed72x160",
    "ufixed72x168",
    "ufixed72x176",
    "ufixed72x184",
    "ufixed80x8",
    "ufixed80x16",
    "ufixed80x24",
    "ufixed80x32",
    "ufixed80x40",
    "ufixed80x48",
    "ufixed80x56",
    "ufixed80x64",
    "ufixed80x72",
    "ufixed80x80",
    "ufixed80x88",
    "ufixed80x96",
    "ufixed80x104",
    "ufixed80x112",
    "ufixed80x120",
    "ufixed80x128",
    "ufixed80x136",
    "ufixed80x144",
    "ufixed80x152",
    "ufixed80x160",
    "ufixed80x168",
    "ufixed80x176",
    "ufixed88x8",
    "ufixed88x16",
    "ufixed88x24",
    "ufixed88x32",
    "ufixed88x40",
    "ufixed88x48",
    "ufixed88x56",
    "ufixed88x64",
    "ufixed88x72",
    "ufixed88x80",
    "ufixed88x88",
    "ufixed88x96",
    "ufixed88x104",
    "ufixed88x112",
    "ufixed88x120",
    "ufixed88x128",
    "ufixed88x136",
    "ufixed88x144",
    "ufixed88x152",
    "ufixed88x160",
    "ufixed88x168",
    "ufixed96x8",
    "ufixed96x16",
    "ufixed96x24",
    "ufixed96x32",
    "ufixed96x40",
    "ufixed96x48",
    "ufixed96x56",
    "ufixed96x64",
    "ufixed96x72",
    "ufixed96x80",
    "ufixed96x88",
    "ufixed96x96",
    "ufixed96x104",
    "ufixed96x112",
    "ufixed96x120",
    "ufixed96x128",
    "ufixed96x136",
    "ufixed96x144",
    "ufixed96x152",
    "ufixed96x160",
    "ufixed104x8",
    "ufixed104x16",
    "ufixed104x24",
    "ufixed104x32",
    "ufixed104x40",
    "ufixed104x48",
    "ufixed104x56",
    "ufixed104x64",
    "ufixed104x72",
    "ufixed104x80",
    "ufixed104x88",
    "ufixed104x96",
    "ufixed104x104",
    "ufixed104x112",
    "ufixed104x120",
    "ufixed104x128",
    "ufixed104x136",
    "ufixed104x144",
    "ufixed104x152",
    "ufixed112x8",
    "ufixed112x16",
    "ufixed112x24",
    "ufixed112x32",
    "ufixed112x40",
    "ufixed112x48",
    "ufixed112x56",
    "ufixed112x64",
    "ufixed112x72",
    "ufixed112x80",
    "ufixed112x88",
    "ufixed112x96",
    "ufixed112x104",
    "ufixed112x112",
    "ufixed112x120",
    "ufixed112x128",
    "ufixed112x136",
    "ufixed112x144",
    "ufixed120x8",
    "ufixed120x16",
    "ufixed120x24",
    "ufixed120x32",
    "ufixed120x40",
    "ufixed120x48",
    "ufixed120x56",
    "ufixed120x64",
    "ufixed120x72",
    "ufixed120x80",
    "ufixed120x88",
    "ufixed120x96",
    "ufixed120x104",
    "ufixed120x112",
    "ufixed120x120",
    "ufixed120x128",
    "ufixed120x136",
    "ufixed128x8",
    "ufixed128x16",
    "ufixed128x24",
    "ufixed128x32",
    "ufixed128x40",
    "ufixed128x48",
    "ufixed128x56",
    "ufixed128x64",
    "ufixed128x72",
    "ufixed128x80",
    "ufixed128x88",
    "ufixed128x96",
    "ufixed128x104",
    "ufixed128x112",
    "ufixed128x120",
    "ufixed128x128",
    "ufixed136x8",
    "ufixed136x16",
    "ufixed136x24",
    "ufixed136x32",
    "ufixed136x40",
    "ufixed136x48",
    "ufixed136x56",
    "ufixed136x64",
    "ufixed136x72",
    "ufixed136x80",
    "ufixed136x88",
    "ufixed136x96",
    "ufixed136x104",
    "ufixed136x112",
    "ufixed136x120",
    "ufixed144x8",
    "ufixed144x16",
    "ufixed144x24",
    "ufixed144x32",
    "ufixed144x40",
    "ufixed144x48",
    "ufixed144x56",
    "ufixed144x64",
    "ufixed144x72",
    "ufixed144x80",
    "ufixed144x88",
    "ufixed144x96",
    "ufixed144x104",
    "ufixed144x112",
    "ufixed152x8",
    "ufixed152x16",
    "ufixed152x24",
    "ufixed152x32",
    "ufixed152x40",
    "ufixed152x48",
    "ufixed152x56",
    "ufixed152x64",
    "ufixed152x72",
    "ufixed152x80",
    "ufixed152x88",
    "ufixed152x96",
    "ufixed152x104",
    "ufixed160x8",
    "ufixed160x16",
    "ufixed160x24",
    "ufixed160x32",
    "ufixed160x40",
    "ufixed160x48",
    "ufixed160x56",
    "ufixed160x64",
    "ufixed160x72",
    "ufixed160x80",
    "ufixed160x88",
    "ufixed160x96",
    "ufixed168x8",
    "ufixed168x16",
    "ufixed168x24",
    "ufixed168x32",
    "ufixed168x40",
    "ufixed168x48",
    "ufixed168x56",
    "ufixed168x64",
    "ufixed168x72",
    "ufixed168x80",
    "ufixed168x88",
    "ufixed176x8",
    "ufixed176x16",
    "ufixed176x24",
    "ufixed176x32",
    "ufixed176x40",
    "ufixed176x48",
    "ufixed176x56",
    "ufixed176x64",
    "ufixed176x72",
    "ufixed176x80",
    "ufixed184x8",
    "ufixed184x16",
    "ufixed184x24",
    "ufixed184x32",
    "ufixed184x40",
    "ufixed184x48",
    "ufixed184x56",
    "ufixed184x64",
    "ufixed184x72",
    "ufixed192x8",
    "ufixed192x16",
    "ufixed192x24",
    "ufixed192x32",
    "ufixed192x40",
    "ufixed192x48",
    "ufixed192x56",
    "ufixed192x64",
    "ufixed200x8",
    "ufixed200x16",
    "ufixed200x24",
    "ufixed200x32",
    "ufixed200x40",
    "ufixed200x48",
    "ufixed200x56",
    "ufixed208x8",
    "ufixed208x16",
    "ufixed208x24",
    "ufixed208x32",
    "ufixed208x40",
    "ufixed208x48",
    "ufixed216x8",
    "ufixed216x16",
    "ufixed216x24",
    "ufixed216x32",
    "ufixed216x40",
    "ufixed224x8",
    "ufixed224x16",
    "ufixed224x24",
    "ufixed224x32",
    "ufixed232x8",
    "ufixed232x16",
    "ufixed232x24",
    "ufixed240x8",
    "ufixed240x16",
    "ufixed248x8",
    "event",
    "enum",
    "let",
    "mapping",
    "private",
    "public",
    "external",
    "inherited",
    "payable",
    "true",
    "false",
    "var",
    "import",
    "constant",
    "if",
    "else",
    "for",
    "else",
    "for",
    "while",
    "do",
    "break",
    "continue",
    "throw",
    "returns",
    "return",
    "suicide",
    "new",
    "is",
    "this",
    "super"
  ],
  operators: [
    "=",
    ">",
    "<",
    "!",
    "~",
    "?",
    ":",
    "==",
    "<=",
    ">=",
    "!=",
    "&&",
    "||",
    "++",
    "--",
    "+",
    "-",
    "*",
    "/",
    "&",
    "|",
    "^",
    "%",
    "<<",
    ">>",
    ">>>",
    "+=",
    "-=",
    "*=",
    "/=",
    "&=",
    "|=",
    "^=",
    "%=",
    "<<=",
    ">>=",
    ">>>="
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  integersuffix: /(ll|LL|u|U|l|L)?(ll|LL|u|U|l|L)?/,
  floatsuffix: /[fFlL]?/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/\[\[.*\]\]/, "annotation"],
      [/^\s*#\w+/, "keyword"],
      [/int\d*/, "keyword"],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\d+[eE]([\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/\d*\.\d+([eE][\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/0[xX][0-9a-fA-F']*[0-9a-fA-F](@integersuffix)/, "number.hex"],
      [/0[0-7']*[0-7](@integersuffix)/, "number.octal"],
      [/0[bB][0-1']*[0-1](@integersuffix)/, "number.binary"],
      [/\d[\d']*\d(@integersuffix)/, "number"],
      [/\d(@integersuffix)/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@doccomment"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    doccomment: [
      [/[^\/*]+/, "comment.doc"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/sophia/sophia.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/sophia/sophia.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/sophia/sophia.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".aes",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" },
    { token: "delimiter.angle", open: "<", close: ">" }
  ],
  keywords: [
    "contract",
    "library",
    "entrypoint",
    "function",
    "stateful",
    "state",
    "hash",
    "signature",
    "tuple",
    "list",
    "address",
    "string",
    "bool",
    "int",
    "record",
    "datatype",
    "type",
    "option",
    "oracle",
    "oracle_query",
    "Call",
    "Bits",
    "Bytes",
    "Oracle",
    "String",
    "Crypto",
    "Address",
    "Auth",
    "Chain",
    "None",
    "Some",
    "bits",
    "bytes",
    "event",
    "let",
    "map",
    "private",
    "public",
    "true",
    "false",
    "var",
    "if",
    "else",
    "throw"
  ],
  operators: [
    "=",
    ">",
    "<",
    "!",
    "~",
    "?",
    "::",
    ":",
    "==",
    "<=",
    ">=",
    "!=",
    "&&",
    "||",
    "++",
    "--",
    "+",
    "-",
    "*",
    "/",
    "&",
    "|",
    "^",
    "%",
    "<<",
    ">>",
    ">>>",
    "+=",
    "-=",
    "*=",
    "/=",
    "&=",
    "|=",
    "^=",
    "%=",
    "<<=",
    ">>=",
    ">>>="
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  integersuffix: /(ll|LL|u|U|l|L)?(ll|LL|u|U|l|L)?/,
  floatsuffix: /[fFlL]?/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/\[\[.*\]\]/, "annotation"],
      [/^\s*#\w+/, "keyword"],
      [/int\d*/, "keyword"],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      [/\d*\d+[eE]([\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/\d*\.\d+([eE][\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/0[xX][0-9a-fA-F']*[0-9a-fA-F](@integersuffix)/, "number.hex"],
      [/0[0-7']*[0-7](@integersuffix)/, "number.octal"],
      [/0[bB][0-1']*[0-1](@integersuffix)/, "number.binary"],
      [/\d[\d']*\d(@integersuffix)/, "number"],
      [/\d(@integersuffix)/, "number"],
      [/[;,.]/, "delimiter"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*\*(?!\/)/, "comment.doc", "@doccomment"],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    doccomment: [
      [/[^\/*]+/, "comment.doc"],
      [/\*\//, "comment.doc", "@pop"],
      [/[\/*]/, "comment.doc"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/sparql/sparql.js":
/*!****************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/sparql/sparql.js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/sparql/sparql.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "'", close: "'", notIn: ["string"] },
    { open: '"', close: '"', notIn: ["string"] },
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".rq",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" },
    { token: "delimiter.angle", open: "<", close: ">" }
  ],
  keywords: [
    "add",
    "as",
    "asc",
    "ask",
    "base",
    "by",
    "clear",
    "construct",
    "copy",
    "create",
    "data",
    "delete",
    "desc",
    "describe",
    "distinct",
    "drop",
    "false",
    "filter",
    "from",
    "graph",
    "group",
    "having",
    "in",
    "insert",
    "limit",
    "load",
    "minus",
    "move",
    "named",
    "not",
    "offset",
    "optional",
    "order",
    "prefix",
    "reduced",
    "select",
    "service",
    "silent",
    "to",
    "true",
    "undef",
    "union",
    "using",
    "values",
    "where",
    "with"
  ],
  builtinFunctions: [
    "a",
    "abs",
    "avg",
    "bind",
    "bnode",
    "bound",
    "ceil",
    "coalesce",
    "concat",
    "contains",
    "count",
    "datatype",
    "day",
    "encode_for_uri",
    "exists",
    "floor",
    "group_concat",
    "hours",
    "if",
    "iri",
    "isblank",
    "isiri",
    "isliteral",
    "isnumeric",
    "isuri",
    "lang",
    "langmatches",
    "lcase",
    "max",
    "md5",
    "min",
    "minutes",
    "month",
    "now",
    "rand",
    "regex",
    "replace",
    "round",
    "sameterm",
    "sample",
    "seconds",
    "sha1",
    "sha256",
    "sha384",
    "sha512",
    "str",
    "strafter",
    "strbefore",
    "strdt",
    "strends",
    "strlang",
    "strlen",
    "strstarts",
    "struuid",
    "substr",
    "sum",
    "timezone",
    "tz",
    "ucase",
    "uri",
    "uuid",
    "year"
  ],
  ignoreCase: true,
  tokenizer: {
    root: [
      [/<[^\s\u00a0>]*>?/, "tag"],
      { include: "@strings" },
      [/#.*/, "comment"],
      [/[{}()\[\]]/, "@brackets"],
      [/[;,.]/, "delimiter"],
      [/[_\w\d]+:(\.(?=[\w_\-\\%])|[:\w_-]|\\[-\\_~.!$&'()*+,;=/?#@%]|%[a-f\d][a-f\d])*/, "tag"],
      [/:(\.(?=[\w_\-\\%])|[:\w_-]|\\[-\\_~.!$&'()*+,;=/?#@%]|%[a-f\d][a-f\d])+/, "tag"],
      [
        /[$?]?[_\w\d]+/,
        {
          cases: {
            "@keywords": { token: "keyword" },
            "@builtinFunctions": { token: "predefined.sql" },
            "@default": "identifier"
          }
        }
      ],
      [/\^\^/, "operator.sql"],
      [/\^[*+\-<>=&|^\/!?]*/, "operator.sql"],
      [/[*+\-<>=&|\/!?]/, "operator.sql"],
      [/@[a-z\d\-]*/, "metatag.html"],
      [/\s+/, "white"]
    ],
    strings: [
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/'$/, "string.sql", "@pop"],
      [/'/, "string.sql", "@stringBody"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"$/, "string.sql", "@pop"],
      [/"/, "string.sql", "@dblStringBody"]
    ],
    stringBody: [
      [/[^\\']+/, "string.sql"],
      [/\\./, "string.escape"],
      [/'/, "string.sql", "@pop"]
    ],
    dblStringBody: [
      [/[^\\"]+/, "string.sql"],
      [/\\./, "string.escape"],
      [/"/, "string.sql", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/sql/sql.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/sql/sql.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/sql/sql.ts
var conf = {
  comments: {
    lineComment: "--",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sql",
  ignoreCase: true,
  brackets: [
    { open: "[", close: "]", token: "delimiter.square" },
    { open: "(", close: ")", token: "delimiter.parenthesis" }
  ],
  keywords: [
    "ABORT",
    "ABSOLUTE",
    "ACTION",
    "ADA",
    "ADD",
    "AFTER",
    "ALL",
    "ALLOCATE",
    "ALTER",
    "ALWAYS",
    "ANALYZE",
    "AND",
    "ANY",
    "ARE",
    "AS",
    "ASC",
    "ASSERTION",
    "AT",
    "ATTACH",
    "AUTHORIZATION",
    "AUTOINCREMENT",
    "AVG",
    "BACKUP",
    "BEFORE",
    "BEGIN",
    "BETWEEN",
    "BIT",
    "BIT_LENGTH",
    "BOTH",
    "BREAK",
    "BROWSE",
    "BULK",
    "BY",
    "CASCADE",
    "CASCADED",
    "CASE",
    "CAST",
    "CATALOG",
    "CHAR",
    "CHARACTER",
    "CHARACTER_LENGTH",
    "CHAR_LENGTH",
    "CHECK",
    "CHECKPOINT",
    "CLOSE",
    "CLUSTERED",
    "COALESCE",
    "COLLATE",
    "COLLATION",
    "COLUMN",
    "COMMIT",
    "COMPUTE",
    "CONFLICT",
    "CONNECT",
    "CONNECTION",
    "CONSTRAINT",
    "CONSTRAINTS",
    "CONTAINS",
    "CONTAINSTABLE",
    "CONTINUE",
    "CONVERT",
    "CORRESPONDING",
    "COUNT",
    "CREATE",
    "CROSS",
    "CURRENT",
    "CURRENT_DATE",
    "CURRENT_TIME",
    "CURRENT_TIMESTAMP",
    "CURRENT_USER",
    "CURSOR",
    "DATABASE",
    "DATE",
    "DAY",
    "DBCC",
    "DEALLOCATE",
    "DEC",
    "DECIMAL",
    "DECLARE",
    "DEFAULT",
    "DEFERRABLE",
    "DEFERRED",
    "DELETE",
    "DENY",
    "DESC",
    "DESCRIBE",
    "DESCRIPTOR",
    "DETACH",
    "DIAGNOSTICS",
    "DISCONNECT",
    "DISK",
    "DISTINCT",
    "DISTRIBUTED",
    "DO",
    "DOMAIN",
    "DOUBLE",
    "DROP",
    "DUMP",
    "EACH",
    "ELSE",
    "END",
    "END-EXEC",
    "ERRLVL",
    "ESCAPE",
    "EXCEPT",
    "EXCEPTION",
    "EXCLUDE",
    "EXCLUSIVE",
    "EXEC",
    "EXECUTE",
    "EXISTS",
    "EXIT",
    "EXPLAIN",
    "EXTERNAL",
    "EXTRACT",
    "FAIL",
    "FALSE",
    "FETCH",
    "FILE",
    "FILLFACTOR",
    "FILTER",
    "FIRST",
    "FLOAT",
    "FOLLOWING",
    "FOR",
    "FOREIGN",
    "FORTRAN",
    "FOUND",
    "FREETEXT",
    "FREETEXTTABLE",
    "FROM",
    "FULL",
    "FUNCTION",
    "GENERATED",
    "GET",
    "GLOB",
    "GLOBAL",
    "GO",
    "GOTO",
    "GRANT",
    "GROUP",
    "GROUPS",
    "HAVING",
    "HOLDLOCK",
    "HOUR",
    "IDENTITY",
    "IDENTITYCOL",
    "IDENTITY_INSERT",
    "IF",
    "IGNORE",
    "IMMEDIATE",
    "IN",
    "INCLUDE",
    "INDEX",
    "INDEXED",
    "INDICATOR",
    "INITIALLY",
    "INNER",
    "INPUT",
    "INSENSITIVE",
    "INSERT",
    "INSTEAD",
    "INT",
    "INTEGER",
    "INTERSECT",
    "INTERVAL",
    "INTO",
    "IS",
    "ISNULL",
    "ISOLATION",
    "JOIN",
    "KEY",
    "KILL",
    "LANGUAGE",
    "LAST",
    "LEADING",
    "LEFT",
    "LEVEL",
    "LIKE",
    "LIMIT",
    "LINENO",
    "LOAD",
    "LOCAL",
    "LOWER",
    "MATCH",
    "MATERIALIZED",
    "MAX",
    "MERGE",
    "MIN",
    "MINUTE",
    "MODULE",
    "MONTH",
    "NAMES",
    "NATIONAL",
    "NATURAL",
    "NCHAR",
    "NEXT",
    "NO",
    "NOCHECK",
    "NONCLUSTERED",
    "NONE",
    "NOT",
    "NOTHING",
    "NOTNULL",
    "NULL",
    "NULLIF",
    "NULLS",
    "NUMERIC",
    "OCTET_LENGTH",
    "OF",
    "OFF",
    "OFFSET",
    "OFFSETS",
    "ON",
    "ONLY",
    "OPEN",
    "OPENDATASOURCE",
    "OPENQUERY",
    "OPENROWSET",
    "OPENXML",
    "OPTION",
    "OR",
    "ORDER",
    "OTHERS",
    "OUTER",
    "OUTPUT",
    "OVER",
    "OVERLAPS",
    "PAD",
    "PARTIAL",
    "PARTITION",
    "PASCAL",
    "PERCENT",
    "PIVOT",
    "PLAN",
    "POSITION",
    "PRAGMA",
    "PRECEDING",
    "PRECISION",
    "PREPARE",
    "PRESERVE",
    "PRIMARY",
    "PRINT",
    "PRIOR",
    "PRIVILEGES",
    "PROC",
    "PROCEDURE",
    "PUBLIC",
    "QUERY",
    "RAISE",
    "RAISERROR",
    "RANGE",
    "READ",
    "READTEXT",
    "REAL",
    "RECONFIGURE",
    "RECURSIVE",
    "REFERENCES",
    "REGEXP",
    "REINDEX",
    "RELATIVE",
    "RELEASE",
    "RENAME",
    "REPLACE",
    "REPLICATION",
    "RESTORE",
    "RESTRICT",
    "RETURN",
    "RETURNING",
    "REVERT",
    "REVOKE",
    "RIGHT",
    "ROLLBACK",
    "ROW",
    "ROWCOUNT",
    "ROWGUIDCOL",
    "ROWS",
    "RULE",
    "SAVE",
    "SAVEPOINT",
    "SCHEMA",
    "SCROLL",
    "SECOND",
    "SECTION",
    "SECURITYAUDIT",
    "SELECT",
    "SEMANTICKEYPHRASETABLE",
    "SEMANTICSIMILARITYDETAILSTABLE",
    "SEMANTICSIMILARITYTABLE",
    "SESSION",
    "SESSION_USER",
    "SET",
    "SETUSER",
    "SHUTDOWN",
    "SIZE",
    "SMALLINT",
    "SOME",
    "SPACE",
    "SQL",
    "SQLCA",
    "SQLCODE",
    "SQLERROR",
    "SQLSTATE",
    "SQLWARNING",
    "STATISTICS",
    "SUBSTRING",
    "SUM",
    "SYSTEM_USER",
    "TABLE",
    "TABLESAMPLE",
    "TEMP",
    "TEMPORARY",
    "TEXTSIZE",
    "THEN",
    "TIES",
    "TIME",
    "TIMESTAMP",
    "TIMEZONE_HOUR",
    "TIMEZONE_MINUTE",
    "TO",
    "TOP",
    "TRAILING",
    "TRAN",
    "TRANSACTION",
    "TRANSLATE",
    "TRANSLATION",
    "TRIGGER",
    "TRIM",
    "TRUE",
    "TRUNCATE",
    "TRY_CONVERT",
    "TSEQUAL",
    "UNBOUNDED",
    "UNION",
    "UNIQUE",
    "UNKNOWN",
    "UNPIVOT",
    "UPDATE",
    "UPDATETEXT",
    "UPPER",
    "USAGE",
    "USE",
    "USER",
    "USING",
    "VACUUM",
    "VALUE",
    "VALUES",
    "VARCHAR",
    "VARYING",
    "VIEW",
    "VIRTUAL",
    "WAITFOR",
    "WHEN",
    "WHENEVER",
    "WHERE",
    "WHILE",
    "WINDOW",
    "WITH",
    "WITHIN GROUP",
    "WITHOUT",
    "WORK",
    "WRITE",
    "WRITETEXT",
    "YEAR",
    "ZONE"
  ],
  operators: [
    "ALL",
    "AND",
    "ANY",
    "BETWEEN",
    "EXISTS",
    "IN",
    "LIKE",
    "NOT",
    "OR",
    "SOME",
    "EXCEPT",
    "INTERSECT",
    "UNION",
    "APPLY",
    "CROSS",
    "FULL",
    "INNER",
    "JOIN",
    "LEFT",
    "OUTER",
    "RIGHT",
    "CONTAINS",
    "FREETEXT",
    "IS",
    "NULL",
    "PIVOT",
    "UNPIVOT",
    "MATCHED"
  ],
  builtinFunctions: [
    "AVG",
    "CHECKSUM_AGG",
    "COUNT",
    "COUNT_BIG",
    "GROUPING",
    "GROUPING_ID",
    "MAX",
    "MIN",
    "SUM",
    "STDEV",
    "STDEVP",
    "VAR",
    "VARP",
    "CUME_DIST",
    "FIRST_VALUE",
    "LAG",
    "LAST_VALUE",
    "LEAD",
    "PERCENTILE_CONT",
    "PERCENTILE_DISC",
    "PERCENT_RANK",
    "COLLATE",
    "COLLATIONPROPERTY",
    "TERTIARY_WEIGHTS",
    "FEDERATION_FILTERING_VALUE",
    "CAST",
    "CONVERT",
    "PARSE",
    "TRY_CAST",
    "TRY_CONVERT",
    "TRY_PARSE",
    "ASYMKEY_ID",
    "ASYMKEYPROPERTY",
    "CERTPROPERTY",
    "CERT_ID",
    "CRYPT_GEN_RANDOM",
    "DECRYPTBYASYMKEY",
    "DECRYPTBYCERT",
    "DECRYPTBYKEY",
    "DECRYPTBYKEYAUTOASYMKEY",
    "DECRYPTBYKEYAUTOCERT",
    "DECRYPTBYPASSPHRASE",
    "ENCRYPTBYASYMKEY",
    "ENCRYPTBYCERT",
    "ENCRYPTBYKEY",
    "ENCRYPTBYPASSPHRASE",
    "HASHBYTES",
    "IS_OBJECTSIGNED",
    "KEY_GUID",
    "KEY_ID",
    "KEY_NAME",
    "SIGNBYASYMKEY",
    "SIGNBYCERT",
    "SYMKEYPROPERTY",
    "VERIFYSIGNEDBYCERT",
    "VERIFYSIGNEDBYASYMKEY",
    "CURSOR_STATUS",
    "DATALENGTH",
    "IDENT_CURRENT",
    "IDENT_INCR",
    "IDENT_SEED",
    "IDENTITY",
    "SQL_VARIANT_PROPERTY",
    "CURRENT_TIMESTAMP",
    "DATEADD",
    "DATEDIFF",
    "DATEFROMPARTS",
    "DATENAME",
    "DATEPART",
    "DATETIME2FROMPARTS",
    "DATETIMEFROMPARTS",
    "DATETIMEOFFSETFROMPARTS",
    "DAY",
    "EOMONTH",
    "GETDATE",
    "GETUTCDATE",
    "ISDATE",
    "MONTH",
    "SMALLDATETIMEFROMPARTS",
    "SWITCHOFFSET",
    "SYSDATETIME",
    "SYSDATETIMEOFFSET",
    "SYSUTCDATETIME",
    "TIMEFROMPARTS",
    "TODATETIMEOFFSET",
    "YEAR",
    "CHOOSE",
    "COALESCE",
    "IIF",
    "NULLIF",
    "ABS",
    "ACOS",
    "ASIN",
    "ATAN",
    "ATN2",
    "CEILING",
    "COS",
    "COT",
    "DEGREES",
    "EXP",
    "FLOOR",
    "LOG",
    "LOG10",
    "PI",
    "POWER",
    "RADIANS",
    "RAND",
    "ROUND",
    "SIGN",
    "SIN",
    "SQRT",
    "SQUARE",
    "TAN",
    "APP_NAME",
    "APPLOCK_MODE",
    "APPLOCK_TEST",
    "ASSEMBLYPROPERTY",
    "COL_LENGTH",
    "COL_NAME",
    "COLUMNPROPERTY",
    "DATABASE_PRINCIPAL_ID",
    "DATABASEPROPERTYEX",
    "DB_ID",
    "DB_NAME",
    "FILE_ID",
    "FILE_IDEX",
    "FILE_NAME",
    "FILEGROUP_ID",
    "FILEGROUP_NAME",
    "FILEGROUPPROPERTY",
    "FILEPROPERTY",
    "FULLTEXTCATALOGPROPERTY",
    "FULLTEXTSERVICEPROPERTY",
    "INDEX_COL",
    "INDEXKEY_PROPERTY",
    "INDEXPROPERTY",
    "OBJECT_DEFINITION",
    "OBJECT_ID",
    "OBJECT_NAME",
    "OBJECT_SCHEMA_NAME",
    "OBJECTPROPERTY",
    "OBJECTPROPERTYEX",
    "ORIGINAL_DB_NAME",
    "PARSENAME",
    "SCHEMA_ID",
    "SCHEMA_NAME",
    "SCOPE_IDENTITY",
    "SERVERPROPERTY",
    "STATS_DATE",
    "TYPE_ID",
    "TYPE_NAME",
    "TYPEPROPERTY",
    "DENSE_RANK",
    "NTILE",
    "RANK",
    "ROW_NUMBER",
    "PUBLISHINGSERVERNAME",
    "OPENDATASOURCE",
    "OPENQUERY",
    "OPENROWSET",
    "OPENXML",
    "CERTENCODED",
    "CERTPRIVATEKEY",
    "CURRENT_USER",
    "HAS_DBACCESS",
    "HAS_PERMS_BY_NAME",
    "IS_MEMBER",
    "IS_ROLEMEMBER",
    "IS_SRVROLEMEMBER",
    "LOGINPROPERTY",
    "ORIGINAL_LOGIN",
    "PERMISSIONS",
    "PWDENCRYPT",
    "PWDCOMPARE",
    "SESSION_USER",
    "SESSIONPROPERTY",
    "SUSER_ID",
    "SUSER_NAME",
    "SUSER_SID",
    "SUSER_SNAME",
    "SYSTEM_USER",
    "USER",
    "USER_ID",
    "USER_NAME",
    "ASCII",
    "CHAR",
    "CHARINDEX",
    "CONCAT",
    "DIFFERENCE",
    "FORMAT",
    "LEFT",
    "LEN",
    "LOWER",
    "LTRIM",
    "NCHAR",
    "PATINDEX",
    "QUOTENAME",
    "REPLACE",
    "REPLICATE",
    "REVERSE",
    "RIGHT",
    "RTRIM",
    "SOUNDEX",
    "SPACE",
    "STR",
    "STUFF",
    "SUBSTRING",
    "UNICODE",
    "UPPER",
    "BINARY_CHECKSUM",
    "CHECKSUM",
    "CONNECTIONPROPERTY",
    "CONTEXT_INFO",
    "CURRENT_REQUEST_ID",
    "ERROR_LINE",
    "ERROR_NUMBER",
    "ERROR_MESSAGE",
    "ERROR_PROCEDURE",
    "ERROR_SEVERITY",
    "ERROR_STATE",
    "FORMATMESSAGE",
    "GETANSINULL",
    "GET_FILESTREAM_TRANSACTION_CONTEXT",
    "HOST_ID",
    "HOST_NAME",
    "ISNULL",
    "ISNUMERIC",
    "MIN_ACTIVE_ROWVERSION",
    "NEWID",
    "NEWSEQUENTIALID",
    "ROWCOUNT_BIG",
    "XACT_STATE",
    "TEXTPTR",
    "TEXTVALID",
    "COLUMNS_UPDATED",
    "EVENTDATA",
    "TRIGGER_NESTLEVEL",
    "UPDATE",
    "CHANGETABLE",
    "CHANGE_TRACKING_CONTEXT",
    "CHANGE_TRACKING_CURRENT_VERSION",
    "CHANGE_TRACKING_IS_COLUMN_IN_MASK",
    "CHANGE_TRACKING_MIN_VALID_VERSION",
    "CONTAINSTABLE",
    "FREETEXTTABLE",
    "SEMANTICKEYPHRASETABLE",
    "SEMANTICSIMILARITYDETAILSTABLE",
    "SEMANTICSIMILARITYTABLE",
    "FILETABLEROOTPATH",
    "GETFILENAMESPACEPATH",
    "GETPATHLOCATOR",
    "PATHNAME",
    "GET_TRANSMISSION_STATUS"
  ],
  builtinVariables: [
    "@@DATEFIRST",
    "@@DBTS",
    "@@LANGID",
    "@@LANGUAGE",
    "@@LOCK_TIMEOUT",
    "@@MAX_CONNECTIONS",
    "@@MAX_PRECISION",
    "@@NESTLEVEL",
    "@@OPTIONS",
    "@@REMSERVER",
    "@@SERVERNAME",
    "@@SERVICENAME",
    "@@SPID",
    "@@TEXTSIZE",
    "@@VERSION",
    "@@CURSOR_ROWS",
    "@@FETCH_STATUS",
    "@@DATEFIRST",
    "@@PROCID",
    "@@ERROR",
    "@@IDENTITY",
    "@@ROWCOUNT",
    "@@TRANCOUNT",
    "@@CONNECTIONS",
    "@@CPU_BUSY",
    "@@IDLE",
    "@@IO_BUSY",
    "@@PACKET_ERRORS",
    "@@PACK_RECEIVED",
    "@@PACK_SENT",
    "@@TIMETICKS",
    "@@TOTAL_ERRORS",
    "@@TOTAL_READ",
    "@@TOTAL_WRITE"
  ],
  pseudoColumns: ["$ACTION", "$IDENTITY", "$ROWGUID", "$PARTITION"],
  tokenizer: {
    root: [
      { include: "@comments" },
      { include: "@whitespace" },
      { include: "@pseudoColumns" },
      { include: "@numbers" },
      { include: "@strings" },
      { include: "@complexIdentifiers" },
      { include: "@scopes" },
      [/[;,.]/, "delimiter"],
      [/[()]/, "@brackets"],
      [
        /[\w@#$]+/,
        {
          cases: {
            "@operators": "operator",
            "@builtinVariables": "predefined",
            "@builtinFunctions": "predefined",
            "@keywords": "keyword",
            "@default": "identifier"
          }
        }
      ],
      [/[<>=!%&+\-*/|~^]/, "operator"]
    ],
    whitespace: [[/\s+/, "white"]],
    comments: [
      [/--+.*/, "comment"],
      [/\/\*/, { token: "comment.quote", next: "@comment" }]
    ],
    comment: [
      [/[^*/]+/, "comment"],
      [/\*\//, { token: "comment.quote", next: "@pop" }],
      [/./, "comment"]
    ],
    pseudoColumns: [
      [
        /[$][A-Za-z_][\w@#$]*/,
        {
          cases: {
            "@pseudoColumns": "predefined",
            "@default": "identifier"
          }
        }
      ]
    ],
    numbers: [
      [/0[xX][0-9a-fA-F]*/, "number"],
      [/[$][+-]*\d*(\.\d*)?/, "number"],
      [/((\d+(\.\d*)?)|(\.\d+))([eE][\-+]?\d+)?/, "number"]
    ],
    strings: [
      [/N'/, { token: "string", next: "@string" }],
      [/'/, { token: "string", next: "@string" }]
    ],
    string: [
      [/[^']+/, "string"],
      [/''/, "string"],
      [/'/, { token: "string", next: "@pop" }]
    ],
    complexIdentifiers: [
      [/\[/, { token: "identifier.quote", next: "@bracketedIdentifier" }],
      [/"/, { token: "identifier.quote", next: "@quotedIdentifier" }]
    ],
    bracketedIdentifier: [
      [/[^\]]+/, "identifier"],
      [/]]/, "identifier"],
      [/]/, { token: "identifier.quote", next: "@pop" }]
    ],
    quotedIdentifier: [
      [/[^"]+/, "identifier"],
      [/""/, "identifier"],
      [/"/, { token: "identifier.quote", next: "@pop" }]
    ],
    scopes: [
      [/BEGIN\s+(DISTRIBUTED\s+)?TRAN(SACTION)?\b/i, "keyword"],
      [/BEGIN\s+TRY\b/i, { token: "keyword.try" }],
      [/END\s+TRY\b/i, { token: "keyword.try" }],
      [/BEGIN\s+CATCH\b/i, { token: "keyword.catch" }],
      [/END\s+CATCH\b/i, { token: "keyword.catch" }],
      [/(BEGIN|CASE)\b/i, { token: "keyword.block" }],
      [/END\b/i, { token: "keyword.block" }],
      [/WHEN\b/i, { token: "keyword.choice" }],
      [/THEN\b/i, { token: "keyword.choice" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/st/st.js":
/*!********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/st/st.js ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/st/st.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["(*", "*)"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["var", "end_var"],
    ["var_input", "end_var"],
    ["var_output", "end_var"],
    ["var_in_out", "end_var"],
    ["var_temp", "end_var"],
    ["var_global", "end_var"],
    ["var_access", "end_var"],
    ["var_external", "end_var"],
    ["type", "end_type"],
    ["struct", "end_struct"],
    ["program", "end_program"],
    ["function", "end_function"],
    ["function_block", "end_function_block"],
    ["action", "end_action"],
    ["step", "end_step"],
    ["initial_step", "end_step"],
    ["transaction", "end_transaction"],
    ["configuration", "end_configuration"],
    ["tcp", "end_tcp"],
    ["recource", "end_recource"],
    ["channel", "end_channel"],
    ["library", "end_library"],
    ["folder", "end_folder"],
    ["binaries", "end_binaries"],
    ["includes", "end_includes"],
    ["sources", "end_sources"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "(", close: ")" },
    { open: "/*", close: "*/" },
    { open: "'", close: "'", notIn: ["string_sq"] },
    { open: '"', close: '"', notIn: ["string_dq"] },
    { open: "var_input", close: "end_var" },
    { open: "var_output", close: "end_var" },
    { open: "var_in_out", close: "end_var" },
    { open: "var_temp", close: "end_var" },
    { open: "var_global", close: "end_var" },
    { open: "var_access", close: "end_var" },
    { open: "var_external", close: "end_var" },
    { open: "type", close: "end_type" },
    { open: "struct", close: "end_struct" },
    { open: "program", close: "end_program" },
    { open: "function", close: "end_function" },
    { open: "function_block", close: "end_function_block" },
    { open: "action", close: "end_action" },
    { open: "step", close: "end_step" },
    { open: "initial_step", close: "end_step" },
    { open: "transaction", close: "end_transaction" },
    { open: "configuration", close: "end_configuration" },
    { open: "tcp", close: "end_tcp" },
    { open: "recource", close: "end_recource" },
    { open: "channel", close: "end_channel" },
    { open: "library", close: "end_library" },
    { open: "folder", close: "end_folder" },
    { open: "binaries", close: "end_binaries" },
    { open: "includes", close: "end_includes" },
    { open: "sources", close: "end_sources" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "var", close: "end_var" },
    { open: "var_input", close: "end_var" },
    { open: "var_output", close: "end_var" },
    { open: "var_in_out", close: "end_var" },
    { open: "var_temp", close: "end_var" },
    { open: "var_global", close: "end_var" },
    { open: "var_access", close: "end_var" },
    { open: "var_external", close: "end_var" },
    { open: "type", close: "end_type" },
    { open: "struct", close: "end_struct" },
    { open: "program", close: "end_program" },
    { open: "function", close: "end_function" },
    { open: "function_block", close: "end_function_block" },
    { open: "action", close: "end_action" },
    { open: "step", close: "end_step" },
    { open: "initial_step", close: "end_step" },
    { open: "transaction", close: "end_transaction" },
    { open: "configuration", close: "end_configuration" },
    { open: "tcp", close: "end_tcp" },
    { open: "recource", close: "end_recource" },
    { open: "channel", close: "end_channel" },
    { open: "library", close: "end_library" },
    { open: "folder", close: "end_folder" },
    { open: "binaries", close: "end_binaries" },
    { open: "includes", close: "end_includes" },
    { open: "sources", close: "end_sources" }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#pragma\\s+region\\b"),
      end: new RegExp("^\\s*#pragma\\s+endregion\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".st",
  ignoreCase: true,
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" }
  ],
  keywords: [
    "if",
    "end_if",
    "elsif",
    "else",
    "case",
    "of",
    "to",
    "__try",
    "__catch",
    "__finally",
    "do",
    "with",
    "by",
    "while",
    "repeat",
    "end_while",
    "end_repeat",
    "end_case",
    "for",
    "end_for",
    "task",
    "retain",
    "non_retain",
    "constant",
    "with",
    "at",
    "exit",
    "return",
    "interval",
    "priority",
    "address",
    "port",
    "on_channel",
    "then",
    "iec",
    "file",
    "uses",
    "version",
    "packagetype",
    "displayname",
    "copyright",
    "summary",
    "vendor",
    "common_source",
    "from",
    "extends",
    "implements"
  ],
  constant: ["false", "true", "null"],
  defineKeywords: [
    "var",
    "var_input",
    "var_output",
    "var_in_out",
    "var_temp",
    "var_global",
    "var_access",
    "var_external",
    "end_var",
    "type",
    "end_type",
    "struct",
    "end_struct",
    "program",
    "end_program",
    "function",
    "end_function",
    "function_block",
    "end_function_block",
    "interface",
    "end_interface",
    "method",
    "end_method",
    "property",
    "end_property",
    "namespace",
    "end_namespace",
    "configuration",
    "end_configuration",
    "tcp",
    "end_tcp",
    "resource",
    "end_resource",
    "channel",
    "end_channel",
    "library",
    "end_library",
    "folder",
    "end_folder",
    "binaries",
    "end_binaries",
    "includes",
    "end_includes",
    "sources",
    "end_sources",
    "action",
    "end_action",
    "step",
    "initial_step",
    "end_step",
    "transaction",
    "end_transaction"
  ],
  typeKeywords: [
    "int",
    "sint",
    "dint",
    "lint",
    "usint",
    "uint",
    "udint",
    "ulint",
    "real",
    "lreal",
    "time",
    "date",
    "time_of_day",
    "date_and_time",
    "string",
    "bool",
    "byte",
    "word",
    "dword",
    "array",
    "pointer",
    "lword"
  ],
  operators: [
    "=",
    ">",
    "<",
    ":",
    ":=",
    "<=",
    ">=",
    "<>",
    "&",
    "+",
    "-",
    "*",
    "**",
    "MOD",
    "^",
    "or",
    "and",
    "not",
    "xor",
    "abs",
    "acos",
    "asin",
    "atan",
    "cos",
    "exp",
    "expt",
    "ln",
    "log",
    "sin",
    "sqrt",
    "tan",
    "sel",
    "max",
    "min",
    "limit",
    "mux",
    "shl",
    "shr",
    "rol",
    "ror",
    "indexof",
    "sizeof",
    "adr",
    "adrinst",
    "bitadr",
    "is_valid",
    "ref",
    "ref_to"
  ],
  builtinVariables: [],
  builtinFunctions: [
    "sr",
    "rs",
    "tp",
    "ton",
    "tof",
    "eq",
    "ge",
    "le",
    "lt",
    "ne",
    "round",
    "trunc",
    "ctd",
    "\u0441tu",
    "ctud",
    "r_trig",
    "f_trig",
    "move",
    "concat",
    "delete",
    "find",
    "insert",
    "left",
    "len",
    "replace",
    "right",
    "rtc"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      [/(\.\.)/, "delimiter"],
      [/\b(16#[0-9A-Fa-f\_]*)+\b/, "number.hex"],
      [/\b(2#[01\_]+)+\b/, "number.binary"],
      [/\b(8#[0-9\_]*)+\b/, "number.octal"],
      [/\b\d*\.\d+([eE][\-+]?\d+)?\b/, "number.float"],
      [/\b(L?REAL)#[0-9\_\.e]+\b/, "number.float"],
      [/\b(BYTE|(?:D|L)?WORD|U?(?:S|D|L)?INT)#[0-9\_]+\b/, "number"],
      [/\d+/, "number"],
      [/\b(T|DT|TOD)#[0-9:-_shmyd]+\b/, "tag"],
      [/\%(I|Q|M)(X|B|W|D|L)[0-9\.]+/, "tag"],
      [/\%(I|Q|M)[0-9\.]*/, "tag"],
      [/\b[A-Za-z]{1,6}#[0-9]+\b/, "tag"],
      [/\b(TO_|CTU_|CTD_|CTUD_|MUX_|SEL_)[A_Za-z]+\b/, "predefined"],
      [/\b[A_Za-z]+(_TO_)[A_Za-z]+\b/, "predefined"],
      [/[;]/, "delimiter"],
      [/[.]/, { token: "delimiter", next: "@params" }],
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@operators": "operators",
            "@keywords": "keyword",
            "@typeKeywords": "type",
            "@defineKeywords": "variable",
            "@constant": "constant",
            "@builtinVariables": "predefined",
            "@builtinFunctions": "predefined",
            "@default": "identifier"
          }
        }
      ],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, { token: "string.quote", bracket: "@open", next: "@string_dq" }],
      [/'/, { token: "string.quote", bracket: "@open", next: "@string_sq" }],
      [/'[^\\']'/, "string"],
      [/(')(@escapes)(')/, ["string", "string.escape", "string"]],
      [/'/, "string.invalid"]
    ],
    params: [
      [/\b[A-Za-z0-9_]+\b(?=\()/, { token: "identifier", next: "@pop" }],
      [/\b[A-Za-z0-9_]+\b/, "variable.name", "@pop"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\/\*/, "comment", "@push"],
      ["\\*/", "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    comment2: [
      [/[^\(*]+/, "comment"],
      [/\(\*/, "comment", "@push"],
      ["\\*\\)", "comment", "@pop"],
      [/[\(*]/, "comment"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/\/\/.*$/, "comment"],
      [/\/\*/, "comment", "@comment"],
      [/\(\*/, "comment", "@comment2"]
    ],
    string_dq: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    string_sq: [
      [/[^\\']+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/'/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/swift/swift.js":
/*!**************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/swift/swift.js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/swift/swift.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "`", close: "`" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".swift",
  identifier: /[a-zA-Z_][\w$]*/,
  attributes: [
    "@GKInspectable",
    "@IBAction",
    "@IBDesignable",
    "@IBInspectable",
    "@IBOutlet",
    "@IBSegueAction",
    "@NSApplicationMain",
    "@NSCopying",
    "@NSManaged",
    "@Sendable",
    "@UIApplicationMain",
    "@autoclosure",
    "@actorIndependent",
    "@asyncHandler",
    "@available",
    "@convention",
    "@derivative",
    "@differentiable",
    "@discardableResult",
    "@dynamicCallable",
    "@dynamicMemberLookup",
    "@escaping",
    "@frozen",
    "@globalActor",
    "@inlinable",
    "@inline",
    "@main",
    "@noDerivative",
    "@nonobjc",
    "@noreturn",
    "@objc",
    "@objcMembers",
    "@preconcurrency",
    "@propertyWrapper",
    "@requires_stored_property_inits",
    "@resultBuilder",
    "@testable",
    "@unchecked",
    "@unknown",
    "@usableFromInline",
    "@warn_unqualified_access"
  ],
  accessmodifiers: ["open", "public", "internal", "fileprivate", "private"],
  keywords: [
    "#available",
    "#colorLiteral",
    "#column",
    "#dsohandle",
    "#else",
    "#elseif",
    "#endif",
    "#error",
    "#file",
    "#fileID",
    "#fileLiteral",
    "#filePath",
    "#function",
    "#if",
    "#imageLiteral",
    "#keyPath",
    "#line",
    "#selector",
    "#sourceLocation",
    "#warning",
    "Any",
    "Protocol",
    "Self",
    "Type",
    "actor",
    "as",
    "assignment",
    "associatedtype",
    "associativity",
    "async",
    "await",
    "break",
    "case",
    "catch",
    "class",
    "continue",
    "convenience",
    "default",
    "defer",
    "deinit",
    "didSet",
    "do",
    "dynamic",
    "dynamicType",
    "else",
    "enum",
    "extension",
    "fallthrough",
    "false",
    "fileprivate",
    "final",
    "for",
    "func",
    "get",
    "guard",
    "higherThan",
    "if",
    "import",
    "in",
    "indirect",
    "infix",
    "init",
    "inout",
    "internal",
    "is",
    "isolated",
    "lazy",
    "left",
    "let",
    "lowerThan",
    "mutating",
    "nil",
    "none",
    "nonisolated",
    "nonmutating",
    "open",
    "operator",
    "optional",
    "override",
    "postfix",
    "precedence",
    "precedencegroup",
    "prefix",
    "private",
    "protocol",
    "public",
    "repeat",
    "required",
    "rethrows",
    "return",
    "right",
    "safe",
    "self",
    "set",
    "some",
    "static",
    "struct",
    "subscript",
    "super",
    "switch",
    "throw",
    "throws",
    "true",
    "try",
    "typealias",
    "unowned",
    "unsafe",
    "var",
    "weak",
    "where",
    "while",
    "willSet",
    "__consuming",
    "__owned"
  ],
  symbols: /[=(){}\[\].,:;@#\_&\-<>`?!+*\\\/]/,
  operatorstart: /[\/=\-+!*%<>&|^~?\u00A1-\u00A7\u00A9\u00AB\u00AC\u00AE\u00B0-\u00B1\u00B6\u00BB\u00BF\u00D7\u00F7\u2016-\u2017\u2020-\u2027\u2030-\u203E\u2041-\u2053\u2055-\u205E\u2190-\u23FF\u2500-\u2775\u2794-\u2BFF\u2E00-\u2E7F\u3001-\u3003\u3008-\u3030]/,
  operatorend: /[\u0300-\u036F\u1DC0-\u1DFF\u20D0-\u20FF\uFE00-\uFE0F\uFE20-\uFE2F\uE0100-\uE01EF]/,
  operators: /(@operatorstart)((@operatorstart)|(@operatorend))*/,
  escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      { include: "@comment" },
      { include: "@attribute" },
      { include: "@literal" },
      { include: "@keyword" },
      { include: "@invokedmethod" },
      { include: "@symbol" }
    ],
    whitespace: [
      [/\s+/, "white"],
      [/"""/, "string.quote", "@endDblDocString"]
    ],
    endDblDocString: [
      [/[^"]+/, "string"],
      [/\\"/, "string"],
      [/"""/, "string.quote", "@popall"],
      [/"/, "string"]
    ],
    symbol: [
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [/[.]/, "delimiter"],
      [/@operators/, "operator"],
      [/@symbols/, "operator"]
    ],
    comment: [
      [/\/\/\/.*$/, "comment.doc"],
      [/\/\*\*/, "comment.doc", "@commentdocbody"],
      [/\/\/.*$/, "comment"],
      [/\/\*/, "comment", "@commentbody"]
    ],
    commentdocbody: [
      [/\/\*/, "comment", "@commentbody"],
      [/\*\//, "comment.doc", "@pop"],
      [/\:[a-zA-Z]+\:/, "comment.doc.param"],
      [/./, "comment.doc"]
    ],
    commentbody: [
      [/\/\*/, "comment", "@commentbody"],
      [/\*\//, "comment", "@pop"],
      [/./, "comment"]
    ],
    attribute: [
      [
        /@@@identifier/,
        {
          cases: {
            "@attributes": "keyword.control",
            "@default": ""
          }
        }
      ]
    ],
    literal: [
      [/"/, { token: "string.quote", next: "@stringlit" }],
      [/0[b]([01]_?)+/, "number.binary"],
      [/0[o]([0-7]_?)+/, "number.octal"],
      [/0[x]([0-9a-fA-F]_?)+([pP][\-+](\d_?)+)?/, "number.hex"],
      [/(\d_?)*\.(\d_?)+([eE][\-+]?(\d_?)+)?/, "number.float"],
      [/(\d_?)+/, "number"]
    ],
    stringlit: [
      [/\\\(/, { token: "operator", next: "@interpolatedexpression" }],
      [/@escapes/, "string"],
      [/\\./, "string.escape.invalid"],
      [/"/, { token: "string.quote", next: "@pop" }],
      [/./, "string"]
    ],
    interpolatedexpression: [
      [/\(/, { token: "operator", next: "@interpolatedexpression" }],
      [/\)/, { token: "operator", next: "@pop" }],
      { include: "@literal" },
      { include: "@keyword" },
      { include: "@symbol" }
    ],
    keyword: [
      [/`/, { token: "operator", next: "@escapedkeyword" }],
      [
        /@identifier/,
        {
          cases: {
            "@keywords": "keyword",
            "[A-Z][a-zA-Z0-9$]*": "type.identifier",
            "@default": "identifier"
          }
        }
      ]
    ],
    escapedkeyword: [
      [/`/, { token: "operator", next: "@pop" }],
      [/./, "identifier"]
    ],
    invokedmethod: [
      [
        /([.])(@identifier)/,
        {
          cases: {
            $2: ["delimeter", "type.identifier"],
            "@default": ""
          }
        }
      ]
    ]
  }
};

/*!---------------------------------------------------------------------------------------------
 *  Copyright (C) David Owens II, owensd.io. All rights reserved.
 *--------------------------------------------------------------------------------------------*/


/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/systemverilog/systemverilog.js":
/*!******************************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/systemverilog/systemverilog.js ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/systemverilog/systemverilog.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["begin", "end"],
    ["case", "endcase"],
    ["casex", "endcase"],
    ["casez", "endcase"],
    ["checker", "endchecker"],
    ["class", "endclass"],
    ["clocking", "endclocking"],
    ["config", "endconfig"],
    ["function", "endfunction"],
    ["generate", "endgenerate"],
    ["group", "endgroup"],
    ["interface", "endinterface"],
    ["module", "endmodule"],
    ["package", "endpackage"],
    ["primitive", "endprimitive"],
    ["program", "endprogram"],
    ["property", "endproperty"],
    ["specify", "endspecify"],
    ["sequence", "endsequence"],
    ["table", "endtable"],
    ["task", "endtask"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "(", close: ")" },
    { open: "'", close: "'", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string"] }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    offSide: false,
    markers: {
      start: new RegExp("^(?:\\s*|.*(?!\\/[\\/\\*])[^\\w])(?:begin|case(x|z)?|class|clocking|config|covergroup|function|generate|interface|module|package|primitive|property|program|sequence|specify|table|task)\\b"),
      end: new RegExp("^(?:\\s*|.*(?!\\/[\\/\\*])[^\\w])(?:end|endcase|endclass|endclocking|endconfig|endgroup|endfunction|endgenerate|endinterface|endmodule|endpackage|endprimitive|endproperty|endprogram|endsequence|endspecify|endtable|endtask)\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".sv",
  brackets: [
    { token: "delimiter.curly", open: "{", close: "}" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.square", open: "[", close: "]" },
    { token: "delimiter.angle", open: "<", close: ">" }
  ],
  keywords: [
    "accept_on",
    "alias",
    "always",
    "always_comb",
    "always_ff",
    "always_latch",
    "and",
    "assert",
    "assign",
    "assume",
    "automatic",
    "before",
    "begin",
    "bind",
    "bins",
    "binsof",
    "bit",
    "break",
    "buf",
    "bufif0",
    "bufif1",
    "byte",
    "case",
    "casex",
    "casez",
    "cell",
    "chandle",
    "checker",
    "class",
    "clocking",
    "cmos",
    "config",
    "const",
    "constraint",
    "context",
    "continue",
    "cover",
    "covergroup",
    "coverpoint",
    "cross",
    "deassign",
    "default",
    "defparam",
    "design",
    "disable",
    "dist",
    "do",
    "edge",
    "else",
    "end",
    "endcase",
    "endchecker",
    "endclass",
    "endclocking",
    "endconfig",
    "endfunction",
    "endgenerate",
    "endgroup",
    "endinterface",
    "endmodule",
    "endpackage",
    "endprimitive",
    "endprogram",
    "endproperty",
    "endspecify",
    "endsequence",
    "endtable",
    "endtask",
    "enum",
    "event",
    "eventually",
    "expect",
    "export",
    "extends",
    "extern",
    "final",
    "first_match",
    "for",
    "force",
    "foreach",
    "forever",
    "fork",
    "forkjoin",
    "function",
    "generate",
    "genvar",
    "global",
    "highz0",
    "highz1",
    "if",
    "iff",
    "ifnone",
    "ignore_bins",
    "illegal_bins",
    "implements",
    "implies",
    "import",
    "incdir",
    "include",
    "initial",
    "inout",
    "input",
    "inside",
    "instance",
    "int",
    "integer",
    "interconnect",
    "interface",
    "intersect",
    "join",
    "join_any",
    "join_none",
    "large",
    "let",
    "liblist",
    "library",
    "local",
    "localparam",
    "logic",
    "longint",
    "macromodule",
    "matches",
    "medium",
    "modport",
    "module",
    "nand",
    "negedge",
    "nettype",
    "new",
    "nexttime",
    "nmos",
    "nor",
    "noshowcancelled",
    "not",
    "notif0",
    "notif1",
    "null",
    "or",
    "output",
    "package",
    "packed",
    "parameter",
    "pmos",
    "posedge",
    "primitive",
    "priority",
    "program",
    "property",
    "protected",
    "pull0",
    "pull1",
    "pulldown",
    "pullup",
    "pulsestyle_ondetect",
    "pulsestyle_onevent",
    "pure",
    "rand",
    "randc",
    "randcase",
    "randsequence",
    "rcmos",
    "real",
    "realtime",
    "ref",
    "reg",
    "reject_on",
    "release",
    "repeat",
    "restrict",
    "return",
    "rnmos",
    "rpmos",
    "rtran",
    "rtranif0",
    "rtranif1",
    "s_always",
    "s_eventually",
    "s_nexttime",
    "s_until",
    "s_until_with",
    "scalared",
    "sequence",
    "shortint",
    "shortreal",
    "showcancelled",
    "signed",
    "small",
    "soft",
    "solve",
    "specify",
    "specparam",
    "static",
    "string",
    "strong",
    "strong0",
    "strong1",
    "struct",
    "super",
    "supply0",
    "supply1",
    "sync_accept_on",
    "sync_reject_on",
    "table",
    "tagged",
    "task",
    "this",
    "throughout",
    "time",
    "timeprecision",
    "timeunit",
    "tran",
    "tranif0",
    "tranif1",
    "tri",
    "tri0",
    "tri1",
    "triand",
    "trior",
    "trireg",
    "type",
    "typedef",
    "union",
    "unique",
    "unique0",
    "unsigned",
    "until",
    "until_with",
    "untyped",
    "use",
    "uwire",
    "var",
    "vectored",
    "virtual",
    "void",
    "wait",
    "wait_order",
    "wand",
    "weak",
    "weak0",
    "weak1",
    "while",
    "wildcard",
    "wire",
    "with",
    "within",
    "wor",
    "xnor",
    "xor"
  ],
  builtin_gates: [
    "and",
    "nand",
    "nor",
    "or",
    "xor",
    "xnor",
    "buf",
    "not",
    "bufif0",
    "bufif1",
    "notif1",
    "notif0",
    "cmos",
    "nmos",
    "pmos",
    "rcmos",
    "rnmos",
    "rpmos",
    "tran",
    "tranif1",
    "tranif0",
    "rtran",
    "rtranif1",
    "rtranif0"
  ],
  operators: [
    "=",
    "+=",
    "-=",
    "*=",
    "/=",
    "%=",
    "&=",
    "|=",
    "^=",
    "<<=",
    ">>+",
    "<<<=",
    ">>>=",
    "?",
    ":",
    "+",
    "-",
    "!",
    "~",
    "&",
    "~&",
    "|",
    "~|",
    "^",
    "~^",
    "^~",
    "+",
    "-",
    "*",
    "/",
    "%",
    "==",
    "!=",
    "===",
    "!==",
    "==?",
    "!=?",
    "&&",
    "||",
    "**",
    "<",
    "<=",
    ">",
    ">=",
    "&",
    "|",
    "^",
    ">>",
    "<<",
    ">>>",
    "<<<",
    "++",
    "--",
    "->",
    "<->",
    "inside",
    "dist",
    "::",
    "+:",
    "-:",
    "*>",
    "&&&",
    "|->",
    "|=>",
    "#=#"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%#]+/,
  escapes: /%%|\\(?:[antvf\\"']|x[0-9A-Fa-f]{1,2}|[0-7]{1,3})/,
  identifier: /(?:[a-zA-Z_][a-zA-Z0-9_$\.]*|\\\S+ )/,
  systemcall: /[$][a-zA-Z0-9_]+/,
  timeunits: /s|ms|us|ns|ps|fs/,
  tokenizer: {
    root: [
      [
        /^(\s*)(@identifier)/,
        [
          "",
          {
            cases: {
              "@builtin_gates": {
                token: "keyword.$2",
                next: "@module_instance"
              },
              table: {
                token: "keyword.$2",
                next: "@table"
              },
              "@keywords": { token: "keyword.$2" },
              "@default": {
                token: "identifier",
                next: "@module_instance"
              }
            }
          }
        ]
      ],
      [/^\s*`include/, { token: "keyword.directive.include", next: "@include" }],
      [/^\s*`\s*\w+/, "keyword"],
      { include: "@identifier_or_keyword" },
      { include: "@whitespace" },
      [/\(\*.*\*\)/, "annotation"],
      [/@systemcall/, "variable.predefined"],
      [/[{}()\[\]]/, "@brackets"],
      [/[<>](?!@symbols)/, "@brackets"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "delimiter",
            "@default": ""
          }
        }
      ],
      { include: "@numbers" },
      [/[;,.]/, "delimiter"],
      { include: "@strings" }
    ],
    identifier_or_keyword: [
      [
        /@identifier/,
        {
          cases: {
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ]
    ],
    numbers: [
      [/\d+?[\d_]*(?:\.[\d_]+)?[eE][\-+]?\d+/, "number.float"],
      [/\d+?[\d_]*\.[\d_]+(?:\s*@timeunits)?/, "number.float"],
      [/(?:\d+?[\d_]*\s*)?'[sS]?[dD]\s*[0-9xXzZ?]+?[0-9xXzZ?_]*/, "number"],
      [/(?:\d+?[\d_]*\s*)?'[sS]?[bB]\s*[0-1xXzZ?]+?[0-1xXzZ?_]*/, "number.binary"],
      [/(?:\d+?[\d_]*\s*)?'[sS]?[oO]\s*[0-7xXzZ?]+?[0-7xXzZ?_]*/, "number.octal"],
      [/(?:\d+?[\d_]*\s*)?'[sS]?[hH]\s*[0-9a-fA-FxXzZ?]+?[0-9a-fA-FxXzZ?_]*/, "number.hex"],
      [/1step/, "number"],
      [/[\dxXzZ]+?[\dxXzZ_]*(?:\s*@timeunits)?/, "number"],
      [/'[01xXzZ]+/, "number"]
    ],
    module_instance: [
      { include: "@whitespace" },
      [/(#?)(\()/, ["", { token: "@brackets", next: "@port_connection" }]],
      [/@identifier\s*[;={}\[\],]/, { token: "@rematch", next: "@pop" }],
      [/@symbols|[;={}\[\],]/, { token: "@rematch", next: "@pop" }],
      [/@identifier/, "type"],
      [/;/, "delimiter", "@pop"]
    ],
    port_connection: [
      { include: "@identifier_or_keyword" },
      { include: "@whitespace" },
      [/@systemcall/, "variable.predefined"],
      { include: "@numbers" },
      { include: "@strings" },
      [/[,]/, "delimiter"],
      [/\(/, "@brackets", "@port_connection"],
      [/\)/, "@brackets", "@pop"]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/\/\*/, "comment", "@comment"],
      [/\/\/.*$/, "comment"]
    ],
    comment: [
      [/[^\/*]+/, "comment"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    strings: [
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/"/, "string", "@string"]
    ],
    string: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    include: [
      [
        /(\s*)(")([\w*\/*]*)(.\w*)(")/,
        [
          "",
          "string.include.identifier",
          "string.include.identifier",
          "string.include.identifier",
          { token: "string.include.identifier", next: "@pop" }
        ]
      ],
      [
        /(\s*)(<)([\w*\/*]*)(.\w*)(>)/,
        [
          "",
          "string.include.identifier",
          "string.include.identifier",
          "string.include.identifier",
          { token: "string.include.identifier", next: "@pop" }
        ]
      ]
    ],
    table: [
      { include: "@whitespace" },
      [/[()]/, "@brackets"],
      [/[:;]/, "delimiter"],
      [/[01\-*?xXbBrRfFpPnN]/, "variable.predefined"],
      ["endtable", "keyword.endtable", "@pop"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/tcl/tcl.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/tcl/tcl.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/tcl/tcl.ts
var conf = {
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ]
};
var language = {
  tokenPostfix: ".tcl",
  specialFunctions: [
    "set",
    "unset",
    "rename",
    "variable",
    "proc",
    "coroutine",
    "foreach",
    "incr",
    "append",
    "lappend",
    "linsert",
    "lreplace"
  ],
  mainFunctions: [
    "if",
    "then",
    "elseif",
    "else",
    "case",
    "switch",
    "while",
    "for",
    "break",
    "continue",
    "return",
    "package",
    "namespace",
    "catch",
    "exit",
    "eval",
    "expr",
    "uplevel",
    "upvar"
  ],
  builtinFunctions: [
    "file",
    "info",
    "concat",
    "join",
    "lindex",
    "list",
    "llength",
    "lrange",
    "lsearch",
    "lsort",
    "split",
    "array",
    "parray",
    "binary",
    "format",
    "regexp",
    "regsub",
    "scan",
    "string",
    "subst",
    "dict",
    "cd",
    "clock",
    "exec",
    "glob",
    "pid",
    "pwd",
    "close",
    "eof",
    "fblocked",
    "fconfigure",
    "fcopy",
    "fileevent",
    "flush",
    "gets",
    "open",
    "puts",
    "read",
    "seek",
    "socket",
    "tell",
    "interp",
    "after",
    "auto_execok",
    "auto_load",
    "auto_mkindex",
    "auto_reset",
    "bgerror",
    "error",
    "global",
    "history",
    "load",
    "source",
    "time",
    "trace",
    "unknown",
    "unset",
    "update",
    "vwait",
    "winfo",
    "wm",
    "bind",
    "event",
    "pack",
    "place",
    "grid",
    "font",
    "bell",
    "clipboard",
    "destroy",
    "focus",
    "grab",
    "lower",
    "option",
    "raise",
    "selection",
    "send",
    "tk",
    "tkwait",
    "tk_bisque",
    "tk_focusNext",
    "tk_focusPrev",
    "tk_focusFollowsMouse",
    "tk_popup",
    "tk_setPalette"
  ],
  symbols: /[=><!~?:&|+\-*\/\^%]+/,
  brackets: [
    { open: "(", close: ")", token: "delimiter.parenthesis" },
    { open: "{", close: "}", token: "delimiter.curly" },
    { open: "[", close: "]", token: "delimiter.square" }
  ],
  escapes: /\\(?:[abfnrtv\\"'\[\]\{\};\$]|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  variables: /(?:\$+(?:(?:\:\:?)?[a-zA-Z_]\w*)+)/,
  tokenizer: {
    root: [
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@specialFunctions": {
              token: "keyword.flow",
              next: "@specialFunc"
            },
            "@mainFunctions": "keyword",
            "@builtinFunctions": "variable",
            "@default": "operator.scss"
          }
        }
      ],
      [/\s+\-+(?!\d|\.)\w*|{\*}/, "metatag"],
      { include: "@whitespace" },
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, "operator"],
      [/\$+(?:\:\:)?\{/, { token: "identifier", next: "@nestedVariable" }],
      [/@variables/, "type.identifier"],
      [/\.(?!\d|\.)[\w\-]*/, "operator.sql"],
      [/\d+(\.\d+)?/, "number"],
      [/\d+/, "number"],
      [/;/, "delimiter"],
      [/"/, { token: "string.quote", bracket: "@open", next: "@dstring" }],
      [/'/, { token: "string.quote", bracket: "@open", next: "@sstring" }]
    ],
    dstring: [
      [/\[/, { token: "@brackets", next: "@nestedCall" }],
      [/\$+(?:\:\:)?\{/, { token: "identifier", next: "@nestedVariable" }],
      [/@variables/, "type.identifier"],
      [/[^\\$\[\]"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/"/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    sstring: [
      [/\[/, { token: "@brackets", next: "@nestedCall" }],
      [/\$+(?:\:\:)?\{/, { token: "identifier", next: "@nestedVariable" }],
      [/@variables/, "type.identifier"],
      [/[^\\$\[\]']+/, "string"],
      [/@escapes/, "string.escape"],
      [/'/, { token: "string.quote", bracket: "@close", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, "white"],
      [/#.*\\$/, { token: "comment", next: "@newlineComment" }],
      [/#.*(?!\\)$/, "comment"]
    ],
    newlineComment: [
      [/.*\\$/, "comment"],
      [/.*(?!\\)$/, { token: "comment", next: "@pop" }]
    ],
    nestedVariable: [
      [/[^\{\}\$]+/, "type.identifier"],
      [/\}/, { token: "identifier", next: "@pop" }]
    ],
    nestedCall: [
      [/\[/, { token: "@brackets", next: "@nestedCall" }],
      [/\]/, { token: "@brackets", next: "@pop" }],
      { include: "root" }
    ],
    specialFunc: [
      [/"/, { token: "string", next: "@dstring" }],
      [/'/, { token: "string", next: "@sstring" }],
      [/\S+/, { token: "type", next: "@pop" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/twig/twig.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/twig/twig.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/twig/twig.ts
var conf = {
  wordPattern: /(-?\d*\.\d\w*)|([^\`\~\!\@\$\^\&\*\(\)\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\s]+)/g,
  comments: {
    blockComment: ["{#", "#}"]
  },
  brackets: [
    ["{#", "#}"],
    ["{%", "%}"],
    ["{{", "}}"],
    ["(", ")"],
    ["[", "]"],
    ["<!--", "-->"],
    ["<", ">"]
  ],
  autoClosingPairs: [
    { open: "{# ", close: " #}" },
    { open: "{% ", close: " %}" },
    { open: "{{ ", close: " }}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: '"', close: '"' },
    { open: "'", close: "'" },
    { open: "<", close: ">" }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: "",
  ignoreCase: true,
  keywords: [
    "apply",
    "autoescape",
    "block",
    "deprecated",
    "do",
    "embed",
    "extends",
    "flush",
    "for",
    "from",
    "if",
    "import",
    "include",
    "macro",
    "sandbox",
    "set",
    "use",
    "verbatim",
    "with",
    "endapply",
    "endautoescape",
    "endblock",
    "endembed",
    "endfor",
    "endif",
    "endmacro",
    "endsandbox",
    "endset",
    "endwith",
    "true",
    "false"
  ],
  tokenizer: {
    root: [
      [/\s+/],
      [/{#/, "comment.twig", "@commentState"],
      [/{%[-~]?/, "delimiter.twig", "@blockState"],
      [/{{[-~]?/, "delimiter.twig", "@variableState"],
      [/<!DOCTYPE/, "metatag.html", "@doctype"],
      [/<!--/, "comment.html", "@comment"],
      [/(<)((?:[\w\-]+:)?[\w\-]+)(\s*)(\/>)/, ["delimiter.html", "tag.html", "", "delimiter.html"]],
      [/(<)(script)/, ["delimiter.html", { token: "tag.html", next: "@script" }]],
      [/(<)(style)/, ["delimiter.html", { token: "tag.html", next: "@style" }]],
      [/(<)((?:[\w\-]+:)?[\w\-]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/(<\/)((?:[\w\-]+:)?[\w\-]+)/, ["delimiter.html", { token: "tag.html", next: "@otherTag" }]],
      [/</, "delimiter.html"],
      [/[^<{]+/]
    ],
    commentState: [
      [/#}/, "comment.twig", "@pop"],
      [/./, "comment.twig"]
    ],
    blockState: [
      [/[-~]?%}/, "delimiter.twig", "@pop"],
      [/\s+/],
      [
        /(verbatim)(\s*)([-~]?%})/,
        ["keyword.twig", "", { token: "delimiter.twig", next: "@rawDataState" }]
      ],
      { include: "expression" }
    ],
    rawDataState: [
      [
        /({%[-~]?)(\s*)(endverbatim)(\s*)([-~]?%})/,
        ["delimiter.twig", "", "keyword.twig", "", { token: "delimiter.twig", next: "@popall" }]
      ],
      [/./, "string.twig"]
    ],
    variableState: [[/[-~]?}}/, "delimiter.twig", "@pop"], { include: "expression" }],
    stringState: [
      [/"/, "string.twig", "@pop"],
      [/#{\s*/, "string.twig", "@interpolationState"],
      [/[^#"\\]*(?:(?:\\.|#(?!\{))[^#"\\]*)*/, "string.twig"]
    ],
    interpolationState: [
      [/}/, "string.twig", "@pop"],
      { include: "expression" }
    ],
    expression: [
      [/\s+/],
      [/\+|-|\/{1,2}|%|\*{1,2}/, "operators.twig"],
      [/(and|or|not|b-and|b-xor|b-or)(\s+)/, ["operators.twig", ""]],
      [/==|!=|<|>|>=|<=/, "operators.twig"],
      [/(starts with|ends with|matches)(\s+)/, ["operators.twig", ""]],
      [/(in)(\s+)/, ["operators.twig", ""]],
      [/(is)(\s+)/, ["operators.twig", ""]],
      [/\||~|:|\.{1,2}|\?{1,2}/, "operators.twig"],
      [
        /[^\W\d][\w]*/,
        {
          cases: {
            "@keywords": "keyword.twig",
            "@default": "variable.twig"
          }
        }
      ],
      [/\d+(\.\d+)?/, "number.twig"],
      [/\(|\)|\[|\]|{|}|,/, "delimiter.twig"],
      [/"([^#"\\]*(?:\\.[^#"\\]*)*)"|\'([^\'\\]*(?:\\.[^\'\\]*)*)\'/, "string.twig"],
      [/"/, "string.twig", "@stringState"],
      [/=>/, "operators.twig"],
      [/=/, "operators.twig"]
    ],
    doctype: [
      [/[^>]+/, "metatag.content.html"],
      [/>/, "metatag.html", "@pop"]
    ],
    comment: [
      [/-->/, "comment.html", "@pop"],
      [/[^-]+/, "comment.content.html"],
      [/./, "comment.content.html"]
    ],
    otherTag: [
      [/\/?>/, "delimiter.html", "@pop"],
      [/"([^"]*)"/, "attribute.value.html"],
      [/'([^']*)'/, "attribute.value.html"],
      [/[\w\-]+/, "attribute.name.html"],
      [/=/, "delimiter.html"],
      [/[ \t\r\n]+/]
    ],
    script: [
      [/type/, "attribute.name.html", "@scriptAfterType"],
      [/"([^"]*)"/, "attribute.value.html"],
      [/'([^']*)'/, "attribute.value.html"],
      [/[\w\-]+/, "attribute.name.html"],
      [/=/, "delimiter.html"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(script\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    scriptAfterType: [
      [/=/, "delimiter.html", "@scriptAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptAfterTypeEquals: [
      [
        /"([^"]*)"/,
        {
          token: "attribute.value.html",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value.html",
          switchTo: "@scriptWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded",
          nextEmbedded: "text/javascript"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptWithCustomType: [
      [
        />/,
        {
          token: "delimiter.html",
          next: "@scriptEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value.html"],
      [/'([^']*)'/, "attribute.value.html"],
      [/[\w\-]+/, "attribute.name.html"],
      [/=/, "delimiter.html"],
      [/[ \t\r\n]+/],
      [/<\/script\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    scriptEmbedded: [
      [/<\/script/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }],
      [/[^<]+/, ""]
    ],
    style: [
      [/type/, "attribute.name.html", "@styleAfterType"],
      [/"([^"]*)"/, "attribute.value.html"],
      [/'([^']*)'/, "attribute.value.html"],
      [/[\w\-]+/, "attribute.name.html"],
      [/=/, "delimiter.html"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [
        /(<\/)(style\s*)(>)/,
        ["delimiter.html", "tag.html", { token: "delimiter.html", next: "@pop" }]
      ]
    ],
    styleAfterType: [
      [/=/, "delimiter.html", "@styleAfterTypeEquals"],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleAfterTypeEquals: [
      [
        /"([^"]*)"/,
        {
          token: "attribute.value.html",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        /'([^']*)'/,
        {
          token: "attribute.value.html",
          switchTo: "@styleWithCustomType.$1"
        }
      ],
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded",
          nextEmbedded: "text/css"
        }
      ],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleWithCustomType: [
      [
        />/,
        {
          token: "delimiter.html",
          next: "@styleEmbedded.$S2",
          nextEmbedded: "$S2"
        }
      ],
      [/"([^"]*)"/, "attribute.value.html"],
      [/'([^']*)'/, "attribute.value.html"],
      [/[\w\-]+/, "attribute.name.html"],
      [/=/, "delimiter.html"],
      [/[ \t\r\n]+/],
      [/<\/style\s*>/, { token: "@rematch", next: "@pop" }]
    ],
    styleEmbedded: [
      [/<\/style/, { token: "@rematch", next: "@pop", nextEmbedded: "@pop" }],
      [/[^<]+/, ""]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/vb/vb.js":
/*!********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/vb/vb.js ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/vb/vb.ts
var conf = {
  comments: {
    lineComment: "'",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"],
    ["<", ">"],
    ["addhandler", "end addhandler"],
    ["class", "end class"],
    ["enum", "end enum"],
    ["event", "end event"],
    ["function", "end function"],
    ["get", "end get"],
    ["if", "end if"],
    ["interface", "end interface"],
    ["module", "end module"],
    ["namespace", "end namespace"],
    ["operator", "end operator"],
    ["property", "end property"],
    ["raiseevent", "end raiseevent"],
    ["removehandler", "end removehandler"],
    ["select", "end select"],
    ["set", "end set"],
    ["structure", "end structure"],
    ["sub", "end sub"],
    ["synclock", "end synclock"],
    ["try", "end try"],
    ["while", "end while"],
    ["with", "end with"],
    ["using", "end using"],
    ["do", "loop"],
    ["for", "next"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string", "comment"] },
    { open: "[", close: "]", notIn: ["string", "comment"] },
    { open: "(", close: ")", notIn: ["string", "comment"] },
    { open: '"', close: '"', notIn: ["string", "comment"] },
    { open: "<", close: ">", notIn: ["string", "comment"] }
  ],
  folding: {
    markers: {
      start: new RegExp("^\\s*#Region\\b"),
      end: new RegExp("^\\s*#End Region\\b")
    }
  }
};
var language = {
  defaultToken: "",
  tokenPostfix: ".vb",
  ignoreCase: true,
  brackets: [
    { token: "delimiter.bracket", open: "{", close: "}" },
    { token: "delimiter.array", open: "[", close: "]" },
    { token: "delimiter.parenthesis", open: "(", close: ")" },
    { token: "delimiter.angle", open: "<", close: ">" },
    {
      token: "keyword.tag-addhandler",
      open: "addhandler",
      close: "end addhandler"
    },
    { token: "keyword.tag-class", open: "class", close: "end class" },
    { token: "keyword.tag-enum", open: "enum", close: "end enum" },
    { token: "keyword.tag-event", open: "event", close: "end event" },
    {
      token: "keyword.tag-function",
      open: "function",
      close: "end function"
    },
    { token: "keyword.tag-get", open: "get", close: "end get" },
    { token: "keyword.tag-if", open: "if", close: "end if" },
    {
      token: "keyword.tag-interface",
      open: "interface",
      close: "end interface"
    },
    { token: "keyword.tag-module", open: "module", close: "end module" },
    {
      token: "keyword.tag-namespace",
      open: "namespace",
      close: "end namespace"
    },
    {
      token: "keyword.tag-operator",
      open: "operator",
      close: "end operator"
    },
    {
      token: "keyword.tag-property",
      open: "property",
      close: "end property"
    },
    {
      token: "keyword.tag-raiseevent",
      open: "raiseevent",
      close: "end raiseevent"
    },
    {
      token: "keyword.tag-removehandler",
      open: "removehandler",
      close: "end removehandler"
    },
    { token: "keyword.tag-select", open: "select", close: "end select" },
    { token: "keyword.tag-set", open: "set", close: "end set" },
    {
      token: "keyword.tag-structure",
      open: "structure",
      close: "end structure"
    },
    { token: "keyword.tag-sub", open: "sub", close: "end sub" },
    {
      token: "keyword.tag-synclock",
      open: "synclock",
      close: "end synclock"
    },
    { token: "keyword.tag-try", open: "try", close: "end try" },
    { token: "keyword.tag-while", open: "while", close: "end while" },
    { token: "keyword.tag-with", open: "with", close: "end with" },
    { token: "keyword.tag-using", open: "using", close: "end using" },
    { token: "keyword.tag-do", open: "do", close: "loop" },
    { token: "keyword.tag-for", open: "for", close: "next" }
  ],
  keywords: [
    "AddHandler",
    "AddressOf",
    "Alias",
    "And",
    "AndAlso",
    "As",
    "Async",
    "Boolean",
    "ByRef",
    "Byte",
    "ByVal",
    "Call",
    "Case",
    "Catch",
    "CBool",
    "CByte",
    "CChar",
    "CDate",
    "CDbl",
    "CDec",
    "Char",
    "CInt",
    "Class",
    "CLng",
    "CObj",
    "Const",
    "Continue",
    "CSByte",
    "CShort",
    "CSng",
    "CStr",
    "CType",
    "CUInt",
    "CULng",
    "CUShort",
    "Date",
    "Decimal",
    "Declare",
    "Default",
    "Delegate",
    "Dim",
    "DirectCast",
    "Do",
    "Double",
    "Each",
    "Else",
    "ElseIf",
    "End",
    "EndIf",
    "Enum",
    "Erase",
    "Error",
    "Event",
    "Exit",
    "False",
    "Finally",
    "For",
    "Friend",
    "Function",
    "Get",
    "GetType",
    "GetXMLNamespace",
    "Global",
    "GoSub",
    "GoTo",
    "Handles",
    "If",
    "Implements",
    "Imports",
    "In",
    "Inherits",
    "Integer",
    "Interface",
    "Is",
    "IsNot",
    "Let",
    "Lib",
    "Like",
    "Long",
    "Loop",
    "Me",
    "Mod",
    "Module",
    "MustInherit",
    "MustOverride",
    "MyBase",
    "MyClass",
    "NameOf",
    "Namespace",
    "Narrowing",
    "New",
    "Next",
    "Not",
    "Nothing",
    "NotInheritable",
    "NotOverridable",
    "Object",
    "Of",
    "On",
    "Operator",
    "Option",
    "Optional",
    "Or",
    "OrElse",
    "Out",
    "Overloads",
    "Overridable",
    "Overrides",
    "ParamArray",
    "Partial",
    "Private",
    "Property",
    "Protected",
    "Public",
    "RaiseEvent",
    "ReadOnly",
    "ReDim",
    "RemoveHandler",
    "Resume",
    "Return",
    "SByte",
    "Select",
    "Set",
    "Shadows",
    "Shared",
    "Short",
    "Single",
    "Static",
    "Step",
    "Stop",
    "String",
    "Structure",
    "Sub",
    "SyncLock",
    "Then",
    "Throw",
    "To",
    "True",
    "Try",
    "TryCast",
    "TypeOf",
    "UInteger",
    "ULong",
    "UShort",
    "Using",
    "Variant",
    "Wend",
    "When",
    "While",
    "Widening",
    "With",
    "WithEvents",
    "WriteOnly",
    "Xor"
  ],
  tagwords: [
    "If",
    "Sub",
    "Select",
    "Try",
    "Class",
    "Enum",
    "Function",
    "Get",
    "Interface",
    "Module",
    "Namespace",
    "Operator",
    "Set",
    "Structure",
    "Using",
    "While",
    "With",
    "Do",
    "Loop",
    "For",
    "Next",
    "Property",
    "Continue",
    "AddHandler",
    "RemoveHandler",
    "Event",
    "RaiseEvent",
    "SyncLock"
  ],
  symbols: /[=><!~?;\.,:&|+\-*\/\^%]+/,
  integersuffix: /U?[DI%L&S@]?/,
  floatsuffix: /[R#F!]?/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      [/next(?!\w)/, { token: "keyword.tag-for" }],
      [/loop(?!\w)/, { token: "keyword.tag-do" }],
      [
        /end\s+(?!for|do)(addhandler|class|enum|event|function|get|if|interface|module|namespace|operator|property|raiseevent|removehandler|select|set|structure|sub|synclock|try|while|with|using)/,
        { token: "keyword.tag-$1" }
      ],
      [
        /[a-zA-Z_]\w*/,
        {
          cases: {
            "@tagwords": { token: "keyword.tag-$0" },
            "@keywords": { token: "keyword.$0" },
            "@default": "identifier"
          }
        }
      ],
      [/^\s*#\w+/, "keyword"],
      [/\d*\d+e([\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/\d*\.\d+(e[\-+]?\d+)?(@floatsuffix)/, "number.float"],
      [/&H[0-9a-f]+(@integersuffix)/, "number.hex"],
      [/&0[0-7]+(@integersuffix)/, "number.octal"],
      [/\d+(@integersuffix)/, "number"],
      [/#.*#/, "number"],
      [/[{}()\[\]]/, "@brackets"],
      [/@symbols/, "delimiter"],
      [/["\u201c\u201d]/, { token: "string.quote", next: "@string" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/(\'|REM(?!\w)).*$/, "comment"]
    ],
    string: [
      [/[^"\u201c\u201d]+/, "string"],
      [/["\u201c\u201d]{2}/, "string.escape"],
      [/["\u201c\u201d]C?/, { token: "string.quote", next: "@pop" }]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/wgsl/wgsl.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/wgsl/wgsl.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

// src/basic-languages/wgsl/wgsl.ts
var conf = {
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "[", close: "]" },
    { open: "{", close: "}" },
    { open: "(", close: ")" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" }
  ]
};
function qw(str) {
  let result = [];
  const words = str.split(/\t+|\r+|\n+| +/);
  for (let i = 0; i < words.length; ++i) {
    if (words[i].length > 0) {
      result.push(words[i]);
    }
  }
  return result;
}
var atoms = qw("true false");
var keywords = qw(`
			  alias
			  break
			  case
			  const
			  const_assert
			  continue
			  continuing
			  default
			  diagnostic
			  discard
			  else
			  enable
			  fn
			  for
			  if
			  let
			  loop
			  override
			  requires
			  return
			  struct
			  switch
			  var
			  while
			  `);
var reserved = qw(`
			  NULL
			  Self
			  abstract
			  active
			  alignas
			  alignof
			  as
			  asm
			  asm_fragment
			  async
			  attribute
			  auto
			  await
			  become
			  binding_array
			  cast
			  catch
			  class
			  co_await
			  co_return
			  co_yield
			  coherent
			  column_major
			  common
			  compile
			  compile_fragment
			  concept
			  const_cast
			  consteval
			  constexpr
			  constinit
			  crate
			  debugger
			  decltype
			  delete
			  demote
			  demote_to_helper
			  do
			  dynamic_cast
			  enum
			  explicit
			  export
			  extends
			  extern
			  external
			  fallthrough
			  filter
			  final
			  finally
			  friend
			  from
			  fxgroup
			  get
			  goto
			  groupshared
			  highp
			  impl
			  implements
			  import
			  inline
			  instanceof
			  interface
			  layout
			  lowp
			  macro
			  macro_rules
			  match
			  mediump
			  meta
			  mod
			  module
			  move
			  mut
			  mutable
			  namespace
			  new
			  nil
			  noexcept
			  noinline
			  nointerpolation
			  noperspective
			  null
			  nullptr
			  of
			  operator
			  package
			  packoffset
			  partition
			  pass
			  patch
			  pixelfragment
			  precise
			  precision
			  premerge
			  priv
			  protected
			  pub
			  public
			  readonly
			  ref
			  regardless
			  register
			  reinterpret_cast
			  require
			  resource
			  restrict
			  self
			  set
			  shared
			  sizeof
			  smooth
			  snorm
			  static
			  static_assert
			  static_cast
			  std
			  subroutine
			  super
			  target
			  template
			  this
			  thread_local
			  throw
			  trait
			  try
			  type
			  typedef
			  typeid
			  typename
			  typeof
			  union
			  unless
			  unorm
			  unsafe
			  unsized
			  use
			  using
			  varying
			  virtual
			  volatile
			  wgsl
			  where
			  with
			  writeonly
			  yield
			  `);
var predeclared_enums = qw(`
		read write read_write
		function private workgroup uniform storage
		perspective linear flat
		center centroid sample
		vertex_index instance_index position front_facing frag_depth
			local_invocation_id local_invocation_index
			global_invocation_id workgroup_id num_workgroups
			sample_index sample_mask
		rgba8unorm
		rgba8snorm
		rgba8uint
		rgba8sint
		rgba16uint
		rgba16sint
		rgba16float
		r32uint
		r32sint
		r32float
		rg32uint
		rg32sint
		rg32float
		rgba32uint
		rgba32sint
		rgba32float
		bgra8unorm
`);
var predeclared_types = qw(`
		bool
		f16
		f32
		i32
		sampler sampler_comparison
		texture_depth_2d
		texture_depth_2d_array
		texture_depth_cube
		texture_depth_cube_array
		texture_depth_multisampled_2d
		texture_external
		texture_external
		u32
		`);
var predeclared_type_generators = qw(`
		array
		atomic
		mat2x2
		mat2x3
		mat2x4
		mat3x2
		mat3x3
		mat3x4
		mat4x2
		mat4x3
		mat4x4
		ptr
		texture_1d
		texture_2d
		texture_2d_array
		texture_3d
		texture_cube
		texture_cube_array
		texture_multisampled_2d
		texture_storage_1d
		texture_storage_2d
		texture_storage_2d_array
		texture_storage_3d
		vec2
		vec3
		vec4
		`);
var predeclared_type_aliases = qw(`
		vec2i vec3i vec4i
		vec2u vec3u vec4u
		vec2f vec3f vec4f
		vec2h vec3h vec4h
		mat2x2f mat2x3f mat2x4f
		mat3x2f mat3x3f mat3x4f
		mat4x2f mat4x3f mat4x4f
		mat2x2h mat2x3h mat2x4h
		mat3x2h mat3x3h mat3x4h
		mat4x2h mat4x3h mat4x4h
		`);
var predeclared_intrinsics = qw(`
  bitcast all any select arrayLength abs acos acosh asin asinh atan atanh atan2
  ceil clamp cos cosh countLeadingZeros countOneBits countTrailingZeros cross
  degrees determinant distance dot exp exp2 extractBits faceForward firstLeadingBit
  firstTrailingBit floor fma fract frexp inverseBits inverseSqrt ldexp length
  log log2 max min mix modf normalize pow quantizeToF16 radians reflect refract
  reverseBits round saturate sign sin sinh smoothstep sqrt step tan tanh transpose
  trunc dpdx dpdxCoarse dpdxFine dpdy dpdyCoarse dpdyFine fwidth fwidthCoarse fwidthFine
  textureDimensions textureGather textureGatherCompare textureLoad textureNumLayers
  textureNumLevels textureNumSamples textureSample textureSampleBias textureSampleCompare
  textureSampleCompareLevel textureSampleGrad textureSampleLevel textureSampleBaseClampToEdge
  textureStore atomicLoad atomicStore atomicAdd atomicSub atomicMax atomicMin
  atomicAnd atomicOr atomicXor atomicExchange atomicCompareExchangeWeak pack4x8snorm
  pack4x8unorm pack2x16snorm pack2x16unorm pack2x16float unpack4x8snorm unpack4x8unorm
  unpack2x16snorm unpack2x16unorm unpack2x16float storageBarrier workgroupBarrier
  workgroupUniformLoad
`);
var operators = qw(`
					 &
					 &&
					 ->
					 /
					 =
					 ==
					 !=
					 >
					 >=
					 <
					 <=
					 %
					 -
					 --
					 +
					 ++
					 |
					 ||
					 *
					 <<
					 >>
					 +=
					 -=
					 *=
					 /=
					 %=
					 &=
					 |=
					 ^=
					 >>=
					 <<=
					 `);
var directive_re = /enable|requires|diagnostic/;
var ident_re = /[_\p{XID_Start}]\p{XID_Continue}*/u;
var predefined_token = "variable.predefined";
var language = {
  tokenPostfix: ".wgsl",
  defaultToken: "invalid",
  unicode: true,
  atoms,
  keywords,
  reserved,
  predeclared_enums,
  predeclared_types,
  predeclared_type_generators,
  predeclared_type_aliases,
  predeclared_intrinsics,
  operators,
  symbols: /[!%&*+\-\.\/:;<=>^|_~,]+/,
  tokenizer: {
    root: [
      [directive_re, "keyword", "@directive"],
      [
        ident_re,
        {
          cases: {
            "@atoms": predefined_token,
            "@keywords": "keyword",
            "@reserved": "invalid",
            "@predeclared_enums": predefined_token,
            "@predeclared_types": predefined_token,
            "@predeclared_type_generators": predefined_token,
            "@predeclared_type_aliases": predefined_token,
            "@predeclared_intrinsics": predefined_token,
            "@default": "identifier"
          }
        }
      ],
      { include: "@commentOrSpace" },
      { include: "@numbers" },
      [/[{}()\[\]]/, "@brackets"],
      ["@", "annotation", "@attribute"],
      [
        /@symbols/,
        {
          cases: {
            "@operators": "operator",
            "@default": "delimiter"
          }
        }
      ],
      [/./, "invalid"]
    ],
    commentOrSpace: [
      [/\s+/, "white"],
      [/\/\*/, "comment", "@blockComment"],
      [/\/\/.*$/, "comment"]
    ],
    blockComment: [
      [/[^\/*]+/, "comment"],
      [/\/\*/, "comment", "@push"],
      [/\*\//, "comment", "@pop"],
      [/[\/*]/, "comment"]
    ],
    attribute: [
      { include: "@commentOrSpace" },
      [/\w+/, "annotation", "@pop"]
    ],
    directive: [
      { include: "@commentOrSpace" },
      [/[()]/, "@brackets"],
      [/,/, "delimiter"],
      [ident_re, "meta.content"],
      [/;/, "delimiter", "@pop"]
    ],
    numbers: [
      [/0[fh]/, "number.float"],
      [/[1-9][0-9]*[fh]/, "number.float"],
      [/[0-9]*\.[0-9]+([eE][+-]?[0-9]+)?[fh]?/, "number.float"],
      [/[0-9]+\.[0-9]*([eE][+-]?[0-9]+)?[fh]?/, "number.float"],
      [/[0-9]+[eE][+-]?[0-9]+[fh]?/, "number.float"],
      [/0[xX][0-9a-fA-F]*\.[0-9a-fA-F]+(?:[pP][+-]?[0-9]+[fh]?)?/, "number.hex"],
      [/0[xX][0-9a-fA-F]+\.[0-9a-fA-F]*(?:[pP][+-]?[0-9]+[fh]?)?/, "number.hex"],
      [/0[xX][0-9a-fA-F]+[pP][+-]?[0-9]+[fh]?/, "number.hex"],
      [/0[xX][0-9a-fA-F]+[iu]?/, "number.hex"],
      [/[1-9][0-9]*[iu]?/, "number"],
      [/0[iu]?/, "number"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/xml/xml.js":
/*!**********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/xml/xml.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/xml/xml.ts
var conf = {
  comments: {
    blockComment: ["<!--", "-->"]
  },
  brackets: [["<", ">"]],
  autoClosingPairs: [
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: '"', close: '"' }
  ],
  surroundingPairs: [
    { open: "<", close: ">" },
    { open: "'", close: "'" },
    { open: '"', close: '"' }
  ],
  onEnterRules: [
    {
      beforeText: new RegExp(`<([_:\\w][_:\\w-.\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      afterText: /^<\/([_:\w][_:\w-.\d]*)\s*>$/i,
      action: {
        indentAction: monaco_editor_core_exports.languages.IndentAction.IndentOutdent
      }
    },
    {
      beforeText: new RegExp(`<(\\w[\\w\\d]*)([^/>]*(?!/)>)[^<]*$`, "i"),
      action: { indentAction: monaco_editor_core_exports.languages.IndentAction.Indent }
    }
  ]
};
var language = {
  defaultToken: "",
  tokenPostfix: ".xml",
  ignoreCase: true,
  qualifiedName: /(?:[\w\.\-]+:)?[\w\.\-]+/,
  tokenizer: {
    root: [
      [/[^<&]+/, ""],
      { include: "@whitespace" },
      [/(<)(@qualifiedName)/, [{ token: "delimiter" }, { token: "tag", next: "@tag" }]],
      [
        /(<\/)(@qualifiedName)(\s*)(>)/,
        [{ token: "delimiter" }, { token: "tag" }, "", { token: "delimiter" }]
      ],
      [/(<\?)(@qualifiedName)/, [{ token: "delimiter" }, { token: "metatag", next: "@tag" }]],
      [/(<\!)(@qualifiedName)/, [{ token: "delimiter" }, { token: "metatag", next: "@tag" }]],
      [/<\!\[CDATA\[/, { token: "delimiter.cdata", next: "@cdata" }],
      [/&\w+;/, "string.escape"]
    ],
    cdata: [
      [/[^\]]+/, ""],
      [/\]\]>/, { token: "delimiter.cdata", next: "@pop" }],
      [/\]/, ""]
    ],
    tag: [
      [/[ \t\r\n]+/, ""],
      [/(@qualifiedName)(\s*=\s*)("[^"]*"|'[^']*')/, ["attribute.name", "", "attribute.value"]],
      [
        /(@qualifiedName)(\s*=\s*)("[^">?\/]*|'[^'>?\/]*)(?=[\?\/]\>)/,
        ["attribute.name", "", "attribute.value"]
      ],
      [/(@qualifiedName)(\s*=\s*)("[^">]*|'[^'>]*)/, ["attribute.name", "", "attribute.value"]],
      [/@qualifiedName/, "attribute.name"],
      [/\?>/, { token: "delimiter", next: "@pop" }],
      [/(\/)(>)/, [{ token: "tag" }, { token: "delimiter", next: "@pop" }]],
      [/>/, { token: "delimiter", next: "@pop" }]
    ],
    whitespace: [
      [/[ \t\r\n]+/, ""],
      [/<!--/, { token: "comment", next: "@comment" }]
    ],
    comment: [
      [/[^<\-]+/, "comment.content"],
      [/-->/, { token: "comment", next: "@pop" }],
      [/<!--/, "comment.content.invalid"],
      [/[<\-]/, "comment.content"]
    ]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/basic-languages/yaml/yaml.js":
/*!************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/basic-languages/yaml/yaml.js ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   conf: function() { return /* binding */ conf; },
/* harmony export */   language: function() { return /* binding */ language; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/basic-languages/yaml/yaml.ts
var conf = {
  comments: {
    lineComment: "#"
  },
  brackets: [
    ["{", "}"],
    ["[", "]"],
    ["(", ")"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  surroundingPairs: [
    { open: "{", close: "}" },
    { open: "[", close: "]" },
    { open: "(", close: ")" },
    { open: '"', close: '"' },
    { open: "'", close: "'" }
  ],
  folding: {
    offSide: true
  },
  onEnterRules: [
    {
      beforeText: /:\s*$/,
      action: {
        indentAction: monaco_editor_core_exports.languages.IndentAction.Indent
      }
    }
  ]
};
var language = {
  tokenPostfix: ".yaml",
  brackets: [
    { token: "delimiter.bracket", open: "{", close: "}" },
    { token: "delimiter.square", open: "[", close: "]" }
  ],
  keywords: ["true", "True", "TRUE", "false", "False", "FALSE", "null", "Null", "Null", "~"],
  numberInteger: /(?:0|[+-]?[0-9]+)/,
  numberFloat: /(?:0|[+-]?[0-9]+)(?:\.[0-9]+)?(?:e[-+][1-9][0-9]*)?/,
  numberOctal: /0o[0-7]+/,
  numberHex: /0x[0-9a-fA-F]+/,
  numberInfinity: /[+-]?\.(?:inf|Inf|INF)/,
  numberNaN: /\.(?:nan|Nan|NAN)/,
  numberDate: /\d{4}-\d\d-\d\d([Tt ]\d\d:\d\d:\d\d(\.\d+)?(( ?[+-]\d\d?(:\d\d)?)|Z)?)?/,
  escapes: /\\(?:[btnfr\\"']|[0-7][0-7]?|[0-3][0-7]{2})/,
  tokenizer: {
    root: [
      { include: "@whitespace" },
      { include: "@comment" },
      [/%[^ ]+.*$/, "meta.directive"],
      [/---/, "operators.directivesEnd"],
      [/\.{3}/, "operators.documentEnd"],
      [/[-?:](?= )/, "operators"],
      { include: "@anchor" },
      { include: "@tagHandle" },
      { include: "@flowCollections" },
      { include: "@blockStyle" },
      [/@numberInteger(?![ \t]*\S+)/, "number"],
      [/@numberFloat(?![ \t]*\S+)/, "number.float"],
      [/@numberOctal(?![ \t]*\S+)/, "number.octal"],
      [/@numberHex(?![ \t]*\S+)/, "number.hex"],
      [/@numberInfinity(?![ \t]*\S+)/, "number.infinity"],
      [/@numberNaN(?![ \t]*\S+)/, "number.nan"],
      [/@numberDate(?![ \t]*\S+)/, "number.date"],
      [/(".*?"|'.*?'|[^#'"]*?)([ \t]*)(:)( |$)/, ["type", "white", "operators", "white"]],
      { include: "@flowScalars" },
      [
        /.+?(?=(\s+#|$))/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "string"
          }
        }
      ]
    ],
    object: [
      { include: "@whitespace" },
      { include: "@comment" },
      [/\}/, "@brackets", "@pop"],
      [/,/, "delimiter.comma"],
      [/:(?= )/, "operators"],
      [/(?:".*?"|'.*?'|[^,\{\[]+?)(?=: )/, "type"],
      { include: "@flowCollections" },
      { include: "@flowScalars" },
      { include: "@tagHandle" },
      { include: "@anchor" },
      { include: "@flowNumber" },
      [
        /[^\},]+/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "string"
          }
        }
      ]
    ],
    array: [
      { include: "@whitespace" },
      { include: "@comment" },
      [/\]/, "@brackets", "@pop"],
      [/,/, "delimiter.comma"],
      { include: "@flowCollections" },
      { include: "@flowScalars" },
      { include: "@tagHandle" },
      { include: "@anchor" },
      { include: "@flowNumber" },
      [
        /[^\],]+/,
        {
          cases: {
            "@keywords": "keyword",
            "@default": "string"
          }
        }
      ]
    ],
    multiString: [[/^( +).+$/, "string", "@multiStringContinued.$1"]],
    multiStringContinued: [
      [
        /^( *).+$/,
        {
          cases: {
            "$1==$S2": "string",
            "@default": { token: "@rematch", next: "@popall" }
          }
        }
      ]
    ],
    whitespace: [[/[ \t\r\n]+/, "white"]],
    comment: [[/#.*$/, "comment"]],
    flowCollections: [
      [/\[/, "@brackets", "@array"],
      [/\{/, "@brackets", "@object"]
    ],
    flowScalars: [
      [/"([^"\\]|\\.)*$/, "string.invalid"],
      [/'([^'\\]|\\.)*$/, "string.invalid"],
      [/'[^']*'/, "string"],
      [/"/, "string", "@doubleQuotedString"]
    ],
    doubleQuotedString: [
      [/[^\\"]+/, "string"],
      [/@escapes/, "string.escape"],
      [/\\./, "string.escape.invalid"],
      [/"/, "string", "@pop"]
    ],
    blockStyle: [[/[>|][0-9]*[+-]?$/, "operators", "@multiString"]],
    flowNumber: [
      [/@numberInteger(?=[ \t]*[,\]\}])/, "number"],
      [/@numberFloat(?=[ \t]*[,\]\}])/, "number.float"],
      [/@numberOctal(?=[ \t]*[,\]\}])/, "number.octal"],
      [/@numberHex(?=[ \t]*[,\]\}])/, "number.hex"],
      [/@numberInfinity(?=[ \t]*[,\]\}])/, "number.infinity"],
      [/@numberNaN(?=[ \t]*[,\]\}])/, "number.nan"],
      [/@numberDate(?=[ \t]*[,\]\}])/, "number.date"]
    ],
    tagHandle: [[/\![^ ]*/, "tag"]],
    anchor: [[/[&*][^ ]+/, "namespace"]]
  }
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/language/css/cssMode.js":
/*!*******************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/language/css/cssMode.js ***!
  \*******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   CompletionAdapter: function() { return /* binding */ CompletionAdapter; },
/* harmony export */   DefinitionAdapter: function() { return /* binding */ DefinitionAdapter; },
/* harmony export */   DiagnosticsAdapter: function() { return /* binding */ DiagnosticsAdapter; },
/* harmony export */   DocumentColorAdapter: function() { return /* binding */ DocumentColorAdapter; },
/* harmony export */   DocumentFormattingEditProvider: function() { return /* binding */ DocumentFormattingEditProvider; },
/* harmony export */   DocumentHighlightAdapter: function() { return /* binding */ DocumentHighlightAdapter; },
/* harmony export */   DocumentLinkAdapter: function() { return /* binding */ DocumentLinkAdapter; },
/* harmony export */   DocumentRangeFormattingEditProvider: function() { return /* binding */ DocumentRangeFormattingEditProvider; },
/* harmony export */   DocumentSymbolAdapter: function() { return /* binding */ DocumentSymbolAdapter; },
/* harmony export */   FoldingRangeAdapter: function() { return /* binding */ FoldingRangeAdapter; },
/* harmony export */   HoverAdapter: function() { return /* binding */ HoverAdapter; },
/* harmony export */   ReferenceAdapter: function() { return /* binding */ ReferenceAdapter; },
/* harmony export */   RenameAdapter: function() { return /* binding */ RenameAdapter; },
/* harmony export */   SelectionRangeAdapter: function() { return /* binding */ SelectionRangeAdapter; },
/* harmony export */   WorkerManager: function() { return /* binding */ WorkerManager; },
/* harmony export */   fromPosition: function() { return /* binding */ fromPosition; },
/* harmony export */   fromRange: function() { return /* binding */ fromRange; },
/* harmony export */   setupMode: function() { return /* binding */ setupMode; },
/* harmony export */   toRange: function() { return /* binding */ toRange; },
/* harmony export */   toTextEdit: function() { return /* binding */ toTextEdit; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/language/css/workerManager.ts
var STOP_WHEN_IDLE_FOR = 2 * 60 * 1e3;
var WorkerManager = class {
  _defaults;
  _idleCheckInterval;
  _lastUsedTime;
  _configChangeListener;
  _worker;
  _client;
  constructor(defaults) {
    this._defaults = defaults;
    this._worker = null;
    this._client = null;
    this._idleCheckInterval = window.setInterval(() => this._checkIfIdle(), 30 * 1e3);
    this._lastUsedTime = 0;
    this._configChangeListener = this._defaults.onDidChange(() => this._stopWorker());
  }
  _stopWorker() {
    if (this._worker) {
      this._worker.dispose();
      this._worker = null;
    }
    this._client = null;
  }
  dispose() {
    clearInterval(this._idleCheckInterval);
    this._configChangeListener.dispose();
    this._stopWorker();
  }
  _checkIfIdle() {
    if (!this._worker) {
      return;
    }
    let timePassedSinceLastUsed = Date.now() - this._lastUsedTime;
    if (timePassedSinceLastUsed > STOP_WHEN_IDLE_FOR) {
      this._stopWorker();
    }
  }
  _getClient() {
    this._lastUsedTime = Date.now();
    if (!this._client) {
      this._worker = monaco_editor_core_exports.editor.createWebWorker({
        moduleId: "vs/language/css/cssWorker",
        label: this._defaults.languageId,
        createData: {
          options: this._defaults.options,
          languageId: this._defaults.languageId
        }
      });
      this._client = this._worker.getProxy();
    }
    return this._client;
  }
  getLanguageServiceWorker(...resources) {
    let _client;
    return this._getClient().then((client) => {
      _client = client;
    }).then((_) => {
      if (this._worker) {
        return this._worker.withSyncedResources(resources);
      }
    }).then((_) => _client);
  }
};

// node_modules/vscode-languageserver-types/lib/esm/main.js
var integer;
(function(integer2) {
  integer2.MIN_VALUE = -2147483648;
  integer2.MAX_VALUE = 2147483647;
})(integer || (integer = {}));
var uinteger;
(function(uinteger2) {
  uinteger2.MIN_VALUE = 0;
  uinteger2.MAX_VALUE = 2147483647;
})(uinteger || (uinteger = {}));
var Position;
(function(Position3) {
  function create(line, character) {
    if (line === Number.MAX_VALUE) {
      line = uinteger.MAX_VALUE;
    }
    if (character === Number.MAX_VALUE) {
      character = uinteger.MAX_VALUE;
    }
    return { line, character };
  }
  Position3.create = create;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Is.uinteger(candidate.line) && Is.uinteger(candidate.character);
  }
  Position3.is = is;
})(Position || (Position = {}));
var Range;
(function(Range3) {
  function create(one, two, three, four) {
    if (Is.uinteger(one) && Is.uinteger(two) && Is.uinteger(three) && Is.uinteger(four)) {
      return { start: Position.create(one, two), end: Position.create(three, four) };
    } else if (Position.is(one) && Position.is(two)) {
      return { start: one, end: two };
    } else {
      throw new Error("Range#create called with invalid arguments[" + one + ", " + two + ", " + three + ", " + four + "]");
    }
  }
  Range3.create = create;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Position.is(candidate.start) && Position.is(candidate.end);
  }
  Range3.is = is;
})(Range || (Range = {}));
var Location;
(function(Location2) {
  function create(uri, range) {
    return { uri, range };
  }
  Location2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.string(candidate.uri) || Is.undefined(candidate.uri));
  }
  Location2.is = is;
})(Location || (Location = {}));
var LocationLink;
(function(LocationLink2) {
  function create(targetUri, targetRange, targetSelectionRange, originSelectionRange) {
    return { targetUri, targetRange, targetSelectionRange, originSelectionRange };
  }
  LocationLink2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.targetRange) && Is.string(candidate.targetUri) && (Range.is(candidate.targetSelectionRange) || Is.undefined(candidate.targetSelectionRange)) && (Range.is(candidate.originSelectionRange) || Is.undefined(candidate.originSelectionRange));
  }
  LocationLink2.is = is;
})(LocationLink || (LocationLink = {}));
var Color;
(function(Color2) {
  function create(red, green, blue, alpha) {
    return {
      red,
      green,
      blue,
      alpha
    };
  }
  Color2.create = create;
  function is(value) {
    var candidate = value;
    return Is.numberRange(candidate.red, 0, 1) && Is.numberRange(candidate.green, 0, 1) && Is.numberRange(candidate.blue, 0, 1) && Is.numberRange(candidate.alpha, 0, 1);
  }
  Color2.is = is;
})(Color || (Color = {}));
var ColorInformation;
(function(ColorInformation2) {
  function create(range, color) {
    return {
      range,
      color
    };
  }
  ColorInformation2.create = create;
  function is(value) {
    var candidate = value;
    return Range.is(candidate.range) && Color.is(candidate.color);
  }
  ColorInformation2.is = is;
})(ColorInformation || (ColorInformation = {}));
var ColorPresentation;
(function(ColorPresentation2) {
  function create(label, textEdit, additionalTextEdits) {
    return {
      label,
      textEdit,
      additionalTextEdits
    };
  }
  ColorPresentation2.create = create;
  function is(value) {
    var candidate = value;
    return Is.string(candidate.label) && (Is.undefined(candidate.textEdit) || TextEdit.is(candidate)) && (Is.undefined(candidate.additionalTextEdits) || Is.typedArray(candidate.additionalTextEdits, TextEdit.is));
  }
  ColorPresentation2.is = is;
})(ColorPresentation || (ColorPresentation = {}));
var FoldingRangeKind;
(function(FoldingRangeKind2) {
  FoldingRangeKind2["Comment"] = "comment";
  FoldingRangeKind2["Imports"] = "imports";
  FoldingRangeKind2["Region"] = "region";
})(FoldingRangeKind || (FoldingRangeKind = {}));
var FoldingRange;
(function(FoldingRange2) {
  function create(startLine, endLine, startCharacter, endCharacter, kind) {
    var result = {
      startLine,
      endLine
    };
    if (Is.defined(startCharacter)) {
      result.startCharacter = startCharacter;
    }
    if (Is.defined(endCharacter)) {
      result.endCharacter = endCharacter;
    }
    if (Is.defined(kind)) {
      result.kind = kind;
    }
    return result;
  }
  FoldingRange2.create = create;
  function is(value) {
    var candidate = value;
    return Is.uinteger(candidate.startLine) && Is.uinteger(candidate.startLine) && (Is.undefined(candidate.startCharacter) || Is.uinteger(candidate.startCharacter)) && (Is.undefined(candidate.endCharacter) || Is.uinteger(candidate.endCharacter)) && (Is.undefined(candidate.kind) || Is.string(candidate.kind));
  }
  FoldingRange2.is = is;
})(FoldingRange || (FoldingRange = {}));
var DiagnosticRelatedInformation;
(function(DiagnosticRelatedInformation2) {
  function create(location, message) {
    return {
      location,
      message
    };
  }
  DiagnosticRelatedInformation2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Location.is(candidate.location) && Is.string(candidate.message);
  }
  DiagnosticRelatedInformation2.is = is;
})(DiagnosticRelatedInformation || (DiagnosticRelatedInformation = {}));
var DiagnosticSeverity;
(function(DiagnosticSeverity2) {
  DiagnosticSeverity2.Error = 1;
  DiagnosticSeverity2.Warning = 2;
  DiagnosticSeverity2.Information = 3;
  DiagnosticSeverity2.Hint = 4;
})(DiagnosticSeverity || (DiagnosticSeverity = {}));
var DiagnosticTag;
(function(DiagnosticTag2) {
  DiagnosticTag2.Unnecessary = 1;
  DiagnosticTag2.Deprecated = 2;
})(DiagnosticTag || (DiagnosticTag = {}));
var CodeDescription;
(function(CodeDescription2) {
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && candidate !== null && Is.string(candidate.href);
  }
  CodeDescription2.is = is;
})(CodeDescription || (CodeDescription = {}));
var Diagnostic;
(function(Diagnostic2) {
  function create(range, message, severity, code, source, relatedInformation) {
    var result = { range, message };
    if (Is.defined(severity)) {
      result.severity = severity;
    }
    if (Is.defined(code)) {
      result.code = code;
    }
    if (Is.defined(source)) {
      result.source = source;
    }
    if (Is.defined(relatedInformation)) {
      result.relatedInformation = relatedInformation;
    }
    return result;
  }
  Diagnostic2.create = create;
  function is(value) {
    var _a;
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && Is.string(candidate.message) && (Is.number(candidate.severity) || Is.undefined(candidate.severity)) && (Is.integer(candidate.code) || Is.string(candidate.code) || Is.undefined(candidate.code)) && (Is.undefined(candidate.codeDescription) || Is.string((_a = candidate.codeDescription) === null || _a === void 0 ? void 0 : _a.href)) && (Is.string(candidate.source) || Is.undefined(candidate.source)) && (Is.undefined(candidate.relatedInformation) || Is.typedArray(candidate.relatedInformation, DiagnosticRelatedInformation.is));
  }
  Diagnostic2.is = is;
})(Diagnostic || (Diagnostic = {}));
var Command;
(function(Command2) {
  function create(title, command) {
    var args = [];
    for (var _i = 2; _i < arguments.length; _i++) {
      args[_i - 2] = arguments[_i];
    }
    var result = { title, command };
    if (Is.defined(args) && args.length > 0) {
      result.arguments = args;
    }
    return result;
  }
  Command2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.title) && Is.string(candidate.command);
  }
  Command2.is = is;
})(Command || (Command = {}));
var TextEdit;
(function(TextEdit2) {
  function replace(range, newText) {
    return { range, newText };
  }
  TextEdit2.replace = replace;
  function insert(position, newText) {
    return { range: { start: position, end: position }, newText };
  }
  TextEdit2.insert = insert;
  function del(range) {
    return { range, newText: "" };
  }
  TextEdit2.del = del;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Is.string(candidate.newText) && Range.is(candidate.range);
  }
  TextEdit2.is = is;
})(TextEdit || (TextEdit = {}));
var ChangeAnnotation;
(function(ChangeAnnotation2) {
  function create(label, needsConfirmation, description) {
    var result = { label };
    if (needsConfirmation !== void 0) {
      result.needsConfirmation = needsConfirmation;
    }
    if (description !== void 0) {
      result.description = description;
    }
    return result;
  }
  ChangeAnnotation2.create = create;
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && Is.objectLiteral(candidate) && Is.string(candidate.label) && (Is.boolean(candidate.needsConfirmation) || candidate.needsConfirmation === void 0) && (Is.string(candidate.description) || candidate.description === void 0);
  }
  ChangeAnnotation2.is = is;
})(ChangeAnnotation || (ChangeAnnotation = {}));
var ChangeAnnotationIdentifier;
(function(ChangeAnnotationIdentifier2) {
  function is(value) {
    var candidate = value;
    return typeof candidate === "string";
  }
  ChangeAnnotationIdentifier2.is = is;
})(ChangeAnnotationIdentifier || (ChangeAnnotationIdentifier = {}));
var AnnotatedTextEdit;
(function(AnnotatedTextEdit2) {
  function replace(range, newText, annotation) {
    return { range, newText, annotationId: annotation };
  }
  AnnotatedTextEdit2.replace = replace;
  function insert(position, newText, annotation) {
    return { range: { start: position, end: position }, newText, annotationId: annotation };
  }
  AnnotatedTextEdit2.insert = insert;
  function del(range, annotation) {
    return { range, newText: "", annotationId: annotation };
  }
  AnnotatedTextEdit2.del = del;
  function is(value) {
    var candidate = value;
    return TextEdit.is(candidate) && (ChangeAnnotation.is(candidate.annotationId) || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  AnnotatedTextEdit2.is = is;
})(AnnotatedTextEdit || (AnnotatedTextEdit = {}));
var TextDocumentEdit;
(function(TextDocumentEdit2) {
  function create(textDocument, edits) {
    return { textDocument, edits };
  }
  TextDocumentEdit2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && OptionalVersionedTextDocumentIdentifier.is(candidate.textDocument) && Array.isArray(candidate.edits);
  }
  TextDocumentEdit2.is = is;
})(TextDocumentEdit || (TextDocumentEdit = {}));
var CreateFile;
(function(CreateFile2) {
  function create(uri, options, annotation) {
    var result = {
      kind: "create",
      uri
    };
    if (options !== void 0 && (options.overwrite !== void 0 || options.ignoreIfExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  CreateFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "create" && Is.string(candidate.uri) && (candidate.options === void 0 || (candidate.options.overwrite === void 0 || Is.boolean(candidate.options.overwrite)) && (candidate.options.ignoreIfExists === void 0 || Is.boolean(candidate.options.ignoreIfExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  CreateFile2.is = is;
})(CreateFile || (CreateFile = {}));
var RenameFile;
(function(RenameFile2) {
  function create(oldUri, newUri, options, annotation) {
    var result = {
      kind: "rename",
      oldUri,
      newUri
    };
    if (options !== void 0 && (options.overwrite !== void 0 || options.ignoreIfExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  RenameFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "rename" && Is.string(candidate.oldUri) && Is.string(candidate.newUri) && (candidate.options === void 0 || (candidate.options.overwrite === void 0 || Is.boolean(candidate.options.overwrite)) && (candidate.options.ignoreIfExists === void 0 || Is.boolean(candidate.options.ignoreIfExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  RenameFile2.is = is;
})(RenameFile || (RenameFile = {}));
var DeleteFile;
(function(DeleteFile2) {
  function create(uri, options, annotation) {
    var result = {
      kind: "delete",
      uri
    };
    if (options !== void 0 && (options.recursive !== void 0 || options.ignoreIfNotExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  DeleteFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "delete" && Is.string(candidate.uri) && (candidate.options === void 0 || (candidate.options.recursive === void 0 || Is.boolean(candidate.options.recursive)) && (candidate.options.ignoreIfNotExists === void 0 || Is.boolean(candidate.options.ignoreIfNotExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  DeleteFile2.is = is;
})(DeleteFile || (DeleteFile = {}));
var WorkspaceEdit;
(function(WorkspaceEdit2) {
  function is(value) {
    var candidate = value;
    return candidate && (candidate.changes !== void 0 || candidate.documentChanges !== void 0) && (candidate.documentChanges === void 0 || candidate.documentChanges.every(function(change) {
      if (Is.string(change.kind)) {
        return CreateFile.is(change) || RenameFile.is(change) || DeleteFile.is(change);
      } else {
        return TextDocumentEdit.is(change);
      }
    }));
  }
  WorkspaceEdit2.is = is;
})(WorkspaceEdit || (WorkspaceEdit = {}));
var TextEditChangeImpl = function() {
  function TextEditChangeImpl2(edits, changeAnnotations) {
    this.edits = edits;
    this.changeAnnotations = changeAnnotations;
  }
  TextEditChangeImpl2.prototype.insert = function(position, newText, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.insert(position, newText);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.insert(position, newText, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.insert(position, newText, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.replace = function(range, newText, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.replace(range, newText);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.replace(range, newText, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.replace(range, newText, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.delete = function(range, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.del(range);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.del(range, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.del(range, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.add = function(edit) {
    this.edits.push(edit);
  };
  TextEditChangeImpl2.prototype.all = function() {
    return this.edits;
  };
  TextEditChangeImpl2.prototype.clear = function() {
    this.edits.splice(0, this.edits.length);
  };
  TextEditChangeImpl2.prototype.assertChangeAnnotations = function(value) {
    if (value === void 0) {
      throw new Error("Text edit change is not configured to manage change annotations.");
    }
  };
  return TextEditChangeImpl2;
}();
var ChangeAnnotations = function() {
  function ChangeAnnotations2(annotations) {
    this._annotations = annotations === void 0 ? /* @__PURE__ */ Object.create(null) : annotations;
    this._counter = 0;
    this._size = 0;
  }
  ChangeAnnotations2.prototype.all = function() {
    return this._annotations;
  };
  Object.defineProperty(ChangeAnnotations2.prototype, "size", {
    get: function() {
      return this._size;
    },
    enumerable: false,
    configurable: true
  });
  ChangeAnnotations2.prototype.manage = function(idOrAnnotation, annotation) {
    var id;
    if (ChangeAnnotationIdentifier.is(idOrAnnotation)) {
      id = idOrAnnotation;
    } else {
      id = this.nextId();
      annotation = idOrAnnotation;
    }
    if (this._annotations[id] !== void 0) {
      throw new Error("Id " + id + " is already in use.");
    }
    if (annotation === void 0) {
      throw new Error("No annotation provided for id " + id);
    }
    this._annotations[id] = annotation;
    this._size++;
    return id;
  };
  ChangeAnnotations2.prototype.nextId = function() {
    this._counter++;
    return this._counter.toString();
  };
  return ChangeAnnotations2;
}();
var WorkspaceChange = function() {
  function WorkspaceChange2(workspaceEdit) {
    var _this = this;
    this._textEditChanges = /* @__PURE__ */ Object.create(null);
    if (workspaceEdit !== void 0) {
      this._workspaceEdit = workspaceEdit;
      if (workspaceEdit.documentChanges) {
        this._changeAnnotations = new ChangeAnnotations(workspaceEdit.changeAnnotations);
        workspaceEdit.changeAnnotations = this._changeAnnotations.all();
        workspaceEdit.documentChanges.forEach(function(change) {
          if (TextDocumentEdit.is(change)) {
            var textEditChange = new TextEditChangeImpl(change.edits, _this._changeAnnotations);
            _this._textEditChanges[change.textDocument.uri] = textEditChange;
          }
        });
      } else if (workspaceEdit.changes) {
        Object.keys(workspaceEdit.changes).forEach(function(key) {
          var textEditChange = new TextEditChangeImpl(workspaceEdit.changes[key]);
          _this._textEditChanges[key] = textEditChange;
        });
      }
    } else {
      this._workspaceEdit = {};
    }
  }
  Object.defineProperty(WorkspaceChange2.prototype, "edit", {
    get: function() {
      this.initDocumentChanges();
      if (this._changeAnnotations !== void 0) {
        if (this._changeAnnotations.size === 0) {
          this._workspaceEdit.changeAnnotations = void 0;
        } else {
          this._workspaceEdit.changeAnnotations = this._changeAnnotations.all();
        }
      }
      return this._workspaceEdit;
    },
    enumerable: false,
    configurable: true
  });
  WorkspaceChange2.prototype.getTextEditChange = function(key) {
    if (OptionalVersionedTextDocumentIdentifier.is(key)) {
      this.initDocumentChanges();
      if (this._workspaceEdit.documentChanges === void 0) {
        throw new Error("Workspace edit is not configured for document changes.");
      }
      var textDocument = { uri: key.uri, version: key.version };
      var result = this._textEditChanges[textDocument.uri];
      if (!result) {
        var edits = [];
        var textDocumentEdit = {
          textDocument,
          edits
        };
        this._workspaceEdit.documentChanges.push(textDocumentEdit);
        result = new TextEditChangeImpl(edits, this._changeAnnotations);
        this._textEditChanges[textDocument.uri] = result;
      }
      return result;
    } else {
      this.initChanges();
      if (this._workspaceEdit.changes === void 0) {
        throw new Error("Workspace edit is not configured for normal text edit changes.");
      }
      var result = this._textEditChanges[key];
      if (!result) {
        var edits = [];
        this._workspaceEdit.changes[key] = edits;
        result = new TextEditChangeImpl(edits);
        this._textEditChanges[key] = result;
      }
      return result;
    }
  };
  WorkspaceChange2.prototype.initDocumentChanges = function() {
    if (this._workspaceEdit.documentChanges === void 0 && this._workspaceEdit.changes === void 0) {
      this._changeAnnotations = new ChangeAnnotations();
      this._workspaceEdit.documentChanges = [];
      this._workspaceEdit.changeAnnotations = this._changeAnnotations.all();
    }
  };
  WorkspaceChange2.prototype.initChanges = function() {
    if (this._workspaceEdit.documentChanges === void 0 && this._workspaceEdit.changes === void 0) {
      this._workspaceEdit.changes = /* @__PURE__ */ Object.create(null);
    }
  };
  WorkspaceChange2.prototype.createFile = function(uri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = CreateFile.create(uri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = CreateFile.create(uri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  WorkspaceChange2.prototype.renameFile = function(oldUri, newUri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = RenameFile.create(oldUri, newUri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = RenameFile.create(oldUri, newUri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  WorkspaceChange2.prototype.deleteFile = function(uri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = DeleteFile.create(uri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = DeleteFile.create(uri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  return WorkspaceChange2;
}();
var TextDocumentIdentifier;
(function(TextDocumentIdentifier2) {
  function create(uri) {
    return { uri };
  }
  TextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri);
  }
  TextDocumentIdentifier2.is = is;
})(TextDocumentIdentifier || (TextDocumentIdentifier = {}));
var VersionedTextDocumentIdentifier;
(function(VersionedTextDocumentIdentifier2) {
  function create(uri, version) {
    return { uri, version };
  }
  VersionedTextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && Is.integer(candidate.version);
  }
  VersionedTextDocumentIdentifier2.is = is;
})(VersionedTextDocumentIdentifier || (VersionedTextDocumentIdentifier = {}));
var OptionalVersionedTextDocumentIdentifier;
(function(OptionalVersionedTextDocumentIdentifier2) {
  function create(uri, version) {
    return { uri, version };
  }
  OptionalVersionedTextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && (candidate.version === null || Is.integer(candidate.version));
  }
  OptionalVersionedTextDocumentIdentifier2.is = is;
})(OptionalVersionedTextDocumentIdentifier || (OptionalVersionedTextDocumentIdentifier = {}));
var TextDocumentItem;
(function(TextDocumentItem2) {
  function create(uri, languageId, version, text) {
    return { uri, languageId, version, text };
  }
  TextDocumentItem2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && Is.string(candidate.languageId) && Is.integer(candidate.version) && Is.string(candidate.text);
  }
  TextDocumentItem2.is = is;
})(TextDocumentItem || (TextDocumentItem = {}));
var MarkupKind;
(function(MarkupKind2) {
  MarkupKind2.PlainText = "plaintext";
  MarkupKind2.Markdown = "markdown";
})(MarkupKind || (MarkupKind = {}));
(function(MarkupKind2) {
  function is(value) {
    var candidate = value;
    return candidate === MarkupKind2.PlainText || candidate === MarkupKind2.Markdown;
  }
  MarkupKind2.is = is;
})(MarkupKind || (MarkupKind = {}));
var MarkupContent;
(function(MarkupContent2) {
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(value) && MarkupKind.is(candidate.kind) && Is.string(candidate.value);
  }
  MarkupContent2.is = is;
})(MarkupContent || (MarkupContent = {}));
var CompletionItemKind;
(function(CompletionItemKind2) {
  CompletionItemKind2.Text = 1;
  CompletionItemKind2.Method = 2;
  CompletionItemKind2.Function = 3;
  CompletionItemKind2.Constructor = 4;
  CompletionItemKind2.Field = 5;
  CompletionItemKind2.Variable = 6;
  CompletionItemKind2.Class = 7;
  CompletionItemKind2.Interface = 8;
  CompletionItemKind2.Module = 9;
  CompletionItemKind2.Property = 10;
  CompletionItemKind2.Unit = 11;
  CompletionItemKind2.Value = 12;
  CompletionItemKind2.Enum = 13;
  CompletionItemKind2.Keyword = 14;
  CompletionItemKind2.Snippet = 15;
  CompletionItemKind2.Color = 16;
  CompletionItemKind2.File = 17;
  CompletionItemKind2.Reference = 18;
  CompletionItemKind2.Folder = 19;
  CompletionItemKind2.EnumMember = 20;
  CompletionItemKind2.Constant = 21;
  CompletionItemKind2.Struct = 22;
  CompletionItemKind2.Event = 23;
  CompletionItemKind2.Operator = 24;
  CompletionItemKind2.TypeParameter = 25;
})(CompletionItemKind || (CompletionItemKind = {}));
var InsertTextFormat;
(function(InsertTextFormat2) {
  InsertTextFormat2.PlainText = 1;
  InsertTextFormat2.Snippet = 2;
})(InsertTextFormat || (InsertTextFormat = {}));
var CompletionItemTag;
(function(CompletionItemTag2) {
  CompletionItemTag2.Deprecated = 1;
})(CompletionItemTag || (CompletionItemTag = {}));
var InsertReplaceEdit;
(function(InsertReplaceEdit2) {
  function create(newText, insert, replace) {
    return { newText, insert, replace };
  }
  InsertReplaceEdit2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.newText) && Range.is(candidate.insert) && Range.is(candidate.replace);
  }
  InsertReplaceEdit2.is = is;
})(InsertReplaceEdit || (InsertReplaceEdit = {}));
var InsertTextMode;
(function(InsertTextMode2) {
  InsertTextMode2.asIs = 1;
  InsertTextMode2.adjustIndentation = 2;
})(InsertTextMode || (InsertTextMode = {}));
var CompletionItem;
(function(CompletionItem2) {
  function create(label) {
    return { label };
  }
  CompletionItem2.create = create;
})(CompletionItem || (CompletionItem = {}));
var CompletionList;
(function(CompletionList2) {
  function create(items, isIncomplete) {
    return { items: items ? items : [], isIncomplete: !!isIncomplete };
  }
  CompletionList2.create = create;
})(CompletionList || (CompletionList = {}));
var MarkedString;
(function(MarkedString2) {
  function fromPlainText(plainText) {
    return plainText.replace(/[\\`*_{}[\]()#+\-.!]/g, "\\$&");
  }
  MarkedString2.fromPlainText = fromPlainText;
  function is(value) {
    var candidate = value;
    return Is.string(candidate) || Is.objectLiteral(candidate) && Is.string(candidate.language) && Is.string(candidate.value);
  }
  MarkedString2.is = is;
})(MarkedString || (MarkedString = {}));
var Hover;
(function(Hover2) {
  function is(value) {
    var candidate = value;
    return !!candidate && Is.objectLiteral(candidate) && (MarkupContent.is(candidate.contents) || MarkedString.is(candidate.contents) || Is.typedArray(candidate.contents, MarkedString.is)) && (value.range === void 0 || Range.is(value.range));
  }
  Hover2.is = is;
})(Hover || (Hover = {}));
var ParameterInformation;
(function(ParameterInformation2) {
  function create(label, documentation) {
    return documentation ? { label, documentation } : { label };
  }
  ParameterInformation2.create = create;
})(ParameterInformation || (ParameterInformation = {}));
var SignatureInformation;
(function(SignatureInformation2) {
  function create(label, documentation) {
    var parameters = [];
    for (var _i = 2; _i < arguments.length; _i++) {
      parameters[_i - 2] = arguments[_i];
    }
    var result = { label };
    if (Is.defined(documentation)) {
      result.documentation = documentation;
    }
    if (Is.defined(parameters)) {
      result.parameters = parameters;
    } else {
      result.parameters = [];
    }
    return result;
  }
  SignatureInformation2.create = create;
})(SignatureInformation || (SignatureInformation = {}));
var DocumentHighlightKind;
(function(DocumentHighlightKind2) {
  DocumentHighlightKind2.Text = 1;
  DocumentHighlightKind2.Read = 2;
  DocumentHighlightKind2.Write = 3;
})(DocumentHighlightKind || (DocumentHighlightKind = {}));
var DocumentHighlight;
(function(DocumentHighlight2) {
  function create(range, kind) {
    var result = { range };
    if (Is.number(kind)) {
      result.kind = kind;
    }
    return result;
  }
  DocumentHighlight2.create = create;
})(DocumentHighlight || (DocumentHighlight = {}));
var SymbolKind;
(function(SymbolKind2) {
  SymbolKind2.File = 1;
  SymbolKind2.Module = 2;
  SymbolKind2.Namespace = 3;
  SymbolKind2.Package = 4;
  SymbolKind2.Class = 5;
  SymbolKind2.Method = 6;
  SymbolKind2.Property = 7;
  SymbolKind2.Field = 8;
  SymbolKind2.Constructor = 9;
  SymbolKind2.Enum = 10;
  SymbolKind2.Interface = 11;
  SymbolKind2.Function = 12;
  SymbolKind2.Variable = 13;
  SymbolKind2.Constant = 14;
  SymbolKind2.String = 15;
  SymbolKind2.Number = 16;
  SymbolKind2.Boolean = 17;
  SymbolKind2.Array = 18;
  SymbolKind2.Object = 19;
  SymbolKind2.Key = 20;
  SymbolKind2.Null = 21;
  SymbolKind2.EnumMember = 22;
  SymbolKind2.Struct = 23;
  SymbolKind2.Event = 24;
  SymbolKind2.Operator = 25;
  SymbolKind2.TypeParameter = 26;
})(SymbolKind || (SymbolKind = {}));
var SymbolTag;
(function(SymbolTag2) {
  SymbolTag2.Deprecated = 1;
})(SymbolTag || (SymbolTag = {}));
var SymbolInformation;
(function(SymbolInformation2) {
  function create(name, kind, range, uri, containerName) {
    var result = {
      name,
      kind,
      location: { uri, range }
    };
    if (containerName) {
      result.containerName = containerName;
    }
    return result;
  }
  SymbolInformation2.create = create;
})(SymbolInformation || (SymbolInformation = {}));
var DocumentSymbol;
(function(DocumentSymbol2) {
  function create(name, detail, kind, range, selectionRange, children) {
    var result = {
      name,
      detail,
      kind,
      range,
      selectionRange
    };
    if (children !== void 0) {
      result.children = children;
    }
    return result;
  }
  DocumentSymbol2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.name) && Is.number(candidate.kind) && Range.is(candidate.range) && Range.is(candidate.selectionRange) && (candidate.detail === void 0 || Is.string(candidate.detail)) && (candidate.deprecated === void 0 || Is.boolean(candidate.deprecated)) && (candidate.children === void 0 || Array.isArray(candidate.children)) && (candidate.tags === void 0 || Array.isArray(candidate.tags));
  }
  DocumentSymbol2.is = is;
})(DocumentSymbol || (DocumentSymbol = {}));
var CodeActionKind;
(function(CodeActionKind2) {
  CodeActionKind2.Empty = "";
  CodeActionKind2.QuickFix = "quickfix";
  CodeActionKind2.Refactor = "refactor";
  CodeActionKind2.RefactorExtract = "refactor.extract";
  CodeActionKind2.RefactorInline = "refactor.inline";
  CodeActionKind2.RefactorRewrite = "refactor.rewrite";
  CodeActionKind2.Source = "source";
  CodeActionKind2.SourceOrganizeImports = "source.organizeImports";
  CodeActionKind2.SourceFixAll = "source.fixAll";
})(CodeActionKind || (CodeActionKind = {}));
var CodeActionContext;
(function(CodeActionContext2) {
  function create(diagnostics, only) {
    var result = { diagnostics };
    if (only !== void 0 && only !== null) {
      result.only = only;
    }
    return result;
  }
  CodeActionContext2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.typedArray(candidate.diagnostics, Diagnostic.is) && (candidate.only === void 0 || Is.typedArray(candidate.only, Is.string));
  }
  CodeActionContext2.is = is;
})(CodeActionContext || (CodeActionContext = {}));
var CodeAction;
(function(CodeAction2) {
  function create(title, kindOrCommandOrEdit, kind) {
    var result = { title };
    var checkKind = true;
    if (typeof kindOrCommandOrEdit === "string") {
      checkKind = false;
      result.kind = kindOrCommandOrEdit;
    } else if (Command.is(kindOrCommandOrEdit)) {
      result.command = kindOrCommandOrEdit;
    } else {
      result.edit = kindOrCommandOrEdit;
    }
    if (checkKind && kind !== void 0) {
      result.kind = kind;
    }
    return result;
  }
  CodeAction2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.title) && (candidate.diagnostics === void 0 || Is.typedArray(candidate.diagnostics, Diagnostic.is)) && (candidate.kind === void 0 || Is.string(candidate.kind)) && (candidate.edit !== void 0 || candidate.command !== void 0) && (candidate.command === void 0 || Command.is(candidate.command)) && (candidate.isPreferred === void 0 || Is.boolean(candidate.isPreferred)) && (candidate.edit === void 0 || WorkspaceEdit.is(candidate.edit));
  }
  CodeAction2.is = is;
})(CodeAction || (CodeAction = {}));
var CodeLens;
(function(CodeLens2) {
  function create(range, data) {
    var result = { range };
    if (Is.defined(data)) {
      result.data = data;
    }
    return result;
  }
  CodeLens2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.undefined(candidate.command) || Command.is(candidate.command));
  }
  CodeLens2.is = is;
})(CodeLens || (CodeLens = {}));
var FormattingOptions;
(function(FormattingOptions2) {
  function create(tabSize, insertSpaces) {
    return { tabSize, insertSpaces };
  }
  FormattingOptions2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.uinteger(candidate.tabSize) && Is.boolean(candidate.insertSpaces);
  }
  FormattingOptions2.is = is;
})(FormattingOptions || (FormattingOptions = {}));
var DocumentLink;
(function(DocumentLink2) {
  function create(range, target, data) {
    return { range, target, data };
  }
  DocumentLink2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.undefined(candidate.target) || Is.string(candidate.target));
  }
  DocumentLink2.is = is;
})(DocumentLink || (DocumentLink = {}));
var SelectionRange;
(function(SelectionRange2) {
  function create(range, parent) {
    return { range, parent };
  }
  SelectionRange2.create = create;
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && Range.is(candidate.range) && (candidate.parent === void 0 || SelectionRange2.is(candidate.parent));
  }
  SelectionRange2.is = is;
})(SelectionRange || (SelectionRange = {}));
var TextDocument;
(function(TextDocument2) {
  function create(uri, languageId, version, content) {
    return new FullTextDocument(uri, languageId, version, content);
  }
  TextDocument2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && (Is.undefined(candidate.languageId) || Is.string(candidate.languageId)) && Is.uinteger(candidate.lineCount) && Is.func(candidate.getText) && Is.func(candidate.positionAt) && Is.func(candidate.offsetAt) ? true : false;
  }
  TextDocument2.is = is;
  function applyEdits(document, edits) {
    var text = document.getText();
    var sortedEdits = mergeSort(edits, function(a, b) {
      var diff = a.range.start.line - b.range.start.line;
      if (diff === 0) {
        return a.range.start.character - b.range.start.character;
      }
      return diff;
    });
    var lastModifiedOffset = text.length;
    for (var i = sortedEdits.length - 1; i >= 0; i--) {
      var e = sortedEdits[i];
      var startOffset = document.offsetAt(e.range.start);
      var endOffset = document.offsetAt(e.range.end);
      if (endOffset <= lastModifiedOffset) {
        text = text.substring(0, startOffset) + e.newText + text.substring(endOffset, text.length);
      } else {
        throw new Error("Overlapping edit");
      }
      lastModifiedOffset = startOffset;
    }
    return text;
  }
  TextDocument2.applyEdits = applyEdits;
  function mergeSort(data, compare) {
    if (data.length <= 1) {
      return data;
    }
    var p = data.length / 2 | 0;
    var left = data.slice(0, p);
    var right = data.slice(p);
    mergeSort(left, compare);
    mergeSort(right, compare);
    var leftIdx = 0;
    var rightIdx = 0;
    var i = 0;
    while (leftIdx < left.length && rightIdx < right.length) {
      var ret = compare(left[leftIdx], right[rightIdx]);
      if (ret <= 0) {
        data[i++] = left[leftIdx++];
      } else {
        data[i++] = right[rightIdx++];
      }
    }
    while (leftIdx < left.length) {
      data[i++] = left[leftIdx++];
    }
    while (rightIdx < right.length) {
      data[i++] = right[rightIdx++];
    }
    return data;
  }
})(TextDocument || (TextDocument = {}));
var FullTextDocument = function() {
  function FullTextDocument2(uri, languageId, version, content) {
    this._uri = uri;
    this._languageId = languageId;
    this._version = version;
    this._content = content;
    this._lineOffsets = void 0;
  }
  Object.defineProperty(FullTextDocument2.prototype, "uri", {
    get: function() {
      return this._uri;
    },
    enumerable: false,
    configurable: true
  });
  Object.defineProperty(FullTextDocument2.prototype, "languageId", {
    get: function() {
      return this._languageId;
    },
    enumerable: false,
    configurable: true
  });
  Object.defineProperty(FullTextDocument2.prototype, "version", {
    get: function() {
      return this._version;
    },
    enumerable: false,
    configurable: true
  });
  FullTextDocument2.prototype.getText = function(range) {
    if (range) {
      var start = this.offsetAt(range.start);
      var end = this.offsetAt(range.end);
      return this._content.substring(start, end);
    }
    return this._content;
  };
  FullTextDocument2.prototype.update = function(event, version) {
    this._content = event.text;
    this._version = version;
    this._lineOffsets = void 0;
  };
  FullTextDocument2.prototype.getLineOffsets = function() {
    if (this._lineOffsets === void 0) {
      var lineOffsets = [];
      var text = this._content;
      var isLineStart = true;
      for (var i = 0; i < text.length; i++) {
        if (isLineStart) {
          lineOffsets.push(i);
          isLineStart = false;
        }
        var ch = text.charAt(i);
        isLineStart = ch === "\r" || ch === "\n";
        if (ch === "\r" && i + 1 < text.length && text.charAt(i + 1) === "\n") {
          i++;
        }
      }
      if (isLineStart && text.length > 0) {
        lineOffsets.push(text.length);
      }
      this._lineOffsets = lineOffsets;
    }
    return this._lineOffsets;
  };
  FullTextDocument2.prototype.positionAt = function(offset) {
    offset = Math.max(Math.min(offset, this._content.length), 0);
    var lineOffsets = this.getLineOffsets();
    var low = 0, high = lineOffsets.length;
    if (high === 0) {
      return Position.create(0, offset);
    }
    while (low < high) {
      var mid = Math.floor((low + high) / 2);
      if (lineOffsets[mid] > offset) {
        high = mid;
      } else {
        low = mid + 1;
      }
    }
    var line = low - 1;
    return Position.create(line, offset - lineOffsets[line]);
  };
  FullTextDocument2.prototype.offsetAt = function(position) {
    var lineOffsets = this.getLineOffsets();
    if (position.line >= lineOffsets.length) {
      return this._content.length;
    } else if (position.line < 0) {
      return 0;
    }
    var lineOffset = lineOffsets[position.line];
    var nextLineOffset = position.line + 1 < lineOffsets.length ? lineOffsets[position.line + 1] : this._content.length;
    return Math.max(Math.min(lineOffset + position.character, nextLineOffset), lineOffset);
  };
  Object.defineProperty(FullTextDocument2.prototype, "lineCount", {
    get: function() {
      return this.getLineOffsets().length;
    },
    enumerable: false,
    configurable: true
  });
  return FullTextDocument2;
}();
var Is;
(function(Is2) {
  var toString = Object.prototype.toString;
  function defined(value) {
    return typeof value !== "undefined";
  }
  Is2.defined = defined;
  function undefined2(value) {
    return typeof value === "undefined";
  }
  Is2.undefined = undefined2;
  function boolean(value) {
    return value === true || value === false;
  }
  Is2.boolean = boolean;
  function string(value) {
    return toString.call(value) === "[object String]";
  }
  Is2.string = string;
  function number(value) {
    return toString.call(value) === "[object Number]";
  }
  Is2.number = number;
  function numberRange(value, min, max) {
    return toString.call(value) === "[object Number]" && min <= value && value <= max;
  }
  Is2.numberRange = numberRange;
  function integer2(value) {
    return toString.call(value) === "[object Number]" && -2147483648 <= value && value <= 2147483647;
  }
  Is2.integer = integer2;
  function uinteger2(value) {
    return toString.call(value) === "[object Number]" && 0 <= value && value <= 2147483647;
  }
  Is2.uinteger = uinteger2;
  function func(value) {
    return toString.call(value) === "[object Function]";
  }
  Is2.func = func;
  function objectLiteral(value) {
    return value !== null && typeof value === "object";
  }
  Is2.objectLiteral = objectLiteral;
  function typedArray(value, check) {
    return Array.isArray(value) && value.every(check);
  }
  Is2.typedArray = typedArray;
})(Is || (Is = {}));

// src/language/common/lspLanguageFeatures.ts
var DiagnosticsAdapter = class {
  constructor(_languageId, _worker, configChangeEvent) {
    this._languageId = _languageId;
    this._worker = _worker;
    const onModelAdd = (model) => {
      let modeId = model.getLanguageId();
      if (modeId !== this._languageId) {
        return;
      }
      let handle;
      this._listener[model.uri.toString()] = model.onDidChangeContent(() => {
        window.clearTimeout(handle);
        handle = window.setTimeout(() => this._doValidate(model.uri, modeId), 500);
      });
      this._doValidate(model.uri, modeId);
    };
    const onModelRemoved = (model) => {
      monaco_editor_core_exports.editor.setModelMarkers(model, this._languageId, []);
      let uriStr = model.uri.toString();
      let listener = this._listener[uriStr];
      if (listener) {
        listener.dispose();
        delete this._listener[uriStr];
      }
    };
    this._disposables.push(monaco_editor_core_exports.editor.onDidCreateModel(onModelAdd));
    this._disposables.push(monaco_editor_core_exports.editor.onWillDisposeModel(onModelRemoved));
    this._disposables.push(monaco_editor_core_exports.editor.onDidChangeModelLanguage((event) => {
      onModelRemoved(event.model);
      onModelAdd(event.model);
    }));
    this._disposables.push(configChangeEvent((_) => {
      monaco_editor_core_exports.editor.getModels().forEach((model) => {
        if (model.getLanguageId() === this._languageId) {
          onModelRemoved(model);
          onModelAdd(model);
        }
      });
    }));
    this._disposables.push({
      dispose: () => {
        monaco_editor_core_exports.editor.getModels().forEach(onModelRemoved);
        for (let key in this._listener) {
          this._listener[key].dispose();
        }
      }
    });
    monaco_editor_core_exports.editor.getModels().forEach(onModelAdd);
  }
  _disposables = [];
  _listener = /* @__PURE__ */ Object.create(null);
  dispose() {
    this._disposables.forEach((d) => d && d.dispose());
    this._disposables.length = 0;
  }
  _doValidate(resource, languageId) {
    this._worker(resource).then((worker) => {
      return worker.doValidation(resource.toString());
    }).then((diagnostics) => {
      const markers = diagnostics.map((d) => toDiagnostics(resource, d));
      let model = monaco_editor_core_exports.editor.getModel(resource);
      if (model && model.getLanguageId() === languageId) {
        monaco_editor_core_exports.editor.setModelMarkers(model, languageId, markers);
      }
    }).then(void 0, (err) => {
      console.error(err);
    });
  }
};
function toSeverity(lsSeverity) {
  switch (lsSeverity) {
    case DiagnosticSeverity.Error:
      return monaco_editor_core_exports.MarkerSeverity.Error;
    case DiagnosticSeverity.Warning:
      return monaco_editor_core_exports.MarkerSeverity.Warning;
    case DiagnosticSeverity.Information:
      return monaco_editor_core_exports.MarkerSeverity.Info;
    case DiagnosticSeverity.Hint:
      return monaco_editor_core_exports.MarkerSeverity.Hint;
    default:
      return monaco_editor_core_exports.MarkerSeverity.Info;
  }
}
function toDiagnostics(resource, diag) {
  let code = typeof diag.code === "number" ? String(diag.code) : diag.code;
  return {
    severity: toSeverity(diag.severity),
    startLineNumber: diag.range.start.line + 1,
    startColumn: diag.range.start.character + 1,
    endLineNumber: diag.range.end.line + 1,
    endColumn: diag.range.end.character + 1,
    message: diag.message,
    code,
    source: diag.source
  };
}
var CompletionAdapter = class {
  constructor(_worker, _triggerCharacters) {
    this._worker = _worker;
    this._triggerCharacters = _triggerCharacters;
  }
  get triggerCharacters() {
    return this._triggerCharacters;
  }
  provideCompletionItems(model, position, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doComplete(resource.toString(), fromPosition(position));
    }).then((info) => {
      if (!info) {
        return;
      }
      const wordInfo = model.getWordUntilPosition(position);
      const wordRange = new monaco_editor_core_exports.Range(position.lineNumber, wordInfo.startColumn, position.lineNumber, wordInfo.endColumn);
      const items = info.items.map((entry) => {
        const item = {
          label: entry.label,
          insertText: entry.insertText || entry.label,
          sortText: entry.sortText,
          filterText: entry.filterText,
          documentation: entry.documentation,
          detail: entry.detail,
          command: toCommand(entry.command),
          range: wordRange,
          kind: toCompletionItemKind(entry.kind)
        };
        if (entry.textEdit) {
          if (isInsertReplaceEdit(entry.textEdit)) {
            item.range = {
              insert: toRange(entry.textEdit.insert),
              replace: toRange(entry.textEdit.replace)
            };
          } else {
            item.range = toRange(entry.textEdit.range);
          }
          item.insertText = entry.textEdit.newText;
        }
        if (entry.additionalTextEdits) {
          item.additionalTextEdits = entry.additionalTextEdits.map(toTextEdit);
        }
        if (entry.insertTextFormat === InsertTextFormat.Snippet) {
          item.insertTextRules = monaco_editor_core_exports.languages.CompletionItemInsertTextRule.InsertAsSnippet;
        }
        return item;
      });
      return {
        isIncomplete: info.isIncomplete,
        suggestions: items
      };
    });
  }
};
function fromPosition(position) {
  if (!position) {
    return void 0;
  }
  return { character: position.column - 1, line: position.lineNumber - 1 };
}
function fromRange(range) {
  if (!range) {
    return void 0;
  }
  return {
    start: {
      line: range.startLineNumber - 1,
      character: range.startColumn - 1
    },
    end: { line: range.endLineNumber - 1, character: range.endColumn - 1 }
  };
}
function toRange(range) {
  if (!range) {
    return void 0;
  }
  return new monaco_editor_core_exports.Range(range.start.line + 1, range.start.character + 1, range.end.line + 1, range.end.character + 1);
}
function isInsertReplaceEdit(edit) {
  return typeof edit.insert !== "undefined" && typeof edit.replace !== "undefined";
}
function toCompletionItemKind(kind) {
  const mItemKind = monaco_editor_core_exports.languages.CompletionItemKind;
  switch (kind) {
    case CompletionItemKind.Text:
      return mItemKind.Text;
    case CompletionItemKind.Method:
      return mItemKind.Method;
    case CompletionItemKind.Function:
      return mItemKind.Function;
    case CompletionItemKind.Constructor:
      return mItemKind.Constructor;
    case CompletionItemKind.Field:
      return mItemKind.Field;
    case CompletionItemKind.Variable:
      return mItemKind.Variable;
    case CompletionItemKind.Class:
      return mItemKind.Class;
    case CompletionItemKind.Interface:
      return mItemKind.Interface;
    case CompletionItemKind.Module:
      return mItemKind.Module;
    case CompletionItemKind.Property:
      return mItemKind.Property;
    case CompletionItemKind.Unit:
      return mItemKind.Unit;
    case CompletionItemKind.Value:
      return mItemKind.Value;
    case CompletionItemKind.Enum:
      return mItemKind.Enum;
    case CompletionItemKind.Keyword:
      return mItemKind.Keyword;
    case CompletionItemKind.Snippet:
      return mItemKind.Snippet;
    case CompletionItemKind.Color:
      return mItemKind.Color;
    case CompletionItemKind.File:
      return mItemKind.File;
    case CompletionItemKind.Reference:
      return mItemKind.Reference;
  }
  return mItemKind.Property;
}
function toTextEdit(textEdit) {
  if (!textEdit) {
    return void 0;
  }
  return {
    range: toRange(textEdit.range),
    text: textEdit.newText
  };
}
function toCommand(c) {
  return c && c.command === "editor.action.triggerSuggest" ? { id: c.command, title: c.title, arguments: c.arguments } : void 0;
}
var HoverAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideHover(model, position, token) {
    let resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doHover(resource.toString(), fromPosition(position));
    }).then((info) => {
      if (!info) {
        return;
      }
      return {
        range: toRange(info.range),
        contents: toMarkedStringArray(info.contents)
      };
    });
  }
};
function isMarkupContent(thing) {
  return thing && typeof thing === "object" && typeof thing.kind === "string";
}
function toMarkdownString(entry) {
  if (typeof entry === "string") {
    return {
      value: entry
    };
  }
  if (isMarkupContent(entry)) {
    if (entry.kind === "plaintext") {
      return {
        value: entry.value.replace(/[\\`*_{}[\]()#+\-.!]/g, "\\$&")
      };
    }
    return {
      value: entry.value
    };
  }
  return { value: "```" + entry.language + "\n" + entry.value + "\n```\n" };
}
function toMarkedStringArray(contents) {
  if (!contents) {
    return void 0;
  }
  if (Array.isArray(contents)) {
    return contents.map(toMarkdownString);
  }
  return [toMarkdownString(contents)];
}
var DocumentHighlightAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentHighlights(model, position, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentHighlights(resource.toString(), fromPosition(position))).then((entries) => {
      if (!entries) {
        return;
      }
      return entries.map((entry) => {
        return {
          range: toRange(entry.range),
          kind: toDocumentHighlightKind(entry.kind)
        };
      });
    });
  }
};
function toDocumentHighlightKind(kind) {
  switch (kind) {
    case DocumentHighlightKind.Read:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Read;
    case DocumentHighlightKind.Write:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Write;
    case DocumentHighlightKind.Text:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Text;
  }
  return monaco_editor_core_exports.languages.DocumentHighlightKind.Text;
}
var DefinitionAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDefinition(model, position, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.findDefinition(resource.toString(), fromPosition(position));
    }).then((definition) => {
      if (!definition) {
        return;
      }
      return [toLocation(definition)];
    });
  }
};
function toLocation(location) {
  return {
    uri: monaco_editor_core_exports.Uri.parse(location.uri),
    range: toRange(location.range)
  };
}
var ReferenceAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideReferences(model, position, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.findReferences(resource.toString(), fromPosition(position));
    }).then((entries) => {
      if (!entries) {
        return;
      }
      return entries.map(toLocation);
    });
  }
};
var RenameAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideRenameEdits(model, position, newName, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doRename(resource.toString(), fromPosition(position), newName);
    }).then((edit) => {
      return toWorkspaceEdit(edit);
    });
  }
};
function toWorkspaceEdit(edit) {
  if (!edit || !edit.changes) {
    return void 0;
  }
  let resourceEdits = [];
  for (let uri in edit.changes) {
    const _uri = monaco_editor_core_exports.Uri.parse(uri);
    for (let e of edit.changes[uri]) {
      resourceEdits.push({
        resource: _uri,
        versionId: void 0,
        textEdit: {
          range: toRange(e.range),
          text: e.newText
        }
      });
    }
  }
  return {
    edits: resourceEdits
  };
}
var DocumentSymbolAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentSymbols(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentSymbols(resource.toString())).then((items) => {
      if (!items) {
        return;
      }
      return items.map((item) => ({
        name: item.name,
        detail: "",
        containerName: item.containerName,
        kind: toSymbolKind(item.kind),
        range: toRange(item.location.range),
        selectionRange: toRange(item.location.range),
        tags: []
      }));
    });
  }
};
function toSymbolKind(kind) {
  let mKind = monaco_editor_core_exports.languages.SymbolKind;
  switch (kind) {
    case SymbolKind.File:
      return mKind.Array;
    case SymbolKind.Module:
      return mKind.Module;
    case SymbolKind.Namespace:
      return mKind.Namespace;
    case SymbolKind.Package:
      return mKind.Package;
    case SymbolKind.Class:
      return mKind.Class;
    case SymbolKind.Method:
      return mKind.Method;
    case SymbolKind.Property:
      return mKind.Property;
    case SymbolKind.Field:
      return mKind.Field;
    case SymbolKind.Constructor:
      return mKind.Constructor;
    case SymbolKind.Enum:
      return mKind.Enum;
    case SymbolKind.Interface:
      return mKind.Interface;
    case SymbolKind.Function:
      return mKind.Function;
    case SymbolKind.Variable:
      return mKind.Variable;
    case SymbolKind.Constant:
      return mKind.Constant;
    case SymbolKind.String:
      return mKind.String;
    case SymbolKind.Number:
      return mKind.Number;
    case SymbolKind.Boolean:
      return mKind.Boolean;
    case SymbolKind.Array:
      return mKind.Array;
  }
  return mKind.Function;
}
var DocumentLinkAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideLinks(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentLinks(resource.toString())).then((items) => {
      if (!items) {
        return;
      }
      return {
        links: items.map((item) => ({
          range: toRange(item.range),
          url: item.target
        }))
      };
    });
  }
};
var DocumentFormattingEditProvider = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentFormattingEdits(model, options, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.format(resource.toString(), null, fromFormattingOptions(options)).then((edits) => {
        if (!edits || edits.length === 0) {
          return;
        }
        return edits.map(toTextEdit);
      });
    });
  }
};
var DocumentRangeFormattingEditProvider = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  canFormatMultipleRanges = false;
  provideDocumentRangeFormattingEdits(model, range, options, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.format(resource.toString(), fromRange(range), fromFormattingOptions(options)).then((edits) => {
        if (!edits || edits.length === 0) {
          return;
        }
        return edits.map(toTextEdit);
      });
    });
  }
};
function fromFormattingOptions(options) {
  return {
    tabSize: options.tabSize,
    insertSpaces: options.insertSpaces
  };
}
var DocumentColorAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentColors(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentColors(resource.toString())).then((infos) => {
      if (!infos) {
        return;
      }
      return infos.map((item) => ({
        color: item.color,
        range: toRange(item.range)
      }));
    });
  }
  provideColorPresentations(model, info, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getColorPresentations(resource.toString(), info.color, fromRange(info.range))).then((presentations) => {
      if (!presentations) {
        return;
      }
      return presentations.map((presentation) => {
        let item = {
          label: presentation.label
        };
        if (presentation.textEdit) {
          item.textEdit = toTextEdit(presentation.textEdit);
        }
        if (presentation.additionalTextEdits) {
          item.additionalTextEdits = presentation.additionalTextEdits.map(toTextEdit);
        }
        return item;
      });
    });
  }
};
var FoldingRangeAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideFoldingRanges(model, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getFoldingRanges(resource.toString(), context)).then((ranges) => {
      if (!ranges) {
        return;
      }
      return ranges.map((range) => {
        const result = {
          start: range.startLine + 1,
          end: range.endLine + 1
        };
        if (typeof range.kind !== "undefined") {
          result.kind = toFoldingRangeKind(range.kind);
        }
        return result;
      });
    });
  }
};
function toFoldingRangeKind(kind) {
  switch (kind) {
    case FoldingRangeKind.Comment:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Comment;
    case FoldingRangeKind.Imports:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Imports;
    case FoldingRangeKind.Region:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Region;
  }
  return void 0;
}
var SelectionRangeAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideSelectionRanges(model, positions, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getSelectionRanges(resource.toString(), positions.map(fromPosition))).then((selectionRanges) => {
      if (!selectionRanges) {
        return;
      }
      return selectionRanges.map((selectionRange) => {
        const result = [];
        while (selectionRange) {
          result.push({ range: toRange(selectionRange.range) });
          selectionRange = selectionRange.parent;
        }
        return result;
      });
    });
  }
};

// src/language/css/cssMode.ts
function setupMode(defaults) {
  const disposables = [];
  const providers = [];
  const client = new WorkerManager(defaults);
  disposables.push(client);
  const worker = (...uris) => {
    return client.getLanguageServiceWorker(...uris);
  };
  function registerProviders() {
    const { languageId, modeConfiguration } = defaults;
    disposeAll(providers);
    if (modeConfiguration.completionItems) {
      providers.push(monaco_editor_core_exports.languages.registerCompletionItemProvider(languageId, new CompletionAdapter(worker, ["/", "-", ":"])));
    }
    if (modeConfiguration.hovers) {
      providers.push(monaco_editor_core_exports.languages.registerHoverProvider(languageId, new HoverAdapter(worker)));
    }
    if (modeConfiguration.documentHighlights) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentHighlightProvider(languageId, new DocumentHighlightAdapter(worker)));
    }
    if (modeConfiguration.definitions) {
      providers.push(monaco_editor_core_exports.languages.registerDefinitionProvider(languageId, new DefinitionAdapter(worker)));
    }
    if (modeConfiguration.references) {
      providers.push(monaco_editor_core_exports.languages.registerReferenceProvider(languageId, new ReferenceAdapter(worker)));
    }
    if (modeConfiguration.documentSymbols) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentSymbolProvider(languageId, new DocumentSymbolAdapter(worker)));
    }
    if (modeConfiguration.rename) {
      providers.push(monaco_editor_core_exports.languages.registerRenameProvider(languageId, new RenameAdapter(worker)));
    }
    if (modeConfiguration.colors) {
      providers.push(monaco_editor_core_exports.languages.registerColorProvider(languageId, new DocumentColorAdapter(worker)));
    }
    if (modeConfiguration.foldingRanges) {
      providers.push(monaco_editor_core_exports.languages.registerFoldingRangeProvider(languageId, new FoldingRangeAdapter(worker)));
    }
    if (modeConfiguration.diagnostics) {
      providers.push(new DiagnosticsAdapter(languageId, worker, defaults.onDidChange));
    }
    if (modeConfiguration.selectionRanges) {
      providers.push(monaco_editor_core_exports.languages.registerSelectionRangeProvider(languageId, new SelectionRangeAdapter(worker)));
    }
    if (modeConfiguration.documentFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentFormattingEditProvider(languageId, new DocumentFormattingEditProvider(worker)));
    }
    if (modeConfiguration.documentRangeFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentRangeFormattingEditProvider(languageId, new DocumentRangeFormattingEditProvider(worker)));
    }
  }
  registerProviders();
  disposables.push(asDisposable(providers));
  return asDisposable(disposables);
}
function asDisposable(disposables) {
  return { dispose: () => disposeAll(disposables) };
}
function disposeAll(disposables) {
  while (disposables.length) {
    disposables.pop().dispose();
  }
}



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/language/html/htmlMode.js":
/*!*********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/language/html/htmlMode.js ***!
  \*********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   CompletionAdapter: function() { return /* binding */ CompletionAdapter; },
/* harmony export */   DefinitionAdapter: function() { return /* binding */ DefinitionAdapter; },
/* harmony export */   DiagnosticsAdapter: function() { return /* binding */ DiagnosticsAdapter; },
/* harmony export */   DocumentColorAdapter: function() { return /* binding */ DocumentColorAdapter; },
/* harmony export */   DocumentFormattingEditProvider: function() { return /* binding */ DocumentFormattingEditProvider; },
/* harmony export */   DocumentHighlightAdapter: function() { return /* binding */ DocumentHighlightAdapter; },
/* harmony export */   DocumentLinkAdapter: function() { return /* binding */ DocumentLinkAdapter; },
/* harmony export */   DocumentRangeFormattingEditProvider: function() { return /* binding */ DocumentRangeFormattingEditProvider; },
/* harmony export */   DocumentSymbolAdapter: function() { return /* binding */ DocumentSymbolAdapter; },
/* harmony export */   FoldingRangeAdapter: function() { return /* binding */ FoldingRangeAdapter; },
/* harmony export */   HoverAdapter: function() { return /* binding */ HoverAdapter; },
/* harmony export */   ReferenceAdapter: function() { return /* binding */ ReferenceAdapter; },
/* harmony export */   RenameAdapter: function() { return /* binding */ RenameAdapter; },
/* harmony export */   SelectionRangeAdapter: function() { return /* binding */ SelectionRangeAdapter; },
/* harmony export */   WorkerManager: function() { return /* binding */ WorkerManager; },
/* harmony export */   fromPosition: function() { return /* binding */ fromPosition; },
/* harmony export */   fromRange: function() { return /* binding */ fromRange; },
/* harmony export */   setupMode: function() { return /* binding */ setupMode; },
/* harmony export */   setupMode1: function() { return /* binding */ setupMode1; },
/* harmony export */   toRange: function() { return /* binding */ toRange; },
/* harmony export */   toTextEdit: function() { return /* binding */ toTextEdit; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/language/html/workerManager.ts
var STOP_WHEN_IDLE_FOR = 2 * 60 * 1e3;
var WorkerManager = class {
  _defaults;
  _idleCheckInterval;
  _lastUsedTime;
  _configChangeListener;
  _worker;
  _client;
  constructor(defaults) {
    this._defaults = defaults;
    this._worker = null;
    this._client = null;
    this._idleCheckInterval = window.setInterval(() => this._checkIfIdle(), 30 * 1e3);
    this._lastUsedTime = 0;
    this._configChangeListener = this._defaults.onDidChange(() => this._stopWorker());
  }
  _stopWorker() {
    if (this._worker) {
      this._worker.dispose();
      this._worker = null;
    }
    this._client = null;
  }
  dispose() {
    clearInterval(this._idleCheckInterval);
    this._configChangeListener.dispose();
    this._stopWorker();
  }
  _checkIfIdle() {
    if (!this._worker) {
      return;
    }
    let timePassedSinceLastUsed = Date.now() - this._lastUsedTime;
    if (timePassedSinceLastUsed > STOP_WHEN_IDLE_FOR) {
      this._stopWorker();
    }
  }
  _getClient() {
    this._lastUsedTime = Date.now();
    if (!this._client) {
      this._worker = monaco_editor_core_exports.editor.createWebWorker({
        moduleId: "vs/language/html/htmlWorker",
        createData: {
          languageSettings: this._defaults.options,
          languageId: this._defaults.languageId
        },
        label: this._defaults.languageId
      });
      this._client = this._worker.getProxy();
    }
    return this._client;
  }
  getLanguageServiceWorker(...resources) {
    let _client;
    return this._getClient().then((client) => {
      _client = client;
    }).then((_) => {
      if (this._worker) {
        return this._worker.withSyncedResources(resources);
      }
    }).then((_) => _client);
  }
};

// node_modules/vscode-languageserver-types/lib/esm/main.js
var integer;
(function(integer2) {
  integer2.MIN_VALUE = -2147483648;
  integer2.MAX_VALUE = 2147483647;
})(integer || (integer = {}));
var uinteger;
(function(uinteger2) {
  uinteger2.MIN_VALUE = 0;
  uinteger2.MAX_VALUE = 2147483647;
})(uinteger || (uinteger = {}));
var Position;
(function(Position3) {
  function create(line, character) {
    if (line === Number.MAX_VALUE) {
      line = uinteger.MAX_VALUE;
    }
    if (character === Number.MAX_VALUE) {
      character = uinteger.MAX_VALUE;
    }
    return { line, character };
  }
  Position3.create = create;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Is.uinteger(candidate.line) && Is.uinteger(candidate.character);
  }
  Position3.is = is;
})(Position || (Position = {}));
var Range;
(function(Range3) {
  function create(one, two, three, four) {
    if (Is.uinteger(one) && Is.uinteger(two) && Is.uinteger(three) && Is.uinteger(four)) {
      return { start: Position.create(one, two), end: Position.create(three, four) };
    } else if (Position.is(one) && Position.is(two)) {
      return { start: one, end: two };
    } else {
      throw new Error("Range#create called with invalid arguments[" + one + ", " + two + ", " + three + ", " + four + "]");
    }
  }
  Range3.create = create;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Position.is(candidate.start) && Position.is(candidate.end);
  }
  Range3.is = is;
})(Range || (Range = {}));
var Location;
(function(Location2) {
  function create(uri, range) {
    return { uri, range };
  }
  Location2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.string(candidate.uri) || Is.undefined(candidate.uri));
  }
  Location2.is = is;
})(Location || (Location = {}));
var LocationLink;
(function(LocationLink2) {
  function create(targetUri, targetRange, targetSelectionRange, originSelectionRange) {
    return { targetUri, targetRange, targetSelectionRange, originSelectionRange };
  }
  LocationLink2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.targetRange) && Is.string(candidate.targetUri) && (Range.is(candidate.targetSelectionRange) || Is.undefined(candidate.targetSelectionRange)) && (Range.is(candidate.originSelectionRange) || Is.undefined(candidate.originSelectionRange));
  }
  LocationLink2.is = is;
})(LocationLink || (LocationLink = {}));
var Color;
(function(Color2) {
  function create(red, green, blue, alpha) {
    return {
      red,
      green,
      blue,
      alpha
    };
  }
  Color2.create = create;
  function is(value) {
    var candidate = value;
    return Is.numberRange(candidate.red, 0, 1) && Is.numberRange(candidate.green, 0, 1) && Is.numberRange(candidate.blue, 0, 1) && Is.numberRange(candidate.alpha, 0, 1);
  }
  Color2.is = is;
})(Color || (Color = {}));
var ColorInformation;
(function(ColorInformation2) {
  function create(range, color) {
    return {
      range,
      color
    };
  }
  ColorInformation2.create = create;
  function is(value) {
    var candidate = value;
    return Range.is(candidate.range) && Color.is(candidate.color);
  }
  ColorInformation2.is = is;
})(ColorInformation || (ColorInformation = {}));
var ColorPresentation;
(function(ColorPresentation2) {
  function create(label, textEdit, additionalTextEdits) {
    return {
      label,
      textEdit,
      additionalTextEdits
    };
  }
  ColorPresentation2.create = create;
  function is(value) {
    var candidate = value;
    return Is.string(candidate.label) && (Is.undefined(candidate.textEdit) || TextEdit.is(candidate)) && (Is.undefined(candidate.additionalTextEdits) || Is.typedArray(candidate.additionalTextEdits, TextEdit.is));
  }
  ColorPresentation2.is = is;
})(ColorPresentation || (ColorPresentation = {}));
var FoldingRangeKind;
(function(FoldingRangeKind2) {
  FoldingRangeKind2["Comment"] = "comment";
  FoldingRangeKind2["Imports"] = "imports";
  FoldingRangeKind2["Region"] = "region";
})(FoldingRangeKind || (FoldingRangeKind = {}));
var FoldingRange;
(function(FoldingRange2) {
  function create(startLine, endLine, startCharacter, endCharacter, kind) {
    var result = {
      startLine,
      endLine
    };
    if (Is.defined(startCharacter)) {
      result.startCharacter = startCharacter;
    }
    if (Is.defined(endCharacter)) {
      result.endCharacter = endCharacter;
    }
    if (Is.defined(kind)) {
      result.kind = kind;
    }
    return result;
  }
  FoldingRange2.create = create;
  function is(value) {
    var candidate = value;
    return Is.uinteger(candidate.startLine) && Is.uinteger(candidate.startLine) && (Is.undefined(candidate.startCharacter) || Is.uinteger(candidate.startCharacter)) && (Is.undefined(candidate.endCharacter) || Is.uinteger(candidate.endCharacter)) && (Is.undefined(candidate.kind) || Is.string(candidate.kind));
  }
  FoldingRange2.is = is;
})(FoldingRange || (FoldingRange = {}));
var DiagnosticRelatedInformation;
(function(DiagnosticRelatedInformation2) {
  function create(location, message) {
    return {
      location,
      message
    };
  }
  DiagnosticRelatedInformation2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Location.is(candidate.location) && Is.string(candidate.message);
  }
  DiagnosticRelatedInformation2.is = is;
})(DiagnosticRelatedInformation || (DiagnosticRelatedInformation = {}));
var DiagnosticSeverity;
(function(DiagnosticSeverity2) {
  DiagnosticSeverity2.Error = 1;
  DiagnosticSeverity2.Warning = 2;
  DiagnosticSeverity2.Information = 3;
  DiagnosticSeverity2.Hint = 4;
})(DiagnosticSeverity || (DiagnosticSeverity = {}));
var DiagnosticTag;
(function(DiagnosticTag2) {
  DiagnosticTag2.Unnecessary = 1;
  DiagnosticTag2.Deprecated = 2;
})(DiagnosticTag || (DiagnosticTag = {}));
var CodeDescription;
(function(CodeDescription2) {
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && candidate !== null && Is.string(candidate.href);
  }
  CodeDescription2.is = is;
})(CodeDescription || (CodeDescription = {}));
var Diagnostic;
(function(Diagnostic2) {
  function create(range, message, severity, code, source, relatedInformation) {
    var result = { range, message };
    if (Is.defined(severity)) {
      result.severity = severity;
    }
    if (Is.defined(code)) {
      result.code = code;
    }
    if (Is.defined(source)) {
      result.source = source;
    }
    if (Is.defined(relatedInformation)) {
      result.relatedInformation = relatedInformation;
    }
    return result;
  }
  Diagnostic2.create = create;
  function is(value) {
    var _a;
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && Is.string(candidate.message) && (Is.number(candidate.severity) || Is.undefined(candidate.severity)) && (Is.integer(candidate.code) || Is.string(candidate.code) || Is.undefined(candidate.code)) && (Is.undefined(candidate.codeDescription) || Is.string((_a = candidate.codeDescription) === null || _a === void 0 ? void 0 : _a.href)) && (Is.string(candidate.source) || Is.undefined(candidate.source)) && (Is.undefined(candidate.relatedInformation) || Is.typedArray(candidate.relatedInformation, DiagnosticRelatedInformation.is));
  }
  Diagnostic2.is = is;
})(Diagnostic || (Diagnostic = {}));
var Command;
(function(Command2) {
  function create(title, command) {
    var args = [];
    for (var _i = 2; _i < arguments.length; _i++) {
      args[_i - 2] = arguments[_i];
    }
    var result = { title, command };
    if (Is.defined(args) && args.length > 0) {
      result.arguments = args;
    }
    return result;
  }
  Command2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.title) && Is.string(candidate.command);
  }
  Command2.is = is;
})(Command || (Command = {}));
var TextEdit;
(function(TextEdit2) {
  function replace(range, newText) {
    return { range, newText };
  }
  TextEdit2.replace = replace;
  function insert(position, newText) {
    return { range: { start: position, end: position }, newText };
  }
  TextEdit2.insert = insert;
  function del(range) {
    return { range, newText: "" };
  }
  TextEdit2.del = del;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Is.string(candidate.newText) && Range.is(candidate.range);
  }
  TextEdit2.is = is;
})(TextEdit || (TextEdit = {}));
var ChangeAnnotation;
(function(ChangeAnnotation2) {
  function create(label, needsConfirmation, description) {
    var result = { label };
    if (needsConfirmation !== void 0) {
      result.needsConfirmation = needsConfirmation;
    }
    if (description !== void 0) {
      result.description = description;
    }
    return result;
  }
  ChangeAnnotation2.create = create;
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && Is.objectLiteral(candidate) && Is.string(candidate.label) && (Is.boolean(candidate.needsConfirmation) || candidate.needsConfirmation === void 0) && (Is.string(candidate.description) || candidate.description === void 0);
  }
  ChangeAnnotation2.is = is;
})(ChangeAnnotation || (ChangeAnnotation = {}));
var ChangeAnnotationIdentifier;
(function(ChangeAnnotationIdentifier2) {
  function is(value) {
    var candidate = value;
    return typeof candidate === "string";
  }
  ChangeAnnotationIdentifier2.is = is;
})(ChangeAnnotationIdentifier || (ChangeAnnotationIdentifier = {}));
var AnnotatedTextEdit;
(function(AnnotatedTextEdit2) {
  function replace(range, newText, annotation) {
    return { range, newText, annotationId: annotation };
  }
  AnnotatedTextEdit2.replace = replace;
  function insert(position, newText, annotation) {
    return { range: { start: position, end: position }, newText, annotationId: annotation };
  }
  AnnotatedTextEdit2.insert = insert;
  function del(range, annotation) {
    return { range, newText: "", annotationId: annotation };
  }
  AnnotatedTextEdit2.del = del;
  function is(value) {
    var candidate = value;
    return TextEdit.is(candidate) && (ChangeAnnotation.is(candidate.annotationId) || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  AnnotatedTextEdit2.is = is;
})(AnnotatedTextEdit || (AnnotatedTextEdit = {}));
var TextDocumentEdit;
(function(TextDocumentEdit2) {
  function create(textDocument, edits) {
    return { textDocument, edits };
  }
  TextDocumentEdit2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && OptionalVersionedTextDocumentIdentifier.is(candidate.textDocument) && Array.isArray(candidate.edits);
  }
  TextDocumentEdit2.is = is;
})(TextDocumentEdit || (TextDocumentEdit = {}));
var CreateFile;
(function(CreateFile2) {
  function create(uri, options, annotation) {
    var result = {
      kind: "create",
      uri
    };
    if (options !== void 0 && (options.overwrite !== void 0 || options.ignoreIfExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  CreateFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "create" && Is.string(candidate.uri) && (candidate.options === void 0 || (candidate.options.overwrite === void 0 || Is.boolean(candidate.options.overwrite)) && (candidate.options.ignoreIfExists === void 0 || Is.boolean(candidate.options.ignoreIfExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  CreateFile2.is = is;
})(CreateFile || (CreateFile = {}));
var RenameFile;
(function(RenameFile2) {
  function create(oldUri, newUri, options, annotation) {
    var result = {
      kind: "rename",
      oldUri,
      newUri
    };
    if (options !== void 0 && (options.overwrite !== void 0 || options.ignoreIfExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  RenameFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "rename" && Is.string(candidate.oldUri) && Is.string(candidate.newUri) && (candidate.options === void 0 || (candidate.options.overwrite === void 0 || Is.boolean(candidate.options.overwrite)) && (candidate.options.ignoreIfExists === void 0 || Is.boolean(candidate.options.ignoreIfExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  RenameFile2.is = is;
})(RenameFile || (RenameFile = {}));
var DeleteFile;
(function(DeleteFile2) {
  function create(uri, options, annotation) {
    var result = {
      kind: "delete",
      uri
    };
    if (options !== void 0 && (options.recursive !== void 0 || options.ignoreIfNotExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  DeleteFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "delete" && Is.string(candidate.uri) && (candidate.options === void 0 || (candidate.options.recursive === void 0 || Is.boolean(candidate.options.recursive)) && (candidate.options.ignoreIfNotExists === void 0 || Is.boolean(candidate.options.ignoreIfNotExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  DeleteFile2.is = is;
})(DeleteFile || (DeleteFile = {}));
var WorkspaceEdit;
(function(WorkspaceEdit2) {
  function is(value) {
    var candidate = value;
    return candidate && (candidate.changes !== void 0 || candidate.documentChanges !== void 0) && (candidate.documentChanges === void 0 || candidate.documentChanges.every(function(change) {
      if (Is.string(change.kind)) {
        return CreateFile.is(change) || RenameFile.is(change) || DeleteFile.is(change);
      } else {
        return TextDocumentEdit.is(change);
      }
    }));
  }
  WorkspaceEdit2.is = is;
})(WorkspaceEdit || (WorkspaceEdit = {}));
var TextEditChangeImpl = function() {
  function TextEditChangeImpl2(edits, changeAnnotations) {
    this.edits = edits;
    this.changeAnnotations = changeAnnotations;
  }
  TextEditChangeImpl2.prototype.insert = function(position, newText, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.insert(position, newText);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.insert(position, newText, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.insert(position, newText, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.replace = function(range, newText, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.replace(range, newText);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.replace(range, newText, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.replace(range, newText, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.delete = function(range, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.del(range);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.del(range, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.del(range, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.add = function(edit) {
    this.edits.push(edit);
  };
  TextEditChangeImpl2.prototype.all = function() {
    return this.edits;
  };
  TextEditChangeImpl2.prototype.clear = function() {
    this.edits.splice(0, this.edits.length);
  };
  TextEditChangeImpl2.prototype.assertChangeAnnotations = function(value) {
    if (value === void 0) {
      throw new Error("Text edit change is not configured to manage change annotations.");
    }
  };
  return TextEditChangeImpl2;
}();
var ChangeAnnotations = function() {
  function ChangeAnnotations2(annotations) {
    this._annotations = annotations === void 0 ? /* @__PURE__ */ Object.create(null) : annotations;
    this._counter = 0;
    this._size = 0;
  }
  ChangeAnnotations2.prototype.all = function() {
    return this._annotations;
  };
  Object.defineProperty(ChangeAnnotations2.prototype, "size", {
    get: function() {
      return this._size;
    },
    enumerable: false,
    configurable: true
  });
  ChangeAnnotations2.prototype.manage = function(idOrAnnotation, annotation) {
    var id;
    if (ChangeAnnotationIdentifier.is(idOrAnnotation)) {
      id = idOrAnnotation;
    } else {
      id = this.nextId();
      annotation = idOrAnnotation;
    }
    if (this._annotations[id] !== void 0) {
      throw new Error("Id " + id + " is already in use.");
    }
    if (annotation === void 0) {
      throw new Error("No annotation provided for id " + id);
    }
    this._annotations[id] = annotation;
    this._size++;
    return id;
  };
  ChangeAnnotations2.prototype.nextId = function() {
    this._counter++;
    return this._counter.toString();
  };
  return ChangeAnnotations2;
}();
var WorkspaceChange = function() {
  function WorkspaceChange2(workspaceEdit) {
    var _this = this;
    this._textEditChanges = /* @__PURE__ */ Object.create(null);
    if (workspaceEdit !== void 0) {
      this._workspaceEdit = workspaceEdit;
      if (workspaceEdit.documentChanges) {
        this._changeAnnotations = new ChangeAnnotations(workspaceEdit.changeAnnotations);
        workspaceEdit.changeAnnotations = this._changeAnnotations.all();
        workspaceEdit.documentChanges.forEach(function(change) {
          if (TextDocumentEdit.is(change)) {
            var textEditChange = new TextEditChangeImpl(change.edits, _this._changeAnnotations);
            _this._textEditChanges[change.textDocument.uri] = textEditChange;
          }
        });
      } else if (workspaceEdit.changes) {
        Object.keys(workspaceEdit.changes).forEach(function(key) {
          var textEditChange = new TextEditChangeImpl(workspaceEdit.changes[key]);
          _this._textEditChanges[key] = textEditChange;
        });
      }
    } else {
      this._workspaceEdit = {};
    }
  }
  Object.defineProperty(WorkspaceChange2.prototype, "edit", {
    get: function() {
      this.initDocumentChanges();
      if (this._changeAnnotations !== void 0) {
        if (this._changeAnnotations.size === 0) {
          this._workspaceEdit.changeAnnotations = void 0;
        } else {
          this._workspaceEdit.changeAnnotations = this._changeAnnotations.all();
        }
      }
      return this._workspaceEdit;
    },
    enumerable: false,
    configurable: true
  });
  WorkspaceChange2.prototype.getTextEditChange = function(key) {
    if (OptionalVersionedTextDocumentIdentifier.is(key)) {
      this.initDocumentChanges();
      if (this._workspaceEdit.documentChanges === void 0) {
        throw new Error("Workspace edit is not configured for document changes.");
      }
      var textDocument = { uri: key.uri, version: key.version };
      var result = this._textEditChanges[textDocument.uri];
      if (!result) {
        var edits = [];
        var textDocumentEdit = {
          textDocument,
          edits
        };
        this._workspaceEdit.documentChanges.push(textDocumentEdit);
        result = new TextEditChangeImpl(edits, this._changeAnnotations);
        this._textEditChanges[textDocument.uri] = result;
      }
      return result;
    } else {
      this.initChanges();
      if (this._workspaceEdit.changes === void 0) {
        throw new Error("Workspace edit is not configured for normal text edit changes.");
      }
      var result = this._textEditChanges[key];
      if (!result) {
        var edits = [];
        this._workspaceEdit.changes[key] = edits;
        result = new TextEditChangeImpl(edits);
        this._textEditChanges[key] = result;
      }
      return result;
    }
  };
  WorkspaceChange2.prototype.initDocumentChanges = function() {
    if (this._workspaceEdit.documentChanges === void 0 && this._workspaceEdit.changes === void 0) {
      this._changeAnnotations = new ChangeAnnotations();
      this._workspaceEdit.documentChanges = [];
      this._workspaceEdit.changeAnnotations = this._changeAnnotations.all();
    }
  };
  WorkspaceChange2.prototype.initChanges = function() {
    if (this._workspaceEdit.documentChanges === void 0 && this._workspaceEdit.changes === void 0) {
      this._workspaceEdit.changes = /* @__PURE__ */ Object.create(null);
    }
  };
  WorkspaceChange2.prototype.createFile = function(uri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = CreateFile.create(uri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = CreateFile.create(uri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  WorkspaceChange2.prototype.renameFile = function(oldUri, newUri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = RenameFile.create(oldUri, newUri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = RenameFile.create(oldUri, newUri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  WorkspaceChange2.prototype.deleteFile = function(uri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = DeleteFile.create(uri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = DeleteFile.create(uri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  return WorkspaceChange2;
}();
var TextDocumentIdentifier;
(function(TextDocumentIdentifier2) {
  function create(uri) {
    return { uri };
  }
  TextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri);
  }
  TextDocumentIdentifier2.is = is;
})(TextDocumentIdentifier || (TextDocumentIdentifier = {}));
var VersionedTextDocumentIdentifier;
(function(VersionedTextDocumentIdentifier2) {
  function create(uri, version) {
    return { uri, version };
  }
  VersionedTextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && Is.integer(candidate.version);
  }
  VersionedTextDocumentIdentifier2.is = is;
})(VersionedTextDocumentIdentifier || (VersionedTextDocumentIdentifier = {}));
var OptionalVersionedTextDocumentIdentifier;
(function(OptionalVersionedTextDocumentIdentifier2) {
  function create(uri, version) {
    return { uri, version };
  }
  OptionalVersionedTextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && (candidate.version === null || Is.integer(candidate.version));
  }
  OptionalVersionedTextDocumentIdentifier2.is = is;
})(OptionalVersionedTextDocumentIdentifier || (OptionalVersionedTextDocumentIdentifier = {}));
var TextDocumentItem;
(function(TextDocumentItem2) {
  function create(uri, languageId, version, text) {
    return { uri, languageId, version, text };
  }
  TextDocumentItem2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && Is.string(candidate.languageId) && Is.integer(candidate.version) && Is.string(candidate.text);
  }
  TextDocumentItem2.is = is;
})(TextDocumentItem || (TextDocumentItem = {}));
var MarkupKind;
(function(MarkupKind2) {
  MarkupKind2.PlainText = "plaintext";
  MarkupKind2.Markdown = "markdown";
})(MarkupKind || (MarkupKind = {}));
(function(MarkupKind2) {
  function is(value) {
    var candidate = value;
    return candidate === MarkupKind2.PlainText || candidate === MarkupKind2.Markdown;
  }
  MarkupKind2.is = is;
})(MarkupKind || (MarkupKind = {}));
var MarkupContent;
(function(MarkupContent2) {
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(value) && MarkupKind.is(candidate.kind) && Is.string(candidate.value);
  }
  MarkupContent2.is = is;
})(MarkupContent || (MarkupContent = {}));
var CompletionItemKind;
(function(CompletionItemKind2) {
  CompletionItemKind2.Text = 1;
  CompletionItemKind2.Method = 2;
  CompletionItemKind2.Function = 3;
  CompletionItemKind2.Constructor = 4;
  CompletionItemKind2.Field = 5;
  CompletionItemKind2.Variable = 6;
  CompletionItemKind2.Class = 7;
  CompletionItemKind2.Interface = 8;
  CompletionItemKind2.Module = 9;
  CompletionItemKind2.Property = 10;
  CompletionItemKind2.Unit = 11;
  CompletionItemKind2.Value = 12;
  CompletionItemKind2.Enum = 13;
  CompletionItemKind2.Keyword = 14;
  CompletionItemKind2.Snippet = 15;
  CompletionItemKind2.Color = 16;
  CompletionItemKind2.File = 17;
  CompletionItemKind2.Reference = 18;
  CompletionItemKind2.Folder = 19;
  CompletionItemKind2.EnumMember = 20;
  CompletionItemKind2.Constant = 21;
  CompletionItemKind2.Struct = 22;
  CompletionItemKind2.Event = 23;
  CompletionItemKind2.Operator = 24;
  CompletionItemKind2.TypeParameter = 25;
})(CompletionItemKind || (CompletionItemKind = {}));
var InsertTextFormat;
(function(InsertTextFormat2) {
  InsertTextFormat2.PlainText = 1;
  InsertTextFormat2.Snippet = 2;
})(InsertTextFormat || (InsertTextFormat = {}));
var CompletionItemTag;
(function(CompletionItemTag2) {
  CompletionItemTag2.Deprecated = 1;
})(CompletionItemTag || (CompletionItemTag = {}));
var InsertReplaceEdit;
(function(InsertReplaceEdit2) {
  function create(newText, insert, replace) {
    return { newText, insert, replace };
  }
  InsertReplaceEdit2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.newText) && Range.is(candidate.insert) && Range.is(candidate.replace);
  }
  InsertReplaceEdit2.is = is;
})(InsertReplaceEdit || (InsertReplaceEdit = {}));
var InsertTextMode;
(function(InsertTextMode2) {
  InsertTextMode2.asIs = 1;
  InsertTextMode2.adjustIndentation = 2;
})(InsertTextMode || (InsertTextMode = {}));
var CompletionItem;
(function(CompletionItem2) {
  function create(label) {
    return { label };
  }
  CompletionItem2.create = create;
})(CompletionItem || (CompletionItem = {}));
var CompletionList;
(function(CompletionList2) {
  function create(items, isIncomplete) {
    return { items: items ? items : [], isIncomplete: !!isIncomplete };
  }
  CompletionList2.create = create;
})(CompletionList || (CompletionList = {}));
var MarkedString;
(function(MarkedString2) {
  function fromPlainText(plainText) {
    return plainText.replace(/[\\`*_{}[\]()#+\-.!]/g, "\\$&");
  }
  MarkedString2.fromPlainText = fromPlainText;
  function is(value) {
    var candidate = value;
    return Is.string(candidate) || Is.objectLiteral(candidate) && Is.string(candidate.language) && Is.string(candidate.value);
  }
  MarkedString2.is = is;
})(MarkedString || (MarkedString = {}));
var Hover;
(function(Hover2) {
  function is(value) {
    var candidate = value;
    return !!candidate && Is.objectLiteral(candidate) && (MarkupContent.is(candidate.contents) || MarkedString.is(candidate.contents) || Is.typedArray(candidate.contents, MarkedString.is)) && (value.range === void 0 || Range.is(value.range));
  }
  Hover2.is = is;
})(Hover || (Hover = {}));
var ParameterInformation;
(function(ParameterInformation2) {
  function create(label, documentation) {
    return documentation ? { label, documentation } : { label };
  }
  ParameterInformation2.create = create;
})(ParameterInformation || (ParameterInformation = {}));
var SignatureInformation;
(function(SignatureInformation2) {
  function create(label, documentation) {
    var parameters = [];
    for (var _i = 2; _i < arguments.length; _i++) {
      parameters[_i - 2] = arguments[_i];
    }
    var result = { label };
    if (Is.defined(documentation)) {
      result.documentation = documentation;
    }
    if (Is.defined(parameters)) {
      result.parameters = parameters;
    } else {
      result.parameters = [];
    }
    return result;
  }
  SignatureInformation2.create = create;
})(SignatureInformation || (SignatureInformation = {}));
var DocumentHighlightKind;
(function(DocumentHighlightKind2) {
  DocumentHighlightKind2.Text = 1;
  DocumentHighlightKind2.Read = 2;
  DocumentHighlightKind2.Write = 3;
})(DocumentHighlightKind || (DocumentHighlightKind = {}));
var DocumentHighlight;
(function(DocumentHighlight2) {
  function create(range, kind) {
    var result = { range };
    if (Is.number(kind)) {
      result.kind = kind;
    }
    return result;
  }
  DocumentHighlight2.create = create;
})(DocumentHighlight || (DocumentHighlight = {}));
var SymbolKind;
(function(SymbolKind2) {
  SymbolKind2.File = 1;
  SymbolKind2.Module = 2;
  SymbolKind2.Namespace = 3;
  SymbolKind2.Package = 4;
  SymbolKind2.Class = 5;
  SymbolKind2.Method = 6;
  SymbolKind2.Property = 7;
  SymbolKind2.Field = 8;
  SymbolKind2.Constructor = 9;
  SymbolKind2.Enum = 10;
  SymbolKind2.Interface = 11;
  SymbolKind2.Function = 12;
  SymbolKind2.Variable = 13;
  SymbolKind2.Constant = 14;
  SymbolKind2.String = 15;
  SymbolKind2.Number = 16;
  SymbolKind2.Boolean = 17;
  SymbolKind2.Array = 18;
  SymbolKind2.Object = 19;
  SymbolKind2.Key = 20;
  SymbolKind2.Null = 21;
  SymbolKind2.EnumMember = 22;
  SymbolKind2.Struct = 23;
  SymbolKind2.Event = 24;
  SymbolKind2.Operator = 25;
  SymbolKind2.TypeParameter = 26;
})(SymbolKind || (SymbolKind = {}));
var SymbolTag;
(function(SymbolTag2) {
  SymbolTag2.Deprecated = 1;
})(SymbolTag || (SymbolTag = {}));
var SymbolInformation;
(function(SymbolInformation2) {
  function create(name, kind, range, uri, containerName) {
    var result = {
      name,
      kind,
      location: { uri, range }
    };
    if (containerName) {
      result.containerName = containerName;
    }
    return result;
  }
  SymbolInformation2.create = create;
})(SymbolInformation || (SymbolInformation = {}));
var DocumentSymbol;
(function(DocumentSymbol2) {
  function create(name, detail, kind, range, selectionRange, children) {
    var result = {
      name,
      detail,
      kind,
      range,
      selectionRange
    };
    if (children !== void 0) {
      result.children = children;
    }
    return result;
  }
  DocumentSymbol2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.name) && Is.number(candidate.kind) && Range.is(candidate.range) && Range.is(candidate.selectionRange) && (candidate.detail === void 0 || Is.string(candidate.detail)) && (candidate.deprecated === void 0 || Is.boolean(candidate.deprecated)) && (candidate.children === void 0 || Array.isArray(candidate.children)) && (candidate.tags === void 0 || Array.isArray(candidate.tags));
  }
  DocumentSymbol2.is = is;
})(DocumentSymbol || (DocumentSymbol = {}));
var CodeActionKind;
(function(CodeActionKind2) {
  CodeActionKind2.Empty = "";
  CodeActionKind2.QuickFix = "quickfix";
  CodeActionKind2.Refactor = "refactor";
  CodeActionKind2.RefactorExtract = "refactor.extract";
  CodeActionKind2.RefactorInline = "refactor.inline";
  CodeActionKind2.RefactorRewrite = "refactor.rewrite";
  CodeActionKind2.Source = "source";
  CodeActionKind2.SourceOrganizeImports = "source.organizeImports";
  CodeActionKind2.SourceFixAll = "source.fixAll";
})(CodeActionKind || (CodeActionKind = {}));
var CodeActionContext;
(function(CodeActionContext2) {
  function create(diagnostics, only) {
    var result = { diagnostics };
    if (only !== void 0 && only !== null) {
      result.only = only;
    }
    return result;
  }
  CodeActionContext2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.typedArray(candidate.diagnostics, Diagnostic.is) && (candidate.only === void 0 || Is.typedArray(candidate.only, Is.string));
  }
  CodeActionContext2.is = is;
})(CodeActionContext || (CodeActionContext = {}));
var CodeAction;
(function(CodeAction2) {
  function create(title, kindOrCommandOrEdit, kind) {
    var result = { title };
    var checkKind = true;
    if (typeof kindOrCommandOrEdit === "string") {
      checkKind = false;
      result.kind = kindOrCommandOrEdit;
    } else if (Command.is(kindOrCommandOrEdit)) {
      result.command = kindOrCommandOrEdit;
    } else {
      result.edit = kindOrCommandOrEdit;
    }
    if (checkKind && kind !== void 0) {
      result.kind = kind;
    }
    return result;
  }
  CodeAction2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.title) && (candidate.diagnostics === void 0 || Is.typedArray(candidate.diagnostics, Diagnostic.is)) && (candidate.kind === void 0 || Is.string(candidate.kind)) && (candidate.edit !== void 0 || candidate.command !== void 0) && (candidate.command === void 0 || Command.is(candidate.command)) && (candidate.isPreferred === void 0 || Is.boolean(candidate.isPreferred)) && (candidate.edit === void 0 || WorkspaceEdit.is(candidate.edit));
  }
  CodeAction2.is = is;
})(CodeAction || (CodeAction = {}));
var CodeLens;
(function(CodeLens2) {
  function create(range, data) {
    var result = { range };
    if (Is.defined(data)) {
      result.data = data;
    }
    return result;
  }
  CodeLens2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.undefined(candidate.command) || Command.is(candidate.command));
  }
  CodeLens2.is = is;
})(CodeLens || (CodeLens = {}));
var FormattingOptions;
(function(FormattingOptions2) {
  function create(tabSize, insertSpaces) {
    return { tabSize, insertSpaces };
  }
  FormattingOptions2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.uinteger(candidate.tabSize) && Is.boolean(candidate.insertSpaces);
  }
  FormattingOptions2.is = is;
})(FormattingOptions || (FormattingOptions = {}));
var DocumentLink;
(function(DocumentLink2) {
  function create(range, target, data) {
    return { range, target, data };
  }
  DocumentLink2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.undefined(candidate.target) || Is.string(candidate.target));
  }
  DocumentLink2.is = is;
})(DocumentLink || (DocumentLink = {}));
var SelectionRange;
(function(SelectionRange2) {
  function create(range, parent) {
    return { range, parent };
  }
  SelectionRange2.create = create;
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && Range.is(candidate.range) && (candidate.parent === void 0 || SelectionRange2.is(candidate.parent));
  }
  SelectionRange2.is = is;
})(SelectionRange || (SelectionRange = {}));
var TextDocument;
(function(TextDocument2) {
  function create(uri, languageId, version, content) {
    return new FullTextDocument(uri, languageId, version, content);
  }
  TextDocument2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && (Is.undefined(candidate.languageId) || Is.string(candidate.languageId)) && Is.uinteger(candidate.lineCount) && Is.func(candidate.getText) && Is.func(candidate.positionAt) && Is.func(candidate.offsetAt) ? true : false;
  }
  TextDocument2.is = is;
  function applyEdits(document, edits) {
    var text = document.getText();
    var sortedEdits = mergeSort(edits, function(a, b) {
      var diff = a.range.start.line - b.range.start.line;
      if (diff === 0) {
        return a.range.start.character - b.range.start.character;
      }
      return diff;
    });
    var lastModifiedOffset = text.length;
    for (var i = sortedEdits.length - 1; i >= 0; i--) {
      var e = sortedEdits[i];
      var startOffset = document.offsetAt(e.range.start);
      var endOffset = document.offsetAt(e.range.end);
      if (endOffset <= lastModifiedOffset) {
        text = text.substring(0, startOffset) + e.newText + text.substring(endOffset, text.length);
      } else {
        throw new Error("Overlapping edit");
      }
      lastModifiedOffset = startOffset;
    }
    return text;
  }
  TextDocument2.applyEdits = applyEdits;
  function mergeSort(data, compare) {
    if (data.length <= 1) {
      return data;
    }
    var p = data.length / 2 | 0;
    var left = data.slice(0, p);
    var right = data.slice(p);
    mergeSort(left, compare);
    mergeSort(right, compare);
    var leftIdx = 0;
    var rightIdx = 0;
    var i = 0;
    while (leftIdx < left.length && rightIdx < right.length) {
      var ret = compare(left[leftIdx], right[rightIdx]);
      if (ret <= 0) {
        data[i++] = left[leftIdx++];
      } else {
        data[i++] = right[rightIdx++];
      }
    }
    while (leftIdx < left.length) {
      data[i++] = left[leftIdx++];
    }
    while (rightIdx < right.length) {
      data[i++] = right[rightIdx++];
    }
    return data;
  }
})(TextDocument || (TextDocument = {}));
var FullTextDocument = function() {
  function FullTextDocument2(uri, languageId, version, content) {
    this._uri = uri;
    this._languageId = languageId;
    this._version = version;
    this._content = content;
    this._lineOffsets = void 0;
  }
  Object.defineProperty(FullTextDocument2.prototype, "uri", {
    get: function() {
      return this._uri;
    },
    enumerable: false,
    configurable: true
  });
  Object.defineProperty(FullTextDocument2.prototype, "languageId", {
    get: function() {
      return this._languageId;
    },
    enumerable: false,
    configurable: true
  });
  Object.defineProperty(FullTextDocument2.prototype, "version", {
    get: function() {
      return this._version;
    },
    enumerable: false,
    configurable: true
  });
  FullTextDocument2.prototype.getText = function(range) {
    if (range) {
      var start = this.offsetAt(range.start);
      var end = this.offsetAt(range.end);
      return this._content.substring(start, end);
    }
    return this._content;
  };
  FullTextDocument2.prototype.update = function(event, version) {
    this._content = event.text;
    this._version = version;
    this._lineOffsets = void 0;
  };
  FullTextDocument2.prototype.getLineOffsets = function() {
    if (this._lineOffsets === void 0) {
      var lineOffsets = [];
      var text = this._content;
      var isLineStart = true;
      for (var i = 0; i < text.length; i++) {
        if (isLineStart) {
          lineOffsets.push(i);
          isLineStart = false;
        }
        var ch = text.charAt(i);
        isLineStart = ch === "\r" || ch === "\n";
        if (ch === "\r" && i + 1 < text.length && text.charAt(i + 1) === "\n") {
          i++;
        }
      }
      if (isLineStart && text.length > 0) {
        lineOffsets.push(text.length);
      }
      this._lineOffsets = lineOffsets;
    }
    return this._lineOffsets;
  };
  FullTextDocument2.prototype.positionAt = function(offset) {
    offset = Math.max(Math.min(offset, this._content.length), 0);
    var lineOffsets = this.getLineOffsets();
    var low = 0, high = lineOffsets.length;
    if (high === 0) {
      return Position.create(0, offset);
    }
    while (low < high) {
      var mid = Math.floor((low + high) / 2);
      if (lineOffsets[mid] > offset) {
        high = mid;
      } else {
        low = mid + 1;
      }
    }
    var line = low - 1;
    return Position.create(line, offset - lineOffsets[line]);
  };
  FullTextDocument2.prototype.offsetAt = function(position) {
    var lineOffsets = this.getLineOffsets();
    if (position.line >= lineOffsets.length) {
      return this._content.length;
    } else if (position.line < 0) {
      return 0;
    }
    var lineOffset = lineOffsets[position.line];
    var nextLineOffset = position.line + 1 < lineOffsets.length ? lineOffsets[position.line + 1] : this._content.length;
    return Math.max(Math.min(lineOffset + position.character, nextLineOffset), lineOffset);
  };
  Object.defineProperty(FullTextDocument2.prototype, "lineCount", {
    get: function() {
      return this.getLineOffsets().length;
    },
    enumerable: false,
    configurable: true
  });
  return FullTextDocument2;
}();
var Is;
(function(Is2) {
  var toString = Object.prototype.toString;
  function defined(value) {
    return typeof value !== "undefined";
  }
  Is2.defined = defined;
  function undefined2(value) {
    return typeof value === "undefined";
  }
  Is2.undefined = undefined2;
  function boolean(value) {
    return value === true || value === false;
  }
  Is2.boolean = boolean;
  function string(value) {
    return toString.call(value) === "[object String]";
  }
  Is2.string = string;
  function number(value) {
    return toString.call(value) === "[object Number]";
  }
  Is2.number = number;
  function numberRange(value, min, max) {
    return toString.call(value) === "[object Number]" && min <= value && value <= max;
  }
  Is2.numberRange = numberRange;
  function integer2(value) {
    return toString.call(value) === "[object Number]" && -2147483648 <= value && value <= 2147483647;
  }
  Is2.integer = integer2;
  function uinteger2(value) {
    return toString.call(value) === "[object Number]" && 0 <= value && value <= 2147483647;
  }
  Is2.uinteger = uinteger2;
  function func(value) {
    return toString.call(value) === "[object Function]";
  }
  Is2.func = func;
  function objectLiteral(value) {
    return value !== null && typeof value === "object";
  }
  Is2.objectLiteral = objectLiteral;
  function typedArray(value, check) {
    return Array.isArray(value) && value.every(check);
  }
  Is2.typedArray = typedArray;
})(Is || (Is = {}));

// src/language/common/lspLanguageFeatures.ts
var DiagnosticsAdapter = class {
  constructor(_languageId, _worker, configChangeEvent) {
    this._languageId = _languageId;
    this._worker = _worker;
    const onModelAdd = (model) => {
      let modeId = model.getLanguageId();
      if (modeId !== this._languageId) {
        return;
      }
      let handle;
      this._listener[model.uri.toString()] = model.onDidChangeContent(() => {
        window.clearTimeout(handle);
        handle = window.setTimeout(() => this._doValidate(model.uri, modeId), 500);
      });
      this._doValidate(model.uri, modeId);
    };
    const onModelRemoved = (model) => {
      monaco_editor_core_exports.editor.setModelMarkers(model, this._languageId, []);
      let uriStr = model.uri.toString();
      let listener = this._listener[uriStr];
      if (listener) {
        listener.dispose();
        delete this._listener[uriStr];
      }
    };
    this._disposables.push(monaco_editor_core_exports.editor.onDidCreateModel(onModelAdd));
    this._disposables.push(monaco_editor_core_exports.editor.onWillDisposeModel(onModelRemoved));
    this._disposables.push(monaco_editor_core_exports.editor.onDidChangeModelLanguage((event) => {
      onModelRemoved(event.model);
      onModelAdd(event.model);
    }));
    this._disposables.push(configChangeEvent((_) => {
      monaco_editor_core_exports.editor.getModels().forEach((model) => {
        if (model.getLanguageId() === this._languageId) {
          onModelRemoved(model);
          onModelAdd(model);
        }
      });
    }));
    this._disposables.push({
      dispose: () => {
        monaco_editor_core_exports.editor.getModels().forEach(onModelRemoved);
        for (let key in this._listener) {
          this._listener[key].dispose();
        }
      }
    });
    monaco_editor_core_exports.editor.getModels().forEach(onModelAdd);
  }
  _disposables = [];
  _listener = /* @__PURE__ */ Object.create(null);
  dispose() {
    this._disposables.forEach((d) => d && d.dispose());
    this._disposables.length = 0;
  }
  _doValidate(resource, languageId) {
    this._worker(resource).then((worker) => {
      return worker.doValidation(resource.toString());
    }).then((diagnostics) => {
      const markers = diagnostics.map((d) => toDiagnostics(resource, d));
      let model = monaco_editor_core_exports.editor.getModel(resource);
      if (model && model.getLanguageId() === languageId) {
        monaco_editor_core_exports.editor.setModelMarkers(model, languageId, markers);
      }
    }).then(void 0, (err) => {
      console.error(err);
    });
  }
};
function toSeverity(lsSeverity) {
  switch (lsSeverity) {
    case DiagnosticSeverity.Error:
      return monaco_editor_core_exports.MarkerSeverity.Error;
    case DiagnosticSeverity.Warning:
      return monaco_editor_core_exports.MarkerSeverity.Warning;
    case DiagnosticSeverity.Information:
      return monaco_editor_core_exports.MarkerSeverity.Info;
    case DiagnosticSeverity.Hint:
      return monaco_editor_core_exports.MarkerSeverity.Hint;
    default:
      return monaco_editor_core_exports.MarkerSeverity.Info;
  }
}
function toDiagnostics(resource, diag) {
  let code = typeof diag.code === "number" ? String(diag.code) : diag.code;
  return {
    severity: toSeverity(diag.severity),
    startLineNumber: diag.range.start.line + 1,
    startColumn: diag.range.start.character + 1,
    endLineNumber: diag.range.end.line + 1,
    endColumn: diag.range.end.character + 1,
    message: diag.message,
    code,
    source: diag.source
  };
}
var CompletionAdapter = class {
  constructor(_worker, _triggerCharacters) {
    this._worker = _worker;
    this._triggerCharacters = _triggerCharacters;
  }
  get triggerCharacters() {
    return this._triggerCharacters;
  }
  provideCompletionItems(model, position, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doComplete(resource.toString(), fromPosition(position));
    }).then((info) => {
      if (!info) {
        return;
      }
      const wordInfo = model.getWordUntilPosition(position);
      const wordRange = new monaco_editor_core_exports.Range(position.lineNumber, wordInfo.startColumn, position.lineNumber, wordInfo.endColumn);
      const items = info.items.map((entry) => {
        const item = {
          label: entry.label,
          insertText: entry.insertText || entry.label,
          sortText: entry.sortText,
          filterText: entry.filterText,
          documentation: entry.documentation,
          detail: entry.detail,
          command: toCommand(entry.command),
          range: wordRange,
          kind: toCompletionItemKind(entry.kind)
        };
        if (entry.textEdit) {
          if (isInsertReplaceEdit(entry.textEdit)) {
            item.range = {
              insert: toRange(entry.textEdit.insert),
              replace: toRange(entry.textEdit.replace)
            };
          } else {
            item.range = toRange(entry.textEdit.range);
          }
          item.insertText = entry.textEdit.newText;
        }
        if (entry.additionalTextEdits) {
          item.additionalTextEdits = entry.additionalTextEdits.map(toTextEdit);
        }
        if (entry.insertTextFormat === InsertTextFormat.Snippet) {
          item.insertTextRules = monaco_editor_core_exports.languages.CompletionItemInsertTextRule.InsertAsSnippet;
        }
        return item;
      });
      return {
        isIncomplete: info.isIncomplete,
        suggestions: items
      };
    });
  }
};
function fromPosition(position) {
  if (!position) {
    return void 0;
  }
  return { character: position.column - 1, line: position.lineNumber - 1 };
}
function fromRange(range) {
  if (!range) {
    return void 0;
  }
  return {
    start: {
      line: range.startLineNumber - 1,
      character: range.startColumn - 1
    },
    end: { line: range.endLineNumber - 1, character: range.endColumn - 1 }
  };
}
function toRange(range) {
  if (!range) {
    return void 0;
  }
  return new monaco_editor_core_exports.Range(range.start.line + 1, range.start.character + 1, range.end.line + 1, range.end.character + 1);
}
function isInsertReplaceEdit(edit) {
  return typeof edit.insert !== "undefined" && typeof edit.replace !== "undefined";
}
function toCompletionItemKind(kind) {
  const mItemKind = monaco_editor_core_exports.languages.CompletionItemKind;
  switch (kind) {
    case CompletionItemKind.Text:
      return mItemKind.Text;
    case CompletionItemKind.Method:
      return mItemKind.Method;
    case CompletionItemKind.Function:
      return mItemKind.Function;
    case CompletionItemKind.Constructor:
      return mItemKind.Constructor;
    case CompletionItemKind.Field:
      return mItemKind.Field;
    case CompletionItemKind.Variable:
      return mItemKind.Variable;
    case CompletionItemKind.Class:
      return mItemKind.Class;
    case CompletionItemKind.Interface:
      return mItemKind.Interface;
    case CompletionItemKind.Module:
      return mItemKind.Module;
    case CompletionItemKind.Property:
      return mItemKind.Property;
    case CompletionItemKind.Unit:
      return mItemKind.Unit;
    case CompletionItemKind.Value:
      return mItemKind.Value;
    case CompletionItemKind.Enum:
      return mItemKind.Enum;
    case CompletionItemKind.Keyword:
      return mItemKind.Keyword;
    case CompletionItemKind.Snippet:
      return mItemKind.Snippet;
    case CompletionItemKind.Color:
      return mItemKind.Color;
    case CompletionItemKind.File:
      return mItemKind.File;
    case CompletionItemKind.Reference:
      return mItemKind.Reference;
  }
  return mItemKind.Property;
}
function toTextEdit(textEdit) {
  if (!textEdit) {
    return void 0;
  }
  return {
    range: toRange(textEdit.range),
    text: textEdit.newText
  };
}
function toCommand(c) {
  return c && c.command === "editor.action.triggerSuggest" ? { id: c.command, title: c.title, arguments: c.arguments } : void 0;
}
var HoverAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideHover(model, position, token) {
    let resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doHover(resource.toString(), fromPosition(position));
    }).then((info) => {
      if (!info) {
        return;
      }
      return {
        range: toRange(info.range),
        contents: toMarkedStringArray(info.contents)
      };
    });
  }
};
function isMarkupContent(thing) {
  return thing && typeof thing === "object" && typeof thing.kind === "string";
}
function toMarkdownString(entry) {
  if (typeof entry === "string") {
    return {
      value: entry
    };
  }
  if (isMarkupContent(entry)) {
    if (entry.kind === "plaintext") {
      return {
        value: entry.value.replace(/[\\`*_{}[\]()#+\-.!]/g, "\\$&")
      };
    }
    return {
      value: entry.value
    };
  }
  return { value: "```" + entry.language + "\n" + entry.value + "\n```\n" };
}
function toMarkedStringArray(contents) {
  if (!contents) {
    return void 0;
  }
  if (Array.isArray(contents)) {
    return contents.map(toMarkdownString);
  }
  return [toMarkdownString(contents)];
}
var DocumentHighlightAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentHighlights(model, position, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentHighlights(resource.toString(), fromPosition(position))).then((entries) => {
      if (!entries) {
        return;
      }
      return entries.map((entry) => {
        return {
          range: toRange(entry.range),
          kind: toDocumentHighlightKind(entry.kind)
        };
      });
    });
  }
};
function toDocumentHighlightKind(kind) {
  switch (kind) {
    case DocumentHighlightKind.Read:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Read;
    case DocumentHighlightKind.Write:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Write;
    case DocumentHighlightKind.Text:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Text;
  }
  return monaco_editor_core_exports.languages.DocumentHighlightKind.Text;
}
var DefinitionAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDefinition(model, position, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.findDefinition(resource.toString(), fromPosition(position));
    }).then((definition) => {
      if (!definition) {
        return;
      }
      return [toLocation(definition)];
    });
  }
};
function toLocation(location) {
  return {
    uri: monaco_editor_core_exports.Uri.parse(location.uri),
    range: toRange(location.range)
  };
}
var ReferenceAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideReferences(model, position, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.findReferences(resource.toString(), fromPosition(position));
    }).then((entries) => {
      if (!entries) {
        return;
      }
      return entries.map(toLocation);
    });
  }
};
var RenameAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideRenameEdits(model, position, newName, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doRename(resource.toString(), fromPosition(position), newName);
    }).then((edit) => {
      return toWorkspaceEdit(edit);
    });
  }
};
function toWorkspaceEdit(edit) {
  if (!edit || !edit.changes) {
    return void 0;
  }
  let resourceEdits = [];
  for (let uri in edit.changes) {
    const _uri = monaco_editor_core_exports.Uri.parse(uri);
    for (let e of edit.changes[uri]) {
      resourceEdits.push({
        resource: _uri,
        versionId: void 0,
        textEdit: {
          range: toRange(e.range),
          text: e.newText
        }
      });
    }
  }
  return {
    edits: resourceEdits
  };
}
var DocumentSymbolAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentSymbols(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentSymbols(resource.toString())).then((items) => {
      if (!items) {
        return;
      }
      return items.map((item) => ({
        name: item.name,
        detail: "",
        containerName: item.containerName,
        kind: toSymbolKind(item.kind),
        range: toRange(item.location.range),
        selectionRange: toRange(item.location.range),
        tags: []
      }));
    });
  }
};
function toSymbolKind(kind) {
  let mKind = monaco_editor_core_exports.languages.SymbolKind;
  switch (kind) {
    case SymbolKind.File:
      return mKind.Array;
    case SymbolKind.Module:
      return mKind.Module;
    case SymbolKind.Namespace:
      return mKind.Namespace;
    case SymbolKind.Package:
      return mKind.Package;
    case SymbolKind.Class:
      return mKind.Class;
    case SymbolKind.Method:
      return mKind.Method;
    case SymbolKind.Property:
      return mKind.Property;
    case SymbolKind.Field:
      return mKind.Field;
    case SymbolKind.Constructor:
      return mKind.Constructor;
    case SymbolKind.Enum:
      return mKind.Enum;
    case SymbolKind.Interface:
      return mKind.Interface;
    case SymbolKind.Function:
      return mKind.Function;
    case SymbolKind.Variable:
      return mKind.Variable;
    case SymbolKind.Constant:
      return mKind.Constant;
    case SymbolKind.String:
      return mKind.String;
    case SymbolKind.Number:
      return mKind.Number;
    case SymbolKind.Boolean:
      return mKind.Boolean;
    case SymbolKind.Array:
      return mKind.Array;
  }
  return mKind.Function;
}
var DocumentLinkAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideLinks(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentLinks(resource.toString())).then((items) => {
      if (!items) {
        return;
      }
      return {
        links: items.map((item) => ({
          range: toRange(item.range),
          url: item.target
        }))
      };
    });
  }
};
var DocumentFormattingEditProvider = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentFormattingEdits(model, options, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.format(resource.toString(), null, fromFormattingOptions(options)).then((edits) => {
        if (!edits || edits.length === 0) {
          return;
        }
        return edits.map(toTextEdit);
      });
    });
  }
};
var DocumentRangeFormattingEditProvider = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  canFormatMultipleRanges = false;
  provideDocumentRangeFormattingEdits(model, range, options, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.format(resource.toString(), fromRange(range), fromFormattingOptions(options)).then((edits) => {
        if (!edits || edits.length === 0) {
          return;
        }
        return edits.map(toTextEdit);
      });
    });
  }
};
function fromFormattingOptions(options) {
  return {
    tabSize: options.tabSize,
    insertSpaces: options.insertSpaces
  };
}
var DocumentColorAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentColors(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentColors(resource.toString())).then((infos) => {
      if (!infos) {
        return;
      }
      return infos.map((item) => ({
        color: item.color,
        range: toRange(item.range)
      }));
    });
  }
  provideColorPresentations(model, info, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getColorPresentations(resource.toString(), info.color, fromRange(info.range))).then((presentations) => {
      if (!presentations) {
        return;
      }
      return presentations.map((presentation) => {
        let item = {
          label: presentation.label
        };
        if (presentation.textEdit) {
          item.textEdit = toTextEdit(presentation.textEdit);
        }
        if (presentation.additionalTextEdits) {
          item.additionalTextEdits = presentation.additionalTextEdits.map(toTextEdit);
        }
        return item;
      });
    });
  }
};
var FoldingRangeAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideFoldingRanges(model, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getFoldingRanges(resource.toString(), context)).then((ranges) => {
      if (!ranges) {
        return;
      }
      return ranges.map((range) => {
        const result = {
          start: range.startLine + 1,
          end: range.endLine + 1
        };
        if (typeof range.kind !== "undefined") {
          result.kind = toFoldingRangeKind(range.kind);
        }
        return result;
      });
    });
  }
};
function toFoldingRangeKind(kind) {
  switch (kind) {
    case FoldingRangeKind.Comment:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Comment;
    case FoldingRangeKind.Imports:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Imports;
    case FoldingRangeKind.Region:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Region;
  }
  return void 0;
}
var SelectionRangeAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideSelectionRanges(model, positions, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getSelectionRanges(resource.toString(), positions.map(fromPosition))).then((selectionRanges) => {
      if (!selectionRanges) {
        return;
      }
      return selectionRanges.map((selectionRange) => {
        const result = [];
        while (selectionRange) {
          result.push({ range: toRange(selectionRange.range) });
          selectionRange = selectionRange.parent;
        }
        return result;
      });
    });
  }
};

// src/language/html/htmlMode.ts
var HTMLCompletionAdapter = class extends CompletionAdapter {
  constructor(worker) {
    super(worker, [".", ":", "<", '"', "=", "/"]);
  }
};
function setupMode1(defaults) {
  const client = new WorkerManager(defaults);
  const worker = (...uris) => {
    return client.getLanguageServiceWorker(...uris);
  };
  let languageId = defaults.languageId;
  monaco_editor_core_exports.languages.registerCompletionItemProvider(languageId, new HTMLCompletionAdapter(worker));
  monaco_editor_core_exports.languages.registerHoverProvider(languageId, new HoverAdapter(worker));
  monaco_editor_core_exports.languages.registerDocumentHighlightProvider(languageId, new DocumentHighlightAdapter(worker));
  monaco_editor_core_exports.languages.registerLinkProvider(languageId, new DocumentLinkAdapter(worker));
  monaco_editor_core_exports.languages.registerFoldingRangeProvider(languageId, new FoldingRangeAdapter(worker));
  monaco_editor_core_exports.languages.registerDocumentSymbolProvider(languageId, new DocumentSymbolAdapter(worker));
  monaco_editor_core_exports.languages.registerSelectionRangeProvider(languageId, new SelectionRangeAdapter(worker));
  monaco_editor_core_exports.languages.registerRenameProvider(languageId, new RenameAdapter(worker));
  if (languageId === "html") {
    monaco_editor_core_exports.languages.registerDocumentFormattingEditProvider(languageId, new DocumentFormattingEditProvider(worker));
    monaco_editor_core_exports.languages.registerDocumentRangeFormattingEditProvider(languageId, new DocumentRangeFormattingEditProvider(worker));
  }
}
function setupMode(defaults) {
  const disposables = [];
  const providers = [];
  const client = new WorkerManager(defaults);
  disposables.push(client);
  const worker = (...uris) => {
    return client.getLanguageServiceWorker(...uris);
  };
  function registerProviders() {
    const { languageId, modeConfiguration } = defaults;
    disposeAll(providers);
    if (modeConfiguration.completionItems) {
      providers.push(monaco_editor_core_exports.languages.registerCompletionItemProvider(languageId, new HTMLCompletionAdapter(worker)));
    }
    if (modeConfiguration.hovers) {
      providers.push(monaco_editor_core_exports.languages.registerHoverProvider(languageId, new HoverAdapter(worker)));
    }
    if (modeConfiguration.documentHighlights) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentHighlightProvider(languageId, new DocumentHighlightAdapter(worker)));
    }
    if (modeConfiguration.links) {
      providers.push(monaco_editor_core_exports.languages.registerLinkProvider(languageId, new DocumentLinkAdapter(worker)));
    }
    if (modeConfiguration.documentSymbols) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentSymbolProvider(languageId, new DocumentSymbolAdapter(worker)));
    }
    if (modeConfiguration.rename) {
      providers.push(monaco_editor_core_exports.languages.registerRenameProvider(languageId, new RenameAdapter(worker)));
    }
    if (modeConfiguration.foldingRanges) {
      providers.push(monaco_editor_core_exports.languages.registerFoldingRangeProvider(languageId, new FoldingRangeAdapter(worker)));
    }
    if (modeConfiguration.selectionRanges) {
      providers.push(monaco_editor_core_exports.languages.registerSelectionRangeProvider(languageId, new SelectionRangeAdapter(worker)));
    }
    if (modeConfiguration.documentFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentFormattingEditProvider(languageId, new DocumentFormattingEditProvider(worker)));
    }
    if (modeConfiguration.documentRangeFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentRangeFormattingEditProvider(languageId, new DocumentRangeFormattingEditProvider(worker)));
    }
  }
  registerProviders();
  disposables.push(asDisposable(providers));
  return asDisposable(disposables);
}
function asDisposable(disposables) {
  return { dispose: () => disposeAll(disposables) };
}
function disposeAll(disposables) {
  while (disposables.length) {
    disposables.pop().dispose();
  }
}



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/language/json/jsonMode.js":
/*!*********************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/language/json/jsonMode.js ***!
  \*********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   CompletionAdapter: function() { return /* binding */ CompletionAdapter; },
/* harmony export */   DefinitionAdapter: function() { return /* binding */ DefinitionAdapter; },
/* harmony export */   DiagnosticsAdapter: function() { return /* binding */ DiagnosticsAdapter; },
/* harmony export */   DocumentColorAdapter: function() { return /* binding */ DocumentColorAdapter; },
/* harmony export */   DocumentFormattingEditProvider: function() { return /* binding */ DocumentFormattingEditProvider; },
/* harmony export */   DocumentHighlightAdapter: function() { return /* binding */ DocumentHighlightAdapter; },
/* harmony export */   DocumentLinkAdapter: function() { return /* binding */ DocumentLinkAdapter; },
/* harmony export */   DocumentRangeFormattingEditProvider: function() { return /* binding */ DocumentRangeFormattingEditProvider; },
/* harmony export */   DocumentSymbolAdapter: function() { return /* binding */ DocumentSymbolAdapter; },
/* harmony export */   FoldingRangeAdapter: function() { return /* binding */ FoldingRangeAdapter; },
/* harmony export */   HoverAdapter: function() { return /* binding */ HoverAdapter; },
/* harmony export */   ReferenceAdapter: function() { return /* binding */ ReferenceAdapter; },
/* harmony export */   RenameAdapter: function() { return /* binding */ RenameAdapter; },
/* harmony export */   SelectionRangeAdapter: function() { return /* binding */ SelectionRangeAdapter; },
/* harmony export */   WorkerManager: function() { return /* binding */ WorkerManager; },
/* harmony export */   fromPosition: function() { return /* binding */ fromPosition; },
/* harmony export */   fromRange: function() { return /* binding */ fromRange; },
/* harmony export */   setupMode: function() { return /* binding */ setupMode; },
/* harmony export */   toRange: function() { return /* binding */ toRange; },
/* harmony export */   toTextEdit: function() { return /* binding */ toTextEdit; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/language/json/workerManager.ts
var STOP_WHEN_IDLE_FOR = 2 * 60 * 1e3;
var WorkerManager = class {
  _defaults;
  _idleCheckInterval;
  _lastUsedTime;
  _configChangeListener;
  _worker;
  _client;
  constructor(defaults) {
    this._defaults = defaults;
    this._worker = null;
    this._client = null;
    this._idleCheckInterval = window.setInterval(() => this._checkIfIdle(), 30 * 1e3);
    this._lastUsedTime = 0;
    this._configChangeListener = this._defaults.onDidChange(() => this._stopWorker());
  }
  _stopWorker() {
    if (this._worker) {
      this._worker.dispose();
      this._worker = null;
    }
    this._client = null;
  }
  dispose() {
    clearInterval(this._idleCheckInterval);
    this._configChangeListener.dispose();
    this._stopWorker();
  }
  _checkIfIdle() {
    if (!this._worker) {
      return;
    }
    let timePassedSinceLastUsed = Date.now() - this._lastUsedTime;
    if (timePassedSinceLastUsed > STOP_WHEN_IDLE_FOR) {
      this._stopWorker();
    }
  }
  _getClient() {
    this._lastUsedTime = Date.now();
    if (!this._client) {
      this._worker = monaco_editor_core_exports.editor.createWebWorker({
        moduleId: "vs/language/json/jsonWorker",
        label: this._defaults.languageId,
        createData: {
          languageSettings: this._defaults.diagnosticsOptions,
          languageId: this._defaults.languageId,
          enableSchemaRequest: this._defaults.diagnosticsOptions.enableSchemaRequest
        }
      });
      this._client = this._worker.getProxy();
    }
    return this._client;
  }
  getLanguageServiceWorker(...resources) {
    let _client;
    return this._getClient().then((client) => {
      _client = client;
    }).then((_) => {
      if (this._worker) {
        return this._worker.withSyncedResources(resources);
      }
    }).then((_) => _client);
  }
};

// node_modules/vscode-languageserver-types/lib/esm/main.js
var integer;
(function(integer2) {
  integer2.MIN_VALUE = -2147483648;
  integer2.MAX_VALUE = 2147483647;
})(integer || (integer = {}));
var uinteger;
(function(uinteger2) {
  uinteger2.MIN_VALUE = 0;
  uinteger2.MAX_VALUE = 2147483647;
})(uinteger || (uinteger = {}));
var Position;
(function(Position3) {
  function create(line, character) {
    if (line === Number.MAX_VALUE) {
      line = uinteger.MAX_VALUE;
    }
    if (character === Number.MAX_VALUE) {
      character = uinteger.MAX_VALUE;
    }
    return { line, character };
  }
  Position3.create = create;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Is.uinteger(candidate.line) && Is.uinteger(candidate.character);
  }
  Position3.is = is;
})(Position || (Position = {}));
var Range;
(function(Range3) {
  function create(one, two, three, four) {
    if (Is.uinteger(one) && Is.uinteger(two) && Is.uinteger(three) && Is.uinteger(four)) {
      return { start: Position.create(one, two), end: Position.create(three, four) };
    } else if (Position.is(one) && Position.is(two)) {
      return { start: one, end: two };
    } else {
      throw new Error("Range#create called with invalid arguments[" + one + ", " + two + ", " + three + ", " + four + "]");
    }
  }
  Range3.create = create;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Position.is(candidate.start) && Position.is(candidate.end);
  }
  Range3.is = is;
})(Range || (Range = {}));
var Location;
(function(Location2) {
  function create(uri, range) {
    return { uri, range };
  }
  Location2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.string(candidate.uri) || Is.undefined(candidate.uri));
  }
  Location2.is = is;
})(Location || (Location = {}));
var LocationLink;
(function(LocationLink2) {
  function create(targetUri, targetRange, targetSelectionRange, originSelectionRange) {
    return { targetUri, targetRange, targetSelectionRange, originSelectionRange };
  }
  LocationLink2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.targetRange) && Is.string(candidate.targetUri) && (Range.is(candidate.targetSelectionRange) || Is.undefined(candidate.targetSelectionRange)) && (Range.is(candidate.originSelectionRange) || Is.undefined(candidate.originSelectionRange));
  }
  LocationLink2.is = is;
})(LocationLink || (LocationLink = {}));
var Color;
(function(Color2) {
  function create(red, green, blue, alpha) {
    return {
      red,
      green,
      blue,
      alpha
    };
  }
  Color2.create = create;
  function is(value) {
    var candidate = value;
    return Is.numberRange(candidate.red, 0, 1) && Is.numberRange(candidate.green, 0, 1) && Is.numberRange(candidate.blue, 0, 1) && Is.numberRange(candidate.alpha, 0, 1);
  }
  Color2.is = is;
})(Color || (Color = {}));
var ColorInformation;
(function(ColorInformation2) {
  function create(range, color) {
    return {
      range,
      color
    };
  }
  ColorInformation2.create = create;
  function is(value) {
    var candidate = value;
    return Range.is(candidate.range) && Color.is(candidate.color);
  }
  ColorInformation2.is = is;
})(ColorInformation || (ColorInformation = {}));
var ColorPresentation;
(function(ColorPresentation2) {
  function create(label, textEdit, additionalTextEdits) {
    return {
      label,
      textEdit,
      additionalTextEdits
    };
  }
  ColorPresentation2.create = create;
  function is(value) {
    var candidate = value;
    return Is.string(candidate.label) && (Is.undefined(candidate.textEdit) || TextEdit.is(candidate)) && (Is.undefined(candidate.additionalTextEdits) || Is.typedArray(candidate.additionalTextEdits, TextEdit.is));
  }
  ColorPresentation2.is = is;
})(ColorPresentation || (ColorPresentation = {}));
var FoldingRangeKind;
(function(FoldingRangeKind2) {
  FoldingRangeKind2["Comment"] = "comment";
  FoldingRangeKind2["Imports"] = "imports";
  FoldingRangeKind2["Region"] = "region";
})(FoldingRangeKind || (FoldingRangeKind = {}));
var FoldingRange;
(function(FoldingRange2) {
  function create(startLine, endLine, startCharacter, endCharacter, kind) {
    var result = {
      startLine,
      endLine
    };
    if (Is.defined(startCharacter)) {
      result.startCharacter = startCharacter;
    }
    if (Is.defined(endCharacter)) {
      result.endCharacter = endCharacter;
    }
    if (Is.defined(kind)) {
      result.kind = kind;
    }
    return result;
  }
  FoldingRange2.create = create;
  function is(value) {
    var candidate = value;
    return Is.uinteger(candidate.startLine) && Is.uinteger(candidate.startLine) && (Is.undefined(candidate.startCharacter) || Is.uinteger(candidate.startCharacter)) && (Is.undefined(candidate.endCharacter) || Is.uinteger(candidate.endCharacter)) && (Is.undefined(candidate.kind) || Is.string(candidate.kind));
  }
  FoldingRange2.is = is;
})(FoldingRange || (FoldingRange = {}));
var DiagnosticRelatedInformation;
(function(DiagnosticRelatedInformation2) {
  function create(location, message) {
    return {
      location,
      message
    };
  }
  DiagnosticRelatedInformation2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Location.is(candidate.location) && Is.string(candidate.message);
  }
  DiagnosticRelatedInformation2.is = is;
})(DiagnosticRelatedInformation || (DiagnosticRelatedInformation = {}));
var DiagnosticSeverity;
(function(DiagnosticSeverity2) {
  DiagnosticSeverity2.Error = 1;
  DiagnosticSeverity2.Warning = 2;
  DiagnosticSeverity2.Information = 3;
  DiagnosticSeverity2.Hint = 4;
})(DiagnosticSeverity || (DiagnosticSeverity = {}));
var DiagnosticTag;
(function(DiagnosticTag2) {
  DiagnosticTag2.Unnecessary = 1;
  DiagnosticTag2.Deprecated = 2;
})(DiagnosticTag || (DiagnosticTag = {}));
var CodeDescription;
(function(CodeDescription2) {
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && candidate !== null && Is.string(candidate.href);
  }
  CodeDescription2.is = is;
})(CodeDescription || (CodeDescription = {}));
var Diagnostic;
(function(Diagnostic2) {
  function create(range, message, severity, code, source, relatedInformation) {
    var result = { range, message };
    if (Is.defined(severity)) {
      result.severity = severity;
    }
    if (Is.defined(code)) {
      result.code = code;
    }
    if (Is.defined(source)) {
      result.source = source;
    }
    if (Is.defined(relatedInformation)) {
      result.relatedInformation = relatedInformation;
    }
    return result;
  }
  Diagnostic2.create = create;
  function is(value) {
    var _a;
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && Is.string(candidate.message) && (Is.number(candidate.severity) || Is.undefined(candidate.severity)) && (Is.integer(candidate.code) || Is.string(candidate.code) || Is.undefined(candidate.code)) && (Is.undefined(candidate.codeDescription) || Is.string((_a = candidate.codeDescription) === null || _a === void 0 ? void 0 : _a.href)) && (Is.string(candidate.source) || Is.undefined(candidate.source)) && (Is.undefined(candidate.relatedInformation) || Is.typedArray(candidate.relatedInformation, DiagnosticRelatedInformation.is));
  }
  Diagnostic2.is = is;
})(Diagnostic || (Diagnostic = {}));
var Command;
(function(Command2) {
  function create(title, command) {
    var args = [];
    for (var _i = 2; _i < arguments.length; _i++) {
      args[_i - 2] = arguments[_i];
    }
    var result = { title, command };
    if (Is.defined(args) && args.length > 0) {
      result.arguments = args;
    }
    return result;
  }
  Command2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.title) && Is.string(candidate.command);
  }
  Command2.is = is;
})(Command || (Command = {}));
var TextEdit;
(function(TextEdit2) {
  function replace(range, newText) {
    return { range, newText };
  }
  TextEdit2.replace = replace;
  function insert(position, newText) {
    return { range: { start: position, end: position }, newText };
  }
  TextEdit2.insert = insert;
  function del(range) {
    return { range, newText: "" };
  }
  TextEdit2.del = del;
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(candidate) && Is.string(candidate.newText) && Range.is(candidate.range);
  }
  TextEdit2.is = is;
})(TextEdit || (TextEdit = {}));
var ChangeAnnotation;
(function(ChangeAnnotation2) {
  function create(label, needsConfirmation, description) {
    var result = { label };
    if (needsConfirmation !== void 0) {
      result.needsConfirmation = needsConfirmation;
    }
    if (description !== void 0) {
      result.description = description;
    }
    return result;
  }
  ChangeAnnotation2.create = create;
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && Is.objectLiteral(candidate) && Is.string(candidate.label) && (Is.boolean(candidate.needsConfirmation) || candidate.needsConfirmation === void 0) && (Is.string(candidate.description) || candidate.description === void 0);
  }
  ChangeAnnotation2.is = is;
})(ChangeAnnotation || (ChangeAnnotation = {}));
var ChangeAnnotationIdentifier;
(function(ChangeAnnotationIdentifier2) {
  function is(value) {
    var candidate = value;
    return typeof candidate === "string";
  }
  ChangeAnnotationIdentifier2.is = is;
})(ChangeAnnotationIdentifier || (ChangeAnnotationIdentifier = {}));
var AnnotatedTextEdit;
(function(AnnotatedTextEdit2) {
  function replace(range, newText, annotation) {
    return { range, newText, annotationId: annotation };
  }
  AnnotatedTextEdit2.replace = replace;
  function insert(position, newText, annotation) {
    return { range: { start: position, end: position }, newText, annotationId: annotation };
  }
  AnnotatedTextEdit2.insert = insert;
  function del(range, annotation) {
    return { range, newText: "", annotationId: annotation };
  }
  AnnotatedTextEdit2.del = del;
  function is(value) {
    var candidate = value;
    return TextEdit.is(candidate) && (ChangeAnnotation.is(candidate.annotationId) || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  AnnotatedTextEdit2.is = is;
})(AnnotatedTextEdit || (AnnotatedTextEdit = {}));
var TextDocumentEdit;
(function(TextDocumentEdit2) {
  function create(textDocument, edits) {
    return { textDocument, edits };
  }
  TextDocumentEdit2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && OptionalVersionedTextDocumentIdentifier.is(candidate.textDocument) && Array.isArray(candidate.edits);
  }
  TextDocumentEdit2.is = is;
})(TextDocumentEdit || (TextDocumentEdit = {}));
var CreateFile;
(function(CreateFile2) {
  function create(uri, options, annotation) {
    var result = {
      kind: "create",
      uri
    };
    if (options !== void 0 && (options.overwrite !== void 0 || options.ignoreIfExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  CreateFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "create" && Is.string(candidate.uri) && (candidate.options === void 0 || (candidate.options.overwrite === void 0 || Is.boolean(candidate.options.overwrite)) && (candidate.options.ignoreIfExists === void 0 || Is.boolean(candidate.options.ignoreIfExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  CreateFile2.is = is;
})(CreateFile || (CreateFile = {}));
var RenameFile;
(function(RenameFile2) {
  function create(oldUri, newUri, options, annotation) {
    var result = {
      kind: "rename",
      oldUri,
      newUri
    };
    if (options !== void 0 && (options.overwrite !== void 0 || options.ignoreIfExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  RenameFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "rename" && Is.string(candidate.oldUri) && Is.string(candidate.newUri) && (candidate.options === void 0 || (candidate.options.overwrite === void 0 || Is.boolean(candidate.options.overwrite)) && (candidate.options.ignoreIfExists === void 0 || Is.boolean(candidate.options.ignoreIfExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  RenameFile2.is = is;
})(RenameFile || (RenameFile = {}));
var DeleteFile;
(function(DeleteFile2) {
  function create(uri, options, annotation) {
    var result = {
      kind: "delete",
      uri
    };
    if (options !== void 0 && (options.recursive !== void 0 || options.ignoreIfNotExists !== void 0)) {
      result.options = options;
    }
    if (annotation !== void 0) {
      result.annotationId = annotation;
    }
    return result;
  }
  DeleteFile2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && candidate.kind === "delete" && Is.string(candidate.uri) && (candidate.options === void 0 || (candidate.options.recursive === void 0 || Is.boolean(candidate.options.recursive)) && (candidate.options.ignoreIfNotExists === void 0 || Is.boolean(candidate.options.ignoreIfNotExists))) && (candidate.annotationId === void 0 || ChangeAnnotationIdentifier.is(candidate.annotationId));
  }
  DeleteFile2.is = is;
})(DeleteFile || (DeleteFile = {}));
var WorkspaceEdit;
(function(WorkspaceEdit2) {
  function is(value) {
    var candidate = value;
    return candidate && (candidate.changes !== void 0 || candidate.documentChanges !== void 0) && (candidate.documentChanges === void 0 || candidate.documentChanges.every(function(change) {
      if (Is.string(change.kind)) {
        return CreateFile.is(change) || RenameFile.is(change) || DeleteFile.is(change);
      } else {
        return TextDocumentEdit.is(change);
      }
    }));
  }
  WorkspaceEdit2.is = is;
})(WorkspaceEdit || (WorkspaceEdit = {}));
var TextEditChangeImpl = function() {
  function TextEditChangeImpl2(edits, changeAnnotations) {
    this.edits = edits;
    this.changeAnnotations = changeAnnotations;
  }
  TextEditChangeImpl2.prototype.insert = function(position, newText, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.insert(position, newText);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.insert(position, newText, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.insert(position, newText, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.replace = function(range, newText, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.replace(range, newText);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.replace(range, newText, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.replace(range, newText, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.delete = function(range, annotation) {
    var edit;
    var id;
    if (annotation === void 0) {
      edit = TextEdit.del(range);
    } else if (ChangeAnnotationIdentifier.is(annotation)) {
      id = annotation;
      edit = AnnotatedTextEdit.del(range, annotation);
    } else {
      this.assertChangeAnnotations(this.changeAnnotations);
      id = this.changeAnnotations.manage(annotation);
      edit = AnnotatedTextEdit.del(range, id);
    }
    this.edits.push(edit);
    if (id !== void 0) {
      return id;
    }
  };
  TextEditChangeImpl2.prototype.add = function(edit) {
    this.edits.push(edit);
  };
  TextEditChangeImpl2.prototype.all = function() {
    return this.edits;
  };
  TextEditChangeImpl2.prototype.clear = function() {
    this.edits.splice(0, this.edits.length);
  };
  TextEditChangeImpl2.prototype.assertChangeAnnotations = function(value) {
    if (value === void 0) {
      throw new Error("Text edit change is not configured to manage change annotations.");
    }
  };
  return TextEditChangeImpl2;
}();
var ChangeAnnotations = function() {
  function ChangeAnnotations2(annotations) {
    this._annotations = annotations === void 0 ? /* @__PURE__ */ Object.create(null) : annotations;
    this._counter = 0;
    this._size = 0;
  }
  ChangeAnnotations2.prototype.all = function() {
    return this._annotations;
  };
  Object.defineProperty(ChangeAnnotations2.prototype, "size", {
    get: function() {
      return this._size;
    },
    enumerable: false,
    configurable: true
  });
  ChangeAnnotations2.prototype.manage = function(idOrAnnotation, annotation) {
    var id;
    if (ChangeAnnotationIdentifier.is(idOrAnnotation)) {
      id = idOrAnnotation;
    } else {
      id = this.nextId();
      annotation = idOrAnnotation;
    }
    if (this._annotations[id] !== void 0) {
      throw new Error("Id " + id + " is already in use.");
    }
    if (annotation === void 0) {
      throw new Error("No annotation provided for id " + id);
    }
    this._annotations[id] = annotation;
    this._size++;
    return id;
  };
  ChangeAnnotations2.prototype.nextId = function() {
    this._counter++;
    return this._counter.toString();
  };
  return ChangeAnnotations2;
}();
var WorkspaceChange = function() {
  function WorkspaceChange2(workspaceEdit) {
    var _this = this;
    this._textEditChanges = /* @__PURE__ */ Object.create(null);
    if (workspaceEdit !== void 0) {
      this._workspaceEdit = workspaceEdit;
      if (workspaceEdit.documentChanges) {
        this._changeAnnotations = new ChangeAnnotations(workspaceEdit.changeAnnotations);
        workspaceEdit.changeAnnotations = this._changeAnnotations.all();
        workspaceEdit.documentChanges.forEach(function(change) {
          if (TextDocumentEdit.is(change)) {
            var textEditChange = new TextEditChangeImpl(change.edits, _this._changeAnnotations);
            _this._textEditChanges[change.textDocument.uri] = textEditChange;
          }
        });
      } else if (workspaceEdit.changes) {
        Object.keys(workspaceEdit.changes).forEach(function(key) {
          var textEditChange = new TextEditChangeImpl(workspaceEdit.changes[key]);
          _this._textEditChanges[key] = textEditChange;
        });
      }
    } else {
      this._workspaceEdit = {};
    }
  }
  Object.defineProperty(WorkspaceChange2.prototype, "edit", {
    get: function() {
      this.initDocumentChanges();
      if (this._changeAnnotations !== void 0) {
        if (this._changeAnnotations.size === 0) {
          this._workspaceEdit.changeAnnotations = void 0;
        } else {
          this._workspaceEdit.changeAnnotations = this._changeAnnotations.all();
        }
      }
      return this._workspaceEdit;
    },
    enumerable: false,
    configurable: true
  });
  WorkspaceChange2.prototype.getTextEditChange = function(key) {
    if (OptionalVersionedTextDocumentIdentifier.is(key)) {
      this.initDocumentChanges();
      if (this._workspaceEdit.documentChanges === void 0) {
        throw new Error("Workspace edit is not configured for document changes.");
      }
      var textDocument = { uri: key.uri, version: key.version };
      var result = this._textEditChanges[textDocument.uri];
      if (!result) {
        var edits = [];
        var textDocumentEdit = {
          textDocument,
          edits
        };
        this._workspaceEdit.documentChanges.push(textDocumentEdit);
        result = new TextEditChangeImpl(edits, this._changeAnnotations);
        this._textEditChanges[textDocument.uri] = result;
      }
      return result;
    } else {
      this.initChanges();
      if (this._workspaceEdit.changes === void 0) {
        throw new Error("Workspace edit is not configured for normal text edit changes.");
      }
      var result = this._textEditChanges[key];
      if (!result) {
        var edits = [];
        this._workspaceEdit.changes[key] = edits;
        result = new TextEditChangeImpl(edits);
        this._textEditChanges[key] = result;
      }
      return result;
    }
  };
  WorkspaceChange2.prototype.initDocumentChanges = function() {
    if (this._workspaceEdit.documentChanges === void 0 && this._workspaceEdit.changes === void 0) {
      this._changeAnnotations = new ChangeAnnotations();
      this._workspaceEdit.documentChanges = [];
      this._workspaceEdit.changeAnnotations = this._changeAnnotations.all();
    }
  };
  WorkspaceChange2.prototype.initChanges = function() {
    if (this._workspaceEdit.documentChanges === void 0 && this._workspaceEdit.changes === void 0) {
      this._workspaceEdit.changes = /* @__PURE__ */ Object.create(null);
    }
  };
  WorkspaceChange2.prototype.createFile = function(uri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = CreateFile.create(uri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = CreateFile.create(uri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  WorkspaceChange2.prototype.renameFile = function(oldUri, newUri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = RenameFile.create(oldUri, newUri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = RenameFile.create(oldUri, newUri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  WorkspaceChange2.prototype.deleteFile = function(uri, optionsOrAnnotation, options) {
    this.initDocumentChanges();
    if (this._workspaceEdit.documentChanges === void 0) {
      throw new Error("Workspace edit is not configured for document changes.");
    }
    var annotation;
    if (ChangeAnnotation.is(optionsOrAnnotation) || ChangeAnnotationIdentifier.is(optionsOrAnnotation)) {
      annotation = optionsOrAnnotation;
    } else {
      options = optionsOrAnnotation;
    }
    var operation;
    var id;
    if (annotation === void 0) {
      operation = DeleteFile.create(uri, options);
    } else {
      id = ChangeAnnotationIdentifier.is(annotation) ? annotation : this._changeAnnotations.manage(annotation);
      operation = DeleteFile.create(uri, options, id);
    }
    this._workspaceEdit.documentChanges.push(operation);
    if (id !== void 0) {
      return id;
    }
  };
  return WorkspaceChange2;
}();
var TextDocumentIdentifier;
(function(TextDocumentIdentifier2) {
  function create(uri) {
    return { uri };
  }
  TextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri);
  }
  TextDocumentIdentifier2.is = is;
})(TextDocumentIdentifier || (TextDocumentIdentifier = {}));
var VersionedTextDocumentIdentifier;
(function(VersionedTextDocumentIdentifier2) {
  function create(uri, version) {
    return { uri, version };
  }
  VersionedTextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && Is.integer(candidate.version);
  }
  VersionedTextDocumentIdentifier2.is = is;
})(VersionedTextDocumentIdentifier || (VersionedTextDocumentIdentifier = {}));
var OptionalVersionedTextDocumentIdentifier;
(function(OptionalVersionedTextDocumentIdentifier2) {
  function create(uri, version) {
    return { uri, version };
  }
  OptionalVersionedTextDocumentIdentifier2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && (candidate.version === null || Is.integer(candidate.version));
  }
  OptionalVersionedTextDocumentIdentifier2.is = is;
})(OptionalVersionedTextDocumentIdentifier || (OptionalVersionedTextDocumentIdentifier = {}));
var TextDocumentItem;
(function(TextDocumentItem2) {
  function create(uri, languageId, version, text) {
    return { uri, languageId, version, text };
  }
  TextDocumentItem2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && Is.string(candidate.languageId) && Is.integer(candidate.version) && Is.string(candidate.text);
  }
  TextDocumentItem2.is = is;
})(TextDocumentItem || (TextDocumentItem = {}));
var MarkupKind;
(function(MarkupKind2) {
  MarkupKind2.PlainText = "plaintext";
  MarkupKind2.Markdown = "markdown";
})(MarkupKind || (MarkupKind = {}));
(function(MarkupKind2) {
  function is(value) {
    var candidate = value;
    return candidate === MarkupKind2.PlainText || candidate === MarkupKind2.Markdown;
  }
  MarkupKind2.is = is;
})(MarkupKind || (MarkupKind = {}));
var MarkupContent;
(function(MarkupContent2) {
  function is(value) {
    var candidate = value;
    return Is.objectLiteral(value) && MarkupKind.is(candidate.kind) && Is.string(candidate.value);
  }
  MarkupContent2.is = is;
})(MarkupContent || (MarkupContent = {}));
var CompletionItemKind;
(function(CompletionItemKind2) {
  CompletionItemKind2.Text = 1;
  CompletionItemKind2.Method = 2;
  CompletionItemKind2.Function = 3;
  CompletionItemKind2.Constructor = 4;
  CompletionItemKind2.Field = 5;
  CompletionItemKind2.Variable = 6;
  CompletionItemKind2.Class = 7;
  CompletionItemKind2.Interface = 8;
  CompletionItemKind2.Module = 9;
  CompletionItemKind2.Property = 10;
  CompletionItemKind2.Unit = 11;
  CompletionItemKind2.Value = 12;
  CompletionItemKind2.Enum = 13;
  CompletionItemKind2.Keyword = 14;
  CompletionItemKind2.Snippet = 15;
  CompletionItemKind2.Color = 16;
  CompletionItemKind2.File = 17;
  CompletionItemKind2.Reference = 18;
  CompletionItemKind2.Folder = 19;
  CompletionItemKind2.EnumMember = 20;
  CompletionItemKind2.Constant = 21;
  CompletionItemKind2.Struct = 22;
  CompletionItemKind2.Event = 23;
  CompletionItemKind2.Operator = 24;
  CompletionItemKind2.TypeParameter = 25;
})(CompletionItemKind || (CompletionItemKind = {}));
var InsertTextFormat;
(function(InsertTextFormat2) {
  InsertTextFormat2.PlainText = 1;
  InsertTextFormat2.Snippet = 2;
})(InsertTextFormat || (InsertTextFormat = {}));
var CompletionItemTag;
(function(CompletionItemTag2) {
  CompletionItemTag2.Deprecated = 1;
})(CompletionItemTag || (CompletionItemTag = {}));
var InsertReplaceEdit;
(function(InsertReplaceEdit2) {
  function create(newText, insert, replace) {
    return { newText, insert, replace };
  }
  InsertReplaceEdit2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.newText) && Range.is(candidate.insert) && Range.is(candidate.replace);
  }
  InsertReplaceEdit2.is = is;
})(InsertReplaceEdit || (InsertReplaceEdit = {}));
var InsertTextMode;
(function(InsertTextMode2) {
  InsertTextMode2.asIs = 1;
  InsertTextMode2.adjustIndentation = 2;
})(InsertTextMode || (InsertTextMode = {}));
var CompletionItem;
(function(CompletionItem2) {
  function create(label) {
    return { label };
  }
  CompletionItem2.create = create;
})(CompletionItem || (CompletionItem = {}));
var CompletionList;
(function(CompletionList2) {
  function create(items, isIncomplete) {
    return { items: items ? items : [], isIncomplete: !!isIncomplete };
  }
  CompletionList2.create = create;
})(CompletionList || (CompletionList = {}));
var MarkedString;
(function(MarkedString2) {
  function fromPlainText(plainText) {
    return plainText.replace(/[\\`*_{}[\]()#+\-.!]/g, "\\$&");
  }
  MarkedString2.fromPlainText = fromPlainText;
  function is(value) {
    var candidate = value;
    return Is.string(candidate) || Is.objectLiteral(candidate) && Is.string(candidate.language) && Is.string(candidate.value);
  }
  MarkedString2.is = is;
})(MarkedString || (MarkedString = {}));
var Hover;
(function(Hover2) {
  function is(value) {
    var candidate = value;
    return !!candidate && Is.objectLiteral(candidate) && (MarkupContent.is(candidate.contents) || MarkedString.is(candidate.contents) || Is.typedArray(candidate.contents, MarkedString.is)) && (value.range === void 0 || Range.is(value.range));
  }
  Hover2.is = is;
})(Hover || (Hover = {}));
var ParameterInformation;
(function(ParameterInformation2) {
  function create(label, documentation) {
    return documentation ? { label, documentation } : { label };
  }
  ParameterInformation2.create = create;
})(ParameterInformation || (ParameterInformation = {}));
var SignatureInformation;
(function(SignatureInformation2) {
  function create(label, documentation) {
    var parameters = [];
    for (var _i = 2; _i < arguments.length; _i++) {
      parameters[_i - 2] = arguments[_i];
    }
    var result = { label };
    if (Is.defined(documentation)) {
      result.documentation = documentation;
    }
    if (Is.defined(parameters)) {
      result.parameters = parameters;
    } else {
      result.parameters = [];
    }
    return result;
  }
  SignatureInformation2.create = create;
})(SignatureInformation || (SignatureInformation = {}));
var DocumentHighlightKind;
(function(DocumentHighlightKind2) {
  DocumentHighlightKind2.Text = 1;
  DocumentHighlightKind2.Read = 2;
  DocumentHighlightKind2.Write = 3;
})(DocumentHighlightKind || (DocumentHighlightKind = {}));
var DocumentHighlight;
(function(DocumentHighlight2) {
  function create(range, kind) {
    var result = { range };
    if (Is.number(kind)) {
      result.kind = kind;
    }
    return result;
  }
  DocumentHighlight2.create = create;
})(DocumentHighlight || (DocumentHighlight = {}));
var SymbolKind;
(function(SymbolKind2) {
  SymbolKind2.File = 1;
  SymbolKind2.Module = 2;
  SymbolKind2.Namespace = 3;
  SymbolKind2.Package = 4;
  SymbolKind2.Class = 5;
  SymbolKind2.Method = 6;
  SymbolKind2.Property = 7;
  SymbolKind2.Field = 8;
  SymbolKind2.Constructor = 9;
  SymbolKind2.Enum = 10;
  SymbolKind2.Interface = 11;
  SymbolKind2.Function = 12;
  SymbolKind2.Variable = 13;
  SymbolKind2.Constant = 14;
  SymbolKind2.String = 15;
  SymbolKind2.Number = 16;
  SymbolKind2.Boolean = 17;
  SymbolKind2.Array = 18;
  SymbolKind2.Object = 19;
  SymbolKind2.Key = 20;
  SymbolKind2.Null = 21;
  SymbolKind2.EnumMember = 22;
  SymbolKind2.Struct = 23;
  SymbolKind2.Event = 24;
  SymbolKind2.Operator = 25;
  SymbolKind2.TypeParameter = 26;
})(SymbolKind || (SymbolKind = {}));
var SymbolTag;
(function(SymbolTag2) {
  SymbolTag2.Deprecated = 1;
})(SymbolTag || (SymbolTag = {}));
var SymbolInformation;
(function(SymbolInformation2) {
  function create(name, kind, range, uri, containerName) {
    var result = {
      name,
      kind,
      location: { uri, range }
    };
    if (containerName) {
      result.containerName = containerName;
    }
    return result;
  }
  SymbolInformation2.create = create;
})(SymbolInformation || (SymbolInformation = {}));
var DocumentSymbol;
(function(DocumentSymbol2) {
  function create(name, detail, kind, range, selectionRange, children) {
    var result = {
      name,
      detail,
      kind,
      range,
      selectionRange
    };
    if (children !== void 0) {
      result.children = children;
    }
    return result;
  }
  DocumentSymbol2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.name) && Is.number(candidate.kind) && Range.is(candidate.range) && Range.is(candidate.selectionRange) && (candidate.detail === void 0 || Is.string(candidate.detail)) && (candidate.deprecated === void 0 || Is.boolean(candidate.deprecated)) && (candidate.children === void 0 || Array.isArray(candidate.children)) && (candidate.tags === void 0 || Array.isArray(candidate.tags));
  }
  DocumentSymbol2.is = is;
})(DocumentSymbol || (DocumentSymbol = {}));
var CodeActionKind;
(function(CodeActionKind2) {
  CodeActionKind2.Empty = "";
  CodeActionKind2.QuickFix = "quickfix";
  CodeActionKind2.Refactor = "refactor";
  CodeActionKind2.RefactorExtract = "refactor.extract";
  CodeActionKind2.RefactorInline = "refactor.inline";
  CodeActionKind2.RefactorRewrite = "refactor.rewrite";
  CodeActionKind2.Source = "source";
  CodeActionKind2.SourceOrganizeImports = "source.organizeImports";
  CodeActionKind2.SourceFixAll = "source.fixAll";
})(CodeActionKind || (CodeActionKind = {}));
var CodeActionContext;
(function(CodeActionContext2) {
  function create(diagnostics, only) {
    var result = { diagnostics };
    if (only !== void 0 && only !== null) {
      result.only = only;
    }
    return result;
  }
  CodeActionContext2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.typedArray(candidate.diagnostics, Diagnostic.is) && (candidate.only === void 0 || Is.typedArray(candidate.only, Is.string));
  }
  CodeActionContext2.is = is;
})(CodeActionContext || (CodeActionContext = {}));
var CodeAction;
(function(CodeAction2) {
  function create(title, kindOrCommandOrEdit, kind) {
    var result = { title };
    var checkKind = true;
    if (typeof kindOrCommandOrEdit === "string") {
      checkKind = false;
      result.kind = kindOrCommandOrEdit;
    } else if (Command.is(kindOrCommandOrEdit)) {
      result.command = kindOrCommandOrEdit;
    } else {
      result.edit = kindOrCommandOrEdit;
    }
    if (checkKind && kind !== void 0) {
      result.kind = kind;
    }
    return result;
  }
  CodeAction2.create = create;
  function is(value) {
    var candidate = value;
    return candidate && Is.string(candidate.title) && (candidate.diagnostics === void 0 || Is.typedArray(candidate.diagnostics, Diagnostic.is)) && (candidate.kind === void 0 || Is.string(candidate.kind)) && (candidate.edit !== void 0 || candidate.command !== void 0) && (candidate.command === void 0 || Command.is(candidate.command)) && (candidate.isPreferred === void 0 || Is.boolean(candidate.isPreferred)) && (candidate.edit === void 0 || WorkspaceEdit.is(candidate.edit));
  }
  CodeAction2.is = is;
})(CodeAction || (CodeAction = {}));
var CodeLens;
(function(CodeLens2) {
  function create(range, data) {
    var result = { range };
    if (Is.defined(data)) {
      result.data = data;
    }
    return result;
  }
  CodeLens2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.undefined(candidate.command) || Command.is(candidate.command));
  }
  CodeLens2.is = is;
})(CodeLens || (CodeLens = {}));
var FormattingOptions;
(function(FormattingOptions2) {
  function create(tabSize, insertSpaces) {
    return { tabSize, insertSpaces };
  }
  FormattingOptions2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.uinteger(candidate.tabSize) && Is.boolean(candidate.insertSpaces);
  }
  FormattingOptions2.is = is;
})(FormattingOptions || (FormattingOptions = {}));
var DocumentLink;
(function(DocumentLink2) {
  function create(range, target, data) {
    return { range, target, data };
  }
  DocumentLink2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Range.is(candidate.range) && (Is.undefined(candidate.target) || Is.string(candidate.target));
  }
  DocumentLink2.is = is;
})(DocumentLink || (DocumentLink = {}));
var SelectionRange;
(function(SelectionRange2) {
  function create(range, parent) {
    return { range, parent };
  }
  SelectionRange2.create = create;
  function is(value) {
    var candidate = value;
    return candidate !== void 0 && Range.is(candidate.range) && (candidate.parent === void 0 || SelectionRange2.is(candidate.parent));
  }
  SelectionRange2.is = is;
})(SelectionRange || (SelectionRange = {}));
var TextDocument;
(function(TextDocument2) {
  function create(uri, languageId, version, content) {
    return new FullTextDocument(uri, languageId, version, content);
  }
  TextDocument2.create = create;
  function is(value) {
    var candidate = value;
    return Is.defined(candidate) && Is.string(candidate.uri) && (Is.undefined(candidate.languageId) || Is.string(candidate.languageId)) && Is.uinteger(candidate.lineCount) && Is.func(candidate.getText) && Is.func(candidate.positionAt) && Is.func(candidate.offsetAt) ? true : false;
  }
  TextDocument2.is = is;
  function applyEdits(document, edits) {
    var text = document.getText();
    var sortedEdits = mergeSort(edits, function(a, b) {
      var diff = a.range.start.line - b.range.start.line;
      if (diff === 0) {
        return a.range.start.character - b.range.start.character;
      }
      return diff;
    });
    var lastModifiedOffset = text.length;
    for (var i = sortedEdits.length - 1; i >= 0; i--) {
      var e = sortedEdits[i];
      var startOffset = document.offsetAt(e.range.start);
      var endOffset = document.offsetAt(e.range.end);
      if (endOffset <= lastModifiedOffset) {
        text = text.substring(0, startOffset) + e.newText + text.substring(endOffset, text.length);
      } else {
        throw new Error("Overlapping edit");
      }
      lastModifiedOffset = startOffset;
    }
    return text;
  }
  TextDocument2.applyEdits = applyEdits;
  function mergeSort(data, compare) {
    if (data.length <= 1) {
      return data;
    }
    var p = data.length / 2 | 0;
    var left = data.slice(0, p);
    var right = data.slice(p);
    mergeSort(left, compare);
    mergeSort(right, compare);
    var leftIdx = 0;
    var rightIdx = 0;
    var i = 0;
    while (leftIdx < left.length && rightIdx < right.length) {
      var ret = compare(left[leftIdx], right[rightIdx]);
      if (ret <= 0) {
        data[i++] = left[leftIdx++];
      } else {
        data[i++] = right[rightIdx++];
      }
    }
    while (leftIdx < left.length) {
      data[i++] = left[leftIdx++];
    }
    while (rightIdx < right.length) {
      data[i++] = right[rightIdx++];
    }
    return data;
  }
})(TextDocument || (TextDocument = {}));
var FullTextDocument = function() {
  function FullTextDocument2(uri, languageId, version, content) {
    this._uri = uri;
    this._languageId = languageId;
    this._version = version;
    this._content = content;
    this._lineOffsets = void 0;
  }
  Object.defineProperty(FullTextDocument2.prototype, "uri", {
    get: function() {
      return this._uri;
    },
    enumerable: false,
    configurable: true
  });
  Object.defineProperty(FullTextDocument2.prototype, "languageId", {
    get: function() {
      return this._languageId;
    },
    enumerable: false,
    configurable: true
  });
  Object.defineProperty(FullTextDocument2.prototype, "version", {
    get: function() {
      return this._version;
    },
    enumerable: false,
    configurable: true
  });
  FullTextDocument2.prototype.getText = function(range) {
    if (range) {
      var start = this.offsetAt(range.start);
      var end = this.offsetAt(range.end);
      return this._content.substring(start, end);
    }
    return this._content;
  };
  FullTextDocument2.prototype.update = function(event, version) {
    this._content = event.text;
    this._version = version;
    this._lineOffsets = void 0;
  };
  FullTextDocument2.prototype.getLineOffsets = function() {
    if (this._lineOffsets === void 0) {
      var lineOffsets = [];
      var text = this._content;
      var isLineStart = true;
      for (var i = 0; i < text.length; i++) {
        if (isLineStart) {
          lineOffsets.push(i);
          isLineStart = false;
        }
        var ch = text.charAt(i);
        isLineStart = ch === "\r" || ch === "\n";
        if (ch === "\r" && i + 1 < text.length && text.charAt(i + 1) === "\n") {
          i++;
        }
      }
      if (isLineStart && text.length > 0) {
        lineOffsets.push(text.length);
      }
      this._lineOffsets = lineOffsets;
    }
    return this._lineOffsets;
  };
  FullTextDocument2.prototype.positionAt = function(offset) {
    offset = Math.max(Math.min(offset, this._content.length), 0);
    var lineOffsets = this.getLineOffsets();
    var low = 0, high = lineOffsets.length;
    if (high === 0) {
      return Position.create(0, offset);
    }
    while (low < high) {
      var mid = Math.floor((low + high) / 2);
      if (lineOffsets[mid] > offset) {
        high = mid;
      } else {
        low = mid + 1;
      }
    }
    var line = low - 1;
    return Position.create(line, offset - lineOffsets[line]);
  };
  FullTextDocument2.prototype.offsetAt = function(position) {
    var lineOffsets = this.getLineOffsets();
    if (position.line >= lineOffsets.length) {
      return this._content.length;
    } else if (position.line < 0) {
      return 0;
    }
    var lineOffset = lineOffsets[position.line];
    var nextLineOffset = position.line + 1 < lineOffsets.length ? lineOffsets[position.line + 1] : this._content.length;
    return Math.max(Math.min(lineOffset + position.character, nextLineOffset), lineOffset);
  };
  Object.defineProperty(FullTextDocument2.prototype, "lineCount", {
    get: function() {
      return this.getLineOffsets().length;
    },
    enumerable: false,
    configurable: true
  });
  return FullTextDocument2;
}();
var Is;
(function(Is2) {
  var toString = Object.prototype.toString;
  function defined(value) {
    return typeof value !== "undefined";
  }
  Is2.defined = defined;
  function undefined2(value) {
    return typeof value === "undefined";
  }
  Is2.undefined = undefined2;
  function boolean(value) {
    return value === true || value === false;
  }
  Is2.boolean = boolean;
  function string(value) {
    return toString.call(value) === "[object String]";
  }
  Is2.string = string;
  function number(value) {
    return toString.call(value) === "[object Number]";
  }
  Is2.number = number;
  function numberRange(value, min, max) {
    return toString.call(value) === "[object Number]" && min <= value && value <= max;
  }
  Is2.numberRange = numberRange;
  function integer2(value) {
    return toString.call(value) === "[object Number]" && -2147483648 <= value && value <= 2147483647;
  }
  Is2.integer = integer2;
  function uinteger2(value) {
    return toString.call(value) === "[object Number]" && 0 <= value && value <= 2147483647;
  }
  Is2.uinteger = uinteger2;
  function func(value) {
    return toString.call(value) === "[object Function]";
  }
  Is2.func = func;
  function objectLiteral(value) {
    return value !== null && typeof value === "object";
  }
  Is2.objectLiteral = objectLiteral;
  function typedArray(value, check) {
    return Array.isArray(value) && value.every(check);
  }
  Is2.typedArray = typedArray;
})(Is || (Is = {}));

// src/language/common/lspLanguageFeatures.ts
var DiagnosticsAdapter = class {
  constructor(_languageId, _worker, configChangeEvent) {
    this._languageId = _languageId;
    this._worker = _worker;
    const onModelAdd = (model) => {
      let modeId = model.getLanguageId();
      if (modeId !== this._languageId) {
        return;
      }
      let handle;
      this._listener[model.uri.toString()] = model.onDidChangeContent(() => {
        window.clearTimeout(handle);
        handle = window.setTimeout(() => this._doValidate(model.uri, modeId), 500);
      });
      this._doValidate(model.uri, modeId);
    };
    const onModelRemoved = (model) => {
      monaco_editor_core_exports.editor.setModelMarkers(model, this._languageId, []);
      let uriStr = model.uri.toString();
      let listener = this._listener[uriStr];
      if (listener) {
        listener.dispose();
        delete this._listener[uriStr];
      }
    };
    this._disposables.push(monaco_editor_core_exports.editor.onDidCreateModel(onModelAdd));
    this._disposables.push(monaco_editor_core_exports.editor.onWillDisposeModel(onModelRemoved));
    this._disposables.push(monaco_editor_core_exports.editor.onDidChangeModelLanguage((event) => {
      onModelRemoved(event.model);
      onModelAdd(event.model);
    }));
    this._disposables.push(configChangeEvent((_) => {
      monaco_editor_core_exports.editor.getModels().forEach((model) => {
        if (model.getLanguageId() === this._languageId) {
          onModelRemoved(model);
          onModelAdd(model);
        }
      });
    }));
    this._disposables.push({
      dispose: () => {
        monaco_editor_core_exports.editor.getModels().forEach(onModelRemoved);
        for (let key in this._listener) {
          this._listener[key].dispose();
        }
      }
    });
    monaco_editor_core_exports.editor.getModels().forEach(onModelAdd);
  }
  _disposables = [];
  _listener = /* @__PURE__ */ Object.create(null);
  dispose() {
    this._disposables.forEach((d) => d && d.dispose());
    this._disposables.length = 0;
  }
  _doValidate(resource, languageId) {
    this._worker(resource).then((worker) => {
      return worker.doValidation(resource.toString());
    }).then((diagnostics) => {
      const markers = diagnostics.map((d) => toDiagnostics(resource, d));
      let model = monaco_editor_core_exports.editor.getModel(resource);
      if (model && model.getLanguageId() === languageId) {
        monaco_editor_core_exports.editor.setModelMarkers(model, languageId, markers);
      }
    }).then(void 0, (err) => {
      console.error(err);
    });
  }
};
function toSeverity(lsSeverity) {
  switch (lsSeverity) {
    case DiagnosticSeverity.Error:
      return monaco_editor_core_exports.MarkerSeverity.Error;
    case DiagnosticSeverity.Warning:
      return monaco_editor_core_exports.MarkerSeverity.Warning;
    case DiagnosticSeverity.Information:
      return monaco_editor_core_exports.MarkerSeverity.Info;
    case DiagnosticSeverity.Hint:
      return monaco_editor_core_exports.MarkerSeverity.Hint;
    default:
      return monaco_editor_core_exports.MarkerSeverity.Info;
  }
}
function toDiagnostics(resource, diag) {
  let code = typeof diag.code === "number" ? String(diag.code) : diag.code;
  return {
    severity: toSeverity(diag.severity),
    startLineNumber: diag.range.start.line + 1,
    startColumn: diag.range.start.character + 1,
    endLineNumber: diag.range.end.line + 1,
    endColumn: diag.range.end.character + 1,
    message: diag.message,
    code,
    source: diag.source
  };
}
var CompletionAdapter = class {
  constructor(_worker, _triggerCharacters) {
    this._worker = _worker;
    this._triggerCharacters = _triggerCharacters;
  }
  get triggerCharacters() {
    return this._triggerCharacters;
  }
  provideCompletionItems(model, position, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doComplete(resource.toString(), fromPosition(position));
    }).then((info) => {
      if (!info) {
        return;
      }
      const wordInfo = model.getWordUntilPosition(position);
      const wordRange = new monaco_editor_core_exports.Range(position.lineNumber, wordInfo.startColumn, position.lineNumber, wordInfo.endColumn);
      const items = info.items.map((entry) => {
        const item = {
          label: entry.label,
          insertText: entry.insertText || entry.label,
          sortText: entry.sortText,
          filterText: entry.filterText,
          documentation: entry.documentation,
          detail: entry.detail,
          command: toCommand(entry.command),
          range: wordRange,
          kind: toCompletionItemKind(entry.kind)
        };
        if (entry.textEdit) {
          if (isInsertReplaceEdit(entry.textEdit)) {
            item.range = {
              insert: toRange(entry.textEdit.insert),
              replace: toRange(entry.textEdit.replace)
            };
          } else {
            item.range = toRange(entry.textEdit.range);
          }
          item.insertText = entry.textEdit.newText;
        }
        if (entry.additionalTextEdits) {
          item.additionalTextEdits = entry.additionalTextEdits.map(toTextEdit);
        }
        if (entry.insertTextFormat === InsertTextFormat.Snippet) {
          item.insertTextRules = monaco_editor_core_exports.languages.CompletionItemInsertTextRule.InsertAsSnippet;
        }
        return item;
      });
      return {
        isIncomplete: info.isIncomplete,
        suggestions: items
      };
    });
  }
};
function fromPosition(position) {
  if (!position) {
    return void 0;
  }
  return { character: position.column - 1, line: position.lineNumber - 1 };
}
function fromRange(range) {
  if (!range) {
    return void 0;
  }
  return {
    start: {
      line: range.startLineNumber - 1,
      character: range.startColumn - 1
    },
    end: { line: range.endLineNumber - 1, character: range.endColumn - 1 }
  };
}
function toRange(range) {
  if (!range) {
    return void 0;
  }
  return new monaco_editor_core_exports.Range(range.start.line + 1, range.start.character + 1, range.end.line + 1, range.end.character + 1);
}
function isInsertReplaceEdit(edit) {
  return typeof edit.insert !== "undefined" && typeof edit.replace !== "undefined";
}
function toCompletionItemKind(kind) {
  const mItemKind = monaco_editor_core_exports.languages.CompletionItemKind;
  switch (kind) {
    case CompletionItemKind.Text:
      return mItemKind.Text;
    case CompletionItemKind.Method:
      return mItemKind.Method;
    case CompletionItemKind.Function:
      return mItemKind.Function;
    case CompletionItemKind.Constructor:
      return mItemKind.Constructor;
    case CompletionItemKind.Field:
      return mItemKind.Field;
    case CompletionItemKind.Variable:
      return mItemKind.Variable;
    case CompletionItemKind.Class:
      return mItemKind.Class;
    case CompletionItemKind.Interface:
      return mItemKind.Interface;
    case CompletionItemKind.Module:
      return mItemKind.Module;
    case CompletionItemKind.Property:
      return mItemKind.Property;
    case CompletionItemKind.Unit:
      return mItemKind.Unit;
    case CompletionItemKind.Value:
      return mItemKind.Value;
    case CompletionItemKind.Enum:
      return mItemKind.Enum;
    case CompletionItemKind.Keyword:
      return mItemKind.Keyword;
    case CompletionItemKind.Snippet:
      return mItemKind.Snippet;
    case CompletionItemKind.Color:
      return mItemKind.Color;
    case CompletionItemKind.File:
      return mItemKind.File;
    case CompletionItemKind.Reference:
      return mItemKind.Reference;
  }
  return mItemKind.Property;
}
function toTextEdit(textEdit) {
  if (!textEdit) {
    return void 0;
  }
  return {
    range: toRange(textEdit.range),
    text: textEdit.newText
  };
}
function toCommand(c) {
  return c && c.command === "editor.action.triggerSuggest" ? { id: c.command, title: c.title, arguments: c.arguments } : void 0;
}
var HoverAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideHover(model, position, token) {
    let resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doHover(resource.toString(), fromPosition(position));
    }).then((info) => {
      if (!info) {
        return;
      }
      return {
        range: toRange(info.range),
        contents: toMarkedStringArray(info.contents)
      };
    });
  }
};
function isMarkupContent(thing) {
  return thing && typeof thing === "object" && typeof thing.kind === "string";
}
function toMarkdownString(entry) {
  if (typeof entry === "string") {
    return {
      value: entry
    };
  }
  if (isMarkupContent(entry)) {
    if (entry.kind === "plaintext") {
      return {
        value: entry.value.replace(/[\\`*_{}[\]()#+\-.!]/g, "\\$&")
      };
    }
    return {
      value: entry.value
    };
  }
  return { value: "```" + entry.language + "\n" + entry.value + "\n```\n" };
}
function toMarkedStringArray(contents) {
  if (!contents) {
    return void 0;
  }
  if (Array.isArray(contents)) {
    return contents.map(toMarkdownString);
  }
  return [toMarkdownString(contents)];
}
var DocumentHighlightAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentHighlights(model, position, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentHighlights(resource.toString(), fromPosition(position))).then((entries) => {
      if (!entries) {
        return;
      }
      return entries.map((entry) => {
        return {
          range: toRange(entry.range),
          kind: toDocumentHighlightKind(entry.kind)
        };
      });
    });
  }
};
function toDocumentHighlightKind(kind) {
  switch (kind) {
    case DocumentHighlightKind.Read:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Read;
    case DocumentHighlightKind.Write:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Write;
    case DocumentHighlightKind.Text:
      return monaco_editor_core_exports.languages.DocumentHighlightKind.Text;
  }
  return monaco_editor_core_exports.languages.DocumentHighlightKind.Text;
}
var DefinitionAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDefinition(model, position, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.findDefinition(resource.toString(), fromPosition(position));
    }).then((definition) => {
      if (!definition) {
        return;
      }
      return [toLocation(definition)];
    });
  }
};
function toLocation(location) {
  return {
    uri: monaco_editor_core_exports.Uri.parse(location.uri),
    range: toRange(location.range)
  };
}
var ReferenceAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideReferences(model, position, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.findReferences(resource.toString(), fromPosition(position));
    }).then((entries) => {
      if (!entries) {
        return;
      }
      return entries.map(toLocation);
    });
  }
};
var RenameAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideRenameEdits(model, position, newName, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.doRename(resource.toString(), fromPosition(position), newName);
    }).then((edit) => {
      return toWorkspaceEdit(edit);
    });
  }
};
function toWorkspaceEdit(edit) {
  if (!edit || !edit.changes) {
    return void 0;
  }
  let resourceEdits = [];
  for (let uri in edit.changes) {
    const _uri = monaco_editor_core_exports.Uri.parse(uri);
    for (let e of edit.changes[uri]) {
      resourceEdits.push({
        resource: _uri,
        versionId: void 0,
        textEdit: {
          range: toRange(e.range),
          text: e.newText
        }
      });
    }
  }
  return {
    edits: resourceEdits
  };
}
var DocumentSymbolAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentSymbols(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentSymbols(resource.toString())).then((items) => {
      if (!items) {
        return;
      }
      return items.map((item) => ({
        name: item.name,
        detail: "",
        containerName: item.containerName,
        kind: toSymbolKind(item.kind),
        range: toRange(item.location.range),
        selectionRange: toRange(item.location.range),
        tags: []
      }));
    });
  }
};
function toSymbolKind(kind) {
  let mKind = monaco_editor_core_exports.languages.SymbolKind;
  switch (kind) {
    case SymbolKind.File:
      return mKind.Array;
    case SymbolKind.Module:
      return mKind.Module;
    case SymbolKind.Namespace:
      return mKind.Namespace;
    case SymbolKind.Package:
      return mKind.Package;
    case SymbolKind.Class:
      return mKind.Class;
    case SymbolKind.Method:
      return mKind.Method;
    case SymbolKind.Property:
      return mKind.Property;
    case SymbolKind.Field:
      return mKind.Field;
    case SymbolKind.Constructor:
      return mKind.Constructor;
    case SymbolKind.Enum:
      return mKind.Enum;
    case SymbolKind.Interface:
      return mKind.Interface;
    case SymbolKind.Function:
      return mKind.Function;
    case SymbolKind.Variable:
      return mKind.Variable;
    case SymbolKind.Constant:
      return mKind.Constant;
    case SymbolKind.String:
      return mKind.String;
    case SymbolKind.Number:
      return mKind.Number;
    case SymbolKind.Boolean:
      return mKind.Boolean;
    case SymbolKind.Array:
      return mKind.Array;
  }
  return mKind.Function;
}
var DocumentLinkAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideLinks(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentLinks(resource.toString())).then((items) => {
      if (!items) {
        return;
      }
      return {
        links: items.map((item) => ({
          range: toRange(item.range),
          url: item.target
        }))
      };
    });
  }
};
var DocumentFormattingEditProvider = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentFormattingEdits(model, options, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.format(resource.toString(), null, fromFormattingOptions(options)).then((edits) => {
        if (!edits || edits.length === 0) {
          return;
        }
        return edits.map(toTextEdit);
      });
    });
  }
};
var DocumentRangeFormattingEditProvider = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  canFormatMultipleRanges = false;
  provideDocumentRangeFormattingEdits(model, range, options, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => {
      return worker.format(resource.toString(), fromRange(range), fromFormattingOptions(options)).then((edits) => {
        if (!edits || edits.length === 0) {
          return;
        }
        return edits.map(toTextEdit);
      });
    });
  }
};
function fromFormattingOptions(options) {
  return {
    tabSize: options.tabSize,
    insertSpaces: options.insertSpaces
  };
}
var DocumentColorAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideDocumentColors(model, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.findDocumentColors(resource.toString())).then((infos) => {
      if (!infos) {
        return;
      }
      return infos.map((item) => ({
        color: item.color,
        range: toRange(item.range)
      }));
    });
  }
  provideColorPresentations(model, info, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getColorPresentations(resource.toString(), info.color, fromRange(info.range))).then((presentations) => {
      if (!presentations) {
        return;
      }
      return presentations.map((presentation) => {
        let item = {
          label: presentation.label
        };
        if (presentation.textEdit) {
          item.textEdit = toTextEdit(presentation.textEdit);
        }
        if (presentation.additionalTextEdits) {
          item.additionalTextEdits = presentation.additionalTextEdits.map(toTextEdit);
        }
        return item;
      });
    });
  }
};
var FoldingRangeAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideFoldingRanges(model, context, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getFoldingRanges(resource.toString(), context)).then((ranges) => {
      if (!ranges) {
        return;
      }
      return ranges.map((range) => {
        const result = {
          start: range.startLine + 1,
          end: range.endLine + 1
        };
        if (typeof range.kind !== "undefined") {
          result.kind = toFoldingRangeKind(range.kind);
        }
        return result;
      });
    });
  }
};
function toFoldingRangeKind(kind) {
  switch (kind) {
    case FoldingRangeKind.Comment:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Comment;
    case FoldingRangeKind.Imports:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Imports;
    case FoldingRangeKind.Region:
      return monaco_editor_core_exports.languages.FoldingRangeKind.Region;
  }
  return void 0;
}
var SelectionRangeAdapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  provideSelectionRanges(model, positions, token) {
    const resource = model.uri;
    return this._worker(resource).then((worker) => worker.getSelectionRanges(resource.toString(), positions.map(fromPosition))).then((selectionRanges) => {
      if (!selectionRanges) {
        return;
      }
      return selectionRanges.map((selectionRange) => {
        const result = [];
        while (selectionRange) {
          result.push({ range: toRange(selectionRange.range) });
          selectionRange = selectionRange.parent;
        }
        return result;
      });
    });
  }
};

// node_modules/jsonc-parser/lib/esm/impl/scanner.js
function createScanner(text, ignoreTrivia) {
  if (ignoreTrivia === void 0) {
    ignoreTrivia = false;
  }
  var len = text.length;
  var pos = 0, value = "", tokenOffset = 0, token = 16, lineNumber = 0, lineStartOffset = 0, tokenLineStartOffset = 0, prevTokenLineStartOffset = 0, scanError = 0;
  function scanHexDigits(count, exact) {
    var digits = 0;
    var value2 = 0;
    while (digits < count || !exact) {
      var ch = text.charCodeAt(pos);
      if (ch >= 48 && ch <= 57) {
        value2 = value2 * 16 + ch - 48;
      } else if (ch >= 65 && ch <= 70) {
        value2 = value2 * 16 + ch - 65 + 10;
      } else if (ch >= 97 && ch <= 102) {
        value2 = value2 * 16 + ch - 97 + 10;
      } else {
        break;
      }
      pos++;
      digits++;
    }
    if (digits < count) {
      value2 = -1;
    }
    return value2;
  }
  function setPosition(newPosition) {
    pos = newPosition;
    value = "";
    tokenOffset = 0;
    token = 16;
    scanError = 0;
  }
  function scanNumber() {
    var start = pos;
    if (text.charCodeAt(pos) === 48) {
      pos++;
    } else {
      pos++;
      while (pos < text.length && isDigit(text.charCodeAt(pos))) {
        pos++;
      }
    }
    if (pos < text.length && text.charCodeAt(pos) === 46) {
      pos++;
      if (pos < text.length && isDigit(text.charCodeAt(pos))) {
        pos++;
        while (pos < text.length && isDigit(text.charCodeAt(pos))) {
          pos++;
        }
      } else {
        scanError = 3;
        return text.substring(start, pos);
      }
    }
    var end = pos;
    if (pos < text.length && (text.charCodeAt(pos) === 69 || text.charCodeAt(pos) === 101)) {
      pos++;
      if (pos < text.length && text.charCodeAt(pos) === 43 || text.charCodeAt(pos) === 45) {
        pos++;
      }
      if (pos < text.length && isDigit(text.charCodeAt(pos))) {
        pos++;
        while (pos < text.length && isDigit(text.charCodeAt(pos))) {
          pos++;
        }
        end = pos;
      } else {
        scanError = 3;
      }
    }
    return text.substring(start, end);
  }
  function scanString() {
    var result = "", start = pos;
    while (true) {
      if (pos >= len) {
        result += text.substring(start, pos);
        scanError = 2;
        break;
      }
      var ch = text.charCodeAt(pos);
      if (ch === 34) {
        result += text.substring(start, pos);
        pos++;
        break;
      }
      if (ch === 92) {
        result += text.substring(start, pos);
        pos++;
        if (pos >= len) {
          scanError = 2;
          break;
        }
        var ch2 = text.charCodeAt(pos++);
        switch (ch2) {
          case 34:
            result += '"';
            break;
          case 92:
            result += "\\";
            break;
          case 47:
            result += "/";
            break;
          case 98:
            result += "\b";
            break;
          case 102:
            result += "\f";
            break;
          case 110:
            result += "\n";
            break;
          case 114:
            result += "\r";
            break;
          case 116:
            result += "	";
            break;
          case 117:
            var ch3 = scanHexDigits(4, true);
            if (ch3 >= 0) {
              result += String.fromCharCode(ch3);
            } else {
              scanError = 4;
            }
            break;
          default:
            scanError = 5;
        }
        start = pos;
        continue;
      }
      if (ch >= 0 && ch <= 31) {
        if (isLineBreak(ch)) {
          result += text.substring(start, pos);
          scanError = 2;
          break;
        } else {
          scanError = 6;
        }
      }
      pos++;
    }
    return result;
  }
  function scanNext() {
    value = "";
    scanError = 0;
    tokenOffset = pos;
    lineStartOffset = lineNumber;
    prevTokenLineStartOffset = tokenLineStartOffset;
    if (pos >= len) {
      tokenOffset = len;
      return token = 17;
    }
    var code = text.charCodeAt(pos);
    if (isWhiteSpace(code)) {
      do {
        pos++;
        value += String.fromCharCode(code);
        code = text.charCodeAt(pos);
      } while (isWhiteSpace(code));
      return token = 15;
    }
    if (isLineBreak(code)) {
      pos++;
      value += String.fromCharCode(code);
      if (code === 13 && text.charCodeAt(pos) === 10) {
        pos++;
        value += "\n";
      }
      lineNumber++;
      tokenLineStartOffset = pos;
      return token = 14;
    }
    switch (code) {
      case 123:
        pos++;
        return token = 1;
      case 125:
        pos++;
        return token = 2;
      case 91:
        pos++;
        return token = 3;
      case 93:
        pos++;
        return token = 4;
      case 58:
        pos++;
        return token = 6;
      case 44:
        pos++;
        return token = 5;
      case 34:
        pos++;
        value = scanString();
        return token = 10;
      case 47:
        var start = pos - 1;
        if (text.charCodeAt(pos + 1) === 47) {
          pos += 2;
          while (pos < len) {
            if (isLineBreak(text.charCodeAt(pos))) {
              break;
            }
            pos++;
          }
          value = text.substring(start, pos);
          return token = 12;
        }
        if (text.charCodeAt(pos + 1) === 42) {
          pos += 2;
          var safeLength = len - 1;
          var commentClosed = false;
          while (pos < safeLength) {
            var ch = text.charCodeAt(pos);
            if (ch === 42 && text.charCodeAt(pos + 1) === 47) {
              pos += 2;
              commentClosed = true;
              break;
            }
            pos++;
            if (isLineBreak(ch)) {
              if (ch === 13 && text.charCodeAt(pos) === 10) {
                pos++;
              }
              lineNumber++;
              tokenLineStartOffset = pos;
            }
          }
          if (!commentClosed) {
            pos++;
            scanError = 1;
          }
          value = text.substring(start, pos);
          return token = 13;
        }
        value += String.fromCharCode(code);
        pos++;
        return token = 16;
      case 45:
        value += String.fromCharCode(code);
        pos++;
        if (pos === len || !isDigit(text.charCodeAt(pos))) {
          return token = 16;
        }
      case 48:
      case 49:
      case 50:
      case 51:
      case 52:
      case 53:
      case 54:
      case 55:
      case 56:
      case 57:
        value += scanNumber();
        return token = 11;
      default:
        while (pos < len && isUnknownContentCharacter(code)) {
          pos++;
          code = text.charCodeAt(pos);
        }
        if (tokenOffset !== pos) {
          value = text.substring(tokenOffset, pos);
          switch (value) {
            case "true":
              return token = 8;
            case "false":
              return token = 9;
            case "null":
              return token = 7;
          }
          return token = 16;
        }
        value += String.fromCharCode(code);
        pos++;
        return token = 16;
    }
  }
  function isUnknownContentCharacter(code) {
    if (isWhiteSpace(code) || isLineBreak(code)) {
      return false;
    }
    switch (code) {
      case 125:
      case 93:
      case 123:
      case 91:
      case 34:
      case 58:
      case 44:
      case 47:
        return false;
    }
    return true;
  }
  function scanNextNonTrivia() {
    var result;
    do {
      result = scanNext();
    } while (result >= 12 && result <= 15);
    return result;
  }
  return {
    setPosition,
    getPosition: function() {
      return pos;
    },
    scan: ignoreTrivia ? scanNextNonTrivia : scanNext,
    getToken: function() {
      return token;
    },
    getTokenValue: function() {
      return value;
    },
    getTokenOffset: function() {
      return tokenOffset;
    },
    getTokenLength: function() {
      return pos - tokenOffset;
    },
    getTokenStartLine: function() {
      return lineStartOffset;
    },
    getTokenStartCharacter: function() {
      return tokenOffset - prevTokenLineStartOffset;
    },
    getTokenError: function() {
      return scanError;
    }
  };
}
function isWhiteSpace(ch) {
  return ch === 32 || ch === 9 || ch === 11 || ch === 12 || ch === 160 || ch === 5760 || ch >= 8192 && ch <= 8203 || ch === 8239 || ch === 8287 || ch === 12288 || ch === 65279;
}
function isLineBreak(ch) {
  return ch === 10 || ch === 13 || ch === 8232 || ch === 8233;
}
function isDigit(ch) {
  return ch >= 48 && ch <= 57;
}

// node_modules/jsonc-parser/lib/esm/impl/parser.js
var ParseOptions;
(function(ParseOptions2) {
  ParseOptions2.DEFAULT = {
    allowTrailingComma: false
  };
})(ParseOptions || (ParseOptions = {}));

// node_modules/jsonc-parser/lib/esm/main.js
var createScanner2 = createScanner;

// src/language/json/tokenization.ts
function createTokenizationSupport(supportComments) {
  return {
    getInitialState: () => new JSONState(null, null, false, null),
    tokenize: (line, state) => tokenize(supportComments, line, state)
  };
}
var TOKEN_DELIM_OBJECT = "delimiter.bracket.json";
var TOKEN_DELIM_ARRAY = "delimiter.array.json";
var TOKEN_DELIM_COLON = "delimiter.colon.json";
var TOKEN_DELIM_COMMA = "delimiter.comma.json";
var TOKEN_VALUE_BOOLEAN = "keyword.json";
var TOKEN_VALUE_NULL = "keyword.json";
var TOKEN_VALUE_STRING = "string.value.json";
var TOKEN_VALUE_NUMBER = "number.json";
var TOKEN_PROPERTY_NAME = "string.key.json";
var TOKEN_COMMENT_BLOCK = "comment.block.json";
var TOKEN_COMMENT_LINE = "comment.line.json";
var ParentsStack = class {
  constructor(parent, type) {
    this.parent = parent;
    this.type = type;
  }
  static pop(parents) {
    if (parents) {
      return parents.parent;
    }
    return null;
  }
  static push(parents, type) {
    return new ParentsStack(parents, type);
  }
  static equals(a, b) {
    if (!a && !b) {
      return true;
    }
    if (!a || !b) {
      return false;
    }
    while (a && b) {
      if (a === b) {
        return true;
      }
      if (a.type !== b.type) {
        return false;
      }
      a = a.parent;
      b = b.parent;
    }
    return true;
  }
};
var JSONState = class {
  _state;
  scanError;
  lastWasColon;
  parents;
  constructor(state, scanError, lastWasColon, parents) {
    this._state = state;
    this.scanError = scanError;
    this.lastWasColon = lastWasColon;
    this.parents = parents;
  }
  clone() {
    return new JSONState(this._state, this.scanError, this.lastWasColon, this.parents);
  }
  equals(other) {
    if (other === this) {
      return true;
    }
    if (!other || !(other instanceof JSONState)) {
      return false;
    }
    return this.scanError === other.scanError && this.lastWasColon === other.lastWasColon && ParentsStack.equals(this.parents, other.parents);
  }
  getStateData() {
    return this._state;
  }
  setStateData(state) {
    this._state = state;
  }
};
function tokenize(comments, line, state, offsetDelta = 0) {
  let numberOfInsertedCharacters = 0;
  let adjustOffset = false;
  switch (state.scanError) {
    case 2 /* UnexpectedEndOfString */:
      line = '"' + line;
      numberOfInsertedCharacters = 1;
      break;
    case 1 /* UnexpectedEndOfComment */:
      line = "/*" + line;
      numberOfInsertedCharacters = 2;
      break;
  }
  const scanner = createScanner2(line);
  let lastWasColon = state.lastWasColon;
  let parents = state.parents;
  const ret = {
    tokens: [],
    endState: state.clone()
  };
  while (true) {
    let offset = offsetDelta + scanner.getPosition();
    let type = "";
    const kind = scanner.scan();
    if (kind === 17 /* EOF */) {
      break;
    }
    if (offset === offsetDelta + scanner.getPosition()) {
      throw new Error("Scanner did not advance, next 3 characters are: " + line.substr(scanner.getPosition(), 3));
    }
    if (adjustOffset) {
      offset -= numberOfInsertedCharacters;
    }
    adjustOffset = numberOfInsertedCharacters > 0;
    switch (kind) {
      case 1 /* OpenBraceToken */:
        parents = ParentsStack.push(parents, 0 /* Object */);
        type = TOKEN_DELIM_OBJECT;
        lastWasColon = false;
        break;
      case 2 /* CloseBraceToken */:
        parents = ParentsStack.pop(parents);
        type = TOKEN_DELIM_OBJECT;
        lastWasColon = false;
        break;
      case 3 /* OpenBracketToken */:
        parents = ParentsStack.push(parents, 1 /* Array */);
        type = TOKEN_DELIM_ARRAY;
        lastWasColon = false;
        break;
      case 4 /* CloseBracketToken */:
        parents = ParentsStack.pop(parents);
        type = TOKEN_DELIM_ARRAY;
        lastWasColon = false;
        break;
      case 6 /* ColonToken */:
        type = TOKEN_DELIM_COLON;
        lastWasColon = true;
        break;
      case 5 /* CommaToken */:
        type = TOKEN_DELIM_COMMA;
        lastWasColon = false;
        break;
      case 8 /* TrueKeyword */:
      case 9 /* FalseKeyword */:
        type = TOKEN_VALUE_BOOLEAN;
        lastWasColon = false;
        break;
      case 7 /* NullKeyword */:
        type = TOKEN_VALUE_NULL;
        lastWasColon = false;
        break;
      case 10 /* StringLiteral */:
        const currentParent = parents ? parents.type : 0 /* Object */;
        const inArray = currentParent === 1 /* Array */;
        type = lastWasColon || inArray ? TOKEN_VALUE_STRING : TOKEN_PROPERTY_NAME;
        lastWasColon = false;
        break;
      case 11 /* NumericLiteral */:
        type = TOKEN_VALUE_NUMBER;
        lastWasColon = false;
        break;
    }
    if (comments) {
      switch (kind) {
        case 12 /* LineCommentTrivia */:
          type = TOKEN_COMMENT_LINE;
          break;
        case 13 /* BlockCommentTrivia */:
          type = TOKEN_COMMENT_BLOCK;
          break;
      }
    }
    ret.endState = new JSONState(state.getStateData(), scanner.getTokenError(), lastWasColon, parents);
    ret.tokens.push({
      startIndex: offset,
      scopes: type
    });
  }
  return ret;
}

// src/language/json/jsonMode.ts
var JSONDiagnosticsAdapter = class extends DiagnosticsAdapter {
  constructor(languageId, worker, defaults) {
    super(languageId, worker, defaults.onDidChange);
    this._disposables.push(monaco_editor_core_exports.editor.onWillDisposeModel((model) => {
      this._resetSchema(model.uri);
    }));
    this._disposables.push(monaco_editor_core_exports.editor.onDidChangeModelLanguage((event) => {
      this._resetSchema(event.model.uri);
    }));
  }
  _resetSchema(resource) {
    this._worker().then((worker) => {
      worker.resetSchema(resource.toString());
    });
  }
};
function setupMode(defaults) {
  const disposables = [];
  const providers = [];
  const client = new WorkerManager(defaults);
  disposables.push(client);
  const worker = (...uris) => {
    return client.getLanguageServiceWorker(...uris);
  };
  function registerProviders() {
    const { languageId, modeConfiguration: modeConfiguration2 } = defaults;
    disposeAll(providers);
    if (modeConfiguration2.documentFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentFormattingEditProvider(languageId, new DocumentFormattingEditProvider(worker)));
    }
    if (modeConfiguration2.documentRangeFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentRangeFormattingEditProvider(languageId, new DocumentRangeFormattingEditProvider(worker)));
    }
    if (modeConfiguration2.completionItems) {
      providers.push(monaco_editor_core_exports.languages.registerCompletionItemProvider(languageId, new CompletionAdapter(worker, [" ", ":", '"'])));
    }
    if (modeConfiguration2.hovers) {
      providers.push(monaco_editor_core_exports.languages.registerHoverProvider(languageId, new HoverAdapter(worker)));
    }
    if (modeConfiguration2.documentSymbols) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentSymbolProvider(languageId, new DocumentSymbolAdapter(worker)));
    }
    if (modeConfiguration2.tokens) {
      providers.push(monaco_editor_core_exports.languages.setTokensProvider(languageId, createTokenizationSupport(true)));
    }
    if (modeConfiguration2.colors) {
      providers.push(monaco_editor_core_exports.languages.registerColorProvider(languageId, new DocumentColorAdapter(worker)));
    }
    if (modeConfiguration2.foldingRanges) {
      providers.push(monaco_editor_core_exports.languages.registerFoldingRangeProvider(languageId, new FoldingRangeAdapter(worker)));
    }
    if (modeConfiguration2.diagnostics) {
      providers.push(new JSONDiagnosticsAdapter(languageId, worker, defaults));
    }
    if (modeConfiguration2.selectionRanges) {
      providers.push(monaco_editor_core_exports.languages.registerSelectionRangeProvider(languageId, new SelectionRangeAdapter(worker)));
    }
  }
  registerProviders();
  disposables.push(monaco_editor_core_exports.languages.setLanguageConfiguration(defaults.languageId, richEditConfiguration));
  let modeConfiguration = defaults.modeConfiguration;
  defaults.onDidChange((newDefaults) => {
    if (newDefaults.modeConfiguration !== modeConfiguration) {
      modeConfiguration = newDefaults.modeConfiguration;
      registerProviders();
    }
  });
  disposables.push(asDisposable(providers));
  return asDisposable(disposables);
}
function asDisposable(disposables) {
  return { dispose: () => disposeAll(disposables) };
}
function disposeAll(disposables) {
  while (disposables.length) {
    disposables.pop().dispose();
  }
}
var richEditConfiguration = {
  wordPattern: /(-?\d*\.\d\w*)|([^\[\{\]\}\:\"\,\s]+)/g,
  comments: {
    lineComment: "//",
    blockComment: ["/*", "*/"]
  },
  brackets: [
    ["{", "}"],
    ["[", "]"]
  ],
  autoClosingPairs: [
    { open: "{", close: "}", notIn: ["string"] },
    { open: "[", close: "]", notIn: ["string"] },
    { open: '"', close: '"', notIn: ["string"] }
  ]
};



/***/ }),

/***/ "./node_modules/monaco-editor/esm/vs/language/typescript/tsMode.js":
/*!*************************************************************************!*\
  !*** ./node_modules/monaco-editor/esm/vs/language/typescript/tsMode.js ***!
  \*************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   Adapter: function() { return /* binding */ Adapter; },
/* harmony export */   CodeActionAdaptor: function() { return /* binding */ CodeActionAdaptor; },
/* harmony export */   DefinitionAdapter: function() { return /* binding */ DefinitionAdapter; },
/* harmony export */   DiagnosticsAdapter: function() { return /* binding */ DiagnosticsAdapter; },
/* harmony export */   DocumentHighlightAdapter: function() { return /* binding */ DocumentHighlightAdapter; },
/* harmony export */   FormatAdapter: function() { return /* binding */ FormatAdapter; },
/* harmony export */   FormatHelper: function() { return /* binding */ FormatHelper; },
/* harmony export */   FormatOnTypeAdapter: function() { return /* binding */ FormatOnTypeAdapter; },
/* harmony export */   InlayHintsAdapter: function() { return /* binding */ InlayHintsAdapter; },
/* harmony export */   Kind: function() { return /* binding */ Kind; },
/* harmony export */   LibFiles: function() { return /* binding */ LibFiles; },
/* harmony export */   OutlineAdapter: function() { return /* binding */ OutlineAdapter; },
/* harmony export */   QuickInfoAdapter: function() { return /* binding */ QuickInfoAdapter; },
/* harmony export */   ReferenceAdapter: function() { return /* binding */ ReferenceAdapter; },
/* harmony export */   RenameAdapter: function() { return /* binding */ RenameAdapter; },
/* harmony export */   SignatureHelpAdapter: function() { return /* binding */ SignatureHelpAdapter; },
/* harmony export */   SuggestAdapter: function() { return /* binding */ SuggestAdapter; },
/* harmony export */   WorkerManager: function() { return /* binding */ WorkerManager; },
/* harmony export */   flattenDiagnosticMessageText: function() { return /* binding */ flattenDiagnosticMessageText; },
/* harmony export */   getJavaScriptWorker: function() { return /* binding */ getJavaScriptWorker; },
/* harmony export */   getTypeScriptWorker: function() { return /* binding */ getTypeScriptWorker; },
/* harmony export */   setupJavaScript: function() { return /* binding */ setupJavaScript; },
/* harmony export */   setupTypeScript: function() { return /* binding */ setupTypeScript; }
/* harmony export */ });
/* harmony import */ var _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../editor/editor.api.js */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.api.js");
/* harmony import */ var _monaco_contribution_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./monaco.contribution.js */ "./node_modules/monaco-editor/esm/vs/language/typescript/monaco.contribution.js");
/*!-----------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Version: 0.44.0(3e047efd345ff102c8c61b5398fb30845aaac166)
 * Released under the MIT license
 * https://github.com/microsoft/monaco-editor/blob/main/LICENSE.txt
 *-----------------------------------------------------------------------------*/

var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));
var __publicField = (obj, key, value) => {
  __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
  return value;
};

// src/fillers/monaco-editor-core.ts
var monaco_editor_core_exports = {};
__reExport(monaco_editor_core_exports, _editor_editor_api_js__WEBPACK_IMPORTED_MODULE_0__);


// src/language/typescript/workerManager.ts
var WorkerManager = class {
  constructor(_modeId, _defaults) {
    this._modeId = _modeId;
    this._defaults = _defaults;
    this._worker = null;
    this._client = null;
    this._configChangeListener = this._defaults.onDidChange(() => this._stopWorker());
    this._updateExtraLibsToken = 0;
    this._extraLibsChangeListener = this._defaults.onDidExtraLibsChange(() => this._updateExtraLibs());
  }
  _configChangeListener;
  _updateExtraLibsToken;
  _extraLibsChangeListener;
  _worker;
  _client;
  dispose() {
    this._configChangeListener.dispose();
    this._extraLibsChangeListener.dispose();
    this._stopWorker();
  }
  _stopWorker() {
    if (this._worker) {
      this._worker.dispose();
      this._worker = null;
    }
    this._client = null;
  }
  async _updateExtraLibs() {
    if (!this._worker) {
      return;
    }
    const myToken = ++this._updateExtraLibsToken;
    const proxy = await this._worker.getProxy();
    if (this._updateExtraLibsToken !== myToken) {
      return;
    }
    proxy.updateExtraLibs(this._defaults.getExtraLibs());
  }
  _getClient() {
    if (!this._client) {
      this._client = (async () => {
        this._worker = monaco_editor_core_exports.editor.createWebWorker({
          moduleId: "vs/language/typescript/tsWorker",
          label: this._modeId,
          keepIdleModels: true,
          createData: {
            compilerOptions: this._defaults.getCompilerOptions(),
            extraLibs: this._defaults.getExtraLibs(),
            customWorkerPath: this._defaults.workerOptions.customWorkerPath,
            inlayHintsOptions: this._defaults.inlayHintsOptions
          }
        });
        if (this._defaults.getEagerModelSync()) {
          return await this._worker.withSyncedResources(monaco_editor_core_exports.editor.getModels().filter((model) => model.getLanguageId() === this._modeId).map((model) => model.uri));
        }
        return await this._worker.getProxy();
      })();
    }
    return this._client;
  }
  async getLanguageServiceWorker(...resources) {
    const client = await this._getClient();
    if (this._worker) {
      await this._worker.withSyncedResources(resources);
    }
    return client;
  }
};

// src/language/typescript/languageFeatures.ts


// src/language/typescript/lib/lib.index.ts
var libFileSet = {};
libFileSet["lib.d.ts"] = true;
libFileSet["lib.decorators.d.ts"] = true;
libFileSet["lib.decorators.legacy.d.ts"] = true;
libFileSet["lib.dom.d.ts"] = true;
libFileSet["lib.dom.iterable.d.ts"] = true;
libFileSet["lib.es2015.collection.d.ts"] = true;
libFileSet["lib.es2015.core.d.ts"] = true;
libFileSet["lib.es2015.d.ts"] = true;
libFileSet["lib.es2015.generator.d.ts"] = true;
libFileSet["lib.es2015.iterable.d.ts"] = true;
libFileSet["lib.es2015.promise.d.ts"] = true;
libFileSet["lib.es2015.proxy.d.ts"] = true;
libFileSet["lib.es2015.reflect.d.ts"] = true;
libFileSet["lib.es2015.symbol.d.ts"] = true;
libFileSet["lib.es2015.symbol.wellknown.d.ts"] = true;
libFileSet["lib.es2016.array.include.d.ts"] = true;
libFileSet["lib.es2016.d.ts"] = true;
libFileSet["lib.es2016.full.d.ts"] = true;
libFileSet["lib.es2017.d.ts"] = true;
libFileSet["lib.es2017.full.d.ts"] = true;
libFileSet["lib.es2017.intl.d.ts"] = true;
libFileSet["lib.es2017.object.d.ts"] = true;
libFileSet["lib.es2017.sharedmemory.d.ts"] = true;
libFileSet["lib.es2017.string.d.ts"] = true;
libFileSet["lib.es2017.typedarrays.d.ts"] = true;
libFileSet["lib.es2018.asyncgenerator.d.ts"] = true;
libFileSet["lib.es2018.asynciterable.d.ts"] = true;
libFileSet["lib.es2018.d.ts"] = true;
libFileSet["lib.es2018.full.d.ts"] = true;
libFileSet["lib.es2018.intl.d.ts"] = true;
libFileSet["lib.es2018.promise.d.ts"] = true;
libFileSet["lib.es2018.regexp.d.ts"] = true;
libFileSet["lib.es2019.array.d.ts"] = true;
libFileSet["lib.es2019.d.ts"] = true;
libFileSet["lib.es2019.full.d.ts"] = true;
libFileSet["lib.es2019.intl.d.ts"] = true;
libFileSet["lib.es2019.object.d.ts"] = true;
libFileSet["lib.es2019.string.d.ts"] = true;
libFileSet["lib.es2019.symbol.d.ts"] = true;
libFileSet["lib.es2020.bigint.d.ts"] = true;
libFileSet["lib.es2020.d.ts"] = true;
libFileSet["lib.es2020.date.d.ts"] = true;
libFileSet["lib.es2020.full.d.ts"] = true;
libFileSet["lib.es2020.intl.d.ts"] = true;
libFileSet["lib.es2020.number.d.ts"] = true;
libFileSet["lib.es2020.promise.d.ts"] = true;
libFileSet["lib.es2020.sharedmemory.d.ts"] = true;
libFileSet["lib.es2020.string.d.ts"] = true;
libFileSet["lib.es2020.symbol.wellknown.d.ts"] = true;
libFileSet["lib.es2021.d.ts"] = true;
libFileSet["lib.es2021.full.d.ts"] = true;
libFileSet["lib.es2021.intl.d.ts"] = true;
libFileSet["lib.es2021.promise.d.ts"] = true;
libFileSet["lib.es2021.string.d.ts"] = true;
libFileSet["lib.es2021.weakref.d.ts"] = true;
libFileSet["lib.es2022.array.d.ts"] = true;
libFileSet["lib.es2022.d.ts"] = true;
libFileSet["lib.es2022.error.d.ts"] = true;
libFileSet["lib.es2022.full.d.ts"] = true;
libFileSet["lib.es2022.intl.d.ts"] = true;
libFileSet["lib.es2022.object.d.ts"] = true;
libFileSet["lib.es2022.regexp.d.ts"] = true;
libFileSet["lib.es2022.sharedmemory.d.ts"] = true;
libFileSet["lib.es2022.string.d.ts"] = true;
libFileSet["lib.es2023.array.d.ts"] = true;
libFileSet["lib.es2023.d.ts"] = true;
libFileSet["lib.es2023.full.d.ts"] = true;
libFileSet["lib.es5.d.ts"] = true;
libFileSet["lib.es6.d.ts"] = true;
libFileSet["lib.esnext.d.ts"] = true;
libFileSet["lib.esnext.full.d.ts"] = true;
libFileSet["lib.esnext.intl.d.ts"] = true;
libFileSet["lib.scripthost.d.ts"] = true;
libFileSet["lib.webworker.d.ts"] = true;
libFileSet["lib.webworker.importscripts.d.ts"] = true;
libFileSet["lib.webworker.iterable.d.ts"] = true;

// src/language/typescript/languageFeatures.ts
function flattenDiagnosticMessageText(diag, newLine, indent = 0) {
  if (typeof diag === "string") {
    return diag;
  } else if (diag === void 0) {
    return "";
  }
  let result = "";
  if (indent) {
    result += newLine;
    for (let i = 0; i < indent; i++) {
      result += "  ";
    }
  }
  result += diag.messageText;
  indent++;
  if (diag.next) {
    for (const kid of diag.next) {
      result += flattenDiagnosticMessageText(kid, newLine, indent);
    }
  }
  return result;
}
function displayPartsToString(displayParts) {
  if (displayParts) {
    return displayParts.map((displayPart) => displayPart.text).join("");
  }
  return "";
}
var Adapter = class {
  constructor(_worker) {
    this._worker = _worker;
  }
  _textSpanToRange(model, span) {
    let p1 = model.getPositionAt(span.start);
    let p2 = model.getPositionAt(span.start + span.length);
    let { lineNumber: startLineNumber, column: startColumn } = p1;
    let { lineNumber: endLineNumber, column: endColumn } = p2;
    return { startLineNumber, startColumn, endLineNumber, endColumn };
  }
};
var LibFiles = class {
  constructor(_worker) {
    this._worker = _worker;
    this._libFiles = {};
    this._hasFetchedLibFiles = false;
    this._fetchLibFilesPromise = null;
  }
  _libFiles;
  _hasFetchedLibFiles;
  _fetchLibFilesPromise;
  isLibFile(uri) {
    if (!uri) {
      return false;
    }
    if (uri.path.indexOf("/lib.") === 0) {
      return !!libFileSet[uri.path.slice(1)];
    }
    return false;
  }
  getOrCreateModel(fileName) {
    const uri = monaco_editor_core_exports.Uri.parse(fileName);
    const model = monaco_editor_core_exports.editor.getModel(uri);
    if (model) {
      return model;
    }
    if (this.isLibFile(uri) && this._hasFetchedLibFiles) {
      return monaco_editor_core_exports.editor.createModel(this._libFiles[uri.path.slice(1)], "typescript", uri);
    }
    const matchedLibFile = _monaco_contribution_js__WEBPACK_IMPORTED_MODULE_1__.typescriptDefaults.getExtraLibs()[fileName];
    if (matchedLibFile) {
      return monaco_editor_core_exports.editor.createModel(matchedLibFile.content, "typescript", uri);
    }
    return null;
  }
  _containsLibFile(uris) {
    for (let uri of uris) {
      if (this.isLibFile(uri)) {
        return true;
      }
    }
    return false;
  }
  async fetchLibFilesIfNecessary(uris) {
    if (!this._containsLibFile(uris)) {
      return;
    }
    await this._fetchLibFiles();
  }
  _fetchLibFiles() {
    if (!this._fetchLibFilesPromise) {
      this._fetchLibFilesPromise = this._worker().then((w) => w.getLibFiles()).then((libFiles) => {
        this._hasFetchedLibFiles = true;
        this._libFiles = libFiles;
      });
    }
    return this._fetchLibFilesPromise;
  }
};
var DiagnosticsAdapter = class extends Adapter {
  constructor(_libFiles, _defaults, _selector, worker) {
    super(worker);
    this._libFiles = _libFiles;
    this._defaults = _defaults;
    this._selector = _selector;
    const onModelAdd = (model) => {
      if (model.getLanguageId() !== _selector) {
        return;
      }
      const maybeValidate = () => {
        const { onlyVisible } = this._defaults.getDiagnosticsOptions();
        if (onlyVisible) {
          if (model.isAttachedToEditor()) {
            this._doValidate(model);
          }
        } else {
          this._doValidate(model);
        }
      };
      let handle;
      const changeSubscription = model.onDidChangeContent(() => {
        clearTimeout(handle);
        handle = window.setTimeout(maybeValidate, 500);
      });
      const visibleSubscription = model.onDidChangeAttached(() => {
        const { onlyVisible } = this._defaults.getDiagnosticsOptions();
        if (onlyVisible) {
          if (model.isAttachedToEditor()) {
            maybeValidate();
          } else {
            monaco_editor_core_exports.editor.setModelMarkers(model, this._selector, []);
          }
        }
      });
      this._listener[model.uri.toString()] = {
        dispose() {
          changeSubscription.dispose();
          visibleSubscription.dispose();
          clearTimeout(handle);
        }
      };
      maybeValidate();
    };
    const onModelRemoved = (model) => {
      monaco_editor_core_exports.editor.setModelMarkers(model, this._selector, []);
      const key = model.uri.toString();
      if (this._listener[key]) {
        this._listener[key].dispose();
        delete this._listener[key];
      }
    };
    this._disposables.push(monaco_editor_core_exports.editor.onDidCreateModel((model) => onModelAdd(model)));
    this._disposables.push(monaco_editor_core_exports.editor.onWillDisposeModel(onModelRemoved));
    this._disposables.push(monaco_editor_core_exports.editor.onDidChangeModelLanguage((event) => {
      onModelRemoved(event.model);
      onModelAdd(event.model);
    }));
    this._disposables.push({
      dispose() {
        for (const model of monaco_editor_core_exports.editor.getModels()) {
          onModelRemoved(model);
        }
      }
    });
    const recomputeDiagostics = () => {
      for (const model of monaco_editor_core_exports.editor.getModels()) {
        onModelRemoved(model);
        onModelAdd(model);
      }
    };
    this._disposables.push(this._defaults.onDidChange(recomputeDiagostics));
    this._disposables.push(this._defaults.onDidExtraLibsChange(recomputeDiagostics));
    monaco_editor_core_exports.editor.getModels().forEach((model) => onModelAdd(model));
  }
  _disposables = [];
  _listener = /* @__PURE__ */ Object.create(null);
  dispose() {
    this._disposables.forEach((d) => d && d.dispose());
    this._disposables = [];
  }
  async _doValidate(model) {
    const worker = await this._worker(model.uri);
    if (model.isDisposed()) {
      return;
    }
    const promises = [];
    const { noSyntaxValidation, noSemanticValidation, noSuggestionDiagnostics } = this._defaults.getDiagnosticsOptions();
    if (!noSyntaxValidation) {
      promises.push(worker.getSyntacticDiagnostics(model.uri.toString()));
    }
    if (!noSemanticValidation) {
      promises.push(worker.getSemanticDiagnostics(model.uri.toString()));
    }
    if (!noSuggestionDiagnostics) {
      promises.push(worker.getSuggestionDiagnostics(model.uri.toString()));
    }
    const allDiagnostics = await Promise.all(promises);
    if (!allDiagnostics || model.isDisposed()) {
      return;
    }
    const diagnostics = allDiagnostics.reduce((p, c) => c.concat(p), []).filter((d) => (this._defaults.getDiagnosticsOptions().diagnosticCodesToIgnore || []).indexOf(d.code) === -1);
    const relatedUris = diagnostics.map((d) => d.relatedInformation || []).reduce((p, c) => c.concat(p), []).map((relatedInformation) => relatedInformation.file ? monaco_editor_core_exports.Uri.parse(relatedInformation.file.fileName) : null);
    await this._libFiles.fetchLibFilesIfNecessary(relatedUris);
    if (model.isDisposed()) {
      return;
    }
    monaco_editor_core_exports.editor.setModelMarkers(model, this._selector, diagnostics.map((d) => this._convertDiagnostics(model, d)));
  }
  _convertDiagnostics(model, diag) {
    const diagStart = diag.start || 0;
    const diagLength = diag.length || 1;
    const { lineNumber: startLineNumber, column: startColumn } = model.getPositionAt(diagStart);
    const { lineNumber: endLineNumber, column: endColumn } = model.getPositionAt(diagStart + diagLength);
    const tags = [];
    if (diag.reportsUnnecessary) {
      tags.push(monaco_editor_core_exports.MarkerTag.Unnecessary);
    }
    if (diag.reportsDeprecated) {
      tags.push(monaco_editor_core_exports.MarkerTag.Deprecated);
    }
    return {
      severity: this._tsDiagnosticCategoryToMarkerSeverity(diag.category),
      startLineNumber,
      startColumn,
      endLineNumber,
      endColumn,
      message: flattenDiagnosticMessageText(diag.messageText, "\n"),
      code: diag.code.toString(),
      tags,
      relatedInformation: this._convertRelatedInformation(model, diag.relatedInformation)
    };
  }
  _convertRelatedInformation(model, relatedInformation) {
    if (!relatedInformation) {
      return [];
    }
    const result = [];
    relatedInformation.forEach((info) => {
      let relatedResource = model;
      if (info.file) {
        relatedResource = this._libFiles.getOrCreateModel(info.file.fileName);
      }
      if (!relatedResource) {
        return;
      }
      const infoStart = info.start || 0;
      const infoLength = info.length || 1;
      const { lineNumber: startLineNumber, column: startColumn } = relatedResource.getPositionAt(infoStart);
      const { lineNumber: endLineNumber, column: endColumn } = relatedResource.getPositionAt(infoStart + infoLength);
      result.push({
        resource: relatedResource.uri,
        startLineNumber,
        startColumn,
        endLineNumber,
        endColumn,
        message: flattenDiagnosticMessageText(info.messageText, "\n")
      });
    });
    return result;
  }
  _tsDiagnosticCategoryToMarkerSeverity(category) {
    switch (category) {
      case 1 /* Error */:
        return monaco_editor_core_exports.MarkerSeverity.Error;
      case 3 /* Message */:
        return monaco_editor_core_exports.MarkerSeverity.Info;
      case 0 /* Warning */:
        return monaco_editor_core_exports.MarkerSeverity.Warning;
      case 2 /* Suggestion */:
        return monaco_editor_core_exports.MarkerSeverity.Hint;
    }
    return monaco_editor_core_exports.MarkerSeverity.Info;
  }
};
var SuggestAdapter = class extends Adapter {
  get triggerCharacters() {
    return ["."];
  }
  async provideCompletionItems(model, position, _context, token) {
    const wordInfo = model.getWordUntilPosition(position);
    const wordRange = new monaco_editor_core_exports.Range(position.lineNumber, wordInfo.startColumn, position.lineNumber, wordInfo.endColumn);
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const info = await worker.getCompletionsAtPosition(resource.toString(), offset);
    if (!info || model.isDisposed()) {
      return;
    }
    const suggestions = info.entries.map((entry) => {
      let range = wordRange;
      if (entry.replacementSpan) {
        const p1 = model.getPositionAt(entry.replacementSpan.start);
        const p2 = model.getPositionAt(entry.replacementSpan.start + entry.replacementSpan.length);
        range = new monaco_editor_core_exports.Range(p1.lineNumber, p1.column, p2.lineNumber, p2.column);
      }
      const tags = [];
      if (entry.kindModifiers !== void 0 && entry.kindModifiers.indexOf("deprecated") !== -1) {
        tags.push(monaco_editor_core_exports.languages.CompletionItemTag.Deprecated);
      }
      return {
        uri: resource,
        position,
        offset,
        range,
        label: entry.name,
        insertText: entry.name,
        sortText: entry.sortText,
        kind: SuggestAdapter.convertKind(entry.kind),
        tags
      };
    });
    return {
      suggestions
    };
  }
  async resolveCompletionItem(item, token) {
    const myItem = item;
    const resource = myItem.uri;
    const position = myItem.position;
    const offset = myItem.offset;
    const worker = await this._worker(resource);
    const details = await worker.getCompletionEntryDetails(resource.toString(), offset, myItem.label);
    if (!details) {
      return myItem;
    }
    return {
      uri: resource,
      position,
      label: details.name,
      kind: SuggestAdapter.convertKind(details.kind),
      detail: displayPartsToString(details.displayParts),
      documentation: {
        value: SuggestAdapter.createDocumentationString(details)
      }
    };
  }
  static convertKind(kind) {
    switch (kind) {
      case Kind.primitiveType:
      case Kind.keyword:
        return monaco_editor_core_exports.languages.CompletionItemKind.Keyword;
      case Kind.variable:
      case Kind.localVariable:
        return monaco_editor_core_exports.languages.CompletionItemKind.Variable;
      case Kind.memberVariable:
      case Kind.memberGetAccessor:
      case Kind.memberSetAccessor:
        return monaco_editor_core_exports.languages.CompletionItemKind.Field;
      case Kind.function:
      case Kind.memberFunction:
      case Kind.constructSignature:
      case Kind.callSignature:
      case Kind.indexSignature:
        return monaco_editor_core_exports.languages.CompletionItemKind.Function;
      case Kind.enum:
        return monaco_editor_core_exports.languages.CompletionItemKind.Enum;
      case Kind.module:
        return monaco_editor_core_exports.languages.CompletionItemKind.Module;
      case Kind.class:
        return monaco_editor_core_exports.languages.CompletionItemKind.Class;
      case Kind.interface:
        return monaco_editor_core_exports.languages.CompletionItemKind.Interface;
      case Kind.warning:
        return monaco_editor_core_exports.languages.CompletionItemKind.File;
    }
    return monaco_editor_core_exports.languages.CompletionItemKind.Property;
  }
  static createDocumentationString(details) {
    let documentationString = displayPartsToString(details.documentation);
    if (details.tags) {
      for (const tag of details.tags) {
        documentationString += `

${tagToString(tag)}`;
      }
    }
    return documentationString;
  }
};
function tagToString(tag) {
  let tagLabel = `*@${tag.name}*`;
  if (tag.name === "param" && tag.text) {
    const [paramName, ...rest] = tag.text;
    tagLabel += `\`${paramName.text}\``;
    if (rest.length > 0)
      tagLabel += ` \u2014 ${rest.map((r) => r.text).join(" ")}`;
  } else if (Array.isArray(tag.text)) {
    tagLabel += ` \u2014 ${tag.text.map((r) => r.text).join(" ")}`;
  } else if (tag.text) {
    tagLabel += ` \u2014 ${tag.text}`;
  }
  return tagLabel;
}
var SignatureHelpAdapter = class extends Adapter {
  signatureHelpTriggerCharacters = ["(", ","];
  static _toSignatureHelpTriggerReason(context) {
    switch (context.triggerKind) {
      case monaco_editor_core_exports.languages.SignatureHelpTriggerKind.TriggerCharacter:
        if (context.triggerCharacter) {
          if (context.isRetrigger) {
            return { kind: "retrigger", triggerCharacter: context.triggerCharacter };
          } else {
            return { kind: "characterTyped", triggerCharacter: context.triggerCharacter };
          }
        } else {
          return { kind: "invoked" };
        }
      case monaco_editor_core_exports.languages.SignatureHelpTriggerKind.ContentChange:
        return context.isRetrigger ? { kind: "retrigger" } : { kind: "invoked" };
      case monaco_editor_core_exports.languages.SignatureHelpTriggerKind.Invoke:
      default:
        return { kind: "invoked" };
    }
  }
  async provideSignatureHelp(model, position, token, context) {
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const info = await worker.getSignatureHelpItems(resource.toString(), offset, {
      triggerReason: SignatureHelpAdapter._toSignatureHelpTriggerReason(context)
    });
    if (!info || model.isDisposed()) {
      return;
    }
    const ret = {
      activeSignature: info.selectedItemIndex,
      activeParameter: info.argumentIndex,
      signatures: []
    };
    info.items.forEach((item) => {
      const signature = {
        label: "",
        parameters: []
      };
      signature.documentation = {
        value: displayPartsToString(item.documentation)
      };
      signature.label += displayPartsToString(item.prefixDisplayParts);
      item.parameters.forEach((p, i, a) => {
        const label = displayPartsToString(p.displayParts);
        const parameter = {
          label,
          documentation: {
            value: displayPartsToString(p.documentation)
          }
        };
        signature.label += label;
        signature.parameters.push(parameter);
        if (i < a.length - 1) {
          signature.label += displayPartsToString(item.separatorDisplayParts);
        }
      });
      signature.label += displayPartsToString(item.suffixDisplayParts);
      ret.signatures.push(signature);
    });
    return {
      value: ret,
      dispose() {
      }
    };
  }
};
var QuickInfoAdapter = class extends Adapter {
  async provideHover(model, position, token) {
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const info = await worker.getQuickInfoAtPosition(resource.toString(), offset);
    if (!info || model.isDisposed()) {
      return;
    }
    const documentation = displayPartsToString(info.documentation);
    const tags = info.tags ? info.tags.map((tag) => tagToString(tag)).join("  \n\n") : "";
    const contents = displayPartsToString(info.displayParts);
    return {
      range: this._textSpanToRange(model, info.textSpan),
      contents: [
        {
          value: "```typescript\n" + contents + "\n```\n"
        },
        {
          value: documentation + (tags ? "\n\n" + tags : "")
        }
      ]
    };
  }
};
var DocumentHighlightAdapter = class extends Adapter {
  async provideDocumentHighlights(model, position, token) {
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const entries = await worker.getDocumentHighlights(resource.toString(), offset, [
      resource.toString()
    ]);
    if (!entries || model.isDisposed()) {
      return;
    }
    return entries.flatMap((entry) => {
      return entry.highlightSpans.map((highlightSpans) => {
        return {
          range: this._textSpanToRange(model, highlightSpans.textSpan),
          kind: highlightSpans.kind === "writtenReference" ? monaco_editor_core_exports.languages.DocumentHighlightKind.Write : monaco_editor_core_exports.languages.DocumentHighlightKind.Text
        };
      });
    });
  }
};
var DefinitionAdapter = class extends Adapter {
  constructor(_libFiles, worker) {
    super(worker);
    this._libFiles = _libFiles;
  }
  async provideDefinition(model, position, token) {
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const entries = await worker.getDefinitionAtPosition(resource.toString(), offset);
    if (!entries || model.isDisposed()) {
      return;
    }
    await this._libFiles.fetchLibFilesIfNecessary(entries.map((entry) => monaco_editor_core_exports.Uri.parse(entry.fileName)));
    if (model.isDisposed()) {
      return;
    }
    const result = [];
    for (let entry of entries) {
      const refModel = this._libFiles.getOrCreateModel(entry.fileName);
      if (refModel) {
        result.push({
          uri: refModel.uri,
          range: this._textSpanToRange(refModel, entry.textSpan)
        });
      }
    }
    return result;
  }
};
var ReferenceAdapter = class extends Adapter {
  constructor(_libFiles, worker) {
    super(worker);
    this._libFiles = _libFiles;
  }
  async provideReferences(model, position, context, token) {
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const entries = await worker.getReferencesAtPosition(resource.toString(), offset);
    if (!entries || model.isDisposed()) {
      return;
    }
    await this._libFiles.fetchLibFilesIfNecessary(entries.map((entry) => monaco_editor_core_exports.Uri.parse(entry.fileName)));
    if (model.isDisposed()) {
      return;
    }
    const result = [];
    for (let entry of entries) {
      const refModel = this._libFiles.getOrCreateModel(entry.fileName);
      if (refModel) {
        result.push({
          uri: refModel.uri,
          range: this._textSpanToRange(refModel, entry.textSpan)
        });
      }
    }
    return result;
  }
};
var OutlineAdapter = class extends Adapter {
  async provideDocumentSymbols(model, token) {
    const resource = model.uri;
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const root = await worker.getNavigationTree(resource.toString());
    if (!root || model.isDisposed()) {
      return;
    }
    const convert = (item, containerLabel) => {
      const result2 = {
        name: item.text,
        detail: "",
        kind: outlineTypeTable[item.kind] || monaco_editor_core_exports.languages.SymbolKind.Variable,
        range: this._textSpanToRange(model, item.spans[0]),
        selectionRange: this._textSpanToRange(model, item.spans[0]),
        tags: [],
        children: item.childItems?.map((child) => convert(child, item.text)),
        containerName: containerLabel
      };
      return result2;
    };
    const result = root.childItems ? root.childItems.map((item) => convert(item)) : [];
    return result;
  }
};
var Kind = class {
};
__publicField(Kind, "unknown", "");
__publicField(Kind, "keyword", "keyword");
__publicField(Kind, "script", "script");
__publicField(Kind, "module", "module");
__publicField(Kind, "class", "class");
__publicField(Kind, "interface", "interface");
__publicField(Kind, "type", "type");
__publicField(Kind, "enum", "enum");
__publicField(Kind, "variable", "var");
__publicField(Kind, "localVariable", "local var");
__publicField(Kind, "function", "function");
__publicField(Kind, "localFunction", "local function");
__publicField(Kind, "memberFunction", "method");
__publicField(Kind, "memberGetAccessor", "getter");
__publicField(Kind, "memberSetAccessor", "setter");
__publicField(Kind, "memberVariable", "property");
__publicField(Kind, "constructorImplementation", "constructor");
__publicField(Kind, "callSignature", "call");
__publicField(Kind, "indexSignature", "index");
__publicField(Kind, "constructSignature", "construct");
__publicField(Kind, "parameter", "parameter");
__publicField(Kind, "typeParameter", "type parameter");
__publicField(Kind, "primitiveType", "primitive type");
__publicField(Kind, "label", "label");
__publicField(Kind, "alias", "alias");
__publicField(Kind, "const", "const");
__publicField(Kind, "let", "let");
__publicField(Kind, "warning", "warning");
var outlineTypeTable = /* @__PURE__ */ Object.create(null);
outlineTypeTable[Kind.module] = monaco_editor_core_exports.languages.SymbolKind.Module;
outlineTypeTable[Kind.class] = monaco_editor_core_exports.languages.SymbolKind.Class;
outlineTypeTable[Kind.enum] = monaco_editor_core_exports.languages.SymbolKind.Enum;
outlineTypeTable[Kind.interface] = monaco_editor_core_exports.languages.SymbolKind.Interface;
outlineTypeTable[Kind.memberFunction] = monaco_editor_core_exports.languages.SymbolKind.Method;
outlineTypeTable[Kind.memberVariable] = monaco_editor_core_exports.languages.SymbolKind.Property;
outlineTypeTable[Kind.memberGetAccessor] = monaco_editor_core_exports.languages.SymbolKind.Property;
outlineTypeTable[Kind.memberSetAccessor] = monaco_editor_core_exports.languages.SymbolKind.Property;
outlineTypeTable[Kind.variable] = monaco_editor_core_exports.languages.SymbolKind.Variable;
outlineTypeTable[Kind.const] = monaco_editor_core_exports.languages.SymbolKind.Variable;
outlineTypeTable[Kind.localVariable] = monaco_editor_core_exports.languages.SymbolKind.Variable;
outlineTypeTable[Kind.variable] = monaco_editor_core_exports.languages.SymbolKind.Variable;
outlineTypeTable[Kind.function] = monaco_editor_core_exports.languages.SymbolKind.Function;
outlineTypeTable[Kind.localFunction] = monaco_editor_core_exports.languages.SymbolKind.Function;
var FormatHelper = class extends Adapter {
  static _convertOptions(options) {
    return {
      ConvertTabsToSpaces: options.insertSpaces,
      TabSize: options.tabSize,
      IndentSize: options.tabSize,
      IndentStyle: 2 /* Smart */,
      NewLineCharacter: "\n",
      InsertSpaceAfterCommaDelimiter: true,
      InsertSpaceAfterSemicolonInForStatements: true,
      InsertSpaceBeforeAndAfterBinaryOperators: true,
      InsertSpaceAfterKeywordsInControlFlowStatements: true,
      InsertSpaceAfterFunctionKeywordForAnonymousFunctions: true,
      InsertSpaceAfterOpeningAndBeforeClosingNonemptyParenthesis: false,
      InsertSpaceAfterOpeningAndBeforeClosingNonemptyBrackets: false,
      InsertSpaceAfterOpeningAndBeforeClosingTemplateStringBraces: false,
      PlaceOpenBraceOnNewLineForControlBlocks: false,
      PlaceOpenBraceOnNewLineForFunctions: false
    };
  }
  _convertTextChanges(model, change) {
    return {
      text: change.newText,
      range: this._textSpanToRange(model, change.span)
    };
  }
};
var FormatAdapter = class extends FormatHelper {
  canFormatMultipleRanges = false;
  async provideDocumentRangeFormattingEdits(model, range, options, token) {
    const resource = model.uri;
    const startOffset = model.getOffsetAt({
      lineNumber: range.startLineNumber,
      column: range.startColumn
    });
    const endOffset = model.getOffsetAt({
      lineNumber: range.endLineNumber,
      column: range.endColumn
    });
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const edits = await worker.getFormattingEditsForRange(resource.toString(), startOffset, endOffset, FormatHelper._convertOptions(options));
    if (!edits || model.isDisposed()) {
      return;
    }
    return edits.map((edit) => this._convertTextChanges(model, edit));
  }
};
var FormatOnTypeAdapter = class extends FormatHelper {
  get autoFormatTriggerCharacters() {
    return [";", "}", "\n"];
  }
  async provideOnTypeFormattingEdits(model, position, ch, options, token) {
    const resource = model.uri;
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const edits = await worker.getFormattingEditsAfterKeystroke(resource.toString(), offset, ch, FormatHelper._convertOptions(options));
    if (!edits || model.isDisposed()) {
      return;
    }
    return edits.map((edit) => this._convertTextChanges(model, edit));
  }
};
var CodeActionAdaptor = class extends FormatHelper {
  async provideCodeActions(model, range, context, token) {
    const resource = model.uri;
    const start = model.getOffsetAt({
      lineNumber: range.startLineNumber,
      column: range.startColumn
    });
    const end = model.getOffsetAt({
      lineNumber: range.endLineNumber,
      column: range.endColumn
    });
    const formatOptions = FormatHelper._convertOptions(model.getOptions());
    const errorCodes = context.markers.filter((m) => m.code).map((m) => m.code).map(Number);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const codeFixes = await worker.getCodeFixesAtPosition(resource.toString(), start, end, errorCodes, formatOptions);
    if (!codeFixes || model.isDisposed()) {
      return { actions: [], dispose: () => {
      } };
    }
    const actions = codeFixes.filter((fix) => {
      return fix.changes.filter((change) => change.isNewFile).length === 0;
    }).map((fix) => {
      return this._tsCodeFixActionToMonacoCodeAction(model, context, fix);
    });
    return {
      actions,
      dispose: () => {
      }
    };
  }
  _tsCodeFixActionToMonacoCodeAction(model, context, codeFix) {
    const edits = [];
    for (const change of codeFix.changes) {
      for (const textChange of change.textChanges) {
        edits.push({
          resource: model.uri,
          versionId: void 0,
          textEdit: {
            range: this._textSpanToRange(model, textChange.span),
            text: textChange.newText
          }
        });
      }
    }
    const action = {
      title: codeFix.description,
      edit: { edits },
      diagnostics: context.markers,
      kind: "quickfix"
    };
    return action;
  }
};
var RenameAdapter = class extends Adapter {
  constructor(_libFiles, worker) {
    super(worker);
    this._libFiles = _libFiles;
  }
  async provideRenameEdits(model, position, newName, token) {
    const resource = model.uri;
    const fileName = resource.toString();
    const offset = model.getOffsetAt(position);
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return;
    }
    const renameInfo = await worker.getRenameInfo(fileName, offset, {
      allowRenameOfImportPath: false
    });
    if (renameInfo.canRename === false) {
      return {
        edits: [],
        rejectReason: renameInfo.localizedErrorMessage
      };
    }
    if (renameInfo.fileToRename !== void 0) {
      throw new Error("Renaming files is not supported.");
    }
    const renameLocations = await worker.findRenameLocations(fileName, offset, false, false, false);
    if (!renameLocations || model.isDisposed()) {
      return;
    }
    const edits = [];
    for (const renameLocation of renameLocations) {
      const model2 = this._libFiles.getOrCreateModel(renameLocation.fileName);
      if (model2) {
        edits.push({
          resource: model2.uri,
          versionId: void 0,
          textEdit: {
            range: this._textSpanToRange(model2, renameLocation.textSpan),
            text: newName
          }
        });
      } else {
        throw new Error(`Unknown file ${renameLocation.fileName}.`);
      }
    }
    return { edits };
  }
};
var InlayHintsAdapter = class extends Adapter {
  async provideInlayHints(model, range, token) {
    const resource = model.uri;
    const fileName = resource.toString();
    const start = model.getOffsetAt({
      lineNumber: range.startLineNumber,
      column: range.startColumn
    });
    const end = model.getOffsetAt({
      lineNumber: range.endLineNumber,
      column: range.endColumn
    });
    const worker = await this._worker(resource);
    if (model.isDisposed()) {
      return null;
    }
    const tsHints = await worker.provideInlayHints(fileName, start, end);
    const hints = tsHints.map((hint) => {
      return {
        ...hint,
        label: hint.text,
        position: model.getPositionAt(hint.position),
        kind: this._convertHintKind(hint.kind)
      };
    });
    return { hints, dispose: () => {
    } };
  }
  _convertHintKind(kind) {
    switch (kind) {
      case "Parameter":
        return monaco_editor_core_exports.languages.InlayHintKind.Parameter;
      case "Type":
        return monaco_editor_core_exports.languages.InlayHintKind.Type;
      default:
        return monaco_editor_core_exports.languages.InlayHintKind.Type;
    }
  }
};

// src/language/typescript/tsMode.ts
var javaScriptWorker;
var typeScriptWorker;
function setupTypeScript(defaults) {
  typeScriptWorker = setupMode(defaults, "typescript");
}
function setupJavaScript(defaults) {
  javaScriptWorker = setupMode(defaults, "javascript");
}
function getJavaScriptWorker() {
  return new Promise((resolve, reject) => {
    if (!javaScriptWorker) {
      return reject("JavaScript not registered!");
    }
    resolve(javaScriptWorker);
  });
}
function getTypeScriptWorker() {
  return new Promise((resolve, reject) => {
    if (!typeScriptWorker) {
      return reject("TypeScript not registered!");
    }
    resolve(typeScriptWorker);
  });
}
function setupMode(defaults, modeId) {
  const disposables = [];
  const providers = [];
  const client = new WorkerManager(modeId, defaults);
  disposables.push(client);
  const worker = (...uris) => {
    return client.getLanguageServiceWorker(...uris);
  };
  const libFiles = new LibFiles(worker);
  function registerProviders() {
    const { modeConfiguration } = defaults;
    disposeAll(providers);
    if (modeConfiguration.completionItems) {
      providers.push(monaco_editor_core_exports.languages.registerCompletionItemProvider(modeId, new SuggestAdapter(worker)));
    }
    if (modeConfiguration.signatureHelp) {
      providers.push(monaco_editor_core_exports.languages.registerSignatureHelpProvider(modeId, new SignatureHelpAdapter(worker)));
    }
    if (modeConfiguration.hovers) {
      providers.push(monaco_editor_core_exports.languages.registerHoverProvider(modeId, new QuickInfoAdapter(worker)));
    }
    if (modeConfiguration.documentHighlights) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentHighlightProvider(modeId, new DocumentHighlightAdapter(worker)));
    }
    if (modeConfiguration.definitions) {
      providers.push(monaco_editor_core_exports.languages.registerDefinitionProvider(modeId, new DefinitionAdapter(libFiles, worker)));
    }
    if (modeConfiguration.references) {
      providers.push(monaco_editor_core_exports.languages.registerReferenceProvider(modeId, new ReferenceAdapter(libFiles, worker)));
    }
    if (modeConfiguration.documentSymbols) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentSymbolProvider(modeId, new OutlineAdapter(worker)));
    }
    if (modeConfiguration.rename) {
      providers.push(monaco_editor_core_exports.languages.registerRenameProvider(modeId, new RenameAdapter(libFiles, worker)));
    }
    if (modeConfiguration.documentRangeFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerDocumentRangeFormattingEditProvider(modeId, new FormatAdapter(worker)));
    }
    if (modeConfiguration.onTypeFormattingEdits) {
      providers.push(monaco_editor_core_exports.languages.registerOnTypeFormattingEditProvider(modeId, new FormatOnTypeAdapter(worker)));
    }
    if (modeConfiguration.codeActions) {
      providers.push(monaco_editor_core_exports.languages.registerCodeActionProvider(modeId, new CodeActionAdaptor(worker)));
    }
    if (modeConfiguration.inlayHints) {
      providers.push(monaco_editor_core_exports.languages.registerInlayHintsProvider(modeId, new InlayHintsAdapter(worker)));
    }
    if (modeConfiguration.diagnostics) {
      providers.push(new DiagnosticsAdapter(libFiles, defaults, modeId, worker));
    }
  }
  registerProviders();
  disposables.push(asDisposable(providers));
  return worker;
}
function asDisposable(disposables) {
  return { dispose: () => disposeAll(disposables) };
}
function disposeAll(disposables) {
  while (disposables.length) {
    disposables.pop().dispose();
  }
}



/***/ })

}]);