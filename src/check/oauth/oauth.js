'use strict';

define("oauth", ['jquery', 'frontia'], function ($, frontia) {
	var vm = avalon.define({
		$id: "oauth"
	});

	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			// API key 从应用信息页面获取
			var AK = 'RqeMWD9G1m8agmxfj6ngCKRG';

			// 初始化 frontia
			frontia.init(AK);

			frontia.social.setLoginCallback({ //登录成功后的回调
				success: function (user) {
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
						post('/api/user/oauth', {
							id: user.getId(),
							token: user.getAccessToken(),
							nickname: result.username,
							image: result.headurl,
							type: user.getMediaType(),
							expire: user.getExpiresIn()
						}, '验证成功', '', function () { // 绑定成功，跳转到绑定页面
							avalon.router.navigate('/user/account/oauth');
						});
					})
				},
				error: function (error) {
					notice('验证失败', 'red');
					avalon.router.navigate('/');
				}
			});
		}

		$ctrl.$onRendered = function () {

		}
	});
});