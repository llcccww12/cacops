import { BoundingBox, Layout } from '~/utils/treelayout';
import { getListValueWithKey, transFileSize, renderSpecObject } from '~/utils';
import { MODEL_ENGINES } from '~/const';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { i18n, lang } from '~/langs';
import { getPromoteData } from '~/apis/modules/common';
function ModelGraph() {
  this.options = {
    horizontalGap: 120,
    verticalGap: 40,
    parentModelNodeWidth: 100,
    parentModelNodeHeight: 60,
    currentModelNodeWidth: 100,
    currentModelNodeHeight: 100,
    deriveModelNodeWidth: 100,
    deriveModelNodeHeight: 60,
    repoNodeWidth: 140,
    repoNodeHeight: 40,
    collapsedIconMarginLeft: 20,
  };

  this.$el = null;
  this.$layout = null;
  this.viewWinWidth = 0;
  this.viewWinHeight = 0;
  this.layoutWidth = 0;
  this.layoutHeight = 0;
  this._layoutWidth = 0;
  this._layoutHeight = 0;
  this.initScale = 1;

  this.drawOffsetLeft = 0;
  this.drawOffsetTop = 0;

  this.treeData = null;
  this.nodeMap = {};
  return this;
}

ModelGraph.prototype.init = function (element, data, options) {
  if (!element) return;
  if (!['absolute', 'relative', 'fixed'].includes(element.style.position)) {
    element.style.position = 'relative';
    element.style.overflow = 'hidden';
  }
  this.$el = element;
  Object.assign(this.options, options);
  this.treeData = data;
  console.log("init",element,data,options)
  this.renderTree();
  this.toolBarInit();
  this.eventInit();
};

ModelGraph.prototype.walkTree = function (pNode, lvl) {
  if (pNode.type == 'model') {
    if (pNode.isParent) {
      pNode.width = this.options.parentModelNodeHeight;
      pNode.height = this.options.parentModelNodeWidth;
    }
    if (pNode.isCurrent) {
      pNode.width = this.options.currentModelNodeHeight;
      pNode.height = this.options.currentModelNodeWidth;
    }
    if (pNode.isDerive) {
      pNode.width = this.options.deriveModelNodeHeight;
      pNode.height = this.options.deriveModelNodeWidth;
    }
  }
  if (pNode.type == 'task') {
    pNode.width = this.options.repoNodeHeight;
    pNode.height = this.options.repoNodeWidth;
  }
  pNode.lvl = lvl;
  const children = pNode.children || [];
  for (let i = 0, iLen = children.length; i < iLen; i++) {
    const child = children[i];
    child.pNode = pNode;
    this.walkTree(child, lvl + 1);
  }
  pNode._children = pNode._children || pNode.children;
  if (pNode.isCollapsed) {
    pNode.children = undefined;
  } else {
    pNode.children = pNode._children;
  }
  if (!pNode._id) {
    pNode._id = `${lvl}-${Math.random()}`;
  }
  this.nodeMap[pNode._id] = pNode;
};

