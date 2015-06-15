'use strict';

define("userAccountEmail", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "userAccountEmail",
		email: '',
		nowEmail: '',
		submit: function () { // 发送验证邮件
			if (vm.email == '') {
				notice('请填写邮箱', 'red');
				return;
			}

			var _this = this;
			$(_this).addClass('disabled').text('邮件发送中..');

			post('/api/user/email', {
				email: vm.email
			}, '已发送验证邮件', '', function () {
				$(_this).removeClass('disabled').text('发送验证邮件');
			}, function () {
				$(_this).removeClass('disabled').text('发送验证邮件');
			});
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			$.when(avalon.vmodels.global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if (avalon.vmodels.global.my.email != '') {
					vm.nowEmail = '绑定邮箱：' + avalon.vmodels.global.my.email;
				} else {
					vm.nowEmail = '暂无绑定邮箱';
				}
			});
		}

		$ctrl.$onRendered = function () {

		}
	});
});