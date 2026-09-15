(function () {
  var VERT = [
    'attribute vec2 uv;',
    'attribute vec2 position;',
    'varying vec2 vUv;',
    'void main() {',
    '  vUv = uv;',
    '  gl_Position = vec4(position, 0, 1);',
    '}'
  ].join('\n');

  var FRAG = [
    'precision highp float;',
    'uniform float uTime;',
    'uniform vec3 uResolution;',
    'uniform vec2 uFocal;',
    'uniform vec2 uRotation;',
    'uniform float uStarSpeed;',
    'uniform float uDensity;',
    'uniform float uHueShift;',
    'uniform float uSpeed;',
    'uniform vec2 uMouse;',
    'uniform float uGlowIntensity;',
    'uniform float uSaturation;',
    'uniform float uMouseRepulsion;',
    'uniform float uTwinkleIntensity;',
    'uniform float uRotationSpeed;',
    'uniform float uRepulsionStrength;',
    'uniform float uMouseActiveFactor;',
    'uniform float uAutoCenterRepulsion;',
    'uniform float uTransparent;',
    'uniform float uLightMode;',
    'varying vec2 vUv;',
    '#define NUM_LAYER 4.0',
    '#define STAR_COLOR_CUTOFF 0.2',
    '#define MAT45 mat2(0.7071, -0.7071, 0.7071, 0.7071)',
    '#define PERIOD 3.0',
    'float Hash21(vec2 p) {',
    '  p = fract(p * vec2(123.34, 456.21));',
    '  p += dot(p, p + 45.32);',
    '  return fract(p.x * p.y);',
    '}',
    'float tri(float x) {',
    '  return abs(fract(x) * 2.0 - 1.0);',
    '}',
    'float tris(float x) {',
    '  float t = fract(x);',
    '  return 1.0 - smoothstep(0.0, 1.0, abs(2.0 * t - 1.0));',
    '}',
    'float trisn(float x) {',
    '  float t = fract(x);',
    '  return 2.0 * (1.0 - smoothstep(0.0, 1.0, abs(2.0 * t - 1.0))) - 1.0;',
    '}',
    'vec3 hsv2rgb(vec3 c) {',
    '  vec4 K = vec4(1.0, 2.0 / 3.0, 1.0 / 3.0, 3.0);',
    '  vec3 p = abs(fract(c.xxx + K.xyz) * 6.0 - K.www);',
    '  return c.z * mix(K.xxx, clamp(p - K.xxx, 0.0, 1.0), c.y);',
    '}',
    'float Star(vec2 uv, float flare) {',
    '  float d = length(uv);',
    '  float m = (0.065 * uGlowIntensity) / d;',
    '  float rays = smoothstep(0.0, 1.0, 1.0 - abs(uv.x * uv.y * 1000.0));',
    '  m += rays * flare * uGlowIntensity * 0.22;',
    '  uv *= MAT45;',
    '  rays = smoothstep(0.0, 1.0, 1.0 - abs(uv.x * uv.y * 1000.0));',
    '  m += rays * 0.08 * flare * uGlowIntensity;',
    '  m *= smoothstep(1.0, 0.2, d);',
    '  return m;',
    '}',
    'vec3 StarLayer(vec2 uv) {',
    '  vec3 col = vec3(0.0);',
    '  vec2 gv = fract(uv) - 0.5;',
    '  vec2 id = floor(uv);',
    '  for (int y = -1; y <= 1; y++) {',
    '    for (int x = -1; x <= 1; x++) {',
    '      vec2 offset = vec2(float(x), float(y));',
    '      vec2 si = id + vec2(float(x), float(y));',
    '      float seed = Hash21(si);',
    '      float size = fract(seed * 345.32);',
    '      float glossLocal = tri(uStarSpeed / (PERIOD * seed + 1.0));',
    '      float flareSize = smoothstep(0.9, 1.0, size) * glossLocal;',
    '      float red = smoothstep(STAR_COLOR_CUTOFF, 1.0, Hash21(si + 1.0)) + STAR_COLOR_CUTOFF;',
    '      float blu = smoothstep(STAR_COLOR_CUTOFF, 1.0, Hash21(si + 3.0)) + STAR_COLOR_CUTOFF;',
    '      float grn = min(red, blu) * seed;',
    '      vec3 base = vec3(red, grn, blu);',
    '      float hue = atan(base.g - base.r, base.b - base.r) / (2.0 * 3.14159) + 0.5;',
    '      hue = fract(hue + uHueShift / 360.0);',
    '      float sat = length(base - vec3(dot(base, vec3(0.299, 0.587, 0.114)))) * uSaturation;',
    '      float val = max(max(base.r, base.g), base.b);',
    '      base = hsv2rgb(vec3(hue, sat, val));',
    '      vec2 pad = vec2(tris(seed * 34.0 + uTime * uSpeed / 10.0), tris(seed * 38.0 + uTime * uSpeed / 30.0)) - 0.5;',
    '      float star = Star(gv - offset - pad, flareSize);',
    '      float twinkle = trisn(uTime * uSpeed + seed * 6.2831) * 0.5 + 1.0;',
    '      twinkle = mix(1.0, twinkle, uTwinkleIntensity);',
    '      star *= twinkle;',
    '      col += star * size * base;',
    '    }',
    '  }',
    '  return col;',
    '}',
    'void main() {',
    '  vec2 focalPx = uFocal * uResolution.xy;',
    '  vec2 uv = (vUv * uResolution.xy - focalPx) / uResolution.y;',
    '  vec2 mouseNorm = uMouse - vec2(0.5);',
    '  if (uAutoCenterRepulsion > 0.0) {',
    '    vec2 centerUV = vec2(0.0, 0.0);',
    '    float centerDist = length(uv - centerUV);',
    '    vec2 repulsion = normalize(uv - centerUV) * (uAutoCenterRepulsion / (centerDist + 0.1));',
    '    uv += repulsion * 0.05;',
    '  } else if (uMouseRepulsion > 0.5) {',
    '    vec2 mousePosUV = (uMouse * uResolution.xy - focalPx) / uResolution.y;',
    '    float mouseDist = length(uv - mousePosUV);',
    '    vec2 repulsion = normalize(uv - mousePosUV) * (uRepulsionStrength / (mouseDist + 0.1));',
    '    uv += repulsion * 0.05 * uMouseActiveFactor;',
    '  } else {',
    '    vec2 mouseOffset = mouseNorm * 0.1 * uMouseActiveFactor;',
    '    uv += mouseOffset;',
    '  }',
    '  float autoRotAngle = uTime * uRotationSpeed;',
    '  mat2 autoRot = mat2(cos(autoRotAngle), -sin(autoRotAngle), sin(autoRotAngle), cos(autoRotAngle));',
    '  uv = autoRot * uv;',
    '  uv = mat2(uRotation.x, -uRotation.y, uRotation.y, uRotation.x) * uv;',
    '  vec3 col = vec3(0.0);',
    '  for (float i = 0.0; i < 1.0; i += 1.0 / NUM_LAYER) {',
    '    float depth = fract(i + uStarSpeed * uSpeed);',
    '    float scale = mix(20.0 * uDensity, 0.5 * uDensity, depth);',
    '    float fade = depth * smoothstep(1.0, 0.9, depth);',
    '    col += StarLayer(uv * scale + i * 453.32) * fade;',
    '  }',
    '  if (uLightMode > 0.5) {',
    '    vec3 paper = vec3(0.82, 0.90, 1.0);',
    '    vec3 tint = vec3(0.22, 0.55, 1.0);',
    '    vec3 glow = mix(tint, vec3(0.85, 0.93, 1.0), 0.18) * col;',
    '    gl_FragColor = vec4(paper + glow * 1.15, 1.0);',
    '  } else if (uTransparent > 0.5) {',
    '    float alpha = length(col);',
    '    alpha = smoothstep(0.0, 0.3, alpha);',
    '    alpha = min(alpha, 1.0);',
    '    gl_FragColor = vec4(col, alpha);',
    '  } else {',
    '    gl_FragColor = vec4(col, 1.0);',
    '  }',
    '}'
  ].join('\n');

  function compile(gl, type, src) {
    var sh = gl.createShader(type);
    gl.shaderSource(sh, src);
    gl.compileShader(sh);
    if (!gl.getShaderParameter(sh, gl.COMPILE_STATUS)) {
      console.warn('[galaxy-bg] shader', gl.getShaderInfoLog(sh));
      gl.deleteShader(sh);
      return null;
    }
    return sh;
  }

  function init() {
    var ctn = document.getElementById('galaxy-bg');
    if (!ctn) return;

    document.documentElement.classList.add('has-galaxy-login');

    var canvas = document.createElement('canvas');
    canvas.setAttribute('aria-hidden', 'true');
    ctn.appendChild(canvas);

    var gl = canvas.getContext('webgl', {
      alpha: false,
      antialias: false,
      premultipliedAlpha: false
    });
    if (!gl) {
      ctn.parentNode && ctn.parentNode.removeChild(ctn);
      document.documentElement.classList.remove('has-galaxy-login');
      return;
    }

    var vs = compile(gl, gl.VERTEX_SHADER, VERT);
    var fs = compile(gl, gl.FRAGMENT_SHADER, FRAG);
    if (!vs || !fs) return;

    var program = gl.createProgram();
    gl.attachShader(program, vs);
    gl.attachShader(program, fs);
    gl.bindAttribLocation(program, 0, 'position');
    gl.bindAttribLocation(program, 1, 'uv');
    gl.linkProgram(program);
    if (!gl.getProgramParameter(program, gl.LINK_STATUS)) {
      console.warn('[galaxy-bg] program', gl.getProgramInfoLog(program));
      return;
    }
    gl.useProgram(program);

    var pos = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER, pos);
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array([-1, -1, 3, -1, -1, 3]), gl.STATIC_DRAW);
    gl.enableVertexAttribArray(0);
    gl.vertexAttribPointer(0, 2, gl.FLOAT, false, 0, 0);

    var uv = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER, uv);
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array([0, 0, 2, 0, 0, 2]), gl.STATIC_DRAW);
    gl.enableVertexAttribArray(1);
    gl.vertexAttribPointer(1, 2, gl.FLOAT, false, 0, 0);

    var loc = {};
    [
      'uTime', 'uResolution', 'uFocal', 'uRotation', 'uStarSpeed', 'uDensity',
      'uHueShift', 'uSpeed', 'uMouse', 'uGlowIntensity', 'uSaturation',
      'uMouseRepulsion', 'uTwinkleIntensity', 'uRotationSpeed', 'uRepulsionStrength',
      'uMouseActiveFactor', 'uAutoCenterRepulsion', 'uTransparent', 'uLightMode'
    ].forEach(function (name) {
      loc[name] = gl.getUniformLocation(program, name);
    });

    var cfg = {
      starSpeed: 0.45,
      density: 1.05,
      hueShift: 210,
      speed: 0.85,
      glowIntensity: 0.4,
      saturation: 0.55,
      twinkleIntensity: 0.28,
      rotationSpeed: 0.06,
      repulsionStrength: 2,
      lightMode: 1
    };

    gl.uniform2f(loc.uFocal, 0.5, 0.5);
    gl.uniform2f(loc.uRotation, 1, 0);
    gl.uniform1f(loc.uDensity, cfg.density);
    gl.uniform1f(loc.uHueShift, cfg.hueShift);
    gl.uniform1f(loc.uSpeed, cfg.speed);
    gl.uniform1f(loc.uGlowIntensity, cfg.glowIntensity);
    gl.uniform1f(loc.uSaturation, cfg.saturation);
    gl.uniform1f(loc.uMouseRepulsion, 1);
    gl.uniform1f(loc.uTwinkleIntensity, cfg.twinkleIntensity);
    gl.uniform1f(loc.uRotationSpeed, cfg.rotationSpeed);
    gl.uniform1f(loc.uRepulsionStrength, cfg.repulsionStrength);
    gl.uniform1f(loc.uAutoCenterRepulsion, 0);
    gl.uniform1f(loc.uTransparent, 0);
    gl.uniform1f(loc.uLightMode, cfg.lightMode);
    gl.uniform2f(loc.uMouse, 0.5, 0.5);
    gl.uniform1f(loc.uMouseActiveFactor, 0);

    var reduce = window.matchMedia && window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    var mouse = { x: 0.5, y: 0.5, tx: 0.5, ty: 0.5, active: 0, tactive: 0 };

    function resize() {
      var dpr = Math.min(window.devicePixelRatio || 1, 1.5);
      var w = Math.max(1, Math.floor(ctn.clientWidth * dpr));
      var h = Math.max(1, Math.floor(ctn.clientHeight * dpr));
      if (canvas.width !== w || canvas.height !== h) {
        canvas.width = w;
        canvas.height = h;
      }
      gl.viewport(0, 0, w, h);
      gl.uniform3f(loc.uResolution, w, h, w / h);
    }

    function onMove(e) {
      var rect = ctn.getBoundingClientRect();
      if (!rect.width || !rect.height) return;
      mouse.tx = (e.clientX - rect.left) / rect.width;
      mouse.ty = 1 - (e.clientY - rect.top) / rect.height;
      mouse.tactive = 1;
    }

    function onLeave() {
      mouse.tactive = 0;
    }

    window.addEventListener('mousemove', onMove, { passive: true });
    window.addEventListener('mouseleave', onLeave);
    window.addEventListener('resize', resize);
    resize();

    var start = performance.now();
    var raf = 0;

    function frame(now) {
      raf = requestAnimationFrame(frame);
      var t = (now - start) * 0.001;
      if (!reduce) {
        gl.uniform1f(loc.uTime, t);
        gl.uniform1f(loc.uStarSpeed, (t * cfg.starSpeed) / 10);
      } else {
        gl.uniform1f(loc.uTime, 0);
        gl.uniform1f(loc.uStarSpeed, 0);
      }
      mouse.x += (mouse.tx - mouse.x) * 0.05;
      mouse.y += (mouse.ty - mouse.y) * 0.05;
      mouse.active += (mouse.tactive - mouse.active) * 0.05;
      gl.uniform2f(loc.uMouse, mouse.x, mouse.y);
      gl.uniform1f(loc.uMouseActiveFactor, mouse.active);
      gl.clearColor(0.82, 0.90, 1.0, 1);
      gl.clear(gl.COLOR_BUFFER_BIT);
      gl.drawArrays(gl.TRIANGLES, 0, 3);
    }

    raf = requestAnimationFrame(frame);

    window.addEventListener('pagehide', function () {
      cancelAnimationFrame(raf);
      window.removeEventListener('mousemove', onMove);
      window.removeEventListener('mouseleave', onLeave);
      window.removeEventListener('resize', resize);
    }, { once: true });
  }

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