ModelGraph.prototype.renderTree = function () {
  const bb = new BoundingBox(this.options.verticalGap, this.options.horizontalGap);
  const layout = new Layout(bb);
  this.walkTree(this.treeData, 0);
  const { result, boundingBox } = layout.layout(this.treeData);
  const viewWinEl = this.$el;
  this.viewWinWidth = viewWinEl.clientWidth;
  this.viewWinHeight = viewWinEl.clientHeight;
  const layoutWidth = boundingBox.bottom;
  const layoutHeight = boundingBox.right;
  this._layoutWidth = layoutWidth;
  this._layoutHeight = layoutHeight;
  if (!this.$layout) {
    this.layoutWidth = Math.max(boundingBox.bottom, boundingBox.right, 1000);
    this.layoutHeight = Math.max(boundingBox.bottom, boundingBox.right, 1000);
    const layoutEl = document.createElement('div');
    layoutEl.classList.add('_tree-layout');
    layoutEl.classList.add('model-graph-layout');
    this.$el.append(layoutEl);
    this.$layout = layoutEl;
    this.$layout.style.width = this.layoutWidth * 4 + 'px';
    this.$layout.style.height = this.layoutHeight * 4 + 'px';
    this.$layout.style.left = - this.layoutWidth * 4 / 2 + this.viewWinWidth / 2 + 'px';
    this.$layout.style.top = - this.layoutHeight * 4 / 2 + this.viewWinHeight / 2 + 'px';
    const scale = Math.min(Math.min((this.viewWinWidth - 100) / layoutWidth, (this.viewWinHeight - 100) / layoutHeight), 1);
    this.initScale = scale;
    this.$layout.style.scale = scale;
    this.$layout.innerHTML = '<svg class="_tree-layout-svg"><g></g></svg>';
    this.drawOffsetLeft = this.layoutWidth * 1.5;
    this.drawOffsetTop = this.layoutHeight * 1.5 + (this.layoutHeight - layoutHeight) / 2 - this.options.verticalGap / 4;
  } else {
    this.drawOffsetLeft = this.layoutWidth * 1.5;
    this.drawOffsetTop = this.layoutHeight * 1.5 + (this.layoutHeight - layoutHeight) / 2 - this.options.verticalGap / 4;
  }
  this.drawTreeNode(this.treeData);
  this.popupInit();
};

ModelGraph.prototype.drawTreeNode = function (treeData, setPoint) {
  console.log("treeData",treeData)
  const children = treeData.children || [];
  let divEl = this.$layout.querySelectorAll(`._tree-node[data-id="${treeData._id}"]`)[0];
  console.log("divEl",divEl)
  if (!divEl) {
    divEl = document.createElement('div');
    divEl.classList = ['_tree-node'];
    let showName = treeData.name;
    if (((treeData.type == 'model' && treeData.isDerive) || treeData.type == 'task')
      && !treeData.isCanOper) {
      showName = `${treeData.name.slice(0, 1)}******`;
      treeData.link = '';
    }
    let innerHTML = treeData.link
      ? `<div class="_tree-node-block"><a class="name" target="_blank" href="${treeData.link}">${showName}</a></div>`
      : `<div class="_tree-node-block"><div class="name">${showName}</div></div>`;
    if (treeData.type == 'model') {
      divEl.classList.add('model');
      if (treeData.isParent) {
        divEl.classList.add('model-parent');
        innerHTML += `<div class="descr">${i18n.t('modelManage.parentModel')}</div>`;
      }
      if (treeData.isCurrent) {
        divEl.classList.add('model-current');
        innerHTML += `<div class="descr">${i18n.t('modelManage.currentModel')}</div>`;
      }
      if (treeData.isDerive) {
        divEl.classList.add('model-derive');
        if (treeData.isPrivate) {
          divEl.classList.add('private');
          innerHTML += `<div class="lock-c"><i class=ri-lock-line></i></div>`;
        }
      }
      if (treeData.pNode) {
        innerHTML += `<div class="creator" style="margin-top:-14px"><a target="_blank" href="/${treeData.creator}">${treeData.creator}</a></div>`;
      }
    }
    if (treeData.type == 'task') {
      divEl.classList.add('model');
      divEl.classList.add('repo');
      if (treeData.isPrivate) {
        divEl.classList.add('private');
        innerHTML += `<div class="lock-c"><i class=ri-lock-line></i></div>`;
      }
    }
    if (children.length || (treeData._children || []).length) {
      innerHTML += `<div class="node-expand-c">
      <div class="node-expand-line"></div>
      <div class="node-expand-icon"></div>
    </div>`;
    }
    divEl.style.width = treeData.height + 'px';
    divEl.style.height = treeData.width + 'px';
    divEl.style.zIndex = 10000 - treeData.lvl;
    divEl.dataset.id = treeData._id;
    divEl.innerHTML = innerHTML;
    console.log("innerHTML",innerHTML)
    this.$layout.append(divEl);
  }
  if (divEl.querySelector('div.node-expand-icon')) {
    divEl.querySelector('div.node-expand-icon').innerHTML = treeData.isCollapsed ? '+' : '-';
  }
  if (setPoint) {
    divEl.style.left = setPoint.x - treeData.height / 2 + 'px';
    divEl.style.top = setPoint.y - treeData.width / 2 + 'px';
    divEl.style.scale = 0;
  } else {
    divEl.style.left = (treeData.y + this.drawOffsetLeft) + 'px';
    divEl.style.top = (treeData.x + this.drawOffsetTop) + 'px';
    divEl.style.scale = 1;
    const creator = divEl.querySelector('div.creator');
    if (creator && treeData.pNode) {
      const marginTop = treeData.pNode.x + treeData.pNode.width / 2 >= treeData.x + treeData.width / 2 ? -14 : 14;
      creator.style.marginTop = marginTop + 'px';
    }
  }
  for (let i = 0, iLen = children.length; i < iLen; i++) {
    this.drawTreeNode(children[i], setPoint);
    this.drawLine(treeData, children[i], setPoint);
  }
};

