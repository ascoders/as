'use strict';

define("checkOauth", ['jquery', 'frontia'], function ($, frontia) {
	var vm = avalon.define({
		$id: "checkOauth",
		showForm: false,
		img: '',
		nickname: '',
		$user: {
			id: null,
			token: null,
			type: null,
			expire: null
		},
		submit: function () {
			if (vm.nickname.length < 3 || vm.nickname.length > 20) {
				notice('昵称长度为3-20', 'red');
				return false;
			}

			post('/api/check/oauthRegister', {
				id: vm.$user.id,
				token: vm.$user.token,
				nickname: vm.nickname,
				image: vm.img,
				type: vm.$user.type,
				expire: vm.$user.expire
			}, '注册成功', '', function (data) {
				data.image = userImage(data.image);
				avalon.vmodels.global.my = data;
				avalon.vmodels.global.myLogin = true;

				// 跳回上个页面
				avalon.router.navigate(avalon.router.getLastPath());
			});
		}
	});
	return avalon.controller(function ($ctrl) {
		//如果已登陆，返回首页
		$.when(global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
			if (global.myLogin) {
				avalon.router.navigate('/');
				return;
			}
		});

		$ctrl.$onEnter = function (param, rs, rj) {
			// API key 从应用信息页面获取
			var AK = 'RqeMWD9G1m8agmxfj6ngCKRG';

			// 初始化 frontia
			frontia.init(AK);

			// 设置登录成功后的回调
			frontia.social.setLoginCallback({
				success: function (user) {
					vm.$user.id = user.getId()
					vm.$user.token = user.getAccessToken()
					vm.$user.type = user.getMediaType()
					vm.$user.expire = user.getExpiresIn()

					post('/api/check/hasOauth', {
						id: user.getId(),
						token: user.getAccessToken(),
						type: user.getMediaType(),
						expire: user.getExpiresIn()
					}, null, '', function (data) {
						if (data == -1) { // 用户不存在
							vm.showForm = true;

							//获取用户详细信息
							$.ajax({
								url: 'https://openapi.baidu.com/social/api/2.0/user/info',
								type: 'POST',
								dataType: 'jsonp',
								jsonp: "callback",
								data: {
									access_token: user.getAccessToken()
								},
							}).done(function (result) {
								vm.img = result.headurl
								vm.nickname = result.username
							});
						} else { //账号已存在，自动登陆
							notice('登陆成功', 'green');

							data.image = userImage(data.image);
							avalon.vmodels.global.my = data;
							avalon.vmodels.global.myLogin = true;
							// 跳回首页
							avalon.router.navigate('/');
						}
					});
				},
				error: function (error) {
					showAlert("验证失败", false);
					location.href = "/";
				}
			});
		}
		$ctrl.$onRendered = function () {

		}
	});
});