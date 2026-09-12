/**
 * 设置元素的粘性定位样式
 * @param {HTMLElement} element - 目标元素
 * @param {Object} options - 定位选项
 * @param {string|number} options.top - 顶部位置
 * @param {string|number} [options.left] - 左侧位置
 * @param {string|number} [options.right] - 右侧位置
 * @param {number} [options.zIndex=10] - 层级
 */
const setStickyStyle = (element, { top, left, right, zIndex = 10 }) => {
  element.style.position = 'sticky';
  element.style.zIndex = zIndex;

  if (top !== undefined) element.style.top = typeof top === 'number' ? `${top}px` : top;
  if (left !== undefined) element.style.left = typeof left === 'number' ? `${left}px` : left;
  if (right !== undefined) element.style.right = typeof right === 'number' ? `${right}px` : right;
};

/**
 * 计算顶部距离
 * @param {*} value - 指令的值
 * @returns {string} - 顶部距离字符串
 */
const calculateTop = (value) => {
  // 如果是数字，直接返回px单位
  if (!isNaN(Number(value))) {
    return `${value}px`;
  }

  // 如果是选择器，计算选择器元素的高度
  if (typeof value === 'string') {
    const element = document.querySelector(value);
    if (element) {
      const rect = element.getBoundingClientRect();
      return `${rect.top + rect.height}px`;
    }
  }

  return '0px';
};

/**
 * 处理固定列
 * @param {Array} fixedColumns - 固定列配置数组
 * @param {HTMLElement} headerWrapper - 表头包装元素
 * @param {string} top - 顶部距离
 * @param {string} direction - 方向: 'left' 或 'right'
 */
const handleFixedColumns = (fixedColumns, headerWrapper, top, direction = 'left') => {
  if (!fixedColumns || !fixedColumns.length) return;

  let accumulatedWidth = 0;
  const columns = direction === 'left' ? fixedColumns : [...fixedColumns].reverse();

  columns.forEach((column, index) => {
    const columnWidth = column.width || column.realWidth || column.minWidth;
    const cell = headerWrapper.querySelector(`th.${column.id}`);

    if (cell) {
      const styleOptions = { top, zIndex: 10 };

      if (direction === 'left') {
        styleOptions.left = accumulatedWidth;
        accumulatedWidth += columnWidth;
      } else {
        styleOptions.right = accumulatedWidth;
        accumulatedWidth += columnWidth;
      }

      setStickyStyle(cell, styleOptions);
    }
  });
};

/**
 * 查找表格组件实例
 * @param {Array} children - Vue组件子实例数组
 * @param {HTMLElement} targetEl - 目标DOM元素
 * @returns {Object|null} - 找到的表格实例
 */
const findTableComponent = (children, targetEl) => {
  if (!Array.isArray(children) || children.length === 0) {
    return null;
  }

  for (const child of children) {
    if (child.$el === targetEl) {
      return child;
    }

    if (child.$children?.length) {
      const found = findTableComponent(child.$children, targetEl);
      if (found) return found;
    }
  }

  return null;
};

/**
 * 修复固定列的隐藏类
 * @param {HTMLElement} headerWrapper - 表头包装元素
 */
const fixFixedColumnClasses = (headerWrapper) => {
  const hiddenThs = headerWrapper.querySelectorAll('th.is-hidden');
  hiddenThs.forEach(th => th.classList.remove('is-hidden'));
};

/**
 * 主函数：设置粘性表头
 * @param {HTMLElement} el - 目标元素
 * @param {Object} binding - 指令绑定对象
 * @param {Object} vnode - 虚拟节点
 */
const stickyThead = (el, binding, vnode) => {
  const top = calculateTop(binding.value);

  // 设置表格容器样式
  el.style.overflow = 'visible';

  // 获取表头包装器并设置粘性
  const headerWrapper = el.querySelector('.el-table__header-wrapper');
  if (!headerWrapper) return;

  setStickyStyle(headerWrapper, { top });
  fixFixedColumnClasses(headerWrapper);

  // 查找表格组件实例
  const table = findTableComponent(vnode.context.$children, el);
  if (!table) return;

  // 处理左右固定列
  handleFixedColumns(table.fixedColumns, headerWrapper, top, 'left');
  handleFixedColumns(table.rightFixedColumns, headerWrapper, top, 'right');
};

// 观察器管理类
class ObserverManager {
  constructor() {
    this.observers = new Map();
  }

  /**
   * 观察元素
   * @param {HTMLElement} element - 要观察的元素
   * @param {Function} callback - 回调函数
   */
  observe(element, callback) {
    if (!element || this.observers.has(element)) return;

    const observer = new ResizeObserver(callback);
    observer.observe(element);
    this.observers.set(element, observer);
  }

  /**
   * 停止观察元素
   * @param {HTMLElement} element - 要停止观察的元素
   */
  unobserve(element) {
    if (!element || !this.observers.has(element)) return;

    const observer = this.observers.get(element);
    observer.unobserve(element);
    observer.disconnect();
    this.observers.delete(element);
  }

  /**
   * 清空所有观察器
   */
  clear() {
    this.observers.forEach(observer => {
      observer.disconnect();
    });
    this.observers.clear();
  }
}

// 创建观察器管理器实例
const observerManager = new ObserverManager();

export default {
  inserted(el, binding, vnode) {
    // 观察表格本身的变化
    observerManager.observe(el, () => stickyThead(el, binding, vnode));

    // 如果绑定值是选择器，观察目标元素的变化
    if (typeof binding.value === 'string') {
      const targetElement = document.querySelector(binding.value);
      if (targetElement) {
        observerManager.observe(targetElement, () => stickyThead(el, binding, vnode));
      }
    }

    // 初始执行一次
    stickyThead(el, binding, vnode);
  },

  componentUpdated(el, binding, vnode) {
    // 组件更新后重新计算
    stickyThead(el, binding, vnode);
  },

  unbind(el, binding) {
    // 停止观察表格本身
    observerManager.unobserve(el);

    // 如果绑定值是选择器，停止观察目标元素
    if (typeof binding.value === 'string') {
      const targetElement = document.querySelector(binding.value);
      if (targetElement) {
        observerManager.unobserve(targetElement);
      }
    }
  },
};
