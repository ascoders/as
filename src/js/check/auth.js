'use strict';

define("auth", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "auth",
		todo: function (type, extend) {
			switch (type) {
			case 'email': // 修改用户email
				avalon.router.navigate("/user/account/email");
				break;
			case 'createAccount': // 注册用户
				avalon.router.navigate("/");
				break;
			}
		}
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			post('/api/check/auth', {
				id: mmState.query.id,
				expire: mmState.query.expire,
				type: mmState.query.type,
				extend: mmState.query.extend,
				sign: mmState.query.sign,
			}, null, '', function (data) {
				//更新用户信息
				data.image = userImage(data.image);
				avalon.vmodels.global.my = data;
				avalon.vmodels.global.myLogin = true;
				vm.todo(mmState.query.type, mmState.query.extend);
			}, function () { // 跳转回首页
				avalon.router.navigate("/");
			});
		}
		$ctrl.$onRendered = function () {

		}
	});
});