ModelGraph.prototype.drawLine = function (pNode, node, setPoint) {
  const svgPathContainer = this.$layout.querySelector('._tree-layout-svg g');
  const node1 = {
    w: pNode.height,
    h: pNode.width,
    l: pNode.y + this.drawOffsetLeft,
    t: pNode.x + this.drawOffsetTop
  }
  const node2 = {
    w: node.height,
    h: node.width,
    l: node.y + this.drawOffsetLeft,
    t: node.x + this.drawOffsetTop
  }
  const point1 = {
    x: node1.l + node1.w + this.options.collapsedIconMarginLeft,
    y: node1.t + node1.h / 2
  };
  const point2 = {
    x: node2.l,
    y: node2.t + node2.h / 2
  }
  let d = `M${point1.x} ${point1.y} C${(point1.x + point2.x) / 2} ${point1.y} ${(point1.x + point2.x) / 2} ${point2.y} ${point2.x} ${point2.y}`;
  let path = svgPathContainer.querySelectorAll(`path[data-id="${node._id}"]`)[0];
  if (!path) {
    path = document.createElementNS('http://www.w3.org/2000/svg', 'path');
    path.setAttribute('fill', 'none');
    path.setAttribute('stroke', 'rgb(122, 184, 251)');
    path.setAttribute('stroke-width', 1);
    path.setAttribute('id', node._id);
    path.dataset.id = node._id;
    svgPathContainer.append(path);
  }
  if (setPoint) {
    d = `M${setPoint.x} ${setPoint.y} C${setPoint.x} ${setPoint.y} ${setPoint.x} ${setPoint.y} ${setPoint.x} ${setPoint.y}`;
  }
  path.setAttribute('d', d);
};

ModelGraph.prototype.removeNode = function (nodeList) {
  const layout = this.$layout;
  for (let i = 0, iLen = nodeList.length; i < iLen; i++) {
    const curNode = nodeList[i];
    const id = curNode._id;
    const nodes = this.$layout.querySelectorAll(`[data-id="${id}"]`);
    nodes.forEach(item => item.remove());
    const children = curNode.children || [];
    this.removeNode(children);
  }
};

