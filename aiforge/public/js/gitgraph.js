"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["gitgraph"],{

/***/ "./web_src/js/vendor/gitgraph.js":
/*!***************************************!*\
  !*** ./web_src/js/vendor/gitgraph.js ***!
  \***************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": function() { return /* binding */ gitGraph; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_fill__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.fill */ "./node_modules/core-js/modules/es.array.fill.js");
/* harmony import */ var core_js_modules_es_array_fill__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_fill__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.splice */ "./node_modules/core-js/modules/es.array.splice.js");
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_5__);







/* This is a customized version of https://github.com/bluef/gitgraph.js/blob/master/gitgraph.js
   Changes include conversion to ES6 and linting fixes */

/*
 * @license magnet:?xt=urn:btih:c80d50af7d3db9be66a4d0a86db0286e4fd33292&dn=bsd-3-clause.txt BSD 3-Clause
 * Copyright (c) 2011, Terrence Lee <kill889@gmail.com>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *     * Neither the name of the <organization> nor the
 *       names of its contributors may be used to endorse or promote products
 *       derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
function gitGraph(canvas, rawGraphList, config) {
  if (!canvas.getContext) {
    return;
  }

  if (typeof config === 'undefined') {
    config = {
      unitSize: 20,
      lineWidth: 3,
      nodeRadius: 4
    };
  }

  var flows = [];
  var graphList = [];
  var ctx = canvas.getContext('2d');
  var devicePixelRatio = window.devicePixelRatio || 1;
  var backingStoreRatio = ctx.webkitBackingStorePixelRatio || ctx.mozBackingStorePixelRatio || ctx.msBackingStorePixelRatio || ctx.oBackingStorePixelRatio || ctx.backingStorePixelRatio || 1;
  var ratio = devicePixelRatio / backingStoreRatio;

  var init = function init() {
    var maxWidth = 0;
    var i;
    var l = rawGraphList.length;
    var row;
    var midStr;

    for (i = 0; i < l; i++) {
      midStr = rawGraphList[i].replace(/\s+/g, ' ').replace(/^\s+|\s+$/g, '');
      maxWidth = Math.max(midStr.replace(/(_|\s)/g, '').length, maxWidth);
      row = midStr.split('');
      graphList.unshift(row);
    }

    var width = maxWidth * config.unitSize;
    var height = graphList.length * config.unitSize;
    canvas.width = width * ratio;
    canvas.height = height * ratio;
    canvas.style.width = "".concat(width, "px");
    canvas.style.height = "".concat(height, "px");
    ctx.lineWidth = config.lineWidth;
    ctx.lineJoin = 'round';
    ctx.lineCap = 'round';
    ctx.scale(ratio, ratio);
  };

  var genRandomStr = function genRandomStr() {
    var chars = '0123456789ABCDEF';
    var stringLength = 6;
    var randomString = '',
        rnum,
        i;

    for (i = 0; i < stringLength; i++) {
      rnum = Math.floor(Math.random() * chars.length);
      randomString += chars.substring(rnum, rnum + 1);
    }

    return randomString;
  };

  var findFlow = function findFlow(id) {
    var i = flows.length;

    while (i-- && flows[i].id !== id) {
      ;
    }

    return i;
  };

  var findColomn = function findColomn(symbol, row) {
    var i = row.length;

    while (i-- && row[i] !== symbol) {
      ;
    }

    return i;
  };

  var findBranchOut = function findBranchOut(row) {
    if (!row) {
      return -1;
    }

    var i = row.length;

    while (i-- && !(row[i - 1] && row[i] === '/' && row[i - 1] === '|') && !(row[i - 2] && row[i] === '_' && row[i - 2] === '|')) {
      ;
    }

    return i;
  };

  var findLineBreak = function findLineBreak(row) {
    if (!row) {
      return -1;
    }

    var i = row.length;

    while (i-- && !(row[i - 1] && row[i - 2] && row[i] === ' ' && row[i - 1] === '|' && row[i - 2] === '_')) {
      ;
    }

    return i;
  };

  var genNewFlow = function genNewFlow() {
    var newId;

    do {
      newId = genRandomStr();
    } while (findFlow(newId) !== -1);

    return {
      id: newId,
      color: "#".concat(newId)
    };
  }; // Draw methods


  var drawLine = function drawLine(moveX, moveY, lineX, lineY, color) {
    ctx.strokeStyle = color;
    ctx.beginPath();
    ctx.moveTo(moveX, moveY);
    ctx.lineTo(lineX, lineY);
    ctx.stroke();
  };

  var drawLineRight = function drawLineRight(x, y, color) {
    drawLine(x, y + config.unitSize / 2, x + config.unitSize, y + config.unitSize / 2, color);
  };

  var drawLineUp = function drawLineUp(x, y, color) {
    drawLine(x, y + config.unitSize / 2, x, y - config.unitSize / 2, color);
  };

  var drawNode = function drawNode(x, y, color) {
    ctx.strokeStyle = color;
    drawLineUp(x, y, color);
    ctx.beginPath();
    ctx.arc(x, y, config.nodeRadius, 0, Math.PI * 2, true);
    ctx.fill();
  };

  var drawLineIn = function drawLineIn(x, y, color) {
    drawLine(x + config.unitSize, y + config.unitSize / 2, x, y - config.unitSize / 2, color);
  };

  var drawLineOut = function drawLineOut(x, y, color) {
    drawLine(x, y + config.unitSize / 2, x + config.unitSize, y - config.unitSize / 2, color);
  };

  var draw = function draw(graphList) {
    var colomn,
        colomnIndex,
        prevColomn,
        condenseIndex,
        breakIndex = -1;
    var x, y;
    var color;
    var nodePos;
    var tempFlow;
    var prevRowLength = 0;
    var flowSwapPos = -1;
    var lastLinePos;
    var i, l;
    var condenseCurrentLength,
        condensePrevLength = 0;
    var inlineIntersect = false; // initiate color array for first row

    for (i = 0, l = graphList[0].length; i < l; i++) {
      if (graphList[0][i] !== '_' && graphList[0][i] !== ' ') {
        flows.push(genNewFlow());
      }
    }

    y = canvas.height / ratio - config.unitSize * 0.5; // iterate

    for (i = 0, l = graphList.length; i < l; i++) {
      x = config.unitSize * 0.5;
      var currentRow = graphList[i];
      var nextRow = graphList[i + 1];
      var prevRow = graphList[i - 1];
      flowSwapPos = -1;
      condenseCurrentLength = currentRow.filter(function (val) {
        return val !== ' ' && val !== '_';
      }).length; // pre process begin
      // use last row for analysing

      if (prevRow) {
        if (!inlineIntersect) {
          // intersect might happen
          for (colomnIndex = 0; colomnIndex < prevRowLength; colomnIndex++) {
            if (prevRow[colomnIndex + 1] && prevRow[colomnIndex] === '/' && prevRow[colomnIndex + 1] === '|' || prevRow[colomnIndex] === '_' && prevRow[colomnIndex + 1] === '|' && prevRow[colomnIndex + 2] === '/') {
              flowSwapPos = colomnIndex; // swap two flow

              tempFlow = {
                id: flows[flowSwapPos].id,
                color: flows[flowSwapPos].color
              };
              flows[flowSwapPos].id = flows[flowSwapPos + 1].id;
              flows[flowSwapPos].color = flows[flowSwapPos + 1].color;
              flows[flowSwapPos + 1].id = tempFlow.id;
              flows[flowSwapPos + 1].color = tempFlow.color;
            }
          }
        }

        if (condensePrevLength < condenseCurrentLength && (nodePos = findColomn('*', currentRow)) !== -1 // eslint-disable-line no-cond-assign
        && findColomn('_', currentRow) === -1) {
          flows.splice(nodePos - 1, 0, genNewFlow());
        }

        if (prevRowLength > currentRow.length && (nodePos = findColomn('*', prevRow)) !== -1) {
          // eslint-disable-line no-cond-assign
          if (findColomn('_', currentRow) === -1 && findColomn('/', currentRow) === -1 && findColomn('\\', currentRow) === -1) {
            flows.splice(nodePos + 1, 1);
          }
        }
      } // done with the previous row


      prevRowLength = currentRow.length; // store for next round

      colomnIndex = 0; // reset index

      condenseIndex = 0;
      condensePrevLength = 0;
      breakIndex = -1; // reset break index

      while (colomnIndex < currentRow.length) {
        colomn = currentRow[colomnIndex];

        if (colomn !== ' ' && colomn !== '_') {
          ++condensePrevLength;
        } // check and fix line break in next row


        if (colomn === '/' && currentRow[colomnIndex - 1] && currentRow[colomnIndex - 1] === '|') {
          /* eslint-disable-next-line */
          if ((breakIndex = findLineBreak(nextRow)) !== -1) {
            nextRow.splice(breakIndex, 1);
          }
        } // if line break found replace all '/' with '|' after breakIndex in previous row


        if (breakIndex !== -1 && colomn === '/' && colomnIndex > breakIndex) {
          currentRow[colomnIndex] = '|';
          colomn = '|';
        }

        if (colomn === ' ' && currentRow[colomnIndex + 1] && currentRow[colomnIndex + 1] === '_' && currentRow[colomnIndex - 1] && currentRow[colomnIndex - 1] === '|') {
          currentRow.splice(colomnIndex, 1);
          currentRow[colomnIndex] = '/';
          colomn = '/';
        } // create new flow only when no intersect happened


        if (flowSwapPos === -1 && colomn === '/' && currentRow[colomnIndex - 1] && currentRow[colomnIndex - 1] === '|') {
          flows.splice(condenseIndex, 0, genNewFlow());
        } // change \ and / to | when it's in the last position of the whole row


        if (colomn === '/' || colomn === '\\') {
          if (!(colomn === '/' && findBranchOut(nextRow) === -1)) {
            /* eslint-disable-next-line */
            if ((lastLinePos = Math.max(findColomn('|', currentRow), findColomn('*', currentRow))) !== -1 && lastLinePos < colomnIndex - 1) {
              while (currentRow[++lastLinePos] === ' ') {
                ;
              }

              if (lastLinePos === colomnIndex) {
                currentRow[colomnIndex] = '|';
              }
            }
          }
        }

        if (colomn === '*' && prevRow && prevRow[condenseIndex + 1] === '\\') {
          flows.splice(condenseIndex + 1, 1);
        }

        if (colomn !== ' ') {
          ++condenseIndex;
        }

        ++colomnIndex;
      }

      condenseCurrentLength = currentRow.filter(function (val) {
        return val !== ' ' && val !== '_';
      }).length; // do some clean up

      if (flows.length > condenseCurrentLength) {
        flows.splice(condenseCurrentLength, flows.length - condenseCurrentLength);
      }

      colomnIndex = 0; // a little inline analysis and draw process

      while (colomnIndex < currentRow.length) {
        colomn = currentRow[colomnIndex];
        prevColomn = currentRow[colomnIndex - 1];

        if (currentRow[colomnIndex] === ' ') {
          currentRow.splice(colomnIndex, 1);
          x += config.unitSize;
          continue;
        } // inline interset


        if ((colomn === '_' || colomn === '/') && currentRow[colomnIndex - 1] === '|' && currentRow[colomnIndex - 2] === '_') {
          inlineIntersect = true;
          tempFlow = flows.splice(colomnIndex - 2, 1)[0];
          flows.splice(colomnIndex - 1, 0, tempFlow);
          currentRow.splice(colomnIndex - 2, 1);
          colomnIndex -= 1;
        } else {
          inlineIntersect = false;
        }

        color = flows[colomnIndex].color;

        switch (colomn) {
          case '_':
            drawLineRight(x, y, color);
            x += config.unitSize;
            break;

          case '*':
            drawNode(x, y, color);
            break;

          case '|':
            drawLineUp(x, y, color);
            break;

          case '/':
            if (prevColomn && (prevColomn === '/' || prevColomn === ' ')) {
              x -= config.unitSize;
            }

            drawLineOut(x, y, color);
            x += config.unitSize;
            break;

          case '\\':
            drawLineIn(x, y, color);
            break;
        }

        ++colomnIndex;
      }

      y -= config.unitSize;
    }
  };

  init();
  draw(graphList);
} // @end-license

/***/ })

}]);