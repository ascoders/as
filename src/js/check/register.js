'use strict';

define("checkRegister", ['jquery'], function ($) {
	var vm = avalon.define({
		$id: "checkRegister",
		email: '',
		nickname: '',
		password: '',
		passwordRepeat: '',
		capid: '', //验证码id
		cap: '',
		stepList: ['填写信息', '邮箱验证', '完成'], //步骤
		step: 0, //当前步骤
		jumpStep: function (step) {
			vm.step = step
		},
		freshCap: function () { //刷新验证码
			//刷新验证码
			wk.post({
				url: '/api/captcha',
				success: function (data) {
					vm.capid = data.captchaCode
				}
			})
		},
		submit: function () { //点击登陆按钮
			if (vm.email == '') {
				return wk.notice('邮箱不能为空', 'red')
			}
			if (vm.nickname == '') {
				return wk.notice('昵称不能为空', 'red')
			}
			if (vm.password == '') {
				return wk.notice('密码不能为空', 'red')
			}
			if (vm.passwordRepeat == '' ||
				vm.passwordRepeat != vm.password) {
				return wk.notice('重复密码不正确', 'red')
			}

			wk.post({
				url: '/api/users/authentication',
				data: {
					email: vm.email,
					nickname: vm.nickname,
					password: vm.password,
					capid: vm.capid,
					cap: vm.cap
				},
				success: function (data) {
					//刷新验证码
					vm.freshCap()
					vm.cap = ''

					//进入下一步
					vm.step = 1
				},
				error: function () {
					//刷新验证码
					vm.freshCap()
					vm.cap = ''
				}
			})
		}
	});
	return avalon.controller(function ($ctrl) {
		$ctrl.$onEnter = function (param, rs, rj) {
			//如果已登陆，返回首页
			$.when(global.temp.myDeferred).done(function () { // 此时获取用户信息完毕
				if (global.myLogin) {
					avalon.router.navigate('/')
					return
				}
			})

			vm.freshCap()
		}
		$ctrl.$onRendered = function () {
			//Enter提交表单
			$('#check-register').bind('keyup', function (event) {
				if (event.keyCode == 13) { //按下Enter
					vm.submit()
				}
			})

			//账号获取焦点
			$('#check-register #email').focus()
		}
	})
})