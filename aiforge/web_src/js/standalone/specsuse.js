window.ACC_CARD_TYPE = [{ k: 'T4', v: 'T4' }, { k: 'A100', v: 'A100' }, { k: 'V100', v: 'V100' }, { k: 'ASCEND910', v: 'Ascend 910' }, { k: 'MLU270', v: 'MLU270' }, { k: 'RTX3080', v: 'RTX3080' }, { k: 'ENFLAME-T20', v: 'ENFLAME-T20' }];

window.getListValueWithKey = (list, key, k = 'k', v = 'v', defaultV = '') => {
  for (let i = 0, iLen = list.length; i < iLen; i++) {
    const listI = list[i];
    if (listI[k] === key) return listI[v];
  }
  return defaultV || key;
};

window.renderSpecStr = (spec, showPoint, langObj) => {
  if (!spec) return '';
  var ngpu = `${spec.ComputeResource}: ${spec.AccCardsNum + '*' + getListValueWithKey(ACC_CARD_TYPE, spec.AccCardType)}`;
  var gpuMemStr = spec.GPUMemGiB != 0 ? `(${langObj.gpu_memory}: ${spec.GPUMemGiB}GB)` : '';
  var sharedMemStr = spec.ShareMemGiB != 0 ? `, ${langObj.shared_memory}: ${spec.ShareMemGiB}GB` : '';
  var pointStr = showPoint ? `, ${spec.UnitPrice == 0 ? langObj.free : spec.UnitPrice.toFixed(2) + langObj.point_hr}` : '';
  var specStr = `${ngpu}${gpuMemStr}, CPU: ${spec.CpuCores}, ${langObj.memory}: ${spec.MemGiB}GB${sharedMemStr}${pointStr}`;
  return specStr;
};

window.renderSpecsSelect = (specsSel, data, showPoint, langObj) => {
  if (data.length === 0) {
    specsSel.append(`<option>${langObj.no_use_resource}</option>`);
    specsSel.addClass('disabled red')
    return
  }
  specsSel.empty();
  data = data || [];
  showPoint = specsSel.attr('blance') ? true : false;
  var oValue = specsSel.attr('ovalue');
  for (var i = 0, iLen = data.length; i < iLen; i++) {
    var spec = data[i];
    var specStr = window.renderSpecStr(spec, showPoint, langObj);
    specsSel.append(`<option name="spec_id" value="${spec.ID}" queueCode="${spec.QueueCode}" unitprice="${spec.UnitPrice}">${specStr}</option>`);
  }
  oValue && specsSel.val(oValue);
  if (showPoint) {
    specsSel.on('change', function (e) {
      var cloudbrain_resource_spec_blance_tip_el = $('.cloudbrain_resource_spec_blance_tip');
      var blance = $(this).attr('blance');
      var unitPrice = $(this).find('option:selected').attr('unitprice');
      var work_server_number = $('#trainjob_work_server_num_select select').val() || $('#trainjob_work_server_num').val() || '1'; // 计算节点数
      if (!blance || !unitPrice) return;
      if (unitPrice == 0) {
        cloudbrain_resource_spec_blance_tip_el.find('.can-use-time').parent().hide();
      } else {
        var canUseTime = Number(blance) / (Number(unitPrice) * Number(work_server_number));
        if (Number(blance) < Number(unitPrice) * Number(work_server_number)) { // 余额不足一个单位单价时可用时间提示为 0
          canUseTime = 0;
        }
        cloudbrain_resource_spec_blance_tip_el.find('.can-use-time').text(canUseTime.toFixed(2)).parent().show();
      }
    }).trigger('change');
    $('#trainjob_work_server_num_select select').on('change', function (e) { // 计算节点数切换重新计算可用时长
      specsSel.trigger('change');
    });
  }
}
