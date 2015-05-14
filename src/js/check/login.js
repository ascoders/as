'use strict'

define("checkLogin", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "checkLogin",
		account: '',
		password: '',
		frontia: function (val) { // 点击第三方登陆
			require(['frontia'], function (frontia) {
				// API key 从应用信息页面获取
				var AK = 'RqeMWD9G1m8agmxfj6ngCKRG';
				// 在应用管理页面下的 社会化服务 - 基础设置中设置该地址
				var redirect_url = 'http://www.wokugame.com/login/oauth'

				// 初始化 frontia
				frontia.init(AK)

				// 初始化登录的配置
				var options = {
					response_type: 'token',
					media_type: val,
					redirect_uri: redirect_url,
					client_type: 'web'
				}

				// 登录
				frontia.social.login(options)
			});
		},
		submit: function () { //点击登陆按钮
			if (avalon.vmodels.checkLogin.account == '') {
				return wk.notice('账号不能为空', 'red')
			}
			if (avalon.vmodels.checkLogin.password == '') {
				return wk.notice('密码不能为空', 'red')
			}

			wk.get({
				url: '/api/users/authentication',
				data: {
					account: avalon.vmodels.checkLogin.account,
					password: avalon.vmodels.checkLogin.password
				},
				success: function (data) {
					data.image = userImage(data.image)
					avalon.vmodels.global.my = data
					avalon.vmodels.global.myLogin = true

					// 信息获取完毕
					avalon.vmodels.global.temp.myDeferred.resolve()

					// 跳回上个页面
					avalon.router.navigate(avalon.router.getLastPath())
				}
			})
		}
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			//如果已登陆，返回首页
			$.when(global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if (global.myLogin) {
					return avalon.router.navigate('/')
				}
			})
		}
		$ctrl.$onRendered = function () {
			//移动到第三方账号按钮上，下部展开
			//鼠标移动显示更多第三方登陆
			$(".other").hover(function () {
				$(".other-hide").show()
			}, function () {
				$(".other-hide").hide()
			})

			//Enter提交表单
			$('#check-login').bind('keyup', function (event) {
				if (event.keyCode == 13) { //按下Enter
					avalon.vmodels.checkLogin.submit()
				}
			})

			//账号获取焦点
			$('#check-login #account').focus()
		}
	});
});