'use strict';

define("userAccountHistory", ['jquery', 'jquery.timeago'], function ($) {
	var vm = avalon.define({
		$id: "userAccountHistory",
		lists: [], // 消息列表
		pagin: '',
		payForm: '',
		rendered: function () { // 列表渲染完毕
			// timeago 插件
			$('#user-account-history .timeago').timeago();

			//倒计时
			$('#user-account-history .timediff').each(function (index) {
				var _this = $(this);
				var time = $(this).attr('time');
				if (!time) {
					return;
				}
				$(this).removeAttr('time');
				timediff($(this), {
					second: time
				}, function () {
					_this.parent().text('');

					vm.lists[index].Status = '<span class="text-muted"><i class="fa fa-times f-mr5"></i>已失效</span>';
				});
			});
		},
		recharge: function (id) { // 继续付款
			console.log(id);
			post('/api/user/rechargeOrder', {
				orderid: id
			}, null, '', function (data) {
				vm.payForm = data;
			});
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			mmState.query.from = mmState.query.from || 0;
			mmState.query.number = mmState.query.number || 20;

			// 请求获取分页信息
			post('/api/user/getHistory', {
				from: mmState.query.from,
				number: mmState.query.number
			}, null, '获取充值记录失败：', function (data) {
				for (var key in data.lists) {
					// 订单状态
					if (data.lists[key].Success) {
						data.lists[key].Status = '<span class="text-success"><i class="fa fa-check f-mr5"></i>已完成</span>';

						// 支付方式
						switch (data.lists[key].PayPlantform) {
						case 'alipay':
							data.lists[key].PayWay = '<img src="/static/img/user/alipay.png" class="payIcon"><span class="payIconText">' + data.lists[key].AlipayNumber + '</span>';
							break;
						}
					} else { // 待支付或者已过期（订单过期时间超过2小时）
						var timeUnix = (Date.parse(new Date()) - Date.parse(data.lists[key].Time)) / 1000; //距离现在时间戳（秒）
						if (timeUnix > 7200) {
							data.lists[key].Status = '<span class="text-muted"><i class="fa fa-times f-mr5"></i>已失效</span>';
						} else {
							data.lists[key].Status = '<a ms-click="recharge(\'' + data.lists[key].Id + '\')" class="text-warning"><i class="fa fa-clock-o f-mr5"></i>待付款</a>';

							data.lists[key].PayWay = '剩余 <span class="timediff" time="' + (7200 - timeUnix) + '"><span class="f-mr5" id="j-hour"></span><span class="f-mr5" id="j-minute"></span><span id="j-second"></span></span>';
						}
					}
				}

				vm.lists = data.lists || [];

				// 生成分页
				vm.pagin = createPagin(mmState.query.from, mmState.query.number, data.count);
			});
		}

		$ctrl.$onRendered = function () {

		}
	});
});