'use strict';

define("userAccountMessage", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "userAccountMessage",
		lists: [], //消息列表
		pagin: ''
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			mmState.query.from = mmState.query.from || 0;
			mmState.query.number = mmState.query.number || 20;

			//请求获取分页信息
			post('/api/user/getMessage', {
				from: mmState.query.from,
				number: mmState.query.number
			}, null, '获取消息失败：', function (data) {
				for (var key in data.lists) {
					switch (data.lists[key].Category) {
					case 'bindOauth':
						data.lists[key].Title = '您的账号使用邮箱注册，可以绑定第三方平台账号';
						data.lists[key].Content = '绑定第三方账号，可以方便您使用一键登陆，忘记密码时也可以快速登录并找回密码！';
						break;
					case 'updateImage':
						data.lists[key].Title = '您的修改了头像';
						data.lists[key].Content = '您的头像修改为<img class="img" src="' + data.lists[key].Info + '">';
						break;
					}
				}

				vm.lists = data.lists || [];

				//生成分页
				vm.pagin = createPagin(mmState.query.from, mmState.query.number, data.count);
			});
		}

		$ctrl.$onRendered = function () {

		}
	});
});