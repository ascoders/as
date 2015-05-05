'use strict';

define("userAccountRecharge", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "userAccountRecharge",
		pay: 1, // 充值金额
		postForm: '', // 提交表单
		submit: function () {
			var _this = this;
			$(_this).addClass('disabled').text('处理中..');
			post('/api/user/recharge', {
				account: avalon.vmodels.global.my.nickName,
				number: vm.pay,
				plantform: 'alipay',
				type: 'web'
			}, null, '', function (data) {
				vm.postForm = data;
			}, function () {
				$(_this).removeClass('disabled').text('确认充值');
			});
		},
		changePay: function (number) { // 改变支付金额
			vm.pay = number;
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {

		}
	});
});