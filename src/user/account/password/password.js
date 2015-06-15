'use strict';

define("userAccountPassword", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "userAccountPassword",
		password: '',
		rpassword: '',
		submit: function () {
			if (avm.password == "") {
				notice('密码不能为空', 'red');
				return;
			}
			if (avm.rpassword == "") {
				notice('确认密码不能为空', 'red');
				return;
			}
			if (avm.password != avm.rpassword) {
				notice('两次输入不一致', 'red');
				return;
			}

			post('/api/user/password', {
				password: avm.password
			}, '修改成功', '');
		}
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {

		}

		$ctrl.$onRendered = function () {

		}
	});
});