ModelGraph.prototype.showModelNodeInfo = async function (nodeEl, nodeData, showOrHide) {
  if (!nodeData.isCanOper) return;
  const self = this;
  const elClass = '_model-info-' + nodeData._id.toString().replace('.', '-');
  let showInfoEl = document.querySelector('.' + elClass);
  console.log("showModelNodeInfo",nodeData,nodeData.type)
  if (showOrHide == 'show') {
    if (showInfoEl) {
      nodeData.delayHideTimer && clearTimeout(nodeData.delayHideTimer);
      return;
    }
    let innerHTML = '';
    showInfoEl = document.createElement('div');
    showInfoEl.classList.add('_model-info');
    showInfoEl.classList.add(elClass);
    if (nodeData.type=='model') {
      let models = [];
      if (nodeData.isParent) {
        // models = nodeData.Models4Parent || [];
        nodeData.Models4Parent.forEach((model) => {
          models.push(model.Model)
        })
        console.log("Paranetmodels",models)
      } else {
        models = nodeData.Model ? [nodeData.Model] : [];
        console.log("models",models)
      }
      
      
      for (let i = 0, iLen = models.length; i < iLen; i++) {
        const model = models[i];
        let modelObj = { ...model };
        modelObj = {
          ...modelObj,
          engineName: getListValueWithKey(MODEL_ENGINES, model.engine),
          modelSize: transFileSize(model.size),
          createTimeStr: formatDate(new Date(model.created_unix * 1000), 'yyyy-MM-dd HH:mm:ss'),
          isPrivateStr: model.is_private ? i18n.t('modelManage.modelAccessPrivate') : i18n.t('modelManage.modelAccessPublic'),
        };
        innerHTML += `
        ${i != 0 ? '<div class="line"></div>' : ''}
        <div class="model-block">
          <div class="tit">${i18n.t('modelManage.modelInfo')}</div>
          <div class="row">
            <div>${i18n.t('modelManage.modelName')}：</div>
            <div><a class="name" target="_blank" href="${modelObj.link}">${lang === 'zh-CN' ? modelObj.alias : modelObj.name}</a></div>
          </div>
          <div class="row">
            <div>${i18n.t('modelManage.modelEngine')}：</div>
            <div>${modelObj.engineName}</div>
          </div>
          <div class="row">
            <div>${i18n.t('modelManage.modelSize')}：</div>
            <div>${modelObj.modelSize}</div>
          </div>
          <div class="row">
            <div>${i18n.t('modelManage.createTime')}：</div>
            <div>${modelObj.createTimeStr}</div>
          </div>
          <div class="row">
            <div>${i18n.t('modelManage.modelAccess')}：</div>
            <div>${modelObj.isPrivateStr}</div>
          </div>
        </div>`;
      }
    } else if (nodeData.type == 'task') {
      let taskObj = { ...nodeData.Task };
      taskObj.sepcStr = renderSpecObject(taskObj.spec, false).specStr;
      innerHTML += `<div class="model-block">
      <div class="tit">${i18n.t('modelManage.trainingInfo')}</div>`
      if (taskObj.dataset_list.length) {
        innerHTML += `
        <div class="row">
          <div>${i18n.t('datasetObj.dataset_name1')}：</div>
          <div>
        `
        taskObj.dataset_list.forEach((dataset, index) => {
          innerHTML += !dataset.is_delete
              ? `<a target="_blank" href="/datasets/detail/${dataset.owner_name}/${dataset.dataset_name}">${dataset.dataset_alias}<span>, </span></a>`
              : `<div>${dataset.dataset_name} (${i18n.t('cloudbrainObj.fileWasDeleted')})<span>, </span></div>`
        })
        innerHTML += `
          </div></div>`
      } else {
        innerHTML += ` 
        <div class="row">
          <div>${i18n.t('dataset')}：</div>
          <div>--</div>
        </div>
        `
      }
      
      if (taskObj.pretrain_model_list.length) {
        innerHTML += `
        <div class="row">
          <div>${i18n.t('modelManage.modelName')}：</div>
          <div>`
        taskObj.pretrain_model_list.forEach((model, index) => {
          innerHTML += !model.is_delete
            ? `<a target="_blank" href="/models/detail/${model.owner_name}/${model.name}">${lang === 'zh-CN' ? modelObj.alias : modelObj.name}<span>, </span></a>`
            : `<div>${lang === 'zh-CN' ? modelObj.alias : modelObj.name} (${i18n.t('cloudbrainObj.fileWasDeleted')})<span>, </span></div>`
        })
        innerHTML += `
          </div></div>`
      } else {
        innerHTML += ` 
        <div class="row">
          <div>${i18n.t('modelManage.modelName')}：</div>
          <div>--</div>
        </div>
        `
      }
      innerHTML += `
          <div class="row">
            <div>${i18n.t('cloudbrainObj.image')}：</div>
            <div>${taskObj.image_name}</div>
          </div>
          <div class="row">
            <div>代码仓：</div>
            <div><a target="_blank" href="/${taskObj.repo_owner_name}/${taskObj.repo_name }">${taskObj.repo_owner_name}/${taskObj.repo_name }</a></div>
          </div>
          <div class="row">
            <div>分支：</div>
            <div>${taskObj.branch_name}</div>
          </div>
          <div class="row">
            <div>${i18n.t('trainDuration')}：</div>
            <div>${taskObj.formatted_duration}</div>
          </div>
          <div class="row">
            <div>${i18n.t('cloudbrainObj.resourceSpec')}：</div>
            <div>${taskObj.sepcStr}</div>
          </div>
        </div>
      `
    }
    console.log(" showInfoEl.innerHTML", innerHTML)
    showInfoEl.innerHTML = innerHTML;
    
    const posInfo = nodeEl.getBoundingClientRect();
    const winW = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
    showInfoEl.style.top = Math.max(posInfo.top - 134 + 6, 20) + 'px';
    showInfoEl.style.left = Math.min(posInfo.left - 6 + posInfo.width, winW - 250 - 20) + 'px';
    document.querySelector('body').append(showInfoEl);
    const mouseEnter = function () {
      self.showModelNodeInfo(nodeEl, nodeData, 'show');
    };
    const mouseLeave = function () {
      self.showModelNodeInfo(nodeEl, nodeData, 'hide');
    };
    showInfoEl.addEventListener('mouseenter', mouseEnter);
    showInfoEl.addEventListener('mouseleave', mouseLeave);
  } else {
    nodeData.delayHideTimer && clearTimeout(nodeData.delayHideTimer);
    nodeData.delayHideTimer = setTimeout(() => {
      if (showInfoEl) {
        showInfoEl.remove();
      }
    }, 200);
  }
};

ModelGraph.prototype.eventInit = function () {
  if (this.$layout) {
    this._drag(this.$layout);
    this._scale(this.$layout);
    const self = this;
    this.$layout.addEventListener('click', function (evt) {
      if (evt.target.classList.contains('node-expand-icon')) {
        const nodeEl = evt.target.parentElement.parentElement;
        const id = nodeEl.dataset.id;
        const curNode = self.nodeMap[id];
        const isCollapsed = curNode.isCollapsed;
        if (isCollapsed) {
          const n = {
            w: curNode.height,
            h: curNode.width,
            l: curNode.y + self.drawOffsetLeft,
            t: curNode.x + self.drawOffsetTop
          }
          const point = {
            x: n.l + n.w + self.options.collapsedIconMarginLeft,
            y: n.t + n.h / 2
          };
          const children = curNode._children || [];
          for (let i = 0, iLen = children.length; i < iLen; i++) {
            self.drawTreeNode(children[i], point);
            self.drawLine(curNode, children[i], point);
          }
          curNode.isCollapsed = !curNode.isCollapsed;
          self.renderTree();
        } else {
          const children = curNode.children || [];
          curNode.isCollapsed = !curNode.isCollapsed;
          self.renderTree();
          const n = {
            w: curNode.height,
            h: curNode.width,
            l: curNode.y + self.drawOffsetLeft,
            t: curNode.x + self.drawOffsetTop
          }
          const point = {
            x: n.l + n.w + self.options.collapsedIconMarginLeft,
            y: n.t + n.h / 2
          };
          for (let i = 0, iLen = children.length; i < iLen; i++) {
            self.drawTreeNode(children[i], point);
            self.drawLine(curNode, children[i], point);
          }
          setTimeout(function () {
            self.removeNode(children);
          }, 260);
        }
      }
    });
  }
};

ModelGraph.prototype.popupInit = function () {
  const modelsEls = this.$layout.querySelectorAll('._tree-node.model ._tree-node-block');
  const self = this;
  console.log("modelsEls",modelsEls)
  const mouseEnterHandler = function (evt) {
    const nodeEl = evt.target.parentElement;
    const id = nodeEl.dataset.id;
    const curNode = self.nodeMap[id];
    self.showModelNodeInfo(nodeEl, curNode, 'show');
  };
  const mouseLeaveHandler = function (evt) {
    const nodeEl = evt.target.parentElement;
    const id = nodeEl.dataset.id;
    const curNode = self.nodeMap[id];
    self.showModelNodeInfo(nodeEl, curNode, 'hide');
  };

  for (let i = 0, iLen = modelsEls.length; i < iLen; i++) {
    const modelEl = modelsEls[i];
    modelEl.removeEventListener('mouseenter', mouseEnterHandler);
    modelEl.removeEventListener('mouseleave', mouseLeaveHandler);
    modelEl.addEventListener('mouseenter', mouseEnterHandler);
    modelEl.addEventListener('mouseleave', mouseLeaveHandler);
  }
}

ModelGraph.prototype.toolBarInit = function () {
  const toolBarEl = document.createElement('div');
  toolBarEl.classList.add('_tree-toolbar');
  toolBarEl.innerHTML = `<div class="legend-c">
  <div class="legend"><div class="legend-icon legend-isprarent"></div><span class="legend-name">${i18n.t('modelManage.parentModel')}</span></div>
  <div class="legend"><div class="legend-icon legend-iscurrent"></div><span class="legend-name">${i18n.t('modelManage.currentModel')}</span></div>
  <div class="legend"><div class="legend-icon legend-isderive"></div><span class="legend-name">${i18n.t('modelManage.publicDerivedModel')}</span></div>  
  <div class="legend"><div class="legend-icon legend-isderive private">
    <div class="lock-c"><i class="ri-lock-line"></i></div>
  </div><span class="legend-name">${i18n.t('modelManage.privateDerivedModel')}</span></div>
  <div class="legend"><div class="legend-icon legend-isrepo"></div><span class="legend-name">${i18n.t('modelManage.trainTask')}</span></div>
  
  </div>`;
  this.$el.append(toolBarEl);
};

ModelGraph.prototype._drag = function (el) {
  var self = this;
  el.onmousedown = mouseDown;
  function mouseDown(event) {
    event = event || window.event;
    var disX = event.clientX - el.offsetLeft;
    var disY = event.clientY - el.offsetTop;
    if (event.target == this || event.target.classList.contains('_tree-layout-svg')) {
      var oldCursor = el.style.cursor;
      el.style.cursor = 'move';
      document.onmousemove = function (event) {
        event = event || window.event;
        MouseMove(event, disX, disY);
      }
      document.onmouseup = function () {
        document.onmousemove = null;
        el.style.cursor = oldCursor;
      }
    }
  }
  function MouseMove(event, x1, y1) {
    event = event || window.event;
    var l = event.clientX - x1,
      t = event.clientY - y1,
      winW = document.documentElement.clientWidth || document.body.clientWidth,
      winH = document.documentElement.clientHeight || document.body.clientHeight,
      maxW = winW - el.offsetWidth,
      maxH = winH - el.offsetHeight;
    // if (l < 0) {
    //   l = 0;
    // } else if (l > maxW - 15) {
    //   l = maxW - 15;
    // }
    // if (t < 0) {
    //   t = 0;
    // } else if (t > maxH - 15) {
    //   t = maxH - 15;
    // }
    el.style.left = l + "px";
    el.style.top = t + "px";
  }
}

ModelGraph.prototype._scale = function (el) {
  const self = this;
  el.onmousewheel = function (event) {
    var scale = parseFloat(el.style.scale || 1);
    if (event.wheelDelta >= 0) {
      scale *= 1.1;
    } else {
      scale /= 1.1;
    }
    if (scale >= self.initScale / 2 && scale < 1.5) {
      el.style.scale = scale;
    }
    event.stopPropagation();
    event.preventDefault();
    event.returnValue = false;
  }
}

export { ModelGraph